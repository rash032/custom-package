package main

// import "../types"

type InData struct {
	V int
}

type OutData struct {
	V int
}

var Input InData
var Output OutData
var Name string

func init() {
        Name = "plugin2"
}
func process() OutData {
        o := OutData{V: Input.V * 20}
        return o
}
func F() {
        Output = process()
}