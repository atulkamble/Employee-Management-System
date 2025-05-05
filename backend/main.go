package main

import (
"fmt"
"log"
"net/http"
"os"
"github.com/gorilla/mux"
)

func main() {
r := mux.NewRouter()
r.HandleFunc("/employees", getEmployees).Methods("GET")
r.HandleFunc("/employees", createEmployee).Methods("POST")

```
port := "8080"
fmt.Printf("Server running on port %s\n", port)
log.Fatal(http.ListenAndServe(":"+port, r))
```

}
