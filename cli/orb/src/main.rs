//#[macro_use]
extern crate structopt;

extern crate clap;

use structopt::StructOpt;

use command;

// This is for autocompletion
//use structopt::clap::Shell;

#[derive(Debug, StructOpt)]
#[structopt(rename_all = "kebab_case")]
pub enum Command {
    // Send build signal
    Build,
    // Send cancel signal
    Cancel,
    // Get logs
    Logs,
    // Actions for Organizations
    Org,
    // Actions for Repos
    Repo,
    // Actions for Polling
    Poll,
    // Do things with secrets for builds
    Secret,
    // Get summary of a repo
    Summary,
    // Administration and service settings
    Operator,
    // Developer level commands and settings
    #[structopt(alias="dev")]
    Developer,
    // Get version string
    Version,
}

#[derive(Debug, StructOpt)]
#[structopt(name = "orb")]
pub struct ApplicationArguments {
    #[structopt(subcommand)]
    pub command: Command,
}

fn main() {
    // generate `bash` completions in "target" directory
    //ApplicationArguments::clap().gen_completions(env!("CARGO_PKG_NAME"), Shell::Bash, "target");

    let matches = ApplicationArguments::from_args();

    // Do stuff with the optional args

    // TODO: Add in logic for getting backend connection info to pass into the subcommand handlers
    // TODO: Subcommands should all return a Result<T>

    // Pass to the subcommand handlers
    match matches.command {
        Command::Build => {}, 
        Command::Cancel => {},
        Command::Logs => {},
        Command::Org => {},
        Command::Repo => {},
        Command::Poll => {},
        Command::Secret => {},
        Command::Summary => {},
        Command::Operator => {},
        Command::Developer => {},
        Command::Version => {},
    }
}
