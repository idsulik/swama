# Swaga CLI Tool

Swaga is a simple CLI tool to list, view, and convert Swagger API endpoints. It helps you explore the API structure from a Swagger/OpenAPI specification file and convert API requests into common formats like `curl`, `fetch`.

## Features

- **List**: Display all API endpoints from a Swagger file in a structured table format.
- **View**: View detailed information about a specific API endpoint.
- **Convert**: Convert an API endpoint into `curl`, `fetch` request formats.

## Installation

You can install `swaga` by cloning the repository and building the binary:

```bash
git clone https://github.com/idsulik/swama
cd swama
go build -o swaga
```

Ensure that the binary is placed in a directory included in your `PATH`.

## Usage

Swaga provides several commands for interacting with your Swagger/OpenAPI file:

### General Usage

```bash
swaga [command]
```

### Available Commands

- **`completion`**: Generate the autocompletion script for your shell.
- **`convert`**: Convert an API endpoint to different request formats (e.g., `curl`, `fetch`).
- **`list`**: Lists all API endpoints in a structured table format.
- **`view`**: View details about a specific API endpoint.
- **`help`**: Show help for commands.

### Global Flags

- **`-f, --file`**: Path to the Swagger JSON/YAML file. If omitted, `swaga` will attempt to locate the Swagger file in the current directory.
- **`-h, --help`**: Display help for the `swaga` CLI.

### Examples

#### List all endpoints

```bash
swaga list -f ./swagger.json
```

This command lists all API endpoints defined in the `swagger.json` file.

#### Output example grouped by tag:
![preview1](https://github.com/user-attachments/assets/c493ce1e-4dcb-4353-9727-b15c07754054)
#### Output example grouped by method:
![preview2](https://github.com/user-attachments/assets/a6d6ed34-d5a6-4645-965e-04881d8eba01)

#### View details of a specific endpoint

```bash
swaga view -f ./swagger.json --endpoint /api/users
```

This command will display detailed information about the `/api/users` endpoint.

#### Convert an endpoint to a `curl` command

```bash
swaga convert -f ./swagger.json --endpoint /api/users --type curl
```

This command converts the `/api/users` endpoint to a `curl` command.

#### Generate shell autocompletion

You can generate autocompletion scripts for your shell (e.g., bash, zsh, fish):

```bash
swaga completion bash > /etc/bash_completion.d/swaga
```

## More Information

For more details on each command, use the `--help` flag with the command:

```bash
swaga [command] --help
```

For example:

```bash
swaga list --help
```

## Contributing

Feel free to contribute to this project by submitting issues or pull requests at the official [GitHub repository](https://github.com/idsulik/swama).

## License

This project is licensed under the MIT License. See the `LICENSE` file for more details.

---

Swaga simplifies exploring and interacting with your Swagger-defined API, making it easier to understand and test API endpoints quickly.