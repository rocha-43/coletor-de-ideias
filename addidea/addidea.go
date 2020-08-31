package addidea

import (
	"fmt"
	"net/http"
)

func AddIdea(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OK -add idea")
}
