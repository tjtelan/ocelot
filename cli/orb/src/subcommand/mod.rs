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

use std::error::Error;
use std::fmt;

#[derive(Debug)]
pub struct SubcommandError {
    details: String,
}

impl SubcommandError {
    pub fn new(msg: &str) -> SubcommandError {
        SubcommandError{details: msg.to_string()}
    }
}

impl fmt::Display for SubcommandError {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f,"{}",self.details)
    }
}

impl Error for SubcommandError {
    fn description(&self) -> &str {
        &self.details
    }
}

#[derive(Debug, StructOpt)]
#[structopt(rename_all = "kebab_case")]
pub enum Subcommand {
    /// Send build signal
    Build(build_cmd::SubcommandOption),
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
    Developer(developer::DeveloperType),
    /// Get version string
    Version,
}

#[derive(Debug, StructOpt)]
pub struct GlobalOption {

}

#[derive(Debug, StructOpt)]
#[structopt(name = "orb")]
pub struct SubcommandContext {
    #[structopt(subcommand)]
    pub subcommand: Subcommand,
    #[structopt(flatten)]
    pub global_option : GlobalOption,
}