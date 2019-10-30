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
#[structopt(rename_all = "keb
ab_case")]
pub struct SubcommandOption {
    /// Path to local repo. Defaults to current working directory
    #[structopt(long)]
    path: Option<String>,

    /// Add env vars to build. Comma-separated with no spaces. ex. "key1=var1,key2,var2"
    #[structopt(long)]
    env: Option<String>,

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

    // If a path isn't given, then use the current working directory based on env var PWD
    let path = &local_option
        .path
        .unwrap_or(env::var("PWD").unwrap_or(".".to_string()));

    // Probably wrap this in a helper function
    let env_str = &local_option
        .env
        .unwrap_or("".to_string());

    let env_vec : Vec<&str> = env_str.split(",").collect();
    let envs_param = match env_vec.len() {
        0 => {
            debug!("No env vars provided");
            None
        },
        _ => {
            debug!("Env var string: {:?}", &env_vec);
            Some(env_vec)
        },
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
    let container_id =
        match docker::container_create(&config.image[..], default_command_w_timeout, envs_param, None) {
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
