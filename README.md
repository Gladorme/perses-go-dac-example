# Perses Golang Dashboard-as-Code example

Example of doing Dashboard-as-Code with Perses using the Golang SDK

## Cheat sheet

```bash
# login first
percli login
# refresh token when needed
percli refesh
# build dashboards
percli dac build -d dashboards/ -ojson
# deploy dashboards
percli apply -d built/
# shortcut build + deploy
percli dac build -d dashboards/ -ojson && percli apply -d built/
```

## References:

Perses:
- [DaC user guide](https://perses.dev/docs/perses/v0.44.0/user-guides/dashboard-as-code.md/)
- [DaC Golang SDK documentation](https://github.com/perses/perses/tree/main/docs/dac/go)

Golang:
- [Go official website](https://go.dev/)
