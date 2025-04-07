package main

func fizzBuzz() {
	var length int
	length = 20

	for i := 1; i <= length; i++ {
		if i%3 == 0 && i%5 == 0 {
			println("FizzBuzz")
		} else if i%3 == 0 {
			println("Fizz")
		} else if i%5 == 0 {
			println("Buzz")
		} else {
			println(i)
		}
	}
}
