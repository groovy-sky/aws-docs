Menu

- [Aws](namespace-aws.md)
- [Api](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.api.html)

## Service     extends [AbstractModel](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.AbstractModel.html)   in package    - [Aws](package-aws.md)

Represents a web service API model.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.AbstractModel.html#method___construct)
: mixed [createErrorParser()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_createErrorParser)
: callable Creates an error parser for the given protocol.[createParser()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_createParser)
: callable Applies the listeners needed to parse client models.[createSerializer()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_createSerializer)
: callable Creates a request serializer for the provided API object.[getApiVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getApiVersion)
: string Get the API version of the service[getClientContextParams()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getClientContextParams)
: array<string\|int, mixed> Get all the context params of the description.[getDefinition()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getDefinition)
: callable Get the service's definition.[getEndpointPrefix()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getEndpointPrefix)
: string Get the API version of the service[getErrorShapes()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getErrorShapes)
: array<string\|int, mixed> Get all of the error shapes of the service[getMetadata()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getMetadata)
: mixed Get all of the service metadata or a specific metadata key value.[getOperation()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getOperation)
: [Operation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html)Get an operation by name.[getOperations()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getOperations)
: array<string\|int, [Operation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html) \> Get all of the operations of the description.[getPaginatorConfig()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getPaginatorConfig)
: array<string\|int, mixed> Retrieve a paginator by name.[getPaginators()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getPaginators)
: array<string\|int, mixed> Gets an associative array of available paginator configurations where
the key is the name of the paginator, and the value is the paginator
configuration.[getProtocol()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getProtocol)
: string Get the protocol used by the service.[getProvider()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getProvider)
: callable Get the service's api provider.[getServiceFullName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getServiceFullName)
: string Get the full name of the service[getServiceId()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getServiceId)
: string Get the service id[getServiceName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getServiceName)
: string Get the service name.[getShapeMap()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getShapeMap)
: [ShapeMap](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html)Get the shape map used by the API.[getSignatureVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getSignatureVersion)
: string Get the default signature version of the service.[getSigningName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getSigningName)
: string Get the signing name used by the service.[getUid()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getUid)
: string Get the uid string used by the service[getWaiterConfig()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getWaiterConfig)
: array<string\|int, mixed> Get a waiter configuration by name.[getWaiters()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getWaiters)
: array<string\|int, mixed> Gets an associative array of available waiter configurations where the
key is the name of the waiter, and the value is the waiter
configuration.[hasOperation()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_hasOperation)
: bool Check if the description has a specific operation by name.[hasPaginator()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_hasPaginator)
: bool Determines if the service has a paginator by name.[hasWaiter()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_hasWaiter)
: bool Determines if the service has a waiter by name.[offsetExists()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.AbstractModel.html#method_offsetExists)
: bool [offsetGet()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.AbstractModel.html#method_offsetGet)
: mixed\|null [offsetSet()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.AbstractModel.html#method_offsetSet)
: void [offsetUnset()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.AbstractModel.html#method_offsetUnset)
: void [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.AbstractModel.html#method_toArray)
: mixed

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.AbstractModel.html\#method___construct)

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

#### createErrorParser()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_createErrorParser)

Creates an error parser for the given protocol.

`
    public
            static        createErrorParser(string $protocol[, Service|null $api = null ]) : callable`

Redundant method signature to preserve backwards compatibility.

##### Parameters

$protocol
: string

Protocol to parse (e.g., query, json, etc.)

$api
: [Service](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html) \|null
= null

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_createErrorParser\#tags)

throwsUnexpectedValueException

##### Return values

callable

#### createParser()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_createParser)

Applies the listeners needed to parse client models.

`
    public
            static        createParser(Service $api) : callable`

##### Parameters

$api
: [Service](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html)

API to create a parser for

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_createParser\#tags)

throwsUnexpectedValueException

##### Return values

callable

#### createSerializer()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_createSerializer)

Creates a request serializer for the provided API object.

`
    public
            static        createSerializer(Service $api, string $endpoint) : callable`

##### Parameters

$api
: [Service](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html)

API that contains a protocol.

$endpoint
: string

Endpoint to send requests to.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_createSerializer\#tags)

throwsUnexpectedValueException

##### Return values

callable

#### getApiVersion()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_getApiVersion)

Get the API version of the service

`
    public
                    getApiVersion() : string`

##### Return values

string

#### getClientContextParams()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_getClientContextParams)

Get all the context params of the description.

`
    public
                    getClientContextParams() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getDefinition()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_getDefinition)

Get the service's definition.

`
    public
                    getDefinition() : callable`

##### Return values

callable

#### getEndpointPrefix()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_getEndpointPrefix)

Get the API version of the service

`
    public
                    getEndpointPrefix() : string`

##### Return values

string

#### getErrorShapes()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_getErrorShapes)

Get all of the error shapes of the service

`
    public
                    getErrorShapes() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getMetadata()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_getMetadata)

Get all of the service metadata or a specific metadata key value.

`
    public
                    getMetadata([string|null $key = null ]) : mixed`

##### Parameters

$key
: string\|null
= null

Key to retrieve or null to retrieve all metadata

##### Return values

mixed
—

Returns the result or null if the key is not found

#### getOperation()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_getOperation)

Get an operation by name.

`
    public
                    getOperation(string $name) : Operation`

##### Parameters

$name
: string

Operation to retrieve by name

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_getOperation\#tags)

throwsInvalidArgumentException

If the operation is not found

##### Return values

[Operation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html)

#### getOperations()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_getOperations)

Get all of the operations of the description.

`
    public
                    getOperations() : array<string|int, Operation>`

##### Return values

array<string\|int, [Operation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Operation.html) >

#### getPaginatorConfig()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_getPaginatorConfig)

Retrieve a paginator by name.

`
    public
                    getPaginatorConfig(string $name) : array<string|int, mixed>`

##### Parameters

$name
: string

Paginator to retrieve by name. This argument is
typically the operation name.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_getPaginatorConfig\#tags)

throwsUnexpectedValueException

if the paginator does not exist.

unstable

The configuration format of paginators may change in the future

##### Return values

array<string\|int, mixed>

#### getPaginators()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_getPaginators)

Gets an associative array of available paginator configurations where
the key is the name of the paginator, and the value is the paginator
configuration.

`
    public
                    getPaginators() : array<string|int, mixed>`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_getPaginators\#tags)

unstable

The configuration format of paginators may change in the future

##### Return values

array<string\|int, mixed>

#### getProtocol()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_getProtocol)

Get the protocol used by the service.

`
    public
                    getProtocol() : string`

##### Return values

string

#### getProvider()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_getProvider)

Get the service's api provider.

`
    public
                    getProvider() : callable`

##### Return values

callable

#### getServiceFullName()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_getServiceFullName)

Get the full name of the service

`
    public
                    getServiceFullName() : string`

##### Return values

string

#### getServiceId()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_getServiceId)

Get the service id

`
    public
                    getServiceId() : string`

##### Return values

string

#### getServiceName()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_getServiceName)

Get the service name.

`
    public
                    getServiceName() : string`

##### Return values

string

#### getShapeMap()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_getShapeMap)

Get the shape map used by the API.

`
    public
                    getShapeMap() : ShapeMap`

##### Return values

[ShapeMap](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.ShapeMap.html)

#### getSignatureVersion()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_getSignatureVersion)

Get the default signature version of the service.

`
    public
                    getSignatureVersion() : string`

Note: this method assumes "v4" when not specified in the model.

##### Return values

string

#### getSigningName()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_getSigningName)

Get the signing name used by the service.

`
    public
                    getSigningName() : string`

##### Return values

string

#### getUid()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_getUid)

Get the uid string used by the service

`
    public
                    getUid() : string`

##### Return values

string

#### getWaiterConfig()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_getWaiterConfig)

Get a waiter configuration by name.

`
    public
                    getWaiterConfig(string $name) : array<string|int, mixed>`

##### Parameters

$name
: string

Name of the waiter by name.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_getWaiterConfig\#tags)

throwsUnexpectedValueException

if the waiter does not exist.

##### Return values

array<string\|int, mixed>

#### getWaiters()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_getWaiters)

Gets an associative array of available waiter configurations where the
key is the name of the waiter, and the value is the waiter
configuration.

`
    public
                    getWaiters() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### hasOperation()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_hasOperation)

Check if the description has a specific operation by name.

`
    public
                    hasOperation(string $name) : bool`

##### Parameters

$name
: string

Operation to check by name

##### Return values

bool

#### hasPaginator()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_hasPaginator)

Determines if the service has a paginator by name.

`
    public
                    hasPaginator(string $name) : bool`

##### Parameters

$name
: string

Name of the paginator.

##### Return values

bool

#### hasWaiter()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html\#method_hasWaiter)

Determines if the service has a waiter by name.

`
    public
                    hasWaiter(string $name) : bool`

##### Parameters

$name
: string

Name of the waiter.

##### Return values

bool

#### offsetExists()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.AbstractModel.html\#method_offsetExists)

`
    public
                    offsetExists(mixed $offset) : bool`

##### Parameters

$offset
: mixed

##### Return values

bool

#### offsetGet()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.AbstractModel.html\#method_offsetGet)

`
    public
                    offsetGet(mixed $offset) : mixed|null`

##### Parameters

$offset
: mixed

##### Return values

mixed\|null

#### offsetSet()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.AbstractModel.html\#method_offsetSet)

`
    public
                    offsetSet(mixed $offset, mixed $value) : void`

##### Parameters

$offset
: mixed$value
: mixed

#### offsetUnset()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.AbstractModel.html\#method_offsetUnset)

`
    public
                    offsetUnset(mixed $offset) : void`

##### Parameters

$offset
: mixed

#### toArray()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.AbstractModel.html\#method_toArray)

`
    public
                    toArray() : mixed`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.AbstractModel.html#method___construct)
  - [createErrorParser()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_createErrorParser)
  - [createParser()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_createParser)
  - [createSerializer()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_createSerializer)
  - [getApiVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getApiVersion)
  - [getClientContextParams()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getClientContextParams)
  - [getDefinition()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getDefinition)
  - [getEndpointPrefix()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getEndpointPrefix)
  - [getErrorShapes()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getErrorShapes)
  - [getMetadata()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getMetadata)
  - [getOperation()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getOperation)
  - [getOperations()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getOperations)
  - [getPaginatorConfig()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getPaginatorConfig)
  - [getPaginators()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getPaginators)
  - [getProtocol()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getProtocol)
  - [getProvider()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getProvider)
  - [getServiceFullName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getServiceFullName)
  - [getServiceId()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getServiceId)
  - [getServiceName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getServiceName)
  - [getShapeMap()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getShapeMap)
  - [getSignatureVersion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getSignatureVersion)
  - [getSigningName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getSigningName)
  - [getUid()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getUid)
  - [getWaiterConfig()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getWaiterConfig)
  - [getWaiters()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_getWaiters)
  - [hasOperation()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_hasOperation)
  - [hasPaginator()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_hasPaginator)
  - [hasWaiter()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#method_hasWaiter)
  - [offsetExists()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.AbstractModel.html#method_offsetExists)
  - [offsetGet()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.AbstractModel.html#method_offsetGet)
  - [offsetSet()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.AbstractModel.html#method_offsetSet)
  - [offsetUnset()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.AbstractModel.html#method_offsetUnset)
  - [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.AbstractModel.html#method_toArray)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html#top)
