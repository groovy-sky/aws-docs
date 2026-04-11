Menu

- [GuzzleHttp](namespace-guzzlehttp.md)
- [Psr7](namespace-guzzlehttp-psr7.md)

## Header        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](class-guzzlehttp-psr7-header-toc.md)

#### Methods  [header link](class-guzzlehttp-psr7-header-toc-methods.md)

[normalize()](class-guzzlehttp-psr7-header-method-normalize.md)
: array<string\|int, mixed> Converts an array of header values that may contain comma separated
headers into an array of headers with no comma separated values.[parse()](class-guzzlehttp-psr7-header-method-parse.md)
: array<string\|int, mixed> Parse an array of header values containing ";" separated data into an
array of associative arrays representing the header key value pair data
of the header. When a parameter does not contain a value, but just
contains a key, this function will inject a key with a '' string value.[splitList()](class-guzzlehttp-psr7-header-method-splitlist.md)
: array<string\|int, string> Splits a HTTP header defined to contain a comma-separated list into
each individual value. Empty values will be removed.

### Methods  [header link](class-guzzlehttp-psr7-header-methods.md)

#### normalize()  [header link](class-guzzlehttp-psr7-header-method-normalize.md)

Converts an array of header values that may contain comma separated
headers into an array of headers with no comma separated values.

`
    public
            static        normalize(string|array<string|int, mixed> $header) : array<string|int, mixed>`

##### Parameters

$header
: string\|array<string\|int, mixed>

Header to normalize.

##### Tags  [header link](class-guzzlehttp-psr7-header-method-normalize-tags.md)

deprecated

Use self::splitList() instead.

##### Return values

array<string\|int, mixed>

#### parse()  [header link](class-guzzlehttp-psr7-header-method-parse.md)

Parse an array of header values containing ";" separated data into an
array of associative arrays representing the header key value pair data
of the header. When a parameter does not contain a value, but just
contains a key, this function will inject a key with a '' string value.

`
    public
            static        parse(string|array<string|int, mixed> $header) : array<string|int, mixed>`

##### Parameters

$header
: string\|array<string\|int, mixed>

Header to parse into components.

##### Return values

array<string\|int, mixed>

#### splitList()  [header link](class-guzzlehttp-psr7-header-method-splitlist.md)

Splits a HTTP header defined to contain a comma-separated list into
each individual value. Empty values will be removed.

`
    public
            static        splitList(string|array<string|int, string> $values) : array<string|int, string>`

Example headers include 'accept', 'cache-control' and 'if-none-match'.

This method must not be used to parse headers that are not defined as
a list, such as 'user-agent' or 'set-cookie'.

##### Parameters

$values
: string\|array<string\|int, string>

Header value as returned by MessageInterface::getHeader()

##### Return values

array<string\|int, string>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-guzzlehttp-psr7-header-toc-methods.md)
- Methods
  - [normalize()](class-guzzlehttp-psr7-header-method-normalize.md)
  - [parse()](class-guzzlehttp-psr7-header-method-parse.md)
  - [splitList()](class-guzzlehttp-psr7-header-method-splitlist.md)

[Back To Top](class-guzzlehttp-psr7-header-top.md)
