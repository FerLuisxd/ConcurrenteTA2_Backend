package main

import (
    // "bufio"
    "fmt"
    "strings"
    "log"
    "os"
    "encoding/csv"
    "strconv"
    // "net"
    
	"encoding/json"
	"math/rand"
	"net/http"
    // "strconv"

    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
    "github.com/bitly/go-simplejson"
)
var adults []Adult
type Adult struct {
	Age     string  `json:"age"`
    Workclass string `json:"workclass"`
    Fnlwgt   int `json:"fnlwgt"`
    Education string    `json:"education"`
    EducationNum int    `json:"educationNum"`
    Marital string  `json:"marital"`
    Occupation string   `json:"occupation"`
    Relationship string `json:"relationship"`
    Race string `json:"race"`
    Sex int  `json:"sex"`
    CapitalGain int  `json:"capitalGain"`
    CapitalLoss int  `json:"capitalLoss"`
    Hours int    `json:"hours"`
    Native string   `json:"native"`
    Ke string `json:"Ke"`
}

func lineToStruc(lines [][]string){
        // Loop through lines & turn into object
        for _, line := range lines {
            Fnlwgt,_ := strconv.Atoi(strings.TrimSpace(line[2]))
            EducationNum,_ := strconv.Atoi(strings.TrimSpace(line[4]))
            // fmt.Println(a)
            CapitalGain,_ := strconv.Atoi(strings.TrimSpace(line[10]))
            CapitalLoss,_ := strconv.Atoi(strings.TrimSpace(line[11]))
            Hours,_ := strconv.Atoi(strings.TrimSpace(line[12]))
            Sex:= 0
            if(strings.TrimSpace(line[9]) == "Male"){
                Sex = 10
            }

            adults = append(adults,Adult{
                Age: strings.TrimSpace(line[0]),
                Workclass: strings.TrimSpace(line[1]),
                Fnlwgt: Fnlwgt,
                Education: strings.TrimSpace(line[3]),
                EducationNum: EducationNum,
                Marital: strings.TrimSpace(line[5]),
                Occupation: strings.TrimSpace(line[6]),
                Relationship: strings.TrimSpace(line[7]),
                Race: strings.TrimSpace(line[8]),
                Sex: Sex,
                CapitalGain: CapitalGain,
                CapitalLoss: CapitalLoss,
                Hours: Hours,
                Native: strings.TrimSpace(line[13]),
                Ke: strings.TrimSpace(line[14]),
            })
        }
}

func readFile(filePath string) ([][]string, error) {

 // Open CSV file
 f, err := os.Open(filePath)
 if err != nil {
     return [][]string{}, err
 }
 defer f.Close()

 // Read File into a Variable
 lines, err := csv.NewReader(f).ReadAll()
 if err != nil {
     return [][]string{}, err
 }

 return lines, nil
}

// Adult struct (Model)



// Get all adults
func getAdults(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(adults)
}

// Get single adult
func getAdult(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through adults and find one with the id from the params
	for _, item := range adults {
        fnlwgt,_ := strconv.Atoi(params["id"])
		if item.Fnlwgt == fnlwgt {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Adult{})
}

// Get single adult
func getCategory(w http.ResponseWriter, r *http.Request) {
    // w.Header().Set("Content-Type", "text/html; charset=utf-8")
    // w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    // w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
 

	w.Header().Set("Content-Type", "application/json")
    // params := mux.Vars(r) // Gets params


    var adult Adult
    _ = json.NewDecoder(r.Body).Decode(&adult)
    
    k := 20 +rand.Intn(20)
    fmt.Println(k)
    result := testCase(adults,adult,k)
    fmt.Printf("Predicted: %s, Actual: %s\n", result[0].key, adult.Ke)
    
	json := simplejson.New()
	json.Set("knn", result[0].key)
	json.Set("actual", adult.Ke)
    json.Set("predicted", result[0].key == adult.Ke)
    
    adults = append(adults, adult)
    
    payload, _ := json.MarshalJSON()
    w.Write(payload)
}


// Add new adult
func createAdult(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var adult Adult
	_ = json.NewDecoder(r.Body).Decode(&adult)
	adults = append(adults, adult)
	json.NewEncoder(w).Encode(adult)
}


func main() {
    lines, err := readFile("dataset/adult.data")
    if err != nil {
        panic(err)
    }
    fmt.Println("Leyo archivos")
    lineToStruc(lines)
    fmt.Println("Parseo Archivos")

	r := mux.NewRouter()

	r.HandleFunc("/adults", getAdults).Methods("GET")
	r.HandleFunc("/adults/{id}", getAdult).Methods("GET")
	r.HandleFunc("/adults", createAdult).Methods("POST")
    r.HandleFunc("/knn", getCategory).Methods("POST")
    
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	//router.HandleFunc("/", RootEndpointGET).Methods("GET")
	//router.HandleFunc("/", RootEndpointPOST).Methods("POST")

    // Start server
    port := ":8000"
    fmt.Println("Escuchando en " + port )
    main3()
    log.Fatal(http.ListenAndServe(port, handlers.CORS(headers, methods, origins)(r)))

}
