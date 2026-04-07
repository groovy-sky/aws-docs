Menu

- [GuzzleHttp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.html)
- [Psr7](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.psr7.html)

## UriResolver        in package    - [Aws](package-aws.md)

FinalYes

Resolves a URI reference in the context of a base URI and the opposite way.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.UriResolver.html\#tags)

author

Tobias Schultze

see[https://datatracker.ietf.org/doc/html/rfc3986#section-5](https://datatracker.ietf.org/doc/html/rfc3986)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.UriResolver.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.UriResolver.html\#toc-methods)

[relativize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.UriResolver.html#method_relativize)
: [UriInterface](class-psr-http-message-uriinterface.md)Returns the target URI as a relative reference from the base URI.[removeDotSegments()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.UriResolver.html#method_removeDotSegments)
: string Removes dot segments from a path and returns the new path.[resolve()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.UriResolver.html#method_resolve)
: [UriInterface](class-psr-http-message-uriinterface.md)Converts the relative URI into a new URI that is resolved against the base URI.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.UriResolver.html\#methods)

#### relativize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.UriResolver.html\#method_relativize)

Returns the target URI as a relative reference from the base URI.

`
    public
            static        relativize(UriInterface $base, UriInterface $target) : UriInterface`

This method is the counterpart to resolve():

(string) $target === (string) UriResolver::resolve($base, UriResolver::relativize($base, $target))

One use-case is to use the current request URI as base URI and then generate relative links in your documents
to reduce the document size or offer self-contained downloadable document archives.

$base = new Uri('http://example.com/a/b/');
echo UriResolver::relativize($base, new Uri('http://example.com/a/b/c')); // prints 'c'.
echo UriResolver::relativize($base, new Uri('http://example.com/a/x/y')); // prints '../x/y'.
echo UriResolver::relativize($base, new Uri('http://example.com/a/b/?q')); // prints '?q'.
echo UriResolver::relativize($base, new Uri('http://example.org/a/b/')); // prints '//example.org/a/b/'.

This method also accepts a target that is already relative and will try to relativize it further. Only a
relative-path reference will be returned as-is.

echo UriResolver::relativize($base, new Uri('/a/b/c')); // prints 'c' as well

##### Parameters

$base
: [UriInterface](class-psr-http-message-uriinterface.md)$target
: [UriInterface](class-psr-http-message-uriinterface.md)

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)

#### removeDotSegments()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.UriResolver.html\#method_removeDotSegments)

Removes dot segments from a path and returns the new path.

`
    public
            static        removeDotSegments(string $path) : string`

##### Parameters

$path
: string

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.UriResolver.html\#method_removeDotSegments\#tags)

see[https://datatracker.ietf.org/doc/html/rfc3986#section-5.2.4](https://datatracker.ietf.org/doc/html/rfc3986)

##### Return values

string

#### resolve()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.UriResolver.html\#method_resolve)

Converts the relative URI into a new URI that is resolved against the base URI.

`
    public
            static        resolve(UriInterface $base, UriInterface $rel) : UriInterface`

##### Parameters

$base
: [UriInterface](class-psr-http-message-uriinterface.md)$rel
: [UriInterface](class-psr-http-message-uriinterface.md)

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.UriResolver.html\#method_resolve\#tags)

see[https://datatracker.ietf.org/doc/html/rfc3986#section-5.2](https://datatracker.ietf.org/doc/html/rfc3986)

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.UriResolver.html#toc-methods)
- Methods
  - [relativize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.UriResolver.html#method_relativize)
  - [removeDotSegments()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.UriResolver.html#method_removeDotSegments)
  - [resolve()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.UriResolver.html#method_resolve)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.UriResolver.html#top)
