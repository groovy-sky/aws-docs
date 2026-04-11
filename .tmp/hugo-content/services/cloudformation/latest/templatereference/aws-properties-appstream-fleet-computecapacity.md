This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppStream::Fleet ComputeCapacity

The desired capacity for a fleet.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DesiredInstances" : Integer,
  "DesiredSessions" : Integer
}

```

### YAML

```yaml

  DesiredInstances: Integer
  DesiredSessions: Integer

```

## Properties

`DesiredInstances`

The desired number of streaming instances.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DesiredSessions`

The desired capacity in terms of number of user sessions, for the multi-session fleet.
This is not allowed for single-session fleets.

When you create a fleet, you must set define either the DesiredSessions or
DesiredInstances attribute, based on the type of fleet you create. You can’t define both
attributes or leave both attributes blank.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::AppStream::Fleet

DomainJoinInfo

All content copied from https://docs.aws.amazon.com/.
