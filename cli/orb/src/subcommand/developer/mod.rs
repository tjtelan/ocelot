use structopt::StructOpt;

use crate::{GlobalOption, SubcommandError};

use log::debug;
use std::env;

use container_builder;
pub mod docker;
pub mod git;
pub mod local_build;
pub mod validate;

#[derive(Debug, StructOpt)]
#[structopt(rename_all = "kebab_case")]
pub enum DeveloperType {
    /// Test git repo metadata parser
    Git(git::SubcommandOption),
    /// Test the docker driver
    Docker(docker::SubcommandOption),
    /// Test running builds
    Build(local_build::SubcommandOption),
    /// Test the config file parsers
    Validate(validate::SubcommandOption),
}

pub fn get_current_workdir() -> String {
    let path = match env::current_dir() {
        Ok(d) => format!("{}", d.display()),
        Err(_) => String::from("."),
    };

    debug!("Current workdir on host: {}", &path);
    path
}

pub fn parse_envs_input(user_input: &Option<String>) -> Option<Vec<&str>> {
    let envs = kv_csv_parser(user_input);
    debug!("Env vars to set: {:?}", envs);
    envs
}

// Automatically add in the docker socket. If we don't pass in any other volumes
// assume passing in the current working directory as well
pub fn parse_volumes_input(user_input: &Option<String>) -> Option<Vec<&str>> {
    let vols = match kv_csv_parser(user_input) {
        Some(v) => {
            let mut new_vec: Vec<&str> = Vec::new();
            new_vec.push(container_builder::DOCKER_SOCKET_VOLMAP);
            new_vec.extend(v.clone());
            Some(new_vec)
        }
        None => {
            let mut new_vec: Vec<&str> = Vec::new();
            new_vec.push(container_builder::DOCKER_SOCKET_VOLMAP);

            // There's got to be a better way to handle this...
            // https://stackoverflow.com/a/30527289/1672638
            new_vec.push(Box::leak(
                format!(
                    "{}:{}",
                    get_current_workdir(),
                    container_builder::ORBITAL_CONTAINER_WORKDIR,
                )
                .into_boxed_str(),
            ));
            Some(new_vec)
        }
    };
    debug!("Volumes to mount: {:?}", &vols);
    vols
}

/// Returns an Option<Vec<&str>> after parsing a comma-separated KEY=VALUE string map from the cli
pub fn kv_csv_parser(kv_str: &Option<String>) -> Option<Vec<&str>> {
    debug!("Parsing Option<String> input: {:?}", &kv_str);
    match kv_str {
        Some(n) => {
            let kv_vec: Vec<&str> = n.split(",").collect();
            return Some(kv_vec);
        }
        None => return None,
    }
}
pub fn subcommand_handler(
    global_option: GlobalOption,
    dev_subcommand: DeveloperType,
) -> Result<(), SubcommandError> {
    match dev_subcommand {
        DeveloperType::Git(sub_option) => git::subcommand_handler(global_option, sub_option),
        DeveloperType::Docker(sub_option) => docker::subcommand_handler(global_option, sub_option),
        DeveloperType::Build(sub_option) => {
            local_build::subcommand_handler(global_option, sub_option)
        }
        DeveloperType::Validate(sub_option) => {
            validate::subcommand_handler(global_option, sub_option)
        }
    }
}
