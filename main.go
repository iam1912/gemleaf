package main

import (
	"github.com/iam1912/gemseries/gemleaf/config"
	"github.com/iam1912/gemseries/gemleaf/service"
)

func main() {
	DB := config.MustGetDB()
	service.Run(DB)
}
