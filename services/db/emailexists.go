package db

import (
  "errors"
  "os"
  "context"
  "encoding/json"
  "database/sql"
  "net/url"
  _ "github.com/denisenkom/go-mssqldb"
)

var (
  ctx context.Context
  db  *sql.DB
  cancel context.CancelFunc
)


// NullableBool is an alias for sql.NullBool data type
type NullableBool struct {
  sql.NullBool
}

// MarshalJSON for NullableBool
func (nb *NullableBool) MarshalJSON() ([]byte, error) {
  if !nb.Valid {
    return []byte("null"), nil
  }
  return json.Marshal(nb.Bool)
}

// #endregion

// #region Type definitions
type Response struct {
  Data bool `json:"data"`
  Meta ResponseMeta `json:"meta"`
}

type ResponseMeta struct {
}

// #endregion

// #region Control

func check(err error) {
  if err != nil {
    panic(err)
  }
}

func responseToJSON(response Response) ([]byte) {
  bytes, err := json.MarshalIndent(&response, "", "  ")
  check(err)
  return bytes
}
// #endregion

// #region Query concerns
func ValidateEmailAddress(email string, caseID string) (bool, error) {
  
  // Try to load .env file - Though remember file is optional; environmental
  // variable values may not be.
  //godotenv.Load()

  // Connect to the database
  connectDB()
  defer db.Close()
  defer cancel()

 
  query :=`select ACEVortex.dbo.DoesEmailExist(@emailAddress,@caseId)`
  
  rows, err := db.QueryContext(ctx, query, sql.Named("emailAddress", email),
       sql.Named("caseId", caseID))
  check(err)

  defer rows.Close()

  isExists := false

  if rows.Next()  {
    err = rows.Scan(&isExists) 
    check(err)
    return isExists, err
  }
  
  return isExists, nil

}

// #endregion

// #region Database concerns
func connectDB() {
  dbConnectURL, err := getDBConnectURL()

  check(err)

  ctx, cancel = context.WithCancel(context.Background())

  db, err = sql.Open("sqlserver", dbConnectURL)
  check(err)

  testDBConnection()

  //log.Print("Database connection established...")
}

func testDBConnection() {
  err := db.PingContext(ctx)
  check(err)
}

func getDBConnectURL() (string, error) {
  // Build SQL server connect string
  // sqlserver://username:password@host/instance?param1=value&param2=value
  // https://github.com/denisenkom/go-mssqldb
  dbHost := os.Getenv("DB_HOST")
  dbUser := os.Getenv("DB_USER")
  dbPass := os.Getenv("DB_PASS")

  if dbHost == "" || dbUser == "" || dbPass == "" {
    return "", errors.New("Invalid environment: DB_HOST, DB_USER, DB_PASS must be defined")
  }

  dbConnectURL := &url.URL{
    Scheme:   "sqlserver",
    User:     url.UserPassword(dbUser, dbPass),
    Host:     dbHost,
  }

  return dbConnectURL.String(), nil
}
// #endregion
