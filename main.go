package main

import (
	"bufio"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"strings"
)

func main() {
	var domain string
	var techFile string
	var techSearch string
	var isCreate bool
	var isDump bool
	var isUpdateTech bool
	var isSearchTech bool
	flag.StringVar(&domain, "domain", "", "-domain")
	flag.BoolVar(&isCreate, "create", false, "-create")
	flag.BoolVar(&isDump, "dump", false, "-dump")
	flag.BoolVar(&isUpdateTech, "update-tech", false, "-update-tech")
	flag.BoolVar(&isSearchTech, "search", false, "-search")
	flag.StringVar(&techSearch, "tech-search", "", "-tech-search")
	flag.StringVar(&techFile, "tech-file", "", "-tech-file")

	domain = strings.ToLower(domain)
	flag.Parse()

	database := Database{DatabaseSource: "targets.sqlite"}
	db, err := database.connect()
	if !database.isTableTargetExist(db) {
		err = database.createTableTargets(db)
		if err != nil {
			fmt.Println(err)
		}
	}

	if isCreate {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			subdomain := strings.ToLower(sc.Text())
			if !strings.Contains(subdomain, "@") && !strings.Contains(subdomain, "--") {
				target := Target{Domain: domain, Subdomain: subdomain}
				database.addTarget(db, target)
			}
		}
	}

	if isDump {
		targets, err := database.getTargets(db, domain)
		if err != nil {
			fmt.Println(err)
		}

		for _, target := range targets {
			fmt.Println(target.Subdomain)
		}
	}

	if isUpdateTech {
		database.cleanTech(db)
		targets := parseTech("tech.txt")
		for _, target := range targets {
			fmt.Println(target)
			err := database.updateTech(db, target.Subdomain, target.Technology)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	if isSearchTech {
		techSearch = strings.TrimSpace(strings.ToLower(techSearch))
		targets, err := database.searchTargetByTech(db, techSearch)
		if err != nil {
			fmt.Println(err)
		}
		for _, target := range targets {
			fmt.Println(target.Subdomain)
		}
	}

}
