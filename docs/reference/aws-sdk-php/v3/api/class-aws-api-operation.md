Menu

- [Aws](namespace-aws.md)
- [Api](namespace-aws-api.md)

## Operation     extends [AbstractModel](class-aws-api-abstractmodel.md)   in package    - [Aws](package-aws.md)

Represents an API operation.

### Table of Contents  [header link](class-aws-api-operation-toc.md)

#### Methods  [header link](class-aws-api-operation-toc-methods.md)

[\_\_construct()](class-aws-api-operation-method-construct.md)
: mixed [getContextParams()](class-aws-api-operation-method-getcontextparams.md)
: array<string\|int, mixed> Gets definition of modeled dynamic values used
for endpoint resolution[getErrors()](class-aws-api-operation-method-geterrors.md)
: array<string\|int, [StructureShape](class-aws-api-structureshape.md) \> Get an array of operation error shapes.[getHttp()](class-aws-api-operation-method-gethttp.md)
: array<string\|int, mixed> Returns an associative array of the HTTP attribute of the operation:[getInput()](class-aws-api-operation-method-getinput.md)
: [StructureShape](class-aws-api-structureshape.md)Get the input shape of the operation.[getOperationContextParams()](class-aws-api-operation-method-getoperationcontextparams.md)
: array<string\|int, mixed> Gets definition of modeled dynamic values used
for endpoint resolution[getOutput()](class-aws-api-operation-method-getoutput.md)
: [StructureShape](class-aws-api-structureshape.md)Get the output shape of the operation.[getStaticContextParams()](class-aws-api-operation-method-getstaticcontextparams.md)
: array<string\|int, mixed> Gets static modeled static values used for
endpoint resolution.[offsetExists()](class-aws-api-abstractmodel.md#method_offsetExists)
: bool [offsetGet()](class-aws-api-abstractmodel.md#method_offsetGet)
: mixed\|null [offsetSet()](class-aws-api-abstractmodel.md#method_offsetSet)
: void [offsetUnset()](class-aws-api-abstractmodel.md#method_offsetUnset)
: void [toArray()](class-aws-api-abstractmodel.md#method_toArray)
: mixed

### Methods  [header link](class-aws-api-operation-methods.md)

#### \_\_construct()  [header link](class-aws-api-operation-method-construct.md)

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

#### getContextParams()  [header link](class-aws-api-operation-method-getcontextparams.md)

Gets definition of modeled dynamic values used
for endpoint resolution

`
    public
                    getContextParams() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getErrors()  [header link](class-aws-api-operation-method-geterrors.md)

Get an array of operation error shapes.

`
    public
                    getErrors() : array<string|int, StructureShape>`

##### Return values

array<string\|int, [StructureShape](class-aws-api-structureshape.md) >

#### getHttp()  [header link](class-aws-api-operation-method-gethttp.md)

Returns an associative array of the HTTP attribute of the operation:

`
    public
                    getHttp() : array<string|int, mixed>`

- method: HTTP method of the operation
- requestUri: URI of the request (can include URI template placeholders)

##### Return values

array<string\|int, mixed>

#### getInput()  [header link](class-aws-api-operation-method-getinput.md)

Get the input shape of the operation.

`
    public
                    getInput() : StructureShape`

##### Return values

[StructureShape](class-aws-api-structureshape.md)

#### getOperationContextParams()  [header link](class-aws-api-operation-method-getoperationcontextparams.md)

Gets definition of modeled dynamic values used
for endpoint resolution

`
    public
                    getOperationContextParams() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getOutput()  [header link](class-aws-api-operation-method-getoutput.md)

Get the output shape of the operation.

`
    public
                    getOutput() : StructureShape`

##### Return values

[StructureShape](class-aws-api-structureshape.md)

#### getStaticContextParams()  [header link](class-aws-api-operation-method-getstaticcontextparams.md)

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
  - [Methods](class-aws-api-operation-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-api-operation-method-construct.md)
  - [getContextParams()](class-aws-api-operation-method-getcontextparams.md)
  - [getErrors()](class-aws-api-operation-method-geterrors.md)
  - [getHttp()](class-aws-api-operation-method-gethttp.md)
  - [getInput()](class-aws-api-operation-method-getinput.md)
  - [getOperationContextParams()](class-aws-api-operation-method-getoperationcontextparams.md)
  - [getOutput()](class-aws-api-operation-method-getoutput.md)
  - [getStaticContextParams()](class-aws-api-operation-method-getstaticcontextparams.md)
  - [offsetExists()](class-aws-api-abstractmodel.md#method_offsetExists)
  - [offsetGet()](class-aws-api-abstractmodel.md#method_offsetGet)
  - [offsetSet()](class-aws-api-abstractmodel.md#method_offsetSet)
  - [offsetUnset()](class-aws-api-abstractmodel.md#method_offsetUnset)
  - [toArray()](class-aws-api-abstractmodel.md#method_toArray)

[Back To Top](class-aws-api-operation-top.md)
