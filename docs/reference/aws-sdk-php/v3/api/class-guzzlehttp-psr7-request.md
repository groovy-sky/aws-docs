Menu

- [GuzzleHttp](namespace-guzzlehttp.md)
- [Psr7](namespace-guzzlehttp-psr7.md)

## Request        in package    - [Aws](package-aws.md)       implements  [RequestInterface](class-psr-http-message-requestinterface.md)  Uses  [MessageTrait](class-guzzlehttp-psr7-messagetrait.md)

PSR-7 request implementation.

### Table of Contents  [header link](class-guzzlehttp-psr7-request-toc.md)

#### Interfaces  [header link](class-guzzlehttp-psr7-request-toc-interfaces.md)

[RequestInterface](class-psr-http-message-requestinterface.md)Representation of an outgoing, client-side request.

#### Methods  [header link](class-guzzlehttp-psr7-request-toc-methods.md)

[\_\_construct()](class-guzzlehttp-psr7-request-method-construct.md)
: mixed [getBody()](class-guzzlehttp-psr7-messagetrait-method-getbody.md)
: [StreamInterface](class-psr-http-message-streaminterface.md)[getHeader()](class-guzzlehttp-psr7-messagetrait-method-getheader.md)
: array<string\|int, mixed> [getHeaderLine()](class-guzzlehttp-psr7-messagetrait-method-getheaderline.md)
: string [getHeaders()](class-guzzlehttp-psr7-messagetrait-method-getheaders.md)
: array<string\|int, mixed> [getMethod()](class-guzzlehttp-psr7-request-method-getmethod.md)
: string Retrieves the HTTP method of the request.[getProtocolVersion()](class-guzzlehttp-psr7-messagetrait-method-getprotocolversion.md)
: string [getRequestTarget()](class-guzzlehttp-psr7-request-method-getrequesttarget.md)
: string Retrieves the message's request target.[getUri()](class-guzzlehttp-psr7-request-method-geturi.md)
: [UriInterface](class-psr-http-message-uriinterface.md)Retrieves the URI instance.[hasHeader()](class-guzzlehttp-psr7-messagetrait-method-hasheader.md)
: bool [withAddedHeader()](class-guzzlehttp-psr7-messagetrait-method-withaddedheader.md)
: static [withBody()](class-guzzlehttp-psr7-messagetrait-method-withbody.md)
: static [withHeader()](class-guzzlehttp-psr7-messagetrait-method-withheader.md)
: static [withMethod()](class-guzzlehttp-psr7-request-method-withmethod.md)
: static Return an instance with the provided HTTP method.[withoutHeader()](class-guzzlehttp-psr7-messagetrait-method-withoutheader.md)
: static [withProtocolVersion()](class-guzzlehttp-psr7-messagetrait-method-withprotocolversion.md)
: static [withRequestTarget()](class-guzzlehttp-psr7-request-method-withrequesttarget.md)
: static Return an instance with the specific request-target.[withUri()](class-guzzlehttp-psr7-request-method-withuri.md)
: static Returns an instance with the provided URI.

### Methods  [header link](class-guzzlehttp-psr7-request-methods.md)

#### \_\_construct()  [header link](class-guzzlehttp-psr7-request-method-construct.md)

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

#### getBody()  [header link](class-guzzlehttp-psr7-messagetrait-method-getbody.md)

`
    public
                    getBody() : StreamInterface`

##### Return values

[StreamInterface](class-psr-http-message-streaminterface.md)

#### getHeader()  [header link](class-guzzlehttp-psr7-messagetrait-method-getheader.md)

`
    public
                    getHeader(mixed $header) : array<string|int, mixed>`

##### Parameters

$header
: mixed

##### Return values

array<string\|int, mixed>

#### getHeaderLine()  [header link](class-guzzlehttp-psr7-messagetrait-method-getheaderline.md)

`
    public
                    getHeaderLine(mixed $header) : string`

##### Parameters

$header
: mixed

##### Return values

string

#### getHeaders()  [header link](class-guzzlehttp-psr7-messagetrait-method-getheaders.md)

`
    public
                    getHeaders() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getMethod()  [header link](class-guzzlehttp-psr7-request-method-getmethod.md)

Retrieves the HTTP method of the request.

`
    public
                    getMethod() : string`

##### Return values

string
—

Returns the request method.

#### getProtocolVersion()  [header link](class-guzzlehttp-psr7-messagetrait-method-getprotocolversion.md)

`
    public
                    getProtocolVersion() : string`

##### Return values

string

#### getRequestTarget()  [header link](class-guzzlehttp-psr7-request-method-getrequesttarget.md)

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

#### getUri()  [header link](class-guzzlehttp-psr7-request-method-geturi.md)

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

#### hasHeader()  [header link](class-guzzlehttp-psr7-messagetrait-method-hasheader.md)

`
    public
                    hasHeader(mixed $header) : bool`

##### Parameters

$header
: mixed

##### Return values

bool

#### withAddedHeader()  [header link](class-guzzlehttp-psr7-messagetrait-method-withaddedheader.md)

`
    public
                    withAddedHeader(mixed $header, mixed $value) : static`

##### Parameters

$header
: mixed$value
: mixed

##### Return values

static

#### withBody()  [header link](class-guzzlehttp-psr7-messagetrait-method-withbody.md)

`
    public
                    withBody(StreamInterface $body) : static`

##### Parameters

$body
: [StreamInterface](class-psr-http-message-streaminterface.md)

##### Return values

static

#### withHeader()  [header link](class-guzzlehttp-psr7-messagetrait-method-withheader.md)

`
    public
                    withHeader(mixed $header, mixed $value) : static`

##### Parameters

$header
: mixed$value
: mixed

##### Return values

static

#### withMethod()  [header link](class-guzzlehttp-psr7-request-method-withmethod.md)

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

#### withoutHeader()  [header link](class-guzzlehttp-psr7-messagetrait-method-withoutheader.md)

`
    public
                    withoutHeader(mixed $header) : static`

##### Parameters

$header
: mixed

##### Return values

static

#### withProtocolVersion()  [header link](class-guzzlehttp-psr7-messagetrait-method-withprotocolversion.md)

`
    public
                    withProtocolVersion(mixed $version) : static`

##### Parameters

$version
: mixed

##### Return values

static

#### withRequestTarget()  [header link](class-guzzlehttp-psr7-request-method-withrequesttarget.md)

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

#### withUri()  [header link](class-guzzlehttp-psr7-request-method-withuri.md)

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
  - [Methods](class-guzzlehttp-psr7-request-toc-methods.md)
- Methods
  - [\_\_construct()](class-guzzlehttp-psr7-request-method-construct.md)
  - [getBody()](class-guzzlehttp-psr7-messagetrait-method-getbody.md)
  - [getHeader()](class-guzzlehttp-psr7-messagetrait-method-getheader.md)
  - [getHeaderLine()](class-guzzlehttp-psr7-messagetrait-method-getheaderline.md)
  - [getHeaders()](class-guzzlehttp-psr7-messagetrait-method-getheaders.md)
  - [getMethod()](class-guzzlehttp-psr7-request-method-getmethod.md)
  - [getProtocolVersion()](class-guzzlehttp-psr7-messagetrait-method-getprotocolversion.md)
  - [getRequestTarget()](class-guzzlehttp-psr7-request-method-getrequesttarget.md)
  - [getUri()](class-guzzlehttp-psr7-request-method-geturi.md)
  - [hasHeader()](class-guzzlehttp-psr7-messagetrait-method-hasheader.md)
  - [withAddedHeader()](class-guzzlehttp-psr7-messagetrait-method-withaddedheader.md)
  - [withBody()](class-guzzlehttp-psr7-messagetrait-method-withbody.md)
  - [withHeader()](class-guzzlehttp-psr7-messagetrait-method-withheader.md)
  - [withMethod()](class-guzzlehttp-psr7-request-method-withmethod.md)
  - [withoutHeader()](class-guzzlehttp-psr7-messagetrait-method-withoutheader.md)
  - [withProtocolVersion()](class-guzzlehttp-psr7-messagetrait-method-withprotocolversion.md)
  - [withRequestTarget()](class-guzzlehttp-psr7-request-method-withrequesttarget.md)
  - [withUri()](class-guzzlehttp-psr7-request-method-withuri.md)

[Back To Top](class-guzzlehttp-psr7-request-top.md)
