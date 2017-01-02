package controllers

func init() {
	mux["get"] = simpleHandler("get")
}
