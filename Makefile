ALL: install

deps:
	yay -S install python-pip nodejs npm

install: post-install.sh

