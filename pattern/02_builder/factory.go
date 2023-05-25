package main

type Factory struct {
	Collector Collector
}

// NewFactory когда создаем новый завод определяем то, чем он будет заниматься. Так как завод может менять выпускаемые бренды,
// должна быть реализованна передача интерфейса аргументом. (чтобы могли задавать любую комплектацию)
func NewFactory(collector Collector) *Factory {
	return &Factory{Collector: collector}
}

// SetCollector будет изменять поведение на уже существующем заводе
func (factory *Factory) SetCollector(collector Collector) {
	factory.Collector = collector
}

// CreateComputer основная функция, которая возвращает уже укомплектованную сборку, на каждом шаге можем что хотим задавать
func (factory Factory) CreateComputer() Computer {
	factory.Collector.SetCore()
	factory.Collector.SetMemory()
	factory.Collector.SetBrand()
	factory.Collector.SetGraphicCard()
	factory.Collector.SetMonitor()
	return factory.Collector.GetComputer()

}
