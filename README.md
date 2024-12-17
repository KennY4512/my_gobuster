# My Gobuster

## Description

Ce programme Go a pour but de recréer Gobuster simplifié afin d'apprendre à développer en Go.  
My Gobuster est un outil permettant de faire du *brute-force* seulement sur les URLs.

## Environnement d'exécution

- Système garantie compatible : Linux
- Version de Go garantie compatible : Go 1.23.2 ou supérieur

## Usage

### 1. Compilation du programme

Le programme peut être compilé avec la commande suivante :

```shell
go build main.go
```

### 2. Appel du programme

Appel du programme compilé :

```shell
./main
```

Appel du programme non-compilé :

```shell
go run main.go
```

### 3. Afficher l'aide (-h)

Utilisation :

```shell
./main -h
```

Sortie :

```plaintext
Usage of mygb:
  -c    Clears terminal at startup
  -csv
        When set to true, write a csv file with the results
  -d string
        Path to dictionnary file (required)
  -e int
        Time before HTTP request expiration in seconds (default 10)
  -h    Display this help
  -q    When set to true, only show HTTP 200
  -t string
        Target to enumerate (required)
  -w int
        Number of workers to run (default 1)
```

### 4. Usage basique

L'usage de base permet d'envoyer des requêtes à partir d'une liste donnée sur une URL donnée.

Utilisation :

```shell
./main -t <Target URL> -d <Path to dictionnary>
```

Exemple de sortie :

```plaintext
go run main.go -t google.com -d testlist.txt 
Starting MyGB

---
Target: http://google.com/
List: testlist.txt
Workers: 1
Timeout: 10s
Quiet mode: false
CSV output: false
---

Starting scan...

/test            404
/on              200
/google          404
/com             404

Scan done in 1.306s
```

## Options

### 1. Nettoyage du terminal (-c)

L'option de nettoyage du terminal permet d'effacer votre terminal avant de démarrer le programme.

Utilisation :

```shell
./main -t <Target URL> -d <Path to dictionnary> -c
```

### 2. Définition du nombre de workers (-w)

L'option de définition du nombre de workers permet d'envoyer plusieurs requêtes simultanément afin d'accélérer le temps de travail avec les dictionnaires longs. Ce processus utilise le principe de *concurence* en Go.  
Par défaut le nombre de workers est défini à 1.

Utilisation :

```shell
./main -t <Target URL> -d <Path to dictionnary> -w <Number of workers>
```

Limites :

Lorsque vous utilisez plusieurs workers, il est possible que le délai d'expiration soit atteint même si le serveur répond dans un délai inférieur à la durée du délai d'expiration. Se référer à seection numéro 3 (Définition du délai d'expiration) pour plus d'informations

### 3. Définition du délai d'expiration (-e)

L'option de définition du délai d'expiration permet de définir, en secondes, la durée maximale d'attente pour obtenir la réponse du serveur.  

Utilisation :

```shell
./main -t <Target URL> -d <Path to dictionnary> -e <Expiration delay>
```

Limites :

Lorsque vous utilisez plusieurs workers, il est possible que le délai d'expiration soit atteint même si le serveur répond dans un délai inférieur à la durée du délai d'expiration, et cela pour deux principales raisons :

1. Saturation des ressources du système (comme les connexions réseau ou les descripteurs de fichiers). Si le nombre de connexions simultanées dépasse la capacité du système, certaines requêtes peuvent être mises en attente, ce qui peut entraîner un dépassement du délai d'expiration.

2. Le serveur peut limiter le nombre de connexions simultanées. Si toutes les connexions sont occupées, les nouvelles requêtes peuvent être mises en attente jusqu'à ce qu'une connexion se libère, ce qui peut entraîner un dépassement du délai d'expiration.

### 4. Option silence (-q)

L'option silence ou quiet permet d'afficher uniquement les requêtes ou la connexion est un succès (code HTTP 200). Cela permet de voir plus clairement les résultats positifs.

Utilisation :

```shell
./main -t <Target URL> -d <Path to dictionnary> -q
```

### 5. Sortie CSV (-csv)

L'option CSV permet de mettre dans un fichier le résultat de chaque requête au format csv. L'option CSV est aussi compatible avec l'option silence. Cela signifie que lors de l'utilisation de cette option, seulement les requêtes donnant un résultat positif seront enregistrées dans le fichier csv.  
Le fichier csv sera déposé dans le répertoire courant avec le nom `mygb_AAAAMMJJ_HHMMSS.csv`.

Utilisation :

```shell
./main -t <Target URL> -d <Path to dictionnary> -csv
```

Exemple de fichier csv :

```plaintext
word,httpcode
test,404
on,200
google,404
com,404
```

## Développement et structure

Le programme est structuré en plusieurs dossiers et fichiers pour une lisibilité et une modularité optimales :

```shell
.
├── main.go # Démarre le programme
├── basics
│   ├── flags.go # Gére et vérifie tout les paramètres
│   ├── printers.go # Imprime tout ce qui doit être affiché dans le terminal
│   └── utilities.go # Contient les fonctions annexes au programme de base (CSV, clear, ...)
└── workers
    ├── starter.go # Appel les fonctions principales qui font fonctionner le programme
    └── worker.go # Réalise le travail de requêtage
```

Le but de cette architecture est de pouvoir modifier et comprendre le code facilement afinde pouvoir ajouter d'éventuelles options supplémentaires.  
L'ensemble du code est commenté en anglais.

## Dépendances

Le programme n'a aucunes dépendances externes.  
Voici l'ensemble des imports locaux auxquels il fait référence :

```plaintext
"flag"
"os"
"strings"
"fmt"
"bufio"
"os/exec"
"time"
"net/http"
"sync"
```

## Auteur

Projet développé par LANG Kenny, dans le cadre du TP évalué de programmation en Go, Université Lyon 1 (LP ESSIR 2024-2025).
