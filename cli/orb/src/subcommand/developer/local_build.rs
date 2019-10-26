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
    // TODO: Change default value to PWD env var
    #[structopt(long, default_value = ".")]
    path: String,

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

    debug!(
        "Git info at path ({:?}): {:?}",
        &local_option.path[..],
        git_info::get_git_info_from_path(&local_option.path[..], &None, &None)
    );

    // TODO: Will want ability to pass in any yaml.
    // TODO: Also handle file being named orb.yaml
    // Look for a file named orb.yml
    debug!("Loading orb.yml from path {:?}", &local_option.path);
    let config = parser::load_orb_yaml(format!("{}/{}", &local_option.path, "orb.yml"))?;

    // writing this down so I don't forget.
    // we probably still want to create containers with a PID 1 process that will live for a while. `sleep 2h`?

    // docker create <image> sleep 2h
    // <output container id>
    // docker exec <container id> sh -c "<command>" (per list item)

    // Though, honestly we will want a way to reference the shell we want, as well as commands with container-level env vars
    // We're going to still want to inject env vars into a build context

    // Logs are still unclear. Am I going to poll on output from the exec?
    // <command> | tee -a /proc/1/fd/1 ?? This seems to give output in docker log, but not sure where the bounds of this are
    // Checked that `tee` is installed in ubuntu, centos, alpine, clearlinux, busybox
    // It seems that even if we're not root, we can print to /proc/1/fd/1

    // Instead of piping to tee, try to get /bin/sh to "give" pid1's stdout/stderr to the exec process./prod/1/fd/1/ or /dev/stdout but we'll need to determine the fd for pid 1

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
    let default_command_w_timeout = vec!["sleep 2h"];
    let container_id = match docker::container_create(&config.image[..], default_command_w_timeout)
    {
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
    //debug!("Sending commands into container");
    //for command in config.command.iter() {
    //    // Build the exec string
    //    let wrapped_command = format!("{} | tee -a /proc/1/fd/1", &command);

    //    let container_command = vec![
    //        "/bin/sh",
    //        "-c",
    //        wrapped_command.as_ref(),
    //    ];

    //    match docker::container_exec(container_id.as_ref(), container_command.clone()) {
    //        Ok(output) => {
    //            debug!("Command: {:?}", &command);
    //            debug!("Output: {:?}", &output);
    //            output
    //        }
    //        Err(_) => {
    //            return Err(SubcommandError::new(&format!(
    //                "Could not create image {}",
    //                &config.image
    //            )))
    //        }
    //    }
    //}

    Ok(())
}
