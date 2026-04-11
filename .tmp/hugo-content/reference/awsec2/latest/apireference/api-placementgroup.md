# PlacementGroup

Describes a placement group.

## Contents

**groupArn**

The Amazon Resource Name (ARN) of the placement group.

Type: String

Required: No

**groupId**

The ID of the placement group.

Type: String

Required: No

**groupName**

The name of the placement group.

Type: String

Required: No

**linkedGroupId**

Reserved for future use.

Type: String

Required: No

**operator**

The service provider that manages the Placement Group.

Type: [OperatorResponse](api-operatorresponse.md) object

Required: No

**partitionCount**

The number of partitions. Valid only if **strategy** is
set to `partition`.

Type: Integer

Required: No

**spreadLevel**

The spread level for the placement group. _Only_ Outpost placement
groups can be spread across hosts.

Type: String

Valid Values: `host | rack`

Required: No

**state**

The state of the placement group.

Type: String

Valid Values: `pending | available | deleting | deleted`

Required: No

**strategy**

The placement strategy.

Type: String

Valid Values: `cluster | spread | partition`

Required: No

**TagSet.N**

Any tags applied to the placement group.

Type: Array of [Tag](api-tag.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/placementgroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/placementgroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/placementgroup.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Placement

PlacementGroupInfo
