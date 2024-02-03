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
# Description: Easily push, pull and encrypt tofu and terraform statefiles from Vault. #
########################################################################################

log_error() {
    echo -e "\033[31mERROR:\033[0m $1" >&2
    exit 1
}

# Detect the package manager used by the Linux distribution
if command -v apt-get &>/dev/null; then
    # Debian/Ubuntu
    package_manager="sudo apt-get"
elif command -v apk &>/dev/null; then
    # Alpine Linux
    package_manager="apk"
elif command -v yum &>/dev/null; then
    # Red Hat/CentOS
    package_manager="yum"
elif command -v dnf &>/dev/null; then
    # Fedora
    package_manager="dnf"
elif command -v zypper &>/dev/null; then
    # openSUSE
    package_manager="zypper"
else
    echo "Unsupported package manager. Please install 'jq', 'wget', and 'gnupg' manually."
    exit 1
fi

# Install jq, wget, and gnupg based on the detected package manager
case $package_manager in
    apt-get)
        echo "jq could not be found, installing..."
        apt-get update && apt-get install jq -y
        ;;
    apk)
        apk add jq
        ;;
    yum)
        yum install jq -y
        ;;
    dnf)
        dnf install jq -y
        ;;
    zypper)
        zypper install jq -y
        ;;
esac

# Install wget and gnupg using the package manager
$package_manager update && $package_manager install wget gnupg -y

OS=$(uname -s)
ARCH=$(uname -m)

case $OS in
    Linux) OS="Linux" ;;
    Darwin) OS="Darwin" ;;
    *) log_error "Unsupported OS: $OS" ;;
esac

case $ARCH in
    x86_64) ARCH="x86_64" ;;
    arm64) ARCH="arm64" ;;
    aarch64) ARCH="arm64" ;;
    *) log_error "Unsupported architecture: $ARCH" ;;
esac

latestVersion=$(curl -s "https://api.github.com/repos/DFW1N/vaultify/releases/latest" | jq -r '.tag_name')
if [ -z "$latestVersion" ]; then
    log_error "Failed to fetch the latest version of vaultify"
fi

currentVersion=$(vaultify --version 2>/dev/null | grep -oE 'v[0-9]+\.[0-9]+\.[0-9]+')
if [ "$currentVersion" = "$latestVersion" ]; then
    echo "vaultify $currentVersion is already installed."
    exit 0
fi

baseURL="https://github.com/DFW1N/vaultify/releases/download/$latestVersion"

# Modify the archiveName based on OS and ARCH
if [ "$OS" = "Linux" ]; then
    if [ "$ARCH" = "x86_64" ]; then
        archiveName="vaultify_linux_x86_64.tar.gz"
    elif [ "$ARCH" = "arm64" ]; then
        archiveName="vaultify_linux_arm64.tar.gz"
    else
        log_error "Unsupported architecture for Linux: $ARCH"
    fi
elif [ "$OS" = "Darwin" ]; then
    if [ "$ARCH" = "x86_64" ]; then
        archiveName="vaultify_darwin_x86_64.tar.gz"
    else
        log_error "Unsupported architecture for Darwin: $ARCH"
    fi
else
    log_error "Unsupported OS: $OS"
fi

echo "Downloading: $baseURL/$checksumsFile"
checksumsFile="checksums.txt"

wget -q "$baseURL/$checksumsFile" || log_error "Failed to download checksums.txt"

# Verify checksum
expectedChecksum=$(grep "$archiveName" "$checksumsFile" | awk '{print $1}')
rm -f "$checksumsFile"
if [ -z "$expectedChecksum" ]; then
    log_error "Failed to find checksum for $archiveName in checksums.txt"
fi

echo "Downloading: $baseURL/$archiveName"
wget -q "$baseURL/$archiveName" || log_error "Failed to download $archiveName"

echo "$expectedChecksum  $archiveName" | sha256sum -c || log_error "Checksum verification failed for $archiveName"

tar -xzf "$archiveName" || log_error "Failed to extract $archiveName"
chmod +x vaultify

mv vaultify /usr/local/bin/ || log_error "Failed to move vaultify to /usr/local/bin/"

rm -f "$archiveName"

echo "Vaultify $latestVersion installed successfully."
