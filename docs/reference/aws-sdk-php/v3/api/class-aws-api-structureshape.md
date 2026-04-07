Menu

- [Aws](namespace-aws.md)
- [Api](namespace-aws-api.md)

## StructureShape     extends [Shape](class-aws-api-shape.md)   in package    - [Aws](package-aws.md)

Represents a structure shape and resolve member shape references.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.StructureShape.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.StructureShape.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.StructureShape.html#method___construct)
: mixed [create()](class-aws-api-shape.md#method_create)
: mixed Get a concrete shape for the given definition.[getContextParam()](class-aws-api-shape.md#method_getContextParam)
: mixed Get a context param definition.[getMember()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.StructureShape.html#method_getMember)
: [Shape](class-aws-api-shape.md)Retrieve a member by name.[getMembers()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.StructureShape.html#method_getMembers)
: array<string\|int, [Shape](class-aws-api-shape.md) \> Gets a list of all members[getName()](class-aws-api-shape.md#method_getName)
: string Get the name of the shape[getOriginalDefinition()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.StructureShape.html#method_getOriginalDefinition)
: array<string\|int, mixed>\|null Used to look up a shape's original definition.[getType()](class-aws-api-shape.md#method_getType)
: string Get the type of the shape[hasMember()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.StructureShape.html#method_hasMember)
: bool Check if a specific member exists by name.[offsetExists()](class-aws-api-abstractmodel.md#method_offsetExists)
: bool [offsetGet()](class-aws-api-abstractmodel.md#method_offsetGet)
: mixed\|null [offsetSet()](class-aws-api-abstractmodel.md#method_offsetSet)
: void [offsetUnset()](class-aws-api-abstractmodel.md#method_offsetUnset)
: void [toArray()](class-aws-api-abstractmodel.md#method_toArray)
: mixed

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.StructureShape.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.StructureShape.html\#method___construct)

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

#### getContextParam()  [header link](class-aws-api-shape.md\#method_getContextParam)

Get a context param definition.

`
    public
                    getContextParam() : mixed`

#### getMember()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.StructureShape.html\#method_getMember)

Retrieve a member by name.

`
    public
                    getMember(string $name) : Shape`

##### Parameters

$name
: string

Name of the member to retrieve

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.StructureShape.html\#method_getMember\#tags)

throwsInvalidArgumentException

if the member is not found.

##### Return values

[Shape](class-aws-api-shape.md)

#### getMembers()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.StructureShape.html\#method_getMembers)

Gets a list of all members

`
    public
                    getMembers() : array<string|int, Shape>`

##### Return values

array<string\|int, [Shape](class-aws-api-shape.md) >

#### getName()  [header link](class-aws-api-shape.md\#method_getName)

Get the name of the shape

`
    public
                    getName() : string`

##### Return values

string

#### getOriginalDefinition()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.StructureShape.html\#method_getOriginalDefinition)

Used to look up a shape's original definition.

`
    public
                    getOriginalDefinition(string $name) : array<string|int, mixed>|null`

##### Parameters

$name
: string

##### Return values

array<string\|int, mixed>\|null

#### getType()  [header link](class-aws-api-shape.md\#method_getType)

Get the type of the shape

`
    public
                    getType() : string`

##### Return values

string

#### hasMember()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.StructureShape.html\#method_hasMember)

Check if a specific member exists by name.

`
    public
                    hasMember(string $name) : bool`

##### Parameters

$name
: string

Name of the member to check

##### Return values

bool

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.StructureShape.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.StructureShape.html#method___construct)
  - [create()](class-aws-api-shape.md#method_create)
  - [getContextParam()](class-aws-api-shape.md#method_getContextParam)
  - [getMember()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.StructureShape.html#method_getMember)
  - [getMembers()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.StructureShape.html#method_getMembers)
  - [getName()](class-aws-api-shape.md#method_getName)
  - [getOriginalDefinition()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.StructureShape.html#method_getOriginalDefinition)
  - [getType()](class-aws-api-shape.md#method_getType)
  - [hasMember()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.StructureShape.html#method_hasMember)
  - [offsetExists()](class-aws-api-abstractmodel.md#method_offsetExists)
  - [offsetGet()](class-aws-api-abstractmodel.md#method_offsetGet)
  - [offsetSet()](class-aws-api-abstractmodel.md#method_offsetSet)
  - [offsetUnset()](class-aws-api-abstractmodel.md#method_offsetUnset)
  - [toArray()](class-aws-api-abstractmodel.md#method_toArray)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.StructureShape.html#top)
