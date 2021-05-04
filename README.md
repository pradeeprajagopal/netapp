# netapp-ocum [![Go Reference](https://pkg.go.dev/badge/github.com/pradeeprajagopal/netapp.svg)](https://pkg.go.dev/github.com/pradeeprajagopal/netapp)

Pakcage performs all read opertion to the Netapp Active IQ Ocum V9.8. 

All Major function has option to connect to older version of the ocum and newer version of the ocum. (Function with V2 prefix)

How to import

```go
import "github.com/pradeeprajagopal/netapp"
```

Example

```go

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

```
