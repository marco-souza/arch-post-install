#! /bin/sh

DATA_LOCAL=./data

for PM in $(ls $DATA_LOCAL); do

  case $PM in
  yay) INSTALL_CDM="-Syu --noconfirm" ;;
  pip) INSTALL_CDM="install" ;;
  npm) INSTALL_CDM="i -g" ;;
  esac

  $PM $INSTALL_CDM $(cat $DATA_LOCAL/$PM)
done