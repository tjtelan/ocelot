//#[macro_use]
extern crate structopt;
use structopt::StructOpt;

pub mod build_cmd;
pub mod cancel;
pub mod logs;
pub mod org;
pub mod repo;
pub mod poll;
pub mod secret;
pub mod summary;
pub mod operator;
pub mod developer;

#[derive(Debug, StructOpt)]
#[structopt(rename_all = "kebab_case")]
pub enum SubCommand {
    /// Send build signal
    Build(build_cmd::SubCommandOption),
    /// Send cancel signal
    Cancel,
    /// Get logs
    Logs,
    /// Actions for Organizations
    Org,
    /// Actions for Repos
    Repo,
    /// Actions for Polling
    Poll,
    /// Do things with secrets for builds
    Secret,
    /// Get summary of a repo
    Summary,
    /// Administration and service settings
    Operator,
    /// Developer level commands and settings
    #[structopt(alias="dev")]
    Developer,
    /// Get version string
    Version,
}

#[derive(Debug, StructOpt)]
pub struct GlobalOption {

}

#[derive(Debug, StructOpt)]
#[structopt(name = "orb")]
pub struct ApplicationArguments {
    #[structopt(subcommand)]
    pub subcommand: SubCommand,
    #[structopt(flatten)]
    pub global_option : GlobalOption,
}