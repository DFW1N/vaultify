# // ########################################################################################
# // # ██████╗ ██╗   ██╗██╗   ██╗███╗   ██╗     ██████╗ ██████╗  ██████╗ ██╗   ██╗██████╗   #
# // # ██╔══██╗██║   ██║██║   ██║████╗  ██║    ██╔════╝ ██╔══██╗██╔═══██╗██║   ██║██╔══██╗  #
# // # ██████╔╝██║   ██║██║   ██║██╔██╗ ██║    ██║  ███╗██████╔╝██║   ██║██║   ██║██████╔╝  #
# // # ██╔══██╗██║   ██║██║   ██║██║╚██╗██║    ██║   ██║██╔══██╗██║   ██║██║   ██║██╔═══╝   #
# // # ██████╔╝╚██████╔╝╚██████╔╝██║ ╚████║    ╚██████╔╝██║  ██║╚██████╔╝╚██████╔╝██║       #
# // # ╚═════╝  ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝     ╚═════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝       #
# // # Author: Sacha Roussakis-Notter														#
# // # Project: Vaultify																	#
# // # Description: Dockerfile used to deploy it locally.                                   #
# // ########################################################################################

# base image
FROM alpine:3.14

# set vault version
ENV VAULT_VERSION 1.15.5

# create a new directory
RUN mkdir /vault

# download dependencies
RUN apk --no-cache add \
      bash \
      ca-certificates \
      wget \
      curl \
      nano \
      jq \
      yq \
      gzip \
      terraform \
      gnupg \
      netcat-openbsd

# download OpenTofu
RUN curl --proto '=https' --tlsv1.2 -fsSL https://get.opentofu.org/install-opentofu.sh -o install-opentofu.sh

RUN latestVersion=$(curl -s "https://api.github.com/repos/DFW1N/vaultify/releases/latest" | jq -r '.tag_name'); baseURL="https://github.com/DFW1N/vaultify/releases/download/$latestVersion"; wget -q "$baseURL/vaultify" && chmod +x vaultify && wget -q "$baseURL/vaultify.sig" && wget -qO- https://vaultify.buungroup.com/public/public.key | gpg --import && gpg --verify vaultify.sig vaultify && rm -f vaultify.sig
RUN chmod +x vaultify
RUN mv vaultify /usr/local/bin/

# download and set up vault
RUN wget --quiet --output-document=/tmp/vault.zip https://releases.hashicorp.com/vault/${VAULT_VERSION}/vault_${VAULT_VERSION}_linux_amd64.zip && \
    unzip /tmp/vault.zip -d /vault && \
    rm -f /tmp/vault.zip && \
    chmod +x /vault

# update PATH
ENV PATH="/vault:$PATH"

# add the config file
COPY ./config/vault-raft-config.hcl /vault/config/vault-config.hcl
COPY ./scripts/docker-entrypoint.sh /vault/config/docker-entrypoint.sh
COPY ./scripts/initialize-vault.sh /vault/config/initialize-vault.sh
RUN chmod +x /vault/config/docker-entrypoint.sh
RUN chmod +x /vault/config/initialize-vault.sh
RUN chmod +x install-opentofu.sh

RUN ./install-opentofu.sh --install-method standalone --skip-verify

# expose ports
EXPOSE 8200 8201 8125

# run vault
ENTRYPOINT ["/vault/config/docker-entrypoint.sh"]
CMD ["server", "-config=/vault/config/vault-config.hcl"]