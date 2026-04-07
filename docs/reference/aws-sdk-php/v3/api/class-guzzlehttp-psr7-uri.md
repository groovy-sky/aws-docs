Menu

- [GuzzleHttp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.html)
- [Psr7](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.psr7.html)

## Uri        in package    - [Aws](package-aws.md)       implements  [UriInterface](class-psr-http-message-uriinterface.md), JsonSerializable

PSR-7 URI implementation.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#tags)

author

Michael Dowling

author

Tobias Schultze

author

Matthew Weier O'Phinney

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#toc-interfaces)

[UriInterface](class-psr-http-message-uriinterface.md)Value object representing a URI.JsonSerializable

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method___construct)
: mixed [\_\_toString()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method___toString)
: string Return the string representation as a URI reference.[composeComponents()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_composeComponents)
: string Composes a URI reference string from its various components.[fromParts()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_fromParts)
: [UriInterface](class-psr-http-message-uriinterface.md)Creates a URI from a hash of \`parse\_url\` components.[getAuthority()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_getAuthority)
: string Retrieve the authority component of the URI.[getFragment()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_getFragment)
: string Retrieve the fragment component of the URI.[getHost()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_getHost)
: string Retrieve the host component of the URI.[getPath()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_getPath)
: string Retrieve the path component of the URI.[getPort()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_getPort)
: null\|int Retrieve the port component of the URI.[getQuery()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_getQuery)
: string Retrieve the query string of the URI.[getScheme()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_getScheme)
: string Retrieve the scheme component of the URI.[getUserInfo()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_getUserInfo)
: string Retrieve the user information component of the URI.[isAbsolute()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_isAbsolute)
: bool Whether the URI is absolute, i.e. it has a scheme.[isAbsolutePathReference()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_isAbsolutePathReference)
: bool Whether the URI is a absolute-path reference.[isDefaultPort()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_isDefaultPort)
: bool Whether the URI has the default port of the current scheme.[isNetworkPathReference()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_isNetworkPathReference)
: bool Whether the URI is a network-path reference.[isRelativePathReference()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_isRelativePathReference)
: bool Whether the URI is a relative-path reference.[isSameDocumentReference()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_isSameDocumentReference)
: bool Whether the URI is a same-document reference.[jsonSerialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_jsonSerialize)
: string [withFragment()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_withFragment)
: static Return an instance with the specified URI fragment.[withHost()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_withHost)
: static Return an instance with the specified host.[withoutQueryValue()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_withoutQueryValue)
: [UriInterface](class-psr-http-message-uriinterface.md)Creates a new URI with a specific query string value removed.[withPath()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_withPath)
: static Return an instance with the specified path.[withPort()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_withPort)
: static Return an instance with the specified port.[withQuery()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_withQuery)
: static Return an instance with the specified query string.[withQueryValue()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_withQueryValue)
: [UriInterface](class-psr-http-message-uriinterface.md)Creates a new URI with a specific query string value.[withQueryValues()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_withQueryValues)
: [UriInterface](class-psr-http-message-uriinterface.md)Creates a new URI with multiple specific query string values.[withScheme()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_withScheme)
: static Return an instance with the specified scheme.[withUserInfo()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_withUserInfo)
: static Return an instance with the specified user information.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method___construct)

`
    public
                    __construct([string $uri = '' ]) : mixed`

##### Parameters

$uri
: string
= ''

#### \_\_toString()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method___toString)

Return the string representation as a URI reference.

`
    public
                    __toString() : string`

Depending on which components of the URI are present, the resulting
string is either a full URI or relative reference according to RFC 3986,
Section 4.1. The method concatenates the various components of the URI,
using the appropriate delimiters:

- If a scheme is present, it MUST be suffixed by ":".
- If an authority is present, it MUST be prefixed by "//".
- The path can be concatenated without delimiters. But there are two
cases where the path has to be adjusted to make the URI reference
valid as PHP does not allow to throw an exception in \_\_toString():
  - If the path is rootless and an authority is present, the path MUST
    be prefixed by "/".
  - If the path is starting with more than one "/" and no authority is
    present, the starting slashes MUST be reduced to one.
- If a query is present, it MUST be prefixed by "?".
- If a fragment is present, it MUST be prefixed by "#".

##### Return values

string

#### composeComponents()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_composeComponents)

Composes a URI reference string from its various components.

`
    public
            static        composeComponents(string|null $scheme, string|null $authority, string $path, string|null $query, string|null $fragment) : string`

Usually this method does not need to be called manually but instead is used indirectly via
`Psr\Http\Message\UriInterface::__toString`.

PSR-7 UriInterface treats an empty component the same as a missing component as
getQuery(), getFragment() etc. always return a string. This explains the slight
difference to RFC 3986 Section 5.3.

Another adjustment is that the authority separator is added even when the authority is missing/empty
for the "file" scheme. This is because PHP stream functions like `file_get_contents` only work with
`file:///myfile` but not with `file:/myfile` although they are equivalent according to RFC 3986. But
`file:///` is the more common syntax for the file scheme anyway (Chrome for example redirects to
that format).

##### Parameters

$scheme
: string\|null$authority
: string\|null$path
: string$query
: string\|null$fragment
: string\|null

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_composeComponents\#tags)

see[https://datatracker.ietf.org/doc/html/rfc3986#section-5.3](https://datatracker.ietf.org/doc/html/rfc3986)

##### Return values

string

#### fromParts()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_fromParts)

Creates a URI from a hash of \`parse\_url\` components.

`
    public
            static        fromParts(array<string|int, mixed> $parts) : UriInterface`

##### Parameters

$parts
: array<string\|int, mixed>

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_fromParts\#tags)

see[https://www.php.net/manual/en/function.parse-url.php](https://www.php.net/manual/en/function.parse-url.php)throws[MalformedUriException](class-guzzlehttp-psr7-exception-malformeduriexception.md)

If the components do not form a valid URI.

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)

#### getAuthority()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_getAuthority)

Retrieve the authority component of the URI.

`
    public
                    getAuthority() : string`

If no authority information is present, this method MUST return an empty
string.

The authority syntax of the URI is:

```
[user-info@]host[:port]

```

If the port component is not set or is the standard port for the current
scheme, it SHOULD NOT be included.

##### Return values

string
—

The URI authority, in "\[user-info@\]host\[:port\]" format.

#### getFragment()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_getFragment)

Retrieve the fragment component of the URI.

`
    public
                    getFragment() : string`

If no fragment is present, this method MUST return an empty string.

The leading "#" character is not part of the fragment and MUST NOT be
added.

The value returned MUST be percent-encoded, but MUST NOT double-encode
any characters. To determine what characters to encode, please refer to
RFC 3986, Sections 2 and 3.5.

##### Return values

string
—

The URI fragment.

#### getHost()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_getHost)

Retrieve the host component of the URI.

`
    public
                    getHost() : string`

If no host is present, this method MUST return an empty string.

The value returned MUST be normalized to lowercase, per RFC 3986
Section 3.2.2.

##### Return values

string
—

The URI host.

#### getPath()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_getPath)

Retrieve the path component of the URI.

`
    public
                    getPath() : string`

The path can either be empty or absolute (starting with a slash) or
rootless (not starting with a slash). Implementations MUST support all
three syntaxes.

Normally, the empty path "" and absolute path "/" are considered equal as
defined in RFC 7230 Section 2.7.3. But this method MUST NOT automatically
do this normalization because in contexts with a trimmed base path, e.g.
the front controller, this difference becomes significant. It's the task
of the user to handle both "" and "/".

The value returned MUST be percent-encoded, but MUST NOT double-encode
any characters. To determine what characters to encode, please refer to
RFC 3986, Sections 2 and 3.3.

As an example, if the value should include a slash ("/") not intended as
delimiter between path segments, that value MUST be passed in encoded
form (e.g., "%2F") to the instance.

##### Return values

string
—

The URI path.

#### getPort()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_getPort)

Retrieve the port component of the URI.

`
    public
                    getPort() : null|int`

If a port is present, and it is non-standard for the current scheme,
this method MUST return it as an integer. If the port is the standard port
used with the current scheme, this method SHOULD return null.

If no port is present, and no scheme is present, this method MUST return
a null value.

If no port is present, but a scheme is present, this method MAY return
the standard port for that scheme, but SHOULD return null.

##### Return values

null\|int
—

The URI port.

#### getQuery()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_getQuery)

Retrieve the query string of the URI.

`
    public
                    getQuery() : string`

If no query string is present, this method MUST return an empty string.

The leading "?" character is not part of the query and MUST NOT be
added.

The value returned MUST be percent-encoded, but MUST NOT double-encode
any characters. To determine what characters to encode, please refer to
RFC 3986, Sections 2 and 3.4.

As an example, if a value in a key/value pair of the query string should
include an ampersand ("&") not intended as a delimiter between values,
that value MUST be passed in encoded form (e.g., "%26") to the instance.

##### Return values

string
—

The URI query string.

#### getScheme()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_getScheme)

Retrieve the scheme component of the URI.

`
    public
                    getScheme() : string`

If no scheme is present, this method MUST return an empty string.

The value returned MUST be normalized to lowercase, per RFC 3986
Section 3.1.

The trailing ":" character is not part of the scheme and MUST NOT be
added.

##### Return values

string
—

The URI scheme.

#### getUserInfo()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_getUserInfo)

Retrieve the user information component of the URI.

`
    public
                    getUserInfo() : string`

If no user information is present, this method MUST return an empty
string.

If a user is present in the URI, this will return that value;
additionally, if the password is also present, it will be appended to the
user value, with a colon (":") separating the values.

The trailing "@" character is not part of the user information and MUST
NOT be added.

##### Return values

string
—

The URI user information, in "username\[:password\]" format.

#### isAbsolute()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_isAbsolute)

Whether the URI is absolute, i.e. it has a scheme.

`
    public
            static        isAbsolute(UriInterface $uri) : bool`

An instance of UriInterface can either be an absolute URI or a relative reference. This method returns true
if it is the former. An absolute URI has a scheme. A relative reference is used to express a URI relative
to another URI, the base URI. Relative references can be divided into several forms:

- network-path references, e.g. '//example.com/path'
- absolute-path references, e.g. '/path'
- relative-path references, e.g. 'subpath'

##### Parameters

$uri
: [UriInterface](class-psr-http-message-uriinterface.md)

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_isAbsolute\#tags)

seeUri::isNetworkPathReferenceseeUri::isAbsolutePathReferenceseeUri::isRelativePathReferencesee[https://datatracker.ietf.org/doc/html/rfc3986#section-4](https://datatracker.ietf.org/doc/html/rfc3986)

##### Return values

bool

#### isAbsolutePathReference()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_isAbsolutePathReference)

Whether the URI is a absolute-path reference.

`
    public
            static        isAbsolutePathReference(UriInterface $uri) : bool`

A relative reference that begins with a single slash character is termed an absolute-path reference.

##### Parameters

$uri
: [UriInterface](class-psr-http-message-uriinterface.md)

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_isAbsolutePathReference\#tags)

see[https://datatracker.ietf.org/doc/html/rfc3986#section-4.2](https://datatracker.ietf.org/doc/html/rfc3986)

##### Return values

bool

#### isDefaultPort()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_isDefaultPort)

Whether the URI has the default port of the current scheme.

`
    public
            static        isDefaultPort(UriInterface $uri) : bool`

`Psr\Http\Message\UriInterface::getPort` may return null or the standard port. This method can be used
independently of the implementation.

##### Parameters

$uri
: [UriInterface](class-psr-http-message-uriinterface.md)

##### Return values

bool

#### isNetworkPathReference()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_isNetworkPathReference)

Whether the URI is a network-path reference.

`
    public
            static        isNetworkPathReference(UriInterface $uri) : bool`

A relative reference that begins with two slash characters is termed an network-path reference.

##### Parameters

$uri
: [UriInterface](class-psr-http-message-uriinterface.md)

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_isNetworkPathReference\#tags)

see[https://datatracker.ietf.org/doc/html/rfc3986#section-4.2](https://datatracker.ietf.org/doc/html/rfc3986)

##### Return values

bool

#### isRelativePathReference()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_isRelativePathReference)

Whether the URI is a relative-path reference.

`
    public
            static        isRelativePathReference(UriInterface $uri) : bool`

A relative reference that does not begin with a slash character is termed a relative-path reference.

##### Parameters

$uri
: [UriInterface](class-psr-http-message-uriinterface.md)

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_isRelativePathReference\#tags)

see[https://datatracker.ietf.org/doc/html/rfc3986#section-4.2](https://datatracker.ietf.org/doc/html/rfc3986)

##### Return values

bool

#### isSameDocumentReference()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_isSameDocumentReference)

Whether the URI is a same-document reference.

`
    public
            static        isSameDocumentReference(UriInterface $uri[, UriInterface|null $base = null ]) : bool`

A same-document reference refers to a URI that is, aside from its fragment
component, identical to the base URI. When no base URI is given, only an empty
URI reference (apart from its fragment) is considered a same-document reference.

##### Parameters

$uri
: [UriInterface](class-psr-http-message-uriinterface.md)

The URI to check

$base
: [UriInterface](class-psr-http-message-uriinterface.md) \|null
= null

An optional base URI to compare against

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_isSameDocumentReference\#tags)

see[https://datatracker.ietf.org/doc/html/rfc3986#section-4.4](https://datatracker.ietf.org/doc/html/rfc3986)

##### Return values

bool

#### jsonSerialize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_jsonSerialize)

`
    public
                    jsonSerialize() : string`

##### Return values

string

#### withFragment()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_withFragment)

Return an instance with the specified URI fragment.

`
    public
                    withFragment(mixed $fragment) : static`

This method MUST retain the state of the current instance, and return
an instance that contains the specified URI fragment.

Users can provide both encoded and decoded fragment characters.
Implementations ensure the correct encoding as outlined in getFragment().

An empty fragment value is equivalent to removing the fragment.

##### Parameters

$fragment
: mixed

The fragment to use with the new instance.

##### Return values

static
—

A new instance with the specified fragment.

#### withHost()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_withHost)

Return an instance with the specified host.

`
    public
                    withHost(mixed $host) : static`

This method MUST retain the state of the current instance, and return
an instance that contains the specified host.

An empty host value is equivalent to removing the host.

##### Parameters

$host
: mixed

The hostname to use with the new instance.

##### Return values

static
—

A new instance with the specified host.

#### withoutQueryValue()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_withoutQueryValue)

Creates a new URI with a specific query string value removed.

`
    public
            static        withoutQueryValue(UriInterface $uri, string $key) : UriInterface`

Any existing query string values that exactly match the provided key are
removed.

##### Parameters

$uri
: [UriInterface](class-psr-http-message-uriinterface.md)

URI to use as a base.

$key
: string

Query string key to remove.

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)

#### withPath()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_withPath)

Return an instance with the specified path.

`
    public
                    withPath(mixed $path) : static`

This method MUST retain the state of the current instance, and return
an instance that contains the specified path.

The path can either be empty or absolute (starting with a slash) or
rootless (not starting with a slash). Implementations MUST support all
three syntaxes.

If the path is intended to be domain-relative rather than path relative then
it must begin with a slash ("/"). Paths not starting with a slash ("/")
are assumed to be relative to some base path known to the application or
consumer.

Users can provide both encoded and decoded path characters.
Implementations ensure the correct encoding as outlined in getPath().

##### Parameters

$path
: mixed

The path to use with the new instance.

##### Return values

static
—

A new instance with the specified path.

#### withPort()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_withPort)

Return an instance with the specified port.

`
    public
                    withPort(mixed $port) : static`

This method MUST retain the state of the current instance, and return
an instance that contains the specified port.

Implementations MUST raise an exception for ports outside the
established TCP and UDP port ranges.

A null value provided for the port is equivalent to removing the port
information.

##### Parameters

$port
: mixed

The port to use with the new instance; a null value
removes the port information.

##### Return values

static
—

A new instance with the specified port.

#### withQuery()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_withQuery)

Return an instance with the specified query string.

`
    public
                    withQuery(mixed $query) : static`

This method MUST retain the state of the current instance, and return
an instance that contains the specified query string.

Users can provide both encoded and decoded query characters.
Implementations ensure the correct encoding as outlined in getQuery().

An empty query string value is equivalent to removing the query string.

##### Parameters

$query
: mixed

The query string to use with the new instance.

##### Return values

static
—

A new instance with the specified query string.

#### withQueryValue()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_withQueryValue)

Creates a new URI with a specific query string value.

`
    public
            static        withQueryValue(UriInterface $uri, string $key, string|null $value) : UriInterface`

Any existing query string values that exactly match the provided key are
removed and replaced with the given key value pair.

A value of null will set the query string key without a value, e.g. "key"
instead of "key=value".

##### Parameters

$uri
: [UriInterface](class-psr-http-message-uriinterface.md)

URI to use as a base.

$key
: string

Key to set.

$value
: string\|null

Value to set

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)

#### withQueryValues()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_withQueryValues)

Creates a new URI with multiple specific query string values.

`
    public
            static        withQueryValues(UriInterface $uri, array<string|int, string|null> $keyValueArray) : UriInterface`

It has the same behavior as withQueryValue() but for an associative array of key => value.

##### Parameters

$uri
: [UriInterface](class-psr-http-message-uriinterface.md)

URI to use as a base.

$keyValueArray
: array<string\|int, string\|null>

Associative array of key and values

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)

#### withScheme()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_withScheme)

Return an instance with the specified scheme.

`
    public
                    withScheme(mixed $scheme) : static`

This method MUST retain the state of the current instance, and return
an instance that contains the specified scheme.

Implementations MUST support the schemes "http" and "https" case
insensitively, and MAY accommodate other schemes if required.

An empty scheme is equivalent to removing the scheme.

##### Parameters

$scheme
: mixed

The scheme to use with the new instance.

##### Return values

static
—

A new instance with the specified scheme.

#### withUserInfo()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html\#method_withUserInfo)

Return an instance with the specified user information.

`
    public
                    withUserInfo(mixed $user[, mixed $password = null ]) : static`

This method MUST retain the state of the current instance, and return
an instance that contains the specified user information.

Password is optional, but the user information MUST include the
user; an empty string for the user is equivalent to removing user
information.

##### Parameters

$user
: mixed

The user name to use for authority.

$password
: mixed
= null

The password associated with $user.

##### Return values

static
—

A new instance with the specified user information.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method___construct)
  - [\_\_toString()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method___toString)
  - [composeComponents()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_composeComponents)
  - [fromParts()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_fromParts)
  - [getAuthority()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_getAuthority)
  - [getFragment()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_getFragment)
  - [getHost()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_getHost)
  - [getPath()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_getPath)
  - [getPort()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_getPort)
  - [getQuery()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_getQuery)
  - [getScheme()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_getScheme)
  - [getUserInfo()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_getUserInfo)
  - [isAbsolute()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_isAbsolute)
  - [isAbsolutePathReference()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_isAbsolutePathReference)
  - [isDefaultPort()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_isDefaultPort)
  - [isNetworkPathReference()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_isNetworkPathReference)
  - [isRelativePathReference()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_isRelativePathReference)
  - [isSameDocumentReference()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_isSameDocumentReference)
  - [jsonSerialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_jsonSerialize)
  - [withFragment()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_withFragment)
  - [withHost()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_withHost)
  - [withoutQueryValue()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_withoutQueryValue)
  - [withPath()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_withPath)
  - [withPort()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_withPort)
  - [withQuery()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_withQuery)
  - [withQueryValue()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_withQueryValue)
  - [withQueryValues()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_withQueryValues)
  - [withScheme()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_withScheme)
  - [withUserInfo()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#method_withUserInfo)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Uri.html#top)
