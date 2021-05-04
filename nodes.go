package netapp

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Nodes struct {
	Embedded struct {
		NetappNodeInventoryList []struct {
			Policies []struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"policies"`
			Utilization       float64 `json:"utilization"`
			Latency           float64 `json:"latency"`
			AvailableCapacity float64 `json:"available_capacity"`
			Cluster           struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"cluster"`
			ClusterFqdn            string      `json:"cluster_fqdn"`
			Iops                   float64     `json:"iops"`
			FlashCacheReadsPercent interface{} `json:"flash_cache_reads_percent"`
			NumberOfHours          int         `json:"number_of_hours"`
			UsedHeadroom           float64     `json:"used_headroom"`
			ID                     int         `json:"id"`
			Node                   struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"node"`
			TotalCapacity float64 `json:"total_capacity"`
			Mbps          float64 `json:"mbps"`
			Status        string  `json:"status"`
		} `json:"netapp:nodeInventoryList"`
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

type NodeV2 struct {
	Records []struct {
		Key      string `json:"key"`
		Name     string `json:"name"`
		UUID     string `json:"uuid"`
		Location string `json:"location"`
		Model    string `json:"model"`
		Uptime   int    `json:"uptime"`
		Cluster  struct {
			Key   string `json:"key"`
			UUID  string `json:"uuid"`
			Name  string `json:"name"`
			links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
		} `json:"cluster"`
		Version struct {
			Full       string `json:"full"`
			Generation int    `json:"generation"`
			Major      int    `json:"major"`
			Minor      int    `json:"minor"`
		} `json:"version"`
		SerialNumber string `json:"serial_number"`
		SystemID     string `json:"system_id"`
		Systemid     string `json:"systemid"`
		Ha           struct {
			Partners []struct {
				Key   string `json:"key"`
				UUID  string `json:"uuid"`
				Name  string `json:"name"`
				links struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
				} `json:"_links"`
			} `json:"partners"`
		} `json:"ha"`
		PerformanceCapacity struct {
			Used float64 `json:"used"`
		} `json:"performance_capacity"`
		Health bool `json:"health"`
		links  struct {
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

//GetAllNodes ...
//Get the performance information for all the nodes.
//
//This API can be used to query about node's health and performance data based on the mime type requested by the client.
func GetAllNodes() (Nodes, error) {
	var nodes Nodes
	client := Auth(server, user, password)
	newUrl := client.URL + "nodes?max_records=15000"
	//logger.Log(newUrl)
	req, err := http.NewRequest(method, newUrl, nil)
	if err != nil {
		return nodes, err
	}
	req.SetBasicAuth(client.UserName, client.Password)
	req.Header.Set("Accept", "application/vnd.netapp.object.inventory.performance.hal+json")
	resp, err := client.Client.Do(req)
	if err != nil {
		return nodes, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nodes, err
	}
	err = json.Unmarshal(bodyText, &nodes)
	if err != nil {
		return nodes, err
	}

	return nodes, nil

}

//getNodeInfoV2 ...
//Function performs all the http query and Unmarshall the data into Node struct
func getNodeInfoV2(query string) (NodeV2, error) {
	var results NodeV2
	bodyText, err := getResponseBody(query)
	if err != nil {
		return NodeV2{}, err
	}
	err = json.Unmarshal(bodyText, &results)
	if err != nil {
		log.Printf("verita-core: Error: %v", err)
		return NodeV2{}, err
	}
	return results, nil
}

//GetNodesV2 ...
//Retrieves all nodes. This can be paginated if the number of nodes are more
func GetNodesV2() (NodeV2, error) {
	query := "/api/datacenter/cluster/nodes"
	return getNodeInfoV2(query)
}

//GetNodesFromNameV2 ...
//Retrieves Node from OCUM based on the Name
func GetNodesFromNameV2(name string) (NodeV2, error) {
	query := "api/datacenter/cluster/nodes?name=" + name
	return getNodeInfoV2(query)

}

//GetNodesFromUUIDV2 ...
//Retrieves Node from OCUM based on the UUID property of the node
func GetNodesFromUUIDV2(uuid string) (NodeV2, error) {
	query := "api/datacenter/cluster/nodes?name=" + uuid
	return getNodeInfoV2(query)

}

//GetNodesFromClusterV2 ...
//Retrieves Node from OCUM based on the cluster.name property
func GetNodesFromClusterV2(cluster string) (NodeV2, error) {
	query := "api/datacenter/cluster/nodes?cluster.name=" + cluster
	return getNodeInfoV2(query)

}

//GetNodesFromKeyV2 ...
//Retrieves Node from OCUM based on the Key of the node
func GetNodesFromKeyV2(key string) (NodeV2, error) {
	query := "/api/datacenter/cluster/nodes/" + key
	return getNodeInfoV2(query)

}

//GetNodeBySerialNumber ...
//Retrieves Node from OCUM based on Serial number.
func GetNodeBySerialNumber(serial_number string) (NodeV2, error) {
	query := "/api/datacenter/cluster/nodes?serial_number=" + serial_number
	return getNodeInfoV2(query)
}
