package main

import (
	"LANG_KENNY_GO25/basics"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	basics.Clear()
	_, t, _, _ := basics.FlagsManager()

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(t)

	if err != nil {
		fmt.Printf("Erreur lors de la requête : %v\n", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	fmt.Println("Code HTTP:", resp.StatusCode)
}
