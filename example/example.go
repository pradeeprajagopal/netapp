package main

import (
	"fmt"
	"os"

	"github.com/pradeeprajagopal/netapp"
)

func main() {
	//SET Environment values if not already done
	os.Setenv("USER", "netapp")
	os.Setenv("SERVER", "ocum.mynetapp.com") //your ocum address goes here
	//get all cifs information
	cifs, err := netapp.GetCifsV2()
	if err != nil {
		panic(err)
	}
	fmt.Println(cifs)

	//get all clusters information
	clusters, err := netapp.GetClustersV2()
	if err != nil {
		panic(err)
	}
	fmt.Println(clusters)
}
