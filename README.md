# Gofycat

# PLEASE NOT THAT A LOT OF ENDPOINTS **DO NOT** WORK. This will be resolved in the future. Please also note that a lot of functions just return a `*http.Response`, since Gfycat's api documentation is pretty lacking and I was unable to test a lot of it. Extended documentation will come later too.

# Usage

```golang
package main

import ("github.com/Krognol/gofycat"
        "fmt"
       )

func main() {
    client_id := "your client id"
    client_secret := "your client secret"

    gfy := gofycat.New(client_id, client_secret, gofycat.Client)
    user, err := gfy.GetUser("Whatever username here")

    if err != nil {
        panic(err)
    }

    fmt.Println(user.URL)
    // >> https://gfycat.com/@whatever_username_here
}
```

# Contributing

1. Fork it ( https://github.com/Krognol/gofycat/fork )
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create a new Pull Request