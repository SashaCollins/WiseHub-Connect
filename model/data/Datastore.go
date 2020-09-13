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

func init() {
    db, err := openDB()
    if err != nil {
        log.Printf("init %q: %v\n", err, db)
        return
    }
    createTables(db)
}

func openDB() (db *gorm.DB, err error) {
    return gorm.Open(sqlite.Open("wisehub.db"), &gorm.Config{})
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


func loadAllPlugins(db *gorm.DB, userID uint) (plugins []Plugin, err error) {
    rows, err := db.Model(&Plugin{}).Where("user_id = ?", userID).Rows()
    if err != nil {
        log.Printf("loadAllPlugins1: %q\n", err)
        return plugins, err
    }
    defer rows.Close()

    var plugin Plugin
    for rows.Next() {
        if err = db.ScanRows(rows, &plugin); err != nil {
            return plugins, err
        }
        plugins = append(plugins, plugin)
    }
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
    user.Plugins, err = loadAllPlugins(db, user.ID)
    if err != nil {
            return user, err
        }
    return user, nil
}

func updatePlugins(db *gorm.DB, userEmail string, updatedPlugins []Plugin) error {
    var user User
    if result := db.Select("id").Where("email = ?", userEmail).First(&user); result.Error != nil {
        log.Printf("updatePlugins: 1. %q\n", result.Error)
        return result.Error
    }
    dbPlugins, err := loadAllPlugins(db, user.ID)
    if err != nil {
        log.Printf("updatePlugins: 1. %q\n", err)
        return err
    }
    for iDBP := range dbPlugins {
        for iUP := range updatedPlugins {
            if dbPlugins[iDBP].PluginName == updatedPlugins[iUP].PluginName {
                db.Model(&Plugin{}).Where("user_id = ? AND plugin_name = ?", user.ID, updatedPlugins[iUP].PluginName).Updates(Plugin{
                    UsernameHost: updatedPlugins[iUP].UsernameHost,
                    Token: updatedPlugins[iUP].Token,
                })
            }
        }

    }
    return nil
}

/* Please note:
Input either zero, one or two arguments
zero loads all users from database
one loads user based on email address
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
        return loadAllUsers(db)
    case 1:
        return loadUserByEmail(db, email[0])
    default:
        fmt.Printf("too many arguments in function load: %v\n", len(email))
        return user, nil
    }
}

func (ds *Datastore) Save(name string, password string, email string) error {
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
    return nil
}

func (ds *Datastore) Update(option string, data map[string]interface{}) error {
    fmt.Println("Update \tin Update")
    fmt.Printf("Data: %s\n", data)
    fmt.Printf("Options: %s\n", option)
    db, err := openDB()
    if err != nil {
        log.Printf("Update %q: %v\n", err, db)
        return err
    }
    switch option {
    case "email":
        db.Model(User{}).Where("email = ?", data["old"]).Updates(User{Email: data["new"].(string)})
    case "password":
        db.Model(User{}).Where("email = ?",  data["email"]).Updates(User{Password: data["password"].(string)})
    case "plugins":
        if err := updatePlugins(db, data["email"].(string), data["updatedPlugins"].([]Plugin)); err != nil {
            log.Printf("Update %q: %v\n", err, db)
            return err
        }
    }
    return nil
}

func (ds *Datastore) Delete(email string) error {
    db, err := openDB()
    if err != nil {
        log.Printf("init %q: %v\n", err, db)
        return err
    }
    db.Where("email = ?", email).Delete(&email)
    return nil
}

