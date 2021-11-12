// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
)

func (s *Server) notFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p := []byte(r.URL.Path)
	if len(p) == 0 {
		s.notFound(w, r)
		return
	}

	var (
		idx  int               // index of next slash
		elem []byte            // current element, without slashes
		args map[string]string // lazily initialized
	)

	// Static code generated router with unwrapped path search.
	switch r.Method {
	case "POST":
		// Root edge.
		if len(p) > 1 && p[0] == '/' {
			p = p[1:]
		}
		if idx = bytes.IndexByte(p[:], '/'); idx < 0 { // looking for next element
			elem, p = p, p[:0] // slash not found, using full path
		} else {
			elem = p[:idx] // slash found, element is path until slash
			p = p[idx:]
		}
		switch string(elem) {
		case "answerCallbackQuery": // -> 1
			// POST /answerCallbackQuery
			s.handleAnswerCallbackQueryPostRequest(args, w, r)
			return
		case "answerPreCheckoutQuery": // -> 2
			// POST /answerPreCheckoutQuery
			s.handleAnswerPreCheckoutQueryPostRequest(args, w, r)
			return
		case "answerShippingQuery": // -> 3
			// POST /answerShippingQuery
			s.handleAnswerShippingQueryPostRequest(args, w, r)
			return
		case "close": // -> 4
			// POST /close
			s.handleClosePostRequest(args, w, r)
			return
		case "deleteStickerFromSet": // -> 5
			// POST /deleteStickerFromSet
			s.handleDeleteStickerFromSetPostRequest(args, w, r)
			return
		case "deleteWebhook": // -> 6
			// POST /deleteWebhook
			s.handleDeleteWebhookPostRequest(args, w, r)
			return
		case "getFile": // -> 7
			// POST /getFile
			s.handleGetFilePostRequest(args, w, r)
			return
		case "getGameHighScores": // -> 8
			// POST /getGameHighScores
			s.handleGetGameHighScoresPostRequest(args, w, r)
			return
		case "getMe": // -> 9
			// POST /getMe
			s.handleGetMePostRequest(args, w, r)
			return
		case "getMyCommands": // -> 10
			// POST /getMyCommands
			s.handleGetMyCommandsPostRequest(args, w, r)
			return
		case "getStickerSet": // -> 11
			// POST /getStickerSet
			s.handleGetStickerSetPostRequest(args, w, r)
			return
		case "getUpdates": // -> 12
			// POST /getUpdates
			s.handleGetUpdatesPostRequest(args, w, r)
			return
		case "getUserProfilePhotos": // -> 13
			// POST /getUserProfilePhotos
			s.handleGetUserProfilePhotosPostRequest(args, w, r)
			return
		case "getWebhookInfo": // -> 14
			// POST /getWebhookInfo
			s.handleGetWebhookInfoPostRequest(args, w, r)
			return
		case "logOut": // -> 15
			// POST /logOut
			s.handleLogOutPostRequest(args, w, r)
			return
		case "sendGame": // -> 16
			// POST /sendGame
			s.handleSendGamePostRequest(args, w, r)
			return
		case "sendInvoice": // -> 17
			// POST /sendInvoice
			s.handleSendInvoicePostRequest(args, w, r)
			return
		case "setMyCommands": // -> 18
			// POST /setMyCommands
			s.handleSetMyCommandsPostRequest(args, w, r)
			return
		case "setStickerPositionInSet": // -> 19
			// POST /setStickerPositionInSet
			s.handleSetStickerPositionInSetPostRequest(args, w, r)
			return
		default:
			s.notFound(w, r)
			return
		}
	default:
		s.notFound(w, r)
		return
	}
}