package storage

import(
	"fmt"

	"go.uber.org/zap"
	//"github.com/jackc/sqlx/types"
	"github.com/jmoiron/sqlx"

	//"github.com/lib/pq"
)

const (
	driverName = "postgres"
	connStr = "user=%s password=%s sslmode=%s dbname=%s host=%s port=%s"
)

type sql struct {
	db *sqlx.DB
	l *zap.Logger
}

func NewSQL(host, username, password, name, port, ssl string, l *zap.logger) (Storage, error) {
	//Connect to DB with driver and
	var db, err = sqlx.Connect(driverName, fmt.Sprintf(connStr, username, password, ssl, name, host, port))
	if err != nil {
		l.Error("NewSQL: error connecting to db", zap.Error(err))
		return nil, err
	}

	//Ping DB to see if the connection is alive and working
	err = db.Ping()
	if err != nil {
		l.Error("NewSQL: error pinging db", zap.Error(err))
		return nil, err
	}
	 return &sql{
		 db:    db,
		 l: l,
	 }, nil
}

func (s *sql) AddCandidate() (string, error) {
	return "", nil
}
