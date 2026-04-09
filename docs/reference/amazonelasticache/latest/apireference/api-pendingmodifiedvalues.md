# PendingModifiedValues

A group of settings that are applied to the cluster in the future, or that are
currently being applied.

## Contents

###### Note

In the following list, the required parameters are described first.

**AuthTokenStatus**

The auth token status

Type: String

Valid Values: `SETTING | ROTATING`

Required: No

**CacheNodeIdsToRemove.CacheNodeId.N**

A list of cache node IDs that are being removed (or will be removed) from the cluster.
A node ID is a 4-digit numeric identifier (0001, 0002, etc.).

Type: Array of strings

Required: No

**CacheNodeType**

The cache node type that this cluster or replication group is scaled to.

Type: String

Required: No

**EngineVersion**

The new cache engine version that the cluster runs.

Type: String

Required: No

**NumCacheNodes**

The new number of cache nodes for the cluster.

For clusters running Valkey or Redis OSS, this value must be 1. For clusters running Memcached, this
value must be between 1 and 40.

Type: Integer

Required: No

**PendingLogDeliveryConfiguration.member.N**

The log delivery configurations being modified

Type: Array of [PendingLogDeliveryConfiguration](api-pendinglogdeliveryconfiguration.md) objects

Required: No

**ScaleConfig**

The scaling configuration changes that are pending for the Memcached cluster.

Type: [ScaleConfig](api-scaleconfig.md) object

Required: No

**TransitEncryptionEnabled**

A flag that enables in-transit encryption when set to true.

Type: Boolean

Required: No

**TransitEncryptionMode**

A setting that allows you to migrate your clients to use in-transit encryption, with
no downtime.

Type: String

Valid Values: `preferred | required`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/pendingmodifiedvalues.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/pendingmodifiedvalues.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/pendingmodifiedvalues.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PendingLogDeliveryConfiguration

ProcessedUpdateAction

All content copied from https://docs.aws.amazon.com/.
