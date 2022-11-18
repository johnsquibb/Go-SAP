# Go-SAP

An emulator for "Simple As Possible" architecture written in Go.

The SAP "Simple As Possible" emulator is inspired by the chapters from Digital Computer Electronics: 3rd Edition by
Albert P. Malvino and Jerald A. Brown. This emulator was written while iterating through each of the sections of the
textbook. It is not designed to be an efficient or concise emulator, the code is written to simulate the concepts of
each of the SAP components in a straightforward, and often verbose manner.

## Install
Download project, then run the example: `go run docs/example.go`

## Architecture

- SAP 8080-like processor architecture. Knowledge of 8085 Assembly is useful.
- 8-bit word size
- 64K memory with 16-bit addressing
- Standard 8080 ALU operations
- Input/Output registers
- Basic Assembler for assembling sample code (See cmd/demo.go or sample code below)

The [System Architecture](docs/SAP_Overview.png) diagram provides a high-level overview of the system components.

## Examples

See [examples](src/examples.src).

During writing, the emulator was tested using examples from the Malvino-Brown textbook, which cannot be provided here
due to copyright. However, many examples can be extracted from the functional tests in the machine/test directory.

The Internet is a great resource for finding 8085 examples that work with the emulator,
e.g: [https://www.tutorialspoint.com/8085-program-to-generate-fibonacci-sequence](https://www.tutorialspoint.com/8085-program-to-generate-fibonacci-sequence)
.

If you'd like to contribute examples, please submit a pull request.

## Usage

See the [Go code](docs/example.go) for setting up the emulator, assembling source code, and executing the
instructions.
