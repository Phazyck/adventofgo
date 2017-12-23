package main

import (
	"github.com/Phazyck/AdventOfGo/day"
	"github.com/Phazyck/AdventOfGo/days"
)

func printAll(days []*day.Day) {
	for _, day := range days {
		day.PrintShort()
	}
}

func printLatest(days []*day.Day) {
	l := len(days)
	i := l - 1
	day := days[i]
	day.PrintShort()
}

func main() {
	days := []*day.Day{
		days.Day01(),
		days.Day02(),
		days.Day03(),
		days.Day04(),
		days.Day05(),
		days.Day06(),
		days.Day07(),
		days.Day08(),
		days.Day09(),
		days.Day10(),
		days.Day11(),
		days.Day12(),
		days.Day13(),
		days.Day14(),
	}

	//printAll(days)
	printLatest(days)
}
