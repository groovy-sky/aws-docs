Menu

- [GuzzleHttp](namespace-guzzlehttp.md)
- [Psr7](namespace-guzzlehttp-psr7.md)

## Query        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](class-guzzlehttp-psr7-query-toc.md)

#### Methods  [header link](class-guzzlehttp-psr7-query-toc-methods.md)

[build()](class-guzzlehttp-psr7-query-method-build.md)
: string Build a query string from an array of key value pairs.[parse()](class-guzzlehttp-psr7-query-method-parse.md)
: array<string\|int, mixed> Parse a query string into an associative array.

### Methods  [header link](class-guzzlehttp-psr7-query-methods.md)

#### build()  [header link](class-guzzlehttp-psr7-query-method-build.md)

Build a query string from an array of key value pairs.

`
    public
            static        build(array<string|int, mixed> $params[, int|false $encoding = PHP_QUERY_RFC3986 ][, bool $treatBoolsAsInts = true ]) : string`

This function can use the return value of `parse()` to build a query
string. This function does not modify the provided keys when an array is
encountered (like `http_build_query()` would).

##### Parameters

$params
: array<string\|int, mixed>

Query string parameters.

$encoding
: int\|false
= PHP\_QUERY\_RFC3986

Set to false to not encode,
PHP\_QUERY\_RFC3986 to encode using
RFC3986, or PHP\_QUERY\_RFC1738 to
encode using RFC1738.

$treatBoolsAsInts
: bool
= true

Set to true to encode as 0/1, and
false as false/true.

##### Return values

string

#### parse()  [header link](class-guzzlehttp-psr7-query-method-parse.md)

Parse a query string into an associative array.

`
    public
            static        parse(string $str[, int|bool $urlEncoding = true ]) : array<string|int, mixed>`

If multiple values are found for the same key, the value of that key
value pair will become an array. This function does not parse nested
PHP style arrays into an associative array (e.g., `foo[a]=1&foo[b]=2`
will be parsed into `['foo[a]' => '1', 'foo[b]' => '2'])`.

##### Parameters

$str
: string

Query string to parse

$urlEncoding
: int\|bool
= true

How the query string is encoded

##### Return values

array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-guzzlehttp-psr7-query-toc-methods.md)
- Methods
  - [build()](class-guzzlehttp-psr7-query-method-build.md)
  - [parse()](class-guzzlehttp-psr7-query-method-parse.md)

[Back To Top](class-guzzlehttp-psr7-query-top.md)
