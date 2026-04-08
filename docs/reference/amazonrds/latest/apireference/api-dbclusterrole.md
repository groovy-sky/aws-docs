# DBClusterRole

Describes an AWS Identity and Access Management (IAM) role that is associated with a DB cluster.

## Contents

###### Note

In the following list, the required parameters are described first.

**FeatureName**

The name of the feature associated with the AWS Identity and Access Management (IAM) role.
For information about supported feature names, see [DBEngineVersion](api-dbengineversion.md).

Type: String

Required: No

**RoleArn**

The Amazon Resource Name (ARN) of the IAM role that is associated with the DB cluster.

Type: String

Required: No

**Status**

Describes the state of association between the IAM role and the DB cluster. The Status property returns one of the following
values:

- `ACTIVE` \- the IAM role ARN is associated with the DB cluster and can be used to
access other Amazon Web Services on your behalf.

- `PENDING` \- the IAM role ARN is being associated with the DB cluster.

- `INVALID` \- the IAM role ARN is associated with the DB cluster, but the DB cluster is unable
to assume the IAM role in order to access other Amazon Web Services on your behalf.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/dbclusterrole.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/dbclusterrole.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/dbclusterrole.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DBClusterParameterGroup

DBClusterSnapshot

All content copied from https://docs.aws.amazon.com/.
