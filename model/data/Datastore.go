/* This Datastore works with sqlite3 Database.
 to use another database exchange this, and only this, file with a 'Database.go' file t
hat works with your desired database. Make sure to implement the Interface DatastoreI.go,
otherwise your dashboard may not work properly.
@author SashaCollins
 */
package data

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
)

var defaultPlugins = []Plugin{
    {PluginName: "Github", UsernameHost: "", Token: "", Description: "", Updated: false},
    {PluginName: "Drone CI", UsernameHost: "", Token: "", Description: "", Updated: false},
    {PluginName: "Heroku", UsernameHost: "", Token: "", Description: "", Updated: false},
}

type Datastore struct{}

func loadAllPlugins(db *gorm.DB, userID uint) (plugins []Plugin, err error) {
    fmt.Println("loadAllPlugins \tstart")
    rows, err := db.Model(&Plugin{}).Where("user_id = ?", userID).Rows()
    if err != nil {
        log.Printf("loadAllPlugins1: %q\n", err)
        return plugins, err
    }
    defer rows.Close()

    var plugin Plugin
    for rows.Next() {
        if err = db.ScanRows(rows, &plugin); err != nil {
            fmt.Println("error")
            return plugins, err
        }
        plugins = append(plugins, plugin)
    }
    fmt.Println("loadAllPlugins \tend")
    return plugins, nil
}

func loadAllUsers(db *gorm.DB) (user User, err error) {
    if result := db.Find(&user); result.Error != nil {
        log.Printf("loadAllUsers: %q\n", result.Error)
        return user, result.Error
    }
    return user, nil
}

func loadUserByEmail(db *gorm.DB, email string) (user User, err error) {
    if result := db.Select("id, name, password, email").Where("email = ?", email).First(&user); result.Error != nil {
        log.Printf("loadUserByEmail: 1. %q\n", result.Error)
        return user, result.Error
    }
    if count := len(user.Plugins); count == 0 {
        user.Plugins, err = loadAllPlugins(db, user.ID)
        if err != nil {
            return user, err
        }
    }
    return user, nil
}

func createTables(db *gorm.DB) {
    if check := db.Migrator().HasTable(&User{}); !check {
        if err := db.Migrator().CreateTable(&User{}); err != nil {
            log.Printf("createTables: 1. %q\n", err)
            return
        }
    }
    if check := db.Migrator().HasTable(&Plugin{}); !check {
        if err := db.Migrator().CreateTable(&Plugin{}); err != nil {
            log.Printf("createTables: 1. %q\n", err)
            return
        }
    }
}

func openDB() (db *gorm.DB, err error) {
    return gorm.Open(sqlite.Open("wisehub.db"), &gorm.Config{})
}

func init() {
    db, err := openDB()
    if err != nil {
        log.Printf("init %q: %v\n", err, db)
        return
    }
    createTables(db)
}

/* Please note:
Input either zero, one or two arguments
zero loads all users from database
one loads user based on email address
two loads user based on email and password
 @param email string
 @param password string
*/
func (ds *Datastore) Load(email ...string) (user User, err error) {
    // Get any parameters passed to us out of the args variable into "real"
    // variables we created for them.
    db, err := openDB()
    if err != nil {
        log.Printf("init %q: %v\n", err, db)
        return user, err
    }
    switch len(email) {
    case 0:
        fmt.Println("Load \tat 0")
        return loadAllUsers(db)
    case 1:
        fmt.Println("Load \tat 1")
        return loadUserByEmail(db, email[0])
    default:
        fmt.Printf("too many arguments in function load: %v\n", len(email))
        return user, nil
    }
}

func (ds *Datastore) Save(name string, password string, email string) error {
    fmt.Println("Save \tin Save")
    db, err := openDB()

    if err != nil {
        log.Printf("Save %q: %v\n", err, db)
        return err
    }
    user := User{Name: name, Password: password, Email: email, Plugins: defaultPlugins}
    if result := db.Create(&user); result.Error != nil {
        log.Printf("Save %q: %v\n", err, db)
        return err
    } // pass pointer of data to Create
    fmt.Println("Save \tend Save")
    return nil
}
// TODO: Check
func (ds *Datastore) Update(option string, data map[string]interface{}) error {
    fmt.Println("Update \tin Update")
    fmt.Printf("%s\n", data)
    fmt.Printf("%s\n", option)
    db, err := openDB()
    if err != nil {
        log.Printf("Update %q: %v\n", err, db)
        return err
    }
    switch option {
    case "email":
        fmt.Println(data["new"].(string))
        db.Model(User{}).Where("email = ?", data["old"]).Updates(User{Email: data["new"].(string)})
    case "password":
        fmt.Println(data["password"].(string))
        db.Model(User{}).Where("email = ?",  data["email"]).Updates(User{Password: data["password"].(string)})
    case "plugins":
        fmt.Println(data["plugins"].([]Plugin))
        db.Model(User{}).Where("email = ?",  data["email"]).Updates(User{Plugins: data["plugins"].([]Plugin)})
    }
    return nil
}

func (ds *Datastore) Delete(email string) error {
    db, err := openDB()
    if err != nil {
        log.Printf("init %q: %v\n", err, db)
        return err
    }
    fmt.Println(email)
    db.Where("email = ?", email).Delete(&email)
    return nil
}

