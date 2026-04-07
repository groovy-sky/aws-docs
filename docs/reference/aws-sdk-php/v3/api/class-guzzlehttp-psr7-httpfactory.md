Menu

- [GuzzleHttp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.html)
- [Psr7](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.psr7.html)

## HttpFactory        in package    - [Aws](package-aws.md)       implements  RequestFactoryInterface, ResponseFactoryInterface, ServerRequestFactoryInterface, StreamFactoryInterface, UploadedFileFactoryInterface, UriFactoryInterface

FinalYes

Implements all of the PSR-17 interfaces.

Note: in consuming code it is recommended to require the implemented interfaces
and inject the instance of this class multiple times.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html\#toc-interfaces)

RequestFactoryInterfaceResponseFactoryInterfaceServerRequestFactoryInterfaceStreamFactoryInterfaceUploadedFileFactoryInterfaceUriFactoryInterface

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html\#toc-methods)

[createRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html#method_createRequest)
: [RequestInterface](class-psr-http-message-requestinterface.md)[createResponse()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html#method_createResponse)
: [ResponseInterface](class-psr-http-message-responseinterface.md)[createServerRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html#method_createServerRequest)
: [ServerRequestInterface](class-psr-http-message-serverrequestinterface.md)[createStream()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html#method_createStream)
: [StreamInterface](class-psr-http-message-streaminterface.md)[createStreamFromFile()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html#method_createStreamFromFile)
: [StreamInterface](class-psr-http-message-streaminterface.md)[createStreamFromResource()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html#method_createStreamFromResource)
: [StreamInterface](class-psr-http-message-streaminterface.md)[createUploadedFile()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html#method_createUploadedFile)
: [UploadedFileInterface](class-psr-http-message-uploadedfileinterface.md)[createUri()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html#method_createUri)
: [UriInterface](class-psr-http-message-uriinterface.md)

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html\#methods)

#### createRequest()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html\#method_createRequest)

`
    public
                    createRequest(string $method, mixed $uri) : RequestInterface`

##### Parameters

$method
: string$uri
: mixed

##### Return values

[RequestInterface](class-psr-http-message-requestinterface.md)

#### createResponse()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html\#method_createResponse)

`
    public
                    createResponse([int $code = 200 ][, string $reasonPhrase = '' ]) : ResponseInterface`

##### Parameters

$code
: int
= 200$reasonPhrase
: string
= ''

##### Return values

[ResponseInterface](class-psr-http-message-responseinterface.md)

#### createServerRequest()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html\#method_createServerRequest)

`
    public
                    createServerRequest(string $method, mixed $uri[, array<string|int, mixed> $serverParams = [] ]) : ServerRequestInterface`

##### Parameters

$method
: string$uri
: mixed$serverParams
: array<string\|int, mixed>
= \[\]

##### Return values

[ServerRequestInterface](class-psr-http-message-serverrequestinterface.md)

#### createStream()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html\#method_createStream)

`
    public
                    createStream([string $content = '' ]) : StreamInterface`

##### Parameters

$content
: string
= ''

##### Return values

[StreamInterface](class-psr-http-message-streaminterface.md)

#### createStreamFromFile()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html\#method_createStreamFromFile)

`
    public
                    createStreamFromFile(string $file[, string $mode = 'r' ]) : StreamInterface`

##### Parameters

$file
: string$mode
: string
= 'r'

##### Return values

[StreamInterface](class-psr-http-message-streaminterface.md)

#### createStreamFromResource()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html\#method_createStreamFromResource)

`
    public
                    createStreamFromResource(mixed $resource) : StreamInterface`

##### Parameters

$resource
: mixed

##### Return values

[StreamInterface](class-psr-http-message-streaminterface.md)

#### createUploadedFile()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html\#method_createUploadedFile)

`
    public
                    createUploadedFile(StreamInterface $stream[, int|null $size = null ][, int $error = UPLOAD_ERR_OK ][, string|null $clientFilename = null ][, string|null $clientMediaType = null ]) : UploadedFileInterface`

##### Parameters

$stream
: [StreamInterface](class-psr-http-message-streaminterface.md)$size
: int\|null
= null$error
: int
= UPLOAD\_ERR\_OK$clientFilename
: string\|null
= null$clientMediaType
: string\|null
= null

##### Return values

[UploadedFileInterface](class-psr-http-message-uploadedfileinterface.md)

#### createUri()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html\#method_createUri)

`
    public
                    createUri([string $uri = '' ]) : UriInterface`

##### Parameters

$uri
: string
= ''

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html#toc-methods)
- Methods
  - [createRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html#method_createRequest)
  - [createResponse()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html#method_createResponse)
  - [createServerRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html#method_createServerRequest)
  - [createStream()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html#method_createStream)
  - [createStreamFromFile()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html#method_createStreamFromFile)
  - [createStreamFromResource()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html#method_createStreamFromResource)
  - [createUploadedFile()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html#method_createUploadedFile)
  - [createUri()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html#method_createUri)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.HttpFactory.html#top)
