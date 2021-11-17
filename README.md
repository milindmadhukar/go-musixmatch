
# Go-Musixmatch!

This is a WIP Go wrapper for working with the [Musixmatch](https://www.musixmatch.com/) API.

It aims to support every task listed in the Web API Endpoint Reference, located [here](https://developer.musixmatch.com/documentation).

*Most of the API endpoints are covered. Docs in progress.*


## 💻 Installation
To install the library simply open a terminal and type:
```
go get github.com/milindmadhukar/go-musixmatch
```

## ️️🛠️ Tools Used

This project was written purely in `Golang` for `Golang`.</br>
The module helps with the usage of the [Musixmatch](https://developer.musixmatch.com/documentation) API.

## ⛏️  Acquiring Musixmatch API Key.

1. Go to the [Musixmatch Developer Page](https://developer.musixmatch.com/plans) and select a plan.
1. Fill in the necessary details to create an account and verify it from your email.
1. Go to your [Account Dashboard](https://developer.musixmatch.com/admin/applications) and you should see an API Key.
1. Note the API Key down and don't reveal it to anyone.


## 🏁 Basic Setup:
For all the endpoints and parameters that can be used, check the [Musixmatch API docs](https://developer.musixmatch.com/documentation)

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	mxm "github.com/milindmadhukar/go-musixmatch"
	"github.com/milindmadhukar/go-musixmatch/params"
)

func main() {

	client := mxm.New("<YOUR API KEY>", http.DefaultClient)

	artists, err := client.SearchArtist(context.Background(), params.QueryArtist("Martin Garrix"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(artists[0])
}
```

### Output
```go
{{24407895 Martin Garrix []  NL [{MARTIJN GARRITSEN}] 71  {[]} 0 2017-02-03 07:02:12 +0000 UTC 1996 1996-05-15  0000-00-00}}
```


## 🧿 Extras

If you face any difficulties contact me [here.](https://milindm.me/contact/)

Thats it, have fun ✚✖
