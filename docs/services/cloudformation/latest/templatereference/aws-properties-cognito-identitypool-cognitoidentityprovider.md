This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::IdentityPool CognitoIdentityProvider

`CognitoIdentityProvider` is a property of the [AWS::Cognito::IdentityPool](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html) resource that represents an Amazon Cognito user
pool and its client ID.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClientId" : String,
  "ProviderName" : String,
  "ServerSideTokenCheck" : Boolean
}

```

### YAML

```yaml

  ClientId: String
  ProviderName: String
  ServerSideTokenCheck: Boolean

```

## Properties

`ClientId`

The client ID for the Amazon Cognito user pool.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProviderName`

The provider name for an Amazon Cognito user pool. For example:
`cognito-idp.us-east-2.amazonaws.com/us-east-2_123456789`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerSideTokenCheck`

TRUE if server-side token validation is enabled for the identity provider’s
token.

After you set the `ServerSideTokenCheck` to TRUE for an identity pool, that
identity pool checks with the integrated user pools to make sure the user has not been
globally signed out or deleted before the identity pool provides an OIDC token or
AWS credentials for the user.

If the user is signed out or deleted, the identity pool returns a 400 Not Authorized
error.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Cognito::IdentityPool

CognitoStreams
