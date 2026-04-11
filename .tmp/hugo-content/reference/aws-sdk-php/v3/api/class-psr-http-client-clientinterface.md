Menu

- [Psr](namespace-psr.md)
- [Http](namespace-psr-http.md)
- [Client](namespace-psr-http-client.md)

## ClientInterface     in    - [Aws](package-aws.md)

### Table of Contents  [header link](class-psr-http-client-clientinterface-toc.md)

#### Methods  [header link](class-psr-http-client-clientinterface-toc-methods.md)

[sendRequest()](class-psr-http-client-clientinterface-method-sendrequest.md)
: [ResponseInterface](class-psr-http-message-responseinterface.md)Sends a PSR-7 request and returns a PSR-7 response.

### Methods  [header link](class-psr-http-client-clientinterface-methods.md)

#### sendRequest()  [header link](class-psr-http-client-clientinterface-method-sendrequest.md)

Sends a PSR-7 request and returns a PSR-7 response.

`
    public
                    sendRequest(RequestInterface $request) : ResponseInterface`

##### Parameters

$request
: [RequestInterface](class-psr-http-message-requestinterface.md)

##### Tags  [header link](class-psr-http-client-clientinterface-method-sendrequest-tags.md)

throws[ClientExceptionInterface](class-psr-http-client-clientexceptioninterface.md)

If an error happens while processing the request.

##### Return values

[ResponseInterface](class-psr-http-message-responseinterface.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-psr-http-client-clientinterface-toc-constants.md)
  - [Methods](class-psr-http-client-clientinterface-toc-methods.md)
- Methods
  - [sendRequest()](class-psr-http-client-clientinterface-method-sendrequest.md)

[Back To Top](class-psr-http-client-clientinterface-top.md)
