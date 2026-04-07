Menu

- [Aws](namespace-aws.md)
- [Api](namespace-aws-api.md)

## ShapeMap        in package    - [Aws](package-aws.md)       implements  ArrayAccess

Builds shape based on shape references.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html\#toc-interfaces)

ArrayAccess

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html#method___construct)
: mixed [getShapeNames()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html#method_getShapeNames)
: array<string\|int, mixed> Get an array of shape names.[offsetExists()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html#method_offsetExists)
: bool [offsetGet()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html#method_offsetGet)
: mixed [offsetSet()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html#method_offsetSet)
: void [offsetUnset()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html#method_offsetUnset)
: void [resolve()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html#method_resolve)
: [Shape](class-aws-api-shape.md)Resolve a shape reference

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html\#method___construct)

`
    public
                    __construct(array<string|int, mixed> $shapeModels) : mixed`

##### Parameters

$shapeModels
: array<string\|int, mixed>

Associative array of shape definitions.

#### getShapeNames()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html\#method_getShapeNames)

Get an array of shape names.

`
    public
                    getShapeNames() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### offsetExists()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html\#method_offsetExists)

`
    public
                    offsetExists(mixed $offset) : bool`

##### Parameters

$offset
: mixed

##### Return values

bool

#### offsetGet()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html\#method_offsetGet)

`
    public
                    offsetGet(mixed $offset) : mixed`

##### Parameters

$offset
: mixed

#### offsetSet()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html\#method_offsetSet)

`
    public
                    offsetSet(mixed $offset, mixed $value) : void`

##### Parameters

$offset
: mixed$value
: mixed

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html\#method_offsetSet\#tags)

throwsBadMethodCallException

#### offsetUnset()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html\#method_offsetUnset)

`
    public
                    offsetUnset(mixed $offset) : void`

##### Parameters

$offset
: mixed

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html\#method_offsetUnset\#tags)

throwsBadMethodCallException

#### resolve()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html\#method_resolve)

Resolve a shape reference

`
    public
                    resolve(array<string|int, mixed> $shapeRef) : Shape`

##### Parameters

$shapeRef
: array<string\|int, mixed>

Shape reference shape

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html\#method_resolve\#tags)

throwsInvalidArgumentException

##### Return values

[Shape](class-aws-api-shape.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html#method___construct)
  - [getShapeNames()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html#method_getShapeNames)
  - [offsetExists()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html#method_offsetExists)
  - [offsetGet()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html#method_offsetGet)
  - [offsetSet()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html#method_offsetSet)
  - [offsetUnset()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html#method_offsetUnset)
  - [resolve()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html#method_resolve)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html#top)
