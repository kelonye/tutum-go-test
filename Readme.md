[Tutum azure demo]()
---

Running locally
---

```bash
fig build
fig up
```

Deploy using Tutum
---

- Sign up at [Tutum.co](http://tutum.co)
- Log into Tutum
- Click on the menu on the upper right corner of the screen
- Select Account info
- Select Api Key
- Link a custom node
- Set the following env vars appropriately:
  - TUTUM_USER
  - TUTUM_APIKEY
- Install [tutum-deploy](https://github.com/kelonye/node-tutum-deploy) with:
  
```bash
npm install -g tutum-deploy
```

- Run `make deploy` to deploy the services
- Go to https://dashboard.tutum.co/node/cluster/list
- Select your linked node and use the provided hostname to access the app
