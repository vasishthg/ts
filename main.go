package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
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

type Travel struct {
	ID          int      `json:"id"`
	User        string   `json:"user"`
	From        string   `json:"from"`
	To          string   `json:"to"`
	Train       int      `json:"train"`
	Seats       []string `json:"seats"`
	Passengers  int      `json:"passengers"`
	Start       string   `json:"start"`
	Duration    string   `json:"duration"`
	Date        string   `json:"date"`
	Hotel       bool     `json:"hotel"`
	Cost        int      `json:"cost"`
	Food        []string `json:"food"`
	Status      string   `json:"status"`
	Hours       int      `json:"hours"`
	Minutes     int      `json:"minutes"`
	D1          int      `json:"d1"`
	D2          int      `json:"d2"`
	Et1         int      `json:"et1"`
	Et2         int      `json:"et2"`
	Timeleft    int      `json:"timeleft"`
	Timeelapsed int      `json:"timeelapsed"`
}

func main() {
	router := gin.Default()
	store := cookie.NewStore([]byte("3214cf255f0728c909157b4395b5fce95a67e46051f2a4138ac5fb573ff0444a"))
	router.Use(sessions.Sessions("ts", store))
	stripe.Key = "sk_test_51Md9OfSElK6vRIkdeUZfai2ParU3JrKolpmWR2KPJRXPHYVTlKZsTQp88bjFBXEWG647LN82YslGMdcc6NL0CiFR002bYVZK1w"

	db, err := sql.Open("mysql", "root:toor@tcp(localhost:3306)/ts")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		date := time.Now().Format("2006-01-02")
		session := sessions.Default(c)
		loggedin := false
		if session.Get("loggedin") == true {
			loggedin = true
			User := User{}
			err = db.QueryRow("SELECT * FROM users WHERE email = ?", session.Get("email")).Scan(&User.ID, &User.Name, &User.Email, &User.Phone, &User.DOB, &User.Password, &User.History, &User.Aadhar, &User.Location, &User.Balance, &User.Points, &User.Ongoing, &User.Events, &User.Reccomended)
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

			rows, err = db.Query("SELECT id, user, `from`, `to`, train, seats, passengers, start, duration, date, hotel, cost, food, status FROM travel WHERE user = ?", session.Get("email"))
			if err != nil {
				fmt.Println("e2:", err.Error())
			}

			var travels []Travel
			var hoursList []int
			var minutesList []int
			type durhrmin struct {
				Hours   int `json:"hours"`
				Minutes int `json:"minutes"`
			}
			type hrmin struct {
				Hours   int `json:"hours"`
				Minutes int `json:"minutes"`
			}

			for rows.Next() {
				var seatsEncoded, foodEncoded string
				var seats []string
				var food []string
				travel := Travel{}
				err = rows.Scan(&travel.ID, &travel.User, &travel.From, &travel.To, &travel.Train, &seatsEncoded, &travel.Passengers, &travel.Start, &travel.Duration, &travel.Date, &travel.Hotel, &travel.Cost, &foodEncoded, &travel.Status)
				if err != nil {
					fmt.Println("e1:", err.Error())
					continue
				}

				if err := json.Unmarshal([]byte(seatsEncoded), &seats); err != nil {
					fmt.Println("e-seats:", err.Error())
				}
				travel.Seats = seats

				if err := json.Unmarshal([]byte(foodEncoded), &food); err != nil {
					fmt.Println("e-food:", err.Error())
				}
				travel.Food = food

				t, err := time.Parse("15:04:05", travel.Start)
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				d, err := time.Parse("15:04:05", travel.Duration)
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				hours := t.Hour()
				minutes := t.Minute()
				Hrmin := hrmin{Hours: hours, Minutes: minutes}
				hoursList = append(hoursList, hours)
				minutesList = append(minutesList, minutes)
				travel.Hours = hours
				travel.Minutes = minutes
				durhr := d.Hour()
				durmin := d.Minute()
				Durhrmin := durhrmin{Hours: durhr, Minutes: durmin}
				travel.D1 = Durhrmin.Hours
				travel.D2 = Durhrmin.Minutes
				endTimehr := durhr + hours
				endTimem := durmin + minutes
				fmt.Println("iewj", endTimehr, endTimem)
				travel.Et1 = endTimehr
				travel.Et2 = endTimem
				fmt.Printf("Travel ID: %d - Hours: %02d, Minutes: %02d\n", travel.ID, Hrmin.Hours, Hrmin.Minutes)
				if travel.Date > date {
					travel.Status = "Upcoming"
				} else if travel.Date == date {
					fmt.Println("same date vhai")
					travelStartTime, err := time.Parse("15:04:05", travel.Start)
					if err != nil {
						fmt.Println(err.Error())
					}
					if travelStartTime.Hour() > time.Now().Hour() {
						fmt.Println("ek hogya")
						travel.Status = "Upcoming"
					} else if travelStartTime.Hour() <= time.Now().Hour() && travel.Et1 > time.Now().Hour() {
						travel.Status = "Ongoing"
					}
				} else {
					travel.Status = "Completed"
				}
				travel.Timeelapsed = time.Now().Hour() - travel.Hours
				travel.Timeleft = travel.Et1 - time.Now().Hour()
				travels = append(travels, travel)

			}
			fmt.Println(travels)

			fmt.Println("Hours List:", hoursList)
			fmt.Println("Minutes List:", minutesList)
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
				"travels":    travels,
				"hours":      hoursList,
				"minutes":    minutesList,
				"date":       date,
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

	router.POST("/ajax/addlocation", func(c *gin.Context) {
		location := c.PostForm("location")
		session := sessions.Default(c)
		email := session.Get("email")
		_, err = db.Exec("UPDATE users SET location = ? WHERE email = ?", location, email)
		if err != nil {
			fmt.Println(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
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
	router.POST("/ajax/getprice", func(c *gin.Context) {
		id := c.PostForm("id")
		var price int
		err = db.QueryRow("SELECT price FROM food WHERE id = ?", id).Scan(&price)
		if err != nil {
			fmt.Println(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"price": price,
		})
	})
	var price float64
	router.GET("/checkout", func(c *gin.Context) {
		sessions := sessions.Default(c)
		fmt.Println("price recieved by server: ", price)
		domain := "http://127.0.0.1:5000"
		params := &stripe.CheckoutSessionParams{
			LineItems: []*stripe.CheckoutSessionLineItemParams{
				{
					PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
						Currency: stripe.String(string(stripe.CurrencyINR)),
						ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
							Name: stripe.String("Indian Railways"),
						},
						UnitAmountDecimal: stripe.Float64(price * 100),
					},
					Quantity: stripe.Int64(1),
				},
			},
			CustomerEmail: stripe.String(sessions.Get("email").(string)),
			Currency:      stripe.String(string(stripe.CurrencyINR)),
			Mode:          stripe.String("payment"),
			SuccessURL:    stripe.String(domain + "/checkout/success?session_id={CHECKOUT_SESSION_ID}"),
			CancelURL:     stripe.String(domain + "/checkout/cancelled"),
		}
		s, _ := session.New(params)
		fmt.Println("params", s)
		c.Redirect(http.StatusFound, s.URL)
	})

	router.POST("/ajax/bookinfo", func(c *gin.Context) {
		priceStr := c.PostForm("price")
		session := sessions.Default(c)
		session.Set("book-from", c.PostForm("from"))
		session.Set("book-to", c.PostForm("to"))
		session.Set("book-train", c.PostForm("train"))
		session.Set("book-seats", c.PostForm("seats"))
		session.Set("book-passengers", c.PostForm("passengers"))
		session.Set("book-start", c.PostForm("start"))
		session.Set("book-duration", c.PostForm("duration"))
		session.Set("book-date", c.PostForm("date"))
		session.Set("book-hotel", c.PostForm("hotel"))
		session.Set("price", priceStr)
		session.Set("book-food", c.PostForm("food"))
		session.Save()
		fmt.Println(c.PostForm("seats"))
		fmt.Println(c.PostForm("food"))

		price, err = strconv.ParseFloat(priceStr, 64)
		if err != nil {
			fmt.Println("error parsing price:", err.Error())
			return
		}

		fmt.Println("price converted by server: ", price)
	})
	router.GET("/checkout/success", func(c *gin.Context) {
		session := sessions.Default(c)
		sessionID := c.Query("session_id")
		if sessionID == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "no session id",
			})
			return
		}
		fmt.Println(session.Get("email"))
		fmt.Println(session.Get("book-from"))
		fmt.Println(session.Get("book-to"))
		fmt.Println(session.Get("book-train"))
		fmt.Println(session.Get("book-seats"))
		fmt.Println(session.Get("book-passengers"))
		fmt.Println(session.Get("book-start"))
		fmt.Println(session.Get("book-duration"))
		fmt.Println(session.Get("book-date"))
		fmt.Println(session.Get("book-hotel"))
		fmt.Println(session.Get("price"))
		fmt.Println(session.Get("book-food"))
		_, err = db.Exec("INSERT INTO travel (`user`, `from`, `to`, `train`, `seats`, `passengers`, `start`, `duration`, `date`, `hotel`, `cost`, `food`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)", session.Get("email"), session.Get("book-from"), session.Get("book-to"), session.Get("book-train"), session.Get("book-seats"), session.Get("book-passengers"), session.Get("book-start"), session.Get("book-duration"), session.Get("book-date"), session.Get("book-hotel"), session.Get("price"), session.Get("book-food"))
		if err != nil {
			fmt.Println(err.Error())
		}
		var points int
		err = db.QueryRow("SELECT points FROM users WHERE email = ?", session.Get("email")).Scan(&points)
		if err != nil {
			fmt.Println(err.Error())
		}
		points += 100
		fmt.Println("points:", points)
		_, err = db.Exec("UPDATE users SET points = ? WHERE email = ?", points, session.Get("email"))
		if err != nil {
			fmt.Println(err.Error())
		}

		c.HTML(http.StatusOK, "success.html", gin.H{
			"session_id": sessionID,
		})
		fmt.Println(sessionID)
	})
	router.GET("/checkout/cancelled", func(c *gin.Context) {

		c.HTML(http.StatusOK, "cancelled.html", gin.H{})

	})
	router.GET("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.Redirect(http.StatusFound, "/")
	})

	router.Run(":5000")
}
