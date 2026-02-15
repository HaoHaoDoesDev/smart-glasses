package database

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func URL() string            { return os.Getenv("SUPABASE_URL") }
func AnonKey() string        { return os.Getenv("SUPABASE_ANON_KEY") }
func ServiceRoleKey() string { return os.Getenv("SUPABASE_SERVICE_ROLE_KEY") }

var httpClient = &http.Client{Timeout: 15 * time.Second}

func Do(method, url string, bearer string, body any) ([]byte, int, error) {
	return doWithKey(method, url, AnonKey(), bearer, body)
}

func DoService(method, url string, body any) ([]byte, int, error) {
	key := ServiceRoleKey()
	if key == "" {
		return nil, http.StatusInternalServerError, fmt.Errorf("SUPABASE_SERVICE_ROLE_KEY is not set")
	}
	return doWithKey(method, url, key, key, body)
}

func doWithKey(method, url string, apiKey string, bearer string, body any) ([]byte, int, error) {
	var buf io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, 0, err
		}
		buf = bytes.NewBuffer(b)
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, 0, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apikey", apiKey)
	req.Header.Set("Content-Profile", "public")
	req.Header.Set("Prefer", "return=representation")
	req.Header.Set("Accept-Profile", "public")

	if bearer != "" {
		req.Header.Set("Authorization", "Bearer "+bearer)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	out, _ := io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {
		return out, resp.StatusCode, fmt.Errorf("supabase error: %s", string(out))
	}

	return out, resp.StatusCode, nil
}
