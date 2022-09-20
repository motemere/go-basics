package main

import (
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Println(Get())

	lo.ForEach(names, func(x string, _ int) {
		logrus.Println(x)
	})
}
