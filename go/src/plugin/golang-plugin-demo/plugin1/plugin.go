package main


import "plugin/golang-plugin-demo/types"

var Input dt.InData
var Output dt.OutData
var Name string

func init() {
        Name = "plugin1"
        Input.V = 1
}

func process() OutData {
        o := OutData{V: Input.V * 2}
        return o
}
func F() {
        Output = process()
}