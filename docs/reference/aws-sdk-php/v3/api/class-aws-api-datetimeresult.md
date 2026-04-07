Menu

- [Aws](namespace-aws.md)
- [Api](namespace-aws-api.md)

## DateTimeResult     extends DateTime   in package    - [Aws](package-aws.md)       implements  JsonSerializable

DateTime overrides that make DateTime work more seamlessly as a string,
with JSON documents, and with JMESPath.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html\#toc-interfaces)

JsonSerializable

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html\#toc-methods)

[\_\_toString()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html#method___toString)
: string Serialize the DateTimeResult as an ISO 8601 date string.[fromEpoch()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html#method_fromEpoch)
: [DateTimeResult](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html)Create a new DateTimeResult from a unix timestamp.[fromISO8601()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html#method_fromISO8601)
: [DateTimeResult](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html)[fromTimestamp()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html#method_fromTimestamp)
: [DateTimeResult](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html)Create a new DateTimeResult from an unknown timestamp.[jsonSerialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html#method_jsonSerialize)
: string Serialize the date as an ISO 8601 date when serializing as JSON.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html\#methods)

#### \_\_toString()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html\#method___toString)

Serialize the DateTimeResult as an ISO 8601 date string.

`
    public
                    __toString() : string`

##### Return values

string

#### fromEpoch()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html\#method_fromEpoch)

Create a new DateTimeResult from a unix timestamp.

`
    public
            static        fromEpoch(mixed $unixTimestamp) : DateTimeResult`

The Unix epoch (or Unix time or POSIX time or Unix
timestamp) is the number of seconds that have elapsed since
January 1, 1970 (midnight UTC/GMT).

##### Parameters

$unixTimestamp
: mixed

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html\#method_fromEpoch\#tags)

throwsException

##### Return values

[DateTimeResult](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html)

#### fromISO8601()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html\#method_fromISO8601)

`
    public
            static        fromISO8601(mixed $iso8601Timestamp) : DateTimeResult`

##### Parameters

$iso8601Timestamp
: mixed

##### Return values

[DateTimeResult](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html)

#### fromTimestamp()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html\#method_fromTimestamp)

Create a new DateTimeResult from an unknown timestamp.

`
    public
            static        fromTimestamp(mixed $timestamp[, mixed $expectedFormat = null ]) : DateTimeResult`

##### Parameters

$timestamp
: mixed$expectedFormat
: mixed
= null

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html\#method_fromTimestamp\#tags)

throwsException

##### Return values

[DateTimeResult](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html)

#### jsonSerialize()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html\#method_jsonSerialize)

Serialize the date as an ISO 8601 date when serializing as JSON.

`
    public
                    jsonSerialize() : string`

##### Return values

string
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html#toc-methods)
- Methods
  - [\_\_toString()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html#method___toString)
  - [fromEpoch()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html#method_fromEpoch)
  - [fromISO8601()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html#method_fromISO8601)
  - [fromTimestamp()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html#method_fromTimestamp)
  - [jsonSerialize()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html#method_jsonSerialize)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.DateTimeResult.html#top)
