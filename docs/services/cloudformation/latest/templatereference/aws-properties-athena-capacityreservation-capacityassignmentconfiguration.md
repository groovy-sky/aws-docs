This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Athena::CapacityReservation CapacityAssignmentConfiguration

Assigns Athena workgroups (and hence their queries) to capacity
reservations. A capacity reservation can have only one capacity assignment
configuration, but the capacity assignment configuration can be made up of multiple
individual assignments. Each assignment specifies how Athena queries can
consume capacity from the capacity reservation that their workgroup is mapped to.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CapacityAssignments" : [ CapacityAssignment, ... ]
}

```

### YAML

```yaml

  CapacityAssignments:
    - CapacityAssignment

```

## Properties

`CapacityAssignments`

The list of assignments that make up the capacity assignment configuration.

_Required_: Yes

_Type_: Array of [CapacityAssignment](aws-properties-athena-capacityreservation-capacityassignment.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CapacityAssignment

Tag

All content copied from https://docs.aws.amazon.com/.
