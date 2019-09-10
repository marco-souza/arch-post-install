ALL: install

deps:
	git clone https://aur.archlinux.org/yay.git
	cd yay
	makepkg -si
	yay -S install python-pip nodejs npm

install: deps
	sh post-install.sh

