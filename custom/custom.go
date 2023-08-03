package ionoscloud

import (
	ionoscloud "github.com/ionos-cloud/sdk-go-dbaas-mongo"
)

type Identifiable interface {
	GetId() *string
	GetIdOk() (*string, bool)
	SetId(v string)
	HasId() bool
}

type Stateful interface {
	GetState() *ionoscloud.State
	GetStateOk() (*ionoscloud.State, bool)
	SetState(ionoscloud.State)
}

type List[T Identifiable] interface {
	GetItems() *[]T
	GetItemsOk() (*[]T, bool)
	SetItems(v []T)
	HasItems() bool
}

type Paginated[T Identifiable] interface {
	GetOffset() *int32
	GetOffsetOk() (*int32, bool)
	SetOffset(int32)
	HasOffset() bool
	GetLimit() *int32
	GetLimitOk() (*int32, bool)
	SetLimit(int32)
	HasLimit() bool
}

type Request[T Identifiable] interface {
	Execute() (T, *ionoscloud.APIResponse, error)
}

type PaginatedRequest[T Identifiable] interface {
	Request[T]
	Limit(int32) PaginatedRequest[T]
	Offset(int32) PaginatedRequest[T]
	FilterName(string) PaginatedRequest[T]
}
