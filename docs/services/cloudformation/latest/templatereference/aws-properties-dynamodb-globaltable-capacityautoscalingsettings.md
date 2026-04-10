This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::GlobalTable CapacityAutoScalingSettings

Configures a scalable target and an autoscaling policy for a table or global secondary
index's read or write capacity.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxCapacity" : Integer,
  "MinCapacity" : Integer,
  "SeedCapacity" : Integer,
  "TargetTrackingScalingPolicyConfiguration" : TargetTrackingScalingPolicyConfiguration
}

```

### YAML

```yaml

  MaxCapacity: Integer
  MinCapacity: Integer
  SeedCapacity: Integer
  TargetTrackingScalingPolicyConfiguration:
    TargetTrackingScalingPolicyConfiguration

```

## Properties

`MaxCapacity`

The maximum provisioned capacity units for the global table.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinCapacity`

The minimum provisioned capacity units for the global table.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SeedCapacity`

When switching billing mode from `PAY_PER_REQUEST` to
`PROVISIONED`, DynamoDB requires you to specify read and write capacity unit
values for the table and for each global secondary index. These values will be applied to
all replicas. The table will use these provisioned values until CloudFormation creates the
autoscaling policies you configured in your template. CloudFormation cannot determine what
capacity the table and its global secondary indexes will require in this time period, since
they are application-dependent.

If you want to switch a table's billing mode from `PAY_PER_REQUEST` to
`PROVISIONED`, you must specify a value for this property for each autoscaled
resource. If you specify different values for the same resource in different regions,
CloudFormation will use the highest value found in either the `SeedCapacity` or
`ReadCapacityUnits` properties. For example, if your global secondary index
`myGSI` has a `SeedCapacity` of 10 in us-east-1 and a fixed
`ReadCapacityUnits` of 20 in eu-west-1, CloudFormation will initially set the
read capacity for `myGSI` to 20. Note that if you disable `ScaleIn`
for `myGSI` in us-east-1, its read capacity units might not be set back to
10.

You must also specify a value for `SeedCapacity` when you plan to switch a
table's billing mode from `PROVISIONED` to `PAY_PER_REQUEST`, because
CloudFormation might need to roll back the operation (reverting the billing mode to
`PROVISIONED`) and this cannot succeed without specifying a value for
`SeedCapacity`.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetTrackingScalingPolicyConfiguration`

Defines a target tracking scaling policy.

_Required_: Yes

_Type_: [TargetTrackingScalingPolicyConfiguration](aws-properties-dynamodb-globaltable-targettrackingscalingpolicyconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AttributeDefinition

ContributorInsightsSpecification

All content copied from https://docs.aws.amazon.com/.
