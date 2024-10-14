package main

import "net/http"

func HandleErrors(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, 400, "Something Went Wrong")
}
