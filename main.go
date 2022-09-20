package main

import (
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Println(GetGreeting())

	lo.ForEach(GetNames(), func(x string, _ int) {
		logrus.Println(x)
	})
}
