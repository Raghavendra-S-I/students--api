package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/Raghavendra/students-api/internal/types"
	"github.com/Raghavendra/students-api/internal/utils/response"
)

func New() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}
		slog.Info("creating a student")

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "ok"})
	}
}
