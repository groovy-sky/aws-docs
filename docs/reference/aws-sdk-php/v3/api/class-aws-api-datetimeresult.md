Menu

- [Aws](namespace-aws.md)
- [Api](namespace-aws-api.md)

## DateTimeResult     extends DateTime   in package    - [Aws](package-aws.md)       implements  JsonSerializable

DateTime overrides that make DateTime work more seamlessly as a string,
with JSON documents, and with JMESPath.

### Table of Contents  [header link](class-aws-api-datetimeresult-toc.md)

#### Interfaces  [header link](class-aws-api-datetimeresult-toc-interfaces.md)

JsonSerializable

#### Methods  [header link](class-aws-api-datetimeresult-toc-methods.md)

[\_\_toString()](class-aws-api-datetimeresult-method-tostring.md)
: string Serialize the DateTimeResult as an ISO 8601 date string.[fromEpoch()](class-aws-api-datetimeresult-method-fromepoch.md)
: [DateTimeResult](class-aws-api-datetimeresult.md)Create a new DateTimeResult from a unix timestamp.[fromISO8601()](class-aws-api-datetimeresult-method-fromiso8601.md)
: [DateTimeResult](class-aws-api-datetimeresult.md)[fromTimestamp()](class-aws-api-datetimeresult-method-fromtimestamp.md)
: [DateTimeResult](class-aws-api-datetimeresult.md)Create a new DateTimeResult from an unknown timestamp.[jsonSerialize()](class-aws-api-datetimeresult-method-jsonserialize.md)
: string Serialize the date as an ISO 8601 date when serializing as JSON.

### Methods  [header link](class-aws-api-datetimeresult-methods.md)

#### \_\_toString()  [header link](class-aws-api-datetimeresult-method-tostring.md)

Serialize the DateTimeResult as an ISO 8601 date string.

`
    public
                    __toString() : string`

##### Return values

string

#### fromEpoch()  [header link](class-aws-api-datetimeresult-method-fromepoch.md)

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

##### Tags  [header link](class-aws-api-datetimeresult-method-fromepoch-tags.md)

throwsException

##### Return values

[DateTimeResult](class-aws-api-datetimeresult.md)

#### fromISO8601()  [header link](class-aws-api-datetimeresult-method-fromiso8601.md)

`
    public
            static        fromISO8601(mixed $iso8601Timestamp) : DateTimeResult`

##### Parameters

$iso8601Timestamp
: mixed

##### Return values

[DateTimeResult](class-aws-api-datetimeresult.md)

#### fromTimestamp()  [header link](class-aws-api-datetimeresult-method-fromtimestamp.md)

Create a new DateTimeResult from an unknown timestamp.

`
    public
            static        fromTimestamp(mixed $timestamp[, mixed $expectedFormat = null ]) : DateTimeResult`

##### Parameters

$timestamp
: mixed$expectedFormat
: mixed
= null

##### Tags  [header link](class-aws-api-datetimeresult-method-fromtimestamp-tags.md)

throwsException

##### Return values

[DateTimeResult](class-aws-api-datetimeresult.md)

#### jsonSerialize()  [header link](class-aws-api-datetimeresult-method-jsonserialize.md)

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
  - [Methods](class-aws-api-datetimeresult-toc-methods.md)
- Methods
  - [\_\_toString()](class-aws-api-datetimeresult-method-tostring.md)
  - [fromEpoch()](class-aws-api-datetimeresult-method-fromepoch.md)
  - [fromISO8601()](class-aws-api-datetimeresult-method-fromiso8601.md)
  - [fromTimestamp()](class-aws-api-datetimeresult-method-fromtimestamp.md)
  - [jsonSerialize()](class-aws-api-datetimeresult-method-jsonserialize.md)

[Back To Top](class-aws-api-datetimeresult-top.md)
