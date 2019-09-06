# ETH RBAC

[WIP] Middle service to manage server access permissions (RBAC) using your inmutable Smart Contract registry

### Install Steps

1. Clone this repository under `go/src/github.com/carlos-buendia/eth-rbac` and `cd` into it

2. Run `go get` to get dependencies for the application

3. Create postgres database and username for this application

4. Add the following environment variables and restart shell/re-source config file:

```bash
export ETH_RBAC_DB_USERNAME={your user you created}
export ETH_RBAC_DB_PASSWORD={password for said user}
export ETH_RBAC_DB_NAME={name of the db you created}
export ETH_RBAC_DB_PORT={port of your local postgres server, usually 5432}
export ETH_RBAC_DB_HOST=localhost
export ETH_RBAC_SECRET={JWT secret to be shared with auth server, likely a keyboard cat locally}
```

5. Start the app with `go run main.go`

Tables will be auto migrated at this point, and you should be good to go!
