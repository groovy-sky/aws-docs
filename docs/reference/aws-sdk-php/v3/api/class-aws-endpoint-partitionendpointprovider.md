Menu

- [Aws](namespace-aws.md)
- [Endpoint](namespace-aws-endpoint.md)

## PartitionEndpointProvider        in package    - [Aws](package-aws.md)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionEndpointProvider.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionEndpointProvider.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionEndpointProvider.html#method___construct)
: mixed The 'options' parameter accepts the following arguments:[\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionEndpointProvider.html#method___invoke)
: mixed [defaultProvider()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionEndpointProvider.html#method_defaultProvider)
: [PartitionEndpointProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionEndpointProvider.html)Creates and returns the default SDK partition provider.[getPartition()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionEndpointProvider.html#method_getPartition)
: [Partition](class-aws-endpoint-partition.md)Returns the partition containing the provided region or the default
partition if no match is found.[getPartitionByName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionEndpointProvider.html#method_getPartitionByName)
: [Partition](class-aws-endpoint-partition.md) \|null Returns the partition with the provided name or null if no partition with
the provided name can be found.[mergePrefixData()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionEndpointProvider.html#method_mergePrefixData)
: array<string\|int, mixed> Copy endpoint data for other prefixes used by a given service

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionEndpointProvider.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionEndpointProvider.html\#method___construct)

The 'options' parameter accepts the following arguments:

`
    public
                    __construct(array<string|int, mixed> $partitions[, string $defaultPartition = 'aws' ][, array<string|int, mixed> $options = [] ]) : mixed`

- sts\_regional\_endpoints: For STS legacy regions, set to 'regional' to
use regional endpoints, 'legacy' to use the legacy global endpoint.
Defaults to 'legacy'.
- s3\_us\_east\_1\_regional\_endpoint: For S3 us-east-1 region, set to 'regional'
to use the regional endpoint, 'legacy' to use the legacy global endpoint.
Defaults to 'legacy'.

##### Parameters

$partitions
: array<string\|int, mixed>$defaultPartition
: string
= 'aws'$options
: array<string\|int, mixed>
= \[\]

#### \_\_invoke()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionEndpointProvider.html\#method___invoke)

`
    public
                    __invoke([array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$args
: array<string\|int, mixed>
= \[\]

#### defaultProvider()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionEndpointProvider.html\#method_defaultProvider)

Creates and returns the default SDK partition provider.

`
    public
            static        defaultProvider([array<string|int, mixed> $options = [] ]) : PartitionEndpointProvider`

##### Parameters

$options
: array<string\|int, mixed>
= \[\]

##### Return values

[PartitionEndpointProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionEndpointProvider.html)

#### getPartition()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionEndpointProvider.html\#method_getPartition)

Returns the partition containing the provided region or the default
partition if no match is found.

`
    public
                    getPartition(string $region, string $service) : Partition`

##### Parameters

$region
: string$service
: string

##### Return values

[Partition](class-aws-endpoint-partition.md)

#### getPartitionByName()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionEndpointProvider.html\#method_getPartitionByName)

Returns the partition with the provided name or null if no partition with
the provided name can be found.

`
    public
                    getPartitionByName(string $name) : Partition|null`

##### Parameters

$name
: string

##### Return values

[Partition](class-aws-endpoint-partition.md) \|null

#### mergePrefixData()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionEndpointProvider.html\#method_mergePrefixData)

Copy endpoint data for other prefixes used by a given service

`
    public
            static        mergePrefixData( $data,  $prefixData) : array<string|int, mixed>`

##### Parameters

$data
: $prefixData
:

##### Return values

array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionEndpointProvider.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionEndpointProvider.html#method___construct)
  - [\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionEndpointProvider.html#method___invoke)
  - [defaultProvider()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionEndpointProvider.html#method_defaultProvider)
  - [getPartition()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionEndpointProvider.html#method_getPartition)
  - [getPartitionByName()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionEndpointProvider.html#method_getPartitionByName)
  - [mergePrefixData()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionEndpointProvider.html#method_mergePrefixData)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.PartitionEndpointProvider.html#top)
