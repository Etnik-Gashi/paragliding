package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_igcInfo(t *testing.T) {

	// instantiate mock HTTP server (just for the purpose of testing
	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the GET request, %s", err)
	}

	resp, _ := client.Do(req)

	//check if the response from the handler is what we except
	if resp.StatusCode != 404 {
		t.Errorf("Expected StatusFound %d, received %d. ", 404, resp.StatusCode)
		return
	}

	// Testing the Status Not Implemented Yet
	req, err = http.NewRequest(http.MethodDelete, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the Delete request, %s", err)
	}

	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("Error executing the Delete request, %s", err)
	}

}

func Test_getAPI_NotImplemented(t *testing.T) {
	// instantiate mock HTTP server (just for the purpose of testing
	ts := httptest.NewServer(http.HandlerFunc(handlerAPI))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the POST request, %s", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error executing the POST request, %s", err)
	}

	//check if the response from the handler is what we except
	if resp.StatusCode != 400 {
		t.Errorf("Expected StatusNotImplemented %d, received %d. ", 400, resp.StatusCode)
		return
	}
}

func Test_getAPIIgc_NotImplemented(t *testing.T) {

	// instantiate mock HTTP server (just for the purpose of testing
	ts := httptest.NewServer(http.HandlerFunc(handlerTrack))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodDelete, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the DELETE request, %s", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error executing the DELETE request, %s", err)
	}

	//check if the response from the handler is what we except
	if resp.StatusCode != http.StatusNotImplemented {
		t.Errorf("Expected StatusNotImplemented %d, received %d. ", http.StatusNotImplemented, resp.StatusCode)
		return
	}

}

func Test_getAPIIgcId_NotImplemented(t *testing.T) {

	// instantiate mock HTTP server (just for the purpose of testing
	ts := httptest.NewServer(http.HandlerFunc(handlerID))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the POST request, %s", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error executing the POST request, %s", err)
	}

	//check if the response from the handler is what we except
	if resp.StatusCode != http.StatusNotImplemented {
		t.Errorf("Expected StatusNotImplemented %d, received %d. ", http.StatusNotImplemented, resp.StatusCode)
		return
	}

}

func Test_getAPIIgcField_NotImplemented(t *testing.T) {

	// instantiate mock HTTP server (just for the purpose of testing
	ts := httptest.NewServer(http.HandlerFunc(handlerField))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the POST request, %s", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error executing the POST request, %s", err)
	}

	//check if the response from the handler is what we except
	if resp.StatusCode != 400 {
		t.Errorf("Expected StatusNotImplemented %d, received %d. ", http.StatusNotImplemented, resp.StatusCode)
		return
	}

}

func Test_webhookNewTrack_NotImplemented(t *testing.T) {

	// instantiate mock HTTP server (just for the purpose of testing
	ts := httptest.NewServer(http.HandlerFunc(webhookNewTrack))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the POST request, %s", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error executing the POST request, %s", err)
	}

	//check if the response from the handler is what we except
	if resp.StatusCode != 200 {
		t.Errorf("Expected StatusNotImplemented %d, received %d. ", 200, resp.StatusCode)
		return
	}

}

func Test_getAPI_MalformedURL(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(handlerAPI))
	defer ts.Close()

	testCases := []string{
		ts.URL,
		ts.URL + "/something/",
		ts.URL + "/something/123/",
	}

	for _, tstring := range testCases {
		resp, err := http.Get(ts.URL)
		if err != nil {
			t.Errorf("Error making the GET request, %s", err)
		}

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("For route: %s, expected StatusCode %d, received %d. ", tstring, http.StatusBadRequest, resp.StatusCode)
			return
		}
	}
}

func Test_getAPIIgc_MalformedURL(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(handlerTrack))
	defer ts.Close()

	testCases := []string{
		ts.URL,
		ts.URL + "/something/",
		ts.URL + "/something/123/",
	}

	for _, tstring := range testCases {
		resp, err := http.Get(ts.URL)
		if err != nil {
			t.Errorf("Error making the GET request, %s", err)
		}

		if resp.StatusCode != 200 {
			t.Errorf("For route: %s, expected StatusCode %d, received %d. ", tstring, http.StatusBadRequest, resp.StatusCode)
			return
		}
	}
}

func Test_getAPIIgcId_MalformedURL(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(handlerID))
	defer ts.Close()

	testCases := []string{
		ts.URL,
		ts.URL + "/something/",
		ts.URL + "/something/123/",
	}

	for _, tstring := range testCases {
		resp, err := http.Get(ts.URL)
		if err != nil {
			t.Errorf("Error making the GET request, %s", err)
		}

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("For route: %s, expected StatusCode %d, received %d. ", tstring, http.StatusBadRequest, resp.StatusCode)
			return
		}
	}
}

func Test_getAPIIgcField_MalformedURL(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(handlerField))
	defer ts.Close()

	testCases := []string{
		ts.URL,
		ts.URL + "/something/",
		ts.URL + "/something/123/",
	}

	for _, tstring := range testCases {
		resp, err := http.Get(ts.URL)
		if err != nil {
			t.Errorf("Error making the GET request, %s", err)
		}

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("For route: %s, expected StatusCode %d, received %d. ", tstring, http.StatusBadRequest, resp.StatusCode)
			return
		}
	}
}

func Test_getAPIIgc_Post(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(handlerTrack))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	apiURLTest := _url{}
	apiURLTest.URL = "http://skypolaris.org/wp-content/uploa/IGS%20Files/Madrid%20to%20Jerez.igc"

	jsonData, _ := json.Marshal(apiURLTest)

	req, err := http.NewRequest("POST", ts.URL, bytes.NewBuffer(jsonData))
	if err != nil {
		t.Errorf("Error making the POST request, %s", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error executing the POST request, %s", err)
	}

	if resp.StatusCode == 400 {
		assert.Equal(t, 400, resp.StatusCode, "OK response is expected")
	} else {
		assert.Equal(t, 200, resp.StatusCode, "OK response is expected")
	}

}

func Test_getNewWebhook(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(webhookNewTrack))
	defer ts.Close()

	client := &http.Client{}

	urlTest := Webhook{}
	urlTest.WebhookURL = "https://discordapp.com:8080/api/webhooks/504970988400803842/cyPUQQw0laWWVSikkV-cwvZKv97xyUkbi-2aDX2fJZccJYmORHOknS155L2lUX3_LPlM"

	jsonData, _ := json.Marshal(urlTest)

	request, err := http.NewRequest("POST", ts.URL, bytes.NewBuffer(jsonData))
	if err != nil {
		t.Errorf("Error making the POST request, %s", err)
	}

	resp, err := client.Do(request)
	if err != nil {
		t.Errorf("Error executing the POST request, %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected StatusOK %d, received %d. ", http.StatusOK, resp.StatusCode)
		return
	}

	if resp.StatusCode == 400 {
		assert.Equal(t, 400, resp.StatusCode, "OK response is expected")
	} else {
		assert.Equal(t, 200, resp.StatusCode, "OK response is expected")
	}

}

func Test_getNewWebhookID(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(webhookNewTrack))
	defer ts.Close()

	client := &http.Client{}

	urlTest := Webhook{}
	urlTest.WebhookID = "0"

	jsonData, _ := json.Marshal(urlTest)

	request, err := http.NewRequest("POST", ts.URL, bytes.NewBuffer(jsonData))
	if err != nil {
		t.Errorf("Error making the POST request, %s", err)
	}

	resp, err := client.Do(request)
	if err != nil {
		t.Errorf("Error executing the POST request, %s", err)
	}

	if resp.StatusCode == 400 {
		assert.Equal(t, 400, resp.StatusCode, "OK response is expected")
	} else {
		assert.Equal(t, 200, resp.StatusCode, "OK response is expected")
	}

}

func Test_getAPIIgc_Post_Empty(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(handlerTrack))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	apiURLTest := _url{}
	apiURLTest.URL = ""

	jsonData, _ := json.Marshal(apiURLTest)

	req, err := http.NewRequest("POST", ts.URL, bytes.NewBuffer(jsonData))
	if err != nil {
		t.Errorf("Error making the POST request, %s", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error executing the POST request, %s", err)
	}

	assert.Equal(t, 200, resp.StatusCode, "OK response is expected")

}

////Webhook tests

func Test_webhookNewTrack(t *testing.T) {
	// instantiate mock HTTP server (just for the purpose of testing
	ts := httptest.NewServer(http.HandlerFunc(webhookNewTrack))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the POST request, %s", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error executing the POST request, %s", err)
	}

	//check if the response from the handler is what we except
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected StatusOK %d, received %d. ", http.StatusOK, resp.StatusCode)
		return
	}

	req, err = http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the GET request, %s", err)
	}

	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("Error executing the GET request, %s", err)
	}

	//check if the response from the handler is what we except
	if resp.StatusCode != http.StatusNotImplemented {
		t.Errorf("Expected StatusNotImplemented %d, received %d. ", http.StatusNotImplemented, resp.StatusCode)
		return
	}

}

func Test_webhookID(t *testing.T) {
	// instantiate mock HTTP server (just for the purpose of testing
	ts := httptest.NewServer(http.HandlerFunc(webhookID))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the POST request, %s", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error executing the POST request, %s", err)
	}

	//check if the response from the handler is what we except
	if resp.StatusCode != http.StatusNotImplemented {
		t.Errorf("Expected StatusNotImplemented %d, received %d. ", http.StatusNotImplemented, resp.StatusCode)
		return
	}
	if resp.StatusCode != 501 {
		t.Errorf("Expected StatusOK %d, received %d. ", 501, resp.StatusCode)
		return
	}

}

////Admin tests

func Test_adminAPITracksCount(t *testing.T) {
	// instantiate mock HTTP server (just for the purpose of testing
	ts := httptest.NewServer(http.HandlerFunc(adminAPITracksCount))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the GET request, %s", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error executing the GET request, %s", err)
	}

	//check if the response from the handler is what we except
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected StatusOK %d, received %d. ", http.StatusOK, resp.StatusCode)
		return
	}

	req, err = http.NewRequest(http.MethodPost, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the POST request, %s", err)
	}

	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("Error executing the POST request, %s", err)
	}

	//check if the response from the handler is what we except
	if resp.StatusCode != http.StatusNotImplemented {
		t.Errorf("Expected StatusNotImplemented %d, received %d. ", http.StatusNotImplemented, resp.StatusCode)
		return
	}

}

func Test_adminAPITracks(t *testing.T) {
	// instantiate mock HTTP server (just for the purpose of testing
	ts := httptest.NewServer(http.HandlerFunc(adminAPITracks))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodDelete, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the DELETE request, %s", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error executing the DELETE request, %s", err)
	}

	//check if the response from the handler is what we except
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected StatusOK %d, received %d. ", http.StatusOK, resp.StatusCode)
		return
	}

	req, err = http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the GET request, %s", err)
	}

	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("Error executing the GET request, %s", err)
	}

	//check if the response from the handler is what we except
	if resp.StatusCode != http.StatusNotImplemented {
		t.Errorf("Expected StatusNotImplemented %d, received %d. ", http.StatusNotImplemented, resp.StatusCode)
		return
	}

}

///Ticker tests

func Test_getAPITickerLatest(t *testing.T) {
	// instantiate mock HTTP server (just for the purpose of testing
	ts := httptest.NewServer(http.HandlerFunc(getApiTickerLatest))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the GET request, %s", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error executing the GET request, %s", err)
	}

	//check if the response from the handler is what we except
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected StatusOK %d, received %d. ", http.StatusOK, resp.StatusCode)
		return
	}

	req, err = http.NewRequest(http.MethodPost, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the POST request, %s", err)
	}

	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("Error executing the POST request, %s", err)
	}

	//check if the response from the handler is what we except
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected StatusNotFound %d, received %d. ", http.StatusNotFound, resp.StatusCode)
		return
	}

}

func Test_getAPITicker(t *testing.T) {
	// instantiate mock HTTP server (just for the purpose of testing
	ts := httptest.NewServer(http.HandlerFunc(getApiTicker))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the GET request, %s", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error executing the GET request, %s", err)
	}

	//check if the response from the handler is what we except
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected StatusOK %d, received %d. ", http.StatusOK, resp.StatusCode)
		return
	}

	req, err = http.NewRequest(http.MethodPost, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the POST request, %s", err)
	}

	resp, err = client.Do(req)
	if err != nil {
		t.Errorf("Error executing the POST request, %s", err)
	}

	//check if the response from the handler is what we except
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected StatusNotFound %d, received %d. ", http.StatusNotFound, resp.StatusCode)
		return
	}

}

func Test_getAPITickerTimestamp(t *testing.T) {
	// instantiate mock HTTP server (just for the purpose of testing
	ts := httptest.NewServer(http.HandlerFunc(getApiTickerTimestamp))
	defer ts.Close()

	//create a request to our mock HTTP server
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Errorf("Error constructing the GET request, %s", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Error executing the GET request, %s", err)
	}

	//check if the response from the handler is what we except
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected StatusBadRequest %d, received %d. ", http.StatusBadRequest, resp.StatusCode)
		return
	}

}

/////////////////Other testing functions

func Test_latestTimestamp(t *testing.T) {
	igcTracks := []tracks{
		tracks{TimeRecorded: time.Date(2018, 4, 25, 12, 32, 1, 0, time.UTC)},
		tracks{TimeRecorded: time.Now()},
		tracks{TimeRecorded: time.Date(2019, 4, 25, 12, 32, 1, 0, time.UTC)},
	}

	latestTS := latestTimestamp(igcTracks)
	if latestTS != igcTracks[2].TimeRecorded {
		t.Error("Not the latest timestamp")
	}
}

func Test_oldestTimestamp(t *testing.T) {
	igcTracks := []tracks{
		tracks{TimeRecorded: time.Date(2018, 4, 25, 12, 32, 1, 0, time.UTC)},
		tracks{TimeRecorded: time.Now()},
		tracks{TimeRecorded: time.Date(2019, 4, 25, 12, 32, 1, 0, time.UTC)},
	}

	oldestTS := oldestTimestamp(igcTracks)
	if oldestTS != igcTracks[0].TimeRecorded {
		t.Error("Not the oldest timestamp")
	}
}

func Test_oldestNewerTimestamp(t *testing.T) {
	igcTracks := []tracks{
		tracks{TimeRecorded: time.Date(2018, 4, 25, 12, 32, 1, 0, time.UTC)},
		tracks{TimeRecorded: time.Date(2018, 4, 26, 12, 32, 1, 0, time.UTC)},
		tracks{TimeRecorded: time.Date(2019, 4, 25, 12, 32, 1, 0, time.UTC)},
	}

	oldestNewTS := oldestNewerTimestamp("25.04.2018 12:34:30.314", igcTracks)

	if oldestNewTS != igcTracks[1].TimeRecorded {
		t.Error("Not the right timestamp")
	}
}

func Test_tickerTimestamps(t *testing.T) {
	igcTracks := []tracks{
		tracks{TimeRecorded: time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)},
	}

	// No connection to the DB :(

	tickerTS := tickerTimestamps("25.04.2018 12:34:30.314")

	if tickerTS.oldestTimestamp != igcTracks[0].TimeRecorded {
		t.Error("Not the right timestamp")
	}
	if tickerTS.oldestNewerTimestamp != igcTracks[0].TimeRecorded {
		t.Error("Not the right timestamp")
	}
	if tickerTS.latestTimestamp != igcTracks[0].TimeRecorded {
		t.Error("Not the right timestamp")
	}
}

/////Mongo tests

func Test_mongoConnect(t *testing.T) {
	if conn := mongoConnect(); conn == nil {
		t.Error("No connection")
	}
}

func Test_urlInMong(t *testing.T) {
	urlExists := urlInMongo(`Some random URL`, mongoConnect().Database("paragliding").Collection("track"))
	if urlExists {
		t.Error("Track should not exist")
	}
}

func Test_getAllTrack(t *testing.T) {
	allTracks := getAllTracks(mongoConnect())

	if len(allTracks) < 0 {
		t.Error("It should be bigger")
	}
}

func Test_getAllWebhooks(t *testing.T) {
	allWebhooks := getAllWebhooks(mongoConnect())

	if len(allWebhooks) < 0 {
		t.Error("It should be bigger")
	}
}

func Test_getTrack(t *testing.T) {
	track := getTrack(mongoConnect(), `url`)

	if track.URL != "" {
		t.Error("It should be empty")
	}
}

func Test_deleteWebhook(t *testing.T) {
	deleteWebhook(mongoConnect(), `noWebhook`)
}
