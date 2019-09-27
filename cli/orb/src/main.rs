use structopt::StructOpt;

extern crate clap;

use subcommand::{self, SubcommandContext, Subcommand, SubcommandError};

// This is for autocompletion
//use structopt::clap::Shell;

// TODO: Create top-level error type to share between services and subcommands
fn main() -> Result<(), SubcommandError> {
    // generate `bash` completions in "target" directory
    //ApplicationArguments::clap().gen_completions(env!("CARGO_PKG_NAME"), Shell::Bash, "target");

    let parsed = SubcommandContext::from_args();

    // Pass to the subcommand handlers
    match parsed.subcommand {
        Subcommand::Build(sub_option) => subcommand::build_cmd::subcommand_handler(parsed.global_option, sub_option), 
        Subcommand::Cancel => Err(SubcommandError::new("Not yet implemented")),
        Subcommand::Logs => Err(SubcommandError::new("Not yet implemented")),
        Subcommand::Org => Err(SubcommandError::new("Not yet implemented")),
        Subcommand::Repo => Err(SubcommandError::new("Not yet implemented")),
        Subcommand::Poll => Err(SubcommandError::new("Not yet implemented")),
        Subcommand::Secret => Err(SubcommandError::new("Not yet implemented")), 
        Subcommand::Summary => Err(SubcommandError::new("Not yet implemented")),
        Subcommand::Operator => Err(SubcommandError::new("Not yet implemented")),
        Subcommand::Developer(sub_command) => subcommand::developer::subcommand_handler(parsed.global_option, sub_command), 
        Subcommand::Version => Err(SubcommandError::new("Not yet implemented")),
    }

}
