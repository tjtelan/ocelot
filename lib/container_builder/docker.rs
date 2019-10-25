use shiplift::{
    builder::ContainerFilter, tty::StreamType, ContainerListOptions, ContainerOptions, Docker,
    ExecContainerOptions, PullOptions,
};

use tokio;
use tokio::prelude::{Future, Stream};


pub fn container_pull(image: Option<String>) -> Result<(), ()> {
    let docker = Docker::new();

    // TODO: The None case needs to error out
    let img = match image {
        Some(i) => i,
        None => { return Err(()) },
    };

    println!("Pulling image: {}", img);

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