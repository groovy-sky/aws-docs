# Security best practices for AWS AppSync

Securing AWS AppSync is more than simply turning on a few levers or setting up logging. The
following sections discuss security best practices that vary depending on how you use the
service.

## Understand authentication methods

AWS AppSync provides multiple ways to authenticate your users to your GraphQL APIs. Each method
has trade-offs in security, auditability, and usability.

The following common authentication methods are available:

- Amazon Cognito user pools allow your GraphQL API to use user attributes for fine-grained access
control and filtering.

- API tokens have a limited lifetime and are appropriate for automated systems, such as
Continuous Integration systems and integration with external
APIs.

- AWS Identity and Access Management (IAM) is appropriate for internal applications managed in your
AWS account.

- OpenID Connect allows you to control and federate access with the OpenID Connect
protocol.

For more information on authentication and authorization in AWS AppSync, see [Configuring authorization and authentication to secure your GraphQL APIs](security-authz.md).

## Understand how API configuration changes propagate

When you save changes to your API configuration, AWS AppSync starts to propagate the changes.
Until your configuration change is propagated, AWS AppSync continues to serve your content from the
previous configuration. After your configuration change is propagated, AWS AppSync immediately starts
to serve your content based on the new configuration. While AWS AppSync is propagating your changes
for an API, we can't determine whether the API is serving your content based on the previous
configuration or the new configuration.

## Use TLS for HTTP resolvers

When using HTTP resolvers, make sure to use TLS-secured (HTTPS) connections wherever
possible. For a full list of the TLS certificates that AWS AppSync trusts, see [Certificate Authorities (CA) Recognized by AWS AppSync for HTTPS Endpoints](http-cert-authorities.md).

## Use roles with the least permissions possible

When using resolvers such as the [DynamoDB\
resolver](tutorial-dynamodb-resolvers.md), use roles that provide the most restrictive view to your
resources, such as your
Amazon DynamoDB
tables.

## IAM policy best practices

Identity-based policies determine whether someone can create, access, or delete AWS AppSync resources in your
account. These actions can incur costs for your AWS account. When you create or edit identity-based policies, follow these guidelines and
recommendations:

- **Get started with AWS managed policies and move toward least-privilege permissions**
– To get started granting permissions to your users and workloads, use the _AWS_
_managed policies_ that grant permissions for many common use cases. They are
available in your AWS account. We recommend that you reduce permissions further by
defining AWS customer managed policies that are specific to your use cases. For more information, see
[AWS managed policies](../../../iam/latest/userguide/access-policies-managed-vs-inline.md#aws-managed-policies) or [AWS managed policies for job functions](../../../iam/latest/userguide/access-policies-job-functions.md) in the _IAM User Guide_.

- **Apply least-privilege permissions** –
When you set permissions with IAM policies, grant only the permissions required to
perform a task. You do this by defining the actions that can be taken on specific resources
under specific conditions, also known as _least-privilege permissions_.
For more information about using IAM to apply permissions, see [Policies and permissions in IAM](../../../iam/latest/userguide/access-policies.md) in the _IAM User Guide_.

- **Use conditions in IAM policies to further restrict access**
– You can add a condition to your policies to limit access to actions and resources. For example, you can write a policy condition to specify that all requests must
be sent using SSL. You can also use conditions to grant access to service actions
if they are used through a specific AWS service, such as CloudFormation. For more information, see
[IAM JSON policy elements: Condition](../../../iam/latest/userguide/reference-policies-elements-condition.md) in the _IAM User Guide_.

- **Use IAM Access Analyzer to validate your IAM policies to ensure secure and functional permissions**
– IAM Access Analyzer validates new and existing policies so that the policies adhere to the IAM policy language (JSON) and IAM best practices.
IAM Access Analyzer provides more than 100 policy checks and actionable recommendations to help
you author secure and functional policies. For more information, see [Validate policies with IAM Access Analyzer](../../../iam/latest/userguide/access-analyzer-policy-validation.md) in the _IAM User Guide_.

- **Require multi-factor authentication (MFA)** –
If you have a scenario that requires IAM users or a root user in your AWS account, turn on MFA for additional security. To require
MFA when API operations are called, add MFA conditions to your policies. For
more information, see [Secure API access with MFA](../../../iam/latest/userguide/id-credentials-mfa-configure-api-require.md) in the _IAM User Guide_.

For more information about best practices in IAM, see [Security best practices in IAM](../../../iam/latest/userguide/best-practices.md) in the _IAM User Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Logging AWS AppSync API calls with AWS CloudTrail

Resolver reference (JavaScript)

All content copied from https://docs.aws.amazon.com/.
