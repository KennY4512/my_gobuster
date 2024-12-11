package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	d, t, w, q := FlagsManager()

	// Affichage des valeurs
	fmt.Println("Liste (d):", d)
	fmt.Println("URL (t):", t)
	fmt.Println("Workers (w):", w)
	fmt.Println("Quiet (q):", q)

	// // Adresse du serveur à tester
	// url := "https://github.com/test"

	// // Envoyer une requête GET
	// resp, err := http.Get(url)
	// if err != nil {
	// 	fmt.Printf("Erreur lors de la requête : %v\n", err)
	// 	os.Exit(1)
	// }
	// defer resp.Body.Close() // S'assurer de fermer la réponse

	// // Vérifier le code de statut HTTP
	// if resp.StatusCode == http.StatusOK {
	// 	fmt.Printf("Le serveur fonctionne correctement. (HTTP %d)\n", resp.StatusCode)
	// } else {
	// 	fmt.Printf("Le serveur a retourné un code HTTP inattendu : %d\n", resp.StatusCode)
	// 	os.Exit(1)
	// }
}

func FlagsManager() (d, t string, w int, q bool) {
	// Déclaration des variables pour stocker les valeurs des options
	var wordList, targetURL string
	var workersNumb int
	var quietFlag, help bool

	flag.StringVar(&wordList, "d", "", "List of words to try")
	flag.StringVar(&targetURL, "t", "", "Target URL")
	flag.IntVar(&workersNumb, "w", 1, "Number of simultaneous workers")
	flag.BoolVar(&quietFlag, "q", false, "Quiet mode")
	flag.BoolVar(&help, "h", false, "Help")

	flag.Parse()

	if help {
		fmt.Println("Usage de l'application :")
		flag.PrintDefaults()
		os.Exit(0)
	}

	return wordList, targetURL, workersNumb, quietFlag
}
