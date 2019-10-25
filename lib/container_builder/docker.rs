use shiplift::{
    builder::ContainerFilter, tty::StreamType, ContainerListOptions, ContainerOptions, Docker,
    ExecContainerOptions, PullOptions,
};

use tokio;
use tokio::prelude::{Future, Stream};

// If we give shiplift an image w/o a tag, it'll download all the tags. Usually the intended behavior is to only pull latest
fn image_tag_sanitizer(image: String) -> Result<String, ()> {
    let split = &image.split(":").collect::<Vec<_>>();

    match split.len()  {
        1 => return Ok(format!("{}:latest", image)),
        2 => return Ok(image),
        _ => return Err(()),
    }
}

pub fn container_pull(image: Option<String>) -> Result<(), ()> {
    let docker = Docker::new();

    let img = match image {
        Some(i) => image_tag_sanitizer(i)?,
        None => return Err(()),
    };

    //println!("Pulling image: {}", img);

    let img_pull = docker
        .images()
        .pull(&PullOptions::builder().image(img.clone()).build())
        .for_each(|output| {
            println!("{:?}", output);
            Ok(())
        })
        .map_err(|e| eprintln!("Error: {}", e));
    Ok(tokio::run(img_pull))
    //Ok(())
}

//pub fn container_create() -> Result<(), ()> {
//
//}
