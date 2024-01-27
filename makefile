.PHONY: setup install-curl install-gzip install-all

install-all: setup install-curl install-gzip

setup:
		sudo apt-get update && sudo apt-get install -y


install-curl:
		sudo apt-get install curl -y && echo "curl installed successfully.

install-gzip:
		sudo apt-get install gzip -y && echo "gzip installed successfully.
