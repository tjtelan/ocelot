use structopt::StructOpt;

extern crate clap;

use subcommand::{self, ApplicationArguments, SubCommand};

// This is for autocompletion
//use structopt::clap::Shell;

fn main() {
    // generate `bash` completions in "target" directory
    //ApplicationArguments::clap().gen_completions(env!("CARGO_PKG_NAME"), Shell::Bash, "target");

    let parsed = ApplicationArguments::from_args();

    // Do stuff with the optional args

    // TODO: Add in logic for getting backend connection info to pass into the subcommand handlers
    // TODO: Subcommands should all return a Result<T>

    // Pass to the subcommand handlers
    match parsed.subcommand {
        SubCommand::Build(local_option) => subcommand::build_cmd::subcommand_handler(parsed.global_option, local_option), 
        SubCommand::Cancel => {},
        SubCommand::Logs => {},
        SubCommand::Org => {},
        SubCommand::Repo => {},
        SubCommand::Poll => {},
        SubCommand::Secret => {},
        SubCommand::Summary => {},
        SubCommand::Operator => {},
        SubCommand::Developer => {},
        SubCommand::Version => {},
    }
}
