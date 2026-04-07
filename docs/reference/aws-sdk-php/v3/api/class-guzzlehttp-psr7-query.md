Menu

- [GuzzleHttp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.html)
- [Psr7](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.psr7.html)

## Query        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Query.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Query.html\#toc-methods)

[build()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Query.html#method_build)
: string Build a query string from an array of key value pairs.[parse()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Query.html#method_parse)
: array<string\|int, mixed> Parse a query string into an associative array.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Query.html\#methods)

#### build()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Query.html\#method_build)

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

#### parse()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Query.html\#method_parse)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Query.html#toc-methods)
- Methods
  - [build()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Query.html#method_build)
  - [parse()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Query.html#method_parse)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Query.html#top)
