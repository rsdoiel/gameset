package dice

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"text/scanner"
	"time"
)

// Dice models a dice of N sides where N must be a postive
// integer greater than zero. If N is greater than zero a
// random value in the range of 1 to N is returned. If not
// a negative one is returned.
func Dice(noOfSides int) int {
	if noOfSides < 1 {
		return -1
	}
	return rand.Intn(noOfSides) + 1
}

func ParseRoll(src string) (int, int, int, error) {
	var (
		cnt, sides, offset int
		err                error
		s                  scanner.Scanner
	)
	s.Init(strings.NewReader(src))
	parts := []string{}
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		parts = append(parts, s.TokenText())
	}
	if len(parts) < 2 {
		return 0, 0, 0, fmt.Errorf("failed parse roll description, %q", src)
	}
	cnt, err = strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, 0, fmt.Errorf("failed to convert dice count, %s", err)
	}
	sides, err = strconv.Atoi(strings.TrimPrefix(parts[1], "d"))
	if err != nil {
		return 0, 0, 0, fmt.Errorf("failed to convert dice side count, %s", err)
	}
	offset = 0
	isNeg := false
	for i := 2; i < len(parts); i += 1 {
		switch parts[i] {
		case "-":
			isNeg = true
		case "+":
			isNeg = false
		default:
			val, err := strconv.Atoi(parts[i])
			if err != nil {
				return 0, 0, 0, fmt.Errorf("failed to convert offset value, %s", err)
			}
			if isNeg {
				val = val * -1
				isNeg = false
			}
			offset += val
		}
	}
	return cnt, sides, offset, nil
}

func RollDice(args []string, displayRoll bool) (int, error) {
	rand.Seed(time.Now().UnixNano())
	vargs := len(args)
	if displayRoll && vargs > 1 {
		fmt.Printf("(")
	}
	result := 0
	for c, arg := range args {
		cnt, sides, offset, err := ParseRoll(arg)
		if err != nil {
			return 0, err
		}
		if displayRoll {
			if vargs > 1 && c > 0 {
				fmt.Printf(" + ")
			}
			fmt.Printf("(%dd%d: ", cnt, sides)
		}
		roll := 0
		for i := 0; i < cnt; i++ {
			roll = Dice(sides)
			if displayRoll {
				if i > 0 {
					fmt.Printf("+")
				}
				fmt.Printf("%d", roll)
			}
			result += roll
		}
		result += offset
		if displayRoll {
			if offset != 0 {
				fmt.Printf(") + %d", offset)
			} else {
				fmt.Printf(")")
			}
		}
	}
	if displayRoll {
		if vargs > 1 {
			fmt.Printf(")")
		}
		fmt.Printf(" = %d\n", result)
	}
	return result, nil
}
