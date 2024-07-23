package pgdb

import "go-transaction-manager/pkg/postgres"

type ReportRepo struct {
	*postgres.Postgres
}

func NewReportRepository(pg *postgres.Postgres) *ReportRepo {
	return &ReportRepo{pg}
}
