// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: product.proto

package protoProduct

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

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID     string          `protobuf:"bytes,1,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
	Name          string          `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Description   string          `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Brand         int32           `protobuf:"varint,4,opt,name=Brand,proto3" json:"Brand,omitempty"`
	Price         float32         `protobuf:"fixed32,5,opt,name=Price,proto3" json:"Price,omitempty"`
	Quantity      int32           `protobuf:"varint,6,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	CategoryID    string          `protobuf:"bytes,7,opt,name=CategoryID,proto3" json:"CategoryID,omitempty"`
	ProductImages []*ProductImage `protobuf:"bytes,8,rep,name=productImages,proto3" json:"productImages,omitempty"`
	Active        bool            `protobuf:"varint,9,opt,name=Active,proto3" json:"Active,omitempty"`
	CreatedAt     string          `protobuf:"bytes,10,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt     string          `protobuf:"bytes,11,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	DeletedAt     string          `protobuf:"bytes,12,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetBrand() int32 {
	if x != nil {
		return x.Brand
	}
	return 0
}

func (x *Product) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Product) GetCategoryID() string {
	if x != nil {
		return x.CategoryID
	}
	return ""
}

func (x *Product) GetProductImages() []*ProductImage {
	if x != nil {
		return x.ProductImages
	}
	return nil
}

func (x *Product) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *Product) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Product) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Product) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type ProductImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ProductID string `protobuf:"bytes,2,opt,name=productID,proto3" json:"productID,omitempty"`
}

func (x *ProductImage) Reset() {
	*x = ProductImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductImage) ProtoMessage() {}

func (x *ProductImage) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductImage.ProtoReflect.Descriptor instead.
func (*ProductImage) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{1}
}

func (x *ProductImage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductImage) GetProductID() string {
	if x != nil {
		return x.ProductID
	}
	return ""
}

var File_product_proto protoreflect.FileDescriptor

var file_product_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xec, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x42, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x51,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x51,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x12, 0x33, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x0d, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x40,
	0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44,
	0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_proto_rawDescOnce sync.Once
	file_product_proto_rawDescData = file_product_proto_rawDesc
)

func file_product_proto_rawDescGZIP() []byte {
	file_product_proto_rawDescOnce.Do(func() {
		file_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_proto_rawDescData)
	})
	return file_product_proto_rawDescData
}

var file_product_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_product_proto_goTypes = []interface{}{
	(*Product)(nil),      // 0: Product
	(*ProductImage)(nil), // 1: ProductImage
}
var file_product_proto_depIdxs = []int32{
	1, // 0: Product.productImages:type_name -> ProductImage
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_product_proto_init() }
func file_product_proto_init() {
	if File_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductImage); i {
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
			RawDescriptor: file_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_product_proto_goTypes,
		DependencyIndexes: file_product_proto_depIdxs,
		MessageInfos:      file_product_proto_msgTypes,
	}.Build()
	File_product_proto = out.File
	file_product_proto_rawDesc = nil
	file_product_proto_goTypes = nil
	file_product_proto_depIdxs = nil
}
