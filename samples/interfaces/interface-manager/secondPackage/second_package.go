package secondManager

import (
	"fmt"
	// Add imports here
)

const (
// Add Constants here
)

type RealSecond struct {
	Application string
}

func New() {
	fmt.Println("I'm the real McCoy. My Name is Mr. New Second.")
}

func (RealSecond) SecondCall() {
	fmt.Println("I'm the real McCoy. My Name is Mr. Second.")
}
