extern crate structopt;
use std::str::FromStr;
use structopt::StructOpt;

use container_builder::docker;

use crate::{GlobalOption, SubcommandError};

#[derive(Debug, StructOpt)]
#[structopt(rename_all = "kebab_case")]
pub struct SubcommandOption {
    /// Docker image
    #[structopt(long)]
    image: Option<String>,

    /// Pull, Create
    action: Action,
}

#[derive(Debug, StructOpt)]
#[structopt(rename_all = "kebab_case")]
pub enum Action {
    Pull,
    Create,
}

impl FromStr for Action {
    type Err = String;
    fn from_str(action: &str) -> Result<Self, Self::Err> {
        match action {
            "pull" => Ok(Action::Pull),
            "create" => Ok(Action::Create),
            _ => Err("Invalid action".to_string()),
        }
    }
}

pub fn subcommand_handler(
    _global_option: GlobalOption,
    local_option: SubcommandOption,
) -> Result<(), SubcommandError> {
    //container_pull(local_option.image);

    match local_option.action {
        Action::Pull => {
            docker::container_pull(local_option.image);
        }
        Action::Create => println!("Placeholder. Create container."),
    }
    Ok(())
}
