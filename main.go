package main

import (
	"checkalive/config"
	"checkalive/dispatcher"
)

func main(){

	c:=config.Config{}
	c.Load()
	d :=dispatcher.Dispatcher{c};
	for _,value  := range c.Monitor.Url {
		d.Url(value)
	}

}

