// this file was auto generated by gotsrpc https://github.com/foomo/gotsrpc
package demo

import (
	tls "crypto/tls"
	gob "encoding/gob"
	fmt "fmt"
	gotsrpc "github.com/foomo/gotsrpc"
	nested "github.com/foomo/gotsrpc/demo/nested"
	gorpc "github.com/valyala/gorpc"
	reflect "reflect"
	strings "strings"
	time "time"
)

type (
	FooGoRPCProxy struct {
		server           *gorpc.Server
		service          *Foo
		callStatsHandler gotsrpc.GoRPCCallStatsHandlerFun
	}

	FooHelloRequest struct {
		Number int64
	}
	FooHelloResponse struct {
		RetHello_0 int
	}
)

func init() {
	gob.Register(FooHelloRequest{})
	gob.Register(FooHelloResponse{})
}

func NewFooGoRPCProxy(addr string, service *Foo, tlsConfig *tls.Config) *FooGoRPCProxy {
	proxy := &FooGoRPCProxy{
		service: service,
	}

	if tlsConfig != nil {
		proxy.server = gorpc.NewTLSServer(addr, proxy.handler, tlsConfig)
	} else {
		proxy.server = gorpc.NewTCPServer(addr, proxy.handler)
	}

	return proxy
}

func (p *FooGoRPCProxy) Start() error {
	return p.server.Start()
}

func (p *FooGoRPCProxy) Serve() error {
	return p.server.Serve()
}

func (p *FooGoRPCProxy) Stop() {
	p.server.Stop()
}

func (p *FooGoRPCProxy) SetCallStatsHandler(handler gotsrpc.GoRPCCallStatsHandlerFun) {
	p.callStatsHandler = handler
}

func (p *FooGoRPCProxy) handler(clientAddr string, request interface{}) (response interface{}) {
	start := time.Now()

	reqType := reflect.TypeOf(request).String()
	funcNameParts := strings.Split(reqType, ".")
	funcName := funcNameParts[len(funcNameParts)-1]

	switch funcName {
	case "FooHelloRequest":
		req := request.(FooHelloRequest)
		retHello_0 := p.service.Hello(req.Number)
		response = FooHelloResponse{RetHello_0: retHello_0}
	default:
		fmt.Println("Unkown request type", reflect.TypeOf(request).String())
	}

	if p.callStatsHandler != nil {
		p.callStatsHandler(&gotsrpc.CallStats{
			Func:      funcName,
			Package:   "github.com/foomo/gotsrpc/demo",
			Service:   "Foo",
			Execution: time.Since(start),
		})
	}

	return
}

type (
	DemoGoRPCProxy struct {
		server           *gorpc.Server
		service          *Demo
		callStatsHandler gotsrpc.GoRPCCallStatsHandlerFun
	}

	DemoExtractAddressRequest struct {
		Person *Person
	}
	DemoExtractAddressResponse struct {
		Addr *Address
		E    *Err
	}

	DemoGiveMeAScalarRequest struct {
	}
	DemoGiveMeAScalarResponse struct {
		Amount nested.Amount
		Wahr   nested.True
		Hier   ScalarInPlace
	}

	DemoHelloRequest struct {
		Name string
	}
	DemoHelloResponse struct {
		RetHello_0 string
		RetHello_1 *Err
	}

	DemoHelloInterfaceRequest struct {
		Anything      interface{}
		AnythingMap   map[string]interface{}
		AnythingSlice []interface{}
	}
	DemoHelloInterfaceResponse struct {
	}

	DemoMapCrapRequest struct {
	}
	DemoMapCrapResponse struct {
		Crap map[string][]int
	}

	DemoNestRequest struct {
	}
	DemoNestResponse struct {
		RetNest_0 *nested.Nested
	}

	DemoTestScalarInPlaceRequest struct {
	}
	DemoTestScalarInPlaceResponse struct {
		RetTestScalarInPlace_0 ScalarInPlace
	}
)

func init() {
	gob.Register(DemoExtractAddressRequest{})
	gob.Register(DemoExtractAddressResponse{})
	gob.Register(DemoGiveMeAScalarRequest{})
	gob.Register(DemoGiveMeAScalarResponse{})
	gob.Register(DemoHelloRequest{})
	gob.Register(DemoHelloResponse{})
	gob.Register(DemoHelloInterfaceRequest{})
	gob.Register(DemoHelloInterfaceResponse{})
	gob.Register(DemoMapCrapRequest{})
	gob.Register(DemoMapCrapResponse{})
	gob.Register(DemoNestRequest{})
	gob.Register(DemoNestResponse{})
	gob.Register(DemoTestScalarInPlaceRequest{})
	gob.Register(DemoTestScalarInPlaceResponse{})
}

func NewDemoGoRPCProxy(addr string, service *Demo, tlsConfig *tls.Config) *DemoGoRPCProxy {
	proxy := &DemoGoRPCProxy{
		service: service,
	}

	if tlsConfig != nil {
		proxy.server = gorpc.NewTLSServer(addr, proxy.handler, tlsConfig)
	} else {
		proxy.server = gorpc.NewTCPServer(addr, proxy.handler)
	}

	return proxy
}

func (p *DemoGoRPCProxy) Start() error {
	return p.server.Start()
}

func (p *DemoGoRPCProxy) Serve() error {
	return p.server.Serve()
}

func (p *DemoGoRPCProxy) Stop() {
	p.server.Stop()
}

func (p *DemoGoRPCProxy) SetCallStatsHandler(handler gotsrpc.GoRPCCallStatsHandlerFun) {
	p.callStatsHandler = handler
}

func (p *DemoGoRPCProxy) handler(clientAddr string, request interface{}) (response interface{}) {
	start := time.Now()

	reqType := reflect.TypeOf(request).String()
	funcNameParts := strings.Split(reqType, ".")
	funcName := funcNameParts[len(funcNameParts)-1]

	switch funcName {
	case "DemoExtractAddressRequest":
		req := request.(DemoExtractAddressRequest)
		addr, e := p.service.ExtractAddress(req.Person)
		response = DemoExtractAddressResponse{Addr: addr, E: e}
	case "DemoGiveMeAScalarRequest":
		amount, wahr, hier := p.service.GiveMeAScalar()
		response = DemoGiveMeAScalarResponse{Amount: amount, Wahr: wahr, Hier: hier}
	case "DemoHelloRequest":
		req := request.(DemoHelloRequest)
		retHello_0, retHello_1 := p.service.Hello(req.Name)
		response = DemoHelloResponse{RetHello_0: retHello_0, RetHello_1: retHello_1}
	case "DemoHelloInterfaceRequest":
		req := request.(DemoHelloInterfaceRequest)
		p.service.HelloInterface(req.Anything, req.AnythingMap, req.AnythingSlice)
		response = DemoHelloInterfaceResponse{}
	case "DemoMapCrapRequest":
		crap := p.service.MapCrap()
		response = DemoMapCrapResponse{Crap: crap}
	case "DemoNestRequest":
		retNest_0 := p.service.Nest()
		response = DemoNestResponse{RetNest_0: retNest_0}
	case "DemoTestScalarInPlaceRequest":
		retTestScalarInPlace_0 := p.service.TestScalarInPlace()
		response = DemoTestScalarInPlaceResponse{RetTestScalarInPlace_0: retTestScalarInPlace_0}
	default:
		fmt.Println("Unkown request type", reflect.TypeOf(request).String())
	}

	if p.callStatsHandler != nil {
		p.callStatsHandler(&gotsrpc.CallStats{
			Func:      funcName,
			Package:   "github.com/foomo/gotsrpc/demo",
			Service:   "Demo",
			Execution: time.Since(start),
		})
	}

	return
}
