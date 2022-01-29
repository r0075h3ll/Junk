//a shitty Go program that pulls off DNS record(s) of a given domain using '-d'
package main

import (
	"fmt"
	"flag"
	"net"
	"sync"
)

var wg sync.WaitGroup

func arec(domain string) {
	a, _ := net.LookupIP(domain)
	defer wg.Done()

	fmt.Println("\nA record:")
	for _, ip := range a {
		fmt.Println(ip)
	}
}

func cnamerec(domain string) {
	cname, _ := net.LookupCNAME(domain)
	defer wg.Done()
	fmt.Println("\nCNAME record:")
	fmt.Println(cname)
}

func mxrec(domain string) {
	mx, _ := net.LookupMX(domain)
	defer wg.Done()
	fmt.Println("\nMX record:")

	for _, mxreco := range mx {
		fmt.Println(mxreco.Host, mxreco.Pref)
	}
}

func txtrec(domain string) {
	txt, _ := net.LookupTXT(domain)
	defer wg.Done()
	fmt.Println("\nTXT record:")
	fmt.Println(txt)
}

func main() {
	var (
		target string
		recordType string
	)
	
	flag.StringVar(&target, "d", "", "single target")
	flag.StringVar(&recordType, "r", "all", "type of record")
	
	flag.Parse()

	switch recordType {
		case "mx":
			go mxrec(target)
			wg.Add(1)
		case "a":
			go arec(target)
			wg.Add(1)
		case "cname":
			go cnamerec(target)
			wg.Add(1)
		case "txt":
			go txtrec(target)
			wg.Add(1)
		default:
			wg.Add(4)
			go mxrec(target)
			go arec(target)
			go cnamerec(target)
			go txtrec(target)
	}

	wg.Wait()
}