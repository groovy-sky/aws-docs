Menu

- [Aws](namespace-aws.md)
- [Api](namespace-aws-api.md)

## Shape     extends [AbstractModel](class-aws-api-abstractmodel.md)   in package    - [Aws](package-aws.md)

Base class representing a modeled shape.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Shape.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Shape.html\#toc-methods)

[\_\_construct()](class-aws-api-abstractmodel.md#method___construct)
: mixed [create()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Shape.html#method_create)
: mixed Get a concrete shape for the given definition.[getContextParam()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Shape.html#method_getContextParam)
: mixed Get a context param definition.[getName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Shape.html#method_getName)
: string Get the name of the shape[getType()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Shape.html#method_getType)
: string Get the type of the shape[offsetExists()](class-aws-api-abstractmodel.md#method_offsetExists)
: bool [offsetGet()](class-aws-api-abstractmodel.md#method_offsetGet)
: mixed\|null [offsetSet()](class-aws-api-abstractmodel.md#method_offsetSet)
: void [offsetUnset()](class-aws-api-abstractmodel.md#method_offsetUnset)
: void [toArray()](class-aws-api-abstractmodel.md#method_toArray)
: mixed

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Shape.html\#methods)

#### \_\_construct()  [header link](class-aws-api-abstractmodel.md\#method___construct)

`
    public
                    __construct(array<string|int, mixed> $definition, ShapeMap $shapeMap) : mixed`

##### Parameters

$definition
: array<string\|int, mixed>

Service description

$shapeMap
: [ShapeMap](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html)

Shapemap used for creating shapes

#### create()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Shape.html\#method_create)

Get a concrete shape for the given definition.

`
    public
            static        create(array<string|int, mixed> $definition, ShapeMap $shapeMap) : mixed`

##### Parameters

$definition
: array<string\|int, mixed>$shapeMap
: [ShapeMap](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html)

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Shape.html\#method_create\#tags)

throwsRuntimeException

if the type is invalid

#### getContextParam()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Shape.html\#method_getContextParam)

Get a context param definition.

`
    public
                    getContextParam() : mixed`

#### getName()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Shape.html\#method_getName)

Get the name of the shape

`
    public
                    getName() : string`

##### Return values

string

#### getType()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Shape.html\#method_getType)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Shape.html#toc-methods)
- Methods
  - [\_\_construct()](class-aws-api-abstractmodel.md#method___construct)
  - [create()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Shape.html#method_create)
  - [getContextParam()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Shape.html#method_getContextParam)
  - [getName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Shape.html#method_getName)
  - [getType()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Shape.html#method_getType)
  - [offsetExists()](class-aws-api-abstractmodel.md#method_offsetExists)
  - [offsetGet()](class-aws-api-abstractmodel.md#method_offsetGet)
  - [offsetSet()](class-aws-api-abstractmodel.md#method_offsetSet)
  - [offsetUnset()](class-aws-api-abstractmodel.md#method_offsetUnset)
  - [toArray()](class-aws-api-abstractmodel.md#method_toArray)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Shape.html#top)
