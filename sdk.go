package sdk

// Hello returns a personalised greeting.
// If name is empty, it falls back to "World".
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name + "!"
}
