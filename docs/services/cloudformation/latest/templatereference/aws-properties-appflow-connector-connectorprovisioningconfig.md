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

_Type_: [LambdaConnectorProvisioningConfig](aws-properties-appflow-connector-lambdaconnectorprovisioningconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::AppFlow::Connector

LambdaConnectorProvisioningConfig

All content copied from https://docs.aws.amazon.com/.
