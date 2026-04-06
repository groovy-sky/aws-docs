# Identity and access management in Amazon Route 53

To perform any operation on Amazon Route 53 resources, such as registering a domain or updating a
record, AWS Identity and Access Management (IAM) requires you to authenticate that you're an approved AWS user.
If you're using the Route 53 console, you authenticate your identity by providing your AWS
user name and a password.

After you authenticate your identity, IAM controls your access to AWS by verifying
that you have permissions to perform operations and to access resources. If you are an
account administrator, you can use IAM to control the access of other users to the
resources that are associated with your account.

This chapter explains how to use [IAM](https://docs.aws.amazon.com/IAM/latest/UserGuide/introduction.html)
and Route 53 to help secure your resources.

###### Topics

- [Authenticating with identities](#security_iam_authentication)

- [Access control](#access-control)

- [Using Service-Linked Roles for Amazon Route 53 Resolver](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/using-service-linked-roles.html)

- [AWS managed policies for Amazon Route 53](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/security-iam-awsmanpol-route53.html)

- [Using IAM policy conditions for fine-grained access control](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/specifying-conditions-route53.html)

- [Amazon Route 53 API permissions: Actions, resources, and conditions reference](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/r53-api-permissions-ref.html)

## Authenticating with identities

Authentication is how you sign in to AWS using your identity credentials. You must be authenticated as the AWS account root user, an IAM user, or by assuming an IAM role.

You can sign in as a federated identity using credentials from an identity source like AWS IAM Identity Center (IAM Identity Center), single sign-on authentication, or Google/Facebook credentials. For more information about signing in, see [How to sign in to your AWS account](https://docs.aws.amazon.com/signin/latest/userguide/how-to-sign-in.html) in the _AWS Sign-In User Guide_.

For programmatic access, AWS provides an SDK and CLI to cryptographically sign requests. For more information, see [AWS Signature Version 4 for API requests](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_sigv.html) in the _IAM User Guide_.

### AWS account root user

When you create an AWS account, you begin with one sign-in identity called the AWS account _root user_ that has complete access to all AWS services and resources. We strongly recommend that you don't use the root user for everyday tasks. For tasks that require root user credentials, see [Tasks that require root user credentials](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_root-user.html#root-user-tasks) in the _IAM User Guide_.

### Federated identity

As a best practice, require human users to use federation with an identity provider to access AWS services using temporary credentials.

A _federated identity_ is a user from your enterprise directory, web identity provider, or Directory Service that accesses AWS services using credentials from an identity source. Federated identities assume roles that provide temporary credentials.

For centralized access management, we recommend AWS IAM Identity Center. For more information, see [What is IAM Identity Center?](https://docs.aws.amazon.com/singlesignon/latest/userguide/what-is.html) in the _AWS IAM Identity Center User Guide_.

### IAM users and groups

An _[IAM user](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users.html)_ is an identity with specific permissions for a single person or application. We recommend using temporary credentials instead of IAM users with long-term credentials. For more information, see [Require human users to use federation with an identity provider to access AWS using temporary credentials](https://docs.aws.amazon.com/IAM/latest/UserGuide/best-practices.html#bp-users-federation-idp) in the _IAM User Guide_.

An [_IAM group_](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_groups.html) specifies a collection of IAM users and makes permissions easier to manage for large sets of users. For more information, see [Use cases for IAM users](https://docs.aws.amazon.com/IAM/latest/UserGuide/gs-identities-iam-users.html) in the _IAM User Guide_.

### IAM roles

An _[IAM role](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html)_ is an identity with specific permissions that provides temporary credentials. You can assume a role by [switching from a user to an IAM role (console)](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_switch-role-console.html) or by calling an AWS CLI or AWS API operation. For more information, see [Methods to assume a role](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_manage-assume.html) in the _IAM User Guide_.

IAM roles are useful for federated user access, temporary IAM user permissions, cross-account access, cross-service access, and applications running on Amazon EC2. For more information, see [Cross account resource access in IAM](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies-cross-account-resource-access.html) in the _IAM User Guide_.

## Access control

To create, update, delete, or list Amazon Route 53 resources, you need permissions to perform the
operation, and you need permission to access the corresponding resources.

The following sections describe how to manage permissions for Route 53. We recommend that
you read the overview first.

###### Topics

- [Overview of managing access permissions to your Amazon Route 53 resources](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/access-control-overview.html)

- [Using identity-based policies (IAM policies) for Amazon Route 53](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/access-control-managing-permissions.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Protection from dangling delegation records

Using Service-Linked Roles
