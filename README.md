# Arch Post Install

A simple program to install my prefered packages after install Antergos.

You are welcome to contribute and implemente installers for whatever Package Manager you want! 😉

## Steps

1. Update sistem with yaourt
2. Setup zsh configuration based on `https://github.com/marco-souza/zshrc.git`
3. Install `yaourt` apps
4. Install `node` apps
5. Install `pip` apps

## How to use it

Just clone the project, edit `arch-post-install/data/apps.json` to put your favorite apps and run `./post-install` from inside the cloned folder.

```sh
git clone https://github.com/marco-souza/arch-post-install.git && \
cd arch-post-install && \
./post-install

```

## Defining apps to be installed

The file `data/apps.json` specifies apps for each installer.

We support 3 kinds of installers:

- yaourt [1]
- npm [2]
- pip [3]

## Next steps

- Install yaourt as dependency
- Create a `release` pipeline in makefile
