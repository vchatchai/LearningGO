package fizzbuzz

func FizzbuzzCheck(value int) string {
	fizz := value % 3
	buzz := value % 5
	if fizz == 0 && buzz == 0 {
		return "fizzbuzz"
	}

	if fizz == 0 {
		return "fizz"
	}

	if buzz == 0 {
		return "buzz"
	}

	return string(value)

}
