# RegionalConfiguration

A list of the replication groups

## Contents

###### Note

In the following list, the required parameters are described first.

**ReplicationGroupId**

The name of the secondary cluster

Type: String

Required: Yes

**ReplicationGroupRegion**

The Amazon region where the cluster is stored

Type: String

Required: Yes

**ReshardingConfiguration.ReshardingConfiguration.N**

A list of `PreferredAvailabilityZones` objects that specifies the
configuration of a node group in the resharded cluster.

Type: Array of [ReshardingConfiguration](api-reshardingconfiguration.md) objects

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/regionalconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/regionalconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/regionalconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RecurringCharge

ReplicationGroup

All content copied from https://docs.aws.amazon.com/.
