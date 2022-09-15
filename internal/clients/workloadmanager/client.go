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

type Client interface {
	Get(ctx context.Context, path string, v interface{}) error
	Post(ctx context.Context, path string, payload interface{}, v interface{}) error
	InstanciateDB(ctx context.Context, path string, payload interface{}, v interface{}) error
}

type Options struct {
	ApiUrl			string
	Verbose			bool
}

type client struct {
	httpClient *http.Client
	options *Options
}

type database struct {
	Name 			string
	Status			string
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

func newClient(ops Options) (*client, error) {
	tr := &http.Transport{
		MaxIdleConns: 10,
		IdleConnTimeout: 30 * time.Second,
		DisableCompression: true,
	}

	httpclient := http.Client{
		Transport: tr,
	}

	return &client{
		httpClient: &httpclient,
		options: &ops,
	}, nil
}

func (c *client) instanciateDB(ctx context.Context, ops Options, db database) (InstanciateDBResponse, error) {
	var result InstanciateDBResponse
	databaseJson, err := json.Marshal(db)
	if err != nil {
		return result, fmt.Errorf("Failed to parse payload, %w", err)
	}

	resp, err := c.httpClient.Post("http://localhost:5000/api/rpc/InstanciateDb", "application/json",
		bytes.NewBuffer(databaseJson))
	if err != nil {
		return result, fmt.Errorf("failed to instanciate database %+v: %w", db, err)
	}


	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(body), &result)
    if err != nil {
        return result, fmt.Errorf("Error unmarshaling data from request.")
    }

	if result.Name == db.Name {
		return result, nil
	}

	return result, fmt.Errorf("failed to instanciate database %+v: %w", db, err)
}

func (c *client) getDB(ctx context.Context, ops Options, db database) (GetDbResponse, error) {
	var result GetDbResponse
	databaseJson, err := json.Marshal(db)
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

	if result.Name != db.Name || err != nil {
		return result, fmt.Errorf("Wrong data retrieved")
	}

	return result, nil
}

func (c *client) deleteDB(ctx context.Context, ops Options, db database) (DeleteDbResponse, error) {
	var result DeleteDbResponse 
	databaseJson, err := json.Marshal(db)
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

	if result.Name != db.Name || err != nil {
		return result, fmt.Errorf("Wrong data retrieved")
	}

	return result, nil
}

// TODO: urgent!! get base path on config file