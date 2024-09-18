# Swama 
[![Go Report Card](https://goreportcard.com/badge/github.com/idsulik/swama)](https://goreportcard.com/report/github.com/idsulik/swama)
[![Version](https://img.shields.io/github/v/release/idsulik/swama)](https://github.com/idsulik/swama/releases)
[![License](https://img.shields.io/github/license/idsulik/swama)](https://github.com/idsulik/swama/blob/main/LICENSE)
[![GoDoc](https://pkg.go.dev/badge/github.com/idsulik/swama)](https://pkg.go.dev/github.com/idsulik/swama)

Swama is a powerful command-line interface (CLI) tool for interacting with Swagger/OpenAPI definitions. It allows you to list, view, convert, and explore API specifications directly from the command line. Swama supports JSON and YAML formats for Swagger files, and it's available for multiple platforms through pre-built binaries.

## Features

- **List and View Endpoints**: Explore API endpoints and their details.
- **Convert Endpoints**: Convert API endpoints to `curl` or `fetch` commands for testing.
- **Explore Tags and Servers**: Easily view API tags and servers.
- **Flexible Filtering**: Filter endpoints by method, tag, or specific endpoint using wildcards.
- **Grouping**: Group endpoint listings by tag or method.
- **Support for Autocompletion**: Enable shell autocompletion for faster workflows.

## Installation

### Download Pre-Built Binaries

Swama provides pre-built binaries for Linux, macOS, and Windows. You can download the appropriate binary from the [releases page](https://github.com/idsulik/swama/releases).

1. **Download the latest release**:
    - Navigate to the [releases page](https://github.com/idsulik/swama/releases).
    - Choose the binary for your platform (Linux, macOS, Windows).

2. **Install the binary**:
    - **Linux/MacOS**: Move the binary to a directory in your `$PATH`:
      ```bash
      sudo mv swama /usr/local/bin/
      sudo chmod +x /usr/local/bin/swama
      ```
    - **Windows**: Add the binary to your system's `PATH` for global access.

### Build from Source

Alternatively, you can build Swama from source:

```bash
git clone https://github.com/idsulik/swama
cd swama
go build -o swama
```

## Usage

After installation, you can use the `swama` command to interact with Swagger/OpenAPI files.

### General Command Usage

```bash
swama [command]
```

### Available Commands

- **`completion`**: Generate the autocompletion script for the specified shell.
- **`endpoints`**: Interact with API endpoints (list, view, convert).
- **`help`**: Get help about a specific command.
- **`info`**: Display general information about the Swagger file.
- **`servers`**: Interact with servers.
- **`tags`**: Interact with API tags.

### Global Flags

- **`-f, --file string`**: Path to the Swagger JSON/YAML file. If not provided, the tool will attempt to locate the Swagger file in the current directory.
- **`-h, --help`**: Displays help for the `swama` command or any subcommand.

---

## Commands Overview

### Endpoints

The `endpoints` command allows you to list, view, and convert API endpoints.

#### List Endpoints

Lists all API endpoints from a Swagger/OpenAPI file.

```bash
swama endpoints list [flags]
```

**Available Flags**:

- `-e, --endpoint string`: Filter by endpoint, supports wildcard.
- `-g, --group string`: Group output by tag or method (default: "tag").
- `-m, --method string`: Filter by method (GET, POST, etc.).
- `-t, --tag string`: Filter by tag.

**Example**:

```bash
swama endpoints list
```

![preview](https://github.com/user-attachments/assets/59937e51-3992-4ee7-b629-a9d004310afc)

#### View Endpoint Details

Displays detailed information for a specific API endpoint.

```bash
swama endpoints view [flags]
```

**Available Flags**:

- `-e, --endpoint string`: Specify the endpoint to view.
- `-m, --method string`: Specify the method (GET, POST, etc.) of the endpoint to view.

**Example**:

```bash
swama endpoints view --method=GET --endpoint=/user
```

![preview](https://github.com/user-attachments/assets/7eff7784-f276-4027-9606-f59fdd6b0951)


#### Convert an Endpoint

Converts an API endpoint to either a `curl` or `fetch` command.

```bash
swama endpoints convert [flags]
```

**Available Flags**:

- `-e, --endpoint string`: Specify the endpoint to convert.
- `-m, --method string`: Specify the method (GET, POST, etc.).
- `-t, --type string`: Type to convert to (`curl`, `fetch`).

**Example**:

```bash
swama endpoints convert --file swagger.yaml --endpoint /api/users --method POST --type curl
```

### Tags

The `tags` command allows you to list API tags in the Swagger/OpenAPI file.

```bash
swama tags list [flags]
```

**Available Flags**:

- `-h, --help`: Displays help for the `tags` command.

**Example**:

```bash
swama tags list --file swagger.yaml
```

### Servers

The `servers` command allows you to list servers from the Swagger/OpenAPI file.

```bash
swama servers list [flags]
```

**Available Flags**:

- `-h, --help`: Displays help for the `servers` command.

**Example**:

```bash
swama servers list --file swagger.yaml
```

### Info

Displays general information about the Swagger/OpenAPI file, such as the version, title, and description.

```bash
swama info view --file swagger.yaml
```

![preview](https://github.com/user-attachments/assets/6fd03077-e7f6-4baa-8b17-17626c5d12a2)
---

## Autocompletion

Swama supports autocompletion for various shells, such as Bash and Zsh. You can generate a script for your shell to enable autocompletion.

### Example: Generate Bash Completion Script

```bash
swama completion bash > /etc/bash_completion.d/swama
```

### Example: Generate Zsh Completion Script

```bash
swama completion zsh > ~/.zsh/completion/_swama
```

## Contributing

Contributions to Swama are welcome! Feel free to submit issues or pull requests on the [GitHub repository](https://github.com/idsulik/swama).

## License

Swama is licensed under the MIT License. See the `LICENSE` file for more details.

---

With Swama, interacting with Swagger/OpenAPI files is straightforward and efficient. Whether you're exploring API endpoints, converting them to testable commands, or managing servers and tags, Swama provides a simple and powerful interface for your needs. Get started by downloading the binary or building from source today!
