#! /bin/env python
""" Script for Install all my packages in my Arch Linux """
from os import system
from sys import argv
import re

#   Installer for platform
update = "yaourt -Syu"

#   List of packages to install
packages = {
    "yaourt -S ": [
        "go",
        "vim",
        "zsh",
        "npm",
        "git",
        "zeal",
        "wine",
        "pencil",
        "kotlin",
        "docker",
        "yakuake",
        "vivaldi",
        "audacity",
        "go-tools",
        "sddm-kcm",
        "franz-bin",
        "filelight",
        "latte-dock",
        "python-pip",
        "gnome-boxes",
        "jdk8-openjdk",
        "jre8-openjdk",
        "google-chrome",
        "oh-my-zsh-git",
        "docker-compose",
        "docker-machine",
        "android-studio",
        "nylas-mail-bin",
        "atom-editor-bin",
        "kdeplasma-addons",
        "google-webdesigner",
        "visual-studio-code",
        "android-studio-canary",
        "intellij-idea-ultimate-edition",
    ],

    "sudo npm i -g ": [
        "gulp",
        "yarn",
        "babel",
        "eslint",
        "lodash",
        "express",
        "webpack",
        "nodemon",
        "typescript",
        "http-server",
        "coffeescript@next",
        "unsplash-wallpaper",
    ],

    "sudo pip install ": [
        "jupyter",
        "jupyterlab",
        "pandas",
        "sklearn",
        "numpy",
        "pylint",
        "autopep8",
    ],
}




#   Main function:
def main():
    """Execute a installation of one or more install packages"""

    # Update system first
    command = update + "; "

    if len(argv) > 1:
        # Install specify set of packages
        for pkg in argv[1:]:
            for install in packages.keys():
                # If finds the package, then install
                if re.search(pkg, install):
                    command += install + " ".join(packages[install]) + "; "
    else:
        # Full Install
        for install in packages.keys():
            command += install + " ".join(packages[install]) + "; "

    # print(command) # test command
    system(command)

if __name__ == "__main__":
    main()
