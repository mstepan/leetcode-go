# LeetCode in Go

Go solutions to LeetCode problems, with a small executable for trying them locally.

## Requirements

- Go 1.27.1 or newer, as specified in `go.mod`.
- Make to use the targets below. You can also run the Go commands directly.
- Optional: VS Code with the Go extension for debugging.

## Getting started

```sh
git clone https://github.com/mstepan/leetcode-go.git
cd leetcode-go
make
```

The default target runs `all`, which formats the code, builds the executable,
runs static checks and tests, and executes the example in `main.go`.

## Commands

| Command | Action |
| --- | --- |
| `make` or `make all` | Run `fmt`, `build`, `vet`, `test`, and `run` |
| `make fmt` | Format all Go packages with `go fmt ./...` |
| `make build` | Build the root package into `./leetcode` |
| `make vet` | Check all packages with `go vet ./...` |
| `make test` | Run tests in all packages with `go test ./...` |
| `make run` | Run the example with `go run .` |
| `make clean` | Remove the `leetcode` executable |

To run the example without Make:

```sh
go run .
```

## Project structure

```text
.
├── .vscode/launch.json              # VS Code debug configuration
├── go.mod                          # Module and Go version
├── main.go                         # Local example runner
├── Makefile                        # Development commands
└── medium/
    ├── repeated_dna_sequences.go    # Repeated DNA Sequences solution
    └── repeated_dna_sequences_test.go
```

The current example calls `medium.FindRepeatedDnaSequences` to find repeated
10-character sequences in a DNA string. Result order is unspecified.
The test file currently contains only a package declaration; no tests are implemented yet.

## Debugging in VS Code

1. Open the repository folder in VS Code.
2. Install the Go extension (`golang.go`).
3. Press **F5** using the **Run leetcode** launch configuration.
4. Install Delve if the extension prompts you.

The configuration builds and debugs the root package. Set breakpoints in
`main.go` or the solution files to inspect execution.

## Adding solutions

Add a solution file to the appropriate difficulty package and a corresponding
`*_test.go` file with test cases. Functions called from `main.go` must be exported
(start with a capital letter) and accessed through their imported package.
Update `main.go` to try the solution, then run `make`.
