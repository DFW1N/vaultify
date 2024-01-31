#!/bin/bash

########################################################################################
# ██████╗ ██╗   ██╗██╗   ██╗███╗   ██╗     ██████╗ ██████╗  ██████╗ ██╗   ██╗██████╗   #
# ██╔══██╗██║   ██║██║   ██║████╗  ██║    ██╔════╝ ██╔══██╗██╔═══██╗██║   ██║██╔══██╗  #
# ██████╔╝██║   ██║██║   ██║██╔██╗ ██║    ██║  ███╗██████╔╝██║   ██║██║   ██║██████╔╝  #
# ██╔══██╗██║   ██║██║   ██║██║╚██╗██║    ██║   ██║██╔══██╗██║   ██║██║   ██║██╔═══╝   #
# ██████╔╝╚██████╔╝╚██████╔╝██║ ╚████║    ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║       #
# ╚═════╝  ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝     ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝       #
# Author: Sacha Roussakis-Notter													   #
# Project: Vaultify																	   #
# Description: Docker Entry Point Shell Script to Setup Vault.                         #
########################################################################################

# NOTE: This should not be used in production, its just for local development or testing to play with `vaultify` CLI.

#################
# Set Variables #
#################

IP_ADDRESS=$(hostname -i)
echo "export VAULT_ADDR=http://${IP_ADDRESS}:8200" >> /root/.bashrc
sed -i "s|api_addr.*|api_addr = \"http://${IP_ADDRESS}:8200\"|g" /vault/config/vault-config.hcl
sed -i "s|cluster_addr.*|cluster_addr = \"http://${IP_ADDRESS}:8201\"|g" /vault/config/vault-config.hcl

#########################
# Start Hashicorp Vault #
#########################

exec vault server -config=/vault/config/vault-config.hcl