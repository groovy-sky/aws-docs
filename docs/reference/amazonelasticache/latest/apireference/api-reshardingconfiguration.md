# ReshardingConfiguration

A list of `PreferredAvailabilityZones` objects that specifies the
configuration of a node group in the resharded cluster.

## Contents

###### Note

In the following list, the required parameters are described first.

**NodeGroupId**

Either the ElastiCache supplied 4-digit id or a user supplied id for the
node group these configuration values apply to.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 4.

Pattern: `\d+`

Required: No

**PreferredAvailabilityZones.AvailabilityZone.N**

A list of preferred availability zones for the nodes in this cluster.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/reshardingconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/reshardingconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/reshardingconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReservedCacheNodesOffering

ReshardingStatus

All content copied from https://docs.aws.amazon.com/.
