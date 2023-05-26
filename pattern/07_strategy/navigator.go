package main

type Navigator struct {
	Strategy
}

// SetStrategy замена поведения на лету
func (nav *Navigator) SetStrategy(str Strategy) {
	nav.Strategy = str

}
