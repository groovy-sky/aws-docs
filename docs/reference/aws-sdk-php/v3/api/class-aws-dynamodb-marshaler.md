Menu

- [Aws](namespace-aws.md)
- [DynamoDb](namespace-aws-dynamodb.md)

## Marshaler        in package    - [Aws](package-aws.md)

Marshals and unmarshals JSON documents and PHP arrays into DynamoDB items.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html#method___construct)
: mixed Instantiates a DynamoDB Marshaler.[binary()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html#method_binary)
: [BinaryValue](class-aws-dynamodb-binaryvalue.md)Creates a special object to represent a DynamoDB binary (B) value.[marshalItem()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html#method_marshalItem)
: array<string\|int, mixed> Marshal a native PHP array of data to a DynamoDB item.[marshalJson()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html#method_marshalJson)
: array<string\|int, mixed> Marshal a JSON document from a string to a DynamoDB item.[marshalValue()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html#method_marshalValue)
: array<string\|int, mixed> Marshal a native PHP value into a DynamoDB attribute value.[number()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html#method_number)
: [NumberValue](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.NumberValue.html)Creates a special object to represent a DynamoDB number (N) value.[set()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html#method_set)
: [SetValue](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SetValue.html)Creates a special object to represent a DynamoDB set (SS/NS/BS) value.[unmarshalItem()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html#method_unmarshalItem)
: array<string\|int, mixed>\|stdClassUnmarshal an item from a DynamoDB operation result into a native PHP
array. If you set $mapAsObject to true, then a stdClass value will be
returned instead.[unmarshalJson()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html#method_unmarshalJson)
: string Unmarshal a document (item) from a DynamoDB operation result into a JSON
document string.[unmarshalValue()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html#method_unmarshalValue)
: mixed Unmarshal a value from a DynamoDB operation result into a native PHP
value. Will return a scalar, array, or (if you set $mapAsObject to true)
stdClass value.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html\#method___construct)

Instantiates a DynamoDB Marshaler.

`
    public
                    __construct([array<string|int, mixed> $options = [] ]) : mixed`

The following options are valid.

- ignore\_invalid: (bool) Set to `true` if invalid values should be
ignored (i.e., not included) during marshaling.
- nullify\_invalid: (bool) Set to `true` if invalid values should be set
to null.
- wrap\_numbers: (bool) Set to `true` to wrap numbers with `NumberValue`
objects during unmarshaling to preserve the precision.

##### Parameters

$options
: array<string\|int, mixed>
= \[\]

Marshaler options

#### binary()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html\#method_binary)

Creates a special object to represent a DynamoDB binary (B) value.

`
    public
                    binary(mixed $value) : BinaryValue`

This helps disambiguate binary values from string (S) values.

##### Parameters

$value
: mixed

A binary value compatible with Guzzle streams.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html\#method_binary\#tags)

seeStream::factory

##### Return values

[BinaryValue](class-aws-dynamodb-binaryvalue.md)

#### marshalItem()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html\#method_marshalItem)

Marshal a native PHP array of data to a DynamoDB item.

`
    public
                    marshalItem(array<string|int, mixed>|stdClass $item) : array<string|int, mixed>`

The result is an array formatted in the proper parameter structure
required by the DynamoDB API for items.

##### Parameters

$item
: array<string\|int, mixed>\|stdClass

An associative array of data.

##### Return values

array<string\|int, mixed>
—

Item formatted for DynamoDB.

#### marshalJson()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html\#method_marshalJson)

Marshal a JSON document from a string to a DynamoDB item.

`
    public
                    marshalJson(string $json) : array<string|int, mixed>`

The result is an array formatted in the proper parameter structure
required by the DynamoDB API for items.

##### Parameters

$json
: string

A valid JSON document.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html\#method_marshalJson\#tags)

throwsInvalidArgumentException

if the JSON is invalid.

##### Return values

array<string\|int, mixed>
—

Item formatted for DynamoDB.

#### marshalValue()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html\#method_marshalValue)

Marshal a native PHP value into a DynamoDB attribute value.

`
    public
                    marshalValue(mixed $value) : array<string|int, mixed>`

The result is an associative array that is formatted in the proper
`[TYPE => VALUE]` parameter structure required by the DynamoDB API.

##### Parameters

$value
: mixed

A scalar, array, or `stdClass` value.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html\#method_marshalValue\#tags)

throwsUnexpectedValueException

if the value cannot be marshaled.

##### Return values

array<string\|int, mixed>
—

Attribute formatted for DynamoDB.

#### number()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html\#method_number)

Creates a special object to represent a DynamoDB number (N) value.

`
    public
                    number(string|int|float $value) : NumberValue`

This helps maintain the precision of large integer/float in PHP.

##### Parameters

$value
: string\|int\|float

A number value.

##### Return values

[NumberValue](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.NumberValue.html)

#### set()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html\#method_set)

Creates a special object to represent a DynamoDB set (SS/NS/BS) value.

`
    public
                    set(array<string|int, mixed> $values) : SetValue`

This helps disambiguate set values from list (L) values.

##### Parameters

$values
: array<string\|int, mixed>

The values of the set.

##### Return values

[SetValue](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.SetValue.html)

#### unmarshalItem()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html\#method_unmarshalItem)

Unmarshal an item from a DynamoDB operation result into a native PHP
array. If you set $mapAsObject to true, then a stdClass value will be
returned instead.

`
    public
                    unmarshalItem(array<string|int, mixed> $data[, bool $mapAsObject = false ]) : array<string|int, mixed>|stdClass`

##### Parameters

$data
: array<string\|int, mixed>

Item from a DynamoDB result.

$mapAsObject
: bool
= false

Whether maps should be represented as stdClass.

##### Return values

array<string\|int, mixed>\|stdClass

#### unmarshalJson()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html\#method_unmarshalJson)

Unmarshal a document (item) from a DynamoDB operation result into a JSON
document string.

`
    public
                    unmarshalJson(array<string|int, mixed> $data[, int $jsonEncodeFlags = 0 ]) : string`

##### Parameters

$data
: array<string\|int, mixed>

Item/document from a DynamoDB result.

$jsonEncodeFlags
: int
= 0

Flags to use with `json_encode()`.

##### Return values

string

#### unmarshalValue()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html\#method_unmarshalValue)

Unmarshal a value from a DynamoDB operation result into a native PHP
value. Will return a scalar, array, or (if you set $mapAsObject to true)
stdClass value.

`
    public
                    unmarshalValue(array<string|int, mixed> $value[, bool $mapAsObject = false ]) : mixed`

##### Parameters

$value
: array<string\|int, mixed>

Value from a DynamoDB result.

$mapAsObject
: bool
= false

Whether maps should be represented as stdClass.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html\#method_unmarshalValue\#tags)

throwsUnexpectedValueException
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html#method___construct)
  - [binary()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html#method_binary)
  - [marshalItem()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html#method_marshalItem)
  - [marshalJson()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html#method_marshalJson)
  - [marshalValue()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html#method_marshalValue)
  - [number()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html#method_number)
  - [set()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html#method_set)
  - [unmarshalItem()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html#method_unmarshalItem)
  - [unmarshalJson()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html#method_unmarshalJson)
  - [unmarshalValue()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html#method_unmarshalValue)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.DynamoDb.Marshaler.html#top)
