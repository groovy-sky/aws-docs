Menu

- [Aws](namespace-aws.md)
- [Api](namespace-aws-api.md)

## ShapeMap        in package    - [Aws](package-aws.md)       implements  ArrayAccess

Builds shape based on shape references.

### Table of Contents  [header link](class-aws-api-shapemap-toc.md)

#### Interfaces  [header link](class-aws-api-shapemap-toc-interfaces.md)

ArrayAccess

#### Methods  [header link](class-aws-api-shapemap-toc-methods.md)

[\_\_construct()](class-aws-api-shapemap-method-construct.md)
: mixed [getShapeNames()](class-aws-api-shapemap-method-getshapenames.md)
: array<string\|int, mixed> Get an array of shape names.[offsetExists()](class-aws-api-shapemap-method-offsetexists.md)
: bool [offsetGet()](class-aws-api-shapemap-method-offsetget.md)
: mixed [offsetSet()](class-aws-api-shapemap-method-offsetset.md)
: void [offsetUnset()](class-aws-api-shapemap-method-offsetunset.md)
: void [resolve()](class-aws-api-shapemap-method-resolve.md)
: [Shape](class-aws-api-shape.md)Resolve a shape reference

### Methods  [header link](class-aws-api-shapemap-methods.md)

#### \_\_construct()  [header link](class-aws-api-shapemap-method-construct.md)

`
    public
                    __construct(array<string|int, mixed> $shapeModels) : mixed`

##### Parameters

$shapeModels
: array<string\|int, mixed>

Associative array of shape definitions.

#### getShapeNames()  [header link](class-aws-api-shapemap-method-getshapenames.md)

Get an array of shape names.

`
    public
                    getShapeNames() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### offsetExists()  [header link](class-aws-api-shapemap-method-offsetexists.md)

`
    public
                    offsetExists(mixed $offset) : bool`

##### Parameters

$offset
: mixed

##### Return values

bool

#### offsetGet()  [header link](class-aws-api-shapemap-method-offsetget.md)

`
    public
                    offsetGet(mixed $offset) : mixed`

##### Parameters

$offset
: mixed

#### offsetSet()  [header link](class-aws-api-shapemap-method-offsetset.md)

`
    public
                    offsetSet(mixed $offset, mixed $value) : void`

##### Parameters

$offset
: mixed$value
: mixed

##### Tags  [header link](class-aws-api-shapemap-method-offsetset-tags.md)

throwsBadMethodCallException

#### offsetUnset()  [header link](class-aws-api-shapemap-method-offsetunset.md)

`
    public
                    offsetUnset(mixed $offset) : void`

##### Parameters

$offset
: mixed

##### Tags  [header link](class-aws-api-shapemap-method-offsetunset-tags.md)

throwsBadMethodCallException

#### resolve()  [header link](class-aws-api-shapemap-method-resolve.md)

Resolve a shape reference

`
    public
                    resolve(array<string|int, mixed> $shapeRef) : Shape`

##### Parameters

$shapeRef
: array<string\|int, mixed>

Shape reference shape

##### Tags  [header link](class-aws-api-shapemap-method-resolve-tags.md)

throwsInvalidArgumentException

##### Return values

[Shape](class-aws-api-shape.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-api-shapemap-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-api-shapemap-method-construct.md)
  - [getShapeNames()](class-aws-api-shapemap-method-getshapenames.md)
  - [offsetExists()](class-aws-api-shapemap-method-offsetexists.md)
  - [offsetGet()](class-aws-api-shapemap-method-offsetget.md)
  - [offsetSet()](class-aws-api-shapemap-method-offsetset.md)
  - [offsetUnset()](class-aws-api-shapemap-method-offsetunset.md)
  - [resolve()](class-aws-api-shapemap-method-resolve.md)

[Back To Top](class-aws-api-shapemap-top.md)
