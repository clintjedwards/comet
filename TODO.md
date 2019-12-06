- User should be able to opt out of a randomly generated name and be able to provide their own name
- Comets should by default crash after the predetermined time in the config, but this should be able to be changed by the user also

* Will use hashicorp go-plugin to provide a plugin interface where backend and auth can be changed easily
* There should be a streaming endpoint when a comet is being created for logs and status updates
* We should make an rpc or maybe a command line flag that causes comet to reload backend plugins so that they can stay up to date
