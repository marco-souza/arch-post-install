#! /bin/sh
if ! [ -x "$(command -v yay)" ]; then
  git clone https://aur.archlinux.org/yay.git
  cd yay && makepkg -si && cd .. && rm -rf yay
fi

if ! [ -x "$(command -v pip3)" ]; then
  yay -Syu python3-pip
fi

if ! [ -x "$(command -v npm)" ] ||  ! [ -x "$(command -v node)" ]; then
  yay -Syu nodejs npm
fi

export NPM_HOME=$HOME/.npm-global
export PATH=$PATH:$NPM_HOME/bin
npm config set prefix $NPM_HOME

echo "Installing system packages..."
DATA_LOCAL=./pkgs
for PM in $(ls $DATA_LOCAL); do
  case $PM in
  yay) INSTALL_CDM="-Syu --noconfirm" ;;
  apt) INSTALL_CDM="install -y" ;;
  pip3) INSTALL_CDM="install --user" ;;
  npm) INSTALL_CDM="i -g" ;;
  esac

  case $PM in
  pip3) INSTALL_CDM="sudo $PM $INSTALL_CDM";;
  *) INSTALL_CDM="$PM $INSTALL_CDM";;
  esac

  $INSTALL_CDM $(cat $DATA_LOCAL/$PM)
done

if [ ! -e 'zshrc' ];
then
  echo "Installing zsh configs..."
  DEST_PATH=../zshrc
  git clone -q https://github.com/marco-souza/zshrc.git $DEST_PATH
  cd $DEST_PATH
  make apply
fi
