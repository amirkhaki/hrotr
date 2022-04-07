package main

func (lk *liker) checkToken() (valid bool,err error) {
	if lk.token == "" {
		return
	}
	req, err := http.NewRequest("GET", API_ROOT+"auth/checktoken")
	if err != nil {
		err = fmt.Errorf("checktoken: Could not create request: %w", err)
		return
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("user-agent", "okhttp/3.12.1")
	req.Header.Add("authorization", "Bearer "+lk.token)
	resp, err := lk.cl.Do(req)
	if err != nil {
		err = fmt.Errorf("checktoken: Could not send request: %w", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("checktoken: Could not read resp body: %w", err)
		return
	}
	var validToken struct {
		Success bool `json:"success"`
		User any `json:"user"`
		Error string `json:"error"`
	}
	err = json.Unmarshal(body, &validToken)
	if err != nil {
		err = fmt.Errorf("checktoken: Could not unmarshal resp body: %w", err)
		return
	}
	if !validToken.Success {
		err = fmt.Errorf("token was not valid: %s", validToken.Error)
		return
	}
	valid = validToken.Success
	return

}
