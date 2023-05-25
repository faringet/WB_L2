package main

import "fmt"

type Device struct {
	Name string
	Next Service
}

func (device *Device) Execute(data *Data) {
	if data.GetSource { // Если данные уже получены
		fmt.Printf("Data from device [%s] already get.\n", device.Name)
		device.Next.Execute(data)
		return
	}
	fmt.Printf("Get data from device [%s].\n", device.Name) // данные приняты успешно и их нужно обработать
	data.GetSource = true
	device.Next.Execute(data) // передаем следующему звену на обработку
}

func (device *Device) SetNext(svc Service) {
	device.Next = svc // задаем следующие звено

}
