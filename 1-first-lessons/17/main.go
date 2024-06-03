package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func sum[T Number](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

func compare[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Alex": 1000, "Jack": 2000, "Mary": 3000}
	m2 := map[string]float64{"Alex": 102.1, "Jack": 203.2, "Mary": 304.9}
	m3 := map[string]MyNumber{"Alex": 1000, "Jack": 2000, "Mary": 3000}

	println(sum(m))
	println(sum(m2))
	println(sum(m3))
	println(compare("as", "as"))
	println(compare("10", "10"))
	println(compare("as", "10"))
}
