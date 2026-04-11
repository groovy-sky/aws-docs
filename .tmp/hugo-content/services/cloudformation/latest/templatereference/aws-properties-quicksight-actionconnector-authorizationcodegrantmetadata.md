This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::ActionConnector AuthorizationCodeGrantMetadata

Metadata for OAuth 2.0 authorization code grant authentication.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizationCodeGrantCredentialsDetails" : AuthorizationCodeGrantCredentialsDetails,
  "AuthorizationCodeGrantCredentialsSource" : String,
  "BaseEndpoint" : String,
  "RedirectUrl" : String
}

```

### YAML

```yaml

  AuthorizationCodeGrantCredentialsDetails:
    AuthorizationCodeGrantCredentialsDetails
  AuthorizationCodeGrantCredentialsSource: String
  BaseEndpoint: String
  RedirectUrl: String

```

## Properties

`AuthorizationCodeGrantCredentialsDetails`

The detailed credentials configuration for authorization code grant.

_Required_: No

_Type_: [AuthorizationCodeGrantCredentialsDetails](aws-properties-quicksight-actionconnector-authorizationcodegrantcredentialsdetails.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthorizationCodeGrantCredentialsSource`

The source of the authorization code grant credentials.

_Required_: No

_Type_: String

_Allowed values_: `PLAIN_CREDENTIALS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BaseEndpoint`

The base URL endpoint for the external service.

_Required_: Yes

_Type_: String

_Pattern_: `^https://.*`

_Minimum_: `1`

_Maximum_: `8192`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RedirectUrl`

The redirect URL for the OAuth authorization flow.

_Required_: Yes

_Type_: String

_Pattern_: `^https://.*`

_Minimum_: `1`

_Maximum_: `8192`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AuthorizationCodeGrantDetails

BasicAuthConnectionMetadata

All content copied from https://docs.aws.amazon.com/.
