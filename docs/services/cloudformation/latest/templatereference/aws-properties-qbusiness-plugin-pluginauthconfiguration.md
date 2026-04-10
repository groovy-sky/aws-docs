This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::Plugin PluginAuthConfiguration

Authentication configuration information for an Amazon Q Business plugin.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BasicAuthConfiguration" : BasicAuthConfiguration,
  "NoAuthConfiguration" : Json,
  "OAuth2ClientCredentialConfiguration" : OAuth2ClientCredentialConfiguration
}

```

### YAML

```yaml

  BasicAuthConfiguration:
    BasicAuthConfiguration
  NoAuthConfiguration: Json
  OAuth2ClientCredentialConfiguration:
    OAuth2ClientCredentialConfiguration

```

## Properties

`BasicAuthConfiguration`

Information about the basic authentication credentials used to configure a
plugin.

_Required_: No

_Type_: [BasicAuthConfiguration](aws-properties-qbusiness-plugin-basicauthconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NoAuthConfiguration`

Information about invoking a custom plugin without any authentication.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OAuth2ClientCredentialConfiguration`

Information about the OAuth 2.0 authentication credential/token used to configure a
plugin.

_Required_: No

_Type_: [OAuth2ClientCredentialConfiguration](aws-properties-qbusiness-plugin-oauth2clientcredentialconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OAuth2ClientCredentialConfiguration

S3

All content copied from https://docs.aws.amazon.com/.
