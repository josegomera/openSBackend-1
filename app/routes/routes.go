// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).URL
}


type tSpeechController struct {}
var SpeechController tSpeechController


func (_ tSpeechController) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("SpeechController.Index", args).URL
}

func (_ tSpeechController) Show(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("SpeechController.Show", args).URL
}

func (_ tSpeechController) Create(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("SpeechController.Create", args).URL
}

func (_ tSpeechController) Update(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("SpeechController.Update", args).URL
}

func (_ tSpeechController) Delete(
		id string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("SpeechController.Delete", args).URL
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}


