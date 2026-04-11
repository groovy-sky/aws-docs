Menu

- [GuzzleHttp](namespace-guzzlehttp.md)
- [Psr7](namespace-guzzlehttp-psr7.md)

## Response        in package    - [Aws](package-aws.md)       implements  [ResponseInterface](class-psr-http-message-responseinterface.md)  Uses  [MessageTrait](class-guzzlehttp-psr7-messagetrait.md)

PSR-7 response implementation.

### Table of Contents  [header link](class-guzzlehttp-psr7-response-toc.md)

#### Interfaces  [header link](class-guzzlehttp-psr7-response-toc-interfaces.md)

[ResponseInterface](class-psr-http-message-responseinterface.md)Representation of an outgoing, server-side response.

#### Methods  [header link](class-guzzlehttp-psr7-response-toc-methods.md)

[\_\_construct()](class-guzzlehttp-psr7-response-method-construct.md)
: mixed [getBody()](class-guzzlehttp-psr7-messagetrait-method-getbody.md)
: [StreamInterface](class-psr-http-message-streaminterface.md)[getHeader()](class-guzzlehttp-psr7-messagetrait-method-getheader.md)
: array<string\|int, mixed> [getHeaderLine()](class-guzzlehttp-psr7-messagetrait-method-getheaderline.md)
: string [getHeaders()](class-guzzlehttp-psr7-messagetrait-method-getheaders.md)
: array<string\|int, mixed> [getProtocolVersion()](class-guzzlehttp-psr7-messagetrait-method-getprotocolversion.md)
: string [getReasonPhrase()](class-guzzlehttp-psr7-response-method-getreasonphrase.md)
: string Gets the response reason phrase associated with the status code.[getStatusCode()](class-guzzlehttp-psr7-response-method-getstatuscode.md)
: int Gets the response status code.[hasHeader()](class-guzzlehttp-psr7-messagetrait-method-hasheader.md)
: bool [withAddedHeader()](class-guzzlehttp-psr7-messagetrait-method-withaddedheader.md)
: static [withBody()](class-guzzlehttp-psr7-messagetrait-method-withbody.md)
: static [withHeader()](class-guzzlehttp-psr7-messagetrait-method-withheader.md)
: static [withoutHeader()](class-guzzlehttp-psr7-messagetrait-method-withoutheader.md)
: static [withProtocolVersion()](class-guzzlehttp-psr7-messagetrait-method-withprotocolversion.md)
: static [withStatus()](class-guzzlehttp-psr7-response-method-withstatus.md)
: static Return an instance with the specified status code and, optionally, reason phrase.

### Methods  [header link](class-guzzlehttp-psr7-response-methods.md)

#### \_\_construct()  [header link](class-guzzlehttp-psr7-response-method-construct.md)

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

#### getProtocolVersion()  [header link](class-guzzlehttp-psr7-messagetrait-method-getprotocolversion.md)

`
    public
                    getProtocolVersion() : string`

##### Return values

string

#### getReasonPhrase()  [header link](class-guzzlehttp-psr7-response-method-getreasonphrase.md)

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

#### getStatusCode()  [header link](class-guzzlehttp-psr7-response-method-getstatuscode.md)

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

#### withStatus()  [header link](class-guzzlehttp-psr7-response-method-withstatus.md)

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
  - [Methods](class-guzzlehttp-psr7-response-toc-methods.md)
- Methods
  - [\_\_construct()](class-guzzlehttp-psr7-response-method-construct.md)
  - [getBody()](class-guzzlehttp-psr7-messagetrait-method-getbody.md)
  - [getHeader()](class-guzzlehttp-psr7-messagetrait-method-getheader.md)
  - [getHeaderLine()](class-guzzlehttp-psr7-messagetrait-method-getheaderline.md)
  - [getHeaders()](class-guzzlehttp-psr7-messagetrait-method-getheaders.md)
  - [getProtocolVersion()](class-guzzlehttp-psr7-messagetrait-method-getprotocolversion.md)
  - [getReasonPhrase()](class-guzzlehttp-psr7-response-method-getreasonphrase.md)
  - [getStatusCode()](class-guzzlehttp-psr7-response-method-getstatuscode.md)
  - [hasHeader()](class-guzzlehttp-psr7-messagetrait-method-hasheader.md)
  - [withAddedHeader()](class-guzzlehttp-psr7-messagetrait-method-withaddedheader.md)
  - [withBody()](class-guzzlehttp-psr7-messagetrait-method-withbody.md)
  - [withHeader()](class-guzzlehttp-psr7-messagetrait-method-withheader.md)
  - [withoutHeader()](class-guzzlehttp-psr7-messagetrait-method-withoutheader.md)
  - [withProtocolVersion()](class-guzzlehttp-psr7-messagetrait-method-withprotocolversion.md)
  - [withStatus()](class-guzzlehttp-psr7-response-method-withstatus.md)

[Back To Top](class-guzzlehttp-psr7-response-top.md)
