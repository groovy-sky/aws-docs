This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::FeatureGroup ThroughputConfig

Used to set feature group throughput configuration. There are two modes:
`ON_DEMAND` and `PROVISIONED`. With on-demand mode, you are
charged for data reads and writes that your application performs on your feature group. You
do not need to specify read and write throughput because Feature Store accommodates your
workloads as they ramp up and down. You can switch a feature group to on-demand only once
in a 24 hour period. With provisioned throughput mode, you specify the read and write
capacity per second that you expect your application to require, and you are billed based
on those limits. Exceeding provisioned throughput will result in your requests being
throttled.

Note: `PROVISIONED` throughput mode is supported only for feature groups that
are offline-only, or use the [`Standard`](../../../../reference/sagemaker/latest/apireference/api-onlinestoreconfig.md#sagemaker-Type-OnlineStoreConfig-StorageType) tier online store.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ProvisionedReadCapacityUnits" : Integer,
  "ProvisionedWriteCapacityUnits" : Integer,
  "ThroughputMode" : String
}

```

### YAML

```yaml

  ProvisionedReadCapacityUnits: Integer
  ProvisionedWriteCapacityUnits: Integer
  ThroughputMode: String

```

## Properties

`ProvisionedReadCapacityUnits`

For provisioned feature groups with online store enabled, this indicates the read
throughput you are billed for and can consume without throttling.

This field is not applicable for on-demand feature groups.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `10000000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProvisionedWriteCapacityUnits`

For provisioned feature groups, this indicates the write throughput you are billed for
and can consume without throttling.

This field is not applicable for on-demand feature groups.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `10000000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ThroughputMode`

The mode used for your feature group throughput: `ON_DEMAND` or
`PROVISIONED`.

_Required_: Yes

_Type_: String

_Allowed values_: `OnDemand | Provisioned`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

TtlDuration

All content copied from https://docs.aws.amazon.com/.
