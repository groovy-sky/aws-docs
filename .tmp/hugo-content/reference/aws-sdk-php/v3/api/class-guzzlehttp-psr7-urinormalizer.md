Menu

- [GuzzleHttp](namespace-guzzlehttp.md)
- [Psr7](namespace-guzzlehttp-psr7.md)

## UriNormalizer        in package    - [Aws](package-aws.md)

FinalYes

Provides methods to normalize and compare URIs.

##### Tags  [header link](class-guzzlehttp-psr7-urinormalizer-tags.md)

author

Tobias Schultze

see[https://datatracker.ietf.org/doc/html/rfc3986#section-6](https://datatracker.ietf.org/doc/html/rfc3986)

### Table of Contents  [header link](class-guzzlehttp-psr7-urinormalizer-toc.md)

#### Constants  [header link](class-guzzlehttp-psr7-urinormalizer-toc-constants.md)

[CAPITALIZE\_PERCENT\_ENCODING](class-guzzlehttp-psr7-urinormalizer-constant-capitalize-percent-encoding.md)
= 1 All letters within a percent-encoding triplet (e.g., "%3A") are case-insensitive, and should be capitalized.[CONVERT\_EMPTY\_PATH](class-guzzlehttp-psr7-urinormalizer-constant-convert-empty-path.md)
= 4 Converts the empty path to "/" for http and https URIs.[DECODE\_UNRESERVED\_CHARACTERS](class-guzzlehttp-psr7-urinormalizer-constant-decode-unreserved-characters.md)
= 2 Decodes percent-encoded octets of unreserved characters.[PRESERVING\_NORMALIZATIONS](class-guzzlehttp-psr7-urinormalizer-constant-preserving-normalizations.md)
= self::CAPITALIZE\_PERCENT\_ENCODING \| self::DECODE\_UNRESERVED\_CHARACTERS \| self::CONVERT\_EMPTY\_PATH \| self::REMOVE\_DEFAULT\_HOST \| self::REMOVE\_DEFAULT\_PORT \| self::REMOVE\_DOT\_SEGMENTS Default normalizations which only include the ones that preserve semantics.[REMOVE\_DEFAULT\_HOST](class-guzzlehttp-psr7-urinormalizer-constant-remove-default-host.md)
= 8 Removes the default host of the given URI scheme from the URI.[REMOVE\_DEFAULT\_PORT](class-guzzlehttp-psr7-urinormalizer-constant-remove-default-port.md)
= 16 Removes the default port of the given URI scheme from the URI.[REMOVE\_DOT\_SEGMENTS](class-guzzlehttp-psr7-urinormalizer-constant-remove-dot-segments.md)
= 32 Removes unnecessary dot-segments.[REMOVE\_DUPLICATE\_SLASHES](class-guzzlehttp-psr7-urinormalizer-constant-remove-duplicate-slashes.md)
= 64 Paths which include two or more adjacent slashes are converted to one.[SORT\_QUERY\_PARAMETERS](class-guzzlehttp-psr7-urinormalizer-constant-sort-query-parameters.md)
= 128 Sort query parameters with their values in alphabetical order.

#### Methods  [header link](class-guzzlehttp-psr7-urinormalizer-toc-methods.md)

[isEquivalent()](class-guzzlehttp-psr7-urinormalizer-method-isequivalent.md)
: bool Whether two URIs can be considered equivalent.[normalize()](class-guzzlehttp-psr7-urinormalizer-method-normalize.md)
: [UriInterface](class-psr-http-message-uriinterface.md)Returns a normalized URI.

### Constants  [header link](class-guzzlehttp-psr7-urinormalizer-constants.md)

#### CAPITALIZE\_PERCENT\_ENCODING  [header link](class-guzzlehttp-psr7-urinormalizer-constant-capitalize-percent-encoding.md)

All letters within a percent-encoding triplet (e.g., "%3A") are case-insensitive, and should be capitalized.

`
    public
        mixed
    CAPITALIZE_PERCENT_ENCODING
    = 1
`

Example: http://example.org/a%c2%b1b → http://example.org/a%C2%B1b

#### CONVERT\_EMPTY\_PATH  [header link](class-guzzlehttp-psr7-urinormalizer-constant-convert-empty-path.md)

Converts the empty path to "/" for http and https URIs.

`
    public
        mixed
    CONVERT_EMPTY_PATH
    = 4
`

Example: http://example.org → http://example.org/

#### DECODE\_UNRESERVED\_CHARACTERS  [header link](class-guzzlehttp-psr7-urinormalizer-constant-decode-unreserved-characters.md)

Decodes percent-encoded octets of unreserved characters.

`
    public
        mixed
    DECODE_UNRESERVED_CHARACTERS
    = 2
`

For consistency, percent-encoded octets in the ranges of ALPHA (%41–%5A and %61–%7A), DIGIT (%30–%39),
hyphen (%2D), period (%2E), underscore (%5F), or tilde (%7E) should not be created by URI producers and,
when found in a URI, should be decoded to their corresponding unreserved characters by URI normalizers.

Example: http://example.org/%7Eusern%61me/ → http://example.org/~username/

#### PRESERVING\_NORMALIZATIONS  [header link](class-guzzlehttp-psr7-urinormalizer-constant-preserving-normalizations.md)

Default normalizations which only include the ones that preserve semantics.

`
    public
        mixed
    PRESERVING_NORMALIZATIONS
    = self::CAPITALIZE_PERCENT_ENCODING | self::DECODE_UNRESERVED_CHARACTERS | self::CONVERT_EMPTY_PATH | self::REMOVE_DEFAULT_HOST | self::REMOVE_DEFAULT_PORT | self::REMOVE_DOT_SEGMENTS
`

#### REMOVE\_DEFAULT\_HOST  [header link](class-guzzlehttp-psr7-urinormalizer-constant-remove-default-host.md)

Removes the default host of the given URI scheme from the URI.

`
    public
        mixed
    REMOVE_DEFAULT_HOST
    = 8
`

Only the "file" scheme defines the default host "localhost".
All of `file:/myfile`, `file:///myfile`, and `file://localhost/myfile`
are equivalent according to RFC 3986. The first format is not accepted
by PHPs stream functions and thus already normalized implicitly to the
second format in the Uri class. See `GuzzleHttp\Psr7\Uri::composeComponents`.

Example: file://localhost/myfile → file:///myfile

#### REMOVE\_DEFAULT\_PORT  [header link](class-guzzlehttp-psr7-urinormalizer-constant-remove-default-port.md)

Removes the default port of the given URI scheme from the URI.

`
    public
        mixed
    REMOVE_DEFAULT_PORT
    = 16
`

Example: http://example.org:80/ → http://example.org/

#### REMOVE\_DOT\_SEGMENTS  [header link](class-guzzlehttp-psr7-urinormalizer-constant-remove-dot-segments.md)

Removes unnecessary dot-segments.

`
    public
        mixed
    REMOVE_DOT_SEGMENTS
    = 32
`

Dot-segments in relative-path references are not removed as it would
change the semantics of the URI reference.

Example: http://example.org/../a/b/../c/./d.html → http://example.org/a/c/d.html

#### REMOVE\_DUPLICATE\_SLASHES  [header link](class-guzzlehttp-psr7-urinormalizer-constant-remove-duplicate-slashes.md)

Paths which include two or more adjacent slashes are converted to one.

`
    public
        mixed
    REMOVE_DUPLICATE_SLASHES
    = 64
`

Webservers usually ignore duplicate slashes and treat those URIs equivalent.
But in theory those URIs do not need to be equivalent. So this normalization
may change the semantics. Encoded slashes (%2F) are not removed.

Example: http://example.org//foo///bar.html → http://example.org/foo/bar.html

#### SORT\_QUERY\_PARAMETERS  [header link](class-guzzlehttp-psr7-urinormalizer-constant-sort-query-parameters.md)

Sort query parameters with their values in alphabetical order.

`
    public
        mixed
    SORT_QUERY_PARAMETERS
    = 128
`

However, the order of parameters in a URI may be significant (this is not defined by the standard).
So this normalization is not safe and may change the semantics of the URI.

Example: ?lang=en&article=fred → ?article=fred&lang=en

Note: The sorting is neither locale nor Unicode aware (the URI query does not get decoded at all) as the
purpose is to be able to compare URIs in a reproducible way, not to have the params sorted perfectly.

### Methods  [header link](class-guzzlehttp-psr7-urinormalizer-methods.md)

#### isEquivalent()  [header link](class-guzzlehttp-psr7-urinormalizer-method-isequivalent.md)

Whether two URIs can be considered equivalent.

`
    public
            static        isEquivalent(UriInterface $uri1, UriInterface $uri2[, int $normalizations = self::PRESERVING_NORMALIZATIONS ]) : bool`

Both URIs are normalized automatically before comparison with the given $normalizations bitmask. The method also
accepts relative URI references and returns true when they are equivalent. This of course assumes they will be
resolved against the same base URI. If this is not the case, determination of equivalence or difference of
relative references does not mean anything.

##### Parameters

$uri1
: [UriInterface](class-psr-http-message-uriinterface.md)

An URI to compare

$uri2
: [UriInterface](class-psr-http-message-uriinterface.md)

An URI to compare

$normalizations
: int
= self::PRESERVING\_NORMALIZATIONS

A bitmask of normalizations to apply, see constants

##### Tags  [header link](class-guzzlehttp-psr7-urinormalizer-method-isequivalent-tags.md)

see[https://datatracker.ietf.org/doc/html/rfc3986#section-6.1](https://datatracker.ietf.org/doc/html/rfc3986)

##### Return values

bool

#### normalize()  [header link](class-guzzlehttp-psr7-urinormalizer-method-normalize.md)

Returns a normalized URI.

`
    public
            static        normalize(UriInterface $uri[, int $flags = self::PRESERVING_NORMALIZATIONS ]) : UriInterface`

The scheme and host component are already normalized to lowercase per PSR-7 UriInterface.
This methods adds additional normalizations that can be configured with the $flags parameter.

PSR-7 UriInterface cannot distinguish between an empty component and a missing component as
getQuery(), getFragment() etc. always return a string. This means the URIs "/?#" and "/" are
treated equivalent which is not necessarily true according to RFC 3986. But that difference
is highly uncommon in reality. So this potential normalization is implied in PSR-7 as well.

##### Parameters

$uri
: [UriInterface](class-psr-http-message-uriinterface.md)

The URI to normalize

$flags
: int
= self::PRESERVING\_NORMALIZATIONS

A bitmask of normalizations to apply, see constants

##### Tags  [header link](class-guzzlehttp-psr7-urinormalizer-method-normalize-tags.md)

see[https://datatracker.ietf.org/doc/html/rfc3986#section-6.2](https://datatracker.ietf.org/doc/html/rfc3986)

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-guzzlehttp-psr7-urinormalizer-toc-constants.md)
  - [Methods](class-guzzlehttp-psr7-urinormalizer-toc-methods.md)
- Constants
  - [CAPITALIZE\_PERCENT\_ENCODING](class-guzzlehttp-psr7-urinormalizer-constant-capitalize-percent-encoding.md)
  - [CONVERT\_EMPTY\_PATH](class-guzzlehttp-psr7-urinormalizer-constant-convert-empty-path.md)
  - [DECODE\_UNRESERVED\_CHARACTERS](class-guzzlehttp-psr7-urinormalizer-constant-decode-unreserved-characters.md)
  - [PRESERVING\_NORMALIZATIONS](class-guzzlehttp-psr7-urinormalizer-constant-preserving-normalizations.md)
  - [REMOVE\_DEFAULT\_HOST](class-guzzlehttp-psr7-urinormalizer-constant-remove-default-host.md)
  - [REMOVE\_DEFAULT\_PORT](class-guzzlehttp-psr7-urinormalizer-constant-remove-default-port.md)
  - [REMOVE\_DOT\_SEGMENTS](class-guzzlehttp-psr7-urinormalizer-constant-remove-dot-segments.md)
  - [REMOVE\_DUPLICATE\_SLASHES](class-guzzlehttp-psr7-urinormalizer-constant-remove-duplicate-slashes.md)
  - [SORT\_QUERY\_PARAMETERS](class-guzzlehttp-psr7-urinormalizer-constant-sort-query-parameters.md)
- Methods
  - [isEquivalent()](class-guzzlehttp-psr7-urinormalizer-method-isequivalent.md)
  - [normalize()](class-guzzlehttp-psr7-urinormalizer-method-normalize.md)

[Back To Top](class-guzzlehttp-psr7-urinormalizer-top.md)
