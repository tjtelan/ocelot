use structopt::StructOpt;

use crate::{GlobalOption, SubcommandError};

pub mod git;

#[derive(Debug, StructOpt)]
#[structopt(rename_all = "kebab_case")]
pub enum DeveloperType {
    Git(git::SubcommandOption),
}

pub fn subcommand_handler(
    global_option: GlobalOption,
    dev_subcommand: DeveloperType,
) -> Result<(), SubcommandError> {
    match dev_subcommand {
        DeveloperType::Git(sub_option) => git::subcommand_handler(global_option, sub_option),
    }
}
