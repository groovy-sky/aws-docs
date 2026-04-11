# Identity and access management for Aurora DSQL

AWS Identity and Access Management (IAM) is an AWS service that helps an administrator securely control access
to AWS resources. IAM administrators control who can be _authenticated_ (signed in) and _authorized_
(have permissions) to use Aurora DSQL resources. IAM is an AWS service that you can
use with no additional charge.

###### Topics

- [Audience](#security_iam_audience)

- [Authenticating with identities](#security_iam_authentication)

- [Managing access using policies](#security_iam_access-manage)

- [How Amazon Aurora DSQL works with IAM](security-iam-service-with-iam.md)

- [Identity-based policy examples for Amazon Aurora DSQL](security-iam-id-based-policy-examples.md)

- [Troubleshooting Amazon Aurora DSQL identity and access](security-iam-troubleshoot.md)

## Audience

How you use AWS Identity and Access Management (IAM) differs based on your role:

- **Service user** \- request permissions from your
administrator if you cannot access features (see [Troubleshooting Amazon Aurora DSQL identity and access](security-iam-troubleshoot.md))

- **Service administrator** \- determine user access and
submit permission requests (see [How Amazon Aurora DSQL works with IAM](security-iam-service-with-iam.md))

- **IAM administrator** \- write policies to manage
access (see [Identity-based policy examples for Amazon Aurora DSQL](security-iam-id-based-policy-examples.md))

## Authenticating with identities

Authentication is how you sign in to AWS using your identity credentials. You must be authenticated as the AWS account root user, an IAM user, or by assuming an IAM role.

You can sign in as a federated identity using credentials from an identity source like AWS IAM Identity Center (IAM Identity Center), single sign-on authentication, or Google/Facebook credentials. For more information about signing in, see [How to sign in to your AWS account](../../../signin/latest/userguide/how-to-sign-in.md) in the _AWS Sign-In User Guide_.

For programmatic access, AWS provides an SDK and CLI to cryptographically sign requests. For more information, see [AWS Signature Version 4 for API requests](../../../iam/latest/userguide/reference-sigv.md) in the _IAM User Guide_.

### AWS account root user

When you create an AWS account, you begin with one sign-in identity called the AWS account _root user_ that has complete access to all AWS services and resources. We strongly recommend that you don't use the root user for everyday tasks. For tasks that require root user credentials, see [Tasks that require root user credentials](../../../iam/latest/userguide/id-root-user.md#root-user-tasks) in the _IAM User Guide_.

### Federated identity

As a best practice, require human users to use federation with an identity provider to access AWS services using temporary credentials.

A _federated identity_ is a user from your enterprise directory, web identity provider, or Directory Service that accesses AWS services using credentials from an identity source. Federated identities assume roles that provide temporary credentials.

For centralized access management, we recommend AWS IAM Identity Center. For more information, see [What is IAM Identity Center?](../../../singlesignon/latest/userguide/what-is.md) in the _AWS IAM Identity Center User Guide_.

### IAM users and groups

An _[IAM user](../../../iam/latest/userguide/id-users.md)_ is an identity with specific permissions for a single person or application. We recommend using temporary credentials instead of IAM users with long-term credentials. For more information, see [Require human users to use federation with an identity provider to access AWS using temporary credentials](../../../iam/latest/userguide/best-practices.md#bp-users-federation-idp) in the _IAM User Guide_.

An [_IAM group_](../../../iam/latest/userguide/id-groups.md) specifies a collection of IAM users and makes permissions easier to manage for large sets of users. For more information, see [Use cases for IAM users](../../../iam/latest/userguide/gs-identities-iam-users.md) in the _IAM User Guide_.

### IAM roles

An _[IAM role](../../../iam/latest/userguide/id-roles.md)_ is an identity with specific permissions that provides temporary credentials. You can assume a role by [switching from a user to an IAM role (console)](../../../iam/latest/userguide/id-roles-use-switch-role-console.md) or by calling an AWS CLI or AWS API operation. For more information, see [Methods to assume a role](../../../iam/latest/userguide/id-roles-manage-assume.md) in the _IAM User Guide_.

IAM roles are useful for federated user access, temporary IAM user permissions, cross-account access, cross-service access, and applications running on Amazon EC2. For more information, see [Cross account resource access in IAM](../../../iam/latest/userguide/access-policies-cross-account-resource-access.md) in the _IAM User Guide_.

## Managing access using policies

You control access in AWS by creating policies and attaching them to AWS identities or resources. A policy defines permissions when associated with an identity or resource. AWS evaluates these policies when a principal makes a request. Most policies are stored in AWS as JSON documents. For more information about JSON policy documents, see [Overview of JSON policies](../../../iam/latest/userguide/access-policies.md#access_policies-json) in the _IAM User Guide_.

Using policies, administrators specify who has access to what by defining which **principal** can perform **actions** on what **resources**, and under what **conditions**.

By default, users and roles have no permissions. An IAM administrator creates IAM policies and adds them to roles, which users can then assume. IAM policies define permissions regardless of the method used to perform the operation.

### Identity-based policies

Identity-based policies are JSON permissions policy documents that you attach to an identity (user, group, or role). These policies control what actions identities can perform, on which resources, and under what conditions. To learn how to create an identity-based policy, see [Define custom IAM permissions with customer managed policies](../../../iam/latest/userguide/access-policies-create.md) in the _IAM User Guide_.

Identity-based policies can be _inline policies_ (embedded directly into a single identity) or _managed policies_ (standalone policies attached to multiple identities). To learn how to choose between managed and inline policies, see [Choose between managed policies and inline policies](../../../iam/latest/userguide/access-policies-choosing-managed-or-inline.md) in the _IAM User Guide_.

### Resource-based policies

Resource-based policies are JSON policy documents that you attach to a resource. Examples include IAM _role trust policies_ and Amazon S3 _bucket policies_. In services that support resource-based policies, service administrators can use them to control access to a specific resource. You must [specify a principal](../../../iam/latest/userguide/reference-policies-elements-principal.md) in a resource-based policy.

Resource-based policies are inline policies that are located in that service. You can't use AWS managed policies from IAM in a resource-based policy.

### Other policy types

AWS supports additional policy types that can set the maximum permissions granted by more common policy types:

- **Permissions boundaries** – Set the maximum permissions that an identity-based policy can grant to an IAM entity. For more information, see [Permissions boundaries for IAM entities](../../../iam/latest/userguide/access-policies-boundaries.md) in the _IAM User Guide_.

- **Service control policies (SCPs)** – Specify the maximum permissions for an organization or organizational unit in AWS Organizations. For more information, see [Service control policies](../../../organizations/latest/userguide/orgs-manage-policies-scps.md) in the _AWS Organizations User Guide_.

- **Resource control policies (RCPs)** – Set the maximum available permissions for resources in your accounts. For more information, see [Resource control policies (RCPs)](../../../organizations/latest/userguide/orgs-manage-policies-rcps.md) in the _AWS Organizations User Guide_.

- **Session policies** – Advanced policies passed as a parameter when creating a temporary session for a role or federated user. For more information, see [Session policies](../../../iam/latest/userguide/access-policies.md#policies_session) in the _IAM User Guide_.

### Multiple policy types

When multiple types of policies apply to a request, the resulting permissions are more complicated to understand. To learn how AWS determines whether to allow a request when multiple policy types are involved, see [Policy evaluation logic](../../../iam/latest/userguide/reference-policies-evaluation-logic.md) in the _IAM User Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Data encryption

How Aurora DSQL works with IAM

All content copied from https://docs.aws.amazon.com/.
