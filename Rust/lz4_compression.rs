//module compress::lzl4 from Rust documentation
//Decompression and Compression. Requires lz4 feature, enabled by default
use compress::lz4;
use std::fs::File;
use std::path::Path;
use std::io::Read;
fn main()
{
	
	let stream = File::open(&Path::new("/home/andrew/Desktop/plain.lz4")).unwrap();
let mut decompressed = Vec::new();

lz4::Decoder::new(stream).read_to_end(&mut decompressed);
	
	
}

