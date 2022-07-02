# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.238.0/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.16, 1.17, 1-bullseye, 1.16-bullseye, 1.17-bullseye, 1-buster, 1.16-buster, 1.17-buster
ARG VARIANT="1.18-bullseye" 
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Choice] Node.js version: none, lts/*, 16, 14, 12, 10
ARG NODE_VERSION="none"
RUN if [ "${NODE_VERSION}" != "none" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# OS packages
RUN set -x \
    && apt-get update \
    && apt-get upgrade -y \
    && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends \
       zsh-autosuggestions \
       zsh-syntax-highlighting \
       pass gnupg2 httpie netcat \
    && apt-get autoremove -y  \
    && apt-get clean -y \
    && rm -r /var/cache/* /var/lib/apt/lists/*

# aws-vault
ENV AWS_VAULT_URL=github.com/99designs/aws-vault/releases/download
ENV AWS_VAULT_VERSION=v6.6.0
RUN set -ex \
    && curl -L -o /usr/local/bin/aws-vault \
        https://${AWS_VAULT_URL}/${AWS_VAULT_VERSION}/aws-vault-linux-amd64 \
    && chmod 755 /usr/local/bin/aws-vault

# Pulumi
RUN set -ex \
    && curl -fsSL https://get.pulumi.com | bash -s -- --version $PULUMI_VERSION \
    && mv ~/.pulumi/bin/* /usr/local/bin
    
# Sentry CLI
RUN set -ex \
    && curl -fsSL https://sentry.io/get-cli/ | bash

# oh-my-zsh plugins
# 	zsh-autosuggestions - suggests commands as you type based on history and completions
#   zsh-syntax-highlighting - enables highlighting of commands as they are typed at a zsh prompt
ARG USERNAME=vscode
ARG PLUGIN_FOLDER="/home/$USERNAME/.oh-my-zsh/custom/plugins"
RUN git clone https://github.com/zsh-users/zsh-syntax-highlighting.git "$PLUGIN_FOLDER"/zsh-syntax-highlighting
RUN git clone https://github.com/zsh-users/zsh-autosuggestions "$PLUGIN_FOLDER"/zsh-autosuggestions
RUN git clone https://github.com/blimmer/zsh-aws-vault.git "$PLUGIN_FOLDER"/zsh-aws-vault
COPY dotfiles/.zshrc /home/$USERNAME/

COPY automator      /workspaces/automator

# Remove library scripts for final image
RUN set -ex \
    && apt-get --purge remove unzip -y \
    && apt-get autoremove --assume-yes \
    && apt-get clean --assume-yes \
    && rm -rf /var/lib/apt/lists/* \
    && rm -rf /tmp/library-scripts \
    && rm -rf /tmp/packages-tmp \
    && rm -rf /tmp/*