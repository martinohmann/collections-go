package mapelement

import "io"

//go:generate collection-gen -i github.com/martinohmann/collections-go/examples/map-element -p github.com/martinohmann/collections-go/examples/map-element

// +collection-gen=true
// +collection-gen:options=underlying,out-name=reader_map
// +collection-gen:options=underlying,immutable,out-name=reader_map

type ReaderMap map[string]io.Reader

// +collection-gen=true
// +collection-gen:options=name=WriterMap,out-name=writer_map
// +collection-gen:options=immutable,name=ImmutableWriterMap,out-name=writer_map

type wm = map[string]io.Writer
