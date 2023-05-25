package main

import "fmt"

type DataService struct {
	Name string
	Next Service
}

// Execute тут данные не передаем, а просто сохраняем
func (upd *DataService) Execute(data *Data) {
	if data.UpdateSource {
		fmt.Println("Data not update")
		return
	}
	fmt.Println("Data saved")
}

func (upd *DataService) SetNext(svc Service) {
	upd.Next = svc

}
