package main

import "LANG_KENNY_GO25/workers"

func main() {
	workers.Execute()

	// client := &http.Client{
	// 	Timeout: time.Duration(expiration) * time.Second,
	// }
	// _, err := client.Get(targetURL)
	// if err != nil {
	// 	fmt.Printf("Erreur lors de la requête : %v\n", err)
	// }
	//defer resp.Body.Close()
	//fmt.Println("Code HTTP:", resp.StatusCode)
}
