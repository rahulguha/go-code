package greetutil

func Greet(s string) string {

	greet := "Mrs"
	if s != "female" {
		greet = "Mr"
	}
	return greet
}
