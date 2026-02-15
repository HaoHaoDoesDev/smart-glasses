package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	apiKey := os.Getenv("GEMINI_API_KEY")

	http.HandleFunc("/verify", func(w http.ResponseWriter, r *http.Request) {
		prompt := r.URL.Query().Get("prompt")
		var data struct {
			Image string `json:"image"`
		}

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			log.Printf("Decode Error: %v", err)
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		success, err := VerifyWithGemini(apiKey, data.Image, prompt)
		if err != nil || !success {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprint(w, "AI Verification Failed")
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Verified")
	})

	log.Fatal(http.ListenAndServe(":8086", nil))
}

func VerifyWithGemini(key, imgB64, userPrompt string) (bool, error) {
	if i := strings.Index(imgB64, ","); i != -1 {
		imgB64 = imgB64[i+1:]
	}

	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1/models/gemini-2.5-flash:generateContent?key=%s", key)

	payload := map[string]interface{}{
		"contents": []interface{}{
			map[string]interface{}{
				"parts": []interface{}{
					map[string]interface{}{"text": fmt.Sprintf("Look at this image. %s Respond with ONLY 'YES' or 'NO'.", userPrompt)},
					map[string]interface{}{
						"inline_data": map[string]interface{}{
							"mime_type": "image/jpeg",
							"data":      imgB64,
						},
					},
				},
			},
		},
	}

	jsonData, _ := json.Marshal(payload)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var result struct {
		Candidates []struct {
			Content struct {
				Parts []struct {
					Text string `json:"text"`
				} `json:"parts"`
			} `json:"content"`
		} `json:"candidates"`
	}

	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		log.Printf("JSON Unmarshal Error: %v", err)
		return false, err
	}

	if len(result.Candidates) > 0 && len(result.Candidates[0].Content.Parts) > 0 {
		answer := strings.TrimSpace(strings.ToUpper(result.Candidates[0].Content.Parts[0].Text))

		return strings.Contains(answer, "YES"), nil
	}

	return false, fmt.Errorf("no valid verification in Gemini response")
}
