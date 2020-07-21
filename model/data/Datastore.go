/* This Datastore works with sqlite3 Database.
 to use another database exchange this, and only this, file with a 'Database.go' file t
hat works with your desired database. Make sure to implement the Interface DatastoreI.go,
otherwise your dashboard may not work properly.
@author SashaCollins
 */
package data

import (
    "database/sql"
    "errors"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
    "log"
    //"os"
)
var (
    db *sql.DB
)

type Datastore struct{}

func init() {
    //_ = os.Remove("./data/wisehub.db")
    var err error
    db, err = sql.Open("sqlite3", "./data/wisehub.db")
    if err != nil {
      fmt.Printf("ERROR: %s\n", err)
    }

    sqlStmt := `
	create table if not exists User (id integer not null primary key autoincrement , name text, password text, email text unique);
-- 	delete from User;
	`
    _, err = db.Exec(sqlStmt)
    if err != nil {
       log.Printf("%q: %s\n", err, sqlStmt)
       fmt.Printf("%q: %s\n", err, sqlStmt)
       return
    }
}

/* Please note:
Input either zero, one or two arguments
zero loads all users from database
one loads user based on email address
two loads user based on email and password
 @param email string
 @param password string
*/
func (ds *Datastore) Load(args ...string) (user User ,err error){
    // We initialize each of the optional parameters to their default value.
    email := "" // ← We initialize email to empty string.
    //password := "" // ← We initialize password to empty string.

    var (
        sqlStmt string
        rows *sql.Rows
        row *sql.Row
    )

    // Get any parameters passed to us out of the args variable into "real"
    // variables we created for them.
    l := len(args)
    switch l {
    //LoadAll
    case 0:
        fmt.Println("\tat 0")
        sqlStmt = `select name, password, email from User;`
        stmt, err := db.Prepare(sqlStmt)
        if err != nil {
            return user, err
        }
        rows, err = stmt.Query()
        if err != nil {
            return user, err
        }
        for rows.Next() {
            if err := row.Scan(&user.Name, &user.Password, &user.Email); err != nil {
                return user, err
            }
            fmt.Println(user.Name, user.Password, user.Email)
        }
        //defer rows.Close()
        err = rows.Err()
        if err != nil {
            err = errors.New("something went wrong while fetching all Users")
            return user, err
        }
        //only last user is returned; multi-user loading for ui not supported yet
        return user, nil

    //register, login, forgot
    case 1: // email
        fmt.Println("\tat 1")
        email = args[0]
        sqlStmt = `
           select name, password, email
           from User
           where User.email = ?;
       `
        stmt, err := db.Prepare(sqlStmt)
        if err != nil {
            return user, err
        }
        defer stmt.Close()
        row = stmt.QueryRow(email)
        if err := row.Scan(&user.Name, &user.Password, &user.Email); err != nil {
            err = errors.New("no matching user was found")
            return user, err
        }
        return user, nil

    ////login
    //case 2: // email, password
    //    fmt.Println("\tat 2")
    //    email = args[0]
    //    password = args[1]
    //    sqlStmt = `
    //       select name, password, email
    //       from User
    //       where User.email = ?
    //       and User.password = ?;
    //   `
    //    stmt, err := db.Prepare(sqlStmt)
    //    if err != nil {
    //        return user, err
    //    }
    //    defer stmt.Close()
    //    row = stmt.QueryRow(email, password)
    //    if err := row.Scan(&user.Name, &user.Password, &user.Email); err != nil {
    //        err = errors.New("no matching user was found")
    //        return user, err
    //    }
    //    return user, nil


    // Since we have 0 mandatory parameter, but a maximum of 3 make sure
    // we have no more than 3.
    //
    default:
        err = errors.New("too many parameters.")
        return user, err
    }
}

func (ds *Datastore) Save(name string, password string, email string) error{

    fmt.Println("\tin Save")
    //fmt.Printf("%s\n", name)
    //fmt.Printf("%s\n", password)
    //fmt.Printf("%s\n", email)
    tx, err := db.Begin()
    if err != nil {
       return err
    }

    sqlStmt := `
       insert 
       into User(name, password, email) 
       values(?, ?, ?);
    `
    stmt, err := tx.Prepare(sqlStmt)
    if err != nil {
       return err
    }
    //fmt.Println("\tafter prepare")
    defer stmt.Close()

    _, err = stmt.Exec(name, password, email)
    if err != nil {
       return err
    }
    //fmt.Println("\tafter exec")

    err = tx.Commit()
    if err != nil {
      return err
    }
    fmt.Println("\tend Save")
    return nil
}
func (ds *Datastore) Del(email string, password string) error{
    tx, err := db.Begin()
    if err != nil {
        return err
    }

    sqlStmt := `
       delete 
       from User
       where User.email = ? 
       and User.password = ?;
    `
    stmt, err := tx.Prepare(sqlStmt)
    if err != nil {
        return err
    }
    defer stmt.Close()
    _, err = ds.Load(email, password)
    if err != nil {
        return err
    }

    _, err = stmt.Exec(email, password)
    if err != nil {
        return err
    }

    err = tx.Commit()
    if err != nil {
        return err
    }
    return nil
}

