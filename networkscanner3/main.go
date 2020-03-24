// nmap sample, nmap installed your pc
// https://github.com/Ullaakut/nmap
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Ullaakut/nmap"
)

func main() {
	target := "192.168.0.1/24"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	scanner, err := nmap.NewScanner(
		nmap.WithTargets(target),
		nmap.WithPorts("80,443"),
		nmap.WithContext(ctx),
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v\n", err)
	}

	result, warnings, err := scanner.Run()
	if err != nil {
		log.Fatalf("unable to run nmap scanner: %v\n", err)
	}
	if warnings != nil {
		log.Printf("Warnings: \n%v\n", warnings)
	}

	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		fmt.Printf("Host %q:\n", host.Addresses[0])
		for _, port := range host.Ports {
			fmt.Printf("\tPort %d/%s %s %s", port.ID, port.Protocol, port.State, port.Service.Name)
		}
	}
	fmt.Printf("Nmap done: %d hosts up scanned in %3f seconds\n", len(result.Hosts), result.Stats.Finished.Elapsed)
}
