package timeComtroller

import (
	"cherish-time-go/controllers"
	"github.com/astaxie/beego/validation"
	"cherish-time-go/logic/weapp/time"
)

type TimeDetailController struct {
	controllers.Controller
	InputData InputData
}

type InputData struct {
	Id string `data:"id" valid:"Required"`
}

func (c *TimeDetailController) Detail() {
	inputData := InputData{}
	c.GetData(&inputData)
	c.Valid(&inputData)

	valid := validation.Validation{}
	valid.Valid(&inputData)

	timeLogic := timeLogic.TimeLogic{}
	logicRet := timeLogic.GetDetail(inputData.Id)

	c.Data["json"] = controllers.RetData{0,"success",logicRet}
	c.ServeJSON()
}