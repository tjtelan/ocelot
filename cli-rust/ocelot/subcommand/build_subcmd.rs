/// This is named build_subcmd.rs bc we can't use build.rs due to overlapping with `cargo` features.

extern crate clap;

use structopt::StructOpt;

//ocelot build -acct-repo <acct>/<repo> -hash <git_hash> -branch <branch> [-latest]

#[derive(Debug, StructOpt)]
#[structopt(rename_all = "kebab_case")]
pub struct SubOption {
    /// Build provided account/repo. Otherwise try to auto-detect from current working directory
    #[structopt(long)]
    acct_repo: Option<String>,
    /// Use provided branch. Default to current active branch
    #[structopt(long)]
    branch : Option<String>,
    /// Build provided commit hash. Otherwise, default to HEAD commit of active branch
    #[structopt(long)]
    hash : Option<String>,
}

// Handle the command line control flow
pub fn subcommand_handler(args: &SubOption) {
    println!("Placeholder for running build");
}