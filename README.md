## LoadEnv - A Simple .env Loader for Shell

### Overview
loadenv is a lightweight Go utility that reads environment variables from a .env file and generates shell export commands. It's designed to be used with eval to temporarily load environment variables into your current shell session.

### Features
- Parses standard .env file format
- Handles quoted values (both single and double quotes)
- Skips comments (lines starting with #) and empty lines
- Can specify custom .env file path

### Installation

```sh
go install github.com/aeilang/loadenv@latest
```


### Usage

#### Basic Usage

```sh
eval "$(loadenv)"
```

This loads variables from .env in the current directory.

#### Specify Custom .env File

```sh
eval "$(loadenv /path/to/custom.env)"
```


#### Preview Output (Without Loading)

```sh
loadenv  # View the export commands that would be executed
```

### Security Notes

1. Always inspect the .env file contents before loading:

```sh
cat .env
```

2. For untrusted sources, preview commands first:

```sh
loadenv untrusted.env
```

3. Remember that environment variables are:
- Only set in the current shell session
- Visible to all child processes
- Not persisted after shell exit

### How It Works

1. The Go program parses the .env file
2. Generates proper shell export commands
3. Outputs them to stdout
4. eval executes these commands in your current shell

### Limitations

1. Only supports simple KEY=VALUE syntax
2. Doesn't handle multi-line values
3. Requires eval to actually set variables
