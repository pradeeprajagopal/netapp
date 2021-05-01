package netapp

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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

type ClusterV2 struct {
	Records []struct {
		Key      string `json:"key"`
		Name     string `json:"name"`
		UUID     string `json:"uuid"`
		Contact  string `json:"contact"`
		Location string `json:"location"`
		Version  struct {
			Full       string `json:"full"`
			Generation int    `json:"generation"`
			Major      int    `json:"major"`
			Minor      int    `json:"minor"`
		} `json:"version"`
		IsSanOptimized  bool          `json:"isSanOptimized"`
		ManagementIP    string        `json:"management_ip"`
		Nodes           []ClusterNode `json:"nodes"`
		StorageCapacity struct {
			Used      int64 `json:"used"`
			Total     int64 `json:"total"`
			Available int64 `json:"available"`
		} `json:"storage_capacity"`
		links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"_links"`
	} `json:"records"`
	numRecords   int `json:"num_records"`
	totalRecords int `json:"total_records"`
	links        struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
}

type ClusterNode struct {
	Key   string `json:"key"`
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	location interface{} `json:"location"`
	version  struct {
		Full       string `json:"full"`
		Generation int    `json:"generation"`
		Major      int    `json:"major"`
		Minor      int    `json:"minor"`
	} `json:"version"`
	model        string `json:"model"`
	Uptime       int    `json:"uptime"`
	serialNumber string `json:"serial_number"`
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

func getClusterInfoV2(query string) (ClusterV2, error) {
	var results ClusterV2
	bodyText, err := getResponseBody(query)
	if err != nil {
		log.Println("Error here")
		return results, err
	}
	err = json.Unmarshal(bodyText, &results)
	if err != nil {
		log.Println("Error here3")
		return results, err
	}
	return results, nil
}


func GetClustersV2() (ClusterV2, error) {
	query := "/api/datacenter/cluster/clusters"
	return getClusterInfoV2(query)

}

func GetClusterV2(name string) (ClusterV2, error) {
	query := "/api/datacenter/cluster/clusters?name=" + strings.ToLower(name)
	return getClusterInfoV2(query)
}

func GetClusterFromKeyV2(name string) (ClusterV2, error) {
	query := "/api/datacenter/cluster/clusters?key=" + strings.ToLower(name)
	return getClusterInfoV2(query)
}

func GetClusterFromUUIDV2(uuid string) (ClusterV2, error) {
	query := "/api/datacenter/cluster/clusters?key=" + uuid
	return getClusterInfoV2(query)
}

func GetClusterFromLocationV2(location string) (ClusterV2, error) {
	query := "/api/datacenter/cluster/clusters?location=" + location
	return getClusterInfoV2(query)
}

func GetClusterFromContactV2(contact string) (ClusterV2, error) {
	query := "/api/datacenter/cluster/clusters?contact=" + contact
	return getClusterInfoV2(query)
}

func GetClusterFromManagementIPV2(management_ip string) (ClusterV2, error) {
	query := "/api/datacenter/cluster/clusters?management_ip=" + management_ip
	return getClusterInfoV2(query)
}

func GetClusterFromMajorVersionV2(version int) (ClusterV2, error) {
	query := "/api/datacenter/cluster/clusters?version.major=" + strconv.Itoa(version)
	return getClusterInfoV2(query)
}

func GetClusterFromMinorVersionV2(version int) (ClusterV2, error) {
	query := "/api/datacenter/cluster/clusters?version.minor=" + strconv.Itoa(version)
	return getClusterInfoV2(query)
}
