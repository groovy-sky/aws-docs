This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Connector ConnectorProvisioningConfig

Contains information about the configuration of the connector being registered.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Lambda" : LambdaConnectorProvisioningConfig
}

```

### YAML

```yaml

  Lambda:
    LambdaConnectorProvisioningConfig

```

## Properties

`Lambda`

Contains information about the configuration of the lambda which is being registered as
the connector.

_Required_: No

_Type_: [LambdaConnectorProvisioningConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-connector-lambdaconnectorprovisioningconfig.html)

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::AppFlow::Connector

LambdaConnectorProvisioningConfig
