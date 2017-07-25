// this file was auto generated by gotsrpc https://github.com/foomo/gotsrpc
package demo

import (
	gotsrpc "github.com/foomo/gotsrpc"
	nested "github.com/foomo/gotsrpc/demo/nested"
)

type FooGoTSRPCClient struct {
	URL      string
	EndPoint string
}

func NewDefaultFooGoTSRPCClient(url string) *FooGoTSRPCClient {
	return NewFooGoTSRPCClient(url, "/service/foo")
}

func NewFooGoTSRPCClient(url string, endpoint string) *FooGoTSRPCClient {
	return &FooGoTSRPCClient{
		URL:      url,
		EndPoint: endpoint,
	}
}

func (goTSRPCClientInstance *FooGoTSRPCClient) Hello(number int64) (retHello_0 int, clientErr error) {
	args := []interface{}{number}
	reply := []interface{}{&retHello_0}
	clientErr = gotsrpc.CallClient(goTSRPCClientInstance.URL, goTSRPCClientInstance.EndPoint, "Hello", args, reply)
	return
}

type DemoGoTSRPCClient struct {
	URL      string
	EndPoint string
}

func NewDefaultDemoGoTSRPCClient(url string) *DemoGoTSRPCClient {
	return NewDemoGoTSRPCClient(url, "/service/demo")
}

func NewDemoGoTSRPCClient(url string, endpoint string) *DemoGoTSRPCClient {
	return &DemoGoTSRPCClient{
		URL:      url,
		EndPoint: endpoint,
	}
}

func (goTSRPCClientInstance *DemoGoTSRPCClient) ExtractAddress(person *Person) (addr *Address, e *Err, clientErr error) {
	args := []interface{}{person}
	reply := []interface{}{&addr, &e}
	clientErr = gotsrpc.CallClient(goTSRPCClientInstance.URL, goTSRPCClientInstance.EndPoint, "ExtractAddress", args, reply)
	return
}

func (goTSRPCClientInstance *DemoGoTSRPCClient) GiveMeAScalar() (amount nested.Amount, wahr nested.True, hier ScalarInPlace, clientErr error) {
	args := []interface{}{}
	reply := []interface{}{&amount, &wahr, &hier}
	clientErr = gotsrpc.CallClient(goTSRPCClientInstance.URL, goTSRPCClientInstance.EndPoint, "GiveMeAScalar", args, reply)
	return
}

func (goTSRPCClientInstance *DemoGoTSRPCClient) Hello(name string) (retHello_0 string, retHello_1 *Err, clientErr error) {
	args := []interface{}{name}
	reply := []interface{}{&retHello_0, &retHello_1}
	clientErr = gotsrpc.CallClient(goTSRPCClientInstance.URL, goTSRPCClientInstance.EndPoint, "Hello", args, reply)
	return
}

func (goTSRPCClientInstance *DemoGoTSRPCClient) HelloInterface(anything interface{}, anythingMap map[string]interface{}, anythingSlice []interface{}) (clientErr error) {
	args := []interface{}{anything, anythingMap, anythingSlice}
	reply := []interface{}{}
	clientErr = gotsrpc.CallClient(goTSRPCClientInstance.URL, goTSRPCClientInstance.EndPoint, "HelloInterface", args, reply)
	return
}

func (goTSRPCClientInstance *DemoGoTSRPCClient) MapCrap() (crap map[string][]int, clientErr error) {
	args := []interface{}{}
	reply := []interface{}{&crap}
	clientErr = gotsrpc.CallClient(goTSRPCClientInstance.URL, goTSRPCClientInstance.EndPoint, "MapCrap", args, reply)
	return
}

func (goTSRPCClientInstance *DemoGoTSRPCClient) Nest() (retNest_0 *nested.Nested, clientErr error) {
	args := []interface{}{}
	reply := []interface{}{&retNest_0}
	clientErr = gotsrpc.CallClient(goTSRPCClientInstance.URL, goTSRPCClientInstance.EndPoint, "Nest", args, reply)
	return
}

func (goTSRPCClientInstance *DemoGoTSRPCClient) TestScalarInPlace() (retTestScalarInPlace_0 ScalarInPlace, clientErr error) {
	args := []interface{}{}
	reply := []interface{}{&retTestScalarInPlace_0}
	clientErr = gotsrpc.CallClient(goTSRPCClientInstance.URL, goTSRPCClientInstance.EndPoint, "TestScalarInPlace", args, reply)
	return
}
