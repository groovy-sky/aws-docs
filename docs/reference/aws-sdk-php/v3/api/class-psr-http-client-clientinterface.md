Menu

- [Psr](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Psr.html)
- [Http](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Psr.http.html)
- [Client](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Psr.http.client.html)

## ClientInterface     in    - [Aws](package-aws.md)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Client.ClientInterface.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Client.ClientInterface.html\#toc-methods)

[sendRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Client.ClientInterface.html#method_sendRequest)
: [ResponseInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.ResponseInterface.html)Sends a PSR-7 request and returns a PSR-7 response.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Client.ClientInterface.html\#methods)

#### sendRequest()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Client.ClientInterface.html\#method_sendRequest)

Sends a PSR-7 request and returns a PSR-7 response.

`
    public
                    sendRequest(RequestInterface $request) : ResponseInterface`

##### Parameters

$request
: [RequestInterface](class-psr-http-message-requestinterface.md)

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Client.ClientInterface.html\#method_sendRequest\#tags)

throws[ClientExceptionInterface](class-psr-http-client-clientexceptioninterface.md)

If an error happens while processing the request.

##### Return values

[ResponseInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.ResponseInterface.html)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Client.ClientInterface.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Client.ClientInterface.html#toc-methods)
- Methods
  - [sendRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Client.ClientInterface.html#method_sendRequest)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Client.ClientInterface.html#top)
