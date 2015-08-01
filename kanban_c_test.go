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

func ParseResponse(res *http.Response) (string, int) {
	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(contents), res.StatusCode
}

var _ = Describe("KanbanC", func() {
	Describe("'/'", func() {
		Context("when get", func() {
			BeforeEach(func() {
			})

			It("should return 200", func() {
				app := web.New()
				Route(app)
				ts := httptest.NewServer(app)
				defer ts.Close()

				res, err := http.Get(ts.URL + "/")
				fmt.Printf("test %i", http.StatusOK)
				if err != nil {
					Fail("http.Response returns error")
				}
				Expect(res.StatusCode).To(Equal(http.StatusOK))
			})
		})
	})
})
