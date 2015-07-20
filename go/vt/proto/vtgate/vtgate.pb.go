// Code generated by protoc-gen-go.
// source: vtgate.proto
// DO NOT EDIT!

/*
Package vtgate is a generated protocol buffer package.

It is generated from these files:
	vtgate.proto

It has these top-level messages:
	Session
	ExecuteRequest
	ExecuteResponse
	ExecuteShardsRequest
	ExecuteShardsResponse
	ExecuteKeyspaceIdsRequest
	ExecuteKeyspaceIdsResponse
	ExecuteKeyRangesRequest
	ExecuteKeyRangesResponse
	ExecuteEntityIdsRequest
	ExecuteEntityIdsResponse
	BoundShardQuery
	ExecuteBatchShardsRequest
	ExecuteBatchShardsResponse
	BoundKeyspaceIdQuery
	ExecuteBatchKeyspaceIdsRequest
	ExecuteBatchKeyspaceIdsResponse
	StreamExecuteRequest
	StreamExecuteResponse
	StreamExecuteShardsRequest
	StreamExecuteShardsResponse
	StreamExecuteKeyspaceIdsRequest
	StreamExecuteKeyspaceIdsResponse
	StreamExecuteKeyRangesRequest
	StreamExecuteKeyRangesResponse
	BeginRequest
	BeginResponse
	CommitRequest
	CommitResponse
	RollbackRequest
	RollbackResponse
	SplitQueryRequest
	SplitQueryResponse
*/
package vtgate

import proto "github.com/golang/protobuf/proto"
import query "github.com/youtube/vitess/go/vt/proto/query"
import topodata "github.com/youtube/vitess/go/vt/proto/topodata"
import vtrpc "github.com/youtube/vitess/go/vt/proto/vtrpc"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type ExecuteEntityIdsRequest_EntityId_Type int32

const (
	ExecuteEntityIdsRequest_EntityId_TYPE_NULL  ExecuteEntityIdsRequest_EntityId_Type = 0
	ExecuteEntityIdsRequest_EntityId_TYPE_BYTES ExecuteEntityIdsRequest_EntityId_Type = 1
	ExecuteEntityIdsRequest_EntityId_TYPE_INT   ExecuteEntityIdsRequest_EntityId_Type = 2
	ExecuteEntityIdsRequest_EntityId_TYPE_UINT  ExecuteEntityIdsRequest_EntityId_Type = 3
	ExecuteEntityIdsRequest_EntityId_TYPE_FLOAT ExecuteEntityIdsRequest_EntityId_Type = 4
)

var ExecuteEntityIdsRequest_EntityId_Type_name = map[int32]string{
	0: "TYPE_NULL",
	1: "TYPE_BYTES",
	2: "TYPE_INT",
	3: "TYPE_UINT",
	4: "TYPE_FLOAT",
}
var ExecuteEntityIdsRequest_EntityId_Type_value = map[string]int32{
	"TYPE_NULL":  0,
	"TYPE_BYTES": 1,
	"TYPE_INT":   2,
	"TYPE_UINT":  3,
	"TYPE_FLOAT": 4,
}

func (x ExecuteEntityIdsRequest_EntityId_Type) String() string {
	return proto.EnumName(ExecuteEntityIdsRequest_EntityId_Type_name, int32(x))
}

// Session objects are session cookies and are invalidated on
// use. Query results will contain updated session values.
// Their content should be opaque to the user.
type Session struct {
	InTransaction bool                    `protobuf:"varint,1,opt,name=in_transaction" json:"in_transaction,omitempty"`
	ShardSessions []*Session_ShardSession `protobuf:"bytes,2,rep,name=shard_sessions" json:"shard_sessions,omitempty"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}

func (m *Session) GetShardSessions() []*Session_ShardSession {
	if m != nil {
		return m.ShardSessions
	}
	return nil
}

type Session_ShardSession struct {
	Target        *query.Target `protobuf:"bytes,1,opt,name=target" json:"target,omitempty"`
	TransactionId int64         `protobuf:"varint,2,opt,name=transaction_id" json:"transaction_id,omitempty"`
}

func (m *Session_ShardSession) Reset()         { *m = Session_ShardSession{} }
func (m *Session_ShardSession) String() string { return proto.CompactTextString(m) }
func (*Session_ShardSession) ProtoMessage()    {}

func (m *Session_ShardSession) GetTarget() *query.Target {
	if m != nil {
		return m.Target
	}
	return nil
}

// ExecuteRequest is the payload to Execute
type ExecuteRequest struct {
	CallerId         *vtrpc.CallerID     `protobuf:"bytes,1,opt,name=caller_id" json:"caller_id,omitempty"`
	Session          *Session            `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
	Query            *query.BoundQuery   `protobuf:"bytes,3,opt,name=query" json:"query,omitempty"`
	TabletType       topodata.TabletType `protobuf:"varint,4,opt,enum=topodata.TabletType" json:"TabletType,omitempty"`
	NotInTransaction bool                `protobuf:"varint,5,opt,name=not_in_transaction" json:"not_in_transaction,omitempty"`
}

func (m *ExecuteRequest) Reset()         { *m = ExecuteRequest{} }
func (m *ExecuteRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteRequest) ProtoMessage()    {}

func (m *ExecuteRequest) GetCallerId() *vtrpc.CallerID {
	if m != nil {
		return m.CallerId
	}
	return nil
}

func (m *ExecuteRequest) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *ExecuteRequest) GetQuery() *query.BoundQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

// ExecuteResponse is the returned value from Execute
type ExecuteResponse struct {
	Error   *vtrpc.RPCError    `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Session *Session           `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
	Result  *query.QueryResult `protobuf:"bytes,3,opt,name=result" json:"result,omitempty"`
}

func (m *ExecuteResponse) Reset()         { *m = ExecuteResponse{} }
func (m *ExecuteResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteResponse) ProtoMessage()    {}

func (m *ExecuteResponse) GetError() *vtrpc.RPCError {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ExecuteResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *ExecuteResponse) GetResult() *query.QueryResult {
	if m != nil {
		return m.Result
	}
	return nil
}

// ExecuteShardsRequest is the payload to ExecuteShards
type ExecuteShardsRequest struct {
	CallerId         *vtrpc.CallerID     `protobuf:"bytes,1,opt,name=caller_id" json:"caller_id,omitempty"`
	Session          *Session            `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
	Query            *query.BoundQuery   `protobuf:"bytes,3,opt,name=query" json:"query,omitempty"`
	Keyspace         string              `protobuf:"bytes,4,opt,name=keyspace" json:"keyspace,omitempty"`
	Shards           []string            `protobuf:"bytes,5,rep,name=shards" json:"shards,omitempty"`
	TabletType       topodata.TabletType `protobuf:"varint,6,opt,enum=topodata.TabletType" json:"TabletType,omitempty"`
	NotInTransaction bool                `protobuf:"varint,7,opt,name=not_in_transaction" json:"not_in_transaction,omitempty"`
}

func (m *ExecuteShardsRequest) Reset()         { *m = ExecuteShardsRequest{} }
func (m *ExecuteShardsRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteShardsRequest) ProtoMessage()    {}

func (m *ExecuteShardsRequest) GetCallerId() *vtrpc.CallerID {
	if m != nil {
		return m.CallerId
	}
	return nil
}

func (m *ExecuteShardsRequest) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *ExecuteShardsRequest) GetQuery() *query.BoundQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

// ExecuteShardsResponse is the returned value from ExecuteShards
type ExecuteShardsResponse struct {
	Error   *vtrpc.RPCError    `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Session *Session           `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
	Result  *query.QueryResult `protobuf:"bytes,3,opt,name=result" json:"result,omitempty"`
}

func (m *ExecuteShardsResponse) Reset()         { *m = ExecuteShardsResponse{} }
func (m *ExecuteShardsResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteShardsResponse) ProtoMessage()    {}

func (m *ExecuteShardsResponse) GetError() *vtrpc.RPCError {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ExecuteShardsResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *ExecuteShardsResponse) GetResult() *query.QueryResult {
	if m != nil {
		return m.Result
	}
	return nil
}

// ExecuteKeyspaceIdsRequest is the payload to ExecuteKeyspaceIds
type ExecuteKeyspaceIdsRequest struct {
	CallerId         *vtrpc.CallerID     `protobuf:"bytes,1,opt,name=caller_id" json:"caller_id,omitempty"`
	Session          *Session            `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
	Query            *query.BoundQuery   `protobuf:"bytes,3,opt,name=query" json:"query,omitempty"`
	Keyspace         string              `protobuf:"bytes,4,opt,name=keyspace" json:"keyspace,omitempty"`
	KeyspaceIds      [][]byte            `protobuf:"bytes,5,rep,name=keyspace_ids,proto3" json:"keyspace_ids,omitempty"`
	TabletType       topodata.TabletType `protobuf:"varint,6,opt,enum=topodata.TabletType" json:"TabletType,omitempty"`
	NotInTransaction bool                `protobuf:"varint,7,opt,name=not_in_transaction" json:"not_in_transaction,omitempty"`
}

func (m *ExecuteKeyspaceIdsRequest) Reset()         { *m = ExecuteKeyspaceIdsRequest{} }
func (m *ExecuteKeyspaceIdsRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteKeyspaceIdsRequest) ProtoMessage()    {}

func (m *ExecuteKeyspaceIdsRequest) GetCallerId() *vtrpc.CallerID {
	if m != nil {
		return m.CallerId
	}
	return nil
}

func (m *ExecuteKeyspaceIdsRequest) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *ExecuteKeyspaceIdsRequest) GetQuery() *query.BoundQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

// ExecuteKeyspaceIdsResponse is the returned value from ExecuteKeyspaceIds
type ExecuteKeyspaceIdsResponse struct {
	Error   *vtrpc.RPCError    `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Session *Session           `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
	Result  *query.QueryResult `protobuf:"bytes,3,opt,name=result" json:"result,omitempty"`
}

func (m *ExecuteKeyspaceIdsResponse) Reset()         { *m = ExecuteKeyspaceIdsResponse{} }
func (m *ExecuteKeyspaceIdsResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteKeyspaceIdsResponse) ProtoMessage()    {}

func (m *ExecuteKeyspaceIdsResponse) GetError() *vtrpc.RPCError {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ExecuteKeyspaceIdsResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *ExecuteKeyspaceIdsResponse) GetResult() *query.QueryResult {
	if m != nil {
		return m.Result
	}
	return nil
}

// ExecuteKeyRangesRequest is the payload to ExecuteKeyRanges
type ExecuteKeyRangesRequest struct {
	CallerId         *vtrpc.CallerID      `protobuf:"bytes,1,opt,name=caller_id" json:"caller_id,omitempty"`
	Session          *Session             `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
	Query            *query.BoundQuery    `protobuf:"bytes,3,opt,name=query" json:"query,omitempty"`
	Keyspace         string               `protobuf:"bytes,4,opt,name=keyspace" json:"keyspace,omitempty"`
	KeyRanges        []*topodata.KeyRange `protobuf:"bytes,5,rep,name=key_ranges" json:"key_ranges,omitempty"`
	TabletType       topodata.TabletType  `protobuf:"varint,6,opt,enum=topodata.TabletType" json:"TabletType,omitempty"`
	NotInTransaction bool                 `protobuf:"varint,7,opt,name=not_in_transaction" json:"not_in_transaction,omitempty"`
}

func (m *ExecuteKeyRangesRequest) Reset()         { *m = ExecuteKeyRangesRequest{} }
func (m *ExecuteKeyRangesRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteKeyRangesRequest) ProtoMessage()    {}

func (m *ExecuteKeyRangesRequest) GetCallerId() *vtrpc.CallerID {
	if m != nil {
		return m.CallerId
	}
	return nil
}

func (m *ExecuteKeyRangesRequest) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *ExecuteKeyRangesRequest) GetQuery() *query.BoundQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *ExecuteKeyRangesRequest) GetKeyRanges() []*topodata.KeyRange {
	if m != nil {
		return m.KeyRanges
	}
	return nil
}

// ExecuteKeyRangesResponse is the returned value from ExecuteKeyRanges
type ExecuteKeyRangesResponse struct {
	Error   *vtrpc.RPCError    `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Session *Session           `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
	Result  *query.QueryResult `protobuf:"bytes,3,opt,name=result" json:"result,omitempty"`
}

func (m *ExecuteKeyRangesResponse) Reset()         { *m = ExecuteKeyRangesResponse{} }
func (m *ExecuteKeyRangesResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteKeyRangesResponse) ProtoMessage()    {}

func (m *ExecuteKeyRangesResponse) GetError() *vtrpc.RPCError {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ExecuteKeyRangesResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *ExecuteKeyRangesResponse) GetResult() *query.QueryResult {
	if m != nil {
		return m.Result
	}
	return nil
}

// ExecuteEntityIdsRequest is the payload to ExecuteEntityIds
type ExecuteEntityIdsRequest struct {
	CallerId          *vtrpc.CallerID                     `protobuf:"bytes,1,opt,name=caller_id" json:"caller_id,omitempty"`
	Session           *Session                            `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
	Query             *query.BoundQuery                   `protobuf:"bytes,3,opt,name=query" json:"query,omitempty"`
	Keyspace          string                              `protobuf:"bytes,4,opt,name=keyspace" json:"keyspace,omitempty"`
	EntityColumnName  string                              `protobuf:"bytes,5,opt,name=entity_column_name" json:"entity_column_name,omitempty"`
	EntityKeyspaceIds []*ExecuteEntityIdsRequest_EntityId `protobuf:"bytes,6,rep,name=entity_keyspace_ids" json:"entity_keyspace_ids,omitempty"`
	TabletType        topodata.TabletType                 `protobuf:"varint,7,opt,enum=topodata.TabletType" json:"TabletType,omitempty"`
	NotInTransaction  bool                                `protobuf:"varint,8,opt,name=not_in_transaction" json:"not_in_transaction,omitempty"`
}

func (m *ExecuteEntityIdsRequest) Reset()         { *m = ExecuteEntityIdsRequest{} }
func (m *ExecuteEntityIdsRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteEntityIdsRequest) ProtoMessage()    {}

func (m *ExecuteEntityIdsRequest) GetCallerId() *vtrpc.CallerID {
	if m != nil {
		return m.CallerId
	}
	return nil
}

func (m *ExecuteEntityIdsRequest) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *ExecuteEntityIdsRequest) GetQuery() *query.BoundQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *ExecuteEntityIdsRequest) GetEntityKeyspaceIds() []*ExecuteEntityIdsRequest_EntityId {
	if m != nil {
		return m.EntityKeyspaceIds
	}
	return nil
}

type ExecuteEntityIdsRequest_EntityId struct {
	XidType    ExecuteEntityIdsRequest_EntityId_Type `protobuf:"varint,1,opt,name=xid_type,enum=vtgate.ExecuteEntityIdsRequest_EntityId_Type" json:"xid_type,omitempty"`
	XidBytes   []byte                                `protobuf:"bytes,2,opt,name=xid_bytes,proto3" json:"xid_bytes,omitempty"`
	XidInt     int64                                 `protobuf:"varint,3,opt,name=xid_int" json:"xid_int,omitempty"`
	XidUint    uint64                                `protobuf:"varint,4,opt,name=xid_uint" json:"xid_uint,omitempty"`
	XidFloat   float64                               `protobuf:"fixed64,5,opt,name=xid_float" json:"xid_float,omitempty"`
	KeyspaceId []byte                                `protobuf:"bytes,6,opt,name=keyspace_id,proto3" json:"keyspace_id,omitempty"`
}

func (m *ExecuteEntityIdsRequest_EntityId) Reset()         { *m = ExecuteEntityIdsRequest_EntityId{} }
func (m *ExecuteEntityIdsRequest_EntityId) String() string { return proto.CompactTextString(m) }
func (*ExecuteEntityIdsRequest_EntityId) ProtoMessage()    {}

// ExecuteEntityIdsResponse is the returned value from ExecuteEntityIds
type ExecuteEntityIdsResponse struct {
	Error   *vtrpc.RPCError    `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Session *Session           `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
	Result  *query.QueryResult `protobuf:"bytes,3,opt,name=result" json:"result,omitempty"`
}

func (m *ExecuteEntityIdsResponse) Reset()         { *m = ExecuteEntityIdsResponse{} }
func (m *ExecuteEntityIdsResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteEntityIdsResponse) ProtoMessage()    {}

func (m *ExecuteEntityIdsResponse) GetError() *vtrpc.RPCError {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ExecuteEntityIdsResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *ExecuteEntityIdsResponse) GetResult() *query.QueryResult {
	if m != nil {
		return m.Result
	}
	return nil
}

// BoundShardQuery represents a single query request for the
// specified list of shards. This is used in a list for
// ExecuteBatchShardsRequest.
type BoundShardQuery struct {
	Query    *query.BoundQuery `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
	Keyspace string            `protobuf:"bytes,2,opt,name=keyspace" json:"keyspace,omitempty"`
	Shards   []string          `protobuf:"bytes,3,rep,name=shards" json:"shards,omitempty"`
}

func (m *BoundShardQuery) Reset()         { *m = BoundShardQuery{} }
func (m *BoundShardQuery) String() string { return proto.CompactTextString(m) }
func (*BoundShardQuery) ProtoMessage()    {}

func (m *BoundShardQuery) GetQuery() *query.BoundQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

// ExecuteBatchShardsRequest is the payload to ExecuteBatchShards
type ExecuteBatchShardsRequest struct {
	CallerId      *vtrpc.CallerID     `protobuf:"bytes,1,opt,name=caller_id" json:"caller_id,omitempty"`
	Session       *Session            `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
	Queries       []*BoundShardQuery  `protobuf:"bytes,3,rep,name=queries" json:"queries,omitempty"`
	TabletType    topodata.TabletType `protobuf:"varint,4,opt,enum=topodata.TabletType" json:"TabletType,omitempty"`
	AsTransaction bool                `protobuf:"varint,5,opt,name=as_transaction" json:"as_transaction,omitempty"`
}

func (m *ExecuteBatchShardsRequest) Reset()         { *m = ExecuteBatchShardsRequest{} }
func (m *ExecuteBatchShardsRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteBatchShardsRequest) ProtoMessage()    {}

func (m *ExecuteBatchShardsRequest) GetCallerId() *vtrpc.CallerID {
	if m != nil {
		return m.CallerId
	}
	return nil
}

func (m *ExecuteBatchShardsRequest) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *ExecuteBatchShardsRequest) GetQueries() []*BoundShardQuery {
	if m != nil {
		return m.Queries
	}
	return nil
}

// ExecuteBatchShardsResponse is the returned value from ExecuteBatchShards
type ExecuteBatchShardsResponse struct {
	Error   *vtrpc.RPCError      `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Session *Session             `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
	Results []*query.QueryResult `protobuf:"bytes,3,rep,name=results" json:"results,omitempty"`
}

func (m *ExecuteBatchShardsResponse) Reset()         { *m = ExecuteBatchShardsResponse{} }
func (m *ExecuteBatchShardsResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteBatchShardsResponse) ProtoMessage()    {}

func (m *ExecuteBatchShardsResponse) GetError() *vtrpc.RPCError {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ExecuteBatchShardsResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *ExecuteBatchShardsResponse) GetResults() []*query.QueryResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// BoundKeyspaceIdQuery represents a single query request for the
// specified list of keyspace ids. This is used in a list for
// ExecuteBatchKeyspaceIdsRequest.
type BoundKeyspaceIdQuery struct {
	Query       *query.BoundQuery `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
	Keyspace    string            `protobuf:"bytes,2,opt,name=keyspace" json:"keyspace,omitempty"`
	KeyspaceIds [][]byte          `protobuf:"bytes,3,rep,name=keyspace_ids,proto3" json:"keyspace_ids,omitempty"`
}

func (m *BoundKeyspaceIdQuery) Reset()         { *m = BoundKeyspaceIdQuery{} }
func (m *BoundKeyspaceIdQuery) String() string { return proto.CompactTextString(m) }
func (*BoundKeyspaceIdQuery) ProtoMessage()    {}

func (m *BoundKeyspaceIdQuery) GetQuery() *query.BoundQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

// ExecuteBatchKeyspaceIdsRequest is the payload to ExecuteBatchKeyspaceId
type ExecuteBatchKeyspaceIdsRequest struct {
	CallerId      *vtrpc.CallerID         `protobuf:"bytes,1,opt,name=caller_id" json:"caller_id,omitempty"`
	Session       *Session                `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
	Queries       []*BoundKeyspaceIdQuery `protobuf:"bytes,3,rep,name=queries" json:"queries,omitempty"`
	TabletType    topodata.TabletType     `protobuf:"varint,4,opt,enum=topodata.TabletType" json:"TabletType,omitempty"`
	AsTransaction bool                    `protobuf:"varint,5,opt,name=as_transaction" json:"as_transaction,omitempty"`
}

func (m *ExecuteBatchKeyspaceIdsRequest) Reset()         { *m = ExecuteBatchKeyspaceIdsRequest{} }
func (m *ExecuteBatchKeyspaceIdsRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteBatchKeyspaceIdsRequest) ProtoMessage()    {}

func (m *ExecuteBatchKeyspaceIdsRequest) GetCallerId() *vtrpc.CallerID {
	if m != nil {
		return m.CallerId
	}
	return nil
}

func (m *ExecuteBatchKeyspaceIdsRequest) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *ExecuteBatchKeyspaceIdsRequest) GetQueries() []*BoundKeyspaceIdQuery {
	if m != nil {
		return m.Queries
	}
	return nil
}

// ExecuteBatchKeyspaceIdsResponse is the returned value from ExecuteBatchKeyspaceId
type ExecuteBatchKeyspaceIdsResponse struct {
	Error   *vtrpc.RPCError      `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Session *Session             `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
	Results []*query.QueryResult `protobuf:"bytes,3,rep,name=results" json:"results,omitempty"`
}

func (m *ExecuteBatchKeyspaceIdsResponse) Reset()         { *m = ExecuteBatchKeyspaceIdsResponse{} }
func (m *ExecuteBatchKeyspaceIdsResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteBatchKeyspaceIdsResponse) ProtoMessage()    {}

func (m *ExecuteBatchKeyspaceIdsResponse) GetError() *vtrpc.RPCError {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ExecuteBatchKeyspaceIdsResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *ExecuteBatchKeyspaceIdsResponse) GetResults() []*query.QueryResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// StreamExecuteRequest is the payload to StreamExecute
type StreamExecuteRequest struct {
	CallerId   *vtrpc.CallerID     `protobuf:"bytes,1,opt,name=caller_id" json:"caller_id,omitempty"`
	Query      *query.BoundQuery   `protobuf:"bytes,2,opt,name=query" json:"query,omitempty"`
	TabletType topodata.TabletType `protobuf:"varint,3,opt,enum=topodata.TabletType" json:"TabletType,omitempty"`
}

func (m *StreamExecuteRequest) Reset()         { *m = StreamExecuteRequest{} }
func (m *StreamExecuteRequest) String() string { return proto.CompactTextString(m) }
func (*StreamExecuteRequest) ProtoMessage()    {}

func (m *StreamExecuteRequest) GetCallerId() *vtrpc.CallerID {
	if m != nil {
		return m.CallerId
	}
	return nil
}

func (m *StreamExecuteRequest) GetQuery() *query.BoundQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

// StreamExecuteResponse is the returned value from StreamExecute
type StreamExecuteResponse struct {
	Error  *vtrpc.RPCError    `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Result *query.QueryResult `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
}

func (m *StreamExecuteResponse) Reset()         { *m = StreamExecuteResponse{} }
func (m *StreamExecuteResponse) String() string { return proto.CompactTextString(m) }
func (*StreamExecuteResponse) ProtoMessage()    {}

func (m *StreamExecuteResponse) GetError() *vtrpc.RPCError {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *StreamExecuteResponse) GetResult() *query.QueryResult {
	if m != nil {
		return m.Result
	}
	return nil
}

// StreamExecuteShardsRequest is the payload to StreamExecuteShards
type StreamExecuteShardsRequest struct {
	CallerId   *vtrpc.CallerID     `protobuf:"bytes,1,opt,name=caller_id" json:"caller_id,omitempty"`
	Query      *query.BoundQuery   `protobuf:"bytes,2,opt,name=query" json:"query,omitempty"`
	Keyspace   string              `protobuf:"bytes,3,opt,name=keyspace" json:"keyspace,omitempty"`
	Shards     []string            `protobuf:"bytes,4,rep,name=shards" json:"shards,omitempty"`
	TabletType topodata.TabletType `protobuf:"varint,5,opt,enum=topodata.TabletType" json:"TabletType,omitempty"`
}

func (m *StreamExecuteShardsRequest) Reset()         { *m = StreamExecuteShardsRequest{} }
func (m *StreamExecuteShardsRequest) String() string { return proto.CompactTextString(m) }
func (*StreamExecuteShardsRequest) ProtoMessage()    {}

func (m *StreamExecuteShardsRequest) GetCallerId() *vtrpc.CallerID {
	if m != nil {
		return m.CallerId
	}
	return nil
}

func (m *StreamExecuteShardsRequest) GetQuery() *query.BoundQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

// StreamExecuteShardsResponse is the returned value from StreamExecuteShards
type StreamExecuteShardsResponse struct {
	Error  *vtrpc.RPCError    `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Result *query.QueryResult `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
}

func (m *StreamExecuteShardsResponse) Reset()         { *m = StreamExecuteShardsResponse{} }
func (m *StreamExecuteShardsResponse) String() string { return proto.CompactTextString(m) }
func (*StreamExecuteShardsResponse) ProtoMessage()    {}

func (m *StreamExecuteShardsResponse) GetError() *vtrpc.RPCError {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *StreamExecuteShardsResponse) GetResult() *query.QueryResult {
	if m != nil {
		return m.Result
	}
	return nil
}

// StreamExecuteKeyspaceIdsRequest is the payload to StreamExecuteKeyspaceIds
type StreamExecuteKeyspaceIdsRequest struct {
	CallerId    *vtrpc.CallerID     `protobuf:"bytes,1,opt,name=caller_id" json:"caller_id,omitempty"`
	Query       *query.BoundQuery   `protobuf:"bytes,2,opt,name=query" json:"query,omitempty"`
	Keyspace    string              `protobuf:"bytes,3,opt,name=keyspace" json:"keyspace,omitempty"`
	KeyspaceIds [][]byte            `protobuf:"bytes,4,rep,name=keyspace_ids,proto3" json:"keyspace_ids,omitempty"`
	TabletType  topodata.TabletType `protobuf:"varint,5,opt,enum=topodata.TabletType" json:"TabletType,omitempty"`
}

func (m *StreamExecuteKeyspaceIdsRequest) Reset()         { *m = StreamExecuteKeyspaceIdsRequest{} }
func (m *StreamExecuteKeyspaceIdsRequest) String() string { return proto.CompactTextString(m) }
func (*StreamExecuteKeyspaceIdsRequest) ProtoMessage()    {}

func (m *StreamExecuteKeyspaceIdsRequest) GetCallerId() *vtrpc.CallerID {
	if m != nil {
		return m.CallerId
	}
	return nil
}

func (m *StreamExecuteKeyspaceIdsRequest) GetQuery() *query.BoundQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

// StreamExecuteKeyspaceIdsResponse is the returned value from StreamExecuteKeyspaceIds
type StreamExecuteKeyspaceIdsResponse struct {
	Error  *vtrpc.RPCError    `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Result *query.QueryResult `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
}

func (m *StreamExecuteKeyspaceIdsResponse) Reset()         { *m = StreamExecuteKeyspaceIdsResponse{} }
func (m *StreamExecuteKeyspaceIdsResponse) String() string { return proto.CompactTextString(m) }
func (*StreamExecuteKeyspaceIdsResponse) ProtoMessage()    {}

func (m *StreamExecuteKeyspaceIdsResponse) GetError() *vtrpc.RPCError {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *StreamExecuteKeyspaceIdsResponse) GetResult() *query.QueryResult {
	if m != nil {
		return m.Result
	}
	return nil
}

// StreamExecuteKeyRangesRequest is the payload to StreamExecuteKeyRanges
type StreamExecuteKeyRangesRequest struct {
	CallerId   *vtrpc.CallerID      `protobuf:"bytes,1,opt,name=caller_id" json:"caller_id,omitempty"`
	Query      *query.BoundQuery    `protobuf:"bytes,2,opt,name=query" json:"query,omitempty"`
	Keyspace   string               `protobuf:"bytes,3,opt,name=keyspace" json:"keyspace,omitempty"`
	KeyRanges  []*topodata.KeyRange `protobuf:"bytes,4,rep,name=key_ranges" json:"key_ranges,omitempty"`
	TabletType topodata.TabletType  `protobuf:"varint,5,opt,enum=topodata.TabletType" json:"TabletType,omitempty"`
}

func (m *StreamExecuteKeyRangesRequest) Reset()         { *m = StreamExecuteKeyRangesRequest{} }
func (m *StreamExecuteKeyRangesRequest) String() string { return proto.CompactTextString(m) }
func (*StreamExecuteKeyRangesRequest) ProtoMessage()    {}

func (m *StreamExecuteKeyRangesRequest) GetCallerId() *vtrpc.CallerID {
	if m != nil {
		return m.CallerId
	}
	return nil
}

func (m *StreamExecuteKeyRangesRequest) GetQuery() *query.BoundQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *StreamExecuteKeyRangesRequest) GetKeyRanges() []*topodata.KeyRange {
	if m != nil {
		return m.KeyRanges
	}
	return nil
}

// StreamExecuteKeyRangesResponse is the returned value from StreamExecuteKeyRanges
type StreamExecuteKeyRangesResponse struct {
	Error  *vtrpc.RPCError    `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Result *query.QueryResult `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
}

func (m *StreamExecuteKeyRangesResponse) Reset()         { *m = StreamExecuteKeyRangesResponse{} }
func (m *StreamExecuteKeyRangesResponse) String() string { return proto.CompactTextString(m) }
func (*StreamExecuteKeyRangesResponse) ProtoMessage()    {}

func (m *StreamExecuteKeyRangesResponse) GetError() *vtrpc.RPCError {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *StreamExecuteKeyRangesResponse) GetResult() *query.QueryResult {
	if m != nil {
		return m.Result
	}
	return nil
}

// BeginRequest is the payload to Begin
type BeginRequest struct {
	CallerId *vtrpc.CallerID `protobuf:"bytes,1,opt,name=caller_id" json:"caller_id,omitempty"`
}

func (m *BeginRequest) Reset()         { *m = BeginRequest{} }
func (m *BeginRequest) String() string { return proto.CompactTextString(m) }
func (*BeginRequest) ProtoMessage()    {}

func (m *BeginRequest) GetCallerId() *vtrpc.CallerID {
	if m != nil {
		return m.CallerId
	}
	return nil
}

// BeginResponse is the returned value from Begin
type BeginResponse struct {
	Error   *vtrpc.RPCError `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Session *Session        `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
}

func (m *BeginResponse) Reset()         { *m = BeginResponse{} }
func (m *BeginResponse) String() string { return proto.CompactTextString(m) }
func (*BeginResponse) ProtoMessage()    {}

func (m *BeginResponse) GetError() *vtrpc.RPCError {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *BeginResponse) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

// CommitRequest is the payload to Commit
type CommitRequest struct {
	CallerId *vtrpc.CallerID `protobuf:"bytes,1,opt,name=caller_id" json:"caller_id,omitempty"`
	Session  *Session        `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
}

func (m *CommitRequest) Reset()         { *m = CommitRequest{} }
func (m *CommitRequest) String() string { return proto.CompactTextString(m) }
func (*CommitRequest) ProtoMessage()    {}

func (m *CommitRequest) GetCallerId() *vtrpc.CallerID {
	if m != nil {
		return m.CallerId
	}
	return nil
}

func (m *CommitRequest) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

// CommitResponse is the returned value from Commit
type CommitResponse struct {
	Error *vtrpc.RPCError `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *CommitResponse) Reset()         { *m = CommitResponse{} }
func (m *CommitResponse) String() string { return proto.CompactTextString(m) }
func (*CommitResponse) ProtoMessage()    {}

func (m *CommitResponse) GetError() *vtrpc.RPCError {
	if m != nil {
		return m.Error
	}
	return nil
}

// RollbackRequest is the payload to Rollback
type RollbackRequest struct {
	CallerId *vtrpc.CallerID `protobuf:"bytes,1,opt,name=caller_id" json:"caller_id,omitempty"`
	Session  *Session        `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
}

func (m *RollbackRequest) Reset()         { *m = RollbackRequest{} }
func (m *RollbackRequest) String() string { return proto.CompactTextString(m) }
func (*RollbackRequest) ProtoMessage()    {}

func (m *RollbackRequest) GetCallerId() *vtrpc.CallerID {
	if m != nil {
		return m.CallerId
	}
	return nil
}

func (m *RollbackRequest) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

// RollbackResponse is the returned value from Rollback
type RollbackResponse struct {
	Error *vtrpc.RPCError `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *RollbackResponse) Reset()         { *m = RollbackResponse{} }
func (m *RollbackResponse) String() string { return proto.CompactTextString(m) }
func (*RollbackResponse) ProtoMessage()    {}

func (m *RollbackResponse) GetError() *vtrpc.RPCError {
	if m != nil {
		return m.Error
	}
	return nil
}

// SplitQueryRequest is the payload to SplitQuery
type SplitQueryRequest struct {
	CallerId   *vtrpc.CallerID   `protobuf:"bytes,1,opt,name=caller_id" json:"caller_id,omitempty"`
	Keyspace   string            `protobuf:"bytes,2,opt,name=keyspace" json:"keyspace,omitempty"`
	Query      *query.BoundQuery `protobuf:"bytes,3,opt,name=query" json:"query,omitempty"`
	SplitCount int64             `protobuf:"varint,4,opt,name=split_count" json:"split_count,omitempty"`
}

func (m *SplitQueryRequest) Reset()         { *m = SplitQueryRequest{} }
func (m *SplitQueryRequest) String() string { return proto.CompactTextString(m) }
func (*SplitQueryRequest) ProtoMessage()    {}

func (m *SplitQueryRequest) GetCallerId() *vtrpc.CallerID {
	if m != nil {
		return m.CallerId
	}
	return nil
}

func (m *SplitQueryRequest) GetQuery() *query.BoundQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

// SplitQueryResponse is the returned value from SplitQuery
type SplitQueryResponse struct {
	Splits []*SplitQueryResponse_Part `protobuf:"bytes,1,rep,name=splits" json:"splits,omitempty"`
}

func (m *SplitQueryResponse) Reset()         { *m = SplitQueryResponse{} }
func (m *SplitQueryResponse) String() string { return proto.CompactTextString(m) }
func (*SplitQueryResponse) ProtoMessage()    {}

func (m *SplitQueryResponse) GetSplits() []*SplitQueryResponse_Part {
	if m != nil {
		return m.Splits
	}
	return nil
}

type SplitQueryResponse_KeyRangePart struct {
	Keyspace  string               `protobuf:"bytes,1,opt,name=keyspace" json:"keyspace,omitempty"`
	KeyRanges []*topodata.KeyRange `protobuf:"bytes,2,rep,name=key_ranges" json:"key_ranges,omitempty"`
}

func (m *SplitQueryResponse_KeyRangePart) Reset()         { *m = SplitQueryResponse_KeyRangePart{} }
func (m *SplitQueryResponse_KeyRangePart) String() string { return proto.CompactTextString(m) }
func (*SplitQueryResponse_KeyRangePart) ProtoMessage()    {}

func (m *SplitQueryResponse_KeyRangePart) GetKeyRanges() []*topodata.KeyRange {
	if m != nil {
		return m.KeyRanges
	}
	return nil
}

type SplitQueryResponse_ShardPart struct {
	Keyspace string   `protobuf:"bytes,1,opt,name=keyspace" json:"keyspace,omitempty"`
	Shards   []string `protobuf:"bytes,2,rep,name=shards" json:"shards,omitempty"`
}

func (m *SplitQueryResponse_ShardPart) Reset()         { *m = SplitQueryResponse_ShardPart{} }
func (m *SplitQueryResponse_ShardPart) String() string { return proto.CompactTextString(m) }
func (*SplitQueryResponse_ShardPart) ProtoMessage()    {}

type SplitQueryResponse_Part struct {
	Query        *query.BoundQuery                `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
	KeyRangePart *SplitQueryResponse_KeyRangePart `protobuf:"bytes,2,opt,name=key_range_part" json:"key_range_part,omitempty"`
	ShardPart    *SplitQueryResponse_ShardPart    `protobuf:"bytes,3,opt,name=shard_part" json:"shard_part,omitempty"`
	Size         int64                            `protobuf:"varint,4,opt,name=size" json:"size,omitempty"`
}

func (m *SplitQueryResponse_Part) Reset()         { *m = SplitQueryResponse_Part{} }
func (m *SplitQueryResponse_Part) String() string { return proto.CompactTextString(m) }
func (*SplitQueryResponse_Part) ProtoMessage()    {}

func (m *SplitQueryResponse_Part) GetQuery() *query.BoundQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

func (m *SplitQueryResponse_Part) GetKeyRangePart() *SplitQueryResponse_KeyRangePart {
	if m != nil {
		return m.KeyRangePart
	}
	return nil
}

func (m *SplitQueryResponse_Part) GetShardPart() *SplitQueryResponse_ShardPart {
	if m != nil {
		return m.ShardPart
	}
	return nil
}

func init() {
	proto.RegisterEnum("vtgate.ExecuteEntityIdsRequest_EntityId_Type", ExecuteEntityIdsRequest_EntityId_Type_name, ExecuteEntityIdsRequest_EntityId_Type_value)
}
