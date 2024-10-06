use serde::{Deserialize, Serialize};
use ssh2::Session;
use std::fmt::{Debug, Formatter};
use std::io::Read;
use std::net::TcpStream;
use warp::reject::Reject;
use warp::Filter;

#[derive(Deserialize)]
struct SshRequest {
    ip: String,
    username: String,
    password: String,
}

#[derive(Serialize)]
struct SshResponse {
    packages: Vec<String>,
}

#[derive(Serialize)]
struct ErrorResponse {
    error: String,
}

impl Debug for ErrorResponse {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        todo!()
    }
}

// Implement the Reject trait for ErrorResponse
impl Reject for ErrorResponse {}

async fn ssh_command_handler(req: SshRequest) -> Result<impl warp::Reply, warp::Rejection> {
    // Create a new SSH session
    let tcp = TcpStream::connect(format!("{}:1024", req.ip)).map_err(|_| {
        warp::reject::custom(ErrorResponse {
            error: "SSH login failed".to_string(),
        })
    })?;

    let mut session = Session::new().unwrap();
    session.set_tcp_stream(tcp);
    session.handshake().map_err(|_| {
        warp::reject::custom(ErrorResponse {
            error: "SSH handshake failed".to_string(),
        })
    })?;

    session
        .userauth_password(&req.username, &req.password)
        .map_err(|_| {
            warp::reject::custom(ErrorResponse {
                error: "SSH authentication failed".to_string(),
            })
        })?;

    // Execute the command
    let mut channel = session.channel_session().map_err(|_| {
        warp::reject::custom(ErrorResponse {
            error: "Failed to open channel".to_string(),
        })
    })?;

    channel.exec("dpkg -l | grep ii").map_err(|_| {
        warp::reject::custom(ErrorResponse {
            error: "Failed to execute command".to_string(),
        })
    })?;

    let mut output = String::new();
    channel.read_to_string(&mut output).map_err(|_| {
        warp::reject::custom(ErrorResponse {
            error: "Failed to read command output".to_string(),
        })
    })?;

    channel.send_eof().map_err(|_| {
        warp::reject::custom(ErrorResponse {
            error: "Failed to send EOF".to_string(),
        })
    })?;
    channel.wait_close().map_err(|_| {
        warp::reject::custom(ErrorResponse {
            error: "Failed to wait for channel close".to_string(),
        })
    })?;

    // Split output into lines and filter out empty lines
    let packages: Vec<String> = output
        .lines()
        .filter_map(|line| {
            let trimmed = line.trim();
            let package_name = trimmed.split_whitespace().next().unwrap_or("").to_string();
            if !package_name.is_empty() {
                Some(package_name)
            } else {
                None
            }
        })
        .collect();
    Ok(warp::reply::json(&SshResponse { packages }))
}

#[tokio::main]
async fn main() {
    // Define the POST route
    let ssh_route = warp::path!("api" / "ssh")
        .and(warp::post())
        .and(warp::body::json()) // Expect JSON body
        .and_then(ssh_command_handler);

    // Start the warp server
    warp::serve(ssh_route).run(([127, 0, 0, 1], 3000)).await;
}
