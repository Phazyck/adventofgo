package main

func dayX() *day {

	part1 := func() interface{} {
		return -1
	}

	part2 := func() interface{} {
		return -1
	}

	solve := func() (interface{}, interface{}) {
		return part1(), part2()
	}

	return &day{-1, "Title", solve}
}
