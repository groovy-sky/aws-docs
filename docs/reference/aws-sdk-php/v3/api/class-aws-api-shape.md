Menu

- [Aws](namespace-aws.md)
- [Api](namespace-aws-api.md)

## Shape     extends [AbstractModel](class-aws-api-abstractmodel.md)   in package    - [Aws](package-aws.md)

Base class representing a modeled shape.

### Table of Contents  [header link](class-aws-api-shape-toc.md)

#### Methods  [header link](class-aws-api-shape-toc-methods.md)

[\_\_construct()](class-aws-api-abstractmodel.md#method___construct)
: mixed [create()](class-aws-api-shape-method-create.md)
: mixed Get a concrete shape for the given definition.[getContextParam()](class-aws-api-shape-method-getcontextparam.md)
: mixed Get a context param definition.[getName()](class-aws-api-shape-method-getname.md)
: string Get the name of the shape[getType()](class-aws-api-shape-method-gettype.md)
: string Get the type of the shape[offsetExists()](class-aws-api-abstractmodel.md#method_offsetExists)
: bool [offsetGet()](class-aws-api-abstractmodel.md#method_offsetGet)
: mixed\|null [offsetSet()](class-aws-api-abstractmodel.md#method_offsetSet)
: void [offsetUnset()](class-aws-api-abstractmodel.md#method_offsetUnset)
: void [toArray()](class-aws-api-abstractmodel.md#method_toArray)
: mixed

### Methods  [header link](class-aws-api-shape-methods.md)

#### \_\_construct()  [header link](class-aws-api-abstractmodel.md\#method___construct)

`
    public
                    __construct(array<string|int, mixed> $definition, ShapeMap $shapeMap) : mixed`

##### Parameters

$definition
: array<string\|int, mixed>

Service description

$shapeMap
: [ShapeMap](class-aws-api-shapemap.md)

Shapemap used for creating shapes

#### create()  [header link](class-aws-api-shape-method-create.md)

Get a concrete shape for the given definition.

`
    public
            static        create(array<string|int, mixed> $definition, ShapeMap $shapeMap) : mixed`

##### Parameters

$definition
: array<string\|int, mixed>$shapeMap
: [ShapeMap](class-aws-api-shapemap.md)

##### Tags  [header link](class-aws-api-shape-method-create-tags.md)

throwsRuntimeException

if the type is invalid

#### getContextParam()  [header link](class-aws-api-shape-method-getcontextparam.md)

Get a context param definition.

`
    public
                    getContextParam() : mixed`

#### getName()  [header link](class-aws-api-shape-method-getname.md)

Get the name of the shape

`
    public
                    getName() : string`

##### Return values

string

#### getType()  [header link](class-aws-api-shape-method-gettype.md)

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
  - [Methods](class-aws-api-shape-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-api-abstractmodel.md#method___construct)
  - [create()](class-aws-api-shape-method-create.md)
  - [getContextParam()](class-aws-api-shape-method-getcontextparam.md)
  - [getName()](class-aws-api-shape-method-getname.md)
  - [getType()](class-aws-api-shape-method-gettype.md)
  - [offsetExists()](class-aws-api-abstractmodel.md#method_offsetExists)
  - [offsetGet()](class-aws-api-abstractmodel.md#method_offsetGet)
  - [offsetSet()](class-aws-api-abstractmodel.md#method_offsetSet)
  - [offsetUnset()](class-aws-api-abstractmodel.md#method_offsetUnset)
  - [toArray()](class-aws-api-abstractmodel.md#method_toArray)

[Back To Top](class-aws-api-shape-top.md)
