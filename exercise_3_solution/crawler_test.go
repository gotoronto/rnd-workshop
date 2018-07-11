package crawler

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestCrawl tests the Crawl(url string) method
func TestCrawl(t *testing.T) {
	var requestCount int32 // A counter to record how many requests our test server received

	// testCrawlerServer is a server that will respond with the same html to every
	// request. The body of the html will be formatted with 5 links that are urls
	// back to itself.
	var testCrawlerServer *httptest.Server
	testCrawlerServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second) // Each request will sleep 1 second before responding
		// bodyTmpl is the html template to be written out on the response when a request
		// is made.
		bodyTmpl, _ := template.New("name").Parse(`<html>
				<body>
					<a href="{{.URL}}/one"></a>
					<a href="{{.URL}}/two"></a>
					<a href="{{.URL}}/three"></a>
					<a href="{{.URL}}/four"></a>
				</body>
		</html>
		`)
		bodyTmpl.Execute(w, struct{ URL string }{testCrawlerServer.URL}) // Write out the html to the response
		atomic.AddInt32(&requestCount, 1)                                // Record that we received a request
	}))
	defer testCrawlerServer.Close() // Shut down the test server after the test completes

	startTime := time.Now()              // The current time before Crawl() is started
	Crawl(testCrawlerServer.URL)         // Calling the Crawl() method that you implement
	timeElapsed := time.Since(startTime) // The time elapsed since Crawl() was called

	// We assert that the test should pass in under 3 seconds if you have implemented
	// concurrency appropriately even though every response should take about 1 second.
	assert.True(t, timeElapsed < 3*time.Second, "Should succeed under 3 seconds")
	// We assert that each link should be request exactly once and no url is requested
	// twice
	assert.Equal(t, int32(5), requestCount, "Should request every url once")
}
