Menu

- [Psr](namespace-psr.md)
- [Http](namespace-psr-http.md)
- [Message](namespace-psr-http-message.md)

## MessageInterface     in    - [Aws](package-aws.md)

HTTP messages consist of requests from a client to a server and responses
from a server to a client. This interface defines the methods common to
each.

Messages are considered immutable; all methods that might change state MUST
be implemented such that they retain the internal state of the current
message and return an instance that contains the changed state.

##### Tags  [header link](class-psr-http-message-messageinterface-tags.md)

link[http://www.ietf.org/rfc/rfc7230.txt](http://www.ietf.org/rfc/rfc7230.txt)link[http://www.ietf.org/rfc/rfc7231.txt](http://www.ietf.org/rfc/rfc7231.txt)

### Table of Contents  [header link](class-psr-http-message-messageinterface-toc.md)

#### Methods  [header link](class-psr-http-message-messageinterface-toc-methods.md)

[getBody()](class-psr-http-message-messageinterface-method-getbody.md)
: [StreamInterface](class-psr-http-message-streaminterface.md)Gets the body of the message.[getHeader()](class-psr-http-message-messageinterface-method-getheader.md)
: array<string\|int, string> Retrieves a message header value by the given case-insensitive name.[getHeaderLine()](class-psr-http-message-messageinterface-method-getheaderline.md)
: string Retrieves a comma-separated string of the values for a single header.[getHeaders()](class-psr-http-message-messageinterface-method-getheaders.md)
: array<string\|int, array<string\|int, string>> Retrieves all message header values.[getProtocolVersion()](class-psr-http-message-messageinterface-method-getprotocolversion.md)
: string Retrieves the HTTP protocol version as a string.[hasHeader()](class-psr-http-message-messageinterface-method-hasheader.md)
: bool Checks if a header exists by the given case-insensitive name.[withAddedHeader()](class-psr-http-message-messageinterface-method-withaddedheader.md)
: static Return an instance with the specified header appended with the given value.[withBody()](class-psr-http-message-messageinterface-method-withbody.md)
: static Return an instance with the specified message body.[withHeader()](class-psr-http-message-messageinterface-method-withheader.md)
: static Return an instance with the provided value replacing the specified header.[withoutHeader()](class-psr-http-message-messageinterface-method-withoutheader.md)
: static Return an instance without the specified header.[withProtocolVersion()](class-psr-http-message-messageinterface-method-withprotocolversion.md)
: static Return an instance with the specified HTTP protocol version.

### Methods  [header link](class-psr-http-message-messageinterface-methods.md)

#### getBody()  [header link](class-psr-http-message-messageinterface-method-getbody.md)

Gets the body of the message.

`
    public
                    getBody() : StreamInterface`

##### Return values

[StreamInterface](class-psr-http-message-streaminterface.md)
—

Returns the body as a stream.

#### getHeader()  [header link](class-psr-http-message-messageinterface-method-getheader.md)

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

#### getHeaderLine()  [header link](class-psr-http-message-messageinterface-method-getheaderline.md)

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

#### getHeaders()  [header link](class-psr-http-message-messageinterface-method-getheaders.md)

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

#### getProtocolVersion()  [header link](class-psr-http-message-messageinterface-method-getprotocolversion.md)

Retrieves the HTTP protocol version as a string.

`
    public
                    getProtocolVersion() : string`

The string MUST contain only the HTTP version number (e.g., "1.1", "1.0").

##### Return values

string
—

HTTP protocol version.

#### hasHeader()  [header link](class-psr-http-message-messageinterface-method-hasheader.md)

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

#### withAddedHeader()  [header link](class-psr-http-message-messageinterface-method-withaddedheader.md)

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

##### Tags  [header link](class-psr-http-message-messageinterface-method-withaddedheader-tags.md)

throwsInvalidArgumentException

for invalid header names or values.

##### Return values

static

#### withBody()  [header link](class-psr-http-message-messageinterface-method-withbody.md)

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
: [StreamInterface](class-psr-http-message-streaminterface.md)

Body.

##### Tags  [header link](class-psr-http-message-messageinterface-method-withbody-tags.md)

throwsInvalidArgumentException

When the body is not valid.

##### Return values

static

#### withHeader()  [header link](class-psr-http-message-messageinterface-method-withheader.md)

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

##### Tags  [header link](class-psr-http-message-messageinterface-method-withheader-tags.md)

throwsInvalidArgumentException

for invalid header names or values.

##### Return values

static

#### withoutHeader()  [header link](class-psr-http-message-messageinterface-method-withoutheader.md)

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

#### withProtocolVersion()  [header link](class-psr-http-message-messageinterface-method-withprotocolversion.md)

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
  - [Constants](class-psr-http-message-messageinterface-toc-constants.md)
  - [Methods](class-psr-http-message-messageinterface-toc-methods.md)
- Methods
  - [getBody()](class-psr-http-message-messageinterface-method-getbody.md)
  - [getHeader()](class-psr-http-message-messageinterface-method-getheader.md)
  - [getHeaderLine()](class-psr-http-message-messageinterface-method-getheaderline.md)
  - [getHeaders()](class-psr-http-message-messageinterface-method-getheaders.md)
  - [getProtocolVersion()](class-psr-http-message-messageinterface-method-getprotocolversion.md)
  - [hasHeader()](class-psr-http-message-messageinterface-method-hasheader.md)
  - [withAddedHeader()](class-psr-http-message-messageinterface-method-withaddedheader.md)
  - [withBody()](class-psr-http-message-messageinterface-method-withbody.md)
  - [withHeader()](class-psr-http-message-messageinterface-method-withheader.md)
  - [withoutHeader()](class-psr-http-message-messageinterface-method-withoutheader.md)
  - [withProtocolVersion()](class-psr-http-message-messageinterface-method-withprotocolversion.md)

[Back To Top](class-psr-http-message-messageinterface-top.md)
