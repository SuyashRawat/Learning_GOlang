package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-crud/packages/models"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUsers(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET Method Called")

		rows, err := database.Query("SELECT * FROM users")
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		defer rows.Close()

		/*
			rows,Next() with rows.Scan()
				rows,Next() advances to the next row:
					Moves the cursor to the next row in the result set.
					Prepares the row's data to be accessed.
				rows.Scan() extracts the data for the current row:
					You pass pointers to variables where you want the data to be stored.
					For example, if a row has id, name, and email columns, you can scan them into variables like this:
		*/

		allUsers := []models.User{}

		for rows.Next() {
			var oneUser models.User
			err := rows.Scan(&oneUser.ID, &oneUser.Name, &oneUser.Email)
			if err != nil {
				log.Fatal(err)
				panic(err)
			}
			allUsers = append(allUsers, oneUser)
		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Println(allUsers)

		json.NewEncoder(w).Encode(allUsers)
		// responsible for converting the allUsers data (which is typically a slice of structs, e.g., []User) into JSON format and sending it directly as an HTTP response to the client.
		/*
			json.NewEncoder(w):
			- Creates a new JSON encoder that writes directly to the http.ResponseWriter (w in this case).
			- w represents the HTTP response stream, so anything written to it is sent as part of the HTTP response.
			.Encode(allUsers):
			- Converts allUsers (a Go data structure) into JSON format.
			- The encoded JSON is immediately written to the http.ResponseWriter.
			Efficient Combination:
			- The json.Encoder combines the conversion (marshal) and writing steps into one operation, making it convenient for HTTP responses.
		*/

		/*
			-> If we didn't want to use the json.NewEncoder(w).Encode(allUsers) method, we could have used the following code to achieve the same result:
			jsonData, err := json.Marshal(allUsers)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)

			-> json.Marshal converts the allUsers struct into a JSON byte slice ([]byte), but it doesn't handle writing the data to the HTTP response.
		*/
	}
}

func GetUserByID(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET Method By ID Called")

		vars := mux.Vars(r)
		id := vars["id"]

		var oneUser models.User
		err := database.QueryRow("Select * FROM users WHERE id = $1", id).Scan(&oneUser.ID, &oneUser.Name, &oneUser.Email)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			// This is a better way to handle the error
			// log.Fatal(err)
			// panic(err)
		}

		jsonData, err := json.Marshal(oneUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
	}
}

func CreateUser(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("POST Method Called")

		var newUser models.User

		// Read the entire request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Unmarshal the JSON into the User struct
		err = json.Unmarshal(body, &newUser)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// Insert the user into the database and get the new user ID
		err = database.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", newUser.Name, newUser.Email).Scan(&newUser.ID)
		if err != nil {
			// Log the error details for debugging
			fmt.Printf("Error saving user: %v\n", err) // Detailed error log

			http.Error(w, "Error saving user", http.StatusInternalServerError)
			return
		}

		// Convert the User struct to JSON
		jsonData, err := json.Marshal(newUser)
		if err != nil {
			http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
			return
		}

		// Write the JSON response
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}

// func createUser(database *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("POST Method Called")

// 		var newUser User

// 		// Read the entire request body
// 		body, err := io.ReadAll(r.Body)
// 		if err != nil {
// 			http.Error(w, "Invalid request body", http.StatusBadRequest)
// 			return
// 		}
// 		defer r.Body.Close()

// 		// Unmarshal the JSON into the User struct
// 		err = json.Unmarshal(body, &newUser)
// 		if err != nil {
// 			http.Error(w, "Invalid JSON", http.StatusBadRequest)
// 			return
// 		}

// 		// Insert the user into the database and get the new user ID
// 		err = database.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", newUser.Name, newUser.Email).Scan(&newUser.ID)
// 		/*
// 			- INSERT INTO users (name, email): This specifies the users table and the columns (name and email) where we want to insert data.
// 			- VALUES ($1, $2): These are placeholders for parameters. The db.QueryRow function will substitute $1 with u.Name and $2 with u.Email.
// 				- Parameterized queries are used to prevent SQL injection.
// 			- RETURNING id: This tells PostgreSQL to return the id of the newly inserted row. This is useful for capturing auto-generated fields (like SERIAL primary keys).
// 		*/

// 		if err != nil {
// 			http.Error(w, "Error saving user", http.StatusInternalServerError)
// 			return
// 		}

// 		// Convert the User struct to JSON
// 		jsonData, err := json.Marshal(newUser)
// 		if err != nil {
// 			http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
// 			return
// 		}

// 		// Write the JSON response
// 		w.Write(jsonData)

// 		/*
// 			Use json.NewDecoder if you are dealing with a stream of JSON data or want to process it efficiently without loading everything into memory at once.
// 			Use json.Unmarshal if you know the payload size is small and prefer a simpler approach.
// 		*/
// 	}
// }

/*
	-> POST Request Using json.NewDecoder and json.NewEncoder
	func createUser(db *sql.DB) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var u User
			json.NewDecoder(r.Body).Decode(&u)

			->  THe incoming JSON data stream from the request body is decoded into the User struct u
				The struct u is being used to store the data that is expected to come in the request body of the POST request in JSON format. Here's a step-by-step explanation:
				The r.Body contains the raw JSON data stream.
				.Decode(&u) maps the JSON keys to the fields in the User struct u.
				JSON Key "name" maps to u.Name.
				JSON Key "email" maps to u.Email.
				The reason json.NewDecoder is used here is to extract the data from the incoming JSON body. Without this, the data from the request body would remain in raw JSON format and wouldn't be usable in the Go application.

			err := db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", u.Name, u.Email).Scan(&u.ID)
			if err != nil {
				log.Fatal(err)
			}

			json.NewEncoder(w).Encode(u)
		}
	}
*/

func UpdateUserByID(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("PUT Method Called")

		vars := mux.Vars(r)
		id := vars["id"]

		var updatedUser models.User

		body, errorr := io.ReadAll(r.Body)
		if errorr != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		err := json.Unmarshal(body, &updatedUser)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		row, _ := database.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", updatedUser.Name, updatedUser.Email, id)
		rowsAffected, _ := row.RowsAffected()
		if rowsAffected == 0 {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		} else {
			fmt.Fprintf(w, "%v rows updated", rowsAffected)
		}

		errr := database.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&updatedUser.ID, &updatedUser.Name, &updatedUser.Email)
		if errr != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		jsonData, err := json.Marshal(updatedUser)

		if err != nil {
			http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
			return
		}

		w.Write(jsonData)
	}
}

/*
	func updateUser(db *sql.DB) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var u User
			json.NewDecoder(r.Body).Decode(&u)

			vars := mux.Vars(r)
			id := vars["id"]

			_, err := db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", u.Name, u.Email, id)
			if err != nil {
				log.Fatal(err)
			}

			json.NewEncoder(w).Encode(u)
		}
	}
*/

func DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("DELETE Method Called")

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		vars := mux.Vars(r)
		id := vars["id"]

		var u models.User
		errr := json.Unmarshal(body, &u)
		if errr != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		queryError := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&u.ID, &u.Name, &u.Email)
		if queryError != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
			if err != nil {
				//todo : fix error handling
				w.WriteHeader(http.StatusNotFound)
				return
			}

			jsonData, err := json.Marshal("User deleted")
			if err != nil {
				http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
				return
			}

			w.Write(jsonData)
		}
	}
}

/*
	func deleteUser(db *sql.DB) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			id := vars["id"]

			var u User
			err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&u.ID, &u.Name, &u.Email)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				return
			} else {
				_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
				if err != nil {
					//todo : fix error handling
					w.WriteHeader(http.StatusNotFound)
					return
				}

				json.NewEncoder(w).Encode("User deleted")
			}
		}
	}
*/
