package service

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"errors"
	"github.com/spf13/viper"
	"metric-hub/config"
	"metric-hub/ent"
	"metric-hub/ent/report"
	"metric-hub/model"
)

type ReportService struct {
	ctx    context.Context
	client *ent.Client
}

func NewReportService(ctx context.Context) *ReportService {
	return &ReportService{
		ctx:    ctx,
		client: config.GetClient(),
	}
}

func (s *ReportService) SaveReport(newReport ent.Report) (*ent.Report, error) {
	return s.client.Report.Create().
		SetServerName(newReport.ServerName).
		SetStartTime(newReport.StartTime).
		SetEndTime(newReport.EndTime).
		SetDurationInSeconds(newReport.EndTime.Sub(newReport.StartTime).Seconds()).
		Save(s.ctx)
}

func (s *ReportService) CountReports() (int, error) {
	return s.client.Report.Query().Count(s.ctx)
}

func (s *ReportService) AggregateReports() (*model.ReportStatistics, error) {
	minNbOfReports := viper.GetInt("MIN_NUMBER_OF_REPORTS")
	nbReports, err := s.CountReports()
	if err != nil {
		return nil, err
	}
	if nbReports < minNbOfReports {
		return nil, errors.New("not enough reports received")
	}
	statistics := make([]*model.ReportStatistics, 0)
	err = s.client.Report.Query().Modify(func(s *sql.Selector) {
		s.Select(
			sql.As("ROUND(AVG(duration_in_seconds))", "mean"),
			sql.As("ROUND(STDDEV_POP(duration_in_seconds))", "stddev"))
	}).Scan(s.ctx, &statistics)
	if err != nil {
		return nil, err
	}
	if len(statistics) >= 1 {
		return statistics[0], nil
	}
	return nil, errors.New("no statistics returned")
}

func (s *ReportService) Outliers() ([]string, error) {
	outlierThreshold := viper.GetFloat64("OUTLIER_THRESHOLD")
	statistics, err := s.AggregateReports()
	if err != nil {
		return nil, err
	}
	mean := statistics.Mean
	stddev := statistics.Stddev

	return s.client.Report.Query().
		Where(report.Or(
			report.DurationInSecondsGT(mean+outlierThreshold*stddev),
			report.DurationInSecondsLT(mean-outlierThreshold*stddev))).
		Modify(func(s *sql.Selector) {
			s.Distinct().Select(report.FieldServerName)
		}).
		Strings(s.ctx)
}
