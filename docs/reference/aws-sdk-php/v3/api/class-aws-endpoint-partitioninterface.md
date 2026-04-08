Menu

- [Aws](namespace-aws.md)
- [Endpoint](namespace-aws-endpoint.md)

## PartitionInterface     in    - [Aws](package-aws.md)

Represents a section of the AWS cloud.

### Table of Contents  [header link](class-aws-endpoint-partitioninterface-toc.md)

#### Methods  [header link](class-aws-endpoint-partitioninterface-toc-methods.md)

[\_\_invoke()](class-aws-endpoint-partitioninterface-method-invoke.md)
: array<string\|int, mixed> A partition must be invokable as an endpoint provider.[getAvailableEndpoints()](class-aws-endpoint-partitioninterface-method-getavailableendpoints.md)
: array<string\|int, string> Return the endpoints supported by a given service.[getName()](class-aws-endpoint-partitioninterface-method-getname.md)
: string Returns the partition's short name, e.g., 'aws,' 'aws-cn,' or
'aws-us-gov.'[isRegionMatch()](class-aws-endpoint-partitioninterface-method-isregionmatch.md)
: bool Determine if this partition contains the provided region. Include the
name of the service to inspect non-regional endpoints

### Methods  [header link](class-aws-endpoint-partitioninterface-methods.md)

#### \_\_invoke()  [header link](class-aws-endpoint-partitioninterface-method-invoke.md)

A partition must be invokable as an endpoint provider.

`
    public
                    __invoke([array<string|int, mixed> $args = [] ]) : array<string|int, mixed>`

##### Parameters

$args
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-endpoint-partitioninterface-method-invoke-tags.md)

see[EndpointProvider](class-aws-endpoint-endpointprovider.md)

##### Return values

array<string\|int, mixed>

#### getAvailableEndpoints()  [header link](class-aws-endpoint-partitioninterface-method-getavailableendpoints.md)

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

#### getName()  [header link](class-aws-endpoint-partitioninterface-method-getname.md)

Returns the partition's short name, e.g., 'aws,' 'aws-cn,' or
'aws-us-gov.'

`
    public
                    getName() : string`

##### Return values

string

#### isRegionMatch()  [header link](class-aws-endpoint-partitioninterface-method-isregionmatch.md)

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
  - [Constants](class-aws-endpoint-partitioninterface-toc-constants.md)
  - [Methods](class-aws-endpoint-partitioninterface-toc-methods.md)
- Methods
  - [\_\_invoke()](class-aws-endpoint-partitioninterface-method-invoke.md)
  - [getAvailableEndpoints()](class-aws-endpoint-partitioninterface-method-getavailableendpoints.md)
  - [getName()](class-aws-endpoint-partitioninterface-method-getname.md)
  - [isRegionMatch()](class-aws-endpoint-partitioninterface-method-isregionmatch.md)

[Back To Top](class-aws-endpoint-partitioninterface-top.md)
