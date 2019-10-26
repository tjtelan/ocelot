use shiplift::{
    builder::ContainerFilter, tty::StreamType, ContainerListOptions, ContainerOptions, Docker,
    ExecContainerOptions, PullOptions,
};

use tokio;
use tokio::prelude::{Future, Stream};

use log::debug;

/// Returns a String to work around shiplift behavior that is different from the docker cli
/// If we give shiplift an image w/o a tag, it'll download all the tags. Usually the intended behavior is to only pull latest
fn image_tag_sanitizer(image: String) -> Result<String, ()> {
    let split = &image.split(":").collect::<Vec<_>>();

    match split.len() {
        1 => {
            return {
                debug!("Image tag was not provided. Assuming {}:latest", &image);
                Ok(format!("{}:latest", image))
            }
        }
        2 => return Ok(image),
        _ => return Err(()),
    }
}

/// Connect to the docker api and pull the provided image
/// if no tag is provided with the image, ":latest" tag will be assumed
pub fn container_pull(image: Option<String>) -> Result<(), ()> {
    let docker = Docker::new();

    let img = match image {
        Some(i) => image_tag_sanitizer(i)?,
        None => return Err(()),
    };

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

//pub fn container_create() -> Result<(), ()> {
//
//}
