Menu

- [Psr](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Psr.html)
- [Http](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Psr.http.html)
- [Message](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Psr.http.message.html)

## MessageInterface     in    - [Aws](package-aws.md)

HTTP messages consist of requests from a client to a server and responses
from a server to a client. This interface defines the methods common to
each.

Messages are considered immutable; all methods that might change state MUST
be implemented such that they retain the internal state of the current
message and return an instance that contains the changed state.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html\#tags)

link[http://www.ietf.org/rfc/rfc7230.txt](http://www.ietf.org/rfc/rfc7230.txt)link[http://www.ietf.org/rfc/rfc7231.txt](http://www.ietf.org/rfc/rfc7231.txt)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html\#toc-methods)

[getBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_getBody)
: [StreamInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html)Gets the body of the message.[getHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_getHeader)
: array<string\|int, string> Retrieves a message header value by the given case-insensitive name.[getHeaderLine()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_getHeaderLine)
: string Retrieves a comma-separated string of the values for a single header.[getHeaders()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_getHeaders)
: array<string\|int, array<string\|int, string>> Retrieves all message header values.[getProtocolVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_getProtocolVersion)
: string Retrieves the HTTP protocol version as a string.[hasHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_hasHeader)
: bool Checks if a header exists by the given case-insensitive name.[withAddedHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_withAddedHeader)
: static Return an instance with the specified header appended with the given value.[withBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_withBody)
: static Return an instance with the specified message body.[withHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_withHeader)
: static Return an instance with the provided value replacing the specified header.[withoutHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_withoutHeader)
: static Return an instance without the specified header.[withProtocolVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_withProtocolVersion)
: static Return an instance with the specified HTTP protocol version.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html\#methods)

#### getBody()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html\#method_getBody)

Gets the body of the message.

`
    public
                    getBody() : StreamInterface`

##### Return values

[StreamInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html)
—

Returns the body as a stream.

#### getHeader()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html\#method_getHeader)

Retrieves a message header value by the given case-insensitive name.

`
    public
                    getHeader(string $name) : array<string|int, string>`

This method returns an array of all the header values of the given
case-insensitive header name.

If the header does not appear in the message, this method MUST return an
empty array.

##### Parameters

$name
: string

Case-insensitive header field name.

##### Return values

array<string\|int, string>
—

An array of string values as provided for the given
header. If the header does not appear in the message, this method MUST
return an empty array.

#### getHeaderLine()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html\#method_getHeaderLine)

Retrieves a comma-separated string of the values for a single header.

`
    public
                    getHeaderLine(string $name) : string`

This method returns all of the header values of the given
case-insensitive header name as a string concatenated together using
a comma.

NOTE: Not all header values may be appropriately represented using
comma concatenation. For such headers, use getHeader() instead
and supply your own delimiter when concatenating.

If the header does not appear in the message, this method MUST return
an empty string.

##### Parameters

$name
: string

Case-insensitive header field name.

##### Return values

string
—

A string of values as provided for the given header
concatenated together using a comma. If the header does not appear in
the message, this method MUST return an empty string.

#### getHeaders()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html\#method_getHeaders)

Retrieves all message header values.

`
    public
                    getHeaders() : array<string|int, array<string|int, string>>`

The keys represent the header name as it will be sent over the wire, and
each value is an array of strings associated with the header.

```prettyprint
// Represent the headers as a string
foreach ($message->getHeaders() as $name => $values) {
    echo $name . ": " . implode(", ", $values);
}

// Emit headers iteratively:
foreach ($message->getHeaders() as $name => $values) {
    foreach ($values as $value) {
        header(sprintf('%s: %s', $name, $value), false);
    }
}

```

While header names are not case-sensitive, getHeaders() will preserve the
exact case in which headers were originally specified.

##### Return values

array<string\|int, array<string\|int, string>>
—

Returns an associative array of the message's headers. Each
key MUST be a header name, and each value MUST be an array of strings
for that header.

#### getProtocolVersion()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html\#method_getProtocolVersion)

Retrieves the HTTP protocol version as a string.

`
    public
                    getProtocolVersion() : string`

The string MUST contain only the HTTP version number (e.g., "1.1", "1.0").

##### Return values

string
—

HTTP protocol version.

#### hasHeader()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html\#method_hasHeader)

Checks if a header exists by the given case-insensitive name.

`
    public
                    hasHeader(string $name) : bool`

##### Parameters

$name
: string

Case-insensitive header field name.

##### Return values

bool
—

Returns true if any header names match the given header
name using a case-insensitive string comparison. Returns false if
no matching header name is found in the message.

#### withAddedHeader()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html\#method_withAddedHeader)

Return an instance with the specified header appended with the given value.

`
    public
                    withAddedHeader(string $name, string|array<string|int, string> $value) : static`

Existing values for the specified header will be maintained. The new
value(s) will be appended to the existing list. If the header did not
exist previously, it will be added.

This method MUST be implemented in such a way as to retain the
immutability of the message, and MUST return an instance that has the
new header and/or value.

##### Parameters

$name
: string

Case-insensitive header field name to add.

$value
: string\|array<string\|int, string>

Header value(s).

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html\#method_withAddedHeader\#tags)

throwsInvalidArgumentException

for invalid header names or values.

##### Return values

static

#### withBody()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html\#method_withBody)

Return an instance with the specified message body.

`
    public
                    withBody(StreamInterface $body) : static`

The body MUST be a StreamInterface object.

This method MUST be implemented in such a way as to retain the
immutability of the message, and MUST return a new instance that has the
new body stream.

##### Parameters

$body
: [StreamInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html)

Body.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html\#method_withBody\#tags)

throwsInvalidArgumentException

When the body is not valid.

##### Return values

static

#### withHeader()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html\#method_withHeader)

Return an instance with the provided value replacing the specified header.

`
    public
                    withHeader(string $name, string|array<string|int, string> $value) : static`

While header names are case-insensitive, the casing of the header will
be preserved by this function, and returned from getHeaders().

This method MUST be implemented in such a way as to retain the
immutability of the message, and MUST return an instance that has the
new and/or updated header and value.

##### Parameters

$name
: string

Case-insensitive header field name.

$value
: string\|array<string\|int, string>

Header value(s).

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html\#method_withHeader\#tags)

throwsInvalidArgumentException

for invalid header names or values.

##### Return values

static

#### withoutHeader()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html\#method_withoutHeader)

Return an instance without the specified header.

`
    public
                    withoutHeader(string $name) : static`

Header resolution MUST be done without case-sensitivity.

This method MUST be implemented in such a way as to retain the
immutability of the message, and MUST return an instance that removes
the named header.

##### Parameters

$name
: string

Case-insensitive header field name to remove.

##### Return values

static

#### withProtocolVersion()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html\#method_withProtocolVersion)

Return an instance with the specified HTTP protocol version.

`
    public
                    withProtocolVersion(string $version) : static`

The version string MUST contain only the HTTP version number (e.g.,
"1.1", "1.0").

This method MUST be implemented in such a way as to retain the
immutability of the message, and MUST return an instance that has the
new protocol version.

##### Parameters

$version
: string

HTTP protocol version

##### Return values

static
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#toc-methods)
- Methods
  - [getBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_getBody)
  - [getHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_getHeader)
  - [getHeaderLine()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_getHeaderLine)
  - [getHeaders()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_getHeaders)
  - [getProtocolVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_getProtocolVersion)
  - [hasHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_hasHeader)
  - [withAddedHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_withAddedHeader)
  - [withBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_withBody)
  - [withHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_withHeader)
  - [withoutHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_withoutHeader)
  - [withProtocolVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_withProtocolVersion)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#top)
