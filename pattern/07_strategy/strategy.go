package main

type Strategy interface {
	Route(startPoint int, endPoint int)
}
