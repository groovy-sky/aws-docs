This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::ActionConnector AuthenticationMetadata

Union type containing authentication metadata for different authentication methods.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApiKeyConnectionMetadata" : APIKeyConnectionMetadata,
  "AuthorizationCodeGrantMetadata" : AuthorizationCodeGrantMetadata,
  "BasicAuthConnectionMetadata" : BasicAuthConnectionMetadata,
  "ClientCredentialsGrantMetadata" : ClientCredentialsGrantMetadata,
  "IamConnectionMetadata" : IAMConnectionMetadata,
  "NoneConnectionMetadata" : NoneConnectionMetadata
}

```

### YAML

```yaml

  ApiKeyConnectionMetadata:
    APIKeyConnectionMetadata
  AuthorizationCodeGrantMetadata:
    AuthorizationCodeGrantMetadata
  BasicAuthConnectionMetadata:
    BasicAuthConnectionMetadata
  ClientCredentialsGrantMetadata:
    ClientCredentialsGrantMetadata
  IamConnectionMetadata:
    IAMConnectionMetadata
  NoneConnectionMetadata:
    NoneConnectionMetadata

```

## Properties

`ApiKeyConnectionMetadata`

API key authentication metadata.

_Required_: No

_Type_: [APIKeyConnectionMetadata](aws-properties-quicksight-actionconnector-apikeyconnectionmetadata.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthorizationCodeGrantMetadata`

OAuth 2.0 authorization code grant authentication metadata.

_Required_: No

_Type_: [AuthorizationCodeGrantMetadata](aws-properties-quicksight-actionconnector-authorizationcodegrantmetadata.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BasicAuthConnectionMetadata`

Basic authentication metadata using username and password.

_Required_: No

_Type_: [BasicAuthConnectionMetadata](aws-properties-quicksight-actionconnector-basicauthconnectionmetadata.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientCredentialsGrantMetadata`

OAuth 2.0 client credentials grant authentication metadata.

_Required_: No

_Type_: [ClientCredentialsGrantMetadata](aws-properties-quicksight-actionconnector-clientcredentialsgrantmetadata.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IamConnectionMetadata`

IAM role-based authentication metadata for AWS services.

_Required_: No

_Type_: [IAMConnectionMetadata](aws-properties-quicksight-actionconnector-iamconnectionmetadata.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoneConnectionMetadata`

No authentication metadata for services that don't require authentication.

_Required_: No

_Type_: [NoneConnectionMetadata](aws-properties-quicksight-actionconnector-noneconnectionmetadata.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuthConfig

AuthorizationCodeGrantCredentialsDetails

All content copied from https://docs.aws.amazon.com/.
