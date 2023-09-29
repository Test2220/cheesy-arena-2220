#!/bin/bash
apt update
apt dist-upgrade -y
apt install git -y

wget https://go.dev/dl/go1.21.1.linux-amd64.tar.gz

tar -C /usr/local -xzf go1.21.1.linux-amd64.tar.gz

export PATH=$PATH:/usr/local/go/bin

mkdir /CheesyArena
cd /CheesyArena
git clone https://github.com/Test2220/cheesy-arena-2220.git /CheesyArena
go build
systemctl link /CheesyArena/Cheesy-arena.service
