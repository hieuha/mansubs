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
	var isCreate bool
	var isDump bool
	var isUpdateTech bool
	flag.StringVar(&domain, "domain", "", "-domain")
	flag.BoolVar(&isCreate, "create", false, "-create")
	flag.BoolVar(&isDump, "dump", false, "-dump")
	flag.BoolVar(&isUpdateTech, "update-tech", false, "-update-tech")
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
		database.getTargets(db, domain)
	}

	if isUpdateTech {
		targets, err := database.getTargets(db, domain)
		if err != nil {
			fmt.Println(err)
		}
		for _, target := range targets {
			err := database.updateTech(db, target.Id, "Web")
			if err != nil {
				fmt.Println(err)
			}
		}
	}

}
