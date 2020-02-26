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

type DBCredentials struct {
	Host string
	User	string
	Pass	string
	Name	string
	Port 	string
	Ssl 	string
}

//NewSQL takes in a logger and dbCredentials struct to establish a connection with a DB
func NewSQL(dbC DBCredentials, l *zap.Logger) (Storage, error) {
	//Connect to DB with driver and
	var db, err = sqlx.Connect(driverName, fmt.Sprintf(connStr, dbC.user, dbC.pass, dbC.ssl, dbC.name, dbC.host, dbC.port))
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
	//Return the connection struct
	 return &sql{
		 db:    db,
		 l: l,
	 }, nil
}

func (s *sql) AddCandidate() (string, error) {
	//TODO: UNIMPLIMENTED
	return "", nil
}

func (s *sql) SubmitVote() (string, error) {
	//TODO: UNIMPLIMENTED
	return "", nil
}
