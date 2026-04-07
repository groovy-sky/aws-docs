Menu

- [GuzzleHttp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.html)
- [Psr7](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.psr7.html)

## MessageTrait

Trait implementing functionality common to requests and responses.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html\#toc-methods)

[getBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getBody)
: [StreamInterface](class-psr-http-message-streaminterface.md)[getHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getHeader)
: array<string\|int, mixed> [getHeaderLine()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getHeaderLine)
: string [getHeaders()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getHeaders)
: array<string\|int, mixed> [getProtocolVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getProtocolVersion)
: string [hasHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_hasHeader)
: bool [withAddedHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withAddedHeader)
: static [withBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withBody)
: static [withHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withHeader)
: static [withoutHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withoutHeader)
: static [withProtocolVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withProtocolVersion)
: static

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html\#methods)

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
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#toc-methods)
- Methods
  - [getBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getBody)
  - [getHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getHeader)
  - [getHeaderLine()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getHeaderLine)
  - [getHeaders()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getHeaders)
  - [getProtocolVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getProtocolVersion)
  - [hasHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_hasHeader)
  - [withAddedHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withAddedHeader)
  - [withBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withBody)
  - [withHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withHeader)
  - [withoutHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withoutHeader)
  - [withProtocolVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withProtocolVersion)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#top)
