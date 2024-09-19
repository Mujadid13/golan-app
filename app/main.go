package main

import (
    "fmt"
    "net/http"
    "html/template"
    "github.com/go-redis/redis/v8"
    "context"
    "math/rand"
    "time"
)

var ctx = context.Background()

func main() {
    // Initialize Redis client
    redisClient := redis.NewClient(&redis.Options{
        Addr:     "redis-server:6379",
        Password: "",
        DB:       0,
    })

    // Seed the random number generator for animal selection
    rand.Seed(time.Now().UnixNano())

    // Define a slice of animal image URLs from Unsplash
    animalImages := []string{
        "https://images.unsplash.com/photo-1574158622682-e40e69881006?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=MnwxfDB8MXxyYW5kb218MHx8YW5pbWFsc3x8fHx8fHwxNjYyMzYwODM5&ixlib=rb-1.2.1&q=80&w=50",
        "https://images.unsplash.com/photo-1561948955-570b270e7c36?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=MnwxfDB8MXxyYW5kb218MHx8YW5pbWFsc3x8fHx8fHwxNjYyMzYwODM5&ixlib=rb-1.2.1&q=80&w=50",
        "https://images.unsplash.com/photo-1535930749574-1399327ce78f?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=MnwxfDB8MXxyYW5kb218MHx8YW5pbWFsc3x8fHx8fHwxNjYyMzYwODM5&ixlib=rb-1.2.1&q=80&w=50",
        "https://images.unsplash.com/photo-1546182990-dffeafbe841d?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=MnwxfDB8MXxyYW5kb218MHx8YW5pbWFsc3x8fHx8fHwxNjYyMzYwODM5&ixlib=rb-1.2.1&q=80&w=50",
        "https://images.unsplash.com/photo-1474511320723-9a56873867b5?crop=entropy&cs=tinysrgb&fit=max&fm=jpg&ixid=MnwxfDB8MXxyYW5kb218MHx8YW5pbWFsc3x8fHx8fHwxNjYyMzYwODM5&ixlib=rb-1.2.1&q=80&w=50",
    }

    // Create a FuncMap to register the template function for random animal images
    funcMap := template.FuncMap{
        "randomAnimal": func() string {
            return animalImages[rand.Intn(len(animalImages))] // Randomly select an animal image
        },
    }

    // Parse the template with custom function for random animal images
    tmpl := template.Must(template.New("form").Funcs(funcMap).Parse(`
        <html>
        <head>
            <title>Fact Website</title>
            <style>
                body { font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; background: #f0f4f8; color: #333; margin: 0; padding: 0; box-sizing: border-box; }
                .container { max-width: 800px; margin: 20px auto; background: white; padding: 20px; border-radius: 8px; box-shadow: 0 2px 15px rgba(0,0,0,0.1); }
                h1 { color: #2c3e50; }
                h2 { color: #16a085; }
                form { margin-top: 20px; }
                label { display: block; margin-bottom: 10px; font-size: 18px; color: #2980b9; }
                input[type="text"], input[type="submit"] { width: calc(100% - 22px); padding: 10px; border-radius: 4px; border: 1px solid #ccc; }
                input[type="submit"] { background: #27ae60; color: white; padding: 10px 20px; border: none; border-radius: 4px; cursor: pointer; font-size: 16px; }
                input[type="submit"]:hover { background: #2ecc71; }
                ul { list-style-type: none; padding: 0; }
                li { display: flex; justify-content: space-between; align-items: center; padding: 10px; background: #ecf0f1; margin-bottom: 8px; border-radius: 4px; }
                img { margin-right: 10px; width: 50px; height: 50px; border-radius: 50%; }
                .delete-btn { background: #e74c3c; color: white; padding: 5px 10px; border: none; border-radius: 4px; cursor: pointer; margin-left: 10px; }
                .delete-btn:hover { background: #c0392b; }
                .edit-btn { background: #f1c40f; color: white; padding: 5px 10px; border: none; border-radius: 4px; cursor: pointer; margin-left: 10px; }
                .edit-btn:hover { background: #f39c12; }
            </style>
        </head>
        <body>
            <div class="container">
                <h1>Welcome to the Fact Website</h1>
                <form action="/" method="post">
                    <label for="name">Enter your name:</label>
                    <input type="text" id="name" name="name"><br><br>
                    <input type="submit" value="Submit">
                </form>
                <h2>Submitted Names:</h2>
                <ul>
                {{range .}}
                    <li>
                        <span><img src="{{randomAnimal}}" alt="Animal">{{.}}</span>
                        <form action="/delete" method="post" style="display:inline;">
                            <input type="hidden" name="name" value="{{.}}">
                            <input type="submit" value="Delete" class="delete-btn">
                        </form>
                        <form action="/edit" method="get" style="display:inline;">
                            <input type="hidden" name="oldName" value="{{.}}">
                            <input type="submit" value="Edit" class="edit-btn">
                        </form>
                    </li>
                {{end}}
                </ul>
            </div>
        </body>
        </html>
    `))

    // Handler for the root page ("/")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            // On GET request, fetch all names and display
            names, err := redisClient.LRange(ctx, "names", 0, -1).Result()
            if err != nil {
                http.Error(w, "Failed to fetch names", http.StatusInternalServerError)
                return
            }
            tmpl.Execute(w, names)
            return
        }

        // Process form submission on POST
        r.ParseForm()
        name := r.FormValue("name")

        // Save the name to Redis list
        err := redisClient.RPush(ctx, "names", name).Err()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Fetch all names including the new one
        names, err := redisClient.LRange(ctx, "names", 0, -1).Result()
        if err != nil {
            http.Error(w, "Failed to fetch names", http.StatusInternalServerError)
            return
        }

        // Display the names with the new one included
        tmpl.Execute(w, names)
    })

    // Handler for the delete request ("/delete")
    http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            r.ParseForm()
            nameToDelete := r.FormValue("name")

            // Remove the name from the Redis list
            err := redisClient.LRem(ctx, "names", 1, nameToDelete).Err()
            if err != nil {
                http.Error(w, "Failed to delete name", http.StatusInternalServerError)
                return
            }

            // Fetch updated list of names
            names, err := redisClient.LRange(ctx, "names", 0, -1).Result()
            if err != nil {
                http.Error(w, "Failed to fetch names", http.StatusInternalServerError)
                return
            }

            // Re-render the template with the updated list
            tmpl.Execute(w, names)
        }
    })

    // Handler for the edit request ("/edit")
    http.HandleFunc("/edit", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            // Get the old name from the query parameters
            oldName := r.URL.Query().Get("oldName")

            // Display an input field with the old name pre-filled for editing
            editForm := `
            <html>
            <head>
                <title>Edit Name</title>
            </head>
            <body>
                <h1>Edit Name</h1>
                <form action="/update" method="post">
                    <input type="hidden" name="oldName" value="` + oldName + `">
                    <label for="newName">New Name:</label>
                    <input type="text" id="newName" name="newName" value="` + oldName + `"><br><br>
                    <input type="submit" value="Update">
                </form>
            </body>
            </html>
            `
            fmt.Fprint(w, editForm)
        }
    })

    // Handler for updating the name ("/update")
    http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            r.ParseForm()
            oldName := r.FormValue("oldName")
            newName := r.FormValue("newName")

            // Find and replace the old name with the new one
            names, err := redisClient.LRange(ctx, "names", 0, -1).Result()
            if err != nil {
                http.Error(w, "Failed to fetch names", http.StatusInternalServerError)
                return
            }

            for i, name := range names {
                if name == oldName {
                    // Update the name in Redis
                    err := redisClient.LSet(ctx, "names", int64(i), newName).Err()
                    if err != nil {
                        http.Error(w, "Failed to update name", http.StatusInternalServerError)
                        return
                    }
                    break
                }
            }

            // Redirect back to the home page to show the updated list
            http.Redirect(w, r, "/", http.StatusSeeOther)
        }
    })

    // Start the HTTP server
    http.ListenAndServe(":8080", nil)
}