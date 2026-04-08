Menu

- [Psr](namespace-psr.md)
- [Http](namespace-psr-http.md)
- [Message](namespace-psr-http-message.md)

## RequestInterface    extends  [MessageInterface](class-psr-http-message-messageinterface.md)   in    - [Aws](package-aws.md)

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

### Table of Contents  [header link](class-psr-http-message-requestinterface-toc.md)

#### Methods  [header link](class-psr-http-message-requestinterface-toc-methods.md)

[getBody()](class-psr-http-message-messageinterface-method-getbody.md)
: [StreamInterface](class-psr-http-message-streaminterface.md)Gets the body of the message.[getHeader()](class-psr-http-message-messageinterface-method-getheader.md)
: array<string\|int, string> Retrieves a message header value by the given case-insensitive name.[getHeaderLine()](class-psr-http-message-messageinterface-method-getheaderline.md)
: string Retrieves a comma-separated string of the values for a single header.[getHeaders()](class-psr-http-message-messageinterface-method-getheaders.md)
: array<string\|int, array<string\|int, string>> Retrieves all message header values.[getMethod()](class-psr-http-message-requestinterface-method-getmethod.md)
: string Retrieves the HTTP method of the request.[getProtocolVersion()](class-psr-http-message-messageinterface-method-getprotocolversion.md)
: string Retrieves the HTTP protocol version as a string.[getRequestTarget()](class-psr-http-message-requestinterface-method-getrequesttarget.md)
: string Retrieves the message's request target.[getUri()](class-psr-http-message-requestinterface-method-geturi.md)
: [UriInterface](class-psr-http-message-uriinterface.md)Retrieves the URI instance.[hasHeader()](class-psr-http-message-messageinterface-method-hasheader.md)
: bool Checks if a header exists by the given case-insensitive name.[withAddedHeader()](class-psr-http-message-messageinterface-method-withaddedheader.md)
: static Return an instance with the specified header appended with the given value.[withBody()](class-psr-http-message-messageinterface-method-withbody.md)
: static Return an instance with the specified message body.[withHeader()](class-psr-http-message-messageinterface-method-withheader.md)
: static Return an instance with the provided value replacing the specified header.[withMethod()](class-psr-http-message-requestinterface-method-withmethod.md)
: static Return an instance with the provided HTTP method.[withoutHeader()](class-psr-http-message-messageinterface-method-withoutheader.md)
: static Return an instance without the specified header.[withProtocolVersion()](class-psr-http-message-messageinterface-method-withprotocolversion.md)
: static Return an instance with the specified HTTP protocol version.[withRequestTarget()](class-psr-http-message-requestinterface-method-withrequesttarget.md)
: static Return an instance with the specific request-target.[withUri()](class-psr-http-message-requestinterface-method-withuri.md)
: static Returns an instance with the provided URI.

### Methods  [header link](class-psr-http-message-requestinterface-methods.md)

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

#### getMethod()  [header link](class-psr-http-message-requestinterface-method-getmethod.md)

Retrieves the HTTP method of the request.

`
    public
                    getMethod() : string`

##### Return values

string
—

Returns the request method.

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

#### getRequestTarget()  [header link](class-psr-http-message-requestinterface-method-getrequesttarget.md)

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

#### getUri()  [header link](class-psr-http-message-requestinterface-method-geturi.md)

Retrieves the URI instance.

`
    public
                    getUri() : UriInterface`

This method MUST return a UriInterface instance.

##### Tags  [header link](class-psr-http-message-requestinterface-method-geturi-tags.md)

link[http://tools.ietf.org/html/rfc3986#section-4.3](http://tools.ietf.org/html/rfc3986)

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)
—

Returns a UriInterface instance
representing the URI of the request.

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

#### withMethod()  [header link](class-psr-http-message-requestinterface-method-withmethod.md)

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

##### Tags  [header link](class-psr-http-message-requestinterface-method-withmethod-tags.md)

throwsInvalidArgumentException

for invalid HTTP methods.

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

#### withRequestTarget()  [header link](class-psr-http-message-requestinterface-method-withrequesttarget.md)

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

##### Tags  [header link](class-psr-http-message-requestinterface-method-withrequesttarget-tags.md)

link[(for the various\
request-target forms allowed in request messages)](http://tools.ietf.org/html/rfc7230)

##### Return values

static

#### withUri()  [header link](class-psr-http-message-requestinterface-method-withuri.md)

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

##### Tags  [header link](class-psr-http-message-requestinterface-method-withuri-tags.md)

link[http://tools.ietf.org/html/rfc3986#section-4.3](http://tools.ietf.org/html/rfc3986)

##### Return values

static
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-psr-http-message-requestinterface-toc-constants.md)
  - [Methods](class-psr-http-message-requestinterface-toc-methods.md)
- Methods
  - [getBody()](class-psr-http-message-messageinterface-method-getbody.md)
  - [getHeader()](class-psr-http-message-messageinterface-method-getheader.md)
  - [getHeaderLine()](class-psr-http-message-messageinterface-method-getheaderline.md)
  - [getHeaders()](class-psr-http-message-messageinterface-method-getheaders.md)
  - [getMethod()](class-psr-http-message-requestinterface-method-getmethod.md)
  - [getProtocolVersion()](class-psr-http-message-messageinterface-method-getprotocolversion.md)
  - [getRequestTarget()](class-psr-http-message-requestinterface-method-getrequesttarget.md)
  - [getUri()](class-psr-http-message-requestinterface-method-geturi.md)
  - [hasHeader()](class-psr-http-message-messageinterface-method-hasheader.md)
  - [withAddedHeader()](class-psr-http-message-messageinterface-method-withaddedheader.md)
  - [withBody()](class-psr-http-message-messageinterface-method-withbody.md)
  - [withHeader()](class-psr-http-message-messageinterface-method-withheader.md)
  - [withMethod()](class-psr-http-message-requestinterface-method-withmethod.md)
  - [withoutHeader()](class-psr-http-message-messageinterface-method-withoutheader.md)
  - [withProtocolVersion()](class-psr-http-message-messageinterface-method-withprotocolversion.md)
  - [withRequestTarget()](class-psr-http-message-requestinterface-method-withrequesttarget.md)
  - [withUri()](class-psr-http-message-requestinterface-method-withuri.md)

[Back To Top](class-psr-http-message-requestinterface-top.md)
