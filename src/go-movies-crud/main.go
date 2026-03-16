/*
GET: Retrieves data from the server. 
It is safe and idempotent (making the same request multiple times yields the same result).

POST: Sends data to the server to create a new resource.
 Repeated calls can create duplicate resources.

PUT: Updates or replaces an existing resource with the provided data.
This is an idempotent method.

DELETE: Removes the specified resource from the server.

PATCH: Applies partial modifications to a resource,
 rather than replacing the entire resource as PUT does.

HEAD: Retrieves only the header information of a resource,
 without the actual body.

OPTIONS: Describes the communication options (allowed methods)
 for a target resource. 

*/

package main



import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

// we  will  create a struct called movie and 
// we will  create a struct  called  director 

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director * Director `json:"director"`

}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`

}

func getMovies(w http.ResponseWriter , r * http.Request ) {
	// 1. Set the Content-Type so the client knows it's receiving JSON
    w.Header().Set("Content-Type", "application/json")

    // 2. Use the JSON encoder to send the 'movies' slice to the ResponseWriter
    // 'movies' is the slice we created earlier
    err := json.NewEncoder(w).Encode(movies)
    
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func deleteMovie(w http.ResponseWriter , r * http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[0:index], movies[index+1:]... )
			break
		}
	}
}

func getMovie(w http.ResponseWriter , r * http.Request) {
	w.Header().Set("Content-Type" , "application/json")
	
	params := mux.Vars(r)

	for _, item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
		}
	}
} 

// What real systems do
// 
// 1️⃣ Use UUID (most common)
// 
// import "github.com/google/uuid"
// 
// movie.ID = uuid.New().String()
// 
// Example ID:
// 
// 550e8400-e29b-41d4-a716-446655440000
// 
// Probability of collision ≈ practically zero.
func createMovie(w http.ResponseWriter , r * http.Request){
	w.Header().Set("Content-Type" , "application/json")
	var movie Movie
	
	err := json.NewDecoder(r.Body).Decode(&movie)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	movie.ID = strconv.Itoa(rand.Intn(1000000000)) 
	movies = append(movies , movie)

	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter , r * http.Request){
	deleteMovie(w,r)

	createMovie(w,r)
}


var movies [] Movie

var nolan * Director = &Director{Firstname: "Christopher", Lastname: "Nolan"}

func main(){
	//using the gorilla mux router 
	router := mux.NewRouter()
	movies = append(movies, 
        Movie{ID: "1", Isbn: "438227", Title: "Inception", Director: nolan},
        Movie{ID: "2", Isbn: "454551", Title: "Interstellar", Director: nolan},
        Movie{ID: "3", Isbn: "882731", Title: "The Dark Knight", Director: nolan},
        Movie{ID: "4", Isbn: "112233", Title: "Pulp Fiction", Director: &Director{Firstname: "Quentin", Lastname: "Tarantino"}},
        Movie{ID: "5", Isbn: "445566", Title: "The Grand Budapest Hotel", Director: &Director{Firstname: "Wes", Lastname: "Anderson"}},
        Movie{ID: "6", Isbn: "778899", Title: "Parasite", Director: &Director{Firstname: "Bong", Lastname: "Joon-ho"}},
        Movie{ID: "7", Isbn: "101112", Title: "The Wolf of Wall Street", Director: &Director{Firstname: "Martin", Lastname: "Scorsese"}},
        Movie{ID: "8", Isbn: "131415", Title: "Arrival", Director: &Director{Firstname: "Denis", Lastname: "Villeneuve"}},
        Movie{ID: "9", Isbn: "161718", Title: "Lady Bird", Director: &Director{Firstname: "Greta", Lastname: "Gerwig"}},
        Movie{ID: "10", Isbn: "192021", Title: "Seven Samurai", Director: &Director{Firstname: "Akira", Lastname: "Kurosawa"}},
    )
	router.HandleFunc("/movies",getMovies).Methods("GET") 
	router.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies",createMovie).Methods("POST") 
	router.HandleFunc("/movies/{id}",updateMovie).Methods("POST") 

	fmt.Println("starting the server @ port:8000")
	log.Fatal(http.ListenAndServe(":8000",router))
	

	
}