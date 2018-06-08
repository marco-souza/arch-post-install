const apps = require("./data/apps.json")
const { exec } = require('child_process')

const update = () => exec("yaourt -Syu")

const install = type => {
  switch (type) {
    case "system":
      return exec(`yaourt -Sy ${apps[type]}`)
    case "node":
      return exec(`npm i -g ${apps[type]}`)
    case "python":
      return exec(`pip install ${apps[type]}`)
  }
}

const loadConfig = () => {
  // clone, cd, apply
  const cmd = `
    git clone https://github.com/marco-souza/zshrc.git &&
    cd zshrc &&
    ./apply.sh &&
    cd -
  `
  exec(cmd)
}

const main = () => {

  // Update
  update()

  // Load config
  loadConfig()

  // Install all apps
  for (let i of Object.keys(apps)) {
    install(i)
  }
}

main()