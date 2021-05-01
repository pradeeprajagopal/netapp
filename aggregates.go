package netapp

import (
	"encoding/json"
	"log"
)

type AggregateV2 struct {
	Records []struct {
		Key     string `json:"key"`
		Name    string `json:"name"` //
		UUID    string `json:"uuid"`
		Cluster struct {
			Key   string `json:"key"`
			Name  string `json:"name"` //
			UUID  string `json:"uuid"`
			links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
		} `json:"cluster"`
		Node struct {
			Key   string `json:"key"`
			UUID  string `json:"uuid"`
			Name  string `json:"name"` //
			links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
			} `json:"_links"`
		} `json:"node"`
		State        string `json:"state"` //
		BlockStorage struct {
			HybridCache struct {
				Enabled bool `json:"enabled"`
				Size    int  `json:"size"` //
			} `json:"hybrid_cache"`
			Primary struct {
				RaidSize int    `json:"raid_size"`
				RaidType string `json:"raid_type"`
			} `json:"primary"`
			Mirror struct {
				State string `json:"state"`
			} `json:"mirror"`
		} `json:"block_storage"`
		DataEncryption struct {
			SoftwareEncryptionEnabled string `json:"software_encryption_enabled"`
		} `json:"data_encryption"`
		SnaplockType string `json:"snaplock_type"`
		Space        struct {
			BlockStorage struct {
				Available int64 `json:"available"`
				Size      int64 `json:"size"`
				Used      int64 `json:"used"`
			} `json:"block_storage"`
			Efficiency struct {
				LogicalUsed int64 `json:"logical_used"`
				Savings     int64 `json:"savings"`
			} `json:"efficiency"`
		} `json:"space"`
		CreateTime          string `json:"create_time"`
		Type                string `json:"type"`
		PerformanceCapacity struct {
			Used float64 `json:"used"`
		} `json:"performance_capacity"`
		links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
		} `json:"_links"`
	} `json:"records"`
}
func getAggregatesV2(query string) (AggregateV2, error) {
	var results AggregateV2
	bodyText, err := getResponseBody(query)
	if err != nil {
		return AggregateV2{}, err
	}
	err = json.Unmarshal(bodyText, &results)
	if err != nil {
		log.Printf("verita-core: Error: %v", err)
		return AggregateV2{}, err
	}
	return results, nil
}

//GetAggregatesV2 Retrieving a list of aggregates in the cluster sorted by "name":
func GetAggregatesV2() (AggregateV2, error) {
	query := "/api/datacenter/storage/aggregates?max_records=50000"
	return getAggregatesV2(query)
}

// GetAggregateFromNameV2 Retrieving a list of aggregates in the datacenter filtered using "name":
func GetAggregateFromNameV2(name string) (AggregateV2, error) {
	query := "/api/datacenter/storage/aggregates?name=" + name
	return getAggregatesV2(query)
}

// GetAggregateFromClusterV2 Retrieving a list of aggregates in the datacenter filtered using "cluster.name"
func GetAggregateFromClusterV2(name string) (AggregateV2, error) {
	query := "/api/datacenter/storage/aggregates?cluster.name=" + name
	return getAggregatesV2(query)
}

// GetAggregateFromKeyV2 Retrieving a list of aggregates in the datacenter filtered using "key"
func GetAggregateFromKeyV2(key string) (AggregateV2, error) {
	query := "/api/datacenter/storage/aggregates?key=" + key
	return getAggregatesV2(query)
}

// GetAggregateFromStateV2 Retrieving a list of aggregates in the datacenter filtered using "state"
func GetAggregateFromStateV2(state string) (AggregateV2, error) {
	query := "/api/datacenter/storage/aggregates?state=" + state
	return getAggregatesV2(query)
}

// GetAggregateFromUUIDV2 Retrieving a list of aggregates in the datacenter filtered using uuid"
func GetAggregateFromUUIDV2(uuid string) (AggregateV2, error) {
	query := "/api/datacenter/storage/aggregates?uuid=" + uuid
	return getAggregatesV2(query)
}

// GetAggregateFromTypeV2 Retrieving a list of aggregates in the datacenter filtered using "type"
func GetAggregateFromTypeV2(t string) (AggregateV2, error) {
	query := "/api/datacenter/storage/aggregates?type=" + t
	return getAggregatesV2(query)
}

// GetAggregateFromClusterKeyV2 Retrieving a list of aggregates in the datacenter filtered using "cluster.key"
func GetAggregateFromClusterKeyV2(key string) (AggregateV2, error) {
	query := "/api/datacenter/storage/aggregates?cluster.key=" + key
	return getAggregatesV2(query)
}

// GetAggregateFromNodeKeyV2 Retrieving a list of aggregates in the datacenter filtered using "Node.key"
func GetAggregateFromNodeKeyV2(key string) (AggregateV2, error) {
	query := "/api/datacenter/storage/aggregates?node.key=" + key
	return getAggregatesV2(query)
}

// GetAggregateFromNodeNameV2 Retrieving a list of aggregates in the datacenter filtered using "Node.name"
func GetAggregateFromNodeNameV2(name string) (AggregateV2, error) {
	query := "/api/datacenter/storage/aggregates?node.name=" + name
	return getAggregatesV2(query)
}

