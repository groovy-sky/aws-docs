# DBInstanceRole

Information about an AWS Identity and Access Management (IAM) role that is associated with a DB instance.

## Contents

###### Note

In the following list, the required parameters are described first.

**FeatureName**

The name of the feature associated with the AWS Identity and Access Management (IAM) role.
For information about supported feature names, see `DBEngineVersion`.

Type: String

Required: No

**RoleArn**

The Amazon Resource Name (ARN) of the IAM role that is associated with the DB
instance.

Type: String

Required: No

**Status**

Information about the state of association between the IAM role and the DB instance. The Status property returns one of the following
values:

- `ACTIVE` \- the IAM role ARN is associated with the DB instance and can be used to
access other AWS services on your behalf.

- `PENDING` \- the IAM role ARN is being associated with the DB instance.

- `INVALID` \- the IAM role ARN is associated with the DB instance, but the DB instance is unable
to assume the IAM role in order to access other AWS services on your behalf.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/dbinstancerole.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/dbinstancerole.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/dbinstancerole.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DBInstanceAutomatedBackupsReplication

DBInstanceStatusInfo

All content copied from https://docs.aws.amazon.com/.
