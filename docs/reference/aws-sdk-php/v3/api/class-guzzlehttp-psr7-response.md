Menu

- [GuzzleHttp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.html)
- [Psr7](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.psr7.html)

## Response        in package    - [Aws](package-aws.md)       implements  [ResponseInterface](class-psr-http-message-responseinterface.md)  Uses  [MessageTrait](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html)

PSR-7 response implementation.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Response.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Response.html\#toc-interfaces)

[ResponseInterface](class-psr-http-message-responseinterface.md)Representation of an outgoing, server-side response.

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Response.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Response.html#method___construct)
: mixed [getBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getBody)
: [StreamInterface](class-psr-http-message-streaminterface.md)[getHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getHeader)
: array<string\|int, mixed> [getHeaderLine()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getHeaderLine)
: string [getHeaders()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getHeaders)
: array<string\|int, mixed> [getProtocolVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getProtocolVersion)
: string [getReasonPhrase()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Response.html#method_getReasonPhrase)
: string Gets the response reason phrase associated with the status code.[getStatusCode()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Response.html#method_getStatusCode)
: int Gets the response status code.[hasHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_hasHeader)
: bool [withAddedHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withAddedHeader)
: static [withBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withBody)
: static [withHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withHeader)
: static [withoutHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withoutHeader)
: static [withProtocolVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withProtocolVersion)
: static [withStatus()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Response.html#method_withStatus)
: static Return an instance with the specified status code and, optionally, reason phrase.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Response.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Response.html\#method___construct)

`
    public
                    __construct([int $status = 200 ][, array<string|int, string|array<string|int, string>> $headers = [] ][, string|resource|StreamInterface|null $body = null ][, string $version = '1.1' ][, string|null $reason = null ]) : mixed`

##### Parameters

$status
: int
= 200

Status code

$headers
: array<string\|int, string\|array<string\|int, string>>
= \[\]

Response headers

$body
: string\|resource\| [StreamInterface](class-psr-http-message-streaminterface.md) \|null
= null

Response body

$version
: string
= '1.1'

Protocol version

$reason
: string\|null
= null

Reason phrase (when empty a default will be used based on the status code)

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

#### getProtocolVersion()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html\#method_getProtocolVersion)

`
    public
                    getProtocolVersion() : string`

##### Return values

string

#### getReasonPhrase()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Response.html\#method_getReasonPhrase)

Gets the response reason phrase associated with the status code.

`
    public
                    getReasonPhrase() : string`

Because a reason phrase is not a required element in a response
status line, the reason phrase value MAY be null. Implementations MAY
choose to return the default RFC 7231 recommended reason phrase (or those
listed in the IANA HTTP Status Code Registry) for the response's
status code.

##### Return values

string
—

Reason phrase; must return an empty string if none present.

#### getStatusCode()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Response.html\#method_getStatusCode)

Gets the response status code.

`
    public
                    getStatusCode() : int`

The status code is a 3-digit integer result code of the server's attempt
to understand and satisfy the request.

##### Return values

int
—

Status code.

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

#### withStatus()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Response.html\#method_withStatus)

Return an instance with the specified status code and, optionally, reason phrase.

`
    public
                    withStatus(mixed $code[, mixed $reasonPhrase = '' ]) : static`

If no reason phrase is specified, implementations MAY choose to default
to the RFC 7231 or IANA recommended reason phrase for the response's
status code.

This method MUST be implemented in such a way as to retain the
immutability of the message, and MUST return an instance that has the
updated status and reason phrase.

##### Parameters

$code
: mixed

The 3-digit integer result code to set.

$reasonPhrase
: mixed
= ''

The reason phrase to use with the
provided status code; if none is provided, implementations MAY
use the defaults as suggested in the HTTP specification.

##### Return values

static
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Response.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Response.html#method___construct)
  - [getBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getBody)
  - [getHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getHeader)
  - [getHeaderLine()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getHeaderLine)
  - [getHeaders()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getHeaders)
  - [getProtocolVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getProtocolVersion)
  - [getReasonPhrase()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Response.html#method_getReasonPhrase)
  - [getStatusCode()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Response.html#method_getStatusCode)
  - [hasHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_hasHeader)
  - [withAddedHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withAddedHeader)
  - [withBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withBody)
  - [withHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withHeader)
  - [withoutHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withoutHeader)
  - [withProtocolVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withProtocolVersion)
  - [withStatus()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Response.html#method_withStatus)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Response.html#top)
