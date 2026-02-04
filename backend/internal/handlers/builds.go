package handlers

import (
	"backend/internal/utils"
	"encoding/json"
	"net/http"
	"time"
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

	build_id, level, class, err := cfg.R2.NewBuild(params.Raw)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Error creating build", err)
		return
	}

	// Append this build to the recent builds stored in memory
	cfg.RecentBuilds = append(cfg.RecentBuilds, RecentBuild{
		ID:        build_id,
		Level:     level,
		Class:     class,
		DateAdded: time.Now().UnixMilli(),
	})

	utils.RespondWithJSON(w, http.StatusOK, build{ID: build_id})
}

func (cfg *APIConfig) GetRecentBuilds(w http.ResponseWriter, r *http.Request) {
	if len(cfg.RecentBuilds) == 0 {
		utils.RespondWithJSON(w, 200, [2]RecentBuild{
			{ID: "cNHxLQX23eEGYD3TAC8vsE", Level: 100, Class: "Hierophant", DateAdded: 1769733065986},
			{ID: "naJZBYXhayPPniJDKE4U7m", Level: 100, Class: "Surfcaster", DateAdded: 1769233765986},
		})
		return
	}
	utils.RespondWithJSON(w, 200, cfg.RecentBuilds)
}
