package sql

import (
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//A number of items are counted
	Count uint32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	//Number of items per page
	Limit uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	//Number of pages to show = round up (count / show)
	Pages uint32 `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	//Current_page and by default = 1
	Page uint32 `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
}



