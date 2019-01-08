package main

import (
	"bytes"
	"github.com/smartystreets/goconvey/convey"
	"log"
	"os"
	"testing"
)

func Test(t *testing.T) {
	os.Args = []string{"cmd", "3", "3", "3"}

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	main()

	convey.Convey("When passing a=3, b=3 and c=3 should print Equilateral", t, func() {
		convey.So(buf.String(), convey.ShouldContainSubstring, "The type of the Triangle is Equilateral")
	})


}
