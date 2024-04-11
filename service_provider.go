package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

func parseServiceProviderCounter(line string, id string) usageCounter {
	/*
	   	"name                             current    min    max   lMin   lMax   lAvg      total",
	       "BT_ACTIVE_CALLS                       0      0      0      0      0      0          0",
	       "CENTREX_ACTIVE_CALLS                  0      0      0      0      0      0          0",
	*/
	parts := strings.Fields(line)
	if len(parts) < 8 {
		return usageCounter{}
	}
	return usageCounter{
		ID:      "0",
		Name:    normalizeMetricName(parts[0]),
		Current: parseUint64(parts[1]),
		Min:     parseUint64(parts[2]),
		Max:     parseUint64(parts[3]),
		LastMin: parseUint64(parts[4]),
		LastMax: parseUint64(parts[5]),
		LastAvg: parseUint64(parts[6]),
		Total:   parseUint64(parts[7]),
	}
}

func fetchServiceProviderCounters(prefix, url string, wg *sync.WaitGroup) {
	defer wg.Done()
	client := http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		logError("Failed to connect", err)
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)

	var counters map[string]interface{}
	error := json.Unmarshal([]byte(bodyString), &counters)
	if error != nil {
		log.Fatal(error)
	}

	re := regexp.MustCompile(`serviceProviderName: ([^"]+)`)

	var clusterInfo = counters["clusterInfo"]
	if clusterInfo == nil {
		logError("Failed to get cluster info")
		return
	}
	dc, cmpGrp := parseClusterInfo(clusterInfo.(string))
	attrs := []MetricAttribute{{"dc", dc}, {"cmpGrp", cmpGrp}}

	for key, value := range counters {
		if strings.HasPrefix(key, "spCounterTable") {
			matches := re.FindStringSubmatch(key)
			if len(matches) > 1 {
				serviceProvider := matches[1]

				for i := range value.([]interface{}) {
					line := value.([]interface{})[i].(string)
					if !strings.HasPrefix(line, "name") {
						ctr := parseServiceProviderCounter(line, serviceProvider)
						setUsageMetric(prefix, ctr, append(attrs, MetricAttribute{"sp", serviceProvider}))
					}
				}
			}
		}
	}
}
