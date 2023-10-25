package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	DOB         []uint8 `json:"dob"`
	Password    string  `json:"password"`
	Aadhar      string  `json:"aadhar"`
	Location    string  `json:"location"`
	History     []byte  `json:"history"`
	Balance     int     `json:"balance"`
	Points      int     `json:"points"`
	Ongoing     []byte  `json:"ongoing"`
	Events      []byte  `json:"events"`
	Reccomended []byte  `json:"reccomended"`
}
type Station struct {
	Code     string `json:"code"`
	Sname    string `json:"sname"`
	Location string `json:"location"`
}

func main() {
	router := gin.Default()
	store := cookie.NewStore([]byte("3214cf255f0728c909157b4395b5fce95a67e46051f2a4138ac5fb573ff0444a"))
	router.Use(sessions.Sessions("ts", store))

	db, err := sql.Open("mysql", "root:toor@tcp(localhost:3306)/ts")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		loggedin := false
		if session.Get("loggedin") == true {
			loggedin = true
			User := User{}
			err = db.QueryRow("SELECT * FROM users ").Scan(&User.ID, &User.Name, &User.Email, &User.Phone, &User.DOB, &User.Password, &User.History, &User.Aadhar, &User.Location, &User.Balance, &User.Points, &User.Ongoing, &User.Events, &User.Reccomended)
			l1 := []rune(User.Name)
			l2 := string(l1[0:1])
			rows, err := db.Query("SELECT code, sname, location FROM stations")
			if err != nil {
				fmt.Println("e2:", err.Error())
			}
			stationcodes := []string{}
			stationname := []string{}
			stationlocation := []string{}
			stations := []Station{}
			for rows.Next() {
				Station := Station{}
				err = rows.Scan(&Station.Code, &Station.Sname, &Station.Location)
				if err != nil {
					fmt.Println("e1:", err.Error())
				}
				stationcodes = append(stationcodes, Station.Code)
				stationname = append(stationname, Station.Sname)
				stationlocation = append(stationlocation, Station.Location)
				stations = append(stations, Station)
			}
			type Food struct {
				ID        int    `json:"id"`
				Title     string `json:"title"`
				Price     int    `json:"price"`
				VegNonVeg string `json:"vegnonveg"`
			}
			rows, err = db.Query("SELECT id, title, price, vegnonveg FROM food")
			var fppd []Food
			for rows.Next() {
				food := Food{}
				err = rows.Scan(&food.ID, &food.Title, &food.Price, &food.VegNonVeg)
				if err != nil {
					fmt.Println("e1:", err.Error())
				}
				fppd = append(fppd, food)
			}

			if err != nil {
				fmt.Println("e2:", err.Error())
			}

			c.HTML(http.StatusOK, "index.html", gin.H{
				"loggedin":   loggedin,
				"name":       User.Name,
				"email":      User.Email,
				"phone":      User.Phone,
				"dob":        User.DOB,
				"aadhar":     User.Aadhar,
				"location":   User.Location,
				"balance":    User.Balance,
				"points":     User.Points,
				"history":    User.History,
				"ongoing":    User.Ongoing,
				"events":     User.Events,
				"recc":       User.Reccomended,
				"l2":         l2,
				"scodes":     stationcodes,
				"snames":     stationname,
				"slocations": stationlocation,
				"stations":   stations,
				"food":       fppd,
			})
		} else {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"loggedin": loggedin,
			})
		}
	})

	router.POST("/ajax/signup", func(c *gin.Context) {
		fname := (c.PostForm("fname"))
		lname := (c.PostForm("lname"))
		email := (c.PostForm("email"))
		phone := (c.PostForm("phone"))
		dob := (c.PostForm("dob"))
		password := (c.PostForm("password"))
		file, err := c.FormFile("file")
		fmt.Println(fname, lname, email, phone, dob, password)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(file.Filename)
		User := User{}
		rows, err := db.Query("SELECT * FROM users WHERE email = ? OR phone = ?", email, phone)
		if err != nil {
			fmt.Println(err.Error())
		}
		emailExists := false
		phoneExists := false
		for rows.Next() {
			err = rows.Scan(&User.ID, &User.Name, &User.Email, &User.Phone, &User.DOB, &User.Password, &User.Aadhar, &User.Location, &User.History, &User.Balance, &User.Points, &User.Ongoing, &User.Events, &User.Reccomended)
			if err != nil {
				fmt.Println(err.Error())
			}
			if User.Email == email {
				emailExists = true
			}
			if User.Phone == phone {

				phoneExists = true
			}
		}

		if emailExists && phoneExists {
			c.JSON(http.StatusConflict, gin.H{
				"status": "email and phone exists",
			})
			return
		}

		if !emailExists && !phoneExists {
			c.SaveUploadedFile(file, "./uploads/aadhar/"+lname+fname+"/"+phone+file.Filename)

			_, err = db.Exec("INSERT INTO users (name, email, phone, dob, password, aadhar) VALUES (?, ?, ?, ?, ?, ?)", fname+" "+lname, email, phone, dob, password, "./uploads/aadhar/"+lname+fname+"/"+phone+file.Filename)
			if err != nil {
				fmt.Println(err.Error())
			}
			session := sessions.Default(c)
			session.Set("loggedin", true)
			session.Set("email", email)
			session.Set("phone", phone)
			session.Set("name", fname+" "+lname)

			session.Save()
			c.JSON(http.StatusOK, gin.H{
				"status": "success",
			})
			return
		}
		if emailExists {
			c.JSON(http.StatusConflict, gin.H{
				"status": "email exists",
			})
			return
		}

		if phoneExists {
			c.JSON(http.StatusConflict, gin.H{
				"status": "phone exists",
			})
			return
		}

	})

	router.POST("/ajax/login", func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")
		User := User{}
		fmt.Println(email, password)
		err = db.QueryRow("SELECT (password) FROM users WHERE email = ?", email).Scan(&User.Password)
		if err != nil {
			fmt.Println(err.Error())
		}
		if err == sql.ErrNoRows {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Invalid Credentials",
			})
		}
		if User.Password == password {
			session := sessions.Default(c)
			session.Set("email", email)
			session.Set("loggedin", true)
			session.Save()
			c.JSON(http.StatusOK, gin.H{
				"status": "authorised",
			})
		} else {
			c.JSON(http.StatusConflict, gin.H{
				"error": "Invalid Credentials",
			})
		}
	})

	router.POST("/ajax/getstation", func(c *gin.Context) {
		code := c.PostForm("id")
		fmt.Println("cie:", code)
		Station := Station{}
		err = db.QueryRow("SELECT code, sname, location FROM stations WHERE code = ?", code).Scan(&Station.Code, &Station.Sname, &Station.Location)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(Station)
		c.JSON(http.StatusOK, gin.H{
			"code":     Station.Code,
			"sname":    Station.Sname,
			"location": Station.Location,
		})
	})

	router.POST("ajax/tripdata", func(c *gin.Context) {
		origin := c.PostForm("origin")
		type train struct {
			ID       int     `json:"id"`
			Stations []byte  `json:"stations"`
			Occupied []uint8 `json:"occupied"`
		}
		Train := train{}
		destination := c.PostForm("destination")
		err = db.QueryRow("SELECT id, stations, occupied FROM trains WHERE `from` = ? AND `to` = ?", origin, destination).Scan(&Train.ID, &Train.Stations, &Train.Occupied)
		if err != nil {
			fmt.Println("error retireving rows: ", err.Error())
		}

		if err == sql.ErrNoRows {
			rows, err := db.Query("SELECT code FROM stations WHERE code != ? AND code != ? ORDER BY RAND() LIMIT 6", origin, destination)
			if err != nil {
				fmt.Println("error retireveringcodes:", err.Error())
			}
			var stationsJSON []byte
			var stations []string
			for rows.Next() {
				var station string
				err = rows.Scan(&station)
				if err != nil {
					fmt.Println("error scanning rows:", err.Error())
				}
				stations = append(stations, station)
				stationsJSON, err = json.Marshal(stations)
				if err != nil {
					fmt.Println("error marshalling stations:", err.Error())
				}

			}
			seats, err := json.Marshal([]int{37, 13, 22, 32, 6, 23, 43, 42, 17, 1, 14, 29, 24, 20, 18, 39, 8, 55, 56, 45, 59, 9, 51, 5, 3, 16, 7, 28, 60, 36, 30, 38, 26, 52, 35, 47, 2, 44, 21, 54, 11, 46, 25, 33})
			if err != nil {
				fmt.Println("error marshalling seats:", err.Error())
			}

			_, err = db.Exec("INSERT INTO trains (`from`, `to`, `stations`, `occupied`) VALUES (?, ?, ?, ?)", origin, destination, stationsJSON, seats)
			if err != nil {
				fmt.Println("error inserting::", err.Error())
			}
			err = db.QueryRow("SELECT id, stations, occupied FROM trains WHERE `from` = ? AND `to` = ?", origin, destination).Scan(&Train.ID, &Train.Stations, &Train.Occupied)
			if err != nil {
				fmt.Println("error retireving rows: ", err.Error())
			}
		}
		var stations []string
		if err := json.Unmarshal(Train.Stations, &stations); err != nil {
			fmt.Println(err.Error())
			return
		}
		var occupied []int
		if err := json.Unmarshal(Train.Occupied, &occupied); err != nil {
			fmt.Println(err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id":       Train.ID,
			"stations": stations,
			"occupied": occupied,
			"len":      len(stations),
		})
	})
	router.POST("/ajax/gettrain", func(c *gin.Context) {
		origincode := c.PostForm("origin")
		destinationcode := c.PostForm("destination")
		var id int
		err := db.QueryRow("SELECT id FROM trains WHERE `from` = ? and `to` = ?", origincode, destinationcode).Scan(&id)
		if err != nil {
			fmt.Println(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	})
	router.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.Redirect(http.StatusFound, "/")
	})
	router.Run(":5000")
}
