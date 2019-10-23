use std::io::Read;
use serde::{Serialize, Deserialize};

#[derive(Debug, PartialEq, Serialize, Deserialize)]
struct OrbitalConfig {
    image: String,
    command: Vec<String>,
}

pub fn load_orb_yaml(path : String) -> Result<(), Box<dyn std::error::Error>> {

    let mut f = std::fs::File::open(path)?;
    let mut file_data = String::new();
    f.read_to_string(&mut file_data)?;

    let deserialized : OrbitalConfig = serde_yaml::from_str(&file_data)?;

    println!("{:?}", deserialized);

    // TODO: Send yaml_data through a parser to validate the config
    Ok(())
}
