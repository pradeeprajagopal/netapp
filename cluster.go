package netapp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Clusters struct {
	Embedded struct {
		NetappClusterInventoryList []struct {
			NodeCount         int           `json:"node_count"`
			Policies          []interface{} `json:"policies"`
			NetworkAddress    string        `json:"network_address"`
			AvailableCapacity float64       `json:"available_capacity"`
			Cluster           struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"cluster"`
			ClusterFqdn   string  `json:"cluster_fqdn"`
			Iops          float64 `json:"iops"`
			NumberOfHours int     `json:"number_of_hours"`
			ID            int     `json:"id"`
			TotalCapacity float64 `json:"total_capacity"`
			Serial        string  `json:"serial"`
			OsVersion     string  `json:"os_version"`
			Mbps          float64 `json:"mbps"`
			Status        string  `json:"status"`
		} `json:"netapp:clusterInventoryList"`
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

//GetClusters connects to ocum and gets all the cluster information
func GetClusters() (Clusters,error) {
	var clusters Clusters
	client := Auth(server,user,password)
	newurl := client.URL + "clusters?max_records=1000"
	req, err := http.NewRequest(method, newurl, nil)
	if err != nil {
		return clusters,err
	}
	req.SetBasicAuth(client.UserName, client.Password)
	req.Header.Set("Accept", "application/vnd.netapp.object.inventory.performance.hal+json")
	resp, err := client.Client.Do(req)
	if err != nil {
		return clusters,err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return clusters,err
	}
	err = json.Unmarshal(bodyText, &clusters)
	if err != nil {
		return clusters,err
	}
	return clusters,nil

}
