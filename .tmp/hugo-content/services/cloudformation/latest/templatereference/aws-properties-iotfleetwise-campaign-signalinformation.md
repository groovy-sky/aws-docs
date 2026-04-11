This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::Campaign SignalInformation

Information about a signal.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataPartitionId" : String,
  "MaxSampleCount" : Number,
  "MinimumSamplingIntervalMs" : Number,
  "Name" : String
}

```

### YAML

```yaml

  DataPartitionId: String
  MaxSampleCount: Number
  MinimumSamplingIntervalMs: Number
  Name: String

```

## Properties

`DataPartitionId`

The ID of the data partition this signal is associated with.

The ID must match one of the IDs provided in `dataPartitions`. This is
accomplished either by specifying a particular data partition ID or by using
`default` for an established default partition. You can establish a
default partition in the `DataPartition` data type.

###### Note

If you upload a signal as a condition for a campaign's data partition,
the same signal must be included in `signalsToCollect`.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxSampleCount`

The maximum number of samples to collect.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `4294967295`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MinimumSamplingIntervalMs`

The minimum duration of time (in milliseconds) between two triggering events to
collect data.

###### Note

If a signal changes often, you might want to collect data at a slower rate.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `4294967295`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the signal.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w|*|-]+(\.[\w|*|-]+)*$`

_Minimum_: `1`

_Maximum_: `150`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SignalFetchInformation

StorageMaximumSize

All content copied from https://docs.aws.amazon.com/.
