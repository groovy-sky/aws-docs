Menu

- [GuzzleHttp](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.html)
- [Psr7](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Guzzlehttp.psr7.html)

## Header        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Header.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Header.html\#toc-methods)

[normalize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Header.html#method_normalize)
: array<string\|int, mixed> Converts an array of header values that may contain comma separated
headers into an array of headers with no comma separated values.[parse()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Header.html#method_parse)
: array<string\|int, mixed> Parse an array of header values containing ";" separated data into an
array of associative arrays representing the header key value pair data
of the header. When a parameter does not contain a value, but just
contains a key, this function will inject a key with a '' string value.[splitList()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Header.html#method_splitList)
: array<string\|int, string> Splits a HTTP header defined to contain a comma-separated list into
each individual value. Empty values will be removed.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Header.html\#methods)

#### normalize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Header.html\#method_normalize)

Converts an array of header values that may contain comma separated
headers into an array of headers with no comma separated values.

`
    public
            static        normalize(string|array<string|int, mixed> $header) : array<string|int, mixed>`

##### Parameters

$header
: string\|array<string\|int, mixed>

Header to normalize.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Header.html\#method_normalize\#tags)

deprecated

Use self::splitList() instead.

##### Return values

array<string\|int, mixed>

#### parse()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Header.html\#method_parse)

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

#### splitList()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Header.html\#method_splitList)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Header.html#toc-methods)
- Methods
  - [normalize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Header.html#method_normalize)
  - [parse()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Header.html#method_parse)
  - [splitList()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Header.html#method_splitList)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Psr7.Header.html#top)
