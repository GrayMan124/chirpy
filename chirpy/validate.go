package main

import (
	"strings"
)

//	type chirp struct {
//		Body string `json:"body"`
//	}
type errorResponse struct {
	Error string `json:"error"`
}

type validResponse struct {
	CleanedBody string `json:"cleaned_body"`
}

func replaceWords(input string) string {
	split := strings.Split(input, " ")

	for idx, word := range split {
		if strings.ToLower(word) == "kerfuffle" || strings.ToLower(word) == "sharbert" || strings.ToLower(word) == "fornax" {
			split[idx] = "****"
		}
	}
	return strings.Join(split, " ")
}

// func validation(w http.ResponseWriter, r *http.Request) {
//
// 	decoder := json.NewDecoder(r.Body)
// 	c := chirp{}
// 	err := decoder.Decode(&c)
// 	if err != nil {
// 		output := errorResponse{
// 			Error: "Something went wrong",
// 		}
// 		w.WriteHeader(500)
// 		out, err := json.Marshal(output)
// 		if err != nil {
// 			log.Fatal("failed marshalling")
// 		}
// 		w.Write(out)
// 		w.Header().Set("Content-Type", "application/json")
// 		return
// 	}
//
// 	if len(c.Body) > 140 {
//
// 		output := errorResponse{
// 			Error: "Chirp is too long",
// 		}
// 		w.WriteHeader(400)
// 		out, err := json.Marshal(output)
// 		if err != nil {
// 			log.Fatal("failed marshalling")
// 		}
// 		w.Write(out)
// 		w.Header().Set("Content-Type", "application/json")
// 		return
// 	}
//
// 	output := validResponse{
// 		CleanedBody: replaceWords(c.Body),
// 	}
// 	w.WriteHeader(200)
// 	out, err := json.Marshal(output)
// 	if err != nil {
// 		log.Fatal("failed marshalling")
// 	}
// 	w.Write(out)
// 	w.Header().Set("Content-Type", "application/json")
//
// }
