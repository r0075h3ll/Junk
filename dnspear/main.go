//a shitty subdomain bruteforcer
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
		bruteForce bool
		domCom []string
	)
	
	flag.StringVar(&target, "d", "", "single target")
	flag.StringVar(&targets, "l", "", "single target")
	flag.StringVar(&recordType, "r", "all", "type of record")
	flag.BoolVar(&bruteForce, "b", false, "brute-force for subdomains")
	
	flag.Parse()

	if target == "" && targets == "" {
		fmt.Println("\n\tdnspear\n")
		flag.PrintDefaults()
	}

	bruteForcer := func(target string) {
		openFile,_ := os.Open("list.txt") //https://gist.github.com/jhaddix/86a06c5dc309d08580a018c66354a056
		reader := bufio.NewScanner(openFile)
		reader.Split(bufio.ScanLines)

		for reader.Scan() {
			newDomain := reader.Text() + target
			domCom = append(domCom,newDomain)
		}

		fmt.Println("\n[!] Target: ", target)

		for _,val := range domCom {
			_,err := net.LookupIP(val)
			if err != nil {
				fmt.Println("[-] Domain not found: ", val)
			} else {
				fmt.Println("[+] Domain found: ", val)
			}
		}
	}

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

	if target != "" && bruteForce == false {
		wg.Add(1)
		go func() {
			switcher(target,recordType)
			wg.Done()
		}()
		wg.Wait()
	} else if targets != "" && bruteForce == false {
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
	} else if bruteForce == true {

		if targets != "" {
			openFile,_ := os.Open(targets)
			reader := bufio.NewScanner(openFile)
			reader.Split(bufio.ScanLines)

			for reader.Scan() {
				wg.Add(1)
				go func() {
					defer wg.Done()
					bruteForcer(reader.Text())
				}()
				wg.Wait()
			}
		} else {
			wg.Add(1)
			go func() {
				defer wg.Done()
				bruteForcer(target)
			}()
			wg.Wait()
		}
	}

}