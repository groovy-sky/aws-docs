# Manage access to Amazon Q Developer with policies

###### Note

The information on this page pertains to accessing Amazon Q Developer. For information about managing
access to Amazon Q Business, see [Identity-based\
policy examples for Amazon Q Business](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/security_iam_id-based-policy-examples.html) in the _Amazon Q Business User Guide_.

The policies and examples in this topic are specific to Amazon Q in the AWS Management Console, AWS Console Mobile Application,
AWS website, AWS Documentation, and in chat applications. Other services integrated with Amazon Q might
require different policies or settings. End users of Amazon Q in third-party IDEs are not required to
use IAM policies. For more information, see the documentation for the service that contains an
Amazon Q feature or integration.

By default, users and roles don't have permission to use Amazon Q. IAM administrators can manage
access to Amazon Q Developer and its features by granting permissions to IAM identities.

The quickest way for an administrator to grant access to users is through an AWS managed policy.
The `AmazonQFullAccess` policy can be attached to IAM identities to grant full access to
Amazon Q Developer and its features. For more information about this policy, see [AWS managed policies for Amazon Q Developer](managed-policy.md).

To manage specific actions that IAM identities can perform with Amazon Q Developer, administrators can
create custom policies that define what permissions a user, group, or role has. You can also use
service control policies (SCPs) to control what Amazon Q features are available in your
organization.

For a list of all Amazon Q permissions you can control with policies, see the see the [Amazon Q Developer permissions reference](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/security_iam_permissions.html).

###### Topics

- [Policy best practices](#security_iam_policy-best-practices)

- [Assign permissions](#setting-up-assign-permissions)

- [Manage access with service control policies (SCPs)](#service-control-policies)

- [Identity-based policy examples for Amazon Q Developer](security-iam-id-based-policy-examples.md)

## Policy best practices

Identity-based policies determine whether someone can create, access, or delete Amazon Q Developer resources in your
account. These actions can incur costs for your AWS account. When you create or edit identity-based policies, follow these guidelines and
recommendations:

- **Get started with AWS managed policies and move toward least-privilege permissions**
– To get started granting permissions to your users and workloads, use the _AWS_
_managed policies_ that grant permissions for many common use cases. They are
available in your AWS account. We recommend that you reduce permissions further by
defining AWS customer managed policies that are specific to your use cases. For more information, see
[AWS managed policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_managed-vs-inline.html#aws-managed-policies) or [AWS managed policies for job functions](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_job-functions.html) in the _IAM User Guide_.

- **Apply least-privilege permissions** –
When you set permissions with IAM policies, grant only the permissions required to
perform a task. You do this by defining the actions that can be taken on specific resources
under specific conditions, also known as _least-privilege permissions_.
For more information about using IAM to apply permissions, see [Policies and permissions in IAM](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html) in the _IAM User Guide_.

- **Use conditions in IAM policies to further restrict access**
– You can add a condition to your policies to limit access to actions and resources. For example, you can write a policy condition to specify that all requests must
be sent using SSL. You can also use conditions to grant access to service actions
if they are used through a specific AWS service, such as CloudFormation. For more information, see
[IAM JSON policy elements: Condition](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_condition.html) in the _IAM User Guide_.

- **Use IAM Access Analyzer to validate your IAM policies to ensure secure and functional permissions**
– IAM Access Analyzer validates new and existing policies so that the policies adhere to the IAM policy language (JSON) and IAM best practices.
IAM Access Analyzer provides more than 100 policy checks and actionable recommendations to help
you author secure and functional policies. For more information, see [Validate policies with IAM Access Analyzer](https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-policy-validation.html) in the _IAM User Guide_.

- **Require multi-factor authentication (MFA)** –
If you have a scenario that requires IAM users or a root user in your AWS account, turn on MFA for additional security. To require
MFA when API operations are called, add MFA conditions to your policies. For
more information, see [Secure API access with MFA](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_mfa_configure-api-require.html) in the _IAM User Guide_.

For more information about best practices in IAM, see [Security best practices in IAM](https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html) in the _IAM User Guide_.

## Assign permissions

To provide access, add permissions to your users, groups, or roles:

- Users and groups in AWS IAM Identity Center:

Create a permission set. Follow the instructions in [Create a permission set](https://docs.aws.amazon.com/singlesignon/latest/userguide/howtocreatepermissionset.html) in the _AWS IAM Identity Center User Guide_.

- Users managed in IAM through an identity provider:

Create a role for identity federation. Follow the instructions in [Create a role for a third-party identity provider (federation)](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-idp.html)
in the _IAM User Guide_.

- IAM users:

- Create a role that your user can assume. Follow the instructions in [Create a role for an IAM user](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user.html) in the _IAM User Guide_.

- (Not recommended) Attach a policy directly to a user or add a user to a user group. Follow the instructions in [Adding permissions to a user (console)](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_change-permissions.html#users_change_permissions-add-console) in the _IAM User Guide_.

## Manage access with service control policies (SCPs)

Service control policies (SCPs) are a type of organization policy that you can use to manage
permissions in your organization. You can control what Amazon Q Developer features are available in your
organization by creating an SCP that specifies permissions for some or all Amazon Q actions.

For more information about using SCPs to control access in your organization, see [Creating, updating, and deleting service control policies](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scps_create.html) and [Attaching and detaching service control policies](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scps_attach.html) in the _AWS Organizations User_
_Guide_.

### Example SCP: Deny access to Amazon Q outside EU Regions

The following SCP denies access to any use of Amazon Q Developer outside of the Europe (Frankfurt) Region
(eu-central-1).

###### Note

The `codewhisperer` prefix is a legacy name from a service that merged
with Amazon Q Developer. For more information, see
[Amazon Q Developer rename - Summary of changes](https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/service-rename.html).

JSON

```json

{
  "Version":"2012-10-17",
  "Statement": [
    {
      "Sid": "DenyAmazonQDeveloperOutsideEU",
      "Effect": "Deny",
      "Action": [
         "codewhisperer:GenerateRecommendations",
         "q:SendMessage",
         "q:GenerateCodeFromCommands",
         "sqlworkbench:GetQSqlRecommendations"
         ],
      "Resource": "*",
      "Condition": {
        "StringNotEquals":
        {"aws:RequestedRegion": [ "eu-central-1"] }
      }
    }
  ]
}

```

### Example SCP: Deny access to Amazon Q

The following SCP denies access to Amazon Q Developer.

###### Note

Denying access to Amazon Q will not disable the Amazon Q icon or chat panel in the AWS
console, AWS website, AWS documentation pages, or AWS Console Mobile Application.

JSON

```json

{
  "Version":"2012-10-17",
  "Statement": [
    {
      "Sid": "DenyAmazonQFullAccess",
      "Effect": "Deny",
      "Action": [
        "q:*"
      ],
      "Resource": "*"
    }
  ]
}

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

How Amazon Q works with IAM

Identity-based policy examples
for Amazon Q
