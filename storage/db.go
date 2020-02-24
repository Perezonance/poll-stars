package storage

import(
	"go.uber.com/zap"
	"github.com/jackc/sqlx/types"
	"github.com/jmoiron.sqlx"

	"github.com/lib/pq"
)

type sql struct {
	db *sqlz.DB
	logger *zap.Logger
}

func NewSQL(host, username, password, name, port, ssl string, l *zap.logger, ) (sql, error) {
	var db, err = sqlx.Connect()
}

func (s *sql) AddCandidate() (string, error) {
	return "", nil
}
