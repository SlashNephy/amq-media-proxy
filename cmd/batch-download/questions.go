package main

import (
	"encoding/json"
	"errors"
	"math/rand"
	"os"

	errors2 "github.com/pkg/errors"
)

type Question struct {
	Songs []struct {
		Examples map[string]string `json:"examples"`
	} `json:"songs"`
}

func parseQuestionsJSON(path string) ([]*Question, error) {
	if path == "" {
		return nil, errors.New("QUESTIONS_JSON_PATH is not set")
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, errors2.WithStack(err)
	}

	var questions []*Question
	if err = json.Unmarshal(bytes, &questions); err != nil {
		return nil, errors2.WithStack(err)
	}

	return questions, nil
}

func LoadMediaURLs(path string) ([]string, error) {
	questions, err := parseQuestionsJSON(path)
	if err != nil {
		return nil, errors2.WithStack(err)
	}

	var urls []string
	for _, q := range questions {
		for _, s := range q.Songs {
			for _, e := range s.Examples {
				urls = append(urls, e)
			}
		}
	}

	rand.Shuffle(len(urls), func(i, j int) {
		urls[i], urls[j] = urls[j], urls[i]
	})

	return urls, nil
}
