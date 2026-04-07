Menu

- [GuzzleHttp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.html)
- [Psr7](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.psr7.html)

## Request        in package    - [Aws](package-aws.md)       implements  [RequestInterface](class-psr-http-message-requestinterface.md)  Uses  [MessageTrait](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html)

PSR-7 request implementation.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Request.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Request.html\#toc-interfaces)

[RequestInterface](class-psr-http-message-requestinterface.md)Representation of an outgoing, client-side request.

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Request.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Request.html#method___construct)
: mixed [getBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getBody)
: [StreamInterface](class-psr-http-message-streaminterface.md)[getHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getHeader)
: array<string\|int, mixed> [getHeaderLine()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getHeaderLine)
: string [getHeaders()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getHeaders)
: array<string\|int, mixed> [getMethod()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Request.html#method_getMethod)
: string Retrieves the HTTP method of the request.[getProtocolVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getProtocolVersion)
: string [getRequestTarget()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Request.html#method_getRequestTarget)
: string Retrieves the message's request target.[getUri()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Request.html#method_getUri)
: [UriInterface](class-psr-http-message-uriinterface.md)Retrieves the URI instance.[hasHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_hasHeader)
: bool [withAddedHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withAddedHeader)
: static [withBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withBody)
: static [withHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withHeader)
: static [withMethod()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Request.html#method_withMethod)
: static Return an instance with the provided HTTP method.[withoutHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withoutHeader)
: static [withProtocolVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withProtocolVersion)
: static [withRequestTarget()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Request.html#method_withRequestTarget)
: static Return an instance with the specific request-target.[withUri()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Request.html#method_withUri)
: static Returns an instance with the provided URI.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Request.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Request.html\#method___construct)

`
    public
                    __construct(string $method, string|UriInterface $uri[, array<string|int, string|array<string|int, string>> $headers = [] ][, string|resource|StreamInterface|null $body = null ][, string $version = '1.1' ]) : mixed`

##### Parameters

$method
: string

HTTP method

$uri
: string\| [UriInterface](class-psr-http-message-uriinterface.md)

URI

$headers
: array<string\|int, string\|array<string\|int, string>>
= \[\]

Request headers

$body
: string\|resource\| [StreamInterface](class-psr-http-message-streaminterface.md) \|null
= null

Request body

$version
: string
= '1.1'

Protocol version

#### getBody()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html\#method_getBody)

`
    public
                    getBody() : StreamInterface`

##### Return values

[StreamInterface](class-psr-http-message-streaminterface.md)

#### getHeader()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html\#method_getHeader)

`
    public
                    getHeader(mixed $header) : array<string|int, mixed>`

##### Parameters

$header
: mixed

##### Return values

array<string\|int, mixed>

#### getHeaderLine()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html\#method_getHeaderLine)

`
    public
                    getHeaderLine(mixed $header) : string`

##### Parameters

$header
: mixed

##### Return values

string

#### getHeaders()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html\#method_getHeaders)

`
    public
                    getHeaders() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getMethod()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Request.html\#method_getMethod)

Retrieves the HTTP method of the request.

`
    public
                    getMethod() : string`

##### Return values

string
—

Returns the request method.

#### getProtocolVersion()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html\#method_getProtocolVersion)

`
    public
                    getProtocolVersion() : string`

##### Return values

string

#### getRequestTarget()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Request.html\#method_getRequestTarget)

Retrieves the message's request target.

`
    public
                    getRequestTarget() : string`

Retrieves the message's request-target either as it will appear (for
clients), as it appeared at request (for servers), or as it was
specified for the instance (see withRequestTarget()).

In most cases, this will be the origin-form of the composed URI,
unless a value was provided to the concrete implementation (see
withRequestTarget() below).

If no URI is available, and no request-target has been specifically
provided, this method MUST return the string "/".

##### Return values

string

#### getUri()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Request.html\#method_getUri)

Retrieves the URI instance.

`
    public
                    getUri() : UriInterface`

This method MUST return a UriInterface instance.

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)
—

Returns a UriInterface instance
representing the URI of the request.

#### hasHeader()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html\#method_hasHeader)

`
    public
                    hasHeader(mixed $header) : bool`

##### Parameters

$header
: mixed

##### Return values

bool

#### withAddedHeader()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html\#method_withAddedHeader)

`
    public
                    withAddedHeader(mixed $header, mixed $value) : static`

##### Parameters

$header
: mixed$value
: mixed

##### Return values

static

#### withBody()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html\#method_withBody)

`
    public
                    withBody(StreamInterface $body) : static`

##### Parameters

$body
: [StreamInterface](class-psr-http-message-streaminterface.md)

##### Return values

static

#### withHeader()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html\#method_withHeader)

`
    public
                    withHeader(mixed $header, mixed $value) : static`

##### Parameters

$header
: mixed$value
: mixed

##### Return values

static

#### withMethod()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Request.html\#method_withMethod)

Return an instance with the provided HTTP method.

`
    public
                    withMethod(mixed $method) : static`

While HTTP method names are typically all uppercase characters, HTTP
method names are case-sensitive and thus implementations SHOULD NOT
modify the given string.

This method MUST be implemented in such a way as to retain the
immutability of the message, and MUST return an instance that has the
changed request method.

##### Parameters

$method
: mixed

Case-sensitive method.

##### Return values

static

#### withoutHeader()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html\#method_withoutHeader)

`
    public
                    withoutHeader(mixed $header) : static`

##### Parameters

$header
: mixed

##### Return values

static

#### withProtocolVersion()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html\#method_withProtocolVersion)

`
    public
                    withProtocolVersion(mixed $version) : static`

##### Parameters

$version
: mixed

##### Return values

static

#### withRequestTarget()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Request.html\#method_withRequestTarget)

Return an instance with the specific request-target.

`
    public
                    withRequestTarget(mixed $requestTarget) : static`

If the request needs a non-origin-form request-target — e.g., for
specifying an absolute-form, authority-form, or asterisk-form —
this method may be used to create an instance with the specified
request-target, verbatim.

This method MUST be implemented in such a way as to retain the
immutability of the message, and MUST return an instance that has the
changed request target.

##### Parameters

$requestTarget
: mixed

##### Return values

static

#### withUri()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Request.html\#method_withUri)

Returns an instance with the provided URI.

`
    public
                    withUri(UriInterface $uri[, mixed $preserveHost = false ]) : static`

This method MUST update the Host header of the returned request by
default if the URI contains a host component. If the URI does not
contain a host component, any pre-existing Host header MUST be carried
over to the returned request.

You can opt-in to preserving the original state of the Host header by
setting `$preserveHost` to `true`. When `$preserveHost` is set to
`true`, this method interacts with the Host header in the following ways:

- If the Host header is missing or empty, and the new URI contains
a host component, this method MUST update the Host header in the returned
request.
- If the Host header is missing or empty, and the new URI does not contain a
host component, this method MUST NOT update the Host header in the returned
request.
- If a Host header is present and non-empty, this method MUST NOT update
the Host header in the returned request.

This method MUST be implemented in such a way as to retain the
immutability of the message, and MUST return an instance that has the
new UriInterface instance.

##### Parameters

$uri
: [UriInterface](class-psr-http-message-uriinterface.md)

New request URI to use.

$preserveHost
: mixed
= false

Preserve the original state of the Host header.

##### Return values

static
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Request.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Request.html#method___construct)
  - [getBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getBody)
  - [getHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getHeader)
  - [getHeaderLine()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getHeaderLine)
  - [getHeaders()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getHeaders)
  - [getMethod()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Request.html#method_getMethod)
  - [getProtocolVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getProtocolVersion)
  - [getRequestTarget()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Request.html#method_getRequestTarget)
  - [getUri()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Request.html#method_getUri)
  - [hasHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_hasHeader)
  - [withAddedHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withAddedHeader)
  - [withBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withBody)
  - [withHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withHeader)
  - [withMethod()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Request.html#method_withMethod)
  - [withoutHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withoutHeader)
  - [withProtocolVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withProtocolVersion)
  - [withRequestTarget()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Request.html#method_withRequestTarget)
  - [withUri()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Request.html#method_withUri)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Request.html#top)
