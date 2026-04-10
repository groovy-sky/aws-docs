This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::AutomationRuleV2 ExternalIntegrationConfiguration

The settings for integrating automation rule actions with external systems or service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConnectorArn" : String
}

```

### YAML

```yaml

  ConnectorArn: String

```

## Properties

`ConnectorArn`

The ARN of the connector that establishes the integration.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DateRange

MapFilter

All content copied from https://docs.aws.amazon.com/.
