/******************************************************************************
 * Copyright (c) Travis Gruber (2017) All rights reserved.
 *
 * file: main.go 
 *
 * Entry point for OrderService.  This service handles all order submission
 * requests from various users.  Order submissions are the only driver for 
 * evaluation and creation of reports, which are based on the contents of 
 * consumer accounts.  The public API also expects authentication headers with
 * valid (cryptographic) tokens, obtained from other services prior to Order 
 * submission.  Orders can also be queried (i.e. http GET), either by orderId,
 * or by lenderId (wih some filtering based on order type).  Order queries 
 * reflect the state of the order, which reports have been fulfilled, and where
 * to retrieve them.
 *****************************************************************************/
package main

import (
    "fmt"
    "os"
    "net/http"

    "github.com/urfav/negroni"
    "github.com/gorilla/mux"
    "github.com/unrolled/render"

    "github.com/grubertw/gimmeyourmoney/OrderService/http_handlers"
    "github.com/grubertw/gimmeyourmoney/OrderService/mq_handlers"
)

func main() {
    port := os.Getenv("PORT")
    if len(port) == 0 {
        port = "3000"
    }

    formatter := render.New(render.Options{IndentJSON: true,})

    server := negroni.Classic()
    mx := mux.NewRouter()

    server.UseHandler(mx)

    server.Run(":" + port)
}