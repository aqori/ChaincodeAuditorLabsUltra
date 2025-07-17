// cmd/chaincodeauditorlabsultra/main.go
package main

import (
"flag"
"log"
"os"

"chaincodeauditorlabsultra/internal/chaincodeauditorlabsultra"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := chaincodeauditorlabsultra.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
