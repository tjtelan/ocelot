extern crate structopt;
use structopt::StructOpt;

use git_meta::git_info;
use config_parser::parser;

use crate::{GlobalOption, SubcommandError};

#[derive(Debug, StructOpt)]
#[structopt(rename_all = "kebab_case")]
pub struct SubcommandOption {
    /// Path to local repo. Defaults to current working directory
    // TODO: Change default value to PWD env var
    #[structopt(long, default_value = ".")]
    path: String,

    /// Use the specified local branch
    #[structopt(long)]
    branch: Option<String>,

    /// Use the specified commit hash
    #[structopt(long)]
    hash: Option<String>,
}

pub fn subcommand_handler(
    _global_option: GlobalOption,
    local_option: SubcommandOption,
) -> Result<(), SubcommandError> {
    // Read options and validate against git repo
    // Read orb.yml

    println!(
        "Git path: {:?}\nInfo: {:?}",
        &local_option.path[..],
        git_info::get_git_info_from_path(&local_option.path[..], &None, &None)
    );

    // TODO: Will want ability to pass in any yaml.
    // TODO: Also handle file being named orb.yaml
    // Look for a file named orb.yml
    let config = parser::load_orb_yaml(format!("{}{}", &local_option.path, "orb.yml"))?;

    Ok(())
}
