// Code generated by protoc-gen-go. DO NOT EDIT.
// source: topodata.proto

/*
Package topodata is a generated protocol buffer package.

It is generated from these files:
	topodata.proto

It has these top-level messages:
	KeyRange
	TabletAlias
	Tablet
	Shard
	Keyspace
	ShardReplication
	ShardReference
	SrvKeyspace
	CellInfo
*/
package topodata

import proto "github.com/sgoby/sqlparser/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// KeyspaceIdType describes the type of the sharding key for a
// range-based sharded keyspace.
type KeyspaceIdType int32

const (
	// UNSET is the default value, when range-based sharding is not used.
	KeyspaceIdType_UNSET KeyspaceIdType = 0
	// UINT64 is when uint64 value is used.
	// This is represented as 'unsigned bigint' in mysql
	KeyspaceIdType_UINT64 KeyspaceIdType = 1
	// BYTES is when an array of bytes is used.
	// This is represented as 'varbinary' in mysql
	KeyspaceIdType_BYTES KeyspaceIdType = 2
)

var KeyspaceIdType_name = map[int32]string{
	0: "UNSET",
	1: "UINT64",
	2: "BYTES",
}
var KeyspaceIdType_value = map[string]int32{
	"UNSET":  0,
	"UINT64": 1,
	"BYTES":  2,
}

func (x KeyspaceIdType) String() string {
	return proto.EnumName(KeyspaceIdType_name, int32(x))
}
func (KeyspaceIdType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// TabletType represents the type of a given tablet.
type TabletType int32

const (
	// UNKNOWN is not a valid value.
	TabletType_UNKNOWN TabletType = 0
	// MASTER is the master server for the shard. Only MASTER allows DMLs.
	TabletType_MASTER TabletType = 1
	// REPLICA is a slave type. It is used to serve live traffic.
	// A REPLICA can be promoted to MASTER. A demoted MASTER will go to REPLICA.
	TabletType_REPLICA TabletType = 2
	// RDONLY (old name) / BATCH (new name) is used to serve traffic for
	// long-running jobs. It is a separate type from REPLICA so
	// long-running queries don't affect web-like traffic.
	TabletType_RDONLY TabletType = 3
	TabletType_BATCH  TabletType = 3
	// SPARE is a type of servers that cannot serve queries, but is available
	// in case an extra server is needed.
	TabletType_SPARE TabletType = 4
	// EXPERIMENTAL is like SPARE, except it can serve queries. This
	// type can be used for usages not planned by Vitess, like online
	// export to another storage engine.
	TabletType_EXPERIMENTAL TabletType = 5
	// BACKUP is the type a server goes to when taking a backup. No queries
	// can be served in BACKUP mode.
	TabletType_BACKUP TabletType = 6
	// RESTORE is the type a server uses when restoring a backup, at
	// startup time.  No queries can be served in RESTORE mode.
	TabletType_RESTORE TabletType = 7
	// DRAINED is the type a server goes into when used by Vitess tools
	// to perform an offline action. It is a serving type (as
	// the tools processes may need to run queries), but it's not used
	// to route queries from Vitess users. In this state,
	// this tablet is dedicated to the process that uses it.
	TabletType_DRAINED TabletType = 8
)

var TabletType_name = map[int32]string{
	0: "UNKNOWN",
	1: "MASTER",
	2: "REPLICA",
	3: "RDONLY",
	// Duplicate value: 3: "BATCH",
	4: "SPARE",
	5: "EXPERIMENTAL",
	6: "BACKUP",
	7: "RESTORE",
	8: "DRAINED",
}
var TabletType_value = map[string]int32{
	"UNKNOWN":      0,
	"MASTER":       1,
	"REPLICA":      2,
	"RDONLY":       3,
	"BATCH":        3,
	"SPARE":        4,
	"EXPERIMENTAL": 5,
	"BACKUP":       6,
	"RESTORE":      7,
	"DRAINED":      8,
}

func (x TabletType) String() string {
	return proto.EnumName(TabletType_name, int32(x))
}
func (TabletType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// KeyRange describes a range of sharding keys, when range-based
// sharding is used.
type KeyRange struct {
	Start []byte `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End   []byte `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (m *KeyRange) Reset()                    { *m = KeyRange{} }
func (m *KeyRange) String() string            { return proto.CompactTextString(m) }
func (*KeyRange) ProtoMessage()               {}
func (*KeyRange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *KeyRange) GetStart() []byte {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *KeyRange) GetEnd() []byte {
	if m != nil {
		return m.End
	}
	return nil
}

// TabletAlias is a globally unique tablet identifier.
type TabletAlias struct {
	// cell is the cell (or datacenter) the tablet is in
	Cell string `protobuf:"bytes,1,opt,name=cell" json:"cell,omitempty"`
	// uid is a unique id for this tablet within the shard
	// (this is the MySQL server id as well).
	Uid uint32 `protobuf:"varint,2,opt,name=uid" json:"uid,omitempty"`
}

func (m *TabletAlias) Reset()                    { *m = TabletAlias{} }
func (m *TabletAlias) String() string            { return proto.CompactTextString(m) }
func (*TabletAlias) ProtoMessage()               {}
func (*TabletAlias) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TabletAlias) GetCell() string {
	if m != nil {
		return m.Cell
	}
	return ""
}

func (m *TabletAlias) GetUid() uint32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

// Tablet represents information about a running instance of vttablet.
type Tablet struct {
	// alias is the unique name of the tablet.
	Alias *TabletAlias `protobuf:"bytes,1,opt,name=alias" json:"alias,omitempty"`
	// Fully qualified domain name of the host.
	Hostname string `protobuf:"bytes,2,opt,name=hostname" json:"hostname,omitempty"`
	// Map of named ports. Normally this should include vt and grpc.
	// Going forward, the mysql port will be stored in mysql_port
	// instead of here.
	// For accessing mysql port, use topoproto.MysqlPort to fetch, and
	// topoproto.SetMysqlPort to set. These wrappers will ensure
	// legacy behavior is supported.
	PortMap map[string]int32 `protobuf:"bytes,4,rep,name=port_map,json=portMap" json:"port_map,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	// Keyspace name.
	Keyspace string `protobuf:"bytes,5,opt,name=keyspace" json:"keyspace,omitempty"`
	// Shard name. If range based sharding is used, it should match
	// key_range.
	Shard string `protobuf:"bytes,6,opt,name=shard" json:"shard,omitempty"`
	// If range based sharding is used, range for the tablet's shard.
	KeyRange *KeyRange `protobuf:"bytes,7,opt,name=key_range,json=keyRange" json:"key_range,omitempty"`
	// type is the current type of the tablet.
	Type TabletType `protobuf:"varint,8,opt,name=type,enum=topodata.TabletType" json:"type,omitempty"`
	// It this is set, it is used as the database name instead of the
	// normal "vt_" + keyspace.
	DbNameOverride string `protobuf:"bytes,9,opt,name=db_name_override,json=dbNameOverride" json:"db_name_override,omitempty"`
	// tablet tags
	Tags map[string]string `protobuf:"bytes,10,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// MySQL hostname.
	MysqlHostname string `protobuf:"bytes,12,opt,name=mysql_hostname,json=mysqlHostname" json:"mysql_hostname,omitempty"`
	// MySQL port. Use topoproto.MysqlPort and topoproto.SetMysqlPort
	// to access this variable. The functions provide support
	// for legacy behavior.
	MysqlPort int32 `protobuf:"varint,13,opt,name=mysql_port,json=mysqlPort" json:"mysql_port,omitempty"`
}

func (m *Tablet) Reset()                    { *m = Tablet{} }
func (m *Tablet) String() string            { return proto.CompactTextString(m) }
func (*Tablet) ProtoMessage()               {}
func (*Tablet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Tablet) GetAlias() *TabletAlias {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *Tablet) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Tablet) GetPortMap() map[string]int32 {
	if m != nil {
		return m.PortMap
	}
	return nil
}

func (m *Tablet) GetKeyspace() string {
	if m != nil {
		return m.Keyspace
	}
	return ""
}

func (m *Tablet) GetShard() string {
	if m != nil {
		return m.Shard
	}
	return ""
}

func (m *Tablet) GetKeyRange() *KeyRange {
	if m != nil {
		return m.KeyRange
	}
	return nil
}

func (m *Tablet) GetType() TabletType {
	if m != nil {
		return m.Type
	}
	return TabletType_UNKNOWN
}

func (m *Tablet) GetDbNameOverride() string {
	if m != nil {
		return m.DbNameOverride
	}
	return ""
}

func (m *Tablet) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Tablet) GetMysqlHostname() string {
	if m != nil {
		return m.MysqlHostname
	}
	return ""
}

func (m *Tablet) GetMysqlPort() int32 {
	if m != nil {
		return m.MysqlPort
	}
	return 0
}

// A Shard contains data about a subset of the data whithin a keyspace.
type Shard struct {
	// No lock is necessary to update this field, when for instance
	// TabletExternallyReparented updates this. However, we lock the
	// shard for reparenting operations (InitShardMaster,
	// PlannedReparentShard,EmergencyReparentShard), to guarantee
	// exclusive operation.
	MasterAlias *TabletAlias `protobuf:"bytes,1,opt,name=master_alias,json=masterAlias" json:"master_alias,omitempty"`
	// key_range is the KeyRange for this shard. It can be unset if:
	// - we are not using range-based sharding in this shard.
	// - the shard covers the entire keyrange.
	// This must match the shard name based on our other conventions, but
	// helpful to have it decomposed here.
	// Once set at creation time, it is never changed.
	KeyRange *KeyRange `protobuf:"bytes,2,opt,name=key_range,json=keyRange" json:"key_range,omitempty"`
	// served_types has at most one entry per TabletType
	// The keyspace lock is always taken when changing this.
	ServedTypes []*Shard_ServedType `protobuf:"bytes,3,rep,name=served_types,json=servedTypes" json:"served_types,omitempty"`
	// SourceShards is the list of shards we're replicating from,
	// using filtered replication.
	// The keyspace lock is always taken when changing this.
	SourceShards []*Shard_SourceShard `protobuf:"bytes,4,rep,name=source_shards,json=sourceShards" json:"source_shards,omitempty"`
	// Cells is the list of cells that contain tablets for this shard.
	// No lock is necessary to update this field.
	Cells []string `protobuf:"bytes,5,rep,name=cells" json:"cells,omitempty"`
	// tablet_controls has at most one entry per TabletType.
	// The keyspace lock is always taken when changing this.
	TabletControls []*Shard_TabletControl `protobuf:"bytes,6,rep,name=tablet_controls,json=tabletControls" json:"tablet_controls,omitempty"`
}

func (m *Shard) Reset()                    { *m = Shard{} }
func (m *Shard) String() string            { return proto.CompactTextString(m) }
func (*Shard) ProtoMessage()               {}
func (*Shard) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Shard) GetMasterAlias() *TabletAlias {
	if m != nil {
		return m.MasterAlias
	}
	return nil
}

func (m *Shard) GetKeyRange() *KeyRange {
	if m != nil {
		return m.KeyRange
	}
	return nil
}

func (m *Shard) GetServedTypes() []*Shard_ServedType {
	if m != nil {
		return m.ServedTypes
	}
	return nil
}

func (m *Shard) GetSourceShards() []*Shard_SourceShard {
	if m != nil {
		return m.SourceShards
	}
	return nil
}

func (m *Shard) GetCells() []string {
	if m != nil {
		return m.Cells
	}
	return nil
}

func (m *Shard) GetTabletControls() []*Shard_TabletControl {
	if m != nil {
		return m.TabletControls
	}
	return nil
}

// ServedType is an entry in the served_types
type Shard_ServedType struct {
	TabletType TabletType `protobuf:"varint,1,opt,name=tablet_type,json=tabletType,enum=topodata.TabletType" json:"tablet_type,omitempty"`
	Cells      []string   `protobuf:"bytes,2,rep,name=cells" json:"cells,omitempty"`
}

func (m *Shard_ServedType) Reset()                    { *m = Shard_ServedType{} }
func (m *Shard_ServedType) String() string            { return proto.CompactTextString(m) }
func (*Shard_ServedType) ProtoMessage()               {}
func (*Shard_ServedType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

func (m *Shard_ServedType) GetTabletType() TabletType {
	if m != nil {
		return m.TabletType
	}
	return TabletType_UNKNOWN
}

func (m *Shard_ServedType) GetCells() []string {
	if m != nil {
		return m.Cells
	}
	return nil
}

// SourceShard represents a data source for filtered replication
// accross shards. When this is used in a destination shard, the master
// of that shard will run filtered replication.
type Shard_SourceShard struct {
	// Uid is the unique ID for this SourceShard object.
	Uid uint32 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	// the source keyspace
	Keyspace string `protobuf:"bytes,2,opt,name=keyspace" json:"keyspace,omitempty"`
	// the source shard
	Shard string `protobuf:"bytes,3,opt,name=shard" json:"shard,omitempty"`
	// the source shard keyrange
	KeyRange *KeyRange `protobuf:"bytes,4,opt,name=key_range,json=keyRange" json:"key_range,omitempty"`
	// the source table list to replicate
	Tables []string `protobuf:"bytes,5,rep,name=tables" json:"tables,omitempty"`
}

func (m *Shard_SourceShard) Reset()                    { *m = Shard_SourceShard{} }
func (m *Shard_SourceShard) String() string            { return proto.CompactTextString(m) }
func (*Shard_SourceShard) ProtoMessage()               {}
func (*Shard_SourceShard) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 1} }

func (m *Shard_SourceShard) GetUid() uint32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *Shard_SourceShard) GetKeyspace() string {
	if m != nil {
		return m.Keyspace
	}
	return ""
}

func (m *Shard_SourceShard) GetShard() string {
	if m != nil {
		return m.Shard
	}
	return ""
}

func (m *Shard_SourceShard) GetKeyRange() *KeyRange {
	if m != nil {
		return m.KeyRange
	}
	return nil
}

func (m *Shard_SourceShard) GetTables() []string {
	if m != nil {
		return m.Tables
	}
	return nil
}

// TabletControl controls tablet's behavior
type Shard_TabletControl struct {
	// which tablet type is affected
	TabletType TabletType `protobuf:"varint,1,opt,name=tablet_type,json=tabletType,enum=topodata.TabletType" json:"tablet_type,omitempty"`
	Cells      []string   `protobuf:"bytes,2,rep,name=cells" json:"cells,omitempty"`
	// what to do
	DisableQueryService bool     `protobuf:"varint,3,opt,name=disable_query_service,json=disableQueryService" json:"disable_query_service,omitempty"`
	BlacklistedTables   []string `protobuf:"bytes,4,rep,name=blacklisted_tables,json=blacklistedTables" json:"blacklisted_tables,omitempty"`
}

func (m *Shard_TabletControl) Reset()                    { *m = Shard_TabletControl{} }
func (m *Shard_TabletControl) String() string            { return proto.CompactTextString(m) }
func (*Shard_TabletControl) ProtoMessage()               {}
func (*Shard_TabletControl) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 2} }

func (m *Shard_TabletControl) GetTabletType() TabletType {
	if m != nil {
		return m.TabletType
	}
	return TabletType_UNKNOWN
}

func (m *Shard_TabletControl) GetCells() []string {
	if m != nil {
		return m.Cells
	}
	return nil
}

func (m *Shard_TabletControl) GetDisableQueryService() bool {
	if m != nil {
		return m.DisableQueryService
	}
	return false
}

func (m *Shard_TabletControl) GetBlacklistedTables() []string {
	if m != nil {
		return m.BlacklistedTables
	}
	return nil
}

// A Keyspace contains data about a keyspace.
type Keyspace struct {
	// name of the column used for sharding
	// empty if the keyspace is not sharded
	ShardingColumnName string `protobuf:"bytes,1,opt,name=sharding_column_name,json=shardingColumnName" json:"sharding_column_name,omitempty"`
	// type of the column used for sharding
	// UNSET if the keyspace is not sharded
	ShardingColumnType KeyspaceIdType `protobuf:"varint,2,opt,name=sharding_column_type,json=shardingColumnType,enum=topodata.KeyspaceIdType" json:"sharding_column_type,omitempty"`
	// ServedFrom will redirect the appropriate traffic to
	// another keyspace.
	ServedFroms []*Keyspace_ServedFrom `protobuf:"bytes,4,rep,name=served_froms,json=servedFroms" json:"served_froms,omitempty"`
}

func (m *Keyspace) Reset()                    { *m = Keyspace{} }
func (m *Keyspace) String() string            { return proto.CompactTextString(m) }
func (*Keyspace) ProtoMessage()               {}
func (*Keyspace) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Keyspace) GetShardingColumnName() string {
	if m != nil {
		return m.ShardingColumnName
	}
	return ""
}

func (m *Keyspace) GetShardingColumnType() KeyspaceIdType {
	if m != nil {
		return m.ShardingColumnType
	}
	return KeyspaceIdType_UNSET
}

func (m *Keyspace) GetServedFroms() []*Keyspace_ServedFrom {
	if m != nil {
		return m.ServedFroms
	}
	return nil
}

// ServedFrom indicates a relationship between a TabletType and the
// keyspace name that's serving it.
type Keyspace_ServedFrom struct {
	// the tablet type (key for the map)
	TabletType TabletType `protobuf:"varint,1,opt,name=tablet_type,json=tabletType,enum=topodata.TabletType" json:"tablet_type,omitempty"`
	// the cells to limit this to
	Cells []string `protobuf:"bytes,2,rep,name=cells" json:"cells,omitempty"`
	// the keyspace name that's serving it
	Keyspace string `protobuf:"bytes,3,opt,name=keyspace" json:"keyspace,omitempty"`
}

func (m *Keyspace_ServedFrom) Reset()                    { *m = Keyspace_ServedFrom{} }
func (m *Keyspace_ServedFrom) String() string            { return proto.CompactTextString(m) }
func (*Keyspace_ServedFrom) ProtoMessage()               {}
func (*Keyspace_ServedFrom) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

func (m *Keyspace_ServedFrom) GetTabletType() TabletType {
	if m != nil {
		return m.TabletType
	}
	return TabletType_UNKNOWN
}

func (m *Keyspace_ServedFrom) GetCells() []string {
	if m != nil {
		return m.Cells
	}
	return nil
}

func (m *Keyspace_ServedFrom) GetKeyspace() string {
	if m != nil {
		return m.Keyspace
	}
	return ""
}

// ShardReplication describes the MySQL replication relationships
// whithin a cell.
type ShardReplication struct {
	// Note there can be only one Node in this array
	// for a given tablet.
	Nodes []*ShardReplication_Node `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *ShardReplication) Reset()                    { *m = ShardReplication{} }
func (m *ShardReplication) String() string            { return proto.CompactTextString(m) }
func (*ShardReplication) ProtoMessage()               {}
func (*ShardReplication) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ShardReplication) GetNodes() []*ShardReplication_Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

// Node describes a tablet instance within the cell
type ShardReplication_Node struct {
	TabletAlias *TabletAlias `protobuf:"bytes,1,opt,name=tablet_alias,json=tabletAlias" json:"tablet_alias,omitempty"`
}

func (m *ShardReplication_Node) Reset()                    { *m = ShardReplication_Node{} }
func (m *ShardReplication_Node) String() string            { return proto.CompactTextString(m) }
func (*ShardReplication_Node) ProtoMessage()               {}
func (*ShardReplication_Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *ShardReplication_Node) GetTabletAlias() *TabletAlias {
	if m != nil {
		return m.TabletAlias
	}
	return nil
}

// ShardReference is used as a pointer from a SrvKeyspace to a Shard
type ShardReference struct {
	// Copied from Shard.
	Name     string    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	KeyRange *KeyRange `protobuf:"bytes,2,opt,name=key_range,json=keyRange" json:"key_range,omitempty"`
}

func (m *ShardReference) Reset()                    { *m = ShardReference{} }
func (m *ShardReference) String() string            { return proto.CompactTextString(m) }
func (*ShardReference) ProtoMessage()               {}
func (*ShardReference) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ShardReference) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ShardReference) GetKeyRange() *KeyRange {
	if m != nil {
		return m.KeyRange
	}
	return nil
}

// SrvKeyspace is a rollup node for the keyspace itself.
type SrvKeyspace struct {
	// The partitions this keyspace is serving, per tablet type.
	Partitions []*SrvKeyspace_KeyspacePartition `protobuf:"bytes,1,rep,name=partitions" json:"partitions,omitempty"`
	// copied from Keyspace
	ShardingColumnName string                    `protobuf:"bytes,2,opt,name=sharding_column_name,json=shardingColumnName" json:"sharding_column_name,omitempty"`
	ShardingColumnType KeyspaceIdType            `protobuf:"varint,3,opt,name=sharding_column_type,json=shardingColumnType,enum=topodata.KeyspaceIdType" json:"sharding_column_type,omitempty"`
	ServedFrom         []*SrvKeyspace_ServedFrom `protobuf:"bytes,4,rep,name=served_from,json=servedFrom" json:"served_from,omitempty"`
}

func (m *SrvKeyspace) Reset()                    { *m = SrvKeyspace{} }
func (m *SrvKeyspace) String() string            { return proto.CompactTextString(m) }
func (*SrvKeyspace) ProtoMessage()               {}
func (*SrvKeyspace) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SrvKeyspace) GetPartitions() []*SrvKeyspace_KeyspacePartition {
	if m != nil {
		return m.Partitions
	}
	return nil
}

func (m *SrvKeyspace) GetShardingColumnName() string {
	if m != nil {
		return m.ShardingColumnName
	}
	return ""
}

func (m *SrvKeyspace) GetShardingColumnType() KeyspaceIdType {
	if m != nil {
		return m.ShardingColumnType
	}
	return KeyspaceIdType_UNSET
}

func (m *SrvKeyspace) GetServedFrom() []*SrvKeyspace_ServedFrom {
	if m != nil {
		return m.ServedFrom
	}
	return nil
}

type SrvKeyspace_KeyspacePartition struct {
	// The type this partition applies to.
	ServedType TabletType `protobuf:"varint,1,opt,name=served_type,json=servedType,enum=topodata.TabletType" json:"served_type,omitempty"`
	// List of non-overlapping continuous shards sorted by range.
	ShardReferences []*ShardReference `protobuf:"bytes,2,rep,name=shard_references,json=shardReferences" json:"shard_references,omitempty"`
}

func (m *SrvKeyspace_KeyspacePartition) Reset()         { *m = SrvKeyspace_KeyspacePartition{} }
func (m *SrvKeyspace_KeyspacePartition) String() string { return proto.CompactTextString(m) }
func (*SrvKeyspace_KeyspacePartition) ProtoMessage()    {}
func (*SrvKeyspace_KeyspacePartition) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{7, 0}
}

func (m *SrvKeyspace_KeyspacePartition) GetServedType() TabletType {
	if m != nil {
		return m.ServedType
	}
	return TabletType_UNKNOWN
}

func (m *SrvKeyspace_KeyspacePartition) GetShardReferences() []*ShardReference {
	if m != nil {
		return m.ShardReferences
	}
	return nil
}

// ServedFrom indicates a relationship between a TabletType and the
// keyspace name that's serving it.
type SrvKeyspace_ServedFrom struct {
	// the tablet type
	TabletType TabletType `protobuf:"varint,1,opt,name=tablet_type,json=tabletType,enum=topodata.TabletType" json:"tablet_type,omitempty"`
	// the keyspace name that's serving it
	Keyspace string `protobuf:"bytes,2,opt,name=keyspace" json:"keyspace,omitempty"`
}

func (m *SrvKeyspace_ServedFrom) Reset()                    { *m = SrvKeyspace_ServedFrom{} }
func (m *SrvKeyspace_ServedFrom) String() string            { return proto.CompactTextString(m) }
func (*SrvKeyspace_ServedFrom) ProtoMessage()               {}
func (*SrvKeyspace_ServedFrom) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 1} }

func (m *SrvKeyspace_ServedFrom) GetTabletType() TabletType {
	if m != nil {
		return m.TabletType
	}
	return TabletType_UNKNOWN
}

func (m *SrvKeyspace_ServedFrom) GetKeyspace() string {
	if m != nil {
		return m.Keyspace
	}
	return ""
}

// CellInfo contains information about a cell. CellInfo objects are
// stored in the global topology server, and describe how to reach
// local topology servers.
type CellInfo struct {
	// ServerAddress contains the address of the server for the cell.
	// The syntax of this field is topology implementation specific.
	// For instance, for Zookeeper, it is a comma-separated list of
	// server addresses.
	ServerAddress string `protobuf:"bytes,1,opt,name=server_address,json=serverAddress" json:"server_address,omitempty"`
	// Root is the path to store data in. It is only used when talking
	// to server_address.
	Root string `protobuf:"bytes,2,opt,name=root" json:"root,omitempty"`
	// Region is a group this cell belongs to. Used by vtgate to route traffic to
	// other cells (in same region) when there is no available tablet in the current cell.
	Region string `protobuf:"bytes,3,opt,name=region" json:"region,omitempty"`
}

func (m *CellInfo) Reset()                    { *m = CellInfo{} }
func (m *CellInfo) String() string            { return proto.CompactTextString(m) }
func (*CellInfo) ProtoMessage()               {}
func (*CellInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CellInfo) GetServerAddress() string {
	if m != nil {
		return m.ServerAddress
	}
	return ""
}

func (m *CellInfo) GetRoot() string {
	if m != nil {
		return m.Root
	}
	return ""
}

func (m *CellInfo) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func init() {
	proto.RegisterType((*KeyRange)(nil), "topodata.KeyRange")
	proto.RegisterType((*TabletAlias)(nil), "topodata.TabletAlias")
	proto.RegisterType((*Tablet)(nil), "topodata.Tablet")
	proto.RegisterType((*Shard)(nil), "topodata.Shard")
	proto.RegisterType((*Shard_ServedType)(nil), "topodata.Shard.ServedType")
	proto.RegisterType((*Shard_SourceShard)(nil), "topodata.Shard.SourceShard")
	proto.RegisterType((*Shard_TabletControl)(nil), "topodata.Shard.TabletControl")
	proto.RegisterType((*Keyspace)(nil), "topodata.Keyspace")
	proto.RegisterType((*Keyspace_ServedFrom)(nil), "topodata.Keyspace.ServedFrom")
	proto.RegisterType((*ShardReplication)(nil), "topodata.ShardReplication")
	proto.RegisterType((*ShardReplication_Node)(nil), "topodata.ShardReplication.Node")
	proto.RegisterType((*ShardReference)(nil), "topodata.ShardReference")
	proto.RegisterType((*SrvKeyspace)(nil), "topodata.SrvKeyspace")
	proto.RegisterType((*SrvKeyspace_KeyspacePartition)(nil), "topodata.SrvKeyspace.KeyspacePartition")
	proto.RegisterType((*SrvKeyspace_ServedFrom)(nil), "topodata.SrvKeyspace.ServedFrom")
	proto.RegisterType((*CellInfo)(nil), "topodata.CellInfo")
	proto.RegisterEnum("topodata.KeyspaceIdType", KeyspaceIdType_name, KeyspaceIdType_value)
	proto.RegisterEnum("topodata.TabletType", TabletType_name, TabletType_value)
}

func init() { proto.RegisterFile("topodata.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdf, 0x6f, 0xe2, 0xc6,
	0x13, 0xff, 0x1a, 0x0c, 0x81, 0x31, 0x70, 0xce, 0x7e, 0x73, 0x95, 0xe5, 0xea, 0x54, 0x84, 0x54,
	0x15, 0x5d, 0x55, 0x5a, 0x71, 0xbd, 0x36, 0x3a, 0xa9, 0x52, 0x08, 0xe1, 0x7a, 0xe4, 0x07, 0xa1,
	0x0b, 0x51, 0x9b, 0x87, 0xca, 0x72, 0xf0, 0x26, 0x67, 0xc5, 0x78, 0xb9, 0xdd, 0x25, 0x12, 0x7f,
	0xc3, 0x3d, 0xf4, 0x9e, 0xfb, 0x9f, 0xf4, 0xa9, 0x8f, 0xfd, 0xb7, 0xaa, 0xdd, 0xb5, 0xc1, 0x90,
	0x26, 0xcd, 0x55, 0x79, 0xca, 0xce, 0xee, 0xcc, 0x78, 0x3e, 0x9f, 0xf9, 0xcc, 0x04, 0xa8, 0x09,
	0x3a, 0xa3, 0x81, 0x2f, 0xfc, 0xd6, 0x8c, 0x51, 0x41, 0x51, 0x29, 0xb5, 0x1b, 0x6d, 0x28, 0x1d,
	0x91, 0x05, 0xf6, 0xe3, 0x2b, 0x82, 0x76, 0xa0, 0xc0, 0x85, 0xcf, 0x84, 0x63, 0xd4, 0x8d, 0x66,
	0x05, 0x6b, 0x03, 0xd9, 0x90, 0x27, 0x71, 0xe0, 0xe4, 0xd4, 0x9d, 0x3c, 0x36, 0x5e, 0x80, 0x35,
	0xf6, 0x2f, 0x22, 0x22, 0x3a, 0x51, 0xe8, 0x73, 0x84, 0xc0, 0x9c, 0x90, 0x28, 0x52, 0x51, 0x65,
	0xac, 0xce, 0x32, 0x68, 0x1e, 0xea, 0xa0, 0x2a, 0x96, 0xc7, 0xc6, 0x1f, 0x26, 0x14, 0x75, 0x14,
	0xfa, 0x12, 0x0a, 0xbe, 0x8c, 0x54, 0x11, 0x56, 0xfb, 0x69, 0x6b, 0x59, 0x5d, 0x26, 0x2d, 0xd6,
	0x3e, 0xc8, 0x85, 0xd2, 0x5b, 0xca, 0x45, 0xec, 0x4f, 0x89, 0x4a, 0x57, 0xc6, 0x4b, 0x1b, 0xed,
	0x42, 0x69, 0x46, 0x99, 0xf0, 0xa6, 0xfe, 0xcc, 0x31, 0xeb, 0xf9, 0xa6, 0xd5, 0x7e, 0xb6, 0x99,
	0xab, 0x35, 0xa4, 0x4c, 0x9c, 0xf8, 0xb3, 0x5e, 0x2c, 0xd8, 0x02, 0x6f, 0xcd, 0xb4, 0x25, 0xb3,
	0x5e, 0x93, 0x05, 0x9f, 0xf9, 0x13, 0xe2, 0x14, 0x74, 0xd6, 0xd4, 0x56, 0x34, 0xbc, 0xf5, 0x59,
	0xe0, 0x14, 0xd5, 0x83, 0x36, 0xd0, 0xd7, 0x50, 0xbe, 0x26, 0x0b, 0x8f, 0x49, 0xa6, 0x9c, 0x2d,
	0x55, 0x38, 0x5a, 0x7d, 0x2c, 0xe5, 0x50, 0xa5, 0xd1, 0x6c, 0x36, 0xc1, 0x14, 0x8b, 0x19, 0x71,
	0x4a, 0x75, 0xa3, 0x59, 0x6b, 0xef, 0x6c, 0x16, 0x36, 0x5e, 0xcc, 0x08, 0x56, 0x1e, 0xa8, 0x09,
	0x76, 0x70, 0xe1, 0x49, 0x44, 0x1e, 0xbd, 0x21, 0x8c, 0x85, 0x01, 0x71, 0xca, 0xea, 0xdb, 0xb5,
	0xe0, 0x62, 0xe0, 0x4f, 0xc9, 0x69, 0x72, 0x8b, 0x5a, 0x60, 0x0a, 0xff, 0x8a, 0x3b, 0xa0, 0xc0,
	0xba, 0xb7, 0xc0, 0x8e, 0xfd, 0x2b, 0xae, 0x91, 0x2a, 0x3f, 0xf4, 0x39, 0xd4, 0xa6, 0x0b, 0xfe,
	0x2e, 0xf2, 0x96, 0x14, 0x56, 0x54, 0xde, 0xaa, 0xba, 0x7d, 0x93, 0xf2, 0xf8, 0x0c, 0x40, 0xbb,
	0x49, 0x7a, 0x9c, 0x6a, 0xdd, 0x68, 0x16, 0x70, 0x59, 0xdd, 0x48, 0xf6, 0xdc, 0x57, 0x50, 0xc9,
	0xb2, 0x28, 0x9b, 0x7b, 0x4d, 0x16, 0x49, 0xbf, 0xe5, 0x51, 0x52, 0x76, 0xe3, 0x47, 0x73, 0xdd,
	0xa1, 0x02, 0xd6, 0xc6, 0xab, 0xdc, 0xae, 0xe1, 0x7e, 0x0f, 0xe5, 0x65, 0x51, 0xff, 0x16, 0x58,
	0xce, 0x04, 0x1e, 0x9a, 0xa5, 0xbc, 0x6d, 0x1e, 0x9a, 0x25, 0xcb, 0xae, 0x34, 0xde, 0x17, 0xa1,
	0x30, 0x52, 0x5d, 0xd8, 0x85, 0xca, 0xd4, 0xe7, 0x82, 0x30, 0xef, 0x01, 0x0a, 0xb2, 0xb4, 0xab,
	0x56, 0xe9, 0x5a, 0xff, 0x72, 0x0f, 0xe8, 0xdf, 0x0f, 0x50, 0xe1, 0x84, 0xdd, 0x90, 0xc0, 0x93,
	0x4d, 0xe2, 0x4e, 0x7e, 0x93, 0x73, 0x55, 0x51, 0x6b, 0xa4, 0x7c, 0x54, 0x37, 0x2d, 0xbe, 0x3c,
	0x73, 0xb4, 0x07, 0x55, 0x4e, 0xe7, 0x6c, 0x42, 0x3c, 0xa5, 0x1f, 0x9e, 0x08, 0xf4, 0xd3, 0x5b,
	0xf1, 0xca, 0x49, 0x9d, 0x71, 0x85, 0xaf, 0x0c, 0x2e, 0xb9, 0x91, 0xb3, 0xc4, 0x9d, 0x42, 0x3d,
	0x2f, 0xb9, 0x51, 0x06, 0x7a, 0x0d, 0x4f, 0x84, 0xc2, 0xe8, 0x4d, 0x68, 0x2c, 0x18, 0x8d, 0xb8,
	0x53, 0xdc, 0x94, 0xbe, 0xce, 0xac, 0xa9, 0xe8, 0x6a, 0x2f, 0x5c, 0x13, 0x59, 0x93, 0xbb, 0xe7,
	0x00, 0xab, 0xd2, 0xd1, 0x4b, 0xb0, 0x92, 0xac, 0x4a, 0xb3, 0xc6, 0x3d, 0x9a, 0x05, 0xb1, 0x3c,
	0xaf, 0x4a, 0xcc, 0x65, 0x4a, 0x74, 0x7f, 0x37, 0xc0, 0xca, 0xc0, 0x4a, 0x97, 0x81, 0xb1, 0x5c,
	0x06, 0x6b, 0xe3, 0x97, 0xbb, 0x6b, 0xfc, 0xf2, 0x77, 0x8e, 0x9f, 0xf9, 0x80, 0xf6, 0x7d, 0x02,
	0x45, 0x55, 0x68, 0x4a, 0x5f, 0x62, 0xb9, 0x7f, 0x1a, 0x50, 0x5d, 0x63, 0xe6, 0x51, 0xb1, 0xa3,
	0x36, 0x3c, 0x0d, 0x42, 0x2e, 0xbd, 0xbc, 0x77, 0x73, 0xc2, 0x16, 0x9e, 0xd4, 0x44, 0x38, 0x21,
	0x0a, 0x4d, 0x09, 0xff, 0x3f, 0x79, 0xfc, 0x49, 0xbe, 0x8d, 0xf4, 0x13, 0xfa, 0x0a, 0xd0, 0x45,
	0xe4, 0x4f, 0xae, 0xa3, 0x90, 0x0b, 0x29, 0x37, 0x5d, 0xb6, 0xa9, 0xd2, 0x6e, 0x67, 0x5e, 0x54,
	0x21, 0xbc, 0xf1, 0x57, 0x4e, 0xed, 0x6c, 0xcd, 0xd6, 0x37, 0xb0, 0xa3, 0x08, 0x0a, 0xe3, 0x2b,
	0x6f, 0x42, 0xa3, 0xf9, 0x34, 0x56, 0x8b, 0x24, 0x99, 0x31, 0x94, 0xbe, 0x75, 0xd5, 0x93, 0xdc,
	0x25, 0xe8, 0xf0, 0x76, 0x84, 0xc2, 0x9d, 0x53, 0xb8, 0x9d, 0x35, 0x52, 0xd5, 0x37, 0xfa, 0x5a,
	0xdd, 0x1b, 0xb9, 0x14, 0x07, 0x7b, 0xcb, 0x19, 0xb9, 0x64, 0x74, 0xca, 0x6f, 0x2f, 0xe1, 0x34,
	0x47, 0x32, 0x26, 0xaf, 0x19, 0x9d, 0xa6, 0x63, 0x22, 0xcf, 0xdc, 0x9d, 0xa7, 0x32, 0x94, 0xe6,
	0xe3, 0xb6, 0x22, 0x2b, 0xb2, 0xfc, 0xba, 0xc8, 0xf4, 0x76, 0x69, 0xbc, 0x37, 0xc0, 0xd6, 0x93,
	0x47, 0x66, 0x51, 0x38, 0xf1, 0x45, 0x48, 0x63, 0xf4, 0x12, 0x0a, 0x31, 0x0d, 0x88, 0xdc, 0x2d,
	0x12, 0xcc, 0x67, 0x1b, 0x63, 0x95, 0x71, 0x6d, 0x0d, 0x68, 0x40, 0xb0, 0xf6, 0x76, 0xf7, 0xc0,
	0x94, 0xa6, 0xdc, 0x50, 0x09, 0x84, 0x87, 0x6c, 0x28, 0xb1, 0x32, 0x1a, 0x67, 0x50, 0x4b, 0xbe,
	0x70, 0x49, 0x18, 0x89, 0x27, 0x44, 0xfe, 0x67, 0xcd, 0x34, 0x53, 0x9d, 0x3f, 0x7a, 0x8f, 0x35,
	0x3e, 0x98, 0x60, 0x8d, 0xd8, 0xcd, 0x52, 0x31, 0x3f, 0x02, 0xcc, 0x7c, 0x26, 0x42, 0x89, 0x20,
	0x05, 0xf9, 0x45, 0x06, 0xe4, 0xca, 0x75, 0xd9, 0xbd, 0x61, 0xea, 0x8f, 0x33, 0xa1, 0x77, 0x4a,
	0x2f, 0xf7, 0xd1, 0xd2, 0xcb, 0xff, 0x07, 0xe9, 0x75, 0xc0, 0xca, 0x48, 0x2f, 0x51, 0x5e, 0xfd,
	0x9f, 0x71, 0x64, 0xc4, 0x07, 0x2b, 0xf1, 0xb9, 0xbf, 0x19, 0xb0, 0x7d, 0x0b, 0xa2, 0xd4, 0x60,
	0x66, 0xef, 0xdf, 0xaf, 0xc1, 0xd5, 0xc2, 0x47, 0x5d, 0xb0, 0x55, 0x95, 0x1e, 0x4b, 0xdb, 0xa7,
	0xe5, 0x68, 0x65, 0x71, 0xad, 0xf7, 0x17, 0x3f, 0xe1, 0x6b, 0x36, 0x77, 0xbd, 0xc7, 0x98, 0x86,
	0x7b, 0x96, 0xeb, 0xa1, 0x59, 0x2a, 0xd8, 0xc5, 0xc6, 0xaf, 0x50, 0xea, 0x92, 0x28, 0xea, 0xc7,
	0x97, 0x54, 0xfe, 0x44, 0x50, 0x28, 0x98, 0xe7, 0x07, 0x01, 0x23, 0x9c, 0x27, 0x6a, 0xab, 0xea,
	0xdb, 0x8e, 0xbe, 0x94, 0x52, 0x64, 0x94, 0x8a, 0x24, 0xa1, 0x3a, 0xcb, 0x15, 0xcb, 0xc8, 0x55,
	0x48, 0xe3, 0x64, 0xbc, 0x12, 0xeb, 0x79, 0x1b, 0x6a, 0xeb, 0x0d, 0x44, 0x65, 0x28, 0x9c, 0x0d,
	0x46, 0xbd, 0xb1, 0xfd, 0x3f, 0x04, 0x50, 0x3c, 0xeb, 0x0f, 0xc6, 0xdf, 0x7d, 0x6b, 0x1b, 0xf2,
	0x7a, 0xff, 0x7c, 0xdc, 0x1b, 0xd9, 0xb9, 0xe7, 0x1f, 0x0c, 0x80, 0x15, 0x1e, 0x64, 0xc1, 0xd6,
	0xd9, 0xe0, 0x68, 0x70, 0xfa, 0xf3, 0x40, 0x87, 0x9c, 0x74, 0x46, 0xe3, 0x1e, 0xb6, 0x0d, 0xf9,
	0x80, 0x7b, 0xc3, 0xe3, 0x7e, 0xb7, 0x63, 0xe7, 0xe4, 0x03, 0x3e, 0x38, 0x1d, 0x1c, 0x9f, 0xdb,
	0x79, 0x95, 0xab, 0x33, 0xee, 0xbe, 0xd1, 0xc7, 0xd1, 0xb0, 0x83, 0x7b, 0xb6, 0x89, 0x6c, 0xa8,
	0xf4, 0x7e, 0x19, 0xf6, 0x70, 0xff, 0xa4, 0x37, 0x18, 0x77, 0x8e, 0xed, 0x82, 0x8c, 0xd9, 0xef,
	0x74, 0x8f, 0xce, 0x86, 0x76, 0x51, 0x27, 0x1b, 0x8d, 0x4f, 0x71, 0xcf, 0xde, 0x92, 0xc6, 0x01,
	0xee, 0xf4, 0x07, 0xbd, 0x03, 0xbb, 0xe4, 0xe6, 0x6c, 0x63, 0x7f, 0x1b, 0x9e, 0x84, 0xb4, 0x75,
	0x13, 0x0a, 0xc2, 0xb9, 0xfe, 0xdd, 0x7c, 0x51, 0x54, 0x7f, 0x5e, 0xfc, 0x1d, 0x00, 0x00, 0xff,
	0xff, 0xc7, 0x53, 0xa5, 0x90, 0x50, 0x0b, 0x00, 0x00,
}
