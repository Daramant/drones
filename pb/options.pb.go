// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: drones/options.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Maze     *Options_Maze  `protobuf:"bytes,1,opt,name=maze,proto3" json:"maze,omitempty"`
	CellSize uint32         `protobuf:"varint,2,opt,name=cellSize,proto3" json:"cellSize,omitempty"`
	Drone    *Options_Drone `protobuf:"bytes,3,opt,name=drone,proto3" json:"drone,omitempty"`
	MaxTicks uint32         `protobuf:"varint,4,opt,name=maxTicks,proto3" json:"maxTicks,omitempty"`
}

func (x *Options) Reset() {
	*x = Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drones_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options) ProtoMessage() {}

func (x *Options) ProtoReflect() protoreflect.Message {
	mi := &file_drones_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options.ProtoReflect.Descriptor instead.
func (*Options) Descriptor() ([]byte, []int) {
	return file_drones_options_proto_rawDescGZIP(), []int{0}
}

func (x *Options) GetMaze() *Options_Maze {
	if x != nil {
		return x.Maze
	}
	return nil
}

func (x *Options) GetCellSize() uint32 {
	if x != nil {
		return x.CellSize
	}
	return 0
}

func (x *Options) GetDrone() *Options_Drone {
	if x != nil {
		return x.Drone
	}
	return nil
}

func (x *Options) GetMaxTicks() uint32 {
	if x != nil {
		return x.MaxTicks
	}
	return 0
}

type Options_CellPos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X uint32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y uint32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Options_CellPos) Reset() {
	*x = Options_CellPos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drones_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options_CellPos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options_CellPos) ProtoMessage() {}

func (x *Options_CellPos) ProtoReflect() protoreflect.Message {
	mi := &file_drones_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options_CellPos.ProtoReflect.Descriptor instead.
func (*Options_CellPos) Descriptor() ([]byte, []int) {
	return file_drones_options_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Options_CellPos) GetX() uint32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Options_CellPos) GetY() uint32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type Options_Maze struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Width  uint32           `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height uint32           `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Walls  []byte           `protobuf:"bytes,3,opt,name=walls,proto3" json:"walls,omitempty"` // Bit array of walls (1 - Wall, 0 - Path)
	Goal   *Options_CellPos `protobuf:"bytes,4,opt,name=goal,proto3" json:"goal,omitempty"`
}

func (x *Options_Maze) Reset() {
	*x = Options_Maze{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drones_options_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options_Maze) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options_Maze) ProtoMessage() {}

func (x *Options_Maze) ProtoReflect() protoreflect.Message {
	mi := &file_drones_options_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options_Maze.ProtoReflect.Descriptor instead.
func (*Options_Maze) Descriptor() ([]byte, []int) {
	return file_drones_options_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Options_Maze) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Options_Maze) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Options_Maze) GetWalls() []byte {
	if x != nil {
		return x.Walls
	}
	return nil
}

func (x *Options_Maze) GetGoal() *Options_CellPos {
	if x != nil {
		return x.Goal
	}
	return nil
}

type Options_Drone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Width    float32 `protobuf:"fixed32,1,opt,name=width,proto3" json:"width,omitempty"`
	Height   float32 `protobuf:"fixed32,2,opt,name=height,proto3" json:"height,omitempty"`
	Weight   float32 `protobuf:"fixed32,3,opt,name=weight,proto3" json:"weight,omitempty"`
	MaxForce float32 `protobuf:"fixed32,4,opt,name=maxForce,proto3" json:"maxForce,omitempty"`
}

func (x *Options_Drone) Reset() {
	*x = Options_Drone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_drones_options_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Options_Drone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options_Drone) ProtoMessage() {}

func (x *Options_Drone) ProtoReflect() protoreflect.Message {
	mi := &file_drones_options_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options_Drone.ProtoReflect.Descriptor instead.
func (*Options_Drone) Descriptor() ([]byte, []int) {
	return file_drones_options_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Options_Drone) GetWidth() float32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Options_Drone) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Options_Drone) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Options_Drone) GetMaxForce() float32 {
	if x != nil {
		return x.MaxForce
	}
	return 0
}

var File_drones_options_proto protoreflect.FileDescriptor

var file_drones_options_proto_rawDesc = []byte{
	0x0a, 0x14, 0x64, 0x72, 0x6f, 0x6e, 0x65, 0x73, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x64, 0x72, 0x6f, 0x6e, 0x65, 0x73, 0x22, 0xa3,
	0x03, 0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x28, 0x0a, 0x04, 0x6d, 0x61,
	0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x72, 0x6f, 0x6e, 0x65,
	0x73, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x61, 0x7a, 0x65, 0x52, 0x04,
	0x6d, 0x61, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x65, 0x6c, 0x6c, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x65, 0x6c, 0x6c, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x2b, 0x0a, 0x05, 0x64, 0x72, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x64, 0x72, 0x6f, 0x6e, 0x65, 0x73, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x44, 0x72, 0x6f, 0x6e, 0x65, 0x52, 0x05, 0x64, 0x72, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x61, 0x78, 0x54, 0x69, 0x63, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x6d, 0x61, 0x78, 0x54, 0x69, 0x63, 0x6b, 0x73, 0x1a, 0x25, 0x0a, 0x07, 0x43, 0x65, 0x6c,
	0x6c, 0x50, 0x6f, 0x73, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x79,
	0x1a, 0x77, 0x0a, 0x04, 0x4d, 0x61, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x61, 0x6c, 0x6c, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x77, 0x61, 0x6c, 0x6c, 0x73, 0x12, 0x2b, 0x0a, 0x04,
	0x67, 0x6f, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x72, 0x6f,
	0x6e, 0x65, 0x73, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x65, 0x6c, 0x6c,
	0x50, 0x6f, 0x73, 0x52, 0x04, 0x67, 0x6f, 0x61, 0x6c, 0x1a, 0x69, 0x0a, 0x05, 0x44, 0x72, 0x6f,
	0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x46,
	0x6f, 0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x46,
	0x6f, 0x72, 0x63, 0x65, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_drones_options_proto_rawDescOnce sync.Once
	file_drones_options_proto_rawDescData = file_drones_options_proto_rawDesc
)

func file_drones_options_proto_rawDescGZIP() []byte {
	file_drones_options_proto_rawDescOnce.Do(func() {
		file_drones_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_drones_options_proto_rawDescData)
	})
	return file_drones_options_proto_rawDescData
}

var file_drones_options_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_drones_options_proto_goTypes = []interface{}{
	(*Options)(nil),         // 0: drones.Options
	(*Options_CellPos)(nil), // 1: drones.Options.CellPos
	(*Options_Maze)(nil),    // 2: drones.Options.Maze
	(*Options_Drone)(nil),   // 3: drones.Options.Drone
}
var file_drones_options_proto_depIdxs = []int32{
	2, // 0: drones.Options.maze:type_name -> drones.Options.Maze
	3, // 1: drones.Options.drone:type_name -> drones.Options.Drone
	1, // 2: drones.Options.Maze.goal:type_name -> drones.Options.CellPos
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_drones_options_proto_init() }
func file_drones_options_proto_init() {
	if File_drones_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_drones_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Options); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_drones_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Options_CellPos); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_drones_options_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Options_Maze); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_drones_options_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Options_Drone); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_drones_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_drones_options_proto_goTypes,
		DependencyIndexes: file_drones_options_proto_depIdxs,
		MessageInfos:      file_drones_options_proto_msgTypes,
	}.Build()
	File_drones_options_proto = out.File
	file_drones_options_proto_rawDesc = nil
	file_drones_options_proto_goTypes = nil
	file_drones_options_proto_depIdxs = nil
}
