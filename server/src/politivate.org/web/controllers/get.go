package controllers

func init() {
	mux["get"] = Beta(simpleHandler("get"))
}
