package start
import (
	"flag"
	"log"
	"net/http"

	"go.uber.org/zap"

	"github.com/alexperez/poll-stars/server"
	"github.com/alexperez/poll-stars/storage"
)

func Start() {
	//Receive db connection credentials from command line argument flags
	var (
		prod = flag.Bool("prod-log", false, "Debug or production logging")
		dbHost = flag.String("db-host", "localhost", "DB host location")
		dbUser = flag.String("db-user", "username", "DB username")
		dbPass = flag.String("db-pass", "password", "DB password")
		dbName = flag.String("db-name", "dbname", "DB name")
		dbPort = flag.String("db-port", "5432", "DB port")
		ssl = flag.String("db-ssl", "0.0.0.0", "DB ssl")

		addy = flag.String("sv-addr", "0.0.0.0", "server address")
		port = flag.String("sv-port", "8080", "server port")
	)

	//Parse the flags after declaring them all
	flag.Parse()

	//Initialize a new DBCreds struct to hold the flag values
	dbC := storage.DBCredentials {
		Host: *dbHost,
		User: *dbUser,
		Pass: *dbPass,
		Name: *dbName,
		Port: *dbPort,
		Ssl:  *ssl,
	}

	//Initialize the zap logger for structured logging
	l, err := getLogger(*prod)
	if err != nil {
		log.Fatalf("Logger failed to be initialized: %v", err)
	}

	//Initialize the connection to the db
	db, err := storage.NewSQL(dbC, l)
	if err != nil {
		log.Fatalf("Db Connection failed to be initialized: %v", err)
	}

	//Create a server struct that contains a logger and database connection
	s := server.NewServer(db, l)

	//Route the handling function for a URL and begin listening
	http.HandleFunc("/", s.VoteHandler)
	err = http.ListenAndServe(*addy + ":" + *port, nil)
	if err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	//Server shut down
	l.Info("Server has Shut Off")
}

//Determines which logger to use
func getLogger(prod bool) (*zap.Logger, error) {
	if prod {
		return zap.NewProduction()
	}
	return zap.NewDevelopment()
}
