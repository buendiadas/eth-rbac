# ETH RBAC

Use an Ethereum Account to manage access permissions (RBAC) using an Smart Contract registry.

Any user with an Ethereum account can sign a message. If the user has been accepted in a Registry Smart Contract, the service will send a `JWT`, which can be used to manage access to traditional web resources. 

### How it works

![Architecture](https://i.ibb.co/rx1cCr4/Eth-Rbac.png)


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
