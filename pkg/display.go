package pkg

type Display interface {
	// Display is the data to show to end user.
	Display() []byte
}
