package controllers

func init() {
	mux["about"] = simpleHandler("about")
}
