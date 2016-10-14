package controllers

import (
	"RESTful/models"
	"github.com/astaxie/beego"
	"strconv"
)

// Operations about object
type RestController struct {
	beego.Controller
}

func (o *RestController) List() {
	objectId := o.Ctx.Input.Param(":objectId")
	if objectId != "" {
		id, err := strconv.Atoi(objectId)
		if err != nil {
			o.Data["json"] = "object no find"
		} else {
			o.Data["json"] = models.GetA(id)
		}

	}
	o.ServeJSON()
}
