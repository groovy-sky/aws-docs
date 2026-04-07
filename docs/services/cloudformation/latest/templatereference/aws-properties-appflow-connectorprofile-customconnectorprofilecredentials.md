This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile CustomConnectorProfileCredentials

The connector-specific profile credentials that are required when using the custom
connector.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApiKey" : ApiKeyCredentials,
  "AuthenticationType" : String,
  "Basic" : BasicAuthCredentials,
  "Custom" : CustomAuthCredentials,
  "Oauth2" : OAuth2Credentials
}

```

### YAML

```yaml

  ApiKey:
    ApiKeyCredentials
  AuthenticationType: String
  Basic:
    BasicAuthCredentials
  Custom:
    CustomAuthCredentials
  Oauth2:
    OAuth2Credentials

```

## Properties

`ApiKey`

The API keys required for the authentication of the user.

_Required_: No

_Type_: [ApiKeyCredentials](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-connectorprofile-apikeycredentials.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthenticationType`

The authentication type that the custom connector uses for authenticating while creating a
connector profile.

_Required_: Yes

_Type_: String

_Allowed values_: `OAUTH2 | APIKEY | BASIC | CUSTOM`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Basic`

The basic credentials that are required for the authentication of the user.

_Required_: No

_Type_: [BasicAuthCredentials](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-connectorprofile-basicauthcredentials.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Custom`

If the connector uses the custom authentication mechanism, this holds the required
credentials.

_Required_: No

_Type_: [CustomAuthCredentials](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-connectorprofile-customauthcredentials.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Oauth2`

The OAuth 2.0 credentials required for the authentication of the user.

_Required_: No

_Type_: [OAuth2Credentials](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-connectorprofile-oauth2credentials.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CustomAuthCredentials

CustomConnectorProfileProperties
