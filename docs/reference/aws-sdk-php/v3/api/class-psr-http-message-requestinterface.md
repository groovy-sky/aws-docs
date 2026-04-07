Menu

- [Psr](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Psr.html)
- [Http](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Psr.http.html)
- [Message](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Psr.http.message.html)

## RequestInterface    extends  [MessageInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html)   in    - [Aws](package-aws.md)

Representation of an outgoing, client-side request.

Per the HTTP specification, this interface includes properties for
each of the following:

- Protocol version
- HTTP method
- URI
- Headers
- Message body

During construction, implementations MUST attempt to set the Host header from
a provided URI if no Host header is provided.

Requests are considered immutable; all methods that might change state MUST
be implemented such that they retain the internal state of the current
message and return an instance that contains the changed state.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html\#toc-methods)

[getBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_getBody)
: [StreamInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.StreamInterface.html)Gets the body of the message.[getHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_getHeader)
: array<string\|int, string> Retrieves a message header value by the given case-insensitive name.[getHeaderLine()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_getHeaderLine)
: string Retrieves a comma-separated string of the values for a single header.[getHeaders()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_getHeaders)
: array<string\|int, array<string\|int, string>> Retrieves all message header values.[getMethod()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html#method_getMethod)
: string Retrieves the HTTP method of the request.[getProtocolVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_getProtocolVersion)
: string Retrieves the HTTP protocol version as a string.[getRequestTarget()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html#method_getRequestTarget)
: string Retrieves the message's request target.[getUri()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html#method_getUri)
: [UriInterface](class-psr-http-message-uriinterface.md)Retrieves the URI instance.[hasHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_hasHeader)
: bool Checks if a header exists by the given case-insensitive name.[withAddedHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_withAddedHeader)
: static Return an instance with the specified header appended with the given value.[withBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_withBody)
: static Return an instance with the specified message body.[withHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_withHeader)
: static Return an instance with the provided value replacing the specified header.[withMethod()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html#method_withMethod)
: static Return an instance with the provided HTTP method.[withoutHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_withoutHeader)
: static Return an instance without the specified header.[withProtocolVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_withProtocolVersion)
: static Return an instance with the specified HTTP protocol version.[withRequestTarget()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html#method_withRequestTarget)
: static Return an instance with the specific request-target.[withUri()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html#method_withUri)
: static Returns an instance with the provided URI.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html\#methods)

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

#### getMethod()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html\#method_getMethod)

Retrieves the HTTP method of the request.

`
    public
                    getMethod() : string`

##### Return values

string
—

Returns the request method.

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

#### getRequestTarget()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html\#method_getRequestTarget)

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

#### getUri()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html\#method_getUri)

Retrieves the URI instance.

`
    public
                    getUri() : UriInterface`

This method MUST return a UriInterface instance.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html\#method_getUri\#tags)

link[http://tools.ietf.org/html/rfc3986#section-4.3](http://tools.ietf.org/html/rfc3986)

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)
—

Returns a UriInterface instance
representing the URI of the request.

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

#### withMethod()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html\#method_withMethod)

Return an instance with the provided HTTP method.

`
    public
                    withMethod(string $method) : static`

While HTTP method names are typically all uppercase characters, HTTP
method names are case-sensitive and thus implementations SHOULD NOT
modify the given string.

This method MUST be implemented in such a way as to retain the
immutability of the message, and MUST return an instance that has the
changed request method.

##### Parameters

$method
: string

Case-sensitive method.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html\#method_withMethod\#tags)

throwsInvalidArgumentException

for invalid HTTP methods.

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

#### withRequestTarget()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html\#method_withRequestTarget)

Return an instance with the specific request-target.

`
    public
                    withRequestTarget(string $requestTarget) : static`

If the request needs a non-origin-form request-target — e.g., for
specifying an absolute-form, authority-form, or asterisk-form —
this method may be used to create an instance with the specified
request-target, verbatim.

This method MUST be implemented in such a way as to retain the
immutability of the message, and MUST return an instance that has the
changed request target.

##### Parameters

$requestTarget
: string

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html\#method_withRequestTarget\#tags)

link[(for the various\
request-target forms allowed in request messages)](http://tools.ietf.org/html/rfc7230)

##### Return values

static

#### withUri()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html\#method_withUri)

Returns an instance with the provided URI.

`
    public
                    withUri(UriInterface $uri[, bool $preserveHost = false ]) : static`

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
: bool
= false

Preserve the original state of the Host header.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html\#method_withUri\#tags)

link[http://tools.ietf.org/html/rfc3986#section-4.3](http://tools.ietf.org/html/rfc3986)

##### Return values

static
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html#toc-methods)
- Methods
  - [getBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_getBody)
  - [getHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_getHeader)
  - [getHeaderLine()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_getHeaderLine)
  - [getHeaders()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_getHeaders)
  - [getMethod()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html#method_getMethod)
  - [getProtocolVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_getProtocolVersion)
  - [getRequestTarget()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html#method_getRequestTarget)
  - [getUri()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html#method_getUri)
  - [hasHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_hasHeader)
  - [withAddedHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_withAddedHeader)
  - [withBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_withBody)
  - [withHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_withHeader)
  - [withMethod()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html#method_withMethod)
  - [withoutHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_withoutHeader)
  - [withProtocolVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.MessageInterface.html#method_withProtocolVersion)
  - [withRequestTarget()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html#method_withRequestTarget)
  - [withUri()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html#method_withUri)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html#top)
