#!/usr/bin/env bash

cd ~/Downloads/
wget https://golang.org/dl/go1.18.linux-amd64.tar.gz
sudo tar -zxvf go1.18.linux-amd64.tar.gz -C /usr/local/
rm go1.18.linux-amd64.tar.gz
