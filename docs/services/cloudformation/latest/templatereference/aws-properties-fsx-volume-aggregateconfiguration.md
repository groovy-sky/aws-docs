This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::Volume AggregateConfiguration

Use to specify configuration options for a volume’s storage aggregate or aggregates.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Aggregates" : [ String, ... ],
  "ConstituentsPerAggregate" : Integer
}

```

### YAML

```yaml

  Aggregates:
    - String
  ConstituentsPerAggregate: Integer

```

## Properties

`Aggregates`

The list of aggregates that this volume resides on. Aggregates are storage pools which make up your primary storage tier. Each high-availability (HA) pair has one aggregate. The names of the aggregates map to the names of the aggregates in the ONTAP CLI and REST API. For FlexVols, there will always be a single entry.

Amazon FSx responds with an HTTP status code 400 (Bad Request) for the following conditions:

- The strings in the value of `Aggregates` are not are not formatted as `aggrX`, where X is a number between 1 and 12.

- The value of `Aggregates` contains aggregates that are not present.

- One or more of the aggregates supplied are too close to the volume limit to support adding more volumes.

_Required_: No

_Type_: Array of String

_Maximum_: `6`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConstituentsPerAggregate`

Used to explicitly set the number of constituents within the FlexGroup per storage aggregate. This field is optional when creating a FlexGroup volume. If unspecified, the default value will be 8. This field cannot be provided when creating a FlexVol volume.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::FSx::Volume

AutocommitPeriod

All content copied from https://docs.aws.amazon.com/.
