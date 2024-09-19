# Lockdiff

Lockdiff is a command-line tool for comparing two package-lock.json files. It identifies shared dependencies and displays their version differences in a tabular format.

## Warning :warning:
This project is still in WIP

## Installation

- Clone the repository:

```bash
git clone https://github.com/mathiasdonoso/lockdiff.git
cd lockdiff
```

- Build the tool:

```bash
go build -o lockdiff .
```

- Install the CLI tool (optional):

```bash
mv lockdiff /usr/local/bin/lockdiff
```

## Usage

### Basic Usage

```bash
./lockdiff <path/to/package-lock1.json> <path/to/package-lock2.json>
```

### Example

```bash
./lockdiff ./package-lock1.json ./package-lock2.json
```

This will output a table showing version differences for shared dependencies between the two package-lock.json files.

## Contributing

Feel free to open issues or submit pull requests to improve the tool. Contributions are welcome!
