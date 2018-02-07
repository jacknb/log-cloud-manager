package controllers

import (
	"strings"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

var langTypes []string

func init() {
	langTypes = strings.Split(beego.AppConfig.String("lang_types"), "|")

	for _, lang := range langTypes {
		beego.Trace("Load language", lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file:", err)
			return
		}
	}
}

type baseController struct {
	beego.Controller
	i18n.Locale
}

type AppController struct {
	baseController
}
