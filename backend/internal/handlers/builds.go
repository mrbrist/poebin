package handlers

import (
	"backend/internal/utils"
	"encoding/json"
	"net/http"
)

type newBuildParams struct {
	Raw string `json:"raw"`
}

type build struct {
	ID string `json:"id"`
}

func (cfg *APIConfig) GetBuild(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	build, err := cfg.R2.GetBuild(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Could not find build", err)
		return
	}

	utils.RespondWithJSON(w, 200, build)
}

func (cfg *APIConfig) NewBuild(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := newBuildParams{}
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	// Not sure if this is actully needed, commented out in case i need it in the future
	// is_valid, err := utils.IsValidBuild(params.Raw)
	// if err != nil {
	// 	utils.RespondWithError(w, http.StatusInternalServerError, "Couldn't validte build", err)
	// 	return
	// }

	// if !is_valid {
	// 	utils.RespondWithError(w, http.StatusInternalServerError, "Build is invalid", nil)
	// 	return
	// }

	build_id, err := cfg.R2.NewBuild(params.Raw)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Error creating build", err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, build{ID: build_id})
}
