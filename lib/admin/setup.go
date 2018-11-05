package admin

import (
	"log"
	"net/http"

	"github.com/haydenmuhl/hambot/lib/database"
)

const initialSetupFormTmpl = `
<html>
<head>
<title>Initial Setup</title>
</head>
<body>
<p>Before I can connect to Twitch chat, I need a bit more information...</p>
<form action="/bootstrap/submit" method="post">
  <div>
    <label for="username">Bot Username:</label>
    <input type="text" id="username" name="username">
  </div>
  <div>
  	<label for="password">OAuth Token:</label>
  	<input type="text" id="password" name="password">
  </div>
  <div>
  	<button type="submit">Submit</button>
  </div>
</form>
</body>
</html>
`

func initialSetupForm(r http.ResponseWriter, req *http.Request) {
	r.Write([]byte(initialSetupFormTmpl))
}

func processSetupForm(r http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	username := req.Form["username"][0]
	password := req.Form["password"][0]

	// TODO: Add validation
	log.Printf("Writing credentials to the database for user %s\n", username)
	db := database.Handle()
	result, err := db.Exec("INSERT INTO bot_credentials (id, username, password) VALUES (1, ?, ?)", username, password)
	if err != nil {
		log.Println(err)
		log.Println(result)
		r.WriteHeader(http.StatusInternalServerError)
		r.Write([]byte("Something went wrong. :-("))
		return
	}

	log.Println("Redirecting to turn the bot on")
	r.Header().Set("Location", "/on")
	r.WriteHeader(http.StatusFound)
	return
}
