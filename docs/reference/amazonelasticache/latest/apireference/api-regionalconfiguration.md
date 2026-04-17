---
title: "RegionalConfiguration"
---

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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/elasticache-2015-02-02/RegionalConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/elasticache-2015-02-02/RegionalConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/elasticache-2015-02-02/RegionalConfiguration)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RecurringCharge

ReplicationGroup

All content copied from https://docs.aws.amazon.com/.
