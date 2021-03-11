package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}
type SumRequest struct {
    A int `json:"a"`
    B int `json:"b"`
}
type SumResult struct {
    Sum int `json:"sum"`
    Request SumRequest `json:"request"`

}
func (c App) Index() revel.Result {
	return c.Render()
}
func (c App) GetTest() revel.Result {
	return c.Render()
}
func (c App) GetTestResult() revel.Result {
	fname := c.Params.Query.Get("fname")
	lname := c.Params.Query.Get("lname")
	return c.Render(fname, lname)
}
func (c App) PostTest() revel.Result {
	return c.Render()
}

func (c App) SumJSON() revel.Result {
	//Getting request data
	data := SumRequest{}
    c.Params.BindJSON(&data)
    
	a := data.A
	b := data.B
	
	//Calculating the result
	result := a + b
	
	//Packing responce to object for future export to JSON
	dataResult := make(map[string]interface{})
    dataResult["error"] = nil
	sumResult := SumResult{Sum: result, Request: data}
	dataResult["result"] = sumResult
	
    return c.RenderJSON(dataResult)
	
}
