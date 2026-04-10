This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::Cluster ManagedScalingPolicy

Managed scaling policy for an Amazon EMR cluster. The policy specifies the
limits for resources that can be added or terminated from a cluster. The policy only
applies to the core and task nodes. The master node cannot be scaled after initial
configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComputeLimits" : ComputeLimits,
  "ScalingStrategy" : String,
  "UtilizationPerformanceIndex" : Integer
}

```

### YAML

```yaml

  ComputeLimits:
    ComputeLimits
  ScalingStrategy: String
  UtilizationPerformanceIndex: Integer

```

## Properties

`ComputeLimits`

The Amazon EC2 unit limits for a managed scaling policy. The managed scaling
activity of a cluster is not allowed to go above or below these limits. The limit only
applies to the core and task nodes. The master node cannot be scaled after initial
configuration.

_Required_: No

_Type_: [ComputeLimits](aws-properties-emr-cluster-computelimits.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScalingStrategy`

Determines whether a custom scaling utilization performance index can be set. Possible values include _ADVANCED_ or _DEFAULT_.

_Required_: No

_Type_: String

_Allowed values_: `DEFAULT | ADVANCED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UtilizationPerformanceIndex`

An integer value that represents an advanced scaling strategy. Setting a higher value optimizes for performance. Setting a lower value optimizes for resource conservation. Setting the value
to 50 balances performance and resource conservation. Possible values are 1, 25, 50, 75, and 100.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KeyValue

MetricDimension

All content copied from https://docs.aws.amazon.com/.
