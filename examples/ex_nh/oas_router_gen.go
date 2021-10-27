// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
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
)

func Register(r chi.Router, s Server) {
	r.MethodFunc("GET", "/api/gallery/{book_id}", NewGetBookHandler(s))
	r.MethodFunc("GET", "/galleries/{media_id}/cover.{format}", NewGetPageCoverImageHandler(s))
	r.MethodFunc("GET", "/galleries/{media_id}/{page}.{format}", NewGetPageImageHandler(s))
	r.MethodFunc("GET", "/galleries/{media_id}/{page}t.{format}", NewGetPageThumbnailImageHandler(s))
	r.MethodFunc("GET", "/api/galleries/search", NewSearchHandler(s))
	r.MethodFunc("GET", "/api/galleries/tagged", NewSearchByTagIDHandler(s))
}