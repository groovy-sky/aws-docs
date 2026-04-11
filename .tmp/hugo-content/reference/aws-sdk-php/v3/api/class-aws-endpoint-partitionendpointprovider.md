Menu

- [Aws](namespace-aws.md)
- [Endpoint](namespace-aws-endpoint.md)

## PartitionEndpointProvider        in package    - [Aws](package-aws.md)

### Table of Contents  [header link](class-aws-endpoint-partitionendpointprovider-toc.md)

#### Methods  [header link](class-aws-endpoint-partitionendpointprovider-toc-methods.md)

[\_\_construct()](class-aws-endpoint-partitionendpointprovider-method-construct.md)
: mixed The 'options' parameter accepts the following arguments:[\_\_invoke()](class-aws-endpoint-partitionendpointprovider-method-invoke.md)
: mixed [defaultProvider()](class-aws-endpoint-partitionendpointprovider-method-defaultprovider.md)
: [PartitionEndpointProvider](class-aws-endpoint-partitionendpointprovider.md)Creates and returns the default SDK partition provider.[getPartition()](class-aws-endpoint-partitionendpointprovider-method-getpartition.md)
: [Partition](class-aws-endpoint-partition.md)Returns the partition containing the provided region or the default
partition if no match is found.[getPartitionByName()](class-aws-endpoint-partitionendpointprovider-method-getpartitionbyname.md)
: [Partition](class-aws-endpoint-partition.md) \|null Returns the partition with the provided name or null if no partition with
the provided name can be found.[mergePrefixData()](class-aws-endpoint-partitionendpointprovider-method-mergeprefixdata.md)
: array<string\|int, mixed> Copy endpoint data for other prefixes used by a given service

### Methods  [header link](class-aws-endpoint-partitionendpointprovider-methods.md)

#### \_\_construct()  [header link](class-aws-endpoint-partitionendpointprovider-method-construct.md)

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

#### \_\_invoke()  [header link](class-aws-endpoint-partitionendpointprovider-method-invoke.md)

`
    public
                    __invoke([array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$args
: array<string\|int, mixed>
= \[\]

#### defaultProvider()  [header link](class-aws-endpoint-partitionendpointprovider-method-defaultprovider.md)

Creates and returns the default SDK partition provider.

`
    public
            static        defaultProvider([array<string|int, mixed> $options = [] ]) : PartitionEndpointProvider`

##### Parameters

$options
: array<string\|int, mixed>
= \[\]

##### Return values

[PartitionEndpointProvider](class-aws-endpoint-partitionendpointprovider.md)

#### getPartition()  [header link](class-aws-endpoint-partitionendpointprovider-method-getpartition.md)

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

#### getPartitionByName()  [header link](class-aws-endpoint-partitionendpointprovider-method-getpartitionbyname.md)

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

#### mergePrefixData()  [header link](class-aws-endpoint-partitionendpointprovider-method-mergeprefixdata.md)

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
  - [Methods](class-aws-endpoint-partitionendpointprovider-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-endpoint-partitionendpointprovider-method-construct.md)
  - [\_\_invoke()](class-aws-endpoint-partitionendpointprovider-method-invoke.md)
  - [defaultProvider()](class-aws-endpoint-partitionendpointprovider-method-defaultprovider.md)
  - [getPartition()](class-aws-endpoint-partitionendpointprovider-method-getpartition.md)
  - [getPartitionByName()](class-aws-endpoint-partitionendpointprovider-method-getpartitionbyname.md)
  - [mergePrefixData()](class-aws-endpoint-partitionendpointprovider-method-mergeprefixdata.md)

[Back To Top](class-aws-endpoint-partitionendpointprovider-top.md)
