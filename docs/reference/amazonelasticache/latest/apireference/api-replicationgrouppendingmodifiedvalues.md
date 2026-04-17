---
title: "ReplicationGroupPendingModifiedValues"
---

# ReplicationGroupPendingModifiedValues

The settings to be applied to the Valkey or Redis OSS replication group, either immediately or
during the next maintenance window.

## Contents

###### Note

In the following list, the required parameters are described first.

**AuthTokenStatus**

The auth token status

Type: String

Valid Values: `SETTING | ROTATING`

Required: No

**AutomaticFailoverStatus**

Indicates the status of automatic failover for this Valkey or Redis OSS replication group.

Type: String

Valid Values: `enabled | disabled`

Required: No

**ClusterMode**

Enabled or Disabled. To modify cluster mode from Disabled to Enabled, you must first
set the cluster mode to Compatible. Compatible mode allows your Valkey or Redis OSS clients to connect
using both cluster mode enabled and cluster mode disabled. After you migrate all Valkey or Redis OSS
clients to use cluster mode enabled, you can then complete cluster mode configuration
and set the cluster mode to Enabled.

Type: String

Valid Values: `enabled | disabled | compatible`

Required: No

**PendingLogDeliveryConfiguration.member.N**

The log delivery configurations being modified

Type: Array of [PendingLogDeliveryConfiguration](api-pendinglogdeliveryconfiguration.md) objects

Required: No

**PrimaryClusterId**

The primary cluster ID that is applied immediately (if
`--apply-immediately` was specified), or during the next maintenance
window.

Type: String

Required: No

**Resharding**

The status of an online resharding operation.

Type: [ReshardingStatus](api-reshardingstatus.md) object

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

**UserGroups**

The user group being modified.

Type: [UserGroupsUpdateStatus](api-usergroupsupdatestatus.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/elasticache-2015-02-02/ReplicationGroupPendingModifiedValues)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/elasticache-2015-02-02/ReplicationGroupPendingModifiedValues)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/elasticache-2015-02-02/ReplicationGroupPendingModifiedValues)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicationGroup

ReservedCacheNode

All content copied from https://docs.aws.amazon.com/.
