Menu

- [Psr](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Psr.html)
- [Http](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Psr.http.html)
- [Message](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Psr.http.message.html)

## UriInterface     in    - [Aws](package-aws.md)

Value object representing a URI.

This interface is meant to represent URIs according to RFC 3986 and to
provide methods for most common operations. Additional functionality for
working with URIs can be provided on top of the interface or externally.
Its primary use is for HTTP requests, but may also be used in other
contexts.

Instances of this interface are considered immutable; all methods that
might change state MUST be implemented such that they retain the internal
state of the current instance and return an instance that contains the
changed state.

Typically the Host header will be also be present in the request message.
For server-side requests, the scheme will typically be discoverable in the
server parameters.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#tags)

link[(the URI specification)](http://tools.ietf.org/html/rfc3986)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#toc-methods)

[\_\_toString()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method___toString)
: string Return the string representation as a URI reference.[getAuthority()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_getAuthority)
: string Retrieve the authority component of the URI.[getFragment()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_getFragment)
: string Retrieve the fragment component of the URI.[getHost()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_getHost)
: string Retrieve the host component of the URI.[getPath()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_getPath)
: string Retrieve the path component of the URI.[getPort()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_getPort)
: null\|int Retrieve the port component of the URI.[getQuery()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_getQuery)
: string Retrieve the query string of the URI.[getScheme()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_getScheme)
: string Retrieve the scheme component of the URI.[getUserInfo()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_getUserInfo)
: string Retrieve the user information component of the URI.[withFragment()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_withFragment)
: static Return an instance with the specified URI fragment.[withHost()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_withHost)
: static Return an instance with the specified host.[withPath()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_withPath)
: static Return an instance with the specified path.[withPort()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_withPort)
: static Return an instance with the specified port.[withQuery()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_withQuery)
: static Return an instance with the specified query string.[withScheme()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_withScheme)
: static Return an instance with the specified scheme.[withUserInfo()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_withUserInfo)
: static Return an instance with the specified user information.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#methods)

#### \_\_toString()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method___toString)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method___toString\#tags)

see[http://tools.ietf.org/html/rfc3986#section-4.1](http://tools.ietf.org/html/rfc3986)

##### Return values

string

#### getAuthority()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method_getAuthority)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method_getAuthority\#tags)

see[https://tools.ietf.org/html/rfc3986#section-3.2](https://tools.ietf.org/html/rfc3986)

##### Return values

string
—

The URI authority, in "\[user-info@\]host\[:port\]" format.

#### getFragment()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method_getFragment)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method_getFragment\#tags)

see[https://tools.ietf.org/html/rfc3986#section-2](https://tools.ietf.org/html/rfc3986)see[https://tools.ietf.org/html/rfc3986#section-3.5](https://tools.ietf.org/html/rfc3986)

##### Return values

string
—

The URI fragment.

#### getHost()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method_getHost)

Retrieve the host component of the URI.

`
    public
                    getHost() : string`

If no host is present, this method MUST return an empty string.

The value returned MUST be normalized to lowercase, per RFC 3986
Section 3.2.2.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method_getHost\#tags)

see[http://tools.ietf.org/html/rfc3986#section-3.2.2](http://tools.ietf.org/html/rfc3986)

##### Return values

string
—

The URI host.

#### getPath()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method_getPath)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method_getPath\#tags)

see[https://tools.ietf.org/html/rfc3986#section-2](https://tools.ietf.org/html/rfc3986)see[https://tools.ietf.org/html/rfc3986#section-3.3](https://tools.ietf.org/html/rfc3986)

##### Return values

string
—

The URI path.

#### getPort()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method_getPort)

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

#### getQuery()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method_getQuery)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method_getQuery\#tags)

see[https://tools.ietf.org/html/rfc3986#section-2](https://tools.ietf.org/html/rfc3986)see[https://tools.ietf.org/html/rfc3986#section-3.4](https://tools.ietf.org/html/rfc3986)

##### Return values

string
—

The URI query string.

#### getScheme()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method_getScheme)

Retrieve the scheme component of the URI.

`
    public
                    getScheme() : string`

If no scheme is present, this method MUST return an empty string.

The value returned MUST be normalized to lowercase, per RFC 3986
Section 3.1.

The trailing ":" character is not part of the scheme and MUST NOT be
added.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method_getScheme\#tags)

see[https://tools.ietf.org/html/rfc3986#section-3.1](https://tools.ietf.org/html/rfc3986)

##### Return values

string
—

The URI scheme.

#### getUserInfo()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method_getUserInfo)

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

#### withFragment()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method_withFragment)

Return an instance with the specified URI fragment.

`
    public
                    withFragment(string $fragment) : static`

This method MUST retain the state of the current instance, and return
an instance that contains the specified URI fragment.

Users can provide both encoded and decoded fragment characters.
Implementations ensure the correct encoding as outlined in getFragment().

An empty fragment value is equivalent to removing the fragment.

##### Parameters

$fragment
: string

The fragment to use with the new instance.

##### Return values

static
—

A new instance with the specified fragment.

#### withHost()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method_withHost)

Return an instance with the specified host.

`
    public
                    withHost(string $host) : static`

This method MUST retain the state of the current instance, and return
an instance that contains the specified host.

An empty host value is equivalent to removing the host.

##### Parameters

$host
: string

The hostname to use with the new instance.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method_withHost\#tags)

throwsInvalidArgumentException

for invalid hostnames.

##### Return values

static
—

A new instance with the specified host.

#### withPath()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method_withPath)

Return an instance with the specified path.

`
    public
                    withPath(string $path) : static`

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
: string

The path to use with the new instance.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method_withPath\#tags)

throwsInvalidArgumentException

for invalid paths.

##### Return values

static
—

A new instance with the specified path.

#### withPort()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method_withPort)

Return an instance with the specified port.

`
    public
                    withPort(null|int $port) : static`

This method MUST retain the state of the current instance, and return
an instance that contains the specified port.

Implementations MUST raise an exception for ports outside the
established TCP and UDP port ranges.

A null value provided for the port is equivalent to removing the port
information.

##### Parameters

$port
: null\|int

The port to use with the new instance; a null value
removes the port information.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method_withPort\#tags)

throwsInvalidArgumentException

for invalid ports.

##### Return values

static
—

A new instance with the specified port.

#### withQuery()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method_withQuery)

Return an instance with the specified query string.

`
    public
                    withQuery(string $query) : static`

This method MUST retain the state of the current instance, and return
an instance that contains the specified query string.

Users can provide both encoded and decoded query characters.
Implementations ensure the correct encoding as outlined in getQuery().

An empty query string value is equivalent to removing the query string.

##### Parameters

$query
: string

The query string to use with the new instance.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method_withQuery\#tags)

throwsInvalidArgumentException

for invalid query strings.

##### Return values

static
—

A new instance with the specified query string.

#### withScheme()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method_withScheme)

Return an instance with the specified scheme.

`
    public
                    withScheme(string $scheme) : static`

This method MUST retain the state of the current instance, and return
an instance that contains the specified scheme.

Implementations MUST support the schemes "http" and "https" case
insensitively, and MAY accommodate other schemes if required.

An empty scheme is equivalent to removing the scheme.

##### Parameters

$scheme
: string

The scheme to use with the new instance.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method_withScheme\#tags)

throwsInvalidArgumentException

for invalid or unsupported schemes.

##### Return values

static
—

A new instance with the specified scheme.

#### withUserInfo()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html\#method_withUserInfo)

Return an instance with the specified user information.

`
    public
                    withUserInfo(string $user[, null|string $password = null ]) : static`

This method MUST retain the state of the current instance, and return
an instance that contains the specified user information.

Password is optional, but the user information MUST include the
user; an empty string for the user is equivalent to removing user
information.

##### Parameters

$user
: string

The user name to use for authority.

$password
: null\|string
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
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#toc-methods)
- Methods
  - [\_\_toString()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method___toString)
  - [getAuthority()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_getAuthority)
  - [getFragment()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_getFragment)
  - [getHost()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_getHost)
  - [getPath()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_getPath)
  - [getPort()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_getPort)
  - [getQuery()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_getQuery)
  - [getScheme()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_getScheme)
  - [getUserInfo()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_getUserInfo)
  - [withFragment()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_withFragment)
  - [withHost()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_withHost)
  - [withPath()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_withPath)
  - [withPort()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_withPort)
  - [withQuery()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_withQuery)
  - [withScheme()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_withScheme)
  - [withUserInfo()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#method_withUserInfo)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html#top)
