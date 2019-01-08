package testdata

const UnaryRPCsClientEndpointInitCode = `// MethodUnaryRPCA calls the "MethodUnaryRPCA" function in
// pb.ServiceUnaryRPCsClient interface.
func (c *Client) MethodUnaryRPCA() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildMethodUnaryRPCAFunc(c.grpccli, c.opts...),
			EncodeMethodUnaryRPCARequest,
			DecodeMethodUnaryRPCAResponse)
		return inv.Invoke(ctx, v)
	}
}

// MethodUnaryRPCB calls the "MethodUnaryRPCB" function in
// pb.ServiceUnaryRPCsClient interface.
func (c *Client) MethodUnaryRPCB() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildMethodUnaryRPCBFunc(c.grpccli, c.opts...),
			EncodeMethodUnaryRPCBRequest,
			DecodeMethodUnaryRPCBResponse)
		return inv.Invoke(ctx, v)
	}
}
`

const UnaryRPCNoPayloadClientEndpointInitCode = `// MethodUnaryRPCNoPayload calls the "MethodUnaryRPCNoPayload" function in
// pb.ServiceUnaryRPCNoPayloadClient interface.
func (c *Client) MethodUnaryRPCNoPayload() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildMethodUnaryRPCNoPayloadFunc(c.grpccli, c.opts...),
			nil,
			DecodeMethodUnaryRPCNoPayloadResponse)
		return inv.Invoke(ctx, v)
	}
}
`

const UnaryRPCNoResultClientEndpointInitCode = `// MethodUnaryRPCNoResult calls the "MethodUnaryRPCNoResult" function in
// pb.ServiceUnaryRPCNoResultClient interface.
func (c *Client) MethodUnaryRPCNoResult() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildMethodUnaryRPCNoResultFunc(c.grpccli, c.opts...),
			EncodeMethodUnaryRPCNoResultRequest,
			nil)
		return inv.Invoke(ctx, v)
	}
}
`

const ServerStreamingRPCClientEndpointInitCode = `// MethodServerStreamingRPC calls the "MethodServerStreamingRPC" function in
// pb.ServiceServerStreamingRPCClient interface.
func (c *Client) MethodServerStreamingRPC() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildMethodServerStreamingRPCFunc(c.grpccli, c.opts...),
			EncodeMethodServerStreamingRPCRequest,
			DecodeMethodServerStreamingRPCResponse)
		return inv.Invoke(ctx, v)
	}
}
`

const ClientStreamingRPCClientEndpointInitCode = `// MethodClientStreamingRPC calls the "MethodClientStreamingRPC" function in
// pb.ServiceClientStreamingRPCClient interface.
func (c *Client) MethodClientStreamingRPC() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildMethodClientStreamingRPCFunc(c.grpccli, c.opts...),
			nil,
			DecodeMethodClientStreamingRPCResponse)
		return inv.Invoke(ctx, v)
	}
}
`

const ClientStreamingRPCWithPayloadClientEndpointInitCode = `// MethodClientStreamingRPCWithPayload calls the
// "MethodClientStreamingRPCWithPayload" function in
// pb.ServiceClientStreamingRPCWithPayloadClient interface.
func (c *Client) MethodClientStreamingRPCWithPayload() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildMethodClientStreamingRPCWithPayloadFunc(c.grpccli, c.opts...),
			EncodeMethodClientStreamingRPCWithPayloadRequest,
			DecodeMethodClientStreamingRPCWithPayloadResponse)
		return inv.Invoke(ctx, v)
	}
}
`

const BidirectionalStreamingRPCClientEndpointInitCode = `// MethodBidirectionalStreamingRPC calls the "MethodBidirectionalStreamingRPC"
// function in pb.ServiceBidirectionalStreamingRPCClient interface.
func (c *Client) MethodBidirectionalStreamingRPC() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildMethodBidirectionalStreamingRPCFunc(c.grpccli, c.opts...),
			nil,
			DecodeMethodBidirectionalStreamingRPCResponse)
		return inv.Invoke(ctx, v)
	}
}
`

const BidirectionalStreamingRPCWithPayloadClientEndpointInitCode = `// MethodBidirectionalStreamingRPCWithPayload calls the
// "MethodBidirectionalStreamingRPCWithPayload" function in
// pb.ServiceBidirectionalStreamingRPCWithPayloadClient interface.
func (c *Client) MethodBidirectionalStreamingRPCWithPayload() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildMethodBidirectionalStreamingRPCWithPayloadFunc(c.grpccli, c.opts...),
			EncodeMethodBidirectionalStreamingRPCWithPayloadRequest,
			DecodeMethodBidirectionalStreamingRPCWithPayloadResponse)
		return inv.Invoke(ctx, v)
	}
}
`
