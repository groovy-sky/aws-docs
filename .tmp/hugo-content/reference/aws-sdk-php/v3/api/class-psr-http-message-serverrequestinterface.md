Menu

- [Psr](namespace-psr.md)
- [Http](namespace-psr-http.md)
- [Message](namespace-psr-http-message.md)

## ServerRequestInterface    extends  [RequestInterface](class-psr-http-message-requestinterface.md)   in    - [Aws](package-aws.md)

Representation of an incoming, server-side HTTP request.

Per the HTTP specification, this interface includes properties for
each of the following:

- Protocol version
- HTTP method
- URI
- Headers
- Message body

Additionally, it encapsulates all data as it has arrived to the
application from the CGI and/or PHP environment, including:

- The values represented in $\_SERVER.
- Any cookies provided (generally via $\_COOKIE)
- Query string arguments (generally via $\_GET, or as parsed via parse\_str())
- Upload files, if any (as represented by $\_FILES)
- Deserialized body parameters (generally from $\_POST)

$\_SERVER values MUST be treated as immutable, as they represent application
state at the time of request; as such, no methods are provided to allow
modification of those values. The other values provide such methods, as they
can be restored from $\_SERVER or the request body, and may need treatment
during the application (e.g., body parameters may be deserialized based on
content type).

Additionally, this interface recognizes the utility of introspecting a
request to derive and match additional parameters (e.g., via URI path
matching, decrypting cookie values, deserializing non-form-encoded body
content, matching authorization headers to users, etc). These parameters
are stored in an "attributes" property.

Requests are considered immutable; all methods that might change state MUST
be implemented such that they retain the internal state of the current
message and return an instance that contains the changed state.

### Table of Contents  [header link](class-psr-http-message-serverrequestinterface-toc.md)

#### Methods  [header link](class-psr-http-message-serverrequestinterface-toc-methods.md)

[getAttribute()](class-psr-http-message-serverrequestinterface-method-getattribute.md)
: mixed Retrieve a single derived request attribute.[getAttributes()](class-psr-http-message-serverrequestinterface-method-getattributes.md)
: array<string\|int, mixed> Retrieve attributes derived from the request.[getBody()](class-psr-http-message-messageinterface.md#method_getBody)
: [StreamInterface](class-psr-http-message-streaminterface.md)Gets the body of the message.[getCookieParams()](class-psr-http-message-serverrequestinterface-method-getcookieparams.md)
: array<string\|int, mixed> Retrieve cookies.[getHeader()](class-psr-http-message-messageinterface.md#method_getHeader)
: array<string\|int, string> Retrieves a message header value by the given case-insensitive name.[getHeaderLine()](class-psr-http-message-messageinterface.md#method_getHeaderLine)
: string Retrieves a comma-separated string of the values for a single header.[getHeaders()](class-psr-http-message-messageinterface.md#method_getHeaders)
: array<string\|int, array<string\|int, string>> Retrieves all message header values.[getMethod()](class-psr-http-message-requestinterface.md#method_getMethod)
: string Retrieves the HTTP method of the request.[getParsedBody()](class-psr-http-message-serverrequestinterface-method-getparsedbody.md)
: null\|array<string\|int, mixed>\|object Retrieve any parameters provided in the request body.[getProtocolVersion()](class-psr-http-message-messageinterface.md#method_getProtocolVersion)
: string Retrieves the HTTP protocol version as a string.[getQueryParams()](class-psr-http-message-serverrequestinterface-method-getqueryparams.md)
: array<string\|int, mixed> Retrieve query string arguments.[getRequestTarget()](class-psr-http-message-requestinterface.md#method_getRequestTarget)
: string Retrieves the message's request target.[getServerParams()](class-psr-http-message-serverrequestinterface-method-getserverparams.md)
: array<string\|int, mixed> Retrieve server parameters.[getUploadedFiles()](class-psr-http-message-serverrequestinterface-method-getuploadedfiles.md)
: array<string\|int, mixed> Retrieve normalized file upload data.[getUri()](class-psr-http-message-requestinterface.md#method_getUri)
: [UriInterface](class-psr-http-message-uriinterface.md)Retrieves the URI instance.[hasHeader()](class-psr-http-message-messageinterface.md#method_hasHeader)
: bool Checks if a header exists by the given case-insensitive name.[withAddedHeader()](class-psr-http-message-messageinterface.md#method_withAddedHeader)
: static Return an instance with the specified header appended with the given value.[withAttribute()](class-psr-http-message-serverrequestinterface-method-withattribute.md)
: static Return an instance with the specified derived request attribute.[withBody()](class-psr-http-message-messageinterface.md#method_withBody)
: static Return an instance with the specified message body.[withCookieParams()](class-psr-http-message-serverrequestinterface-method-withcookieparams.md)
: static Return an instance with the specified cookies.[withHeader()](class-psr-http-message-messageinterface.md#method_withHeader)
: static Return an instance with the provided value replacing the specified header.[withMethod()](class-psr-http-message-requestinterface.md#method_withMethod)
: static Return an instance with the provided HTTP method.[withoutAttribute()](class-psr-http-message-serverrequestinterface-method-withoutattribute.md)
: static Return an instance that removes the specified derived request attribute.[withoutHeader()](class-psr-http-message-messageinterface.md#method_withoutHeader)
: static Return an instance without the specified header.[withParsedBody()](class-psr-http-message-serverrequestinterface-method-withparsedbody.md)
: static Return an instance with the specified body parameters.[withProtocolVersion()](class-psr-http-message-messageinterface.md#method_withProtocolVersion)
: static Return an instance with the specified HTTP protocol version.[withQueryParams()](class-psr-http-message-serverrequestinterface-method-withqueryparams.md)
: static Return an instance with the specified query string arguments.[withRequestTarget()](class-psr-http-message-requestinterface.md#method_withRequestTarget)
: static Return an instance with the specific request-target.[withUploadedFiles()](class-psr-http-message-serverrequestinterface-method-withuploadedfiles.md)
: static Create a new instance with the specified uploaded files.[withUri()](class-psr-http-message-requestinterface.md#method_withUri)
: static Returns an instance with the provided URI.

### Methods  [header link](class-psr-http-message-serverrequestinterface-methods.md)

#### getAttribute()  [header link](class-psr-http-message-serverrequestinterface-method-getattribute.md)

Retrieve a single derived request attribute.

`
    public
                    getAttribute(string $name[, mixed $default = null ]) : mixed`

Retrieves a single derived request attribute as described in
getAttributes(). If the attribute has not been previously set, returns
the default value as provided.

This method obviates the need for a hasAttribute() method, as it allows
specifying a default value to return if the attribute is not found.

##### Parameters

$name
: string

The attribute name.

$default
: mixed
= null

Default value to return if the attribute does not exist.

##### Tags  [header link](class-psr-http-message-serverrequestinterface-method-getattribute-tags.md)

seegetAttributes()

#### getAttributes()  [header link](class-psr-http-message-serverrequestinterface-method-getattributes.md)

Retrieve attributes derived from the request.

`
    public
                    getAttributes() : array<string|int, mixed>`

The request "attributes" may be used to allow injection of any
parameters derived from the request: e.g., the results of path
match operations; the results of decrypting cookies; the results of
deserializing non-form-encoded message bodies; etc. Attributes
will be application and request specific, and CAN be mutable.

##### Return values

array<string\|int, mixed>
—

Attributes derived from the request.

#### getBody()  [header link](class-psr-http-message-messageinterface.md\#method_getBody)

Gets the body of the message.

`
    public
                    getBody() : StreamInterface`

##### Return values

[StreamInterface](class-psr-http-message-streaminterface.md)
—

Returns the body as a stream.

#### getCookieParams()  [header link](class-psr-http-message-serverrequestinterface-method-getcookieparams.md)

Retrieve cookies.

`
    public
                    getCookieParams() : array<string|int, mixed>`

Retrieves cookies sent by the client to the server.

The data MUST be compatible with the structure of the $\_COOKIE
superglobal.

##### Return values

array<string\|int, mixed>

#### getHeader()  [header link](class-psr-http-message-messageinterface.md\#method_getHeader)

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

#### getHeaderLine()  [header link](class-psr-http-message-messageinterface.md\#method_getHeaderLine)

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

#### getHeaders()  [header link](class-psr-http-message-messageinterface.md\#method_getHeaders)

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

#### getMethod()  [header link](class-psr-http-message-requestinterface.md\#method_getMethod)

Retrieves the HTTP method of the request.

`
    public
                    getMethod() : string`

##### Return values

string
—

Returns the request method.

#### getParsedBody()  [header link](class-psr-http-message-serverrequestinterface-method-getparsedbody.md)

Retrieve any parameters provided in the request body.

`
    public
                    getParsedBody() : null|array<string|int, mixed>|object`

If the request Content-Type is either application/x-www-form-urlencoded
or multipart/form-data, and the request method is POST, this method MUST
return the contents of $\_POST.

Otherwise, this method may return any results of deserializing
the request body content; as parsing returns structured content, the
potential types MUST be arrays or objects only. A null value indicates
the absence of body content.

##### Return values

null\|array<string\|int, mixed>\|object
—

The deserialized body parameters, if any.
These will typically be an array or object.

#### getProtocolVersion()  [header link](class-psr-http-message-messageinterface.md\#method_getProtocolVersion)

Retrieves the HTTP protocol version as a string.

`
    public
                    getProtocolVersion() : string`

The string MUST contain only the HTTP version number (e.g., "1.1", "1.0").

##### Return values

string
—

HTTP protocol version.

#### getQueryParams()  [header link](class-psr-http-message-serverrequestinterface-method-getqueryparams.md)

Retrieve query string arguments.

`
    public
                    getQueryParams() : array<string|int, mixed>`

Retrieves the deserialized query string arguments, if any.

Note: the query params might not be in sync with the URI or server
params. If you need to ensure you are only getting the original
values, you may need to parse the query string from `getUri()->getQuery()`
or from the `QUERY_STRING` server param.

##### Return values

array<string\|int, mixed>

#### getRequestTarget()  [header link](class-psr-http-message-requestinterface.md\#method_getRequestTarget)

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

#### getServerParams()  [header link](class-psr-http-message-serverrequestinterface-method-getserverparams.md)

Retrieve server parameters.

`
    public
                    getServerParams() : array<string|int, mixed>`

Retrieves data related to the incoming request environment,
typically derived from PHP's $\_SERVER superglobal. The data IS NOT
REQUIRED to originate from $\_SERVER.

##### Return values

array<string\|int, mixed>

#### getUploadedFiles()  [header link](class-psr-http-message-serverrequestinterface-method-getuploadedfiles.md)

Retrieve normalized file upload data.

`
    public
                    getUploadedFiles() : array<string|int, mixed>`

This method returns upload metadata in a normalized tree, with each leaf
an instance of Psr\\Http\\Message\\UploadedFileInterface.

These values MAY be prepared from $\_FILES or the message body during
instantiation, or MAY be injected via withUploadedFiles().

##### Return values

array<string\|int, mixed>
—

An array tree of UploadedFileInterface instances; an empty
array MUST be returned if no data is present.

#### getUri()  [header link](class-psr-http-message-requestinterface.md\#method_getUri)

Retrieves the URI instance.

`
    public
                    getUri() : UriInterface`

This method MUST return a UriInterface instance.

##### Tags  [header link](class-psr-http-message-requestinterface.md\#method_getUri\#tags)

link[http://tools.ietf.org/html/rfc3986#section-4.3](http://tools.ietf.org/html/rfc3986)

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)
—

Returns a UriInterface instance
representing the URI of the request.

#### hasHeader()  [header link](class-psr-http-message-messageinterface.md\#method_hasHeader)

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

#### withAddedHeader()  [header link](class-psr-http-message-messageinterface.md\#method_withAddedHeader)

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

##### Tags  [header link](class-psr-http-message-messageinterface.md\#method_withAddedHeader\#tags)

throwsInvalidArgumentException

for invalid header names or values.

##### Return values

static

#### withAttribute()  [header link](class-psr-http-message-serverrequestinterface-method-withattribute.md)

Return an instance with the specified derived request attribute.

`
    public
                    withAttribute(string $name, mixed $value) : static`

This method allows setting a single derived request attribute as
described in getAttributes().

This method MUST be implemented in such a way as to retain the
immutability of the message, and MUST return an instance that has the
updated attribute.

##### Parameters

$name
: string

The attribute name.

$value
: mixed

The value of the attribute.

##### Tags  [header link](class-psr-http-message-serverrequestinterface-method-withattribute-tags.md)

seegetAttributes()

##### Return values

static

#### withBody()  [header link](class-psr-http-message-messageinterface.md\#method_withBody)

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

##### Tags  [header link](class-psr-http-message-messageinterface.md\#method_withBody\#tags)

throwsInvalidArgumentException

When the body is not valid.

##### Return values

static

#### withCookieParams()  [header link](class-psr-http-message-serverrequestinterface-method-withcookieparams.md)

Return an instance with the specified cookies.

`
    public
                    withCookieParams(array<string|int, mixed> $cookies) : static`

The data IS NOT REQUIRED to come from the $\_COOKIE superglobal, but MUST
be compatible with the structure of $\_COOKIE. Typically, this data will
be injected at instantiation.

This method MUST NOT update the related Cookie header of the request
instance, nor related values in the server params.

This method MUST be implemented in such a way as to retain the
immutability of the message, and MUST return an instance that has the
updated cookie values.

##### Parameters

$cookies
: array<string\|int, mixed>

Array of key/value pairs representing cookies.

##### Return values

static

#### withHeader()  [header link](class-psr-http-message-messageinterface.md\#method_withHeader)

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

##### Tags  [header link](class-psr-http-message-messageinterface.md\#method_withHeader\#tags)

throwsInvalidArgumentException

for invalid header names or values.

##### Return values

static

#### withMethod()  [header link](class-psr-http-message-requestinterface.md\#method_withMethod)

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

##### Tags  [header link](class-psr-http-message-requestinterface.md\#method_withMethod\#tags)

throwsInvalidArgumentException

for invalid HTTP methods.

##### Return values

static

#### withoutAttribute()  [header link](class-psr-http-message-serverrequestinterface-method-withoutattribute.md)

Return an instance that removes the specified derived request attribute.

`
    public
                    withoutAttribute(string $name) : static`

This method allows removing a single derived request attribute as
described in getAttributes().

This method MUST be implemented in such a way as to retain the
immutability of the message, and MUST return an instance that removes
the attribute.

##### Parameters

$name
: string

The attribute name.

##### Tags  [header link](class-psr-http-message-serverrequestinterface-method-withoutattribute-tags.md)

seegetAttributes()

##### Return values

static

#### withoutHeader()  [header link](class-psr-http-message-messageinterface.md\#method_withoutHeader)

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

#### withParsedBody()  [header link](class-psr-http-message-serverrequestinterface-method-withparsedbody.md)

Return an instance with the specified body parameters.

`
    public
                    withParsedBody(null|array<string|int, mixed>|object $data) : static`

These MAY be injected during instantiation.

If the request Content-Type is either application/x-www-form-urlencoded
or multipart/form-data, and the request method is POST, use this method
ONLY to inject the contents of $\_POST.

The data IS NOT REQUIRED to come from $\_POST, but MUST be the results of
deserializing the request body content. Deserialization/parsing returns
structured data, and, as such, this method ONLY accepts arrays or objects,
or a null value if nothing was available to parse.

As an example, if content negotiation determines that the request data
is a JSON payload, this method could be used to create a request
instance with the deserialized parameters.

This method MUST be implemented in such a way as to retain the
immutability of the message, and MUST return an instance that has the
updated body parameters.

##### Parameters

$data
: null\|array<string\|int, mixed>\|object

The deserialized body data. This will
typically be in an array or object.

##### Tags  [header link](class-psr-http-message-serverrequestinterface-method-withparsedbody-tags.md)

throwsInvalidArgumentException

if an unsupported argument type is
provided.

##### Return values

static

#### withProtocolVersion()  [header link](class-psr-http-message-messageinterface.md\#method_withProtocolVersion)

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

#### withQueryParams()  [header link](class-psr-http-message-serverrequestinterface-method-withqueryparams.md)

Return an instance with the specified query string arguments.

`
    public
                    withQueryParams(array<string|int, mixed> $query) : static`

These values SHOULD remain immutable over the course of the incoming
request. They MAY be injected during instantiation, such as from PHP's
$\_GET superglobal, or MAY be derived from some other value such as the
URI. In cases where the arguments are parsed from the URI, the data
MUST be compatible with what PHP's parse\_str() would return for
purposes of how duplicate query parameters are handled, and how nested
sets are handled.

Setting query string arguments MUST NOT change the URI stored by the
request, nor the values in the server params.

This method MUST be implemented in such a way as to retain the
immutability of the message, and MUST return an instance that has the
updated query string arguments.

##### Parameters

$query
: array<string\|int, mixed>

Array of query string arguments, typically from
$\_GET.

##### Return values

static

#### withRequestTarget()  [header link](class-psr-http-message-requestinterface.md\#method_withRequestTarget)

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

##### Tags  [header link](class-psr-http-message-requestinterface.md\#method_withRequestTarget\#tags)

link[(for the various\
request-target forms allowed in request messages)](http://tools.ietf.org/html/rfc7230)

##### Return values

static

#### withUploadedFiles()  [header link](class-psr-http-message-serverrequestinterface-method-withuploadedfiles.md)

Create a new instance with the specified uploaded files.

`
    public
                    withUploadedFiles(array<string|int, mixed> $uploadedFiles) : static`

This method MUST be implemented in such a way as to retain the
immutability of the message, and MUST return an instance that has the
updated body parameters.

##### Parameters

$uploadedFiles
: array<string\|int, mixed>

An array tree of UploadedFileInterface instances.

##### Tags  [header link](class-psr-http-message-serverrequestinterface-method-withuploadedfiles-tags.md)

throwsInvalidArgumentException

if an invalid structure is provided.

##### Return values

static

#### withUri()  [header link](class-psr-http-message-requestinterface.md\#method_withUri)

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

##### Tags  [header link](class-psr-http-message-requestinterface.md\#method_withUri\#tags)

link[http://tools.ietf.org/html/rfc3986#section-4.3](http://tools.ietf.org/html/rfc3986)

##### Return values

static
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-psr-http-message-serverrequestinterface-toc-constants.md)
  - [Methods](class-psr-http-message-serverrequestinterface-toc-methods.md)
- Methods
  - [getAttribute()](class-psr-http-message-serverrequestinterface-method-getattribute.md)
  - [getAttributes()](class-psr-http-message-serverrequestinterface-method-getattributes.md)
  - [getBody()](class-psr-http-message-messageinterface.md#method_getBody)
  - [getCookieParams()](class-psr-http-message-serverrequestinterface-method-getcookieparams.md)
  - [getHeader()](class-psr-http-message-messageinterface.md#method_getHeader)
  - [getHeaderLine()](class-psr-http-message-messageinterface.md#method_getHeaderLine)
  - [getHeaders()](class-psr-http-message-messageinterface.md#method_getHeaders)
  - [getMethod()](class-psr-http-message-requestinterface.md#method_getMethod)
  - [getParsedBody()](class-psr-http-message-serverrequestinterface-method-getparsedbody.md)
  - [getProtocolVersion()](class-psr-http-message-messageinterface.md#method_getProtocolVersion)
  - [getQueryParams()](class-psr-http-message-serverrequestinterface-method-getqueryparams.md)
  - [getRequestTarget()](class-psr-http-message-requestinterface.md#method_getRequestTarget)
  - [getServerParams()](class-psr-http-message-serverrequestinterface-method-getserverparams.md)
  - [getUploadedFiles()](class-psr-http-message-serverrequestinterface-method-getuploadedfiles.md)
  - [getUri()](class-psr-http-message-requestinterface.md#method_getUri)
  - [hasHeader()](class-psr-http-message-messageinterface.md#method_hasHeader)
  - [withAddedHeader()](class-psr-http-message-messageinterface.md#method_withAddedHeader)
  - [withAttribute()](class-psr-http-message-serverrequestinterface-method-withattribute.md)
  - [withBody()](class-psr-http-message-messageinterface.md#method_withBody)
  - [withCookieParams()](class-psr-http-message-serverrequestinterface-method-withcookieparams.md)
  - [withHeader()](class-psr-http-message-messageinterface.md#method_withHeader)
  - [withMethod()](class-psr-http-message-requestinterface.md#method_withMethod)
  - [withoutAttribute()](class-psr-http-message-serverrequestinterface-method-withoutattribute.md)
  - [withoutHeader()](class-psr-http-message-messageinterface.md#method_withoutHeader)
  - [withParsedBody()](class-psr-http-message-serverrequestinterface-method-withparsedbody.md)
  - [withProtocolVersion()](class-psr-http-message-messageinterface.md#method_withProtocolVersion)
  - [withQueryParams()](class-psr-http-message-serverrequestinterface-method-withqueryparams.md)
  - [withRequestTarget()](class-psr-http-message-requestinterface.md#method_withRequestTarget)
  - [withUploadedFiles()](class-psr-http-message-serverrequestinterface-method-withuploadedfiles.md)
  - [withUri()](class-psr-http-message-requestinterface.md#method_withUri)

[Back To Top](class-psr-http-message-serverrequestinterface-top.md)
