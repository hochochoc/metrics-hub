package handler

import (
	"encoding/json"
	"metric-hub/ent"
	"metric-hub/service"
	"metric-hub/utils"
	"net/http"
)

func ProcessReportHandler(w http.ResponseWriter, r *http.Request) {
	var newReport ent.Report
	err := json.NewDecoder(r.Body).Decode(&newReport)
	if err != nil {
		utils.HttpResponse(w, false, http.StatusBadRequest, err, nil)
		return
	}
	report, err := service.NewReportService(r.Context()).SaveReport(newReport)
	if err != nil {
		utils.HttpResponse(w, false, http.StatusInternalServerError, err, nil)
		return
	}
	utils.HttpResponse(w, true, http.StatusOK, err, report)
	return
}

func ProcessStatisticsHandler(w http.ResponseWriter, r *http.Request) {
	statistics, err := service.NewReportService(r.Context()).AggregateReports()
	if err != nil {
		utils.HttpResponse(w, false, http.StatusInternalServerError, err, nil)
		return
	}
	utils.HttpResponse(w, true, http.StatusOK, err, statistics)
	return
}

func ProcessOutliersHandler(w http.ResponseWriter, r *http.Request) {
	outliers, err := service.NewReportService(r.Context()).Outliers()
	if err != nil {
		utils.HttpResponse(w, false, http.StatusInternalServerError, err, nil)
		return
	}
	utils.HttpResponse(w, true, http.StatusOK, err, outliers)
	return
}
