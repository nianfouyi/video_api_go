package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/nianfouyi/video-user-api/internal/models"
	"github.com/nianfouyi/video-user-api/internal/services"
	"net/http"
)

type PlaybackRecordHandler struct {
	recordService *services.PlaybackRecordService
}

func NewPlaybackRecordHandler(recordService *services.PlaybackRecordService) *PlaybackRecordHandler {
	return &PlaybackRecordHandler{recordService: recordService}
}

func (h *PlaybackRecordHandler) RecordPlayback(c *gin.Context) {
	var requestBody struct {
		UserID    uint `json:"user_id"`
		VideoID   uint `json:"video_id"`
		WatchTime int  `json:"watch_time"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.service.RecordPlayback(requestBody.UserID, requestBody.VideoID, requestBody.WatchTime)
	if err != nil {
		http.Error(w, "Failed to record playback", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Playback recorded successfully"})
}

func (h *PlaybackRecordHandler) CreateRecordPlayback(c *gin.Context) {
	var record models.PlaybackRecord
	if err:= c.ShouldBindJSON(&record);err!= nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	createRecord,err:= h.


}
