Menu

- [GuzzleHttp](namespace-guzzlehttp.md)
- [Psr7](namespace-guzzlehttp-psr7.md)

## Message        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](class-guzzlehttp-psr7-message-toc.md)

#### Methods  [header link](class-guzzlehttp-psr7-message-toc-methods.md)

[bodySummary()](class-guzzlehttp-psr7-message-method-bodysummary.md)
: string\|null Get a short summary of the message body.[parseMessage()](class-guzzlehttp-psr7-message-method-parsemessage.md)
: array<string\|int, mixed> Parses an HTTP message into an associative array.[parseRequest()](class-guzzlehttp-psr7-message-method-parserequest.md)
: [RequestInterface](class-psr-http-message-requestinterface.md)Parses a request message string into a request object.[parseRequestUri()](class-guzzlehttp-psr7-message-method-parserequesturi.md)
: string Constructs a URI for an HTTP request message.[parseResponse()](class-guzzlehttp-psr7-message-method-parseresponse.md)
: [ResponseInterface](class-psr-http-message-responseinterface.md)Parses a response message string into a response object.[rewindBody()](class-guzzlehttp-psr7-message-method-rewindbody.md)
: void Attempts to rewind a message body and throws an exception on failure.[toString()](class-guzzlehttp-psr7-message-method-tostring.md)
: string Returns the string representation of an HTTP message.

### Methods  [header link](class-guzzlehttp-psr7-message-methods.md)

#### bodySummary()  [header link](class-guzzlehttp-psr7-message-method-bodysummary.md)

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

#### parseMessage()  [header link](class-guzzlehttp-psr7-message-method-parsemessage.md)

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

#### parseRequest()  [header link](class-guzzlehttp-psr7-message-method-parserequest.md)

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

#### parseRequestUri()  [header link](class-guzzlehttp-psr7-message-method-parserequesturi.md)

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

#### parseResponse()  [header link](class-guzzlehttp-psr7-message-method-parseresponse.md)

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

#### rewindBody()  [header link](class-guzzlehttp-psr7-message-method-rewindbody.md)

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

##### Tags  [header link](class-guzzlehttp-psr7-message-method-rewindbody-tags.md)

throwsRuntimeException

#### toString()  [header link](class-guzzlehttp-psr7-message-method-tostring.md)

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
  - [Methods](class-guzzlehttp-psr7-message-toc-methods.md)
- Methods
  - [bodySummary()](class-guzzlehttp-psr7-message-method-bodysummary.md)
  - [parseMessage()](class-guzzlehttp-psr7-message-method-parsemessage.md)
  - [parseRequest()](class-guzzlehttp-psr7-message-method-parserequest.md)
  - [parseRequestUri()](class-guzzlehttp-psr7-message-method-parserequesturi.md)
  - [parseResponse()](class-guzzlehttp-psr7-message-method-parseresponse.md)
  - [rewindBody()](class-guzzlehttp-psr7-message-method-rewindbody.md)
  - [toString()](class-guzzlehttp-psr7-message-method-tostring.md)

[Back To Top](class-guzzlehttp-psr7-message-top.md)
