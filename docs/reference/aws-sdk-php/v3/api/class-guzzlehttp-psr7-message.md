Menu

- [GuzzleHttp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.html)
- [Psr7](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.psr7.html)

## Message        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Message.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Message.html\#toc-methods)

[bodySummary()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Message.html#method_bodySummary)
: string\|null Get a short summary of the message body.[parseMessage()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Message.html#method_parseMessage)
: array<string\|int, mixed> Parses an HTTP message into an associative array.[parseRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Message.html#method_parseRequest)
: [RequestInterface](class-psr-http-message-requestinterface.md)Parses a request message string into a request object.[parseRequestUri()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Message.html#method_parseRequestUri)
: string Constructs a URI for an HTTP request message.[parseResponse()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Message.html#method_parseResponse)
: [ResponseInterface](class-psr-http-message-responseinterface.md)Parses a response message string into a response object.[rewindBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Message.html#method_rewindBody)
: void Attempts to rewind a message body and throws an exception on failure.[toString()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Message.html#method_toString)
: string Returns the string representation of an HTTP message.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Message.html\#methods)

#### bodySummary()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Message.html\#method_bodySummary)

Get a short summary of the message body.

`
    public
            static        bodySummary(MessageInterface $message[, int $truncateAt = 120 ]) : string|null`

Will return `null` if the response is not printable.

##### Parameters

$message
: [MessageInterface](class-psr-http-message-messageinterface.md)

The message to get the body summary

$truncateAt
: int
= 120

The maximum allowed size of the summary

##### Return values

string\|null

#### parseMessage()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Message.html\#method_parseMessage)

Parses an HTTP message into an associative array.

`
    public
            static        parseMessage(string $message) : array<string|int, mixed>`

The array contains the "start-line" key containing the start line of
the message, "headers" key containing an associative array of header
array values, and a "body" key containing the body of the message.

##### Parameters

$message
: string

HTTP request or response to parse.

##### Return values

array<string\|int, mixed>

#### parseRequest()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Message.html\#method_parseRequest)

Parses a request message string into a request object.

`
    public
            static        parseRequest(string $message) : RequestInterface`

##### Parameters

$message
: string

Request message string.

##### Return values

[RequestInterface](class-psr-http-message-requestinterface.md)

#### parseRequestUri()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Message.html\#method_parseRequestUri)

Constructs a URI for an HTTP request message.

`
    public
            static        parseRequestUri(string $path, array<string|int, mixed> $headers) : string`

##### Parameters

$path
: string

Path from the start-line

$headers
: array<string\|int, mixed>

Array of headers (each value an array).

##### Return values

string

#### parseResponse()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Message.html\#method_parseResponse)

Parses a response message string into a response object.

`
    public
            static        parseResponse(string $message) : ResponseInterface`

##### Parameters

$message
: string

Response message string.

##### Return values

[ResponseInterface](class-psr-http-message-responseinterface.md)

#### rewindBody()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Message.html\#method_rewindBody)

Attempts to rewind a message body and throws an exception on failure.

`
    public
            static        rewindBody(MessageInterface $message) : void`

The body of the message will only be rewound if a call to `tell()`
returns a value other than `0`.

##### Parameters

$message
: [MessageInterface](class-psr-http-message-messageinterface.md)

Message to rewind

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Message.html\#method_rewindBody\#tags)

throwsRuntimeException

#### toString()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Message.html\#method_toString)

Returns the string representation of an HTTP message.

`
    public
            static        toString(MessageInterface $message) : string`

##### Parameters

$message
: [MessageInterface](class-psr-http-message-messageinterface.md)

Message to convert to a string.

##### Return values

string
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Message.html#toc-methods)
- Methods
  - [bodySummary()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Message.html#method_bodySummary)
  - [parseMessage()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Message.html#method_parseMessage)
  - [parseRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Message.html#method_parseRequest)
  - [parseRequestUri()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Message.html#method_parseRequestUri)
  - [parseResponse()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Message.html#method_parseResponse)
  - [rewindBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Message.html#method_rewindBody)
  - [toString()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Message.html#method_toString)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Message.html#top)
