package netapp

import (
	"encoding/json"
	"log"
)

type FileShares struct {
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
		AssignedPerformanceServiceLevel struct {
			ExpectedIops int    `json:"expected_iops"`
			Key          string `json:"key"`
			Name         string `json:"name"`
			PeakIops     int    `json:"peak_iops"`
		} `json:"assigned_performance_service_level"`
		AssignedStorageEfficiencyPolicy struct {
			Key  string `json:"key"`
			Name string `json:"name"`
		} `json:"assigned_storage_efficiency_policy"`
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
		Key                                string `json:"key"`
		Name                               string `json:"name"`
		RecommendedPerformanceServiceLevel struct {
			ExpectedIops int    `json:"expected_iops"`
			Key          string `json:"key"`
			Name         string `json:"name"`
			PeakIops     int    `json:"peak_iops"`
		} `json:"recommended_performance_service_level"`
		Space struct {
			Size int `json:"size"`
		} `json:"space"`
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

func getFSInfoV2(query string) (FileShares, error) {
	var results FileShares
	bodyText, err := getResponseBody(query)
	if err != nil {
		return FileShares{}, err
	}
	err = json.Unmarshal(bodyText, &results)
	if err != nil {
		log.Printf("Error: %v", err)
		return results, err
	}
	return results, nil
}

//Retrieves a list of all files shares (CIFS shares and NFS file shares).
//GetFileSharesV2 Retrieving a list of file shares
func GetFileSharesV2() (FileShares, error) {
	query := "/api/storage-provider/file-shares?max_records=50000"
	return getFSInfoV2(query)
}

//GetFileSharesFromUUIDV2 Retrieving a list of file shares using "uuid
func GetFileSharesFromUUIDV2(uuid string) (FileShares, error) {
	query := "/api/storage-provider/file-shares?uuid=" + uuid
	return getFSInfoV2(query)
}

//GetFileSharesFromNameV2 Retrieving a list of file shares using "name"
func GetFileSharesFromNameV2(name string) (FileShares, error) {
	query := "/api/storage-provider/file-shares?name=" + name
	return getFSInfoV2(query)
}

//GetFileSharesFromSVMKeyV2 Retrieving a list of file shares using "svm key"
func GetFileSharesFromSVMKeyV2(svm string) (FileShares, error) {
	query := "/api/storage-provider/file-shares?svm.key=" + svm
	return getFSInfoV2(query)
}

//GetFileSharesFromSVMClusterV2 Retrieving a list of file shares using "svm and cluster name"
func GetFileSharesFromSVMClusterV2(svm, cluster string) (FileShares, error) {
	query := "/api/storage-provider/file-shares?svm.name=" + svm + "&&cluster.name=" + cluster
	return getFSInfoV2(query)
}

