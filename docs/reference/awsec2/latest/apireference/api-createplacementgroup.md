# CreatePlacementGroup

Creates a placement group in which to launch instances. The strategy of the placement
group determines how the instances are organized within the group.

A `cluster` placement group is a logical grouping of instances within a
single Availability Zone that benefit from low network latency, high network throughput.
A `spread` placement group places instances on distinct hardware. A
`partition` placement group places groups of instances in different
partitions, where instances in one partition do not share the same hardware with
instances in another partition.

For more information, see [Placement groups](../../../../services/ec2/latest/userguide/placement-groups.md) in the
_Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the operation, without actually making the
request, and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**GroupName**

A name for the placement group. Must be unique within the scope of your account for
the Region.

Constraints: Up to 255 ASCII characters

Type: String

Required: No

**LinkedGroupId**

Reserved for future use.

Type: String

Required: No

**Operator**

Reserved for internal use.

Type: [OperatorRequest](api-operatorrequest.md) object

Required: No

**PartitionCount**

The number of partitions. Valid only when **Strategy** is
set to `partition`.

Type: Integer

Required: No

**SpreadLevel**

Determines how placement groups spread instances.

- Host – You can use `host` only with Outpost placement
groups.

- Rack – No usage restrictions.

Type: String

Valid Values: `host | rack`

Required: No

**Strategy**

The placement strategy.

Type: String

Valid Values: `cluster | spread | partition`

Required: No

**TagSpecification.N**

The tags to apply to the new placement group.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**placementGroup**

Information about the placement group.

Type: [PlacementGroup](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_PlacementGroup.html) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example creates a cluster placement group named `XYZ-cluster`,
and applies a tag with a key of `purpose` and a value of
`production`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CreatePlacementGroup
&GroupName=XYZ-cluster
&Strategy=cluster
&TagSpecification.1.ResourceType=placement-group
&TagSpecification.1.Tag.1.Key=purpose
&TagSpecification.1.Tag.1.Value=production
&AUTHPARAMS
```

#### Sample Response

```

<CreatePlacementGroupResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>1bbcaf48-7155-4154-a7ac-c6031EXAMPLE</requestId>
    <return>true</return>
    <placementGroup>
        <groupName>XYZ-cluster</groupName>
        <groupId>pg-0bea00ad0bexample</groupId>
        <strategy>cluster</strategy>
        <state>available</state>
        <tagSet>
            <item>
                <key>purpose</key>
                <value>production</value>
            </item>
        </tagSet>
    </placementGroup>
</CreatePlacementGroupResponse>
```

### Example

This example creates a partition placement group named
`HDFS-Group-A` with five partitions.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CreatePlacementGroup
&GroupName=HDFS-Group-A
&Strategy=partition
&PartitionCount=5
&AUTHPARAMS
```

#### Sample Response

```

<CreatePlacementGroupResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>1bbcaf48-7155-4154-a7ac-c6031EXAMPLE</requestId>
    <return>true</return>
    <placementGroup>
        <groupName>HDFS-Group-A</groupName>
        <groupId>pg-0fc13f6eb3example</groupId>
        <strategy>partition</strategy>
        <state>available</state>
        <partitionCount>5</partitionCount>
    </placementGroup>
</CreatePlacementGroupResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreatePlacementGroup)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreatePlacementGroup)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/CreatePlacementGroup)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/CreatePlacementGroup)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/CreatePlacementGroup)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/CreatePlacementGroup)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/CreatePlacementGroup)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/CreatePlacementGroup)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreatePlacementGroup)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/CreatePlacementGroup)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CreateNetworkInterfacePermission

CreatePublicIpv4Pool
