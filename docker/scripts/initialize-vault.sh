#!/bin/sh

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

VAULT_STATUS=$(vault status)

if echo "$VAULT_STATUS" | grep -q 'Initialized.*true'; then
    echo "Vault is already initialized."
    if echo "$VAULT_STATUS" | grep -q 'Sealed.*true'; then
        echo "export UNSEAL_KEY=$(cat /vault/unseal_key.txt)" >> /root/.bashrc
        vault operator unseal $UNSEAL_KEY
    fi
else
    INIT_OUTPUT=$(vault operator init -key-shares=1 -key-threshold=1)
    echo "$INIT_OUTPUT" > /vault/init_output.txt
    UNSEAL_KEY=$(echo "$INIT_OUTPUT" | grep 'Unseal Key 1:' | awk '{print $NF}')
    ROOT_TOKEN=$(echo "$INIT_OUTPUT" | grep 'Initial Root Token:' | awk '{print $NF}')
    echo $UNSEAL_KEY > /vault/unseal_key.txt
    echo "export VAULT_UNSEAL_KEY=$(cat /vault/unseal_key.txt)" >> /root/.bashrc
    echo $ROOT_TOKEN > /vault/root_token.txt
    echo "export VAULT_TOKEN=$(cat /vault/root_token.txt)" >> /root/.bashrc
    rm -rf /vault/root_token.txt /vault/unseal_key.txt /vault/unseal_key.txt
    vault operator unseal $UNSEAL_KEY
fi
