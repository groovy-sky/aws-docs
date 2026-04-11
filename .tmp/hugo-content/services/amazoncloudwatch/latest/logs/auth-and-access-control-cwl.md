# Identity and access management for Amazon CloudWatch Logs

Access to Amazon CloudWatch Logs requires credentials that AWS can use to authenticate your requests.
Those credentials must have permissions to access AWS resources, such as to retrieve CloudWatch Logs
data about your cloud resources. The following sections provide details on how you can use
[AWS Identity and Access Management (IAM)](../../../iam/latest/userguide/introduction.md) and CloudWatch Logs to help
secure your resources by controlling who can access them:

- [Authentication](#authentication-cwl)

- [Access control](#access-control-cwl)

## Authentication

To provide access, add permissions to your users, groups, or roles:

- Users and groups in AWS IAM Identity Center:

Create a permission set. Follow the instructions in [Create a permission set](../../../singlesignon/latest/userguide/howtocreatepermissionset.md) in the _AWS IAM Identity Center User Guide_.

- Users managed in IAM through an identity provider:

Create a role for identity federation. Follow the instructions in [Create a role for a third-party identity provider (federation)](../../../iam/latest/userguide/id-roles-create-for-idp.md)
in the _IAM User Guide_.

- IAM users:

- Create a role that your user can assume. Follow the instructions in [Create a role for an IAM user](../../../iam/latest/userguide/id-roles-create-for-user.md) in the _IAM User Guide_.

- (Not recommended) Attach a policy directly to a user or add a user to a user group. Follow the instructions in [Adding permissions to a user (console)](../../../iam/latest/userguide/id-users-change-permissions.md#users_change_permissions-add-console) in the _IAM User Guide_.

## Access control

You can have valid credentials to authenticate your requests, but unless you have
permissions you cannot create or access CloudWatch Logs resources. For example, you must have
permissions to create log streams, create log groups, and so on.

The following sections describe how to manage permissions for CloudWatch Logs. We recommend that
you read the overview first.

- [Overview of managing access permissions to your CloudWatch Logs resources](iam-access-control-overview-cwl.md)

- [Using identity-based policies (IAM policies) for CloudWatch Logs](iam-identity-based-access-control-cwl.md)

- [CloudWatch Logs permissions reference](permissions-reference-cwl.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Data protection

Overview of managing access

All content copied from https://docs.aws.amazon.com/.
