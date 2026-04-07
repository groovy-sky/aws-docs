Menu

- [Aws](namespace-aws.md)
- [Endpoint](namespace-aws-endpoint.md)

## Partition        in package    - [Aws](package-aws.md)       implements  ArrayAccess, [PartitionInterface](class-aws-endpoint-partitioninterface.md)  Uses  [HasDataTrait](class-aws-hasdatatrait.md)

FinalYes

Default implementation of an AWS partition.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.Partition.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.Partition.html\#toc-interfaces)

ArrayAccess[PartitionInterface](class-aws-endpoint-partitioninterface.md)Represents a section of the AWS cloud.

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.Partition.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.Partition.html#method___construct)
: mixed The partition constructor accepts the following options:[\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.Partition.html#method___invoke)
: array<string\|int, mixed> A partition must be invokable as an endpoint provider.[count()](class-aws-hasdatatrait.md#method_count)
: int [getAvailableEndpoints()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.Partition.html#method_getAvailableEndpoints)
: array<string\|int, string> Return the endpoints supported by a given service.[getIterator()](class-aws-hasdatatrait.md#method_getIterator)
: Traversable[getName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.Partition.html#method_getName)
: string Returns the partition's short name, e.g., 'aws,' 'aws-cn,' or
'aws-us-gov.'[isRegionMatch()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.Partition.html#method_isRegionMatch)
: bool Determine if this partition contains the provided region. Include the
name of the service to inspect non-regional endpoints[offsetExists()](class-aws-hasdatatrait.md#method_offsetExists)
: bool [offsetGet()](class-aws-hasdatatrait.md#method_offsetGet)
: mixed\|null This method returns a reference to the variable to allow for indirect
array modification (e.g., $foo\['bar'\]\['baz'\] = 'qux').[offsetSet()](class-aws-hasdatatrait.md#method_offsetSet)
: void [offsetUnset()](class-aws-hasdatatrait.md#method_offsetUnset)
: void [toArray()](class-aws-hasdatatrait.md#method_toArray)
: mixed

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.Partition.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.Partition.html\#method___construct)

The partition constructor accepts the following options:

`
    public
                    __construct(array<string|int, mixed> $definition) : mixed`

- `partition`: (string, required) The partition name as specified in an
ARN (e.g., `aws`)
- `partitionName`: (string) The human readable name of the partition
(e.g., "AWS Standard")
- `dnsSuffix`: (string, required) The DNS suffix of the partition. This
value is used to determine how endpoints in the partition are resolved.
- `regionRegex`: (string) A PCRE regular expression that specifies the
pattern that region names in the endpoint adhere to.
- `regions`: (array, required) A map of the regions in the partition.
Each key is the region as present in a hostname (e.g., `us-east-1`),
and each value is a structure containing region information.
- `defaults`: (array) A map of default key value pairs to apply to each
endpoint of the partition. Any value in an `endpoint` definition will
supersede any values specified in `defaults`.
- `services`: (array, required) A map of service endpoint prefix name
(the value found in a hostname) to information about the service.

##### Parameters

$definition
: array<string\|int, mixed>

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.Partition.html\#method___construct\#tags)

throwsInvalidArgumentException

if any required options are missing

#### \_\_invoke()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.Partition.html\#method___invoke)

A partition must be invokable as an endpoint provider.

`
    public
                    __invoke([array<string|int, mixed> $args = [] ]) : array<string|int, mixed>`

##### Parameters

$args
: array<string\|int, mixed>
= \[\]

##### Return values

array<string\|int, mixed>

#### count()  [header link](class-aws-hasdatatrait.md\#method_count)

`
    public
                    count() : int`

##### Return values

int

#### getAvailableEndpoints()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.Partition.html\#method_getAvailableEndpoints)

Return the endpoints supported by a given service.

`
    public
                    getAvailableEndpoints(mixed $service[, mixed $allowNonRegionalEndpoints = false ]) : array<string|int, string>`

##### Parameters

$service
: mixed

Identifier of the service
whose endpoints should be
listed (e.g., 's3' or 'ses')

$allowNonRegionalEndpoints
: mixed
= false

Set to `true` to include
endpoints that are not AWS
regions (e.g., 'local' for
DynamoDB or
'fips-us-gov-west-1' for S3)

##### Return values

array<string\|int, string>

#### getIterator()  [header link](class-aws-hasdatatrait.md\#method_getIterator)

`
    public
                    getIterator() : Traversable`

##### Return values

Traversable

#### getName()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.Partition.html\#method_getName)

Returns the partition's short name, e.g., 'aws,' 'aws-cn,' or
'aws-us-gov.'

`
    public
                    getName() : string`

##### Return values

string

#### isRegionMatch()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.Partition.html\#method_isRegionMatch)

Determine if this partition contains the provided region. Include the
name of the service to inspect non-regional endpoints

`
    public
                    isRegionMatch(mixed $region, mixed $service) : bool`

##### Parameters

$region
: mixed$service
: mixed

##### Return values

bool

#### offsetExists()  [header link](class-aws-hasdatatrait.md\#method_offsetExists)

`
    public
                    offsetExists(mixed $offset) : bool`

##### Parameters

$offset
: mixed

##### Return values

bool

#### offsetGet()  [header link](class-aws-hasdatatrait.md\#method_offsetGet)

This method returns a reference to the variable to allow for indirect
array modification (e.g., $foo\['bar'\]\['baz'\] = 'qux').

`
    public
                &    offsetGet( $offset) : mixed|null`

##### Parameters

$offset
:

##### Return values

mixed\|null

#### offsetSet()  [header link](class-aws-hasdatatrait.md\#method_offsetSet)

`
    public
                    offsetSet(mixed $offset, mixed $value) : void`

##### Parameters

$offset
: mixed$value
: mixed

#### offsetUnset()  [header link](class-aws-hasdatatrait.md\#method_offsetUnset)

`
    public
                    offsetUnset(mixed $offset) : void`

##### Parameters

$offset
: mixed

#### toArray()  [header link](class-aws-hasdatatrait.md\#method_toArray)

`
    public
                    toArray() : mixed`

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.Partition.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.Partition.html#method___construct)
  - [\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.Partition.html#method___invoke)
  - [count()](class-aws-hasdatatrait.md#method_count)
  - [getAvailableEndpoints()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.Partition.html#method_getAvailableEndpoints)
  - [getIterator()](class-aws-hasdatatrait.md#method_getIterator)
  - [getName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.Partition.html#method_getName)
  - [isRegionMatch()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.Partition.html#method_isRegionMatch)
  - [offsetExists()](class-aws-hasdatatrait.md#method_offsetExists)
  - [offsetGet()](class-aws-hasdatatrait.md#method_offsetGet)
  - [offsetSet()](class-aws-hasdatatrait.md#method_offsetSet)
  - [offsetUnset()](class-aws-hasdatatrait.md#method_offsetUnset)
  - [toArray()](class-aws-hasdatatrait.md#method_toArray)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.Partition.html#top)
