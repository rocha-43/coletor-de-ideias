package see

import (
	"fmt"
	"net/http"
)

func See(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OK -view")
}
