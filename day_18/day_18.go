package day_18

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type tokenType int

const (
	OP_ADD tokenType = iota
	OP_MUL
	OP_OPEN
	OP_CLOSE
	NUMBER
)

type token struct {
	t tokenType
	v int
}

func (t token) String() string {
	switch t.t {
	case OP_ADD:
		return "+"
	case OP_MUL:
		return "*"
	case OP_OPEN:
		return "("
	case OP_CLOSE:
		return ")"
	case NUMBER:
		return strconv.Itoa(t.v)
	}

	return ""
}

type Day_18 struct{}

func New() *Day_18 {
	return &Day_18{}
}

func getLines(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func evaluateExpression(expr string, precedence map[tokenType]int) int {
	tokens := toTokens(expr)
	tokens = shuntingYard(tokens, precedence)
	return calculate(tokens)
}

func isDigit(c byte) bool {
	return c >= 48 && c < 58
}

func toTokens(expr string) []token {
	tokens := make([]token, 0)

	for i := 0; i < len(expr); i++ {
		if expr[i] != ' ' {
			if isDigit(expr[i]) {
				n := int(expr[i]) - 48
				for i+1 < len(expr) && isDigit(expr[i+1]) {
					i++
					n *= 10
					n += int(expr[i]) - 48
				}
				tokens = append(tokens, token{NUMBER, n})
			} else {
				switch expr[i] {
				case '+':
					tokens = append(tokens, token{OP_ADD, 0})
				case '*':
					tokens = append(tokens, token{OP_MUL, 0})
				case '(':
					tokens = append(tokens, token{OP_OPEN, 0})
				case ')':
					tokens = append(tokens, token{OP_CLOSE, 0})
				}
			}
		}
	}

	return tokens
}

func isOperator(t token) bool {
	return t.t == OP_ADD || t.t == OP_MUL
}

func shuntingYard(tokens []token, precedence map[tokenType]int) []token {
	operatorStack := make([]token, 0)
	outputQueue := make([]token, 0)

	for len(tokens) > 0 {
		t := tokens[0]
		tokens = tokens[1:]

		if t.t == NUMBER {
			outputQueue = append(outputQueue, t)
		} else if isOperator(t) {
			for len(operatorStack) > 0 &&
				isOperator(operatorStack[len(operatorStack)-1]) &&
				precedence[operatorStack[len(operatorStack)-1].t] >= precedence[t.t] {
				outputQueue = append(outputQueue, operatorStack[len(operatorStack)-1])
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
			operatorStack = append(operatorStack, t)
		} else if t.t == OP_OPEN {
			operatorStack = append(operatorStack, t)
		} else if t.t == OP_CLOSE {
			for operatorStack[len(operatorStack)-1].t != OP_OPEN {
				outputQueue = append(outputQueue, operatorStack[len(operatorStack)-1])
				operatorStack = operatorStack[:len(operatorStack)-1]
			}
			operatorStack = operatorStack[:len(operatorStack)-1]
		}

	}

	for len(operatorStack) > 0 {
		outputQueue = append(outputQueue, operatorStack[len(operatorStack)-1])
		operatorStack = operatorStack[:len(operatorStack)-1]
	}

	return outputQueue
}

func calculate(tokens []token) int {
	stack := make([]token, 0)
	for len(tokens) > 0 {
		current := tokens[0]
		if isOperator(current) {
			res := 0
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]

			switch current.t {
			case OP_ADD:
				res = a.v + b.v
			case OP_MUL:
				res = a.v * b.v
			}

			stack = append(stack[:len(stack)-2], token{NUMBER, res})
		} else {
			stack = append(stack, current)
		}
		tokens = tokens[1:]
	}

	return stack[0].v
}

func (d *Day_18) PartA(input string) string {
	precedence := map[tokenType]int{OP_ADD: 1, OP_MUL: 1}
	total := 0

	for _, line := range getLines(input) {
		total += evaluateExpression(line, precedence)
	}

	return strconv.Itoa(total)
}

func (d *Day_18) PartB(input string) string {
	precedence := map[tokenType]int{OP_ADD: 2, OP_MUL: 1}
	total := 0

	for _, line := range getLines(input) {
		total += evaluateExpression(line, precedence)
	}

	return strconv.Itoa(total)
}
