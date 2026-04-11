Menu

- [GuzzleHttp](namespace-guzzlehttp.md)
- [Psr7](namespace-guzzlehttp-psr7.md)

## UriResolver        in package    - [Aws](package-aws.md)

FinalYes

Resolves a URI reference in the context of a base URI and the opposite way.

##### Tags  [header link](class-guzzlehttp-psr7-uriresolver-tags.md)

author

Tobias Schultze

see[https://datatracker.ietf.org/doc/html/rfc3986#section-5](https://datatracker.ietf.org/doc/html/rfc3986)

### Table of Contents  [header link](class-guzzlehttp-psr7-uriresolver-toc.md)

#### Methods  [header link](class-guzzlehttp-psr7-uriresolver-toc-methods.md)

[relativize()](class-guzzlehttp-psr7-uriresolver-method-relativize.md)
: [UriInterface](class-psr-http-message-uriinterface.md)Returns the target URI as a relative reference from the base URI.[removeDotSegments()](class-guzzlehttp-psr7-uriresolver-method-removedotsegments.md)
: string Removes dot segments from a path and returns the new path.[resolve()](class-guzzlehttp-psr7-uriresolver-method-resolve.md)
: [UriInterface](class-psr-http-message-uriinterface.md)Converts the relative URI into a new URI that is resolved against the base URI.

### Methods  [header link](class-guzzlehttp-psr7-uriresolver-methods.md)

#### relativize()  [header link](class-guzzlehttp-psr7-uriresolver-method-relativize.md)

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

#### removeDotSegments()  [header link](class-guzzlehttp-psr7-uriresolver-method-removedotsegments.md)

Removes dot segments from a path and returns the new path.

`
    public
            static        removeDotSegments(string $path) : string`

##### Parameters

$path
: string

##### Tags  [header link](class-guzzlehttp-psr7-uriresolver-method-removedotsegments-tags.md)

see[https://datatracker.ietf.org/doc/html/rfc3986#section-5.2.4](https://datatracker.ietf.org/doc/html/rfc3986)

##### Return values

string

#### resolve()  [header link](class-guzzlehttp-psr7-uriresolver-method-resolve.md)

Converts the relative URI into a new URI that is resolved against the base URI.

`
    public
            static        resolve(UriInterface $base, UriInterface $rel) : UriInterface`

##### Parameters

$base
: [UriInterface](class-psr-http-message-uriinterface.md)$rel
: [UriInterface](class-psr-http-message-uriinterface.md)

##### Tags  [header link](class-guzzlehttp-psr7-uriresolver-method-resolve-tags.md)

see[https://datatracker.ietf.org/doc/html/rfc3986#section-5.2](https://datatracker.ietf.org/doc/html/rfc3986)

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-guzzlehttp-psr7-uriresolver-toc-methods.md)
- Methods
  - [relativize()](class-guzzlehttp-psr7-uriresolver-method-relativize.md)
  - [removeDotSegments()](class-guzzlehttp-psr7-uriresolver-method-removedotsegments.md)
  - [resolve()](class-guzzlehttp-psr7-uriresolver-method-resolve.md)

[Back To Top](class-guzzlehttp-psr7-uriresolver-top.md)
