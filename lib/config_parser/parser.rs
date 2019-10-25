use serde::{Deserialize, Serialize};

#[derive(Debug, PartialEq, Serialize, Deserialize)]
pub struct OrbitalConfig {
    pub image: String,
    pub command: Vec<String>,
}

pub fn load_orb_yaml(path: String) -> Result<OrbitalConfig, Box<dyn std::error::Error>> {
    let f = std::fs::File::open(path)?;
    let parsed: OrbitalConfig = serde_yaml::from_reader(&f)?;

    // TODO: Place this behind a debug flag
    //println!("{:?}", parsed);

    Ok(parsed)
}
