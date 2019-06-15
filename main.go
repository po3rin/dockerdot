package main

import (
	"fmt"
	"syscall/js"

	"github.com/po3rin/dockerdot/docker2dot"
)

func registerCallbacks() {
	var cb js.Func
	document := js.Global().Get("document")
	element := document.Call("getElementById", "textarea")

	cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		text := element.Get("value").String()
		dockerfile := []byte(text)

		// https://github.com/golang/go/issues/26382
		// should wrap func with gorutine.
		go func() {
			dot, err := docker2dot.Docker2Dot(dockerfile)
			if err != nil {
				fmt.Println(err)
			}
			showGraph := js.Global().Get("showGraph")
			showGraph.Invoke(string(dot))
		}()
		return nil
	})
	js.Global().Get("document").Call("getElementById", "button").Call("addEventListener", "click", cb)
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
