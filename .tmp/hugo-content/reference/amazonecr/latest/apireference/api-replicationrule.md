# ReplicationRule

An array of objects representing the replication destinations and repository filters
for a replication configuration.

## Contents

**destinations**

An array of objects representing the destination for a replication rule.

Type: Array of [ReplicationDestination](api-replicationdestination.md) objects

Array Members: Minimum number of 0 items. Maximum number of 100 items.

Required: Yes

**repositoryFilters**

An array of objects representing the filters for a replication rule. Specifying a
repository filter for a replication rule provides a method for controlling which
repositories in a private registry are replicated.

Type: Array of [RepositoryFilter](api-repositoryfilter.md) objects

Array Members: Minimum number of 1 item. Maximum number of 100 items.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ecr-2015-09-21/replicationrule.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ecr-2015-09-21/replicationrule.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ecr-2015-09-21/replicationrule.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicationDestination

Repository

All content copied from https://docs.aws.amazon.com/.
