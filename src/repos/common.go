package repos
import (
    "database/sql"
    "gopkg.in/gorp.v1"
    "../entity"
)

var db *sql.DB
var dbmap *gorp.DbMap

type Repo struct {
    Db *sql.DB
}

func (v Repo) Init() {
    db = v.Db
    dbmap = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
    dbmap.AddTableWithName(entity.Employee{}, "employee")
}
