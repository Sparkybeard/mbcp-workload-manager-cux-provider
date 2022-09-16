package workloadmanagerclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Options struct {
	ApiUrl			string
	Verbose			bool
}

type Client struct {
	httpClient *http.Client
	options *Options
}

type InstanciateDBResponse struct {
	Name			string
	Id				string
}

type GetDbResponse struct {
	Name			string
	Id				string
}

type DeleteDbResponse struct {
	Name			string
	Id				string
}

func NewClient(ops Options) (*Client, error) {
	tr := &http.Transport{
		MaxIdleConns: 10,
		IdleConnTimeout: 30 * time.Second,
		DisableCompression: true,
	}

	httpclient := http.Client{
		Transport: tr,
	}

	return &Client{
		httpClient: &httpclient,
		options: &ops,
	}, nil
}

func (c *Client) InstanciateDB(ctx context.Context, dbName string) (InstanciateDBResponse, error) {
	var result InstanciateDBResponse
	databaseJson, err := json.Marshal(dbName)
	if err != nil {
		return result, fmt.Errorf("Failed to parse payload, %w", err)
	}

	resp, err := c.httpClient.Post("http://localhost:5000/api/rpc/InstanciateDb", "application/json",
		bytes.NewBuffer(databaseJson))
	if err != nil {
		return result, fmt.Errorf("failed to instanciate database %+v: %w", dbName, err)
	}


	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(body), &result)
    if err != nil {
        return result, fmt.Errorf("Error unmarshaling data from request.")
    }

	if result.Name == dbName {
		return result, nil
	}

	return result, fmt.Errorf("failed to instanciate database %+v: %w", dbName, err)
}

func (c *Client) GetDB(ctx context.Context, dbName string) (GetDbResponse, error) {
	var result GetDbResponse
	databaseJson, err := json.Marshal(dbName)
	if err != nil {
		return result, fmt.Errorf("failed to parse payload, %w", err)
	}
	
	resp, err := c.httpClient.Post("http://localhost:5000/api/rpc/GetDb", "application/json",
		bytes.NewBuffer(databaseJson))
	if err != nil {
		return result, fmt.Errorf("failed to retrieve database data") 
	}
	
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(body), &result)

	if result.Name != dbName || err != nil {
		return result, fmt.Errorf("Wrong data retrieved")
	}

	return result, nil
}

func (c *Client) DeleteDB(ctx context.Context, dbName string) (DeleteDbResponse, error) {
	var result DeleteDbResponse 
	databaseJson, err := json.Marshal(dbName)
	if err != nil {
		return result, fmt.Errorf("failed to parse payload, %w", err)
	}
	
	resp, err := c.httpClient.Post("http://localhost:5000/api/rpc/DeleteDb", "application/json",
		bytes.NewBuffer(databaseJson))
	if err != nil {
		return result, fmt.Errorf("failed to retrieve database data") 
	}
	
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(body), &result)

	if result.Name != dbName || err != nil {
		return result, fmt.Errorf("Wrong data retrieved")
	}

	return result, nil
}

// TODO: urgent!! get base path on config file