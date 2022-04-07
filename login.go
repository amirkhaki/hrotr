package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"bytes"
	"io"
)
const API_ROOT = "https://virgool.io/api/v1.1/"
type liker struct {
	cl    *http.Client
	token string
}

func (lk *liker) login(username, password string) error {
	credentials := map[string]string{"username":username, "password":password}
	data, err := json.Marshal(credentials)
	if err != nil {
		return fmt.Errorf("login: Could not marshal credentials: %w", err)
	}
	req, err := http.NewRequest("POST", API_ROOT+"login", bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("login: Could not create request: %w", err)
	}
	req.Header.Add("content-type","application/json")
	req.Header.Add("user-agent","okhttp/3.12.1")
	resp, err := lk.cl.Do(req)
	if err != nil {
		return fmt.Errorf("login: Could not send request: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("login: Could not read resp body: %w", err)
	}
	var token struct {
		Success bool `json:"success"`
		Msg string `json:"msg"`
		Token string `json:"token"`
		User any `json:"user"`
	}
	err = json.Unmarshal(body, &token)
	if err != nil {
		return fmt.Errorf("login: Could not unmarshal resp body %s: %w",string(body), err)
	}
	if !token.Success {
		return fmt.Errorf("login: unsuccessful: %s %s", token.Msg, string(body))
	}
	lk.token = token.Token
	return nil
}

func New() (*liker, error) {
	client := &http.Client{}
	lk := liker{cl:client}
	return &lk, nil
}


func main() {
	lk, err := New()
	if err != nil {
		panic(err)
	}
	if err = lk.login("m_e", "AmIrReZa1383"); err != nil {
		panic(err)
	}
	fmt.Println(lk)
}
