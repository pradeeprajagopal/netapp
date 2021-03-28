package netapp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Ports struct {
	Embedded struct {
		NetappPortInventoryList []struct {
			Policies    []interface{} `json:"policies"`
			Role        string        `json:"role"`
			Utilization float64       `json:"utilization"`
			PortType    string        `json:"port_type"`
			Cluster     struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"cluster"`
			ClusterFqdn   string `json:"cluster_fqdn"`
			Speed         int    `json:"speed"`
			NumberOfHours int    `json:"number_of_hours"`
			ID            int    `json:"id"`
			Port          struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"port"`
			Node struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"node"`
			Mbps   float64 `json:"mbps"`
			Status string  `json:"status"`
		} `json:"netapp:portInventoryList"`
	} `json:"_embedded"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Curies []struct {
			Href      string `json:"href"`
			Name      string `json:"name"`
			Templated bool   `json:"templated"`
		} `json:"curies"`
	} `json:"_links"`
	TotalCount int `json:"totalCount"`
}


//GetAllPorts ... gets all port information
func GetAllPorts() (Ports,error) {
	var ports Ports
	client := Auth(server,user,password)
	newurl := client.URL + "ports?max_records=15000"
	req, err := http.NewRequest(method, newurl, nil)
	if err != nil {
		return ports,err
	}
	req.SetBasicAuth(client.UserName, client.Password)
	req.Header.Set("Accept", "application/vnd.netapp.object.inventory.performance.hal+json")
	resp, err := client.Client.Do(req)
	if err != nil {
		return ports,err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ports,err
	}
	//fmt.Println(string(bodyText))
	err = json.Unmarshal(bodyText, &ports)
	if err != nil {
		return ports,err
	}
	//return output
	return ports,nil

}
