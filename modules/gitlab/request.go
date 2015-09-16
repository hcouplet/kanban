package gitlab

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Execute request to gitlab and chek err
func (g *ApiGitlab) Do(cl *http.Client, req *http.Request, v interface{}) error {
	res, err := cl.Do(req)

	if err != nil {
		log.Printf("%s", err.Error())
		return err
	}

	if res.StatusCode != http.StatusOK {
		log.Printf("%s: %i", "Bad response code", res.StatusCode)
		return errors.New(fmt.Sprintf("%s: %i", "Bad response code", res.StatusCode))
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&v); err != nil {
		log.Printf("%s: %i", "Bad response code", err.Error())
		return err
	}

	return nil
}

func (g *ApiGitlab) GetUrl(p []string) string {
	return g.config.BasePath + "/" + strings.Join(p, "/")
}
