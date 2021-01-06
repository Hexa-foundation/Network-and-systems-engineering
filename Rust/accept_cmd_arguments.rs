//Gotten from Stack overflow & error prone
use std::process::Command;
fn main()
{
	let cmd = Command::new("ls")
	           .arg("-l")
	           .arg("-a")
	           .spawn()
	           .expect("Failed to start");
	
println!("status: {}", cmd.status);
println!("stdout: {}", String::from_utf8_lossy(&cmd.stdout));
println!("stderr: {}", String::from_utf8_lossy(&cmd.stderr));

assert!(output.status.success());

}
