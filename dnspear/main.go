//a shitty Go program that uses goroutines to pull off DNS record(s
//use -d for a single domain, use -l to specify the domain file
package main

import (
	"fmt"
	"flag"
	"net"
	"os"
	"bufio"
	"sync"
)

var wg sync.WaitGroup

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
		wg.Add(1)
		go func() {
			switcher(target,recordType)
			wg.Done()
		}()
		wg.Wait()
	} else {
		openFile,_ := os.Open(targets)
		reader := bufio.NewScanner(openFile)
		reader.Split(bufio.ScanLines)

		for reader.Scan() {
			targetDomain := reader.Text()
			fmt.Println("")
			fmt.Printf("Domain: %s", targetDomain)
			wg.Add(1)
			go func() {
				switcher(reader.Text(),recordType)
				wg.Done()
			}()
			wg.Wait()
		}
	}

}