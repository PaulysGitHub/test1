package support

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/tebeka/selenium"
)

var driver selenium.WebDriver

// WDInit retunrs an instance of WebDriver
func WDInit() selenium.WebDriver {
	var err error
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	driver, err = selenium.NewRemote(caps, "")

	if err != nil {
		fmt.Println("Erro ao instanciar o driver:", err.Error())
	}

	driver.SetImplicitWaitTimeout(time.Second * 10)
	driver.ResizeWindow("note", 1280, 800)

	return driver
}

// SaveImage pga o print do webdriver em bytes, converte para png e salva no projeto
// To see the log report type this command in terminal
// godog --format=cucumber > log/report.json
// make sure you have the cucumber-html-reporter first
// npm install cucumber-html-reporter --save-dev
// make sure node.js is install too
// make sure to create a reporter.js file
func SaveImage(foto []byte, name string) {

	img, _, _ := image.Decode(bytes.NewReader(foto))

	out, err := os.Create("./log/screenshots" + name + ".png")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = png.Encode(out, img)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
