# AWS managed policies for Amazon Athena

An AWS managed policy is a standalone policy that is created and administered by AWS. AWS managed policies are designed
to provide permissions for many common use cases so that you can start assigning permissions to users, groups, and roles.

Keep in mind that AWS managed policies might not grant least-privilege permissions for your specific use cases because
they're available for all AWS customers to use. We recommend that you reduce permissions further by defining
[customer managed policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_managed-vs-inline.html#customer-managed-policies) that are specific to your use cases.

You cannot change the permissions defined in AWS managed policies. If AWS updates the permissions defined in an AWS
managed policy, the update affects all principal identities (users, groups, and roles) that the policy is attached to. AWS is
most likely to update an AWS managed policy when a new AWS service is launched or new API operations become available for
existing services.

For more information, see [AWS managed policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_managed-vs-inline.html#aws-managed-policies) in the
_IAM User Guide_.

## Considerations when using managed policies with Athena

Managed policies are easy to use and are updated automatically with the required
actions as the service evolves. When using managed policies with Athena, keep the
following points in mind:

- To allow or deny Amazon Athena service actions for yourself or other users using
AWS Identity and Access Management (IAM), you attach identity-based policies to principals, such as
users or groups.

- Each identity-based policy consists of statements that define the actions that
are allowed or denied. For more information and step-by-step instructions for
attaching a policy to a user, see [Attaching managed policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_managed-using.html#attach-managed-policy-console) in the
_IAM User Guide_. For a list of actions, see the
[Amazon Athena API Reference](../../../../reference/athena/latest/apireference.md).

- _Customer-managed_ and _inline_
identity-based policies allow you to specify more detailed Athena actions within
a policy to fine-tune access. We recommend that you use the
`AmazonAthenaFullAccess` policy as a starting point and then
allow or deny specific actions listed in the [Amazon Athena API Reference](../../../../reference/athena/latest/apireference.md). For more
information about inline policies, see [Managed policies\
and inline policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_managed-vs-inline.html) in the
_IAM User Guide_.

- If you also have principals that connect using JDBC, you must provide the JDBC
driver credentials to your application. For more information, see [Control access through JDBC and ODBC connections](https://docs.aws.amazon.com/athena/latest/ug/policy-actions.html).

- If you have encrypted the AWS Glue Data Catalog, you must specify additional
actions in the identity-based IAM policies for Athena. For more information, see
[Configure access from Athena to encrypted metadata in the AWS Glue Data Catalog](https://docs.aws.amazon.com/athena/latest/ug/access-encrypted-data-glue-data-catalog.html).

- If you create and use workgroups, make sure your policies include relevant
access to workgroup actions. For detailed information, see [Use IAM policies to control workgroup access](https://docs.aws.amazon.com/athena/latest/ug/workgroups-iam-policy.html) and [Example workgroup policies](https://docs.aws.amazon.com/athena/latest/ug/example-policies-workgroup.html).

## AWS managed policy: AmazonAthenaFullAccess

The `AmazonAthenaFullAccess` managed policy grants full access to
Athena.

To provide access, add permissions to your users, groups, or roles:

- Users and groups in AWS IAM Identity Center:

Create a permission set. Follow the instructions in [Create a permission set](https://docs.aws.amazon.com/singlesignon/latest/userguide/howtocreatepermissionset.html) in the _AWS IAM Identity Center User Guide_.

- Users managed in IAM through an identity provider:

Create a role for identity federation. Follow the instructions in [Create a role for a third-party identity provider (federation)](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-idp.html)
in the _IAM User Guide_.

- IAM users:

- Create a role that your user can assume. Follow the instructions in [Create a role for an IAM user](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user.html) in the _IAM User Guide_.

- (Not recommended) Attach a policy directly to a user or add a user to a user group. Follow the instructions in [Adding permissions to a user (console)](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_change-permissions.html#users_change_permissions-add-console) in the _IAM User Guide_.

### Permissions groupings

The `AmazonAthenaFullAccess` policy is grouped into the following sets
of permissions.

- `athena` – Allows
principals access to Athena resources.

- `glue` – Allows
principals access to AWS Glue Catalogs, databases, tables, and partitions. This
is required so that the principal can use the AWS Glue Data Catalogs with
Athena.

- `s3` – Allows the
principal to write and read query results from Amazon S3, to read publically
available Athena data examples that reside in Amazon S3, and to list buckets. This
is required so that the principal can use Athena to work with Amazon S3.

- `sns` – Allows principals
to list Amazon SNS topics and get topic attributes. This enables principals to
use Amazon SNS topics with Athena for monitoring and alert purposes.

- `cloudwatch` – Allows
principals to create, read, and delete CloudWatch alarms. For more information,
see [Use CloudWatch and EventBridge to monitor queries and control costs](https://docs.aws.amazon.com/athena/latest/ug/workgroups-control-limits.html).

- `lakeformation` – Allows
principals to request temporary credentials to access data in a data lake
location that is registered with Lake Formation. For more information, see [Underlying data access control](https://docs.aws.amazon.com/lake-formation/latest/dg/access-control-underlying-data.html) in the _AWS Lake Formation Developer Guide_.

- `datazone` – Allows
principals to list Amazon DataZone projects, domains, and environments. For
information about using DataZone in Athena, see [Use Amazon DataZone in Athena](https://docs.aws.amazon.com/athena/latest/ug/datazone-using.html).

- `pricing` – Provides
access to AWS Billing and Cost Management. For more information, see [GetProducts](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_pricing_GetProducts.html) in the _AWS Billing and Cost Management API Reference_.

To view the permissions for this policy, see [AmazonAthenaFullAccess](https://docs.aws.amazon.com/aws-managed-policy/latest/reference/AmazonAthenaFullAccess.html) in the AWS Managed Policy Reference.

###### Note

You must explicitly allow access to service-owned Amazon S3 buckets to store example queries and
sample dataset. For more information, see [Data perimeters](https://docs.aws.amazon.com/athena/latest/ug/data-perimeters.html).

## AWS managed policy: AWSQuicksightAthenaAccess

`AWSQuicksightAthenaAccess` grants access to actions that Quick requires for
integration with Athena. You can attach the `AWSQuicksightAthenaAccess` policy
to your IAM identities. Attach this policy only to principals who use Quick with Athena.
This policy includes some actions for Athena that are either deprecated and not included
in the current public API, or that are used only with the JDBC and ODBC drivers.

### Permissions groupings

The `AWSQuicksightAthenaAccess` policy is grouped into the following
sets of permissions.

- `athena` – Allows the
principal to run queries on Athena resources.

- `glue` – Allows
principals access to AWS Glue Catalogs, databases, tables, and partitions. This
is required so that the principal can use the AWS Glue Data Catalogs with
Athena.

- `s3` – Allows the
principal to write and read query results from Amazon S3.

- `lakeformation` – Allows
principals to request temporary credentials to access data in a data lake
location that is registered with Lake Formation. For more information, see [Underlying data access control](https://docs.aws.amazon.com/lake-formation/latest/dg/access-control-underlying-data.html) in the _AWS Lake Formation Developer Guide_.

To view the permissions for this policy, see [AWSQuicksightAthenaAccess](https://docs.aws.amazon.com/aws-managed-policy/latest/reference/AWSQuicksightAthenaAccess.html) in the AWS Managed Policy Reference.

## Athena updates to AWS managed policies

View details about updates to AWS managed policies for Athena since this service
began tracking these changes.

ChangeDescriptionDate[AWSQuicksightAthenaAccess](#awsquicksightathenaaccess-managed-policy) – Updates to existing
policiesThe `glue:GetCatalog` and `glue:GetCatalogs`
permissions were added to enable Athena users to access to SageMaker AI Lakehouse
catalogs.January 02, 2025[AmazonAthenaFullAccess](#amazonathenafullaccess-managed-policy) – Update to existing
policyThe `glue:GetCatalog` and `glue:GetCatalogs`
permissions were added to enable Athena users to access to SageMaker AI Lakehouse
catalogs.January 02, 2025[AmazonAthenaFullAccess](#amazonathenafullaccess-managed-policy) – Update to existing
policy

Enables Athena to use the publicly documented AWS Glue
`GetCatalogImportStatus` API to retrieve catalog
import status.

June 18, 2024

[AmazonAthenaFullAccess](#amazonathenafullaccess-managed-policy) – Update to existing
policy

The `datazone:ListDomains`,
`datazone:ListProjects`, and
`datazone:ListAccountEnvironments` permissions were
added to enable Athena users to work with Amazon DataZone domains, projects,
and environments. For more information, see [Use Amazon DataZone in Athena](https://docs.aws.amazon.com/athena/latest/ug/datazone-using.html).

January 3, 2024

[AmazonAthenaFullAccess](#amazonathenafullaccess-managed-policy) – Update to existing
policy

The `glue:StartColumnStatisticsTaskRun`,
`glue:GetColumnStatisticsTaskRun`, and
`glue:GetColumnStatisticsTaskRuns` permissions were
added to give Athena the right to call AWS Glue to retrieve statistics
for the cost-based optimizer feature. For more information, see
[Use the cost-based optimizer](https://docs.aws.amazon.com/athena/latest/ug/cost-based-optimizer.html).

January 3, 2024

[AmazonAthenaFullAccess](#amazonathenafullaccess-managed-policy) – Update to existing
policy

Athena added `pricing:GetProducts` to provide access to
AWS Billing and Cost Management. For more information, see [GetProducts](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_pricing_GetProducts.html) in the _AWS Billing and Cost Management API Reference_.

January 25, 2023

[AmazonAthenaFullAccess](#amazonathenafullaccess-managed-policy) – Update to existing
policy

Athena added `cloudwatch:GetMetricData` to retrieve CloudWatch
metric values. For more information, see [GetMetricData](https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_GetMetricData.html) in the
_Amazon CloudWatch API Reference_.

November 14, 2022

[AmazonAthenaFullAccess](#amazonathenafullaccess-managed-policy) and [AWSQuicksightAthenaAccess](#awsquicksightathenaaccess-managed-policy) – Updates to existing
policies

Athena added `s3:PutBucketPublicAccessBlock` to enable
the blocking of public access on the buckets created by
Athena.

July 7, 2021

Athena started tracking changes

Athena started tracking changes for its AWS managed
policies.

July 7, 2021

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Identity and access
management

Data perimeters
