package netapp

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Qtrees struct {
	Embedded struct {
		NetappQtreeInventoryList []struct {
			Qtree struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"qtree"`
			FilesSoftLimit int `json:"files_soft_limit"`
			Volume         struct {
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
			Svm struct {
				ID                        int           `json:"id"`
				Label                     string        `json:"label"`
				UnsupportedStatisticTypes []interface{} `json:"unsupported_statistic_types"`
				ViewID                    interface{}   `json:"view_id"`
			} `json:"svm"`
			ClusterFqdn         string      `json:"cluster_fqdn"`
			QuotaType           string      `json:"quota_type"`
			UserOrGroup         string      `json:"user_or_group"`
			DiskUsedPercentage  float64     `json:"disk_used_percentage"`
			FilesUsedPercentage interface{} `json:"files_used_percentage"`
			DiskHardLimit       float64     `json:"disk_hard_limit"`
			DiskSoftLimit       float64     `json:"disk_soft_limit"`
			ID                  int         `json:"id"`
			FilesHardLimit      float64     `json:"files_hard_limit"`
			Status              string      `json:"status"`
		} `json:"netapp:qtreeInventoryList"`
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

type QtreeV2 struct {
	links struct {
		Next struct {
			Href string `json:"href"`
		} `json:"next"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	numRecords int `json:"num_records"`
	Records    []struct {
		links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"_links"`
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
		ExportPolicy struct {
			links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
			ID   string `json:"id"`
			Key  string `json:"key"`
			Name string `json:"name"`
		} `json:"export_policy"`
		ID            int    `json:"id"`
		Key           string `json:"key"`
		Name          string `json:"name"`
		SecurityStyle string `json:"security_style"`
		Svm           struct {
			links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
			Key  string `json:"key"`
			Name string `json:"name"`
			UUID string `json:"uuid"`
		} `json:"svm"`
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
	} `json:"records"`
	totalRecords int `json:"total_records"`
}

//GetAllQtrees ...
func GetAllQtrees() (Qtrees, error) {
	var qtrees Qtrees
	client := Auth(server,user,password)
	newUrl := client.URL + "qtrees?max_records=30000"
	req, err := http.NewRequest(method, newUrl, nil)
	if err != nil {
		return qtrees, err
	}
	req.SetBasicAuth(client.UserName, client.Password)
	req.Header.Set("Accept", "application/vnd.netapp.object.inventory.capacity.hal+json")
	resp, err := client.Client.Do(req)
	if err != nil {
		return qtrees, err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return qtrees, err
	}
	//fmt.Println(string(bodyText))
	err = json.Unmarshal(bodyText, &qtrees)
	if err != nil {
		return qtrees, err
	}
	//return output

	return qtrees, nil

}


func getQtreesV2(query string) (QtreeV2, error) {
	var qtrees QtreeV2
	bodyText, err := getResponseBody(query)
	if err != nil {
		return QtreeV2{}, err
	}
	err = json.Unmarshal(bodyText, &qtrees)
	if err != nil {
		log.Printf("verita-core: Error: %v", err)
		return qtrees, err
	}
	return qtrees, nil
}

//Retrieving a list of Qtrees in the datacenter
func GetQtreesV2() (QtreeV2, error) {
	query := "/api/datacenter/storage/qtrees?max_records=50000"
	return getQtreesV2(query)
}

//Retrieving a list of Qtrees in the datacenter filtered using "name"
func GetQtreesFromNameV2(name string) (QtreeV2, error) {
	query := "/api/datacenter/storage/qtrees?name=" + strings.ToLower(name)
	return getQtreesV2(query)
}

//Retrieving a list of LUNs in the datacenter filtered using "cluster.name"
func GetQtreesFromClusterV2(name string) (QtreeV2, error) {
	query := "/api/datacenter/storage/qtrees?cluster.name=" + strings.ToLower(name)
	return getQtreesV2(query)
}

//Retrieving a list of LUNs in the datacenter filtered using "id"
func GetQtreesFromidV2(id int) (QtreeV2, error) {
	query := "/api/datacenter/qtrees?id=" + strconv.Itoa(id)
	return getQtreesV2(query)
}

//Retrieving a list of LUNs in the datacenter filtered using "key"
func GetQtreesFromKeyV2(key string) (QtreeV2, error) {
	query := "/api/datacenter/storage/qtrees?key=" + key
	return getQtreesV2(query)
}

//Retrieving a list of LUNs in the datacenter filtered using "svm.name"
func GetQtreesFromSvmNameV2(name string) (QtreeV2, error) {
	query := "/api/datacenter/storage/qtrees?svm.name=" + name
	return getQtreesV2(query)
}

//Retrieving a list of LUNs in the datacenter filtered using "svm.uuid"
func GetQtreesFromSvmUUIDV2(uuid string) (QtreeV2, error) {
	query := "/api/datacenter/storage/qtrees?svm.uuid=" + uuid
	return getQtreesV2(query)
}

//Retrieving a list of LUNs in the datacenter filtered using "volume.name"
func GetQtreesFromVolumeNameV2(name string) (QtreeV2, error) {
	query := "/api/datacenter/storage/qtrees?volume.name=" + name
	return getQtreesV2(query)
}

//Retrieving a list of LUNs in the datacenter filtered using "volume.uuid"
func GetQtreesFromVolumeUUIDV2(uuid string) (QtreeV2, error) {
	query := "/api/datacenter/storage/qtrees?volume.uuid=" + uuid
	return getQtreesV2(query)
}
