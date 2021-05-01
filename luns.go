package netapp

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Luns struct {
	Embedded struct {
		NetappLunInventoryList []struct {
			Policies          []interface{} `json:"policies"`
			Latency           float64       `json:"latency"`
			AvailableCapacity float64       `json:"available_capacity"`
			Volume            struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"volume"`
			Cluster struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"cluster"`
			Aggregate struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"aggregate"`
			Svm struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"svm"`
			ClusterFqdn    string `json:"cluster_fqdn"`
			QosPolicyGroup struct {
				ID                        interface{}   `json:"id"`
				Label                     interface{}   `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"qos_policy_group"`
			Iops          float64 `json:"iops"`
			NumberOfHours float64 `json:"number_of_hours"`
			ID            int     `json:"id"`
			Node          struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"node"`
			TotalCapacity float64 `json:"total_capacity"`
			Mbps          float64 `json:"mbps"`
			Status        string  `json:"status"`
			Lun           struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"lun"`
		} `json:"netapp:lunInventoryList"`
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

type LunV2 struct {
	links struct {
		Next struct {
			Href string `json:"href"`
		} `json:"next"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	NumRecords int `json:"num_records"`
	Records    []struct {
		links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"_links"`
		Class   string `json:"class"`
		Cluster struct {
			links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
			Key  string `json:"key"`
			Name string `json:"name"`
			UUID string `json:"uuid"`
		} `json:"cluster"`
		Comment  string `json:"comment"`
		Key      string `json:"key"`
		Location struct {
			Qtree struct {
				links struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
				} `json:"_links"`
				ID   string `json:"id"`
				Key  string `json:"key"`
				Name string `json:"name"`
			} `json:"qtree"`
			Volume struct {
				links struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
				} `json:"_links"`
				Key  string `json:"key"`
				Name string `json:"name"`
				UUID string `json:"uuid"`
			} `json:"volume"`
		} `json:"location"`
		LunMaps []struct {
			Igroup struct {
				links struct {
					Self struct {
						Href string `json:"href"`
					} `json:"self"`
				} `json:"_links"`
				Key  string `json:"key"`
				Name string `json:"name"`
				UUID string `json:"uuid"`
			} `json:"igroup"`
			LogicalUnitNumber int `json:"logical_unit_number"`
		} `json:"lun_maps"`
		Name      string `json:"name"`
		OsType    string `json:"os_type"`
		QosPolicy struct {
			links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
			Key  string `json:"key"`
			Name string `json:"name"`
			UUID string `json:"uuid"`
		} `json:"qos_policy"`
		SerialNumber string `json:"serial_number"`
		Space        struct {
			Size int `json:"size"`
			Used int `json:"used"`
		} `json:"space"`
		Status struct {
			Mapped bool   `json:"mapped"`
			State  string `json:"state"`
		} `json:"status"`
		Svm struct {
			links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
			Key  string `json:"key"`
			Name string `json:"name"`
			UUID string `json:"uuid"`
		} `json:"svm"`
		UUID string `json:"uuid"`
	} `json:"records"`
	totalRecords int `json:"total_records"`
}

//GetLuns ... retrieves all Luns information
func GetLuns() (Luns, error) {
	var luns Luns
	client := Auth(server,user,password)
	newurl := client.URL + "luns?max_records=15000"
	req, err := http.NewRequest(method, newurl, nil)
	if err != nil {
		return luns, err
	}
	req.SetBasicAuth(client.UserName, client.Password)
	req.Header.Set("Accept", "application/vnd.netapp.object.inventory.performance.hal+json")
	resp, err := client.Client.Do(req)
	if err != nil {
		return luns, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return luns, err
	}
	//fmt.Println(string(bodyText))
	err = json.Unmarshal(bodyText, &luns)
	if err != nil {
		return luns, err
	}
	//return output
	return luns, nil

}

func getLunsInfoV2(query string) (LunV2, error) {
	var results LunV2
	bodyText, err := getResponseBody(query)
	if err != nil {
		return LunV2{}, err
	}
	err = json.Unmarshal(bodyText, &results)
	if err != nil {
		log.Printf("verita-core: Error: %v", err)
		return LunV2{}, err
	}
	return results, nil
}

//Retrieving a list of LUNs in the datacenter
func GetLunsV2() (LunV2, error) {
	query := "/api/datacenter/storage/luns?max_records=50000"
	return getLunsInfoV2(query)
}

//Retrieving a list of LUNs in the datacenter filtered using "name"
func GetLunsFromNameV2(name string) (LunV2, error) {
	query := "/api/datacenter/storage/luns?name=" + strings.ToLower(name)
	return getLunsInfoV2(query)
}

//Retrieving a list of LUNs in the datacenter filtered using "cluster.name"
func GetLunsFromClusterV2(name string) (LunV2, error) {
	query := "/api/datacenter/storage/luns?cluster.name=" + strings.ToLower(name)
	return getLunsInfoV2(query)
}

//Retrieving a list of LUNs in the datacenter filtered using "uuid"
func GetLunsFromuuidV2(uuid string) (LunV2, error) {
	query := "/api/datacenter/storage/luns?uuid=" + uuid
	return getLunsInfoV2(query)
}

//Retrieving a list of LUNs in the datacenter filtered using "key"
func GetLunsFromKeyV2(key string) (LunV2, error) {
	query := "/api/datacenter/storage/luns?key=" + key
	return getLunsInfoV2(query)
}

//Retrieving a list of LUNs in the datacenter filtered using "serialnumber"
func GetLunsFromSerialNumberV2(serial_number string) (LunV2, error) {
	query := "/api/datacenter/storage/luns?serial_number=" + serial_number
	return getLunsInfoV2(query)
}

//Retrieving a list of LUNs in the datacenter filtered using "svm.name"
func GetLunsFromSvmNameV2(name string) (LunV2, error) {
	query := "/api/datacenter/storage/luns?svm.name=" + name
	return getLunsInfoV2(query)
}

//Retrieving a list of LUNs in the datacenter filtered using "svm.uuid"
func GetLunsFromSvmUUIDV2(uuid string) (LunV2, error) {
	query := "/api/datacenter/storage/luns?svm.uuid=" + uuid
	return getLunsInfoV2(query)
}

//Retrieving a list of LUNs in the datacenter filtered using "volume.name"
func GetLunsFromVolumeNameV2(name string) (LunV2, error) {
	query := "/api/datacenter/storage/luns?volume.name=" + name
	return getLunsInfoV2(query)
}

//Retrieving a list of LUNs in the datacenter filtered using "volume.uuid"
func GetLunsFromVolumeUUIDV2(uuid string) (LunV2, error) {
	query := "/api/datacenter/storage/luns?volume.uuid=" + uuid
	return getLunsInfoV2(query)
}
