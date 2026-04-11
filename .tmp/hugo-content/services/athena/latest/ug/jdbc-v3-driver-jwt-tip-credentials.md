# JWT with identity center integration

This authentication type allows you to use a JSON web token (JWT) obtained from an
external identity provider as a connection parameter to authenticate with Athena. You can
use this plugin, to enable support for corporate identities via trusted identity propagation. For more
information on how to use trusted identity propagation with drivers, see [Use Trusted identity propagation with Amazon Athena drivers](using-trusted-identity-propagation.md). You can also [configure and deploy\
resources using CloudFormation](using-trusted-identity-propagation-cloudformation.md).

With trusted identity propagation, identity context is added to an IAM role to identify the user requesting
access to AWS resources. For information on enabling and using trusted identity propagation, see [What is trusted identity propagation?](../../../singlesignon/latest/userguide/trustedidentitypropagation.md).

## Credentials provider

The credentials provider that will be used to authenticate requests to AWS. Set
the value of this parameter to `JWT_TIP`.

Parameter nameAliasParameter typeDefault valueValue to useCredentialsProviderAWSCredentialsProviderClass (deprecated)Requirednone`JWT_TIP`

## JWT web identity token

The JWT token obtained from an external federated identity provider. This token
will be used to authenticate with Athena. Token Caching is enabled by default and
allows the same Identity Center access token to be used across driver connections.
We recommend to provide a fresh JWT token upon "Testing Connection" as the exchanged
token is present only during driver instance is active.

Parameter nameAliasParameter typeDefault valueJwtWebIdentityTokenweb\_identity\_token (deprecated)Requirednone

## WorkgroupArn

The Amazon Resource Name (ARN) of the Amazon Athena workgroup. For more information
about workgroups, see [WorkGroup](../../../../reference/athena/latest/apireference/api-workgroup.md).

Parameter nameAliasParameter typeDefault valueWorkGroupArnnoneRequiredprimary

## JWT application role ARN

The ARN of the role to assume. This role is used for JWT exchange, getting IAM Identity Center
customer managed application ARN through workgroup tags, and getting access role
ARN. For more information about assuming roles, see [AssumeRole](../../../../reference/sts/latest/apireference/api-assumerole.md).

Parameter nameAliasParameter typeDefault valueApplicationRoleArnnoneRequirednone

## JWT role session name

The name of the session when authenticating with JWT credentials. It can be any
name of your choice.

Parameter nameAliasParameter typeDefault valueJwtRoleSessionNamerole\_session\_name (deprecated)Requirednone

## Role session duration

The duration, in seconds, of the role session. For more information, see [AssumeRoleWithWebIdentity](../../../../reference/sts/latest/apireference/api-assumerolewithwebidentity.md).

Parameter nameAliasParameter typeDefault valueRoleSessionDurationDuration (deprecated)Optional3600

## JWT access role ARN

The ARN of the role to assume. This is the role assumed by the Athena service to
make calls on the behalf of you. For more information about assuming roles, see
[AssumeRole](../../../../reference/sts/latest/apireference/api-assumerole.md) in the _AWS Security Token Service API Reference_.

Parameter nameAliasParameter typeDefault valueAccessRoleArnnoneOptionalnone

## IAM Identity Center customer managed application ARN

The ARN of IAM Identity Center customer managed application. For more information, see [customer managed\
applications](../../../singlesignon/latest/userguide/customermanagedapps.md).

Parameter nameAliasParameter typeDefault valueCustomerIdcApplicationArnnoneOptionalnone

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JWT

Browser trusted identity propagation

All content copied from https://docs.aws.amazon.com/.
