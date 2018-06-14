extern crate walkdir;

use std::env;
use std::io::Error;
use std::path::Path;

use walkdir::WalkDir;

fn walk<P: AsRef<Path>>(path: P) -> Result<(), Error> {
  let mut count = 0;
  for entry in WalkDir::new(path) {
    let entry = entry?;
    if entry.file_type().is_dir() {
      continue;
    }
    let name = entry.file_name();
    if let Some(s) = name.to_str() {
      if s.to_ascii_lowercase().ends_with(".jpg") {
        count += 1;
      }
    }
  }
  println!("{}", count);
  Ok(())
}

fn main() -> Result<(), Error> {
  walk(env::args().nth(1).unwrap())
}
