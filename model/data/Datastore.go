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
)

var (
    db *sql.DB
    err error
    rows *sql.Rows
    row *sql.Row
)

type Datastore struct{}

func loadAllPlugins() (plugins []Plugin, err error) {
    loadPlugins := `
	select name, description from Plugin;
	`
    tx, err := db.Begin()
    if err != nil {
        return nil, err
    }
    stmt, err := tx.Prepare(loadPlugins)
    if err != nil {
       return nil, err
    }
    var plugin Plugin
    rows, err = stmt.Query()
    if err != nil {
        return plugins, err
    }
    defer rows.Close()
    for rows.Next() {
        if err := rows.Scan(&plugin.Name, &plugin.Description); err != nil {
            return plugins, err
        }
        plugin.Token = "****"
        plugins = append(plugins, plugin)
        fmt.Println(plugin.Name, plugin.Description)
    }
    err = rows.Err()
    if err != nil {
        err = errors.New("something went wrong while fetching all Users")
        return plugins, err
    }
    return plugins, nil
}

func loadAllUsers() (user User, plugins []Plugin, err error) {
    sqlStmt := `select name, password, email, plugins from User;`
    stmt, err := db.Prepare(sqlStmt)
    if err != nil {
        return user, plugins, err
    }
    rows, err = stmt.Query()
    if err != nil {
        return user, plugins, err
    }
    defer rows.Close()
    for rows.Next() {
        if err := rows.Scan(&user.Name, &user.Password, &user.Email, &user.Plugins); err != nil {
            return user, plugins, err
        }
        if check := len(user.Plugins); check > 0 {
            plugins, err = loadAllPlugins()
            if err != nil {
                return user, plugins, err
            }
        }
        fmt.Println(user.Name, user.Password, user.Email, user.Plugins)
    }
    err = rows.Err()
    if err != nil {
        err = errors.New("something went wrong while fetching all Users")
        return user, plugins, err
    }
    return user, plugins,nil
}

func loadUserByEmail(email string) (user User, plugins []Plugin, err error) {
    sqlStmt := `
           select name, password, email, plugins
           from User
           where User.email = ?;
       `
    stmt, err := db.Prepare(sqlStmt)
    if err != nil {
        return user, plugins, err
    }
    row = stmt.QueryRow(email)
    if err := row.Scan(&user.Name, &user.Password, &user.Email, &user.Plugins); err != nil {
        return user, plugins, err
    }
    if check := len(user.Plugins); check > 0 {
        plugins, err = loadAllPlugins()
        if err != nil {
            return user, plugins, err
        }
    }
    return user, plugins,nil
}

func openDb() {
    db, err = sql.Open("sqlite3", "./data/wisehub.db")
    if err != nil {
        fmt.Printf("openDB: %s\n", err)
    }
}

func init() {
    var err error
    openDb()

    createUser := `
	create table if not exists User (id integer not null primary key autoincrement , name text, password text, email text unique, plugins text);
    delete from User;
	`
    createPlugins := `
	create table if not exists Plugin (id integer not null primary key autoincrement , name text, token text, description text unique);
    delete from Plugin;
	`
    _, err = db.Exec(createUser)
    if err != nil {
       log.Printf("init %q: %s\n", err, createUser)
       return
    }
    _, err = db.Exec(createPlugins)
    if err != nil {
       log.Printf("init %q: %s\n", err, createPlugins)
       return
    }

    initGithub := `
	insert into Plugin (name, token, description) values(?, ?, ?);
	`
    initDrone := `
	insert into Plugin (name, token, description) values(?, ?, ?);
	`
    initHeroku := `
	insert into Plugin (name, token, description) values(?, ?, ?);
	`
    _, err = db.Exec(initGithub, "", "", "Github")
    if err != nil {
        log.Printf("init %q: %s\n", err, createPlugins)
        return
    }
    _, err = db.Exec(initDrone, "", "", "Drone CI")
    if err != nil {
        log.Printf("init %q: %s\n", err, createPlugins)
        return
    }
    _, err = db.Exec(initHeroku, "", "", "Heroku")
    if err != nil {
        log.Printf("init %q: %s\n", err, createPlugins)
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
func (ds *Datastore) Load(email ...string) (user User ,plugins []Plugin, err error) {
    // open DB
    openDb()
    //defer db.Close()

    // Get any parameters passed to us out of the args variable into "real"
    // variables we created for them.
    check := len(email)
    switch check {
    case 0:
        fmt.Println("Load \tat 0")
        return loadAllUsers()
    case 1:
        fmt.Println("Load \tat 1")
        return loadUserByEmail(email[0])
    default:
        err := fmt.Errorf("too many arguments in function load: %v", len(email))
        return user, plugins, err
    }
}

//TODO Plugin Save
func (ds *Datastore) Save(name string, password string, email string) error {
    fmt.Println("Save \tin Save")
    // open DB
    openDb()
    //defer db.Close()

    tx, err := db.Begin()
    if err != nil {
       return err
    }

    defaultPlugins := "1,2,3"

    saveUser := `
       insert 
       into User(name, password, email, plugins) 
       values(?, ?, ?, ?);
    `
    stmt, err := tx.Prepare(saveUser)
    if err != nil {
       return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(name, password, email, defaultPlugins)
    if err != nil {
       return err
    }

    err = tx.Commit()
    if err != nil {
      return err
    }
    fmt.Println("Save \tend Save")
    return nil
}

// TODO plugin update
func (ds *Datastore) Update(option string, data ...string) error {
    // open DB
    openDb()
    //defer db.Close()

    fmt.Println("Update \tin Update")
    //fmt.Printf("%s\n", data)
    //fmt.Printf("%s\n", option)
    tx, err := db.Begin()
    if err != nil {
       return err
    }

    sqlStmt := ""
    switch option {
    case "email":
        sqlStmt = `
            update User
            set User.email = ?
            where User.email = ?;
        `
    case "password":
        sqlStmt = `
            update User
            set User.password = ?
            where User.email = ?;
        `
    case "plugins":
        sqlStmt = `
            update User
            set User.plugins = ?
            where User.email = ?;
        `
    }
    stmt, err := tx.Prepare(sqlStmt)
    if err != nil {
       return err
    }
    //fmt.Println("\tafter prepare")
    defer stmt.Close()

    _, err = stmt.Exec(data[0], data[1])
    if err != nil {
       return err
    }
    //fmt.Println("\tafter exec")

    err = tx.Commit()
    if err != nil {
      return err
    }
    fmt.Println("Update \tend Save")
    return nil
}

func (ds *Datastore) Delete(email string) error {
    // open DB
    openDb()
    //defer db.Close()

    tx, err := db.Begin()
    if err != nil {
        return err
    }

    deleteProfile := `
       delete 
       from User
       where User.email = ?;
    `
    stmt, err := tx.Prepare(deleteProfile)
    if err != nil {
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(email)
    if err != nil {
        return err
    }
    return nil
}

