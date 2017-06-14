package model

import (
	"fmt"
)

// ResourceKey is the reference of a specific game resource.
type ResourceKey struct {
	Type  ResourceType
	Index uint16
}

// MakeResourceKey returns a combined resource identifier.
func MakeResourceKey(resourceType ResourceType, index uint16) ResourceKey {
	return ResourceKey{resourceType, index}
}

// ResourceKeyFromInt returns a resource identifier wrapping the provided integer.
func ResourceKeyFromInt(value int) ResourceKey {
	return ResourceKey{
		Type:  ResourceType(uint16((value >> 16) & 0xFFFF)),
		Index: uint16(value & 0xFFFF)}
}

// ToInt returns a single integer representation of the ID.
func (id ResourceKey) ToInt() int {
	return (int(id.Type) << 16) | int(id.Index)
}

// String implements the Stringer interface.
func (id ResourceKey) String() string {
	return fmt.Sprintf("0x%04X:%03d", uint16(id.Type), id.Index)
}
