Menu

- [Aws](namespace-aws.md)
- [Api](namespace-aws-api.md)

## Service     extends [AbstractModel](class-aws-api-abstractmodel.md)   in package    - [Aws](package-aws.md)

Represents a web service API model.

### Table of Contents  [header link](class-aws-api-service-toc.md)

#### Methods  [header link](class-aws-api-service-toc-methods.md)

[\_\_construct()](class-aws-api-abstractmodel-method-construct.md)
: mixed [createErrorParser()](class-aws-api-service-method-createerrorparser.md)
: callable Creates an error parser for the given protocol.[createParser()](class-aws-api-service-method-createparser.md)
: callable Applies the listeners needed to parse client models.[createSerializer()](class-aws-api-service-method-createserializer.md)
: callable Creates a request serializer for the provided API object.[getApiVersion()](class-aws-api-service-method-getapiversion.md)
: string Get the API version of the service[getClientContextParams()](class-aws-api-service-method-getclientcontextparams.md)
: array<string\|int, mixed> Get all the context params of the description.[getDefinition()](class-aws-api-service-method-getdefinition.md)
: callable Get the service's definition.[getEndpointPrefix()](class-aws-api-service-method-getendpointprefix.md)
: string Get the API version of the service[getErrorShapes()](class-aws-api-service-method-geterrorshapes.md)
: array<string\|int, mixed> Get all of the error shapes of the service[getMetadata()](class-aws-api-service-method-getmetadata.md)
: mixed Get all of the service metadata or a specific metadata key value.[getOperation()](class-aws-api-service-method-getoperation.md)
: [Operation](class-aws-api-operation.md)Get an operation by name.[getOperations()](class-aws-api-service-method-getoperations.md)
: array<string\|int, [Operation](class-aws-api-operation.md) \> Get all of the operations of the description.[getPaginatorConfig()](class-aws-api-service-method-getpaginatorconfig.md)
: array<string\|int, mixed> Retrieve a paginator by name.[getPaginators()](class-aws-api-service-method-getpaginators.md)
: array<string\|int, mixed> Gets an associative array of available paginator configurations where
the key is the name of the paginator, and the value is the paginator
configuration.[getProtocol()](class-aws-api-service-method-getprotocol.md)
: string Get the protocol used by the service.[getProvider()](class-aws-api-service-method-getprovider.md)
: callable Get the service's api provider.[getServiceFullName()](class-aws-api-service-method-getservicefullname.md)
: string Get the full name of the service[getServiceId()](class-aws-api-service-method-getserviceid.md)
: string Get the service id[getServiceName()](class-aws-api-service-method-getservicename.md)
: string Get the service name.[getShapeMap()](class-aws-api-service-method-getshapemap.md)
: [ShapeMap](class-aws-api-shapemap.md)Get the shape map used by the API.[getSignatureVersion()](class-aws-api-service-method-getsignatureversion.md)
: string Get the default signature version of the service.[getSigningName()](class-aws-api-service-method-getsigningname.md)
: string Get the signing name used by the service.[getUid()](class-aws-api-service-method-getuid.md)
: string Get the uid string used by the service[getWaiterConfig()](class-aws-api-service-method-getwaiterconfig.md)
: array<string\|int, mixed> Get a waiter configuration by name.[getWaiters()](class-aws-api-service-method-getwaiters.md)
: array<string\|int, mixed> Gets an associative array of available waiter configurations where the
key is the name of the waiter, and the value is the waiter
configuration.[hasOperation()](class-aws-api-service-method-hasoperation.md)
: bool Check if the description has a specific operation by name.[hasPaginator()](class-aws-api-service-method-haspaginator.md)
: bool Determines if the service has a paginator by name.[hasWaiter()](class-aws-api-service-method-haswaiter.md)
: bool Determines if the service has a waiter by name.[offsetExists()](class-aws-api-abstractmodel-method-offsetexists.md)
: bool [offsetGet()](class-aws-api-abstractmodel-method-offsetget.md)
: mixed\|null [offsetSet()](class-aws-api-abstractmodel-method-offsetset.md)
: void [offsetUnset()](class-aws-api-abstractmodel-method-offsetunset.md)
: void [toArray()](class-aws-api-abstractmodel-method-toarray.md)
: mixed

### Methods  [header link](class-aws-api-service-methods.md)

#### \_\_construct()  [header link](class-aws-api-abstractmodel-method-construct.md)

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

#### createErrorParser()  [header link](class-aws-api-service-method-createerrorparser.md)

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
: [Service](class-aws-api-service.md) \|null
= null

##### Tags  [header link](class-aws-api-service-method-createerrorparser-tags.md)

throwsUnexpectedValueException

##### Return values

callable

#### createParser()  [header link](class-aws-api-service-method-createparser.md)

Applies the listeners needed to parse client models.

`
    public
            static        createParser(Service $api) : callable`

##### Parameters

$api
: [Service](class-aws-api-service.md)

API to create a parser for

##### Tags  [header link](class-aws-api-service-method-createparser-tags.md)

throwsUnexpectedValueException

##### Return values

callable

#### createSerializer()  [header link](class-aws-api-service-method-createserializer.md)

Creates a request serializer for the provided API object.

`
    public
            static        createSerializer(Service $api, string $endpoint) : callable`

##### Parameters

$api
: [Service](class-aws-api-service.md)

API that contains a protocol.

$endpoint
: string

Endpoint to send requests to.

##### Tags  [header link](class-aws-api-service-method-createserializer-tags.md)

throwsUnexpectedValueException

##### Return values

callable

#### getApiVersion()  [header link](class-aws-api-service-method-getapiversion.md)

Get the API version of the service

`
    public
                    getApiVersion() : string`

##### Return values

string

#### getClientContextParams()  [header link](class-aws-api-service-method-getclientcontextparams.md)

Get all the context params of the description.

`
    public
                    getClientContextParams() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getDefinition()  [header link](class-aws-api-service-method-getdefinition.md)

Get the service's definition.

`
    public
                    getDefinition() : callable`

##### Return values

callable

#### getEndpointPrefix()  [header link](class-aws-api-service-method-getendpointprefix.md)

Get the API version of the service

`
    public
                    getEndpointPrefix() : string`

##### Return values

string

#### getErrorShapes()  [header link](class-aws-api-service-method-geterrorshapes.md)

Get all of the error shapes of the service

`
    public
                    getErrorShapes() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getMetadata()  [header link](class-aws-api-service-method-getmetadata.md)

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

#### getOperation()  [header link](class-aws-api-service-method-getoperation.md)

Get an operation by name.

`
    public
                    getOperation(string $name) : Operation`

##### Parameters

$name
: string

Operation to retrieve by name

##### Tags  [header link](class-aws-api-service-method-getoperation-tags.md)

throwsInvalidArgumentException

If the operation is not found

##### Return values

[Operation](class-aws-api-operation.md)

#### getOperations()  [header link](class-aws-api-service-method-getoperations.md)

Get all of the operations of the description.

`
    public
                    getOperations() : array<string|int, Operation>`

##### Return values

array<string\|int, [Operation](class-aws-api-operation.md) >

#### getPaginatorConfig()  [header link](class-aws-api-service-method-getpaginatorconfig.md)

Retrieve a paginator by name.

`
    public
                    getPaginatorConfig(string $name) : array<string|int, mixed>`

##### Parameters

$name
: string

Paginator to retrieve by name. This argument is
typically the operation name.

##### Tags  [header link](class-aws-api-service-method-getpaginatorconfig-tags.md)

throwsUnexpectedValueException

if the paginator does not exist.

unstable

The configuration format of paginators may change in the future

##### Return values

array<string\|int, mixed>

#### getPaginators()  [header link](class-aws-api-service-method-getpaginators.md)

Gets an associative array of available paginator configurations where
the key is the name of the paginator, and the value is the paginator
configuration.

`
    public
                    getPaginators() : array<string|int, mixed>`

##### Tags  [header link](class-aws-api-service-method-getpaginators-tags.md)

unstable

The configuration format of paginators may change in the future

##### Return values

array<string\|int, mixed>

#### getProtocol()  [header link](class-aws-api-service-method-getprotocol.md)

Get the protocol used by the service.

`
    public
                    getProtocol() : string`

##### Return values

string

#### getProvider()  [header link](class-aws-api-service-method-getprovider.md)

Get the service's api provider.

`
    public
                    getProvider() : callable`

##### Return values

callable

#### getServiceFullName()  [header link](class-aws-api-service-method-getservicefullname.md)

Get the full name of the service

`
    public
                    getServiceFullName() : string`

##### Return values

string

#### getServiceId()  [header link](class-aws-api-service-method-getserviceid.md)

Get the service id

`
    public
                    getServiceId() : string`

##### Return values

string

#### getServiceName()  [header link](class-aws-api-service-method-getservicename.md)

Get the service name.

`
    public
                    getServiceName() : string`

##### Return values

string

#### getShapeMap()  [header link](class-aws-api-service-method-getshapemap.md)

Get the shape map used by the API.

`
    public
                    getShapeMap() : ShapeMap`

##### Return values

[ShapeMap](class-aws-api-shapemap.md)

#### getSignatureVersion()  [header link](class-aws-api-service-method-getsignatureversion.md)

Get the default signature version of the service.

`
    public
                    getSignatureVersion() : string`

Note: this method assumes "v4" when not specified in the model.

##### Return values

string

#### getSigningName()  [header link](class-aws-api-service-method-getsigningname.md)

Get the signing name used by the service.

`
    public
                    getSigningName() : string`

##### Return values

string

#### getUid()  [header link](class-aws-api-service-method-getuid.md)

Get the uid string used by the service

`
    public
                    getUid() : string`

##### Return values

string

#### getWaiterConfig()  [header link](class-aws-api-service-method-getwaiterconfig.md)

Get a waiter configuration by name.

`
    public
                    getWaiterConfig(string $name) : array<string|int, mixed>`

##### Parameters

$name
: string

Name of the waiter by name.

##### Tags  [header link](class-aws-api-service-method-getwaiterconfig-tags.md)

throwsUnexpectedValueException

if the waiter does not exist.

##### Return values

array<string\|int, mixed>

#### getWaiters()  [header link](class-aws-api-service-method-getwaiters.md)

Gets an associative array of available waiter configurations where the
key is the name of the waiter, and the value is the waiter
configuration.

`
    public
                    getWaiters() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### hasOperation()  [header link](class-aws-api-service-method-hasoperation.md)

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

#### hasPaginator()  [header link](class-aws-api-service-method-haspaginator.md)

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

#### hasWaiter()  [header link](class-aws-api-service-method-haswaiter.md)

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

#### offsetExists()  [header link](class-aws-api-abstractmodel-method-offsetexists.md)

`
    public
                    offsetExists(mixed $offset) : bool`

##### Parameters

$offset
: mixed

##### Return values

bool

#### offsetGet()  [header link](class-aws-api-abstractmodel-method-offsetget.md)

`
    public
                    offsetGet(mixed $offset) : mixed|null`

##### Parameters

$offset
: mixed

##### Return values

mixed\|null

#### offsetSet()  [header link](class-aws-api-abstractmodel-method-offsetset.md)

`
    public
                    offsetSet(mixed $offset, mixed $value) : void`

##### Parameters

$offset
: mixed$value
: mixed

#### offsetUnset()  [header link](class-aws-api-abstractmodel-method-offsetunset.md)

`
    public
                    offsetUnset(mixed $offset) : void`

##### Parameters

$offset
: mixed

#### toArray()  [header link](class-aws-api-abstractmodel-method-toarray.md)

`
    public
                    toArray() : mixed`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-api-service-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-api-abstractmodel-method-construct.md)
  - [createErrorParser()](class-aws-api-service-method-createerrorparser.md)
  - [createParser()](class-aws-api-service-method-createparser.md)
  - [createSerializer()](class-aws-api-service-method-createserializer.md)
  - [getApiVersion()](class-aws-api-service-method-getapiversion.md)
  - [getClientContextParams()](class-aws-api-service-method-getclientcontextparams.md)
  - [getDefinition()](class-aws-api-service-method-getdefinition.md)
  - [getEndpointPrefix()](class-aws-api-service-method-getendpointprefix.md)
  - [getErrorShapes()](class-aws-api-service-method-geterrorshapes.md)
  - [getMetadata()](class-aws-api-service-method-getmetadata.md)
  - [getOperation()](class-aws-api-service-method-getoperation.md)
  - [getOperations()](class-aws-api-service-method-getoperations.md)
  - [getPaginatorConfig()](class-aws-api-service-method-getpaginatorconfig.md)
  - [getPaginators()](class-aws-api-service-method-getpaginators.md)
  - [getProtocol()](class-aws-api-service-method-getprotocol.md)
  - [getProvider()](class-aws-api-service-method-getprovider.md)
  - [getServiceFullName()](class-aws-api-service-method-getservicefullname.md)
  - [getServiceId()](class-aws-api-service-method-getserviceid.md)
  - [getServiceName()](class-aws-api-service-method-getservicename.md)
  - [getShapeMap()](class-aws-api-service-method-getshapemap.md)
  - [getSignatureVersion()](class-aws-api-service-method-getsignatureversion.md)
  - [getSigningName()](class-aws-api-service-method-getsigningname.md)
  - [getUid()](class-aws-api-service-method-getuid.md)
  - [getWaiterConfig()](class-aws-api-service-method-getwaiterconfig.md)
  - [getWaiters()](class-aws-api-service-method-getwaiters.md)
  - [hasOperation()](class-aws-api-service-method-hasoperation.md)
  - [hasPaginator()](class-aws-api-service-method-haspaginator.md)
  - [hasWaiter()](class-aws-api-service-method-haswaiter.md)
  - [offsetExists()](class-aws-api-abstractmodel-method-offsetexists.md)
  - [offsetGet()](class-aws-api-abstractmodel-method-offsetget.md)
  - [offsetSet()](class-aws-api-abstractmodel-method-offsetset.md)
  - [offsetUnset()](class-aws-api-abstractmodel-method-offsetunset.md)
  - [toArray()](class-aws-api-abstractmodel-method-toarray.md)

[Back To Top](class-aws-api-service-top.md)
