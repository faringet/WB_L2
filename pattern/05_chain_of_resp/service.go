package main

type Service interface {
	Execute(*Data)
	SetNext(Service)
}

type Data struct {
	GetSource    bool // Признак того были ли приняты данные (выполнился ли прием данных)
	UpdateSource bool // Ставит отметку тот сервис, который обработал данные
}
