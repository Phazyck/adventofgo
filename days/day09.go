package days

import (
	"bufio"

	"github.com/Phazyck/AdventOfGo/assert"
	"github.com/Phazyck/AdventOfGo/day"
	"github.com/Phazyck/AdventOfGo/input"
)

// Day09 is the 9th day in Advent of Code.
func Day09() *day.Day {

	countGarbage := func(r *bufio.Reader) int {
		count := 0
		for {
			c, _, err := r.ReadRune()
			assert.IsNil(err)
			switch c {
			case '>':
				return count
			case '!':
				_, _, err = r.ReadRune()
				assert.IsNil(err)
			default:
				count++
			}
		}
	}

	solve := func() (interface{}, interface{}) {
		f := input.Open(9)
		r := bufio.NewReader(f)

		total := 0
		score := 0
		garbage := 0
		for {
			c, _, err := r.ReadRune()

			if err != nil {
				break
			}

			switch c {
			case '{':
				score++

			case '}':
				total += score
				score--
			case '!':
				_, _, err = r.ReadRune()
				assert.IsNil(err)
			case '<':
				garbage += countGarbage(r)
			}
		}

		return total, garbage
	}

	return day.NewDay(9, "Stream Processing", solve)
}