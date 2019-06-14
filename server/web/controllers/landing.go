package controllers

func init() {
	mux[""] = simpleHandler("landing")
}
