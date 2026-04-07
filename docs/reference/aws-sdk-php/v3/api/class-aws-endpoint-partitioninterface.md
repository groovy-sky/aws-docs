Menu

- [Aws](namespace-aws.md)
- [Endpoint](namespace-aws-endpoint.md)

## PartitionInterface     in    - [Aws](package-aws.md)

Represents a section of the AWS cloud.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionInterface.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionInterface.html\#toc-methods)

[\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionInterface.html#method___invoke)
: array<string\|int, mixed> A partition must be invokable as an endpoint provider.[getAvailableEndpoints()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionInterface.html#method_getAvailableEndpoints)
: array<string\|int, string> Return the endpoints supported by a given service.[getName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionInterface.html#method_getName)
: string Returns the partition's short name, e.g., 'aws,' 'aws-cn,' or
'aws-us-gov.'[isRegionMatch()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionInterface.html#method_isRegionMatch)
: bool Determine if this partition contains the provided region. Include the
name of the service to inspect non-regional endpoints

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionInterface.html\#methods)

#### \_\_invoke()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionInterface.html\#method___invoke)

A partition must be invokable as an endpoint provider.

`
    public
                    __invoke([array<string|int, mixed> $args = [] ]) : array<string|int, mixed>`

##### Parameters

$args
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionInterface.html\#method___invoke\#tags)

see[EndpointProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.EndpointProvider.html)

##### Return values

array<string\|int, mixed>

#### getAvailableEndpoints()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionInterface.html\#method_getAvailableEndpoints)

Return the endpoints supported by a given service.

`
    public
                    getAvailableEndpoints(string $service[, bool $allowNonRegionalEndpoints = false ]) : array<string|int, string>`

##### Parameters

$service
: string

Identifier of the service
whose endpoints should be
listed (e.g., 's3' or 'ses')

$allowNonRegionalEndpoints
: bool
= false

Set to `true` to include
endpoints that are not AWS
regions (e.g., 'local' for
DynamoDB or
'fips-us-gov-west-1' for S3)

##### Return values

array<string\|int, string>

#### getName()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionInterface.html\#method_getName)

Returns the partition's short name, e.g., 'aws,' 'aws-cn,' or
'aws-us-gov.'

`
    public
                    getName() : string`

##### Return values

string

#### isRegionMatch()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionInterface.html\#method_isRegionMatch)

Determine if this partition contains the provided region. Include the
name of the service to inspect non-regional endpoints

`
    public
                    isRegionMatch(string $region, string $service) : bool`

##### Parameters

$region
: string$service
: string

##### Return values

bool
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionInterface.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionInterface.html#toc-methods)
- Methods
  - [\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionInterface.html#method___invoke)
  - [getAvailableEndpoints()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionInterface.html#method_getAvailableEndpoints)
  - [getName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionInterface.html#method_getName)
  - [isRegionMatch()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionInterface.html#method_isRegionMatch)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionInterface.html#top)
