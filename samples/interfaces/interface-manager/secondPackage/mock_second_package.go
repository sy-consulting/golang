package secondManager

import (
	"fmt"
	// Add imports here
)

const (
// Add Constants here
)

type MockSecond struct {
	Application string
}

func (MockSecond) SecondCall() {
	fmt.Println("I'm an imposter. My Name is Mr. Second Mock.")
}
