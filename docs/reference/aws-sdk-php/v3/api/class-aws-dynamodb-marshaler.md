Menu

- [Aws](namespace-aws.md)
- [DynamoDb](namespace-aws-dynamodb.md)

## Marshaler        in package    - [Aws](package-aws.md)

Marshals and unmarshals JSON documents and PHP arrays into DynamoDB items.

### Table of Contents  [header link](class-aws-dynamodb-marshaler-toc.md)

#### Methods  [header link](class-aws-dynamodb-marshaler-toc-methods.md)

[\_\_construct()](class-aws-dynamodb-marshaler-method-construct.md)
: mixed Instantiates a DynamoDB Marshaler.[binary()](class-aws-dynamodb-marshaler-method-binary.md)
: [BinaryValue](class-aws-dynamodb-binaryvalue.md)Creates a special object to represent a DynamoDB binary (B) value.[marshalItem()](class-aws-dynamodb-marshaler-method-marshalitem.md)
: array<string\|int, mixed> Marshal a native PHP array of data to a DynamoDB item.[marshalJson()](class-aws-dynamodb-marshaler-method-marshaljson.md)
: array<string\|int, mixed> Marshal a JSON document from a string to a DynamoDB item.[marshalValue()](class-aws-dynamodb-marshaler-method-marshalvalue.md)
: array<string\|int, mixed> Marshal a native PHP value into a DynamoDB attribute value.[number()](class-aws-dynamodb-marshaler-method-number.md)
: [NumberValue](class-aws-dynamodb-numbervalue.md)Creates a special object to represent a DynamoDB number (N) value.[set()](class-aws-dynamodb-marshaler-method-set.md)
: [SetValue](class-aws-dynamodb-setvalue.md)Creates a special object to represent a DynamoDB set (SS/NS/BS) value.[unmarshalItem()](class-aws-dynamodb-marshaler-method-unmarshalitem.md)
: array<string\|int, mixed>\|stdClassUnmarshal an item from a DynamoDB operation result into a native PHP
array. If you set $mapAsObject to true, then a stdClass value will be
returned instead.[unmarshalJson()](class-aws-dynamodb-marshaler-method-unmarshaljson.md)
: string Unmarshal a document (item) from a DynamoDB operation result into a JSON
document string.[unmarshalValue()](class-aws-dynamodb-marshaler-method-unmarshalvalue.md)
: mixed Unmarshal a value from a DynamoDB operation result into a native PHP
value. Will return a scalar, array, or (if you set $mapAsObject to true)
stdClass value.

### Methods  [header link](class-aws-dynamodb-marshaler-methods.md)

#### \_\_construct()  [header link](class-aws-dynamodb-marshaler-method-construct.md)

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

#### binary()  [header link](class-aws-dynamodb-marshaler-method-binary.md)

Creates a special object to represent a DynamoDB binary (B) value.

`
    public
                    binary(mixed $value) : BinaryValue`

This helps disambiguate binary values from string (S) values.

##### Parameters

$value
: mixed

A binary value compatible with Guzzle streams.

##### Tags  [header link](class-aws-dynamodb-marshaler-method-binary-tags.md)

seeStream::factory

##### Return values

[BinaryValue](class-aws-dynamodb-binaryvalue.md)

#### marshalItem()  [header link](class-aws-dynamodb-marshaler-method-marshalitem.md)

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

#### marshalJson()  [header link](class-aws-dynamodb-marshaler-method-marshaljson.md)

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

##### Tags  [header link](class-aws-dynamodb-marshaler-method-marshaljson-tags.md)

throwsInvalidArgumentException

if the JSON is invalid.

##### Return values

array<string\|int, mixed>
—

Item formatted for DynamoDB.

#### marshalValue()  [header link](class-aws-dynamodb-marshaler-method-marshalvalue.md)

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

##### Tags  [header link](class-aws-dynamodb-marshaler-method-marshalvalue-tags.md)

throwsUnexpectedValueException

if the value cannot be marshaled.

##### Return values

array<string\|int, mixed>
—

Attribute formatted for DynamoDB.

#### number()  [header link](class-aws-dynamodb-marshaler-method-number.md)

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

[NumberValue](class-aws-dynamodb-numbervalue.md)

#### set()  [header link](class-aws-dynamodb-marshaler-method-set.md)

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

[SetValue](class-aws-dynamodb-setvalue.md)

#### unmarshalItem()  [header link](class-aws-dynamodb-marshaler-method-unmarshalitem.md)

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

#### unmarshalJson()  [header link](class-aws-dynamodb-marshaler-method-unmarshaljson.md)

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

#### unmarshalValue()  [header link](class-aws-dynamodb-marshaler-method-unmarshalvalue.md)

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

##### Tags  [header link](class-aws-dynamodb-marshaler-method-unmarshalvalue-tags.md)

throwsUnexpectedValueException
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-dynamodb-marshaler-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-dynamodb-marshaler-method-construct.md)
  - [binary()](class-aws-dynamodb-marshaler-method-binary.md)
  - [marshalItem()](class-aws-dynamodb-marshaler-method-marshalitem.md)
  - [marshalJson()](class-aws-dynamodb-marshaler-method-marshaljson.md)
  - [marshalValue()](class-aws-dynamodb-marshaler-method-marshalvalue.md)
  - [number()](class-aws-dynamodb-marshaler-method-number.md)
  - [set()](class-aws-dynamodb-marshaler-method-set.md)
  - [unmarshalItem()](class-aws-dynamodb-marshaler-method-unmarshalitem.md)
  - [unmarshalJson()](class-aws-dynamodb-marshaler-method-unmarshaljson.md)
  - [unmarshalValue()](class-aws-dynamodb-marshaler-method-unmarshalvalue.md)

[Back To Top](class-aws-dynamodb-marshaler-top.md)
