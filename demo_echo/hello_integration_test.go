// go test -run Integration ./... -v -count=1

package demo_test

import (
	"context"
	"demo"
	"io"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationHelloSuccess(t *testing.T) {
	e := echo.New()
	e.Use(middleware.Recover())

	// Register the handler
	e.GET("/panic", demo.TryToFail)

	// Start the server
	errCh := make(chan error)
	go func() {
		errCh <- e.Start(":1323")
	}()

	err := waitForServerStart(e, errCh, false)
	assert.NoError(t, err)

	// Make a request to the server
	if resp, err := http.Get("http://127.0.0.1:1323/panic"); err == nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				assert.Fail(t, err.Error())
			}
		}(resp.Body)
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

		expect := `{"message":"Internal Server Error"}`
		if body, err := io.ReadAll(resp.Body); err == nil {
			assert.Equal(t, expect, strings.TrimSpace(string(body)))
		} else {
			assert.Fail(t, err.Error())
		}
	} else {
		assert.Fail(t, err.Error())
	}

	// Shutdown the server
	if err := e.Close(); err != nil {
		t.Fatal(err)
	}
}

func waitForServerStart(e *echo.Echo, errChan <-chan error, isTLS bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	ticker := time.NewTicker(5 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			var addr net.Addr
			if isTLS {
				addr = e.TLSListenerAddr()
			} else {
				addr = e.ListenerAddr()
			}
			if addr != nil && strings.Contains(addr.String(), ":") {
				return nil // was started
			}
		case err := <-errChan:
			if err == http.ErrServerClosed {
				return nil
			}
			return err
		}
	}
}
