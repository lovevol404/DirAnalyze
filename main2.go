package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
	"strconv"
)

func main() {
	var path string

	a := app.New()
	w := a.NewWindow("DirAnalyze")

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
	vbox2 := container.NewGridWithRows(1)

	button := widget.NewButton("Start Analyze", func() {
		numChan := make(chan int, 100)
		vbox2.RemoveAll()
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
		}()
		d := getDirInfo(path, path, numChan)
		close(numChan)
		if d == nil {
			hello.SetText("can not find path:" + path)
			return
		} else {
			data := binding.NewStringTree()
			err := data.Append("", d.Path, fmt.Sprintf("%s:%s", d.Path, d.sizeH))
			if err != nil {
				hello.SetText(err.Error())
				return
			}
			appendDataTimes(d, data, 3, 1)
			tree := widget.NewTreeWithData(data,
				func(bool) fyne.CanvasObject {
					return widget.NewLabel("Template Object")
				},
				func(data binding.DataItem, _ bool, item fyne.CanvasObject) {
					item.(*widget.Label).Bind(data.(binding.String))
				})
			tree.OnSelected = func(path string) {
				text.SetText(path)
			}
			vbox2.Add(tree)
		}
	})
	vbox.Add(button)
	w.SetContent(container.New(layout.NewBorderLayout(vbox, nil, nil, nil), vbox, vbox2))
	w.Resize(fyne.NewSize(1000, 600))
	w.ShowAndRun()
}

func appendData(d *DirInfo, data binding.StringTree) {
	dirs := d.subDirs
	for _, dir := range dirs {
		err := data.Append(d.Path, dir.Path, fmt.Sprintf("%s:%s", dir.Path, dir.sizeH))
		if err != nil {
			log.Println(err)
			continue
		}
		appendData(dir, data)
	}
}

func appendDataTimes(d *DirInfo, data binding.StringTree, maxTimes, currentTimes int) {
	if currentTimes >= maxTimes {
		return
	}
	dirs := d.subDirs
	for _, dir := range dirs {
		err := data.Append(d.Path, dir.Path, fmt.Sprintf("%s:%s", dir.Path, dir.sizeH))
		if err != nil {
			log.Println(err)
		}
		appendDataTimes(dir, data, maxTimes, currentTimes+1)
	}
}
