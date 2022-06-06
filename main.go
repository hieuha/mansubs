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
	flag.StringVar(&domain, "domain", "", "domain")
	flag.BoolVar(&isCreate, "create", false, "method")
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
			if !strings.Contains(subdomain, "@") {
				target := Target{Domain: domain, Subdomain: subdomain}
				err = database.addTarget(db, target)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}

}
