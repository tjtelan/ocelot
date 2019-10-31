extern crate structopt;
use structopt::StructOpt;

use config_parser::parser;
use container_builder::docker;
use git_meta::git_info;

use log::debug;

use crate::{GlobalOption, SubcommandError};

use std::env;

//static CURRENT_DIRECTORY : String = env::var("PWD").unwrap_or(".".to_string());

#[derive(Debug, StructOpt)]
#[structopt(rename_all = "kebab_case")]
pub struct SubcommandOption {
    /// Path to local repo. Defaults to current working directory
    #[structopt(long)]
    path: Option<String>,

    /// Add env vars to build. Comma-separated with no spaces. ex. "key1=var1,key2=var2"
    #[structopt(long, short)]
    env: Option<String>,

    /// Add volume mapping from host to container. Comma-separated with no spaces. ex. "/host/path1:/container/path1,/host/path2:/container/path2"
    #[structopt(long, short)]
    volume: Option<String>,

    /// Use the specified local branch
    #[structopt(long)]
    branch: Option<String>,

    /// Use the specified commit hash
    #[structopt(long)]
    hash: Option<String>,
}

/// Returns an Option<Vec<&str>> after parsing a comma-separated KEY=VALUE string map from the cli
fn kv_csv_parser(kv_str: &Option<String>) -> Option<Vec<&str>> {
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
    _global_option: GlobalOption,
    local_option: SubcommandOption,
) -> Result<(), SubcommandError> {
    // Read options and validate against git repo
    // Read orb.yml

    // If a path isn't given, then use the current working directory based on env var PWD
    let path = &local_option
        .path
        .unwrap_or(env::var("PWD").unwrap_or_default());

    const DOCKER_SOCKET_VOLMAP: &str = "/var/run/docker.sock:/var/run/docker.sock";
    let envs_vec = kv_csv_parser(&local_option.env);
    let vols_vec = match kv_csv_parser(&local_option.volume) {
        Some(v) => {
            let mut new_vec: Vec<&str> = Vec::new();
            new_vec.push(DOCKER_SOCKET_VOLMAP);
            new_vec.extend(v.clone());
            Some(new_vec)
        }
        None => Some(vec![DOCKER_SOCKET_VOLMAP]),
    };

    debug!(
        "Git info at path ({:?}): {:?}",
        &path,
        git_info::get_git_info_from_path(&path[..], &None, &None)
    );

    // TODO: Will want ability to pass in any yaml.
    // TODO: Also handle file being named orb.yaml
    // Look for a file named orb.yml
    debug!("Loading orb.yml from path {:?}", &path);
    let config = parser::load_orb_yaml(format!("{}/{}", &path, "orb.yml"))?;

    debug!("Pulling container: {:?}", config.image.clone());
    match docker::container_pull(&config.image[..]) {
        Ok(ok) => ok, // The successful result doesn't matter
        Err(_) => {
            return Err(SubcommandError::new(&format!(
                "Could not pull image {}",
                &config.image
            )))
        }
    };

    // Create a new container
    debug!("Creating container");
    let default_command_w_timeout = vec!["sleep", "2h"];
    let container_id = match docker::container_create(
        &config.image[..],
        default_command_w_timeout,
        envs_vec,
        vols_vec,
    ) {
        Ok(container_id) => container_id,
        Err(_) => {
            return Err(SubcommandError::new(&format!(
                "Could not create image {}",
                &config.image
            )))
        }
    };

    // Start the new container

    match docker::container_start(&container_id) {
        Ok(container_id) => container_id,
        Err(_) => {
            return Err(SubcommandError::new(&format!(
                "Could not start image {}",
                &config.image
            )))
        }
    }

    // TODO: Make sure tests try to exec w/o starting the container
    // Exec into the new container
    debug!("Sending commands into container");
    for command in config.command.iter() {
        // Build the exec string
        let wrapped_command = format!("{} | tee -a /proc/1/fd/1", &command);

        let container_command = vec!["/bin/sh", "-c", wrapped_command.as_ref()];

        match docker::container_exec(container_id.as_ref(), container_command.clone()) {
            Ok(output) => {
                debug!("Command: {:?}", &command);
                debug!("Output: {:?}", &output);
                output
            }
            Err(_) => {
                return Err(SubcommandError::new(&format!(
                    "Could not create image {}",
                    &config.image
                )))
            }
        }
    }

    Ok(())
}
