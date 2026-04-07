Menu

- [Psr](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Psr.html)
- [Http](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Psr.http.html)
- [Client](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Psr.http.client.html)

## RequestExceptionInterface    extends  [ClientExceptionInterface](class-psr-http-client-clientexceptioninterface.md)   in    - [Aws](package-aws.md)

Exception for when a request failed.

Examples:

- Request is invalid (e.g. method is missing)
- Runtime request errors (e.g. the body stream is not seekable)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Client.RequestExceptionInterface.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Client.RequestExceptionInterface.html\#toc-methods)

[getRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Client.RequestExceptionInterface.html#method_getRequest)
: [RequestInterface](class-psr-http-message-requestinterface.md)Returns the request.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Client.RequestExceptionInterface.html\#methods)

#### getRequest()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Client.RequestExceptionInterface.html\#method_getRequest)

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
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Client.RequestExceptionInterface.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Client.RequestExceptionInterface.html#toc-methods)
- Methods
  - [getRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Client.RequestExceptionInterface.html#method_getRequest)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Client.RequestExceptionInterface.html#top)
