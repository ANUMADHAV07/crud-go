package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct {
}

func NewWorkoutHandler() *WorkoutHandler {
	return &WorkoutHandler{}
}

func (wh *WorkoutHandler) HandleGetWorkoutByID(w http.ResponseWriter, r *http.Request) {
	paramsWokoutID := chi.URLParam(r, "id")

	if paramsWokoutID == "" {
		http.NotFound(w, r)
		return
	}

	workourID, err := strconv.ParseInt(paramsWokoutID, 10, 64)

	if err != nil {
		http.NotFound(w, r)
	}

	fmt.Fprintf(w, "this is the workout id: %d\n", workourID)

}

func (wh *WorkoutHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "workout created \n")
}
