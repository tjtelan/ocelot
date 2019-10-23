extern crate structopt;
use structopt::StructOpt;

use config_parser::parser;

use crate::{GlobalOption, SubcommandError};

#[derive(Debug, StructOpt)]
#[structopt(rename_all = "kebab_case")]
pub struct SubcommandOption {
    /// Path to local repo. Defaults to current working directory
    #[structopt(long)]
    path: Option<String>,
}

// TODO: Do we want to return the config?
pub fn subcommand_handler(
    _global_option: GlobalOption,
    local_option: SubcommandOption,
) -> Result<(), SubcommandError> {
    if let Some(path) = local_option.path {
        match parser::load_orb_yaml(path) {
            Ok(c) => Ok(()),
            Err(e) => Err(SubcommandError::new("Config file failed validation")),
        }

    }
    else {
        Err(SubcommandError::new("No config file specified"))
    }
}

