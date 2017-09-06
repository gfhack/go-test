package main

import (
	"time"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID        int       `db:"user_id, primarykey, autoincrement" json:"id"`
	Email     string    `db:"email" json:"email"`
	Name      string    `db:"username" json:"username"`
	Profile   string    `db:"profile" json:"profile"`
	Team      string    `db:"team" json:"team"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

func main() {
	// Disable Console Color
	// gin.DisableConsoleColor()
	router := gin.Default()

	// Ping test
	router.GET("/users", func(c *gin.Context) {
		db, err := sql.Open("postgres", "host=localhost port=5432 user=macbook dbname=postgres sslmode=disable")

		rows, err := db.Query("SELECT users.id as user_id, users.name as username, users.email as email, profiles.title as profile, teams.name as team, users.created_at, users.updated_at FROM users JOIN profiles ON users.profile_id = profiles.id JOIN teams ON users.team_id = teams.id")

		if err != nil {
			c.JSON(406, gin.H{"status": err.Error()})
		}

		defer rows.Close()

		var users []*User;

		for rows.Next() {
			user := new(User)
			err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Profile, &user.Team, &user.CreatedAt, &user.UpdatedAt)
			if err != nil {
				c.JSON(406, gin.H{"status": err.Error()})
			}
			users = append(users, user)
		}

		c.JSON(200, gin.H{"data": users})
	})

	// Get user value
	// router.GET("/user/:name", func(c *gin.Context) {
	// 	user := c.Params.ByName("name")
	// 	value, ok := DB[user]
	// 	if ok {
	// 		c.JSON(200, gin.H{"user": user, "value": value})
	// 	} else {
	// 		c.JSON(200, gin.H{"user": user, "status": "no value"})
	// 	}
	// })

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	// authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
	// 	"foo":  "bar", // user:foo password:bar
	// 	"manu": "123", // user:manu password:123
	// }))
	//
	// authorized.POST("admin", func(c *gin.Context) {
	// 	user := c.MustGet(gin.AuthUserKey).(string)
	//
	// 	// Parse JSON
	// 	var json struct {
	// 		Value string `json:"value" binding:"required"`
	// 	}
	//
	// 	if c.Bind(&json) == nil {
	// 		DB[user] = json.Value
	// 		c.JSON(200, gin.H{"status": "ok"})
	// 	}
	// })

	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}
