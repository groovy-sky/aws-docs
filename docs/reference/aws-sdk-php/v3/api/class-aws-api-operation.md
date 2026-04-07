Menu

- [Aws](namespace-aws.md)
- [Api](namespace-aws-api.md)

## Operation     extends [AbstractModel](class-aws-api-abstractmodel.md)   in package    - [Aws](package-aws.md)

Represents an API operation.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html#method___construct)
: mixed [getContextParams()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html#method_getContextParams)
: array<string\|int, mixed> Gets definition of modeled dynamic values used
for endpoint resolution[getErrors()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html#method_getErrors)
: array<string\|int, [StructureShape](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.StructureShape.html) \> Get an array of operation error shapes.[getHttp()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html#method_getHttp)
: array<string\|int, mixed> Returns an associative array of the HTTP attribute of the operation:[getInput()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html#method_getInput)
: [StructureShape](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.StructureShape.html)Get the input shape of the operation.[getOperationContextParams()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html#method_getOperationContextParams)
: array<string\|int, mixed> Gets definition of modeled dynamic values used
for endpoint resolution[getOutput()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html#method_getOutput)
: [StructureShape](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.StructureShape.html)Get the output shape of the operation.[getStaticContextParams()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html#method_getStaticContextParams)
: array<string\|int, mixed> Gets static modeled static values used for
endpoint resolution.[offsetExists()](class-aws-api-abstractmodel.md#method_offsetExists)
: bool [offsetGet()](class-aws-api-abstractmodel.md#method_offsetGet)
: mixed\|null [offsetSet()](class-aws-api-abstractmodel.md#method_offsetSet)
: void [offsetUnset()](class-aws-api-abstractmodel.md#method_offsetUnset)
: void [toArray()](class-aws-api-abstractmodel.md#method_toArray)
: mixed

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html\#method___construct)

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

#### getContextParams()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html\#method_getContextParams)

Gets definition of modeled dynamic values used
for endpoint resolution

`
    public
                    getContextParams() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getErrors()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html\#method_getErrors)

Get an array of operation error shapes.

`
    public
                    getErrors() : array<string|int, StructureShape>`

##### Return values

array<string\|int, [StructureShape](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.StructureShape.html) >

#### getHttp()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html\#method_getHttp)

Returns an associative array of the HTTP attribute of the operation:

`
    public
                    getHttp() : array<string|int, mixed>`

- method: HTTP method of the operation
- requestUri: URI of the request (can include URI template placeholders)

##### Return values

array<string\|int, mixed>

#### getInput()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html\#method_getInput)

Get the input shape of the operation.

`
    public
                    getInput() : StructureShape`

##### Return values

[StructureShape](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.StructureShape.html)

#### getOperationContextParams()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html\#method_getOperationContextParams)

Gets definition of modeled dynamic values used
for endpoint resolution

`
    public
                    getOperationContextParams() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getOutput()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html\#method_getOutput)

Get the output shape of the operation.

`
    public
                    getOutput() : StructureShape`

##### Return values

[StructureShape](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.StructureShape.html)

#### getStaticContextParams()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html\#method_getStaticContextParams)

Gets static modeled static values used for
endpoint resolution.

`
    public
                    getStaticContextParams() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html#method___construct)
  - [getContextParams()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html#method_getContextParams)
  - [getErrors()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html#method_getErrors)
  - [getHttp()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html#method_getHttp)
  - [getInput()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html#method_getInput)
  - [getOperationContextParams()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html#method_getOperationContextParams)
  - [getOutput()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html#method_getOutput)
  - [getStaticContextParams()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html#method_getStaticContextParams)
  - [offsetExists()](class-aws-api-abstractmodel.md#method_offsetExists)
  - [offsetGet()](class-aws-api-abstractmodel.md#method_offsetGet)
  - [offsetSet()](class-aws-api-abstractmodel.md#method_offsetSet)
  - [offsetUnset()](class-aws-api-abstractmodel.md#method_offsetUnset)
  - [toArray()](class-aws-api-abstractmodel.md#method_toArray)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html#top)
