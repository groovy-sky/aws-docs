This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cassandra::Table AutoScalingSpecification

The optional auto scaling capacity settings for a table in provisioned capacity mode.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ReadCapacityAutoScaling" : AutoScalingSetting,
  "WriteCapacityAutoScaling" : AutoScalingSetting
}

```

### YAML

```yaml

  ReadCapacityAutoScaling:
    AutoScalingSetting
  WriteCapacityAutoScaling:
    AutoScalingSetting

```

## Properties

`ReadCapacityAutoScaling`

The auto scaling settings for the table's read capacity.

_Required_: No

_Type_: [AutoScalingSetting](aws-properties-cassandra-table-autoscalingsetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WriteCapacityAutoScaling`

The auto scaling settings for the table's write capacity.

_Required_: No

_Type_: [AutoScalingSetting](aws-properties-cassandra-table-autoscalingsetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AutoScalingSetting

BillingMode

All content copied from https://docs.aws.amazon.com/.
