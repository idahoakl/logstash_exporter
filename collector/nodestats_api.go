package collector

import (
	"encoding/json"
	"github.com/prometheus/common/log"
	"net/http"
)

type NodeStatsResponse struct {
	Host        string `json:"host"`
	Version     string `json:"version"`
	HTTPAddress string `json:"http_address"`
	Jvm         struct {
		Threads struct {
			Count     int `json:"count"`
			PeakCount int `json:"peak_count"`
		} `json:"threads"`
		Mem struct {
			HeapUsedInBytes         int `json:"heap_used_in_bytes"`
			HeapUsedPercent         int `json:"heap_used_percent"`
			HeapCommittedInBytes    int `json:"heap_committed_in_bytes"`
			HeapMaxInBytes          int `json:"heap_max_in_bytes"`
			NonHeapUsedInBytes      int `json:"non_heap_used_in_bytes"`
			NonHeapCommittedInBytes int `json:"non_heap_committed_in_bytes"`
			Pools                   struct {
				Survivor struct {
					PeakUsedInBytes  int `json:"peak_used_in_bytes"`
					UsedInBytes      int `json:"used_in_bytes"`
					PeakMaxInBytes   int `json:"peak_max_in_bytes"`
					MaxInBytes       int `json:"max_in_bytes"`
					CommittedInBytes int `json:"committed_in_bytes"`
				} `json:"survivor"`
				Old struct {
					PeakUsedInBytes  int `json:"peak_used_in_bytes"`
					UsedInBytes      int `json:"used_in_bytes"`
					PeakMaxInBytes   int `json:"peak_max_in_bytes"`
					MaxInBytes       int `json:"max_in_bytes"`
					CommittedInBytes int `json:"committed_in_bytes"`
				} `json:"old"`
				Young struct {
					PeakUsedInBytes  int `json:"peak_used_in_bytes"`
					UsedInBytes      int `json:"used_in_bytes"`
					PeakMaxInBytes   int `json:"peak_max_in_bytes"`
					MaxInBytes       int `json:"max_in_bytes"`
					CommittedInBytes int `json:"committed_in_bytes"`
				} `json:"young"`
			} `json:"pools"`
		} `json:"mem"`
		Gc struct {
			Collectors struct {
				Old struct {
					CollectionTimeInMillis int `json:"collection_time_in_millis"`
					CollectionCount        int `json:"collection_count"`
				} `json:"old"`
				Young struct {
					CollectionTimeInMillis int `json:"collection_time_in_millis"`
					CollectionCount        int `json:"collection_count"`
				} `json:"young"`
			} `json:"collectors"`
		} `json:"gc"`
	} `json:"jvm"`
	Process struct {
		OpenFileDescriptors     int `json:"open_file_descriptors"`
		PeakOpenFileDescriptors int `json:"peak_open_file_descriptors"`
		MaxFileDescriptors      int `json:"max_file_descriptors"`
		Mem                     struct {
			TotalVirtualInBytes int64 `json:"total_virtual_in_bytes"`
		} `json:"mem"`
		CPU struct {
			TotalInMillis int64 `json:"total_in_millis"`
			Percent       int   `json:"percent"`
			LoadAvg       struct {
				OneMin     float64 `json:"1m"`
				FiveMin    float64 `json:"5m"`
				FifteenMin float64 `json:"15m"`
			} `json:"load_average"`
		} `json:"cpu"`
	} `json:"process"`
	Pipeline struct {
		Events struct {
			DurationInMillis int `json:"duration_in_millis"`
			In               int `json:"in"`
			Filtered         int `json:"filtered"`
			Out              int `json:"out"`
		} `json:"events"`
		Plugins struct {
			Inputs []struct {
				ID     string `json:"id"`
				Events struct {
					In  int `json:"in"`
					Out int `json:"out"`
				} `json:"events"`
				Name string `json:"name"`
			} `json:"inputs,omitempty"`
			Filters []struct {
				ID     string `json:"id"`
				Events struct {
					DurationInMillis int `json:"duration_in_millis"`
					In               int `json:"in"`
					Out              int `json:"out"`
				} `json:"events,omitempty"`
				Name             string `json:"name"`
				Matches          int    `json:"matches,omitempty"`
				Failures         int    `json:"failures,omitempty"`
				PatternsPerField struct {
					CapturedRequestHeaders int `json:"captured_request_headers"`
				} `json:"patterns_per_field,omitempty"`
				Formats int `json:"formats,omitempty"`
			} `json:"filters"`
			Outputs []struct {
				ID     string `json:"id"`
				Events struct {
					In  int `json:"in"`
					Out int `json:"out"`
				} `json:"events"`
				Name string `json:"name"`
			} `json:"outputs"`
		} `json:"plugins"`
		Reloads struct {
			LastError            interface{} `json:"last_error"`
			Successes            int         `json:"successes"`
			LastSuccessTimestamp interface{} `json:"last_success_timestamp"`
			LastFailureTimestamp interface{} `json:"last_failure_timestamp"`
			Failures             int         `json:"failures"`
		} `json:"reloads"`
		Queue struct {
			Events   int    `json:"events,omitempty"`
			Type     string `json:"type"`
			Capacity struct {
				PageCapacityBytes int64 `json:"page_capacity_in_bytes,omitempty"`
				MaxQueueSizeBytes int64 `json:"max_queue_size_in_bytes,omitempty"`
				MaxUnreadEvents   int64 `json:"max_unread_events,omitempty"`
			} `json:"capacity"`
			Data struct {
				Path           string `json:"path,omitempty"`
				FreeSpaceBytes int64  `json:"free_space_in_bytes,omitempty"`
				StorageType    string `json:"storage_type,omitempty"`
			} `json:"data"`
		} `json:"queue"`
	} `json:"pipeline"`
}

func getNodeStats(endpoint string, target interface{}) error {
	response, err := http.Get(endpoint + "/_node/stats")
	if err != nil {
		log.Errorf("Cannot retrieve metrics: %s", err)
		return err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(target); err != nil {
		log.Errorf("Cannot parse Logstash response json: %s", err)
	}

	return nil
}

func NodeStats(endpoint string) (NodeStatsResponse, error) {
	var response NodeStatsResponse
	err := getNodeStats(endpoint, &response)

	return response, err
}
