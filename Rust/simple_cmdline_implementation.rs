//very simple command line impelementation from the Rust documentation
use std::env::args;
fn main() {
  
  let pattern = std::env::args().nth(1).expect("pattern provided");
  let path = std::env::args().nth(2).expect("Path not given");
 
 struct Cli {
    pattern: String,
    path: std::path::PathBuf,
};

}
