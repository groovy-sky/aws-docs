# ReplicationConfiguration

A container for replication rules. You can add up to 1,000 rules. The maximum size of a replication
configuration is 2 MB.

## Contents

**Role**

The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that Amazon S3 assumes when replicating
objects. For more information, see [How to Set Up Replication](../dev/replication-how-setup.md) in the
_Amazon S3 User Guide_.

Type: String

Required: Yes

**Rules**

A container for one or more replication rules. A replication configuration must have at least one
rule and can contain a maximum of 1,000 rules.

Type: Array of [ReplicationRule](api-replicationrule.md) data types

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/replicationconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/replicationconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/replicationconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicaModifications

ReplicationRule

All content copied from https://docs.aws.amazon.com/.
