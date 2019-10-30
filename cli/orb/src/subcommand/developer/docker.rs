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

    /// command
    #[structopt(long)]
    command: Option<String>,

    /// Pull, Create
    action: Action,
}

#[derive(Debug, StructOpt)]
#[structopt(rename_all = "kebab_case")]
pub enum Action {
    Pull,
    Create,
    Start,
    Stop,
    Exec,
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
    match local_option.action {
        Action::Pull => {
            match docker::container_pull(
                &local_option.image.clone().expect("No image provided")[..],
            ) {
                Ok(_) => return Ok(()),
                Err(_) => {
                    return Err(SubcommandError::new(&format!(
                        "Could not pull image {:?}",
                        &local_option.image
                    )))
                }
            };
        }
        Action::Create => {
            let unwrapped_command = local_option.command.clone().expect("No command provided");

            // FIXME
            // This is going to be a stupid parsed command on whitespace only.
            // Embedded commands with $() or backtics not expected to work with this parsing
            let command_vec_slice: Vec<&str> = unwrapped_command.split_whitespace().collect();

            match docker::container_create(
                &local_option.image.clone().expect("No image provided")[..],
                command_vec_slice,
                None,
                None,
            ) {
                Ok(container_id) => {
                    println!("{}", container_id);
                    return Ok(());
                }
                Err(_) => {
                    return Err(SubcommandError::new(&format!(
                        "Could not pull image {:?}",
                        &local_option.image
                    )))
                }
            };
        }

        Action::Start => {
            //debug!("Placeholder. Start container");
            Ok(())
        }
        Action::Stop => {
            //debug!("Placeholder. Stop container");
            Ok(())
        }
        Action::Exec => {
            //debug!("Placeholder. Exec container");
            Ok(())
        }
    }
}
