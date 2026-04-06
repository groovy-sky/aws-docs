This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile ConnectorProfileConfig

Defines the connector-specific configuration and credentials for the connector profile.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConnectorProfileCredentials" : ConnectorProfileCredentials,
  "ConnectorProfileProperties" : ConnectorProfileProperties
}

```

### YAML

```yaml

  ConnectorProfileCredentials:
    ConnectorProfileCredentials
  ConnectorProfileProperties:
    ConnectorProfileProperties

```

## Properties

`ConnectorProfileCredentials`

The connector-specific credentials required by each connector.

_Required_: No

_Type_: [ConnectorProfileCredentials](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-connectorprofile-connectorprofilecredentials.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectorProfileProperties`

The connector-specific properties of the profile configuration.

_Required_: No

_Type_: [ConnectorProfileProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-connectorprofile-connectorprofileproperties.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ConnectorProfileConfig](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_ConnectorProfileConfig.html) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ConnectorOAuthRequest

ConnectorProfileCredentials
