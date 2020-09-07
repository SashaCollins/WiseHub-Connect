/* This Datastore works with sqlite3 Database.
 to use another database exchange this, and only this, file with a 'Database.go' file t
hat works with your desired database. Make sure to implement the Interface DatastoreI.go,
otherwise your dashboard may not work properly.
@author SashaCollins
 */
package data

import (
    "fmt"
    _ "github.com/mattn/go-sqlite3"
    "gorm.io/gorm"
    "log"

    "gorm.io/driver/mysql"
    "gorm.io/driver/postgres"
    "gorm.io/driver/sqlite"
    "gorm.io/driver/sqlserver"
)

var (
    db *gorm.DB
    err error
)

type Datastore struct{}

func loadAllPlugins() (plugins []Plugin) {
    var plugin Plugin
    result := db.Find(&plugin)
    fmt.Println(result)
    if result.Error != nil {
        log.Printf("loadAllPlugins: %q\n", result.Error)
        return
    }
    return plugins



    //loadPlugins := `
	//select name, description from Plugin;
	//`
    //tx, err := db.Begin()
    //if err != nil {
    //    return nil, err
    //}
    //stmt, err := tx.Prepare(loadPlugins)
    //if err != nil {
    //   return nil, err
    //}
    //var plugin Plugin
    //rows, err = stmt.Query()
    //if err != nil {
    //    return plugins, err
    //}
    //defer rows.Close()
    //for rows.Next() {
    //    if err := rows.Scan(&plugin.Name, &plugin.Description); err != nil {
    //        return plugins, err
    //    }
    //    plugins = append(plugins, plugin)
    //    fmt.Println(plugin.Name, plugin.Description)
    //}
    //err = rows.Err()
    //if err != nil {
    //    err = errors.New("something went wrong while fetching all Users")
    //    return plugins, err
    //}
    //return plugins, nil
}

func loadAllUsers() (user User) {
    result := db.Find(&user)
    fmt.Println(result)
    if result.Error != nil {
        log.Printf("loadAllUsers: %q\n", result.Error)
        return
    }
    return user


    //sqlStmt := `select name, password, email, plugins from User;`
    //stmt, err := db.Prepare(sqlStmt)
    //if err != nil {
    //    return user, plugins, err
    //}
    //rows, err = stmt.Query()
    //if err != nil {
    //    return user, plugins, err
    //}
    //defer rows.Close()
    //for rows.Next() {
    //    if err := rows.Scan(&user.Name, &user.Password, &user.Email, &user.Plugins); err != nil {
    //        return user, plugins, err
    //    }
    //    if check := len(user.Plugins); check > 0 {
    //        plugins, err = loadAllPlugins()
    //        if err != nil {
    //            return user, plugins, err
    //        }
    //    }
    //    fmt.Println(user.Name, user.Password, user.Email, user.Plugins)
    //}
    //err = rows.Err()
    //if err != nil {
    //    err = errors.New("something went wrong while fetching all Users")
    //    return user, plugins, err
    //}
    //return user, plugins,nil
}

func loadUserByEmail(email string) (user User) {
    result := db.Where("email = ?", email).First(&user)
    if result.Error != nil {
        log.Printf("loadUserByEmail: 1. %q\n", result.Error)
        return
    }
    if check := len(user.Plugins); check > 0 {
        plugins := loadAllPlugins()
        if err == nil {
            log.Printf("loadUserByEmail: 2. %q\n", result.Error)
            return
        }
        user.Plugins = plugins
    }
    return user

    //sqlStmt := `
    //       select name, password, email, plugins
    //       from User
    //       where User.email = ?;
    //   `
    //stmt, err := db.Prepare(sqlStmt)
    //if err != nil {
    //    return user, plugins, err
    //}
    //row = stmt.QueryRow(email)
    //if err := row.Scan(&user.Name, &user.Password, &user.Email, &user.Plugins); err != nil {
    //    return user, plugins, err
    //}
    //if check := len(user.Plugins); check > 0 {
    //    plugins, err = loadAllPlugins()
    //    if err != nil {
    //        return user, plugins, err
    //    }
    //}
    //return user, plugins,nil
}

func createTables(db *gorm.DB) {
    plugins := loadAllPlugins()
    if plugins == nil {
        log.Printf("createTables: 1. %q\n", err)
        return
    }

    admin := User{
        Name: "Admin",
        Password: "",
        Email: "test@test.de",
        Plugins: plugins,
    }
    result := db.Create(&admin) // pass pointer of data to Create
    if result.Error != nil {
        log.Printf("createTables: 2. %q\n", result.Error)
        return
    }




    //createUser := `
    //create table if not exists User (id integer not null primary key autoincrement , name text, password text, email text unique, plugins text);
    //`
    //createPlugins := `
    //create table if not exists Plugin (id integer not null primary key autoincrement , name text, token text, description text unique);
    //`
    //_, err = db.Exec(createUser)
    //if err != nil {
    //    log.Printf("init %q: %s\n", err, createUser)
    //    return
    //}
    //_, err = db.Exec(createPlugins)
    //if err != nil {
    //    log.Printf("init %q: %s\n", err, createPlugins)
    //    return
    //}
    //
    //initGithub := `
    //insert into Plugin (name, token, description) values(?, ?, ?);
    //`
    //initDrone := `
    //insert into Plugin (name, token, description) values(?, ?, ?);
    //`
    //initHeroku := `
    //insert into Plugin (name, token, description) values(?, ?, ?);
    //`
    //_, err = db.Exec(initGithub, "", "", "Github")
    //if err != nil {
    //    log.Printf("init %q: %s\n", err, createPlugins)
    //    return
    //}
    //_, err = db.Exec(initDrone, "", "", "Drone CI")
    //if err != nil {
    //    log.Printf("init %q: %s\n", err, createPlugins)
    //    return
    //}
    //_, err = db.Exec(initHeroku, "", "", "Heroku")
    //if err != nil {
    //    log.Printf("init %q: %s\n", err, createPlugins)
    //    return
    //}
}

func openDB(driver string, data map[string]interface{}) (db *gorm.DB, err error) {
    switch driver {
    case "mysql":
        if length := len(data); length < 5 {
            log.Printf("data is too short. Length: %d", length)
            return
        }
        user := fmt.Sprintf("%s", data["user"])
        password := fmt.Sprintf("%s", data["password"])
        dbname := fmt.Sprintf("%s", data["dbname"])
        host := fmt.Sprintf("%s", data["host"])
        port := fmt.Sprintf("%d", data["port"])
        dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname
        return gorm.Open(mysql.Open(dsn), &gorm.Config{})
    case "sqlite":
        return gorm.Open(sqlite.Open("wisehub.db"), &gorm.Config{})
    case "postgres":
        if length := len(data); length < 6 {
            log.Printf("data is too short. Length: %d", length)
            return
        }
        user := fmt.Sprintf("user=%v ", data["user"])
        password := fmt.Sprintf("password=%v ", data["password"])
        dbname := fmt.Sprintf("dbname=%v ", data["dbname"])
        port := fmt.Sprintf("port=%v ", data["port"])
        sslMode := fmt.Sprintf("sslmode=%v ", data["sslMode"])
        timeZone := fmt.Sprintf("TimeZone=%v", data["timeZone"])
        dsn := user + password + dbname + port + sslMode + timeZone
        return gorm.Open(postgres.New(postgres.Config{
            DSN: dsn,
            PreferSimpleProtocol: true, // disables implicit prepared statement usage
        }), &gorm.Config{})
    case "sqlserver":
        if length := len(data); length < 5 {
            log.Printf("data is too short. Length: %d", length)
            return
        }
        user := fmt.Sprintf("%s", data["user"])
        password := fmt.Sprintf("%s", data["password"])
        dbname := fmt.Sprintf("%s", data["dbname"])
        host := fmt.Sprintf("%s", data["host"])
        port := fmt.Sprintf("%d", data["port"])
        dsn := "sqlserver://" + user + ":" + password + "@" + host + ":" + port + "?database=" + dbname
        return gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
    default:
        log.Printf("Driver does not exists. Driver: %s", driver)
        return
    }
}

func (ds *Datastore) New(driver string, data  map[string]interface{}) (db *gorm.DB) {
    db, err = openDB(driver, data)
    if err != nil {
        log.Printf("New %q: %v\n", err, db)
        return
    }
    createTables(db)
    return db
}

/* Please note:
Input either zero, one or two arguments
zero loads all users from database
one loads user based on email address
two loads user based on email and password
 @param email string
 @param password string
*/
func (ds *Datastore) Load(email ...string) (user User) {
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
        fmt.Printf("too many arguments in function load: %v\n", len(email))
        return user
    }
}

//TODO Plugin Save
func (ds *Datastore) Save(name string, password string, email string) error {
    fmt.Println("Save \tin Save")

    //tx, err := db.Begin()
    //if err != nil {
    //   return err
    //}
    //
    //defaultPlugins := "1,2,3"
    //
    //saveUser := `
    //   insert
    //   into User(name, password, email, plugins)
    //   values(?, ?, ?, ?);
    //`
    //stmt, err := tx.Prepare(saveUser)
    //if err != nil {
    //   return err
    //}
    //defer stmt.Close()
    //
    //_, err = stmt.Exec(name, password, email, defaultPlugins)
    //if err != nil {
    //   return err
    //}
    //
    //err = tx.Commit()
    //if err != nil {
    //  return err
    //}
    //fmt.Println("Save \tend Save")
    return nil
}

// TODO plugin update
func (ds *Datastore) Update(option string, data ...interface{}) error {
    fmt.Println("Update \tin Update")
    //fmt.Printf("%s\n", data)
    //fmt.Printf("%s\n", option)
    //tx, err := db.Begin()
    //if err != nil {
    //   return err
    //}

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
    fmt.Println(sqlStmt)
    //stmt, err := tx.Prepare(sqlStmt)
    //if err != nil {
    //   return err
    //}
    ////fmt.Println("\tafter prepare")
    //defer stmt.Close()
    //
    //_, err = stmt.Exec(data[0], data[1])
    //if err != nil {
    //   return err
    //}
    ////fmt.Println("\tafter exec")
    //
    //err = tx.Commit()
    //if err != nil {
    //  return err
    //}
    //fmt.Println("Update \tend Save")
    return nil
}

func (ds *Datastore) Delete(email string) error {
    //tx, err := db.Begin()
    //if err != nil {
    //    return err
    //}
    //
    //deleteProfile := `
    //   delete
    //   from User
    //   where User.email = ?;
    //`
    //stmt, err := tx.Prepare(deleteProfile)
    //if err != nil {
    //    return err
    //}
    //defer stmt.Close()
    //
    //_, err = stmt.Exec(email)
    //if err != nil {
    //    return err
    //}
    //return nil
    return nil
}

