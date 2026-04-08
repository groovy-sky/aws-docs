Menu

- [Psr](namespace-psr.md)
- [Http](namespace-psr-http.md)
- [Client](namespace-psr-http-client.md)

## NetworkExceptionInterface    extends  [ClientExceptionInterface](class-psr-http-client-clientexceptioninterface.md)   in    - [Aws](package-aws.md)

Thrown when the request cannot be completed because of network issues.

There is no response object as this exception is thrown when no response has been received.

Example: the target host name can not be resolved or the connection failed.

### Table of Contents  [header link](class-psr-http-client-networkexceptioninterface-toc.md)

#### Methods  [header link](class-psr-http-client-networkexceptioninterface-toc-methods.md)

[getRequest()](class-psr-http-client-networkexceptioninterface-method-getrequest.md)
: [RequestInterface](class-psr-http-message-requestinterface.md)Returns the request.

### Methods  [header link](class-psr-http-client-networkexceptioninterface-methods.md)

#### getRequest()  [header link](class-psr-http-client-networkexceptioninterface-method-getrequest.md)

Returns the request.

`
    public
                    getRequest() : RequestInterface`

The request object MAY be a different object from the one passed to ClientInterface::sendRequest()

##### Return values

[RequestInterface](class-psr-http-message-requestinterface.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-psr-http-client-networkexceptioninterface-toc-constants.md)
  - [Methods](class-psr-http-client-networkexceptioninterface-toc-methods.md)
- Methods
  - [getRequest()](class-psr-http-client-networkexceptioninterface-method-getrequest.md)

[Back To Top](class-psr-http-client-networkexceptioninterface-top.md)
