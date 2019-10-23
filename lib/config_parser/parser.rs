use serde::{Deserialize, Serialize};

#[derive(Debug, PartialEq, Serialize, Deserialize)]
pub struct OrbitalConfig {
    image: String,
    command: Vec<String>,
}

pub fn load_orb_yaml(path: String) -> Result<OrbitalConfig, Box<dyn std::error::Error>> {
    let f = std::fs::File::open(path)?;
    let parsed: OrbitalConfig = serde_yaml::from_reader(&f)?;

    println!("{:?}", parsed);

    Ok(parsed)
}
