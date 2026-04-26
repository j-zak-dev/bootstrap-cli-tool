# bootstrap-cli-tool

A simple CLI tool which can help scaffold projects in minutes across multiple OSes covering major patterns like DDD, BDD and TDD!

## Usage:

#### Pre-requisites:

- `GO` installed on device.

#### How to run the CLI:

- Run the `go build` command.
- Run the `go install` command.
- Use `bct` to run commands.

#### Implmeneted commands:

[Major commands]:

- `createNewDomain` - creates a new domain where business, entrypoint, application and infrastructure logic lives.
  - `-n` flag accepts astring for the domain name ("newDomain" by default).
  - `-d` flag enables the DDD style config.

[Minor commands]:

- `creategitignore` - Creates a new `.gitignore`file
  - `-f` flag accepts file path for pre configured gitignore template.
