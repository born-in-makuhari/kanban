//
// Specs of controller
//

package main_test

import (
	"fmt"
	. "github.com/born-in-makuhari/kanban"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/zenazn/goji/web"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

// --------------------------------------------------------------------
// My helper functions

func ParseResponse(res *http.Response) (string, int) {
	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(contents), res.StatusCode
}

func Requester(url string) *http.Response {
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		Fail("http.Response returns error")
	}
	return res
}

// --------------------------------------------------------------------
// Describes

var _ = Describe("KanbanC", func() {
	var app *web.Mux
	var ts *httptest.Server
	var res *http.Response

	// create Server
	BeforeEach(func() {
		app = web.New()
		Route(app)
		ts = httptest.NewServer(app)
	})

	// close Server
	AfterEach(func() {
		ts.Close()
	})

	//
	// GET /
	//
	Describe("/", func() {
		Context("when get", func() {

			It("should return 200", func() {
				res = Requester(ts.URL + "/")
				Expect(res.StatusCode).To(Equal(http.StatusOK))
			})
		})
	})
	//
	// GET /hello/:name
	//
	Describe("/hello/:name", func() {
		Context("when get with name", func() {
			It("should return 200", func() {
				res = Requester(ts.URL + "/hello/kanban")
				Expect(res.StatusCode).To(Equal(http.StatusOK))
			})
		})
		Context("when get without name", func() {
			It("should return 404", func() {
				res = Requester(ts.URL + "/hello/")
				Expect(res.StatusCode).To(Equal(http.StatusNotFound))
			})
		})
	})
})
