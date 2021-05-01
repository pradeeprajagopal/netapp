package netapp

import (
	"encoding/json"
	"log"
	"strings"
)

type Cifsv2 struct {
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
		Acls []struct {
			Permission  string `json:"permission"`
			UserOrGroup string `json:"user_or_group"`
		} `json:"acls"`
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
		comment string `json:"comment"`
		Key     string `json:"key"`
		Name    string `json:"name"`
		Path    string `json:"path"`
		Svm     struct {
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

func getCifsV2(query string) (Cifsv2, error) {
	var results Cifsv2

	bodyText, err := getResponseBody(query)
	if err != nil {
		return Cifsv2{}, err
	}
	err = json.Unmarshal(bodyText, &results)
	if err != nil {
		log.Printf("verita-core: Error: %v", err)
		return Cifsv2{}, err
	}
	return results, nil
}

//GetCifsV2 Retrieves CIFS shares.
func GetCifsV2() (Cifsv2, error) {
	query := "/api/datacenter/protocols/cifs/shares?max_records=50000"
	return getCifsV2(query)
}

//Retrieving a list of CIFS shares in the datacenter filtered using "cluster.name
func GetCifsFromClusterV2(name string) (Cifsv2, error) {
	query := "/api/datacenter/protocols/cifs/shares?cluster.name=" + strings.ToLower(name)
	return getCifsV2(query)

}

//Retrieving a list of CIFS shares in the datacenter filtered using "svm.name
func GetCifsFromSvmV2(name string) (Cifsv2, error) {
	query := "/api/datacenter/protocols/cifs/shares?svm.name=" + strings.ToLower(name)
	return getCifsV2(query)
}

//Retrieving a CIFS share using the specified key:
func GetCifsNameV2(key string) (Cifsv2, error) {
	query := "/api/datacenter/protocols/cifs/shares/" + strings.ToLower(key)
	return getCifsV2(query)

}

//Retrieving a CIFS share using the specified path:
func GetCifsNameFromPathV2(path string) (Cifsv2, error) {
	query := "/api/datacenter/protocols/cifs/shares?path=" + path
	return getCifsV2(query)

}

//Retrieving a CIFS share using the specified volume name:
func GetCifsNameFromVolumeV2(name string) (Cifsv2, error) {
	query := "/api/datacenter/protocols/cifs/shares?volume.name=" + name
	return getCifsV2(query)

}

