Menu

- [GuzzleHttp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.html)
- [Psr7](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.psr7.html)

## ServerRequest     extends [Request](class-guzzlehttp-psr7-request.md)   in package    - [Aws](package-aws.md)       implements  [ServerRequestInterface](class-psr-http-message-serverrequestinterface.md)

Server-side HTTP request

Extends the Request definition to add methods for accessing incoming data,
specifically server parameters, cookies, matched path parameters, query
string arguments, body parameters, and upload file information.

"Attributes" are discovered via decomposing the request (and usually
specifically the URI path), and typically will be injected by the application.

Requests are considered immutable; all methods that might change state are
implemented such that they retain the internal state of the current
message and return a new instance that contains the changed state.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html\#toc-interfaces)

[ServerRequestInterface](class-psr-http-message-serverrequestinterface.md)Representation of an incoming, server-side HTTP request.

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method___construct)
: mixed [fromGlobals()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_fromGlobals)
: [ServerRequestInterface](class-psr-http-message-serverrequestinterface.md)Return a ServerRequest populated with superglobals:
$\_GET
$\_POST
$\_COOKIE
$\_FILES
$\_SERVER[getAttribute()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_getAttribute)
: mixed Retrieve a single derived request attribute.[getAttributes()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_getAttributes)
: array<string\|int, mixed> Retrieve attributes derived from the request.[getBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getBody)
: [StreamInterface](class-psr-http-message-streaminterface.md)[getCookieParams()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_getCookieParams)
: array<string\|int, mixed> Retrieve cookies.[getHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getHeader)
: array<string\|int, mixed> [getHeaderLine()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getHeaderLine)
: string [getHeaders()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getHeaders)
: array<string\|int, mixed> [getMethod()](class-guzzlehttp-psr7-request.md#method_getMethod)
: string Retrieves the HTTP method of the request.[getParsedBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_getParsedBody)
: array<string\|int, mixed>\|object\|null Retrieve any parameters provided in the request body.[getProtocolVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getProtocolVersion)
: string [getQueryParams()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_getQueryParams)
: array<string\|int, mixed> Retrieve query string arguments.[getRequestTarget()](class-guzzlehttp-psr7-request.md#method_getRequestTarget)
: string Retrieves the message's request target.[getServerParams()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_getServerParams)
: array<string\|int, mixed> Retrieve server parameters.[getUploadedFiles()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_getUploadedFiles)
: array<string\|int, mixed> Retrieve normalized file upload data.[getUri()](class-guzzlehttp-psr7-request.md#method_getUri)
: [UriInterface](class-psr-http-message-uriinterface.md)Retrieves the URI instance.[getUriFromGlobals()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_getUriFromGlobals)
: [UriInterface](class-psr-http-message-uriinterface.md)Get a Uri populated with values from $\_SERVER.[hasHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_hasHeader)
: bool [normalizeFiles()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_normalizeFiles)
: array<string\|int, mixed> Return an UploadedFile instance array.[withAddedHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withAddedHeader)
: static [withAttribute()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_withAttribute)
: static Return an instance with the specified derived request attribute.[withBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withBody)
: static [withCookieParams()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_withCookieParams)
: static Return an instance with the specified cookies.[withHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withHeader)
: static [withMethod()](class-guzzlehttp-psr7-request.md#method_withMethod)
: static Return an instance with the provided HTTP method.[withoutAttribute()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_withoutAttribute)
: static Return an instance that removes the specified derived request attribute.[withoutHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withoutHeader)
: static [withParsedBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_withParsedBody)
: static Return an instance with the specified body parameters.[withProtocolVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withProtocolVersion)
: static [withQueryParams()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_withQueryParams)
: static Return an instance with the specified query string arguments.[withRequestTarget()](class-guzzlehttp-psr7-request.md#method_withRequestTarget)
: static Return an instance with the specific request-target.[withUploadedFiles()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_withUploadedFiles)
: static Create a new instance with the specified uploaded files.[withUri()](class-guzzlehttp-psr7-request.md#method_withUri)
: static Returns an instance with the provided URI.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html\#method___construct)

`
    public
                    __construct(string $method, string|UriInterface $uri[, array<string|int, string|array<string|int, string>> $headers = [] ][, string|resource|StreamInterface|null $body = null ][, string $version = '1.1' ][, array<string|int, mixed> $serverParams = [] ]) : mixed`

##### Parameters

$method
: string

HTTP method

$uri
: string\| [UriInterface](class-psr-http-message-uriinterface.md)

URI

$headers
: array<string\|int, string\|array<string\|int, string>>
= \[\]

Request headers

$body
: string\|resource\| [StreamInterface](class-psr-http-message-streaminterface.md) \|null
= null

Request body

$version
: string
= '1.1'

Protocol version

$serverParams
: array<string\|int, mixed>
= \[\]

Typically the $\_SERVER superglobal

#### fromGlobals()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html\#method_fromGlobals)

Return a ServerRequest populated with superglobals:
$\_GET
$\_POST
$\_COOKIE
$\_FILES
$\_SERVER

`
    public
            static        fromGlobals() : ServerRequestInterface`

##### Return values

[ServerRequestInterface](class-psr-http-message-serverrequestinterface.md)

#### getAttribute()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html\#method_getAttribute)

Retrieve a single derived request attribute.

`
    public
                    getAttribute(mixed $attribute[, mixed $default = null ]) : mixed`

##### Parameters

$attribute
: mixed$default
: mixed
= null

Default value to return if the attribute does not exist.

#### getAttributes()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html\#method_getAttributes)

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

#### getBody()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html\#method_getBody)

`
    public
                    getBody() : StreamInterface`

##### Return values

[StreamInterface](class-psr-http-message-streaminterface.md)

#### getCookieParams()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html\#method_getCookieParams)

Retrieve cookies.

`
    public
                    getCookieParams() : array<string|int, mixed>`

Retrieves cookies sent by the client to the server.

The data MUST be compatible with the structure of the $\_COOKIE
superglobal.

##### Return values

array<string\|int, mixed>

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

#### getMethod()  [header link](class-guzzlehttp-psr7-request.md\#method_getMethod)

Retrieves the HTTP method of the request.

`
    public
                    getMethod() : string`

##### Return values

string
—

Returns the request method.

#### getParsedBody()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html\#method_getParsedBody)

Retrieve any parameters provided in the request body.

`
    public
                    getParsedBody() : array<string|int, mixed>|object|null`

##### Return values

array<string\|int, mixed>\|object\|null

#### getProtocolVersion()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html\#method_getProtocolVersion)

`
    public
                    getProtocolVersion() : string`

##### Return values

string

#### getQueryParams()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html\#method_getQueryParams)

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

#### getRequestTarget()  [header link](class-guzzlehttp-psr7-request.md\#method_getRequestTarget)

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

#### getServerParams()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html\#method_getServerParams)

Retrieve server parameters.

`
    public
                    getServerParams() : array<string|int, mixed>`

Retrieves data related to the incoming request environment,
typically derived from PHP's $\_SERVER superglobal. The data IS NOT
REQUIRED to originate from $\_SERVER.

##### Return values

array<string\|int, mixed>

#### getUploadedFiles()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html\#method_getUploadedFiles)

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

#### getUri()  [header link](class-guzzlehttp-psr7-request.md\#method_getUri)

Retrieves the URI instance.

`
    public
                    getUri() : UriInterface`

This method MUST return a UriInterface instance.

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)
—

Returns a UriInterface instance
representing the URI of the request.

#### getUriFromGlobals()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html\#method_getUriFromGlobals)

Get a Uri populated with values from $\_SERVER.

`
    public
            static        getUriFromGlobals() : UriInterface`

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)

#### hasHeader()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html\#method_hasHeader)

`
    public
                    hasHeader(mixed $header) : bool`

##### Parameters

$header
: mixed

##### Return values

bool

#### normalizeFiles()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html\#method_normalizeFiles)

Return an UploadedFile instance array.

`
    public
            static        normalizeFiles(array<string|int, mixed> $files) : array<string|int, mixed>`

##### Parameters

$files
: array<string\|int, mixed>

An array which respect $\_FILES structure

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html\#method_normalizeFiles\#tags)

throwsInvalidArgumentException

for unrecognized values

##### Return values

array<string\|int, mixed>

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

#### withAttribute()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html\#method_withAttribute)

Return an instance with the specified derived request attribute.

`
    public
                    withAttribute(mixed $attribute, mixed $value) : static`

This method allows setting a single derived request attribute as
described in getAttributes().

This method MUST be implemented in such a way as to retain the
immutability of the message, and MUST return an instance that has the
updated attribute.

##### Parameters

$attribute
: mixed$value
: mixed

The value of the attribute.

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

#### withCookieParams()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html\#method_withCookieParams)

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

#### withMethod()  [header link](class-guzzlehttp-psr7-request.md\#method_withMethod)

Return an instance with the provided HTTP method.

`
    public
                    withMethod(mixed $method) : static`

While HTTP method names are typically all uppercase characters, HTTP
method names are case-sensitive and thus implementations SHOULD NOT
modify the given string.

This method MUST be implemented in such a way as to retain the
immutability of the message, and MUST return an instance that has the
changed request method.

##### Parameters

$method
: mixed

Case-sensitive method.

##### Return values

static

#### withoutAttribute()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html\#method_withoutAttribute)

Return an instance that removes the specified derived request attribute.

`
    public
                    withoutAttribute(mixed $attribute) : static`

This method allows removing a single derived request attribute as
described in getAttributes().

This method MUST be implemented in such a way as to retain the
immutability of the message, and MUST return an instance that removes
the attribute.

##### Parameters

$attribute
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

#### withParsedBody()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html\#method_withParsedBody)

Return an instance with the specified body parameters.

`
    public
                    withParsedBody(mixed $data) : static`

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
: mixed

The deserialized body data. This will
typically be in an array or object.

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

#### withQueryParams()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html\#method_withQueryParams)

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

#### withRequestTarget()  [header link](class-guzzlehttp-psr7-request.md\#method_withRequestTarget)

Return an instance with the specific request-target.

`
    public
                    withRequestTarget(mixed $requestTarget) : static`

If the request needs a non-origin-form request-target — e.g., for
specifying an absolute-form, authority-form, or asterisk-form —
this method may be used to create an instance with the specified
request-target, verbatim.

This method MUST be implemented in such a way as to retain the
immutability of the message, and MUST return an instance that has the
changed request target.

##### Parameters

$requestTarget
: mixed

##### Return values

static

#### withUploadedFiles()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html\#method_withUploadedFiles)

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

##### Return values

static

#### withUri()  [header link](class-guzzlehttp-psr7-request.md\#method_withUri)

Returns an instance with the provided URI.

`
    public
                    withUri(UriInterface $uri[, mixed $preserveHost = false ]) : static`

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
: mixed
= false

Preserve the original state of the Host header.

##### Return values

static
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method___construct)
  - [fromGlobals()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_fromGlobals)
  - [getAttribute()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_getAttribute)
  - [getAttributes()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_getAttributes)
  - [getBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getBody)
  - [getCookieParams()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_getCookieParams)
  - [getHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getHeader)
  - [getHeaderLine()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getHeaderLine)
  - [getHeaders()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getHeaders)
  - [getMethod()](class-guzzlehttp-psr7-request.md#method_getMethod)
  - [getParsedBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_getParsedBody)
  - [getProtocolVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_getProtocolVersion)
  - [getQueryParams()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_getQueryParams)
  - [getRequestTarget()](class-guzzlehttp-psr7-request.md#method_getRequestTarget)
  - [getServerParams()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_getServerParams)
  - [getUploadedFiles()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_getUploadedFiles)
  - [getUri()](class-guzzlehttp-psr7-request.md#method_getUri)
  - [getUriFromGlobals()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_getUriFromGlobals)
  - [hasHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_hasHeader)
  - [normalizeFiles()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_normalizeFiles)
  - [withAddedHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withAddedHeader)
  - [withAttribute()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_withAttribute)
  - [withBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withBody)
  - [withCookieParams()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_withCookieParams)
  - [withHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withHeader)
  - [withMethod()](class-guzzlehttp-psr7-request.md#method_withMethod)
  - [withoutAttribute()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_withoutAttribute)
  - [withoutHeader()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withoutHeader)
  - [withParsedBody()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_withParsedBody)
  - [withProtocolVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.MessageTrait.html#method_withProtocolVersion)
  - [withQueryParams()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_withQueryParams)
  - [withRequestTarget()](class-guzzlehttp-psr7-request.md#method_withRequestTarget)
  - [withUploadedFiles()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#method_withUploadedFiles)
  - [withUri()](class-guzzlehttp-psr7-request.md#method_withUri)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.ServerRequest.html#top)
