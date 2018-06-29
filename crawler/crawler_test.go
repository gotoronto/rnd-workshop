package crawler

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"os"
	"sync/atomic"
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
	var count int32
	done := make(chan os.Signal, 1)

	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		bodyTmpl.Execute(w, struct{ URL string }{server.URL})
		if atomic.AddInt32(&count, 1) == 5 {
			done <- os.Interrupt
		}
	}))
	defer server.Close()

	crawler := New(server.URL)
	crawler.Crawl(done)

	elapsed := time.Since(start)
	assert.True(t, elapsed < 3*time.Second, "Should succeed under 3 seconds")
}
