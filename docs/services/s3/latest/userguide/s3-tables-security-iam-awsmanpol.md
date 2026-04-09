# AWS managed policies for S3 Tables

An AWS managed policy is a standalone policy that is created and administered by AWS. AWS managed policies are designed
to provide permissions for many common use cases so that you can start assigning permissions to users, groups, and roles.

Keep in mind that AWS managed policies might not grant least-privilege permissions for your specific use cases because
they're available for all AWS customers to use. We recommend that you reduce permissions further by defining
[customer managed policies](../../../iam/latest/userguide/access-policies-managed-vs-inline.md#customer-managed-policies) that are specific to your use cases.

You cannot change the permissions defined in AWS managed policies. If AWS updates the permissions defined in an AWS
managed policy, the update affects all principal identities (users, groups, and roles) that the policy is attached to. AWS is
most likely to update an AWS managed policy when a new AWS service is launched or new API operations become available for
existing services.

For more information, see [AWS managed policies](../../../iam/latest/userguide/access-policies-managed-vs-inline.md#aws-managed-policies) in the
_IAM User Guide_.

## AWS managed policy: AmazonS3TablesFullAccess

You can attach the `AmazonS3TablesFullAccess` policy to your IAM
identities. This policy grants permissions that allow full access to Amazon S3
Tables. For more information about this policy, see
[AmazonS3TablesFullAccess](../../../aws-managed-policy/latest/reference/amazons3tablesfullaccess.md).

## AWS managed policy: AmazonS3TablesReadOnlyAccess

You can attach the `AmazonS3TablesReadOnlyAccess` policy to your IAM
identities. This policy grants permissions that allow read-only access to Amazon S3
Tables. For more information about this policy, see
[AmazonS3TablesReadOnlyAccess](../../../aws-managed-policy/latest/reference/amazons3tablesreadonlyaccess.md).

## AWS managed policy: AmazonS3TablesLakeFormationServiceRole

You can attach the `AmazonS3TablesLakeFormationServiceRole` policy to your IAM
identities. This policy grants permissions that allow the AWS Lake Formation service role access to S3 Tables. AWS KMS permissions are used to allow Lake Formation to access
encrypted tables. For more information about this policy, see
[AmazonS3TablesLakeFormationServiceRole](../../../aws-managed-policy/latest/reference/amazons3tableslakeformationservicerole.md).

## Amazon S3 Tables updates to AWS managed policies

View details about updates to AWS managed policies for Amazon S3 Tables since S3 Tables began tracking these changes.

ChangeDescriptionDate

Amazon S3 Tables updated `AmazonS3TablesFullAccess`.

S3 Tables updated the AWS-managed policy called `AmazonS3TablesFullAccess`. This update grants permission to pass a role to the S3 Tables Replication service.

December 2, 2025

Amazon S3 Tables added `AmazonS3TablesLakeFormationServiceRole`.

S3 Tables added a new AWS-managed policy called `AmazonS3TablesLakeFormationServiceRole`.
This policy grants permissions that allows the Lake Formation service role access to S3 Tables.

May 19, 2025

Amazon S3 Tables added `AmazonS3TablesFullAccess`.

S3 Tables added a new AWS-managed policy called `AmazonS3TablesFullAccess`.
This policy grants permissions that allow full access to Amazon S3
Tables.

December 03, 2024

Amazon S3 Tables added `AmazonS3TablesReadOnlyAccess`.

S3 Tables added a new AWS-managed policy called `AmazonS3TablesReadOnlyAccess`.
This policy grants permissions to allow read-only access to Amazon S3
Tables.

December 03, 2024

Amazon S3 Tables started tracking changes.

Amazon S3 Tables started tracking changes for its AWS managed policies.

December 03, 2024

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Resource-based
policies

Granting access with SQL semantics

All content copied from https://docs.aws.amazon.com/.
