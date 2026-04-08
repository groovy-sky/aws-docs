Menu

- [Psr](namespace-psr.md)
- [Http](namespace-psr-http.md)
- [Client](namespace-psr-http-client.md)

## RequestExceptionInterface    extends  [ClientExceptionInterface](class-psr-http-client-clientexceptioninterface.md)   in    - [Aws](package-aws.md)

Exception for when a request failed.

Examples:

- Request is invalid (e.g. method is missing)
- Runtime request errors (e.g. the body stream is not seekable)

### Table of Contents  [header link](class-psr-http-client-requestexceptioninterface-toc.md)

#### Methods  [header link](class-psr-http-client-requestexceptioninterface-toc-methods.md)

[getRequest()](class-psr-http-client-requestexceptioninterface-method-getrequest.md)
: [RequestInterface](class-psr-http-message-requestinterface.md)Returns the request.

### Methods  [header link](class-psr-http-client-requestexceptioninterface-methods.md)

#### getRequest()  [header link](class-psr-http-client-requestexceptioninterface-method-getrequest.md)

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
  - [Constants](class-psr-http-client-requestexceptioninterface-toc-constants.md)
  - [Methods](class-psr-http-client-requestexceptioninterface-toc-methods.md)
- Methods
  - [getRequest()](class-psr-http-client-requestexceptioninterface-method-getrequest.md)

[Back To Top](class-psr-http-client-requestexceptioninterface-top.md)
