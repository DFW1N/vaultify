<!-- ########################################################################################
# ██████╗ ██╗   ██╗██╗   ██╗███╗   ██╗     ██████╗ ██████╗  ██████╗ ██╗   ██╗██████╗        #
# ██╔══██╗██║   ██║██║   ██║████╗  ██║    ██╔════╝ ██╔══██╗██╔═══██╗██║   ██║██╔══██╗       #
# ██████╔╝██║   ██║██║   ██║██╔██╗ ██║    ██║  ███╗██████╔╝██║   ██║██║   ██║██████╔╝       #
# ██╔══██╗██║   ██║██║   ██║██║╚██╗██║    ██║   ██║██╔══██╗██║   ██║██║   ██║██╔═══╝        #
# ██████╔╝╚██████╔╝╚██████╔╝██║ ╚████║    ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║            #
# ╚═════╝  ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝     ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝            #
# Author: Sacha Roussakis-Notter													        #
# Project: Vaultify																	        #
# Description: Docker Entry Point Shell Script to Setup Vault.                              #
############################################################################################# -->

```bash
██╗   ██╗ █████╗ ██╗   ██╗██╗  ████████╗██╗███████╗██╗   ██╗
██║   ██║██╔══██╗██║   ██║██║  ╚══██╔══╝██║██╔════╝╚██╗ ██╔╝
██║   ██║███████║██║   ██║██║     ██║   ██║█████╗   ╚████╔╝ 
╚██╗ ██╔╝██╔══██║██║   ██║██║     ██║   ██║██╔══╝    ╚██╔╝  
 ╚████╔╝ ██║  ██║╚██████╔╝███████╗██║   ██║██║        ██║   
  ╚═══╝  ╚═╝  ╚═╝ ╚═════╝ ╚══════╝╚═╝   ╚═╝╚═╝        ╚═╝   
                                                            
```

> NOTE: This should not be used in production, its just for local development or testing to play with `vaultify` CLI.

All `Buun Group`, Hashicorp Vault, [Docker Hub](https://hub.docker.com/r/buungroup/vault-raft) containers come with `Vaultify` installed

The following commands will execute:

# Docker Start Up

```bash
git clone https://github.com/DFW1N/vaultify.git && cd docker
docker-compose up -d
docker-compose logs # <-- Wait until its ready to initialize
docker exec -it vault-raft-backend /bin/bash
/vault/config/initialize-vault.sh
source ~/.bashrc
```

### Docker Cheatsheet

```bash
docker ps # <--- List all Docker containers
docker volume ls # <--- List all Docker containers volumesdocker
docker network ls # <--- List all Docker networks
docker exec -it <container-name> /bin/bash # <--- Spawn into your Docker container with /bin/bash
docker-compose logs # <--- List the logs of your docker container while in the working direcotry of your docker-compose.yml file.
docker-compose stop # <--- Stop your docker-compose container.
docker-compose rm # <--- Delete your docker-compose container.
docker-compose restart # <--- Restart your docker-compose container.
```

---

## Building Docker Containers

This section is optional if you wanted to edit the `dockerFile`, and host your own repository on docker to push images to or to your own docker hub.

```bash
docker build --no-cache -t buungroup/vault-raft:0.18 . -f dockerFile # <--- Build your dockerFile and tag it.
docker tag buungroup/vault-raft:0.18 buungroup/vault-raft:0.18 # <--- Change this to your Docker Hub Repository. 
docker push buungroup/vault-raft:0.18 # <--- To publish to your Docker Hub Repository.
```

---

# Run Vaultify

Now that you have initialized Vault, you can run `vaultify` commands as all environment variables should be set automatically.

```bash
vaultify status
vaultify init
```

---

# Clean Up

Run the following commands to delete and clean up `Hashicorp Vault`

> Note: Please make sure, that you are inside the directory where the docker-compose.yml file exists.

```bash
docker-compose stop
docker-compose rm
docker volume rm docker_vault_data
```