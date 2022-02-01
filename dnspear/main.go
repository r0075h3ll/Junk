//a shitty Go program that pulls off DNS record(s) of a given domain using '-d'
package main

import (
	"fmt"
	"flag"
	"net"
	"os"
	"bufio"
)


func arec(domain string) {
	a, _ := net.LookupIP(domain)

	fmt.Println("\nA record:")
	for _, ip := range a {
		fmt.Println(ip)
	}
}

func cnamerec(domain string) {
	cname, _ := net.LookupCNAME(domain)
	fmt.Println("\nCNAME record:")
	fmt.Println(cname)
}

func mxrec(domain string) {
	mx, _ := net.LookupMX(domain)
	fmt.Println("\nMX record:")

	for _, mxreco := range mx {
		fmt.Println(mxreco.Host, mxreco.Pref)
	}
}

func txtrec(domain string) {
	txt, _ := net.LookupTXT(domain)
	fmt.Println("\nTXT record:")
	fmt.Println(txt)
}

func main() {
	var (
		target,targets string
		recordType string
	)
	
	flag.StringVar(&target, "d", "", "single target")
	flag.StringVar(&targets, "l", "", "single target")
	flag.StringVar(&recordType, "r", "all", "type of record")
	
	flag.Parse()

	switcher := func(domain,record string) {
			switch record {
				case "mx":
					mxrec(domain)
				case "a":
					arec(domain)
				case "cname":
					cnamerec(domain)
				case "txt":
					txtrec(domain)
				default:
					mxrec(domain)
					arec(domain)
					cnamerec(domain)
					txtrec(domain)
			}
		}

	if target != "" {
		switcher(target,recordType)
	} else {
		openFile,_ := os.Open(targets)
		reader := bufio.NewScanner(openFile)
		reader.Split(bufio.ScanLines)

		for reader.Scan() {
			fmt.Println("\n")
			fmt.Printf("\033[91mDOMAIN:\033[00m %s", reader.Text())
			switcher(reader.Text(),recordType)
		}
	}

}