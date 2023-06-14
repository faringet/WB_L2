package main

// конкретная программа

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}
