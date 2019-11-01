extern crate structopt;
use structopt::StructOpt;

use config_parser::parser;
use container_builder::docker;
use git_meta::git_info;

use log::debug;

use crate::{GlobalOption, SubcommandError};

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

pub fn subcommand_handler(
    _global_option: GlobalOption,
    local_option: SubcommandOption,
) -> Result<(), SubcommandError> {
    // Read options and validate against git repo
    // Read orb.yml

    // If a path isn't given, then use the current working directory based on env var PWD
    let path = &local_option.path.unwrap_or(super::get_current_workdir());

    let envs_vec = super::parse_envs_input(&local_option.env);
    let vols_vec = super::parse_volumes_input(&local_option.volume);

    debug!(
        "Git info at path ({:?}): {:?}",
        &path,
        git_info::get_git_info_from_path(path.as_str(), &None, &None)
    );

    // TODO: Will want ability to pass in any yaml.
    // TODO: Also handle file being named orb.yaml
    // Look for a file named orb.yml
    debug!("Loading orb.yml from path {:?}", &path);
    let config = parser::load_orb_yaml(format!("{}/{}", &path, "orb.yml"))?;

    debug!("Pulling container: {:?}", config.image.clone());
    match docker::container_pull(config.image.as_str()) {
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
        config.image.as_str(),
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
