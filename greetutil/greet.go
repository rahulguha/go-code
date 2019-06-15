package greetutil

func Greet(s string) string {

	greet := "Mr"
	if s != "male" {
		greet = "Mrs"
	} else {
		greet = "Mr"
	}
	return greet
}
