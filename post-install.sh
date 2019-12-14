#! /bin/sh
if ! [ -x "$(command -v yay)" ]; then
  git clone https://aur.archlinux.org/yay.git
  cd yay && makepkg -si && cd .. && rm -rf yay
fi

if ! [ -x "$(command -v npm)" ]; then
  yay -S install nodejs npm
fi

export NPM_HOME=$HOME/.npm-global
export PATH=$PATH:$NPM_HOME/bin
npm config set prefix $NPM_HOME

echo "Installing system packages..."
DATA_LOCAL=./pkgs
for PM in $(ls $DATA_LOCAL); do
  case $PM in
  yay) INSTALL_CDM="-Syu --noconfirm" ;;
  pip) INSTALL_CDM="install --user" ;;
  npm) INSTALL_CDM="i -g" ;;
  esac

  $PM $INSTALL_CDM $(cat $DATA_LOCAL/$PM)
done

echo "Installing zsh configs..."
git clone -q https://github.com/marco-souza/zshrc.git
cd zshrc
make
