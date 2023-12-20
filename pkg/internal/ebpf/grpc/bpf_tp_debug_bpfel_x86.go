// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64
// +build 386 amd64

package grpc

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpf_tp_debugGoroutineMetadata struct {
	Parent    uint64
	Timestamp uint64
}

type bpf_tp_debugGrpcClientFuncInvocationT struct {
	StartMonotimeNs uint64
	Cc              uint64
	Method          uint64
	MethodLen       uint64
	Tp              bpf_tp_debugTpInfoT
	Flags           uint64
}

type bpf_tp_debugGrpcSrvFuncInvocationT struct {
	StartMonotimeNs uint64
	Stream          uint64
	Tp              bpf_tp_debugTpInfoT
}

type bpf_tp_debugPidKeyT struct {
	Pid       uint32
	Namespace uint32
}

type bpf_tp_debugTpInfoT struct {
	TraceId  [16]uint8
	SpanId   [8]uint8
	ParentId [8]uint8
	Epoch    uint64
	Flags    uint8
	_        [7]byte
}

// loadBpf_tp_debug returns the embedded CollectionSpec for bpf_tp_debug.
func loadBpf_tp_debug() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_Bpf_tp_debugBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf_tp_debug: %w", err)
	}

	return spec, err
}

// loadBpf_tp_debugObjects loads bpf_tp_debug and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpf_tp_debugObjects
//	*bpf_tp_debugPrograms
//	*bpf_tp_debugMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpf_tp_debugObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf_tp_debug()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpf_tp_debugSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_tp_debugSpecs struct {
	bpf_tp_debugProgramSpecs
	bpf_tp_debugMapSpecs
}

// bpf_tp_debugSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_tp_debugProgramSpecs struct {
	UprobeClientConnClose                       *ebpf.ProgramSpec `ebpf:"uprobe_ClientConn_Close"`
	UprobeClientConnInvoke                      *ebpf.ProgramSpec `ebpf:"uprobe_ClientConn_Invoke"`
	UprobeClientConnInvokeReturn                *ebpf.ProgramSpec `ebpf:"uprobe_ClientConn_Invoke_return"`
	UprobeClientConnNewStream                   *ebpf.ProgramSpec `ebpf:"uprobe_ClientConn_NewStream"`
	UprobeHpackEncoderWriteField                *ebpf.ProgramSpec `ebpf:"uprobe_hpack_Encoder_WriteField"`
	UprobeServerHandleStream                    *ebpf.ProgramSpec `ebpf:"uprobe_server_handleStream"`
	UprobeServerHandleStreamReturn              *ebpf.ProgramSpec `ebpf:"uprobe_server_handleStream_return"`
	UprobeTransportHttp2ClientNewStream         *ebpf.ProgramSpec `ebpf:"uprobe_transport_http2Client_NewStream"`
	UprobeTransportLoopyWriterWriteHeader       *ebpf.ProgramSpec `ebpf:"uprobe_transport_loopyWriter_writeHeader"`
	UprobeTransportLoopyWriterWriteHeaderReturn *ebpf.ProgramSpec `ebpf:"uprobe_transport_loopyWriter_writeHeader_return"`
	UprobeTransportWriteStatus                  *ebpf.ProgramSpec `ebpf:"uprobe_transport_writeStatus"`
}

// bpf_tp_debugMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_tp_debugMapSpecs struct {
	Events                    *ebpf.MapSpec `ebpf:"events"`
	GoTraceMap                *ebpf.MapSpec `ebpf:"go_trace_map"`
	GolangMapbucketStorageMap *ebpf.MapSpec `ebpf:"golang_mapbucket_storage_map"`
	OngoingGoroutines         *ebpf.MapSpec `ebpf:"ongoing_goroutines"`
	OngoingGrpcClientRequests *ebpf.MapSpec `ebpf:"ongoing_grpc_client_requests"`
	OngoingGrpcHeaderWrites   *ebpf.MapSpec `ebpf:"ongoing_grpc_header_writes"`
	OngoingGrpcRequestStatus  *ebpf.MapSpec `ebpf:"ongoing_grpc_request_status"`
	OngoingGrpcServerRequests *ebpf.MapSpec `ebpf:"ongoing_grpc_server_requests"`
	OngoingStreams            *ebpf.MapSpec `ebpf:"ongoing_streams"`
	PidCache                  *ebpf.MapSpec `ebpf:"pid_cache"`
	ValidPids                 *ebpf.MapSpec `ebpf:"valid_pids"`
}

// bpf_tp_debugObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpf_tp_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_tp_debugObjects struct {
	bpf_tp_debugPrograms
	bpf_tp_debugMaps
}

func (o *bpf_tp_debugObjects) Close() error {
	return _Bpf_tp_debugClose(
		&o.bpf_tp_debugPrograms,
		&o.bpf_tp_debugMaps,
	)
}

// bpf_tp_debugMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpf_tp_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_tp_debugMaps struct {
	Events                    *ebpf.Map `ebpf:"events"`
	GoTraceMap                *ebpf.Map `ebpf:"go_trace_map"`
	GolangMapbucketStorageMap *ebpf.Map `ebpf:"golang_mapbucket_storage_map"`
	OngoingGoroutines         *ebpf.Map `ebpf:"ongoing_goroutines"`
	OngoingGrpcClientRequests *ebpf.Map `ebpf:"ongoing_grpc_client_requests"`
	OngoingGrpcHeaderWrites   *ebpf.Map `ebpf:"ongoing_grpc_header_writes"`
	OngoingGrpcRequestStatus  *ebpf.Map `ebpf:"ongoing_grpc_request_status"`
	OngoingGrpcServerRequests *ebpf.Map `ebpf:"ongoing_grpc_server_requests"`
	OngoingStreams            *ebpf.Map `ebpf:"ongoing_streams"`
	PidCache                  *ebpf.Map `ebpf:"pid_cache"`
	ValidPids                 *ebpf.Map `ebpf:"valid_pids"`
}

func (m *bpf_tp_debugMaps) Close() error {
	return _Bpf_tp_debugClose(
		m.Events,
		m.GoTraceMap,
		m.GolangMapbucketStorageMap,
		m.OngoingGoroutines,
		m.OngoingGrpcClientRequests,
		m.OngoingGrpcHeaderWrites,
		m.OngoingGrpcRequestStatus,
		m.OngoingGrpcServerRequests,
		m.OngoingStreams,
		m.PidCache,
		m.ValidPids,
	)
}

// bpf_tp_debugPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpf_tp_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_tp_debugPrograms struct {
	UprobeClientConnClose                       *ebpf.Program `ebpf:"uprobe_ClientConn_Close"`
	UprobeClientConnInvoke                      *ebpf.Program `ebpf:"uprobe_ClientConn_Invoke"`
	UprobeClientConnInvokeReturn                *ebpf.Program `ebpf:"uprobe_ClientConn_Invoke_return"`
	UprobeClientConnNewStream                   *ebpf.Program `ebpf:"uprobe_ClientConn_NewStream"`
	UprobeHpackEncoderWriteField                *ebpf.Program `ebpf:"uprobe_hpack_Encoder_WriteField"`
	UprobeServerHandleStream                    *ebpf.Program `ebpf:"uprobe_server_handleStream"`
	UprobeServerHandleStreamReturn              *ebpf.Program `ebpf:"uprobe_server_handleStream_return"`
	UprobeTransportHttp2ClientNewStream         *ebpf.Program `ebpf:"uprobe_transport_http2Client_NewStream"`
	UprobeTransportLoopyWriterWriteHeader       *ebpf.Program `ebpf:"uprobe_transport_loopyWriter_writeHeader"`
	UprobeTransportLoopyWriterWriteHeaderReturn *ebpf.Program `ebpf:"uprobe_transport_loopyWriter_writeHeader_return"`
	UprobeTransportWriteStatus                  *ebpf.Program `ebpf:"uprobe_transport_writeStatus"`
}

func (p *bpf_tp_debugPrograms) Close() error {
	return _Bpf_tp_debugClose(
		p.UprobeClientConnClose,
		p.UprobeClientConnInvoke,
		p.UprobeClientConnInvokeReturn,
		p.UprobeClientConnNewStream,
		p.UprobeHpackEncoderWriteField,
		p.UprobeServerHandleStream,
		p.UprobeServerHandleStreamReturn,
		p.UprobeTransportHttp2ClientNewStream,
		p.UprobeTransportLoopyWriterWriteHeader,
		p.UprobeTransportLoopyWriterWriteHeaderReturn,
		p.UprobeTransportWriteStatus,
	)
}

func _Bpf_tp_debugClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_tp_debug_bpfel_x86.o
var _Bpf_tp_debugBytes []byte