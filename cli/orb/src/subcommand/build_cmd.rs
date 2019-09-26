extern crate structopt;
use structopt::StructOpt;

#[derive(Debug, StructOpt)]
#[structopt(rename_all = "kebab_case")]
pub struct SubCommandOption {
    /// Path to local repo. Defaults to current working directory
    #[structopt(long)]
    path: Option<String>,
}

pub fn subcommand_handler(_global_option : super::GlobalOption, _local_option : SubCommandOption) {
//pub fn command_handler(args : SubCommandOption) {
    unimplemented!();
}