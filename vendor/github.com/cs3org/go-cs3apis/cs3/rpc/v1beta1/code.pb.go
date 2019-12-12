// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/rpc/v1beta1/code.proto

package rpcv1beta1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// These are the canonical error codes used by CS3 APIs.
//
// Adapted from Google APIs:
// https://github.com/googleapis/googleapis/
//
// Sometimes multiple error codes may apply.  Services should return
// the most specific error code that applies.  For example, prefer
// `OUT_OF_RANGE` over `FAILED_PRECONDITION` if both codes apply.
// Similarly prefer `NOT_FOUND` or `ALREADY_EXISTS` over `FAILED_PRECONDITION`.
type Code int32

const (
	// A programmer would not intentionally set the code to CODE_INVALID.
	// This code exists to force service implementors to set
	// a specific code for the API call and to not rely on defaults.
	//
	// HTTP Mapping: 500 Internal Server Error
	Code_CODE_INVALID Code = 0
	// Not an error; returned on success
	//
	// HTTP Mapping: 200 OK
	Code_CODE_OK Code = 1
	// The operation was cancelled, typically by the caller.
	//
	// HTTP Mapping: 499 Client Closed Request
	Code_CODE_CANCELLED Code = 2
	// Unknown error.  For example, this error may be returned when
	// a `Status` value received from another address space belongs to
	// an error space that is not known in this address space.  Also
	// errors raised by APIs that do not return enough error information
	// may be converted to this error.
	//
	// HTTP Mapping: 500 Internal Server Error
	Code_CODE_UNKNOWN Code = 3
	// The client specified an invalid argument.  Note that this differs
	// from `FAILED_PRECONDITION`.  `INVALID_ARGUMENT` indicates arguments
	// that are problematic regardless of the state of the system
	// (e.g., a malformed file name).
	//
	// HTTP Mapping: 400 Bad Request
	Code_CODE_INVALID_ARGUMENT Code = 4
	// The deadline expired before the operation could complete. For operations
	// that change the state of the system, this error may be returned
	// even if the operation has completed successfully.  For example, a
	// successful response from a server could have been delayed long
	// enough for the deadline to expire.
	//
	// HTTP Mapping: 504 Gateway Timeout
	Code_CODE_DEADLINE_EXCEEDED Code = 5
	// Some requested entity (e.g., file or directory) was not found.
	//
	// Note to server developers: if a request is denied for an entire class
	// of users, such as gradual feature rollout or undocumented whitelist,
	// `NOT_FOUND` may be used. If a request is denied for some users within
	// a class of users, such as user-based access control, `PERMISSION_DENIED`
	// must be used.
	//
	// HTTP Mapping: 404 Not Found
	Code_CODE_NOT_FOUND Code = 6
	// The entity that a client attempted to create (e.g., file or directory)
	// already exists.
	//
	// HTTP Mapping: 409 Conflict
	Code_CODE_ALREADY_EXISTS Code = 7
	// The caller does not have permission to execute the specified
	// operation. `PERMISSION_DENIED` must not be used for rejections
	// caused by exhausting some resource (use `RESOURCE_EXHAUSTED`
	// instead for those errors). `PERMISSION_DENIED` must not be
	// used if the caller can not be identified (use `UNAUTHENTICATED`
	// instead for those errors). This error code does not imply the
	// request is valid or the requested entity exists or satisfies
	// other pre-conditions.
	//
	// HTTP Mapping: 403 Forbidden
	Code_CODE_PERMISSION_DENIED Code = 8
	// The request does not have valid authentication credentials for the
	// operation.
	//
	// HTTP Mapping: 401 Unauthorized
	Code_CODE_UNAUTHENTICATED Code = 9
	// Some resource has been exhausted, perhaps a per-user quota, or
	// perhaps the entire file system is out of space.
	//
	// HTTP Mapping: 429 Too Many Requests
	Code_CODE_RESOURCE_EXHAUSTED Code = 10
	// The operation was rejected because the system is not in a state
	// required for the operation's execution.  For example, the directory
	// to be deleted is non-empty, an rmdir operation is applied to
	// a non-directory, etc.
	//
	// Service implementors can use the following guidelines to decide
	// between `FAILED_PRECONDITION`, `ABORTED`, and `UNAVAILABLE`:
	//  (a) Use `UNAVAILABLE` if the client can retry just the failing call.
	//  (b) Use `ABORTED` if the client should retry at a higher level
	//      (e.g., when a client-specified test-and-set fails, indicating the
	//      client should restart a read-modify-write sequence).
	//  (c) Use `FAILED_PRECONDITION` if the client should not retry until
	//      the system state has been explicitly fixed.  E.g., if an "rmdir"
	//      fails because the directory is non-empty, `FAILED_PRECONDITION`
	//      should be returned since the client should not retry unless
	//      the files are deleted from the directory.
	//
	// HTTP Mapping: 400 Bad Request
	Code_CODE_FAILED_PRECONDITION Code = 11
	// The operation was aborted, typically due to a concurrency issue such as
	// a sequencer check failure or transaction abort.
	//
	// See the guidelines above for deciding between `FAILED_PRECONDITION`,
	// `ABORTED`, and `UNAVAILABLE`.
	//
	// HTTP Mapping: 409 Conflict
	Code_CODE_ABORTED Code = 12
	// The operation was attempted past the valid range.  E.g., seeking or
	// reading past end-of-file.
	//
	// Unlike `INVALID_ARGUMENT`, this error indicates a problem that may
	// be fixed if the system state changes. For example, a 32-bit file
	// system will generate `INVALID_ARGUMENT` if asked to read at an
	// offset that is not in the range [0,2^32-1], but it will generate
	// `OUT_OF_RANGE` if asked to read from an offset past the current
	// file size.
	//
	// There is a fair bit of overlap between `FAILED_PRECONDITION` and
	// `OUT_OF_RANGE`.  We recommend using `OUT_OF_RANGE` (the more specific
	// error) when it applies so that callers who are iterating through
	// a space can easily look for an `OUT_OF_RANGE` error to detect when
	// they are done.
	//
	// HTTP Mapping: 400 Bad Request
	Code_CODE_OUT_OF_RANGE Code = 13
	// The operation is not implemented or is not supported/enabled in this
	// service.
	//
	// HTTP Mapping: 501 Not Implemented
	Code_CODE_UNIMPLEMENTED Code = 14
	// Internal errors.  This means that some invariants expected by the
	// underlying system have been broken.  This error code is reserved
	// for serious errors.
	//
	// HTTP Mapping: 500 Internal Server Error
	Code_CODE_INTERNAL Code = 15
	// The service is currently unavailable.  This is most likely a
	// transient condition, which can be corrected by retrying with
	// a backoff.
	//
	// See the guidelines above for deciding between `FAILED_PRECONDITION`,
	// `ABORTED`, and `UNAVAILABLE`.
	//
	// HTTP Mapping: 503 Service Unavailable
	Code_CODE_UNAVAILABLE Code = 16
	// Unrecoverable data loss or corruption.
	//
	// HTTP Mapping: 500 Internal Server Error
	Code_CODE_DATA_LOSS Code = 17
	// Redirects the operation to another location.
	// Used in a Status reponse with a reference to the target URI.
	Code_CODE_REDIRECTION Code = 18
)

var Code_name = map[int32]string{
	0:  "CODE_INVALID",
	1:  "CODE_OK",
	2:  "CODE_CANCELLED",
	3:  "CODE_UNKNOWN",
	4:  "CODE_INVALID_ARGUMENT",
	5:  "CODE_DEADLINE_EXCEEDED",
	6:  "CODE_NOT_FOUND",
	7:  "CODE_ALREADY_EXISTS",
	8:  "CODE_PERMISSION_DENIED",
	9:  "CODE_UNAUTHENTICATED",
	10: "CODE_RESOURCE_EXHAUSTED",
	11: "CODE_FAILED_PRECONDITION",
	12: "CODE_ABORTED",
	13: "CODE_OUT_OF_RANGE",
	14: "CODE_UNIMPLEMENTED",
	15: "CODE_INTERNAL",
	16: "CODE_UNAVAILABLE",
	17: "CODE_DATA_LOSS",
	18: "CODE_REDIRECTION",
}

var Code_value = map[string]int32{
	"CODE_INVALID":             0,
	"CODE_OK":                  1,
	"CODE_CANCELLED":           2,
	"CODE_UNKNOWN":             3,
	"CODE_INVALID_ARGUMENT":    4,
	"CODE_DEADLINE_EXCEEDED":   5,
	"CODE_NOT_FOUND":           6,
	"CODE_ALREADY_EXISTS":      7,
	"CODE_PERMISSION_DENIED":   8,
	"CODE_UNAUTHENTICATED":     9,
	"CODE_RESOURCE_EXHAUSTED":  10,
	"CODE_FAILED_PRECONDITION": 11,
	"CODE_ABORTED":             12,
	"CODE_OUT_OF_RANGE":        13,
	"CODE_UNIMPLEMENTED":       14,
	"CODE_INTERNAL":            15,
	"CODE_UNAVAILABLE":         16,
	"CODE_DATA_LOSS":           17,
	"CODE_REDIRECTION":         18,
}

func (x Code) String() string {
	return proto.EnumName(Code_name, int32(x))
}

func (Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d7978aa548ead3ec, []int{0}
}

func init() {
	proto.RegisterEnum("cs3.rpc.v1beta1.Code", Code_name, Code_value)
}

func init() { proto.RegisterFile("cs3/rpc/v1beta1/code.proto", fileDescriptor_d7978aa548ead3ec) }

var fileDescriptor_d7978aa548ead3ec = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc7, 0x59, 0x3b, 0x36, 0x76, 0xf6, 0xd1, 0x53, 0xef, 0x93, 0xc1, 0x13, 0x70, 0x91, 0xaa,
	0xea, 0x13, 0x38, 0xf6, 0xe9, 0x66, 0xcd, 0xb5, 0x23, 0xc7, 0x29, 0x05, 0x4d, 0xb2, 0x98, 0xd7,
	0x4b, 0x94, 0x28, 0x9d, 0x78, 0x20, 0x2e, 0xb9, 0xe6, 0x29, 0x78, 0x2a, 0x14, 0x13, 0x22, 0xc4,
	0x65, 0xfe, 0xbf, 0x93, 0xf3, 0xf1, 0xff, 0x1b, 0x6e, 0xe3, 0x6e, 0x31, 0x6b, 0x9b, 0x38, 0xfb,
	0x36, 0x7f, 0xda, 0xbe, 0x7c, 0x99, 0xcf, 0x62, 0xfd, 0xbc, 0xcd, 0x9a, 0xb6, 0x7e, 0xa9, 0xd9,
	0x24, 0xee, 0x16, 0x59, 0xdb, 0xc4, 0xac, 0x67, 0x1f, 0x7e, 0x8e, 0x61, 0x5f, 0xd4, 0xcf, 0x5b,
	0x86, 0x70, 0x22, 0xac, 0xa4, 0xa0, 0xcc, 0x9a, 0x6b, 0x25, 0xf1, 0x15, 0x3b, 0x86, 0xc3, 0xa4,
	0xd8, 0x07, 0xdc, 0x63, 0x0c, 0xce, 0xd2, 0x87, 0xe0, 0x46, 0x90, 0xd6, 0x24, 0x71, 0x34, 0xfc,
	0x52, 0x99, 0x07, 0x63, 0x3f, 0x1a, 0x1c, 0xb3, 0xb7, 0x70, 0xf9, 0x6f, 0x93, 0xc0, 0xdd, 0x5d,
	0xb5, 0x22, 0xe3, 0x71, 0x9f, 0xdd, 0xc2, 0x55, 0x42, 0x92, 0xb8, 0xd4, 0xca, 0x50, 0xa0, 0x8d,
	0x20, 0x92, 0x24, 0xf1, 0xf5, 0xd0, 0xdc, 0x58, 0x1f, 0x96, 0xb6, 0x32, 0x12, 0x0f, 0xd8, 0x35,
	0x9c, 0x27, 0x8d, 0x6b, 0x47, 0x5c, 0x7e, 0x0a, 0xb4, 0x51, 0xa5, 0x2f, 0xf1, 0x70, 0x68, 0x54,
	0x90, 0x5b, 0xa9, 0xb2, 0x54, 0xd6, 0x04, 0x49, 0x46, 0x91, 0xc4, 0x37, 0xec, 0x06, 0x2e, 0xfa,
	0x8d, 0x78, 0xe5, 0xef, 0xc9, 0x78, 0x25, 0xb8, 0x27, 0x89, 0x47, 0xec, 0x1d, 0x5c, 0x27, 0xe2,
	0xa8, 0xb4, 0x95, 0x13, 0xdd, 0xf8, 0x7b, 0x5e, 0x95, 0x1d, 0x04, 0xf6, 0x1e, 0x6e, 0x12, 0x5c,
	0x72, 0xa5, 0x49, 0x86, 0xc2, 0x91, 0xb0, 0x46, 0x2a, 0xaf, 0xac, 0xc1, 0xe3, 0xe1, 0x4c, 0x9e,
	0x5b, 0xd7, 0xd5, 0x9f, 0xb0, 0x4b, 0x98, 0xfe, 0x71, 0xa6, 0xf2, 0xc1, 0x2e, 0x83, 0xe3, 0xe6,
	0x8e, 0xf0, 0x94, 0x5d, 0x01, 0xeb, 0xa7, 0xab, 0x55, 0xa1, 0xa9, 0x3b, 0x9c, 0x24, 0x9e, 0xb1,
	0x29, 0x9c, 0xf6, 0xae, 0x78, 0x72, 0x86, 0x6b, 0x9c, 0xb0, 0x0b, 0xc0, 0xbf, 0x8b, 0xae, 0xb9,
	0xd2, 0x3c, 0xd7, 0x84, 0x38, 0xf8, 0x20, 0xb9, 0xe7, 0x41, 0xdb, 0xb2, 0xc4, 0xe9, 0x50, 0xe9,
	0x48, 0x2a, 0x47, 0x22, 0xed, 0xc4, 0xf2, 0x0d, 0x9c, 0xc7, 0xfa, 0x6b, 0xf6, 0x5f, 0x9a, 0xf9,
	0x51, 0x17, 0x65, 0xd1, 0x25, 0x5d, 0xec, 0x7d, 0x86, 0xb6, 0x89, 0x3d, 0xf8, 0x3e, 0x1a, 0x0b,
	0xb7, 0xf9, 0x31, 0x9a, 0x88, 0xdd, 0x22, 0x73, 0x4d, 0xcc, 0xd6, 0xf3, 0xbc, 0xd3, 0x7f, 0x25,
	0xe5, 0xd1, 0x35, 0xf1, 0xb1, 0x57, 0x9e, 0x0e, 0xd2, 0x43, 0x59, 0xfc, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0x5d, 0xe4, 0xb1, 0x9c, 0x46, 0x02, 0x00, 0x00,
}