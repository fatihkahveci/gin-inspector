# Gin Inspector

![Gin Inspector HTML Preview](https://raw.githubusercontent.com/fatihkahveci/gin-inspector/master/preview-html.png)

Gin middleware for investigating http request.

## Usage


```sh
$ go get github.com/fatihkahveci/gin-inspector
```

### JSON Response

```
package main

import (
	"github.com/fatihkahveci/gin-inspector"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	debug := true

	if debug {
		r.Use(inspector.InspectorStats())
		r.GET("/_inspector", func(c *gin.Context) {
			c.JSON(200, inspector.GetPaginator())
		})
	}

	r.Run()
}
```

### Html Template

```
package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/fatihkahveci/gin-inspector"
	"github.com/gin-gonic/gin"
)

func formatDate(t time.Time) string {
	return t.Format(time.RFC822)
}

func main() {
	r := gin.Default()
	r.Delims("{{", "}}")

	r.SetFuncMap(template.FuncMap{
		"formatDate": formatDate,
	})

	r.LoadHTMLFiles("inspector.html")
	debug := true

	if debug {
		r.Use(inspector.InspectorStats())

		r.GET("/_inspector", func(c *gin.Context) {
			c.HTML(http.StatusOK, "inspector.html", map[string]interface{}{
				"title":      "Gin Inspector",
				"pagination": inspector.GetPaginator(),
			})

		})
	}

	r.Run(":8080")
}

```