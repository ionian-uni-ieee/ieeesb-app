# Install docker
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# Get compose files
mkdir ./ieeesb-app -f
cd ./ieeesb-app
wget https://raw.githubusercontent.com/ionian-uni-ieee/ieeesb-app/dev/build/docker/docker-compose.yml 
wget https://raw.githubusercontent.com/ionian-uni-ieee/ieeesb-app/dev/build/docker/docker-compose.prod.ym
