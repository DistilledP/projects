package main

import (
	"github.com/DistilledP/lungfish/internal/libs"
)

func main() {
	libs.GetServices().GetConfig().ShowSettings()
}
