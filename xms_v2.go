package main

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

// JSON struct for XMS REST API v2 (MediaServer 5.2 or newer)
type LicensesV2 struct {
    LicenseUsages []LicenseUsages `json:"feature_usage,omitempty"`
}

type LicenseUsages struct {
	Id 			string `json:"id,omitempty"`
	In_use 		uint64 `json:"in_use"`
	In_use_pc 	uint64 `json:"in_use_pc"`
	Free		uint64 `json:"free"`
}

type SessionsV2 struct {
    Stats SessionUsages `json:"stats"`
}

type SessionUsages struct {
	SignalingSessions 		uint64 `json:"signaling_sessions"`
	SignalingSessionsMax	uint64 `json:"signaling_sessions_max"`
	RtpSessions				uint64 `json:"rtp_sessions"`
	RtpSessionsMax			uint64 `json:"rtp_sessions_max"`
	FaxSessions 			uint64 `json:"fax_sessions"`
	FaxSessionsMax			uint64 `json:"fax_sessions_max"`
	SpeechSessions			uint64 `json:"speech_sessions"`
	SpeechSessionsMax		uint64 `json:"speech_sessions_max"`
	ConferenceSessions		uint64 `json:"conference_sessions"`
	ConferenceSessionsMax	uint64 `json:"conference_sessions_max"`
}

// ---------------------------- Fetch For XMS REST API v2
func fetchXmsV2Metrics(prefix, url string, user string, pwd string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Disable of certificate checks required for XMS in case HTTPS is used
	tr := &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		DisableKeepAlives: true,
	}
	client := http.Client{Timeout: 2 * time.Second, Transport: tr}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(user, pwd)

	resp, err := client.Do(req)
	if err != nil {
		logError("Failed to connect", err)
		setMetricValue("xms_up", 0)
		return
	}
	defer resp.Body.Close()

	setMetricValue("xms_up", 1)

	if prefix == "xms_counter" {
		processXmsV2SessionMetrics(prefix, resp)
	} else {
		processXmsV2LicenseMetrics(prefix, resp)
	}
}

func processXmsV2SessionMetrics(prefix string, resp *http.Response) {
	val := &SessionsV2{}
	decoder := json.NewDecoder(resp.Body)

	err := decoder.Decode(val)
	if err != nil {
		log.Fatal(prefix, "Failed to decode XMS response:", err)
		return
	}

	setMetricValue(prefix + "_signaling_sessions", val.Stats.SignalingSessions)
	setMetricValue(prefix + "_signaling_sessions_max", val.Stats.SignalingSessionsMax)
	setMetricValue(prefix + "_fax_sessions", val.Stats.FaxSessions)
	setMetricValue(prefix + "_fax_sessions_max", val.Stats.FaxSessionsMax)
	setMetricValue(prefix + "_rtp_sessions", val.Stats.RtpSessions)
	setMetricValue(prefix + "_rtp_sessions_max", val.Stats.RtpSessionsMax)
	setMetricValue(prefix + "_speech_sessions", val.Stats.SpeechSessions)
	setMetricValue(prefix + "_speech_sessions_max", val.Stats.SpeechSessionsMax)
	setMetricValue(prefix + "_conference_sessions", val.Stats.ConferenceSessions)
	setMetricValue(prefix + "_conference_sessions_max", val.Stats.ConferenceSessionsMax)
}

func processXmsV2LicenseMetrics(prefix string, resp *http.Response) {
	val := &LicensesV2{}
	decoder := json.NewDecoder(resp.Body)

	err := decoder.Decode(val)
	if err != nil {
		log.Fatal(prefix, "Failed to decode XMS response:", err)
		return
	}

	for _, item := range val.LicenseUsages {
		name := strings.ToLower(strings.ReplaceAll(item.Id, " ", "_"))
		basename := prefix + "_" + name

		setMetricValue(basename + "_free", item.Free)
		setMetricValue(basename + "_allocated", item.Free + item.In_use)
		setMetricValue(basename + "_total", item.Free + item.In_use)
		setMetricValue(basename + "_used", item.In_use)
		setMetricValue(basename + "_percent_used", item.In_use_pc)
	}
}