package asm

import (
	"Go-SAP3/machine/asm/opcodes"
	"Go-SAP3/machine/types"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Line struct {
	Number        int
	Size          int
	MemoryAddress int
	Value         string
	Tokens        []string
	Label         string
}

// Directives allows author to provide additional information to assembler.
type Directives struct {
	PC      int // Set program counter at start up.
	Address int // Starting address for program instructions.
}

// Assemble translates assembly source code mnemonics into op code instructions.
func Assemble(source string) ([]types.Word, Directives, error) {

	// Get source lines into usable format.
	lines := ParseSourceLines(source)
	directives := ParseDirectives(lines)
	lines = TokenizeLines(lines)
	lines = CalculateMemoryAddresses(lines, directives)
	lines = FindLabels(lines)
	lines = ReplaceLabels(lines)

	// Build Words from parsed lines.
	values := ParseMnemonics(lines)

	return values, directives, nil
}

func ParseDirectives(lines []Line) Directives {
	dir := Directives{}

	for _, line := range lines {
		toks := strings.Fields(line.Value)
		if len(toks) < 3 {
			continue
		}

		if toks[0] != "#dir" {
			continue
		}

		switch toks[1] {
		case "ADDR":
			hex, err := strconv.ParseInt(toks[2], 0, 32)
			if err != nil {
				log.Fatal(fmt.Sprintf("Invalid directive %v", line))
			}
			dir.Address = int(hex)
		case "PC":
			hex, err := strconv.ParseInt(toks[2], 0, 32)
			if err != nil {
				log.Fatal(fmt.Sprintf("Invalid directive %v", line))
			}
			dir.PC = int(hex)
		}
	}

	return dir
}

func TokenizeLines(lines []Line) []Line {
	for idx, line := range lines {
		lines[idx] = TokenizeLine(line)
	}

	return lines
}

func TokenizeLine(line Line) Line {
	toks := strings.Fields(line.Value)

	for idx, tok := range toks {
		// For commands like 'MVI A,0x44' or 'MOV A,C', "LXI B,0x1234"
		if InString(tok, ",") {
			subToks := strings.Split(tok, ",")
			replaceToks := append(toks[0:idx], subToks...)
			toks = replaceToks
		}
	}

	line.Tokens = toks
	line.Size = GetByteSize(line.Tokens)

	return line
}

func GetByteSize(tokens []string) int {
	for i := 0; i < len(tokens); i++ {
		max := 4 // Max length of any particular mnemonic
		if len(tokens[i]) < 4 {
			max = len(tokens[i])
		}
		needle := tokens[i][0:max]
		if size, ok := ByteSizes[needle]; ok {
			return size
		}
	}

	return 0
}

func CalculateMemoryAddresses(lines []Line, directives Directives) []Line {
	byteCounter := directives.Address

	for idx, line := range lines {
		lines[idx].MemoryAddress = byteCounter
		byteCounter += line.Size
	}

	return lines
}

func ParseSourceLines(source string) []Line {
	lines := BuildLines(source)

	// Initial rinse
	lines = TrimLines(lines)

	// Strip all comments
	lines = StripLeadingComments(lines)
	lines = StripTrailingComments(lines)

	// Final rinse
	lines = TrimLines(lines)

	return lines
}

func BuildLines(source string) []Line {
	sourceLines := strings.Split(source, "\n")

	var lines []Line

	for counter, source := range sourceLines {
		v := strings.TrimSpace(source)
		// Skip blank lines.
		if len(v) == 0 {
			continue
		}

		line := Line{
			Number: counter,
			Value:  v,
		}
		lines = append(lines, line)
	}

	return lines
}

func TrimLines(lines []Line) []Line {
	var clean []Line

	for _, line := range lines {
		line.Value = strings.TrimSpace(line.Value)
		clean = append(clean, line)
	}

	return clean
}

func StripLeadingComments(lines []Line) []Line {
	var clean []Line

	for _, line := range lines {
		arr := []rune(line.Value)
		if len(arr) > 0 {
			if arr[0] != ';' {
				clean = append(clean, line)
			}
		}
	}

	return clean
}

var maxLabelLength = 6

func FindLabels(lines []Line) []Line {
	for idx, line := range lines {
		// Label must be first word on line.
		arr := []rune(line.Tokens[0])
		delimiterPos := StrPos(line.Value, ":")

		if delimiterPos != -1 {

			// First character must not be a number
			if isNumeric(string(arr[0])) {
				log.Fatal("Invalid label, must begin with A-Z")
			}

			if delimiterPos <= maxLabelLength+1 {

				// Record the found label.
				lines[idx].Label = string(arr[0:delimiterPos])

				// Remove label from tokens.
				lines[idx].Tokens = line.Tokens[1:]
			} else {
				log.Fatal("Invalid label length")
			}
		}
	}

	return lines
}

func ReplaceLabels(lines []Line) []Line {
	for _, line := range lines {
		if len(line.Label) > 0 {
			for idx, searchLine := range lines {
				for tIdx, tok := range searchLine.Tokens {
					if tok == line.Label {
						// Replace the label with the memory address in hex.
						lines[idx].Tokens[tIdx] = fmt.Sprintf("0x%X", line.MemoryAddress)
					}
				}
			}
		}
	}

	return lines
}

func StripTrailingComments(lines []Line) []Line {
	for idx, line := range lines {
		words := strings.Fields(line.Value)
		var keepWords []string
		for _, word := range words {
			// Ignoring everything after comment token
			arr := []rune(word)

			if len(arr) > 0 {
				if arr[0] == ';' {
					break
				}
			}
			keepWords = append(keepWords, word)
		}
		lines[idx].Value = strings.Join(keepWords, " ")
	}

	return lines
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func ParseMnemonics(lines []Line) []types.Word {
	var words []types.Word

	for _, line := range lines {

		if found, opCode := ParseLiteralInstructionMatch(line); found {
			words = append(words, opCode)
			continue
		}

		if found, opCode, msb, lsb := ParseMemoryReferenceInstructionMatch(line); found {
			words = append(words, opCode)

			// SAP is Little Endian, expects LSB first.
			words = append(words, lsb)
			words = append(words, msb)
			continue
		}

		if found, opCode, word := ParseImmediateInstructionMatch(line); found {
			words = append(words, opCode)
			words = append(words, word)
			continue
		}

		if found, opCode, msb, lsb := ParseExtendedInstructionMatch(line); found {
			words = append(words, opCode)

			// SAP is Little Endian, expects LSB first.
			words = append(words, lsb)
			words = append(words, msb)
			continue
		}
	}

	return words
}

func ParseLiteralInstructionMatch(line Line) (bool, types.Word) {
	var found bool
	var value types.Word

	// Handle special cases like "MOV A,C"
	var check string

	if len(line.Tokens) == 0 {
		return found, value
	}

	switch line.Tokens[0] {
	case "MOV":
		check = strings.Join([]string{line.Tokens[0], line.Tokens[1]}, " ")
		check = strings.Join([]string{check, line.Tokens[2]}, ",")
	default:
		check = strings.Join(line.Tokens, " ")
	}

	if exact, ok := Mnemonics[check]; ok {
		for _, literal := range opcodes.LiteralInstructions {
			if literal == exact {
				value = literal
				found = true
				break
			}
		}
	}

	return found, value
}

func ParseMemoryReferenceInstructionMatch(line Line) (found bool, opCode, msb, lsb types.Word) {

	if len(line.Tokens) < 2 {
		return
	}

	for _, word := range opcodes.MemoryReferenceInstructions {
		// OpCode is first byte.
		if word == Mnemonics[line.Tokens[0]] {
			opCode = word
			found = true

			hex, err := strconv.ParseInt(line.Tokens[1], 0, 32)
			if err != nil {
				log.Fatal("Error parsing memory reference instruction: ", err, " ", line.Value, " ", hex)
			}

			// MSB is second byte.
			msb = types.Word(hex >> 8)

			// LSB is third byte.
			lsb = types.Word(hex)

			break
		}
	}

	return
}

func ParseImmediateInstructionMatch(line Line) (found bool, opCode, w types.Word) {
	if len(line.Tokens) < 2 {
		return
	}

	// Handle special cases like "MVI A"
	var check string
	hexTokenPos := 1

	switch line.Tokens[0] {
	case "MVI":
		check = strings.Join([]string{line.Tokens[0], line.Tokens[1]}, " ")
		hexTokenPos = 2
	default:

		check = line.Tokens[0]
	}

	for _, word := range opcodes.ImmediateInstructions {
		// OpCode is first byte.
		if word == Mnemonics[check] {
			opCode = word
			found = true
			// Value is second byte.
			hex, err := strconv.ParseInt(line.Tokens[hexTokenPos], 0, 32)
			if err != nil {
				log.Fatal("Error parsing immediate instruction: ", err, " ", line.Value, " ", hex)
			}
			w = types.Word(hex)

			break
		}
	}

	return
}

func ParseExtendedInstructionMatch(line Line) (found bool, opCode, msb types.Word, lsb types.Word) {
	if len(line.Tokens) < 3 {
		return
	}

	// Handle special cases like "LXI B"
	var check string
	hexTokenPos := 1

	switch line.Tokens[0] {
	case "LXI":
		check = strings.Join([]string{line.Tokens[0], line.Tokens[1]}, " ")
		hexTokenPos = 2
	default:

		check = line.Tokens[0]
	}

	for _, word := range opcodes.ExtendedInstructions {
		// OpCode is first byte.
		if word == Mnemonics[check] {
			opCode = word
			found = true
			// MSB is second byte.
			hex, err := strconv.ParseInt(line.Tokens[hexTokenPos], 0, 32)
			if err != nil {
				log.Fatal("Error parsing extended instruction: ", err, " ", line.Value, " ", hex)
			}

			// MSB is second byte.
			msb = types.Word(hex >> 8)

			// LSB is third byte.
			lsb = types.Word(hex)

			break
		}
	}

	return
}

func InString(haystack, needle string) bool {
	if needle == "" {
		return false
	}
	idx := strings.Index(haystack, needle)
	if idx == -1 {
		return false
	}
	return true
}

func StrPos(haystack, needle string) int {
	return strings.Index(haystack, needle)
}
