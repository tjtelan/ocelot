# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "generic/ubuntu1804"

  # config.vm.network "forwarded_port", guest: 80, host: 8080
  # config.vm.provider "virtualbox" do |vb|
  #   vb.memory = "1024"
  # end

  #config.ssh.username = 'ubuntu'

  config.vm.synced_folder ".", "/home/vagrant/orbitalci",
    type: "virtualbox",
    disabled: false

  config.vm.provision "shell", inline: <<-SHELL
    apt-get update
    apt-get install -y curl git pkg-config libssl-dev build-essential
    # Install rust via rustup script
    curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y
    echo -n 'export PATH=${HOME}/.cargo/bin:${PATH}' >> ~/.bashrc
    # Configure current session to add cargo to path
    source $HOME/.cargo/env
    # Install docker via get docker script
    curl -fsSL https://get.docker.com | sh
    usermod -aG docker vagrant
    # Compile orb
    pushd orbitalci
    make
  SHELL
end