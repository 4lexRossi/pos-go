package main

func main() {
	a := 1
	b := 2
	c := 3
	switch a {
	case 1:
		println("a")
	default:
		println("not a")
	}

	if a > b && c > a {
		println("a > b and c > a")
	}
}
