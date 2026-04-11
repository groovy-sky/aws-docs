This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KafkaConnect::Connector ScaleOutPolicy

The scale-out policy for the connector.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CpuUtilizationPercentage" : Integer
}

```

### YAML

```yaml

  CpuUtilizationPercentage: Integer

```

## Properties

`CpuUtilizationPercentage`

The CPU utilization percentage threshold at which you want connector scale out to be
triggered.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScaleInPolicy

Tag

All content copied from https://docs.aws.amazon.com/.
