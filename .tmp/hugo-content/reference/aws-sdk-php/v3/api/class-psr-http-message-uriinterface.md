Menu

- [Psr](namespace-psr.md)
- [Http](namespace-psr-http.md)
- [Message](namespace-psr-http-message.md)

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

##### Tags  [header link](class-psr-http-message-uriinterface-tags.md)

link[(the URI specification)](http://tools.ietf.org/html/rfc3986)

### Table of Contents  [header link](class-psr-http-message-uriinterface-toc.md)

#### Methods  [header link](class-psr-http-message-uriinterface-toc-methods.md)

[\_\_toString()](class-psr-http-message-uriinterface-method-tostring.md)
: string Return the string representation as a URI reference.[getAuthority()](class-psr-http-message-uriinterface-method-getauthority.md)
: string Retrieve the authority component of the URI.[getFragment()](class-psr-http-message-uriinterface-method-getfragment.md)
: string Retrieve the fragment component of the URI.[getHost()](class-psr-http-message-uriinterface-method-gethost.md)
: string Retrieve the host component of the URI.[getPath()](class-psr-http-message-uriinterface-method-getpath.md)
: string Retrieve the path component of the URI.[getPort()](class-psr-http-message-uriinterface-method-getport.md)
: null\|int Retrieve the port component of the URI.[getQuery()](class-psr-http-message-uriinterface-method-getquery.md)
: string Retrieve the query string of the URI.[getScheme()](class-psr-http-message-uriinterface-method-getscheme.md)
: string Retrieve the scheme component of the URI.[getUserInfo()](class-psr-http-message-uriinterface-method-getuserinfo.md)
: string Retrieve the user information component of the URI.[withFragment()](class-psr-http-message-uriinterface-method-withfragment.md)
: static Return an instance with the specified URI fragment.[withHost()](class-psr-http-message-uriinterface-method-withhost.md)
: static Return an instance with the specified host.[withPath()](class-psr-http-message-uriinterface-method-withpath.md)
: static Return an instance with the specified path.[withPort()](class-psr-http-message-uriinterface-method-withport.md)
: static Return an instance with the specified port.[withQuery()](class-psr-http-message-uriinterface-method-withquery.md)
: static Return an instance with the specified query string.[withScheme()](class-psr-http-message-uriinterface-method-withscheme.md)
: static Return an instance with the specified scheme.[withUserInfo()](class-psr-http-message-uriinterface-method-withuserinfo.md)
: static Return an instance with the specified user information.

### Methods  [header link](class-psr-http-message-uriinterface-methods.md)

#### \_\_toString()  [header link](class-psr-http-message-uriinterface-method-tostring.md)

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

##### Tags  [header link](class-psr-http-message-uriinterface-method-tostring-tags.md)

see[http://tools.ietf.org/html/rfc3986#section-4.1](http://tools.ietf.org/html/rfc3986)

##### Return values

string

#### getAuthority()  [header link](class-psr-http-message-uriinterface-method-getauthority.md)

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

##### Tags  [header link](class-psr-http-message-uriinterface-method-getauthority-tags.md)

see[https://tools.ietf.org/html/rfc3986#section-3.2](https://tools.ietf.org/html/rfc3986)

##### Return values

string
—

The URI authority, in "\[user-info@\]host\[:port\]" format.

#### getFragment()  [header link](class-psr-http-message-uriinterface-method-getfragment.md)

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

##### Tags  [header link](class-psr-http-message-uriinterface-method-getfragment-tags.md)

see[https://tools.ietf.org/html/rfc3986#section-2](https://tools.ietf.org/html/rfc3986)see[https://tools.ietf.org/html/rfc3986#section-3.5](https://tools.ietf.org/html/rfc3986)

##### Return values

string
—

The URI fragment.

#### getHost()  [header link](class-psr-http-message-uriinterface-method-gethost.md)

Retrieve the host component of the URI.

`
    public
                    getHost() : string`

If no host is present, this method MUST return an empty string.

The value returned MUST be normalized to lowercase, per RFC 3986
Section 3.2.2.

##### Tags  [header link](class-psr-http-message-uriinterface-method-gethost-tags.md)

see[http://tools.ietf.org/html/rfc3986#section-3.2.2](http://tools.ietf.org/html/rfc3986)

##### Return values

string
—

The URI host.

#### getPath()  [header link](class-psr-http-message-uriinterface-method-getpath.md)

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

##### Tags  [header link](class-psr-http-message-uriinterface-method-getpath-tags.md)

see[https://tools.ietf.org/html/rfc3986#section-2](https://tools.ietf.org/html/rfc3986)see[https://tools.ietf.org/html/rfc3986#section-3.3](https://tools.ietf.org/html/rfc3986)

##### Return values

string
—

The URI path.

#### getPort()  [header link](class-psr-http-message-uriinterface-method-getport.md)

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

#### getQuery()  [header link](class-psr-http-message-uriinterface-method-getquery.md)

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

##### Tags  [header link](class-psr-http-message-uriinterface-method-getquery-tags.md)

see[https://tools.ietf.org/html/rfc3986#section-2](https://tools.ietf.org/html/rfc3986)see[https://tools.ietf.org/html/rfc3986#section-3.4](https://tools.ietf.org/html/rfc3986)

##### Return values

string
—

The URI query string.

#### getScheme()  [header link](class-psr-http-message-uriinterface-method-getscheme.md)

Retrieve the scheme component of the URI.

`
    public
                    getScheme() : string`

If no scheme is present, this method MUST return an empty string.

The value returned MUST be normalized to lowercase, per RFC 3986
Section 3.1.

The trailing ":" character is not part of the scheme and MUST NOT be
added.

##### Tags  [header link](class-psr-http-message-uriinterface-method-getscheme-tags.md)

see[https://tools.ietf.org/html/rfc3986#section-3.1](https://tools.ietf.org/html/rfc3986)

##### Return values

string
—

The URI scheme.

#### getUserInfo()  [header link](class-psr-http-message-uriinterface-method-getuserinfo.md)

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

#### withFragment()  [header link](class-psr-http-message-uriinterface-method-withfragment.md)

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

#### withHost()  [header link](class-psr-http-message-uriinterface-method-withhost.md)

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

##### Tags  [header link](class-psr-http-message-uriinterface-method-withhost-tags.md)

throwsInvalidArgumentException

for invalid hostnames.

##### Return values

static
—

A new instance with the specified host.

#### withPath()  [header link](class-psr-http-message-uriinterface-method-withpath.md)

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

##### Tags  [header link](class-psr-http-message-uriinterface-method-withpath-tags.md)

throwsInvalidArgumentException

for invalid paths.

##### Return values

static
—

A new instance with the specified path.

#### withPort()  [header link](class-psr-http-message-uriinterface-method-withport.md)

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

##### Tags  [header link](class-psr-http-message-uriinterface-method-withport-tags.md)

throwsInvalidArgumentException

for invalid ports.

##### Return values

static
—

A new instance with the specified port.

#### withQuery()  [header link](class-psr-http-message-uriinterface-method-withquery.md)

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

##### Tags  [header link](class-psr-http-message-uriinterface-method-withquery-tags.md)

throwsInvalidArgumentException

for invalid query strings.

##### Return values

static
—

A new instance with the specified query string.

#### withScheme()  [header link](class-psr-http-message-uriinterface-method-withscheme.md)

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

##### Tags  [header link](class-psr-http-message-uriinterface-method-withscheme-tags.md)

throwsInvalidArgumentException

for invalid or unsupported schemes.

##### Return values

static
—

A new instance with the specified scheme.

#### withUserInfo()  [header link](class-psr-http-message-uriinterface-method-withuserinfo.md)

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
  - [Constants](class-psr-http-message-uriinterface-toc-constants.md)
  - [Methods](class-psr-http-message-uriinterface-toc-methods.md)
- Methods
  - [\_\_toString()](class-psr-http-message-uriinterface-method-tostring.md)
  - [getAuthority()](class-psr-http-message-uriinterface-method-getauthority.md)
  - [getFragment()](class-psr-http-message-uriinterface-method-getfragment.md)
  - [getHost()](class-psr-http-message-uriinterface-method-gethost.md)
  - [getPath()](class-psr-http-message-uriinterface-method-getpath.md)
  - [getPort()](class-psr-http-message-uriinterface-method-getport.md)
  - [getQuery()](class-psr-http-message-uriinterface-method-getquery.md)
  - [getScheme()](class-psr-http-message-uriinterface-method-getscheme.md)
  - [getUserInfo()](class-psr-http-message-uriinterface-method-getuserinfo.md)
  - [withFragment()](class-psr-http-message-uriinterface-method-withfragment.md)
  - [withHost()](class-psr-http-message-uriinterface-method-withhost.md)
  - [withPath()](class-psr-http-message-uriinterface-method-withpath.md)
  - [withPort()](class-psr-http-message-uriinterface-method-withport.md)
  - [withQuery()](class-psr-http-message-uriinterface-method-withquery.md)
  - [withScheme()](class-psr-http-message-uriinterface-method-withscheme.md)
  - [withUserInfo()](class-psr-http-message-uriinterface-method-withuserinfo.md)

[Back To Top](class-psr-http-message-uriinterface-top.md)
