package secondManager

type SecondIF interface {
	SecondCall()
}

func SecondCall(message SecondIF) {
	message.SecondCall()
}
