use shiplift::{
    builder::ContainerFilter, tty::StreamType, ContainerListOptions, ContainerOptions, Docker,
    ExecContainerOptions, PullOptions,
};

use tokio;
use tokio::prelude::{Future, Stream};

use log::{debug, error};

/// Returns a String to work around shiplift behavior that is different from the docker cli
/// If we give shiplift an image w/o a tag, it'll download all the tags. Usually the intended behavior is to only pull latest
fn image_tag_sanitizer(image: &str) -> Result<String, ()> {
    let split = &image.split(":").collect::<Vec<_>>();

    match split.len() {
        1 => {
            return {
                debug!("Image tag was not provided. Assuming {}:latest", &image);
                Ok(format!("{}:latest", image))
            }
        }
        2 => return Ok(image.to_string()),
        _ => return Err(()),
    }
}

/// Connect to the docker api and pull the provided image
/// if no tag is provided with the image, ":latest" tag will be assumed
pub fn container_pull(image: &str) -> Result<(), ()> {
    let docker = Docker::new();

    let img = image_tag_sanitizer(image)?;

    debug!("Pulling image: {}", img);

    let img_pull = docker
        .images()
        .pull(&PullOptions::builder().image(img.clone()).build())
        .for_each(|output| {
            debug!("{:?}", output);
            Ok(())
        })
        .map_err(|e| eprintln!("Error: {}", e));
    Ok(tokio::run(img_pull))
}

/// Returns the id of the container that is created
pub fn container_create(image: &str, command: Vec<&str>) -> Result<String, ()> {
    let docker = Docker::new();

    // TODO: Need a naming convention
    let container_spec = ContainerOptions::builder(image)
        //.name("test-container-name")
        .attach_stdout(true)
        .attach_stderr(true)
        .cmd(command)
        .build();

    let new_container = docker
        .containers()
        .create(&container_spec)
        .map(|info| {
            debug!("{:?}", info);
            info
        })
        .map_err(|e| {
            error!("Error: {}", e);
            e
        });

    let mut container_runtime = tokio::runtime::Runtime::new().expect("Unable to create a runtime");

    // Wait for the container to be created so we can get its container id
    let container = match container_runtime.block_on(new_container) {
        Ok(runtime) => runtime,
        Err(_) => return Err(()),
    };

    Ok(container.id)
}
