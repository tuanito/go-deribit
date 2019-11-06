package deribit

import (
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/go-openapi/strfmt"

	"github.com/adampointer/go-deribit/v3/client/operations"
	"github.com/gorilla/websocket"
)

const (
	liveURL = "wss://www.deribit.com/ws/api/v2/"
	testURL = "wss://test.deribit.com/ws/api/v2/"
)

// ErrTimeout - request timed out
var ErrTimeout = errors.New("timed out waiting for a response")

// Exchange is an API wrapper with the exchange
type Exchange struct {
	url    string
	test   bool
	client *operations.Client
	RPCCore
}

// NewExchange creates a new API wrapper
// key and secret can be ignored if you are only calling public endpoints
func NewExchange(test bool, errs chan error, stop chan bool) (*Exchange, error) {
	exc := &Exchange{
		RPCCore: RPCCore{
			calls: &callManager{
				pending:       make(map[uint64]*RPCCall, 1),
				subscriptions: make(map[string]*RPCSubscription),
				counter:       1,
			},
			connMgr: &connManager{},
			errors:  errs,
			stop:    stop,
		},
	}
	exc.onDisconnect = exc.Reconnect
	exc.url = liveURL
	if test {
		exc.test = true
		exc.url = testURL
	}
	return exc, nil
}

// NewExchangeFromCore creates a new exchange from an existing RPCCore
func NewExchangeFromCore(test bool, core RPCCore) *Exchange {
	exc := &Exchange{RPCCore: core}
	exc.onDisconnect = exc.Reconnect
	exc.url = liveURL
	if test {
		exc.test = true
		exc.url = testURL
	}
	return exc
}

// Connect to the websocket API
func (e *Exchange) Connect() error {
	c, _, err := websocket.DefaultDialer.Dial(e.url, nil)
	if err != nil {
		return err
	}
	e.connMgr.conn = c
	// Start listening for responses
	go e.read()
	go e.heartbeat()

	authed := false
	if clientID != "" && clientSecret != "" {
		// re-authenticate on reconnect
		if err := e.Authenticate(); err != nil {
			return fmt.Errorf("Error re-authenticating: %s", err)
		}
		authed = true
	}
	e.resubscribe(authed)
	return nil
}

// Close the websocket connection
func (e *Exchange) Close() error {
	return e.close()
}

// SetLogOutput set log output
func (e *Exchange) SetLogOutput(w io.Writer) {
	log.SetOutput(w)
}

// SetDisconnectHandler overrides the default disconnect handler
func (e *Exchange) SetDisconnectHandler(f func(*RPCCore)) {
	e.onDisconnect = f
}

// Reconnect reconnect is already built-in on OnDisconnect. Use this method only within OnDisconnect to override it
func (e *Exchange) Reconnect(core *RPCCore) {
	if err := e.Connect(); err != nil {
		log.Printf("reconnect failed %v", err)
		e.errors <- fmt.Errorf("reconnect failed: %w", err)
	}
}

// Client returns an initialised API client
func (e *Exchange) Client() *operations.Client {
	if e.client == nil {
		e.client = operations.New(e, strfmt.Default)
	}
	return e.client
}

func (e *Exchange) heartbeat() {
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				e.testConnection()
			case <-e.stop:
				ticker.Stop()
			}
		}
	}()
}

func (e *Exchange) testConnection() {
	if _, err := e.Client().GetPublicTest(&operations.GetPublicTestParams{}); err != nil {
		// We've got an error, so we reconnect
		e.onDisconnect(&e.RPCCore)
	}
}

func (e *Exchange) resubscribe(authed bool) {
	if !authed {
		// We re-wire public subscriptions
		for chan0 := range e.calls.getSubscriptions() {
			log.Printf("Attempt at reconnecting subscription: %v", chan0)
			req := &operations.GetPublicSubscribeParams{Channels: []string{chan0}}

			if _, err := e.Client().GetPublicSubscribe(req); err != nil {
				log.Printf("Reconnection failed: %v", err)
				e.calls.deleteSubscription(chan0)
				continue
			}
			log.Printf("Subscription %v successfully re-wired", chan0)
		}
		return
	}

	// We re-wire private subscriptions
	for chan0 := range e.calls.getSubscriptions() {
		log.Printf("Attempt at reconnecting subscription: %v", chan0)
		req := &operations.GetPrivateSubscribeParams{Channels: []string{chan0}}

		if _, err := e.Client().GetPrivateSubscribe(req); err != nil {
			log.Printf("Reconnection failed: %v", err)
			e.calls.deleteSubscription(chan0)
			continue
		}
		log.Printf("Subscription %v successfully re-wired", chan0)
	}
}
