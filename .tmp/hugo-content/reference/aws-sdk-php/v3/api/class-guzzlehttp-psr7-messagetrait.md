Menu

- [GuzzleHttp](namespace-guzzlehttp.md)
- [Psr7](namespace-guzzlehttp-psr7.md)

## MessageTrait

Trait implementing functionality common to requests and responses.

### Table of Contents  [header link](class-guzzlehttp-psr7-messagetrait-toc.md)

#### Methods  [header link](class-guzzlehttp-psr7-messagetrait-toc-methods.md)

[getBody()](class-guzzlehttp-psr7-messagetrait-method-getbody.md)
: [StreamInterface](class-psr-http-message-streaminterface.md)[getHeader()](class-guzzlehttp-psr7-messagetrait-method-getheader.md)
: array<string\|int, mixed> [getHeaderLine()](class-guzzlehttp-psr7-messagetrait-method-getheaderline.md)
: string [getHeaders()](class-guzzlehttp-psr7-messagetrait-method-getheaders.md)
: array<string\|int, mixed> [getProtocolVersion()](class-guzzlehttp-psr7-messagetrait-method-getprotocolversion.md)
: string [hasHeader()](class-guzzlehttp-psr7-messagetrait-method-hasheader.md)
: bool [withAddedHeader()](class-guzzlehttp-psr7-messagetrait-method-withaddedheader.md)
: static [withBody()](class-guzzlehttp-psr7-messagetrait-method-withbody.md)
: static [withHeader()](class-guzzlehttp-psr7-messagetrait-method-withheader.md)
: static [withoutHeader()](class-guzzlehttp-psr7-messagetrait-method-withoutheader.md)
: static [withProtocolVersion()](class-guzzlehttp-psr7-messagetrait-method-withprotocolversion.md)
: static

### Methods  [header link](class-guzzlehttp-psr7-messagetrait-methods.md)

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
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-guzzlehttp-psr7-messagetrait-toc-methods.md)
- Methods
  - [getBody()](class-guzzlehttp-psr7-messagetrait-method-getbody.md)
  - [getHeader()](class-guzzlehttp-psr7-messagetrait-method-getheader.md)
  - [getHeaderLine()](class-guzzlehttp-psr7-messagetrait-method-getheaderline.md)
  - [getHeaders()](class-guzzlehttp-psr7-messagetrait-method-getheaders.md)
  - [getProtocolVersion()](class-guzzlehttp-psr7-messagetrait-method-getprotocolversion.md)
  - [hasHeader()](class-guzzlehttp-psr7-messagetrait-method-hasheader.md)
  - [withAddedHeader()](class-guzzlehttp-psr7-messagetrait-method-withaddedheader.md)
  - [withBody()](class-guzzlehttp-psr7-messagetrait-method-withbody.md)
  - [withHeader()](class-guzzlehttp-psr7-messagetrait-method-withheader.md)
  - [withoutHeader()](class-guzzlehttp-psr7-messagetrait-method-withoutheader.md)
  - [withProtocolVersion()](class-guzzlehttp-psr7-messagetrait-method-withprotocolversion.md)

[Back To Top](class-guzzlehttp-psr7-messagetrait-top.md)
