#!/usr/bin/env bash

echo "Installing the latest version of NVM"
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash
source ~/.bashrc
echo "Installed NVM: $(nvm --version)"

echo "Installing the latest version of Node.js"
nvm install node
echo "Installed Node.js: $(node --version)"

echo "Upgrading NPM to the latest version"
npm install -g npm@latest
echo "Upgraded NPM: $(npm --version)"
