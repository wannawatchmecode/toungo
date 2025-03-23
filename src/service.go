package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func toungeGenerateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " - " + r.RemoteAddr + ": Request received in toungeGenerateHandler from ip")
	rawInput, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading request body in toungeGenerateHandler")
		fmt.Fprintln(w, "Error handling the request")
		return
		//		return err
	}
	input := string(rawInput)
	//	fmt.Println(input)
	packaged := PackageInput(input)
	//	fmt.Println(packaged)
	output := GenerateInOrderOfMaxValsRemove(packaged)
	//	fmt.Println(output)
	fmt.Fprintln(w, output)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " - " + r.RemoteAddr + ": toungeGenerateHandler: 200 Success")
}

func RunService(port string) {
	http.HandleFunc("/tounge/generate", toungeGenerateHandler)

	fmt.Println("Server started on " + port)
	http.ListenAndServe(":"+port, nil)
}
