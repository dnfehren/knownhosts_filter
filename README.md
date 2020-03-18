# KnownHosts file filter tool

> assumes your knownhosts file is in `~/.ssh/knownhosts`

- `go build`
- copy the binary to where your shell will find it
- `knownhosts_filter dev` => remove all the knownhosts file entries with dev
- `knownhosts_filter dev web` => remove all the knownhosts file entries with dev AND web

