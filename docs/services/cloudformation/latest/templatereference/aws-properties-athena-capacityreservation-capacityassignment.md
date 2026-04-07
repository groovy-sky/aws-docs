This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Athena::CapacityReservation CapacityAssignment

A mapping between one or more workgroups and a capacity reservation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "WorkgroupNames" : [ String, ... ]
}

```

### YAML

```yaml

  WorkgroupNames:
    - String

```

## Properties

`WorkgroupNames`

The list of workgroup names for the capacity assignment.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Athena::CapacityReservation

CapacityAssignmentConfiguration
