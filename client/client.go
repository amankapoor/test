package main

import "honnef.co/go/js/dom"

func main() {

	println("The sample Isomorphic Go skeleton web app successfully printed to your web browser console using GopherJS!")
	d := dom.GetWindow().Document()
	h := d.GetElementByID("welcomeMessage")
	h.SetInnerHTML("<p>Space for Gist and vision</p><h3>Featured Articles here</h3><h3>Featured HowTo Guides here</h3>")
}
