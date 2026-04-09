# JWT credentials

With this authentication type, you can use a JSON web token (JWT) obtained from an
external identity provider as a connection parameter to authenticate with Athena. The
external credentials provider must already be federated with AWS.

## Credentials provider

The credentials provider that will be used to authenticate requests to AWS. Set
the value of this parameter to `JWT`.

Parameter nameAliasParameter typeDefault valueValue to useCredentialsProviderAWSCredentialsProviderClass (deprecated)Requirednone`JWT`

## JWT web identity token

The JWT token obtained from an external federated identity provider. This token
will be used to authenticate with Athena.

Parameter nameAliasParameter typeDefault valueJwtWebIdentityTokenweb\_identity\_token (deprecated)Requirednone

## JWT role ARN

The Amazon Resource Name (ARN) of the role to assume. For information about
assuming roles, see [AssumeRole](../../../../reference/sts/latest/apireference/api-assumerole.md) in the
_AWS Security Token Service API Reference_.

Parameter nameAliasParameter typeDefault valueJwtRoleArnrole\_arn (deprecated)Requirednone

## JWT role session name

The name of the session when you use JWT credentials for authentication. The name
can be any name that you choose.

Parameter nameAliasParameter typeDefault valueJwtRoleSessionNamerole\_session\_name (deprecated)Requirednone

## Role session duration

The duration, in seconds, of the role session. For more information, see [AssumeRoleWithWebIdentity](../../../../reference/sts/latest/apireference/api-assumerolewithwebidentity.md) in the
_AWS Security Token Service API Reference_.

Parameter nameAliasParameter typeDefault valueRoleSessionDurationDuration (deprecated)Optional3600

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Custom

JWT trusted identity propagation

All content copied from https://docs.aws.amazon.com/.
