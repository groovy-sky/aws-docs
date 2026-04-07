This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile CustomConnectorProfileProperties

The profile properties required by the custom connector.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OAuth2Properties" : OAuth2Properties,
  "ProfileProperties" : {Key: Value, ...}
}

```

### YAML

```yaml

  OAuth2Properties:
    OAuth2Properties
  ProfileProperties:
    Key: Value

```

## Properties

`OAuth2Properties`

The OAuth 2.0 properties required for OAuth 2.0 authentication.

_Required_: No

_Type_: [OAuth2Properties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-connectorprofile-oauth2properties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProfileProperties`

A map of properties that are required to create a profile for the custom connector.

_Required_: No

_Type_: Object of String

_Pattern_: `^[\w]{1,256}$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CustomConnectorProfileCredentials

DatadogConnectorProfileCredentials
