Menu

- [Aws](namespace-aws.md)
- [Api](namespace-aws-api.md)

## TimestampShape     extends [Shape](class-aws-api-shape.md)   in package    - [Aws](package-aws.md)

Represents a timestamp shape.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.TimestampShape.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.TimestampShape.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.TimestampShape.html#method___construct)
: mixed [create()](class-aws-api-shape.md#method_create)
: mixed Get a concrete shape for the given definition.[format()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.TimestampShape.html#method_format)
: int\|string Formats a timestamp value for a service.[getContextParam()](class-aws-api-shape.md#method_getContextParam)
: mixed Get a context param definition.[getName()](class-aws-api-shape.md#method_getName)
: string Get the name of the shape[getType()](class-aws-api-shape.md#method_getType)
: string Get the type of the shape[offsetExists()](class-aws-api-abstractmodel.md#method_offsetExists)
: bool [offsetGet()](class-aws-api-abstractmodel.md#method_offsetGet)
: mixed\|null [offsetSet()](class-aws-api-abstractmodel.md#method_offsetSet)
: void [offsetUnset()](class-aws-api-abstractmodel.md#method_offsetUnset)
: void [toArray()](class-aws-api-abstractmodel.md#method_toArray)
: mixed

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.TimestampShape.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.TimestampShape.html\#method___construct)

`
    public
                    __construct(array<string|int, mixed> $definition, ShapeMap $shapeMap) : mixed`

##### Parameters

$definition
: array<string\|int, mixed>$shapeMap
: [ShapeMap](class-aws-api-shapemap.md)

#### create()  [header link](class-aws-api-shape.md\#method_create)

Get a concrete shape for the given definition.

`
    public
            static        create(array<string|int, mixed> $definition, ShapeMap $shapeMap) : mixed`

##### Parameters

$definition
: array<string\|int, mixed>$shapeMap
: [ShapeMap](class-aws-api-shapemap.md)

##### Tags  [header link](class-aws-api-shape.md\#method_create\#tags)

throwsRuntimeException

if the type is invalid

#### format()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.TimestampShape.html\#method_format)

Formats a timestamp value for a service.

`
    public
            static        format(mixed $value, string $format) : int|string`

##### Parameters

$value
: mixed

Value to format

$format
: string

Format used to serialize the value

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.TimestampShape.html\#method_format\#tags)

throwsUnexpectedValueException

if the format is unknown.

throwsInvalidArgumentException

if the value is an unsupported type.

##### Return values

int\|string

#### getContextParam()  [header link](class-aws-api-shape.md\#method_getContextParam)

Get a context param definition.

`
    public
                    getContextParam() : mixed`

#### getName()  [header link](class-aws-api-shape.md\#method_getName)

Get the name of the shape

`
    public
                    getName() : string`

##### Return values

string

#### getType()  [header link](class-aws-api-shape.md\#method_getType)

Get the type of the shape

`
    public
                    getType() : string`

##### Return values

string

#### offsetExists()  [header link](class-aws-api-abstractmodel.md\#method_offsetExists)

`
    public
                    offsetExists(mixed $offset) : bool`

##### Parameters

$offset
: mixed

##### Return values

bool

#### offsetGet()  [header link](class-aws-api-abstractmodel.md\#method_offsetGet)

`
    public
                    offsetGet(mixed $offset) : mixed|null`

##### Parameters

$offset
: mixed

##### Return values

mixed\|null

#### offsetSet()  [header link](class-aws-api-abstractmodel.md\#method_offsetSet)

`
    public
                    offsetSet(mixed $offset, mixed $value) : void`

##### Parameters

$offset
: mixed$value
: mixed

#### offsetUnset()  [header link](class-aws-api-abstractmodel.md\#method_offsetUnset)

`
    public
                    offsetUnset(mixed $offset) : void`

##### Parameters

$offset
: mixed

#### toArray()  [header link](class-aws-api-abstractmodel.md\#method_toArray)

`
    public
                    toArray() : mixed`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.TimestampShape.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.TimestampShape.html#method___construct)
  - [create()](class-aws-api-shape.md#method_create)
  - [format()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.TimestampShape.html#method_format)
  - [getContextParam()](class-aws-api-shape.md#method_getContextParam)
  - [getName()](class-aws-api-shape.md#method_getName)
  - [getType()](class-aws-api-shape.md#method_getType)
  - [offsetExists()](class-aws-api-abstractmodel.md#method_offsetExists)
  - [offsetGet()](class-aws-api-abstractmodel.md#method_offsetGet)
  - [offsetSet()](class-aws-api-abstractmodel.md#method_offsetSet)
  - [offsetUnset()](class-aws-api-abstractmodel.md#method_offsetUnset)
  - [toArray()](class-aws-api-abstractmodel.md#method_toArray)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.TimestampShape.html#top)
