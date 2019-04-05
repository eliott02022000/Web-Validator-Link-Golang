package main

import (  
    "fmt"
	"os"
	"regexp"
	"net/http"
	"io/ioutil"
	"strings"
	

)
var lien []string

var done map[string]bool

func main() {

	done = make(map[string]bool)
	// recupere les args
	argument := os.Args[1]
	e := url(argument)
// rappel
	check(e, argument)

	// creation du txt
	creative, err := os.Create("final.txt")
    if err != nil {         
        fmt.Println("File reading error", err)
    }
	defer  creative.Close()
	// boucle recuperer et mettre les valeurs fichier
	for i := 0 ; i < len(e); i++ {
		fmt.Println(e[i])
		fmt.Fprintln(creative, string(e[i]))
	}
	// afficher
	for i := 0 ; i < len(lien); i++ {
		fmt.Fprintln(creative, lien[i])
	}
	
}

func url(url string) []string {

	// recuperer url
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	} 
		
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	// regex http et href
	re := regexp.MustCompile(`(http|ftp|https)://([\w_-]+(?:(?:\.[\w_-]+)+))([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?`)
	href := regexp.MustCompile(`"/([\w-.,@?^=%&amp;:/~+#]*[\w-@?^=%&amp;/~+#])+"`)

	// recup le body
	e := re.FindAllString(string(robots), -1)
	h := href.FindAllString(string(robots), -1)

	// enlever a Gauche
	for idx, names := range h {
		h[idx] = strings.TrimLeft(names, "\"")
	}
	// enlever a droite
	for idx, names := range h {
		h[idx] = strings.TrimRight(names, "\"")
	}

	// reel de base et href
	for idx, names := range h {
		h[idx] = url + names
	}
	// associe 
	eh := append(e,h...)
	return eh
}

func check(check []string, urll string) {

	if 	done[urll] == true {
		return
	}

	for idx := 0; idx < len(check); idx++ {

			if done[urll] == false {

				name := url(check[idx])
				for _, names := range name {
					fmt.Println(urll, string(test(names)))
					lien = append(lien, check[idx])
				}
				done[check[idx]] = true
			}
		}
	
	done[urll] = true
}

func test(url string)string {
    resp, err := http.Get(url)
    panicc(err)
    status := ""
	fmt.Println(resp.StatusCode)
    if resp.StatusCode == 200 {
        status="->Ok"
    } else {
        status ="->error 404"
    }

    return status
}

func panicc(e error) {
    if e != nil {
        panic(e)
    }
}