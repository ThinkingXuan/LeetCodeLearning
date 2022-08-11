package main

func main() {

}

func gcd(a, b int64) int64 {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
