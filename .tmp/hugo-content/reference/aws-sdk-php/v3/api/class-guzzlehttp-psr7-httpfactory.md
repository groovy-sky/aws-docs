Menu

- [GuzzleHttp](namespace-guzzlehttp.md)
- [Psr7](namespace-guzzlehttp-psr7.md)

## HttpFactory        in package    - [Aws](package-aws.md)       implements  RequestFactoryInterface, ResponseFactoryInterface, ServerRequestFactoryInterface, StreamFactoryInterface, UploadedFileFactoryInterface, UriFactoryInterface

FinalYes

Implements all of the PSR-17 interfaces.

Note: in consuming code it is recommended to require the implemented interfaces
and inject the instance of this class multiple times.

### Table of Contents  [header link](class-guzzlehttp-psr7-httpfactory-toc.md)

#### Interfaces  [header link](class-guzzlehttp-psr7-httpfactory-toc-interfaces.md)

RequestFactoryInterfaceResponseFactoryInterfaceServerRequestFactoryInterfaceStreamFactoryInterfaceUploadedFileFactoryInterfaceUriFactoryInterface

#### Methods  [header link](class-guzzlehttp-psr7-httpfactory-toc-methods.md)

[createRequest()](class-guzzlehttp-psr7-httpfactory-method-createrequest.md)
: [RequestInterface](class-psr-http-message-requestinterface.md)[createResponse()](class-guzzlehttp-psr7-httpfactory-method-createresponse.md)
: [ResponseInterface](class-psr-http-message-responseinterface.md)[createServerRequest()](class-guzzlehttp-psr7-httpfactory-method-createserverrequest.md)
: [ServerRequestInterface](class-psr-http-message-serverrequestinterface.md)[createStream()](class-guzzlehttp-psr7-httpfactory-method-createstream.md)
: [StreamInterface](class-psr-http-message-streaminterface.md)[createStreamFromFile()](class-guzzlehttp-psr7-httpfactory-method-createstreamfromfile.md)
: [StreamInterface](class-psr-http-message-streaminterface.md)[createStreamFromResource()](class-guzzlehttp-psr7-httpfactory-method-createstreamfromresource.md)
: [StreamInterface](class-psr-http-message-streaminterface.md)[createUploadedFile()](class-guzzlehttp-psr7-httpfactory-method-createuploadedfile.md)
: [UploadedFileInterface](class-psr-http-message-uploadedfileinterface.md)[createUri()](class-guzzlehttp-psr7-httpfactory-method-createuri.md)
: [UriInterface](class-psr-http-message-uriinterface.md)

### Methods  [header link](class-guzzlehttp-psr7-httpfactory-methods.md)

#### createRequest()  [header link](class-guzzlehttp-psr7-httpfactory-method-createrequest.md)

`
    public
                    createRequest(string $method, mixed $uri) : RequestInterface`

##### Parameters

$method
: string$uri
: mixed

##### Return values

[RequestInterface](class-psr-http-message-requestinterface.md)

#### createResponse()  [header link](class-guzzlehttp-psr7-httpfactory-method-createresponse.md)

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

#### createServerRequest()  [header link](class-guzzlehttp-psr7-httpfactory-method-createserverrequest.md)

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

#### createStream()  [header link](class-guzzlehttp-psr7-httpfactory-method-createstream.md)

`
    public
                    createStream([string $content = '' ]) : StreamInterface`

##### Parameters

$content
: string
= ''

##### Return values

[StreamInterface](class-psr-http-message-streaminterface.md)

#### createStreamFromFile()  [header link](class-guzzlehttp-psr7-httpfactory-method-createstreamfromfile.md)

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

#### createStreamFromResource()  [header link](class-guzzlehttp-psr7-httpfactory-method-createstreamfromresource.md)

`
    public
                    createStreamFromResource(mixed $resource) : StreamInterface`

##### Parameters

$resource
: mixed

##### Return values

[StreamInterface](class-psr-http-message-streaminterface.md)

#### createUploadedFile()  [header link](class-guzzlehttp-psr7-httpfactory-method-createuploadedfile.md)

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

#### createUri()  [header link](class-guzzlehttp-psr7-httpfactory-method-createuri.md)

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
  - [Methods](class-guzzlehttp-psr7-httpfactory-toc-methods.md)
- Methods
  - [createRequest()](class-guzzlehttp-psr7-httpfactory-method-createrequest.md)
  - [createResponse()](class-guzzlehttp-psr7-httpfactory-method-createresponse.md)
  - [createServerRequest()](class-guzzlehttp-psr7-httpfactory-method-createserverrequest.md)
  - [createStream()](class-guzzlehttp-psr7-httpfactory-method-createstream.md)
  - [createStreamFromFile()](class-guzzlehttp-psr7-httpfactory-method-createstreamfromfile.md)
  - [createStreamFromResource()](class-guzzlehttp-psr7-httpfactory-method-createstreamfromresource.md)
  - [createUploadedFile()](class-guzzlehttp-psr7-httpfactory-method-createuploadedfile.md)
  - [createUri()](class-guzzlehttp-psr7-httpfactory-method-createuri.md)

[Back To Top](class-guzzlehttp-psr7-httpfactory-top.md)
