This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cassandra::Table ScalingPolicy

Amazon Keyspaces supports the `target tracking` auto scaling policy. With this policy, Amazon Keyspaces auto scaling
ensures that the table's ratio of consumed to provisioned capacity stays at or near the target value that you specify. You
define the target value as a percentage between 20 and 90.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TargetTrackingScalingPolicyConfiguration" : TargetTrackingScalingPolicyConfiguration
}

```

### YAML

```yaml

  TargetTrackingScalingPolicyConfiguration:
    TargetTrackingScalingPolicyConfiguration

```

## Properties

`TargetTrackingScalingPolicyConfiguration`

The auto scaling policy that scales a table based on the ratio of consumed to provisioned capacity.

_Required_: No

_Type_: [TargetTrackingScalingPolicyConfiguration](aws-properties-cassandra-table-targettrackingscalingpolicyconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicaSpecification

Tag

All content copied from https://docs.aws.amazon.com/.
