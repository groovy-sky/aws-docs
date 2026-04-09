# ProcessedUpdateAction

Update action that has been processed for the corresponding apply/stop request

## Contents

###### Note

In the following list, the required parameters are described first.

**CacheClusterId**

The ID of the cache cluster

Type: String

Required: No

**ReplicationGroupId**

The ID of the replication group

Type: String

Required: No

**ServiceUpdateName**

The unique ID of the service update

Type: String

Required: No

**UpdateActionStatus**

The status of the update action on the Valkey or Redis OSS cluster

Type: String

Valid Values: `not-applied | waiting-to-start | in-progress | stopping | stopped | complete | scheduling | scheduled | not-applicable`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/processedupdateaction.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/processedupdateaction.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/processedupdateaction.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PendingModifiedValues

RecurringCharge

All content copied from https://docs.aws.amazon.com/.
