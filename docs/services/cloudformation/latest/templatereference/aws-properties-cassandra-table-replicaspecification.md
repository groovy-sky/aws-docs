This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cassandra::Table ReplicaSpecification

The AWS Region specific settings of a multi-Region table.

For a multi-Region table, you can configure the table's read capacity differently per AWS Region. You can do this by configuring the following parameters.

- `region`: The Region where these settings are applied. (Required)

- `readCapacityUnits`: The provisioned read capacity units. (Optional)

- `readCapacityAutoScaling`: The read capacity auto scaling settings for the table. (Optional)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ReadCapacityAutoScaling" : AutoScalingSetting,
  "ReadCapacityUnits" : Integer,
  "Region" : String
}

```

### YAML

```yaml

  ReadCapacityAutoScaling:
    AutoScalingSetting
  ReadCapacityUnits: Integer
  Region: String

```

## Properties

`ReadCapacityAutoScaling`

The read capacity auto scaling settings for the multi-Region
table in the specified AWS Region.

_Required_: No

_Type_: [AutoScalingSetting](aws-properties-cassandra-table-autoscalingsetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReadCapacityUnits`

The provisioned read capacity units for the multi-Region table in the specified AWS Region.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Region`

The AWS Region.

_Required_: Yes

_Type_: String

_Minimum_: `2`

_Maximum_: `25`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProvisionedThroughput

ScalingPolicy

All content copied from https://docs.aws.amazon.com/.
