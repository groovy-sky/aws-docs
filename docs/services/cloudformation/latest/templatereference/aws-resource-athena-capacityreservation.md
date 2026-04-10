This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Athena::CapacityReservation

Specifies a capacity reservation with the provided name and number of requested data
processing units.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Athena::CapacityReservation",
  "Properties" : {
      "CapacityAssignmentConfiguration" : CapacityAssignmentConfiguration,
      "Name" : String,
      "Tags" : [ Tag, ... ],
      "TargetDpus" : Integer
    }
}

```

### YAML

```yaml

Type: AWS::Athena::CapacityReservation
Properties:
  CapacityAssignmentConfiguration:
    CapacityAssignmentConfiguration
  Name: String
  Tags:
    - Tag
  TargetDpus: Integer

```

## Properties

`CapacityAssignmentConfiguration`

Assigns Athena workgroups (and hence their queries) to capacity reservations. A
capacity reservation can have only one capacity assignment configuration, but the
capacity assignment configuration can be made up of multiple individual assignments.
Each assignment specifies how Athena queries can consume capacity from the capacity
reservation that their workgroup is mapped to.

_Required_: No

_Type_: [CapacityAssignmentConfiguration](aws-properties-athena-capacityreservation-capacityassignmentconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the capacity reservation.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9._-]{1,128}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to the capacity reservation.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-athena-capacityreservation-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetDpus`

The number of data processing units requested.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the capacity reservation.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AllocatedDpus`

The number of data processing units currently allocated.

`Arn`

The ARN of the capacity reservation.

`CreationTime`

The time in UTC epoch millis when the capacity reservation was created.

`LastSuccessfulAllocationTime`

The time of the most recent capacity allocation that succeeded.

`Status`

The status of the capacity reservation.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Athena

CapacityAssignment

All content copied from https://docs.aws.amazon.com/.
