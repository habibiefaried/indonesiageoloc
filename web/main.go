package web

import (
	"database/sql"
    "net/http"
    "encoding/json"
    "strconv"
    "fmt"
    "strings"
    _ "github.com/go-sql-driver/mysql"
)

var DBString string

type Address struct {
        ID      string  `json:"id"`
        Full    []string `json:"full"`
}

type ResultResponse struct {
        Total   int      `json:"total"`
        Address []Address `json:"address"`
        Error   string   `json:"error"`
}

func dbConn() (db *sql.DB) {
    db, err := sql.Open("mysql", DBString)
    if err != nil {
        panic(err.Error())
    }
    return db
}

func topAreaQuery(lat string, lon string, limit int, tipe string) []Address{
        // variable tipe validation
        ret := make([]Address, limit)
        if (strings.Compare(tipe,"village") == 0) || (strings.Compare(tipe,"city") == 0) || (strings.Compare(tipe,"area") == 0) {
                db := dbConn()
                var sf float32 = 3.14159 / 180
                selDB, err := db.Query("SELECT * FROM "+tipe+" ORDER BY ACOS(SIN(lat * ?) *SIN(? * ?) + COS(lat * ?) * COS(? * ?) * COS( (lon-?) * ? )) ASC LIMIT ?", sf, lat, sf, sf, lat, sf, lon, sf, limit)
                if err != nil {
                fmt.Println(err)
            }

            count := 0
            for selDB.Next() {
                var id, name, areaid string
                var latq, longq float32
                err = selDB.Scan(&id, &areaid, &name, &latq, &longq)
                ret[count].ID = id
                ret[count].Full = strings.SplitN(name, "/", -1)
                count = count + 1
                if err != nil {
                        fmt.Println(err)
                }
            }
            db.Close()
        }
        return ret
}

func index(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
    lattitude, ok := r.URL.Query()["lat"]
    if !ok || len(lattitude[0]) < 1 {
        rr := ResultResponse {
                Total: 0,
                Error: "Url Param 'lat' is missing",
        }
        json.NewEncoder(w).Encode(rr)
        return
    }

    longitude, ok := r.URL.Query()["long"]
    if !ok || len(longitude[0]) < 1 {
        rr := ResultResponse {
                Total: 0,
                Error: "Url Param 'long' is missing",
        }
        json.NewEncoder(w).Encode(rr)
        return
    }

    limit, ok := r.URL.Query()["limit"]
    if !ok || len(limit[0]) < 1 {
        rr := ResultResponse {
                Total: 0,
                Error: "Url Param 'limit' is missing",
        }
        json.NewEncoder(w).Encode(rr)
        return
    }

    i, err := strconv.Atoi(limit[0]) // conversion from limit string to int
    if err != nil {
        rr := ResultResponse {
                Total: 0,
                Error: err.Error(),
        }
        json.NewEncoder(w).Encode(rr)
        return
    }

    if ((i < 1) || (i > 50)) {
        rr := ResultResponse {
                Total: 0,
                Error: "Limit is between 1-50",
        }
        json.NewEncoder(w).Encode(rr)
        return
    }

    tipe, ok := r.URL.Query()["tipe"]
    if !ok || len(tipe[0]) < 1 {
        rr := ResultResponse {
                Total: 0,
                Error: "Url Param 'tipe' is missing",
        }
        json.NewEncoder(w).Encode(rr)
        return
    }

    retVillage := topAreaQuery(lattitude[0],longitude[0], i, tipe[0])
    rr := ResultResponse {
                Total: len(retVillage),
                Address: retVillage,
        }
    json.NewEncoder(w).Encode(rr)
}

func Serve(p string) {
	DBString = p
	fmt.Println("Listening on port 18081")
    http.HandleFunc("/", index)
    http.ListenAndServe(":18081", nil)
}
