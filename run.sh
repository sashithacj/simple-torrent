#!/bin/bash
docker stop cloud-torrent cloud-torrent-files 
docker rm cloud-torrent cloud-torrent-files
docker system prune -a
systemctl enable docker
cd /
rm -r simple-torrent
git clone https://github.com/sashithacj/simple-torrent.git
cd simple-torrent
docker-compose up
