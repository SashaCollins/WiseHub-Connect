package data

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
    "log"
    "os"
)
type Datastore struct{}

func (ds *Datastore) init() {

}
func (ds *Datastore) Load(){

}
func (ds *Datastore) Save(){

}
func (ds *Datastore) Del(){

}
func (ds *Datastore) Start(){
    fmt.Println("1")
    os.Remove("./foo.db")

    fmt.Println("1")
    db, err := sql.Open("sqlite3", "./foo.db")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("1")
    defer db.Close()

    sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
    fmt.Println("1")
    _, err = db.Exec(sqlStmt)
    if err != nil {
        log.Printf("%q: %s\n", err, sqlStmt)
        return
    }
    fmt.Println("1")
    tx, err := db.Begin()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("1")
    stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("1")
    defer stmt.Close()
    for i := 0; i < 100; i++ {
        _, err = stmt.Exec(i, fmt.Sprintf("こんにちわ世界%03d", i))
        if err != nil {
            log.Fatal(err)
        }
    }
    fmt.Println("1")
    tx.Commit()

    rows, err := db.Query("select id, name from foo")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("1")
    defer rows.Close()
    for rows.Next() {
        var id int
        var name string
        err = rows.Scan(&id, &name)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(id, name)
    }
    fmt.Println("1")
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }

    stmt, err = db.Prepare("select name from foo where id = ?")
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()
    var name string
    err = stmt.QueryRow("3").Scan(&name)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(name)

    _, err = db.Exec("delete from foo")
    if err != nil {
        log.Fatal(err)
    }

    _, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
    if err != nil {
        log.Fatal(err)
    }

    rows, err = db.Query("select id, name from foo")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    for rows.Next() {
        var id int
        var name string
        err = rows.Scan(&id, &name)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(id, name)
    }
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }
}
