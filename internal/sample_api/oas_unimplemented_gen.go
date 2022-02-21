// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"math/big"
	"math/bits"
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
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
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
	_ = bits.LeadingZeros64
	_ = big.Rat{}
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
	_ = otelogen.Version
	_ = attribute.KeyValue{}
	_ = trace.TraceIDFromHex
	_ = otel.GetTracerProvider
	_ = metric.NewNoopMeterProvider
	_ = regexp.MustCompile
	_ = jx.Null
	_ = sync.Pool{}
	_ = codes.Unset
)

var _ Handler = UnimplementedHandler{}

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

// DataGetFormat implements dataGetFormat operation.
//
// Retrieve data.
//
// GET /name/{id}/{foo}1234{bar}-{baz}!{kek}
func (UnimplementedHandler) DataGetFormat(ctx context.Context, params DataGetFormatParams) (r string, _ error) {
	return r, ht.ErrNotImplemented
}

// DefaultTest implements defaultTest operation.
//
// POST /defaultTest
func (UnimplementedHandler) DefaultTest(ctx context.Context, req DefaultTest, params DefaultTestParams) (r int32, _ error) {
	return r, ht.ErrNotImplemented
}

// ErrorGet implements errorGet operation.
//
// Returns error.
//
// GET /error
func (UnimplementedHandler) ErrorGet(ctx context.Context) (r ErrorStatusCode, _ error) {
	return r, ht.ErrNotImplemented
}

// FoobarGet implements foobarGet operation.
//
// Dumb endpoint for testing things.
//
// GET /foobar
func (UnimplementedHandler) FoobarGet(ctx context.Context, params FoobarGetParams) (r FoobarGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// FoobarPost implements foobarPost operation.
//
// Dumb endpoint for testing things.
//
// POST /foobar
func (UnimplementedHandler) FoobarPost(ctx context.Context, req OptPet) (r FoobarPostRes, _ error) {
	return r, ht.ErrNotImplemented
}

// FoobarPut implements  operation.
//
// PUT /foobar
func (UnimplementedHandler) FoobarPut(ctx context.Context) (r FoobarPutDefStatusCode, _ error) {
	return r, ht.ErrNotImplemented
}

// GetHeader implements getHeader operation.
//
// Test for header param.
//
// GET /test/header
func (UnimplementedHandler) GetHeader(ctx context.Context, params GetHeaderParams) (r Hash, _ error) {
	return r, ht.ErrNotImplemented
}

// MultipleGenericResponses implements multipleGenericResponses operation.
//
// GET /multipleGenericResponses
func (UnimplementedHandler) MultipleGenericResponses(ctx context.Context) (r MultipleGenericResponsesRes, _ error) {
	return r, ht.ErrNotImplemented
}

// NullableDefaultResponse implements nullableDefaultResponse operation.
//
// GET /nullableDefaultResponse
func (UnimplementedHandler) NullableDefaultResponse(ctx context.Context) (r NilIntStatusCode, _ error) {
	return r, ht.ErrNotImplemented
}

// OctetStreamBinaryStringSchema implements octetStreamBinaryStringSchema operation.
//
// GET /octetStreamBinaryStringSchema
func (UnimplementedHandler) OctetStreamBinaryStringSchema(ctx context.Context) (r OctetStreamBinaryStringSchemaOK, _ error) {
	return r, ht.ErrNotImplemented
}

// OctetStreamEmptySchema implements octetStreamEmptySchema operation.
//
// GET /octetStreamEmptySchema
func (UnimplementedHandler) OctetStreamEmptySchema(ctx context.Context) (r OctetStreamEmptySchemaOK, _ error) {
	return r, ht.ErrNotImplemented
}

// OneofBug implements oneofBug operation.
//
// POST /oneofBug
func (UnimplementedHandler) OneofBug(ctx context.Context, req OneOfBugs) (r OneofBugOK, _ error) {
	return r, ht.ErrNotImplemented
}

// PetCreate implements petCreate operation.
//
// Creates pet.
//
// POST /pet
func (UnimplementedHandler) PetCreate(ctx context.Context, req OptPet) (r Pet, _ error) {
	return r, ht.ErrNotImplemented
}

// PetFriendsNamesByID implements petFriendsNamesByID operation.
//
// Returns names of all friends of pet.
//
// GET /pet/friendNames/{id}
func (UnimplementedHandler) PetFriendsNamesByID(ctx context.Context, params PetFriendsNamesByIDParams) (r []string, _ error) {
	return r, ht.ErrNotImplemented
}

// PetGet implements petGet operation.
//
// Returns pet from the system that the user has access to.
//
// GET /pet
func (UnimplementedHandler) PetGet(ctx context.Context, params PetGetParams) (r PetGetRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PetGetAvatarByID implements petGetAvatarByID operation.
//
// Returns pet avatar by id.
//
// GET /pet/avatar
func (UnimplementedHandler) PetGetAvatarByID(ctx context.Context, params PetGetAvatarByIDParams) (r PetGetAvatarByIDRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PetGetAvatarByName implements petGetAvatarByName operation.
//
// Returns pet's avatar by name.
//
// GET /pet/{name}/avatar
func (UnimplementedHandler) PetGetAvatarByName(ctx context.Context, params PetGetAvatarByNameParams) (r PetGetAvatarByNameRes, _ error) {
	return r, ht.ErrNotImplemented
}

// PetGetByName implements petGetByName operation.
//
// Returns pet by name from the system that the user has access to.
//
// GET /pet/{name}
func (UnimplementedHandler) PetGetByName(ctx context.Context, params PetGetByNameParams) (r Pet, _ error) {
	return r, ht.ErrNotImplemented
}

// PetNameByID implements petNameByID operation.
//
// Returns pet name by pet id.
//
// GET /pet/name/{id}
func (UnimplementedHandler) PetNameByID(ctx context.Context, params PetNameByIDParams) (r string, _ error) {
	return r, ht.ErrNotImplemented
}

// PetUpdateNameAliasPost implements  operation.
//
// POST /pet/updateNameAlias
func (UnimplementedHandler) PetUpdateNameAliasPost(ctx context.Context, req OptPetName) (r PetUpdateNameAliasPostDefStatusCode, _ error) {
	return r, ht.ErrNotImplemented
}

// PetUpdateNamePost implements  operation.
//
// POST /pet/updateName
func (UnimplementedHandler) PetUpdateNamePost(ctx context.Context, req OptString) (r PetUpdateNamePostDefStatusCode, _ error) {
	return r, ht.ErrNotImplemented
}

// PetUploadAvatarByID implements petUploadAvatarByID operation.
//
// Uploads pet avatar by id.
//
// POST /pet/avatar
func (UnimplementedHandler) PetUploadAvatarByID(ctx context.Context, req PetUploadAvatarByIDReq, params PetUploadAvatarByIDParams) (r PetUploadAvatarByIDRes, _ error) {
	return r, ht.ErrNotImplemented
}

// RecursiveArrayGet implements  operation.
//
// GET /recursiveArray
func (UnimplementedHandler) RecursiveArrayGet(ctx context.Context) (r RecursiveArray, _ error) {
	return r, ht.ErrNotImplemented
}

// RecursiveMapGet implements  operation.
//
// GET /recursiveMap
func (UnimplementedHandler) RecursiveMapGet(ctx context.Context) (r RecursiveMap, _ error) {
	return r, ht.ErrNotImplemented
}

// TestFloatValidation implements testFloatValidation operation.
//
// POST /testFloatValidation
func (UnimplementedHandler) TestFloatValidation(ctx context.Context, req TestFloatValidation) (r TestFloatValidationOK, _ error) {
	return r, ht.ErrNotImplemented
}

// TestObjectQueryParameter implements testObjectQueryParameter operation.
//
// GET /testObjectQueryParameter
func (UnimplementedHandler) TestObjectQueryParameter(ctx context.Context, params TestObjectQueryParameterParams) (r TestObjectQueryParameterOK, _ error) {
	return r, ht.ErrNotImplemented
}