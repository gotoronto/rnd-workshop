package crawler

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var bodyTmpl *template.Template

func init() {
	var err error
	bodyTmpl, err = template.New("name").Parse(`<html>
			<body>
				<a href="{{.URL}}/one"></a>
				<a href="{{.URL}}/two"></a>
				<a href="{{.URL}}/three"></a>
				<a href="{{.URL}}/four"></a>
			</body>
	</html>
	`)
	if err != nil {
		panic(err)
	}
}

func TestCrawl(t *testing.T) {
	start := time.Now()
	count := 0
	done := make(chan int, 1)
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		count++
		bodyTmpl.Execute(w, struct{ URL string }{server.URL})
		if count == 5 {
			done <- 1
		}
	}))
	defer server.Close()
	crawler := New(server.URL)
	crawler.Crawl(done)

	elapsed := time.Since(start)
	assert.True(t, elapsed < 1*time.Second)
}
