This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::ActionConnector AuthConfig

Authentication configuration for connecting to external services.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthenticationMetadata" : AuthenticationMetadata,
  "AuthenticationType" : String
}

```

### YAML

```yaml

  AuthenticationMetadata:
    AuthenticationMetadata
  AuthenticationType: String

```

## Properties

`AuthenticationMetadata`

The authentication metadata containing the specific configuration for the chosen authentication type.

_Required_: Yes

_Type_: [AuthenticationMetadata](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-actionconnector-authenticationmetadata.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthenticationType`

The type of authentication method.

_Required_: Yes

_Type_: String

_Allowed values_: `BASIC | API_KEY | OAUTH2_CLIENT_CREDENTIALS | NONE | IAM | OAUTH2_AUTHORIZATION_CODE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

APIKeyConnectionMetadata

AuthenticationMetadata
