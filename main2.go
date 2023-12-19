package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func main() {
	var path string

	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("waiting start")
	text := widget.NewEntry()
	text.SetPlaceHolder("input path")
	text.OnChanged = func(s string) {
		path = s
	}

	vbox := container.NewVBox(
		hello,
		text,
	)
	button := widget.NewButton("Start Analyze", func() {
		dirMap = make(map[string]*DirInfo)
		numChan = make(chan int, 100)
		stopGetChan := make(chan int)
		go func() {
			num := 0
			for {
				n, ok := <-numChan
				if !ok {
					hello.SetText("files(paths) count:" + strconv.Itoa(num))
					break
				}
				num += n
				hello.SetText("files(paths) count:" + strconv.Itoa(num))
			}
			stopGetChan <- 1
		}()
		d := getDirInfo(path, path)
		close(numChan)

		<-stopGetChan
		if d == nil {
			hello.SetText("can not find path:" + path)
			return
		} else {
			data := binding.NewStringTree()
			tree := widget.NewTreeWithData(data,
				func(bool) fyne.CanvasObject {
					return widget.NewLabel("Template Object")
				},
				func(data binding.DataItem, _ bool, item fyne.CanvasObject) {
					item.(*widget.Label).Bind(data.(binding.String))
				})
			data.Append("", d.Path, fmt.Sprintf("%s:%s", d.Path, d.sizeH))
			appendData(d, data)
			tree.Resize(fyne.NewSize(1000, 1000))
			vbox.Add(tree)

		}
	})
	vbox.Add(button)
	w.SetContent(vbox)
	w.Resize(fyne.NewSize(1000, 600))
	w.ShowAndRun()
}

func appendData(d *DirInfo, data binding.StringTree) {
	dirs := d.subDirs
	for _, dir := range dirs {
		data.Append(d.Path, dir.Path, fmt.Sprintf("%s:%s", dir.Path, dir.sizeH))
		appendData(dir, data)
	}
}
