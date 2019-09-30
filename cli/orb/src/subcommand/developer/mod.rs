use structopt::StructOpt;

use crate::{GlobalOption, SubcommandError};

pub mod git;

#[derive(Debug, StructOpt)]
#[structopt(rename_all = "kebab_case")]
pub enum DeveloperType {
    Git(git::SubcommandOption)
}


pub fn subcommand_handler(_global_option : GlobalOption, dev_subcommand : DeveloperType) -> Result<(), SubcommandError> {
    Err(SubcommandError::new("Not yet implemented"))
}