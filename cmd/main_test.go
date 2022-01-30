package main

import (
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	n   int
	exp *big.Int
}

func TestWithoutN(t *testing.T) {
	r := setupServer("../templates/*.html")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Send required fibonacci sequence element number as <i>?n=[your number]</i>")
}

func TestWithInvalidN(t *testing.T) {
	r := setupServer("../templates/*.html")

	testCases := []testCase{
		{
			n: -10,
		},
		{
			n: 2000001,
		},
	}

	for _, tc := range testCases {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/?n=%d", tc.n), nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Contains(t, w.Body.String(), "Send required fibonacci sequence element number as <i>?n=[your number]</i>")
	}
}

func TestWithValidN(t *testing.T) {
	r := setupServer("../templates/*.html")

	// This calculator https://keisan.casio.com/exec/system/1180573404 is used as a criterion of correct results.
	// The results of different calculators might differ depending on what is considered to be the first number: 0 or 1.
	fib400 := new(big.Int)
	fib400.SetString("176023680645013966468226945392411250770384383304492191886725992896575345044216019675", 10)
	fib1000 := new(big.Int)
	fib1000.SetString("43466557686937456435688527675040625802564660517371780402481729089536555417949051890403879840079255169295922593080322634775209689623239873322471161642996440906533187938298969649928516003704476137795166849228875", 10)

	testCases := []testCase{
		{
			n:   0,
			exp: big.NewInt(0),
		},
		{
			n:   1,
			exp: big.NewInt(1),
		},
		{
			n:   42,
			exp: big.NewInt(267914296),
		},
		{
			n:   400,
			exp: fib400,
		},
		{
			n:   1000,
			exp: fib1000,
		},
	}

	for _, tc := range testCases {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", fmt.Sprintf("/?n=%d", tc.n), nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Contains(t, w.Body.String(), fmt.Sprintf("<b>Requested number:</b> %d", tc.n))
		assert.Contains(t, w.Body.String(), fmt.Sprintf("<b>Fibonacci number:</b> %d", tc.exp))
	}
}
