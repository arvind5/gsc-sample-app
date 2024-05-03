#!/bin/bash

# Install required packages
sudo apt-get install build-essential -y
sudo apt-get install docker.io docker-buildx-plugin python3 python3-pip -y

# Install Python Packages
pip3 install docker jinja2 tomli tomli-w pyyaml