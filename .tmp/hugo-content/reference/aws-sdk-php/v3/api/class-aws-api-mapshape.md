Menu

- [Aws](namespace-aws.md)
- [Api](namespace-aws-api.md)

## MapShape     extends [Shape](class-aws-api-shape.md)   in package    - [Aws](package-aws.md)

Represents a map shape.

### Table of Contents  [header link](class-aws-api-mapshape-toc.md)

#### Methods  [header link](class-aws-api-mapshape-toc-methods.md)

[\_\_construct()](class-aws-api-mapshape-method-construct.md)
: mixed [create()](class-aws-api-shape-method-create.md)
: mixed Get a concrete shape for the given definition.[getContextParam()](class-aws-api-shape-method-getcontextparam.md)
: mixed Get a context param definition.[getKey()](class-aws-api-mapshape-method-getkey.md)
: [Shape](class-aws-api-shape.md)[getName()](class-aws-api-shape-method-getname.md)
: string Get the name of the shape[getType()](class-aws-api-shape-method-gettype.md)
: string Get the type of the shape[getValue()](class-aws-api-mapshape-method-getvalue.md)
: [Shape](class-aws-api-shape.md)[offsetExists()](class-aws-api-abstractmodel.md#method_offsetExists)
: bool [offsetGet()](class-aws-api-abstractmodel.md#method_offsetGet)
: mixed\|null [offsetSet()](class-aws-api-abstractmodel.md#method_offsetSet)
: void [offsetUnset()](class-aws-api-abstractmodel.md#method_offsetUnset)
: void [toArray()](class-aws-api-abstractmodel.md#method_toArray)
: mixed

### Methods  [header link](class-aws-api-mapshape-methods.md)

#### \_\_construct()  [header link](class-aws-api-mapshape-method-construct.md)

`
    public
                    __construct(array<string|int, mixed> $definition, ShapeMap $shapeMap) : mixed`

##### Parameters

$definition
: array<string\|int, mixed>$shapeMap
: [ShapeMap](class-aws-api-shapemap.md)

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

#### getKey()  [header link](class-aws-api-mapshape-method-getkey.md)

`
    public
                    getKey() : Shape`

##### Return values

[Shape](class-aws-api-shape.md)

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

#### getValue()  [header link](class-aws-api-mapshape-method-getvalue.md)

`
    public
                    getValue() : Shape`

##### Tags  [header link](class-aws-api-mapshape-method-getvalue-tags.md)

throwsRuntimeException

if no value is specified

##### Return values

[Shape](class-aws-api-shape.md)

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
  - [Methods](class-aws-api-mapshape-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-api-mapshape-method-construct.md)
  - [create()](class-aws-api-shape-method-create.md)
  - [getContextParam()](class-aws-api-shape-method-getcontextparam.md)
  - [getKey()](class-aws-api-mapshape-method-getkey.md)
  - [getName()](class-aws-api-shape-method-getname.md)
  - [getType()](class-aws-api-shape-method-gettype.md)
  - [getValue()](class-aws-api-mapshape-method-getvalue.md)
  - [offsetExists()](class-aws-api-abstractmodel.md#method_offsetExists)
  - [offsetGet()](class-aws-api-abstractmodel.md#method_offsetGet)
  - [offsetSet()](class-aws-api-abstractmodel.md#method_offsetSet)
  - [offsetUnset()](class-aws-api-abstractmodel.md#method_offsetUnset)
  - [toArray()](class-aws-api-abstractmodel.md#method_toArray)

[Back To Top](class-aws-api-mapshape-top.md)
