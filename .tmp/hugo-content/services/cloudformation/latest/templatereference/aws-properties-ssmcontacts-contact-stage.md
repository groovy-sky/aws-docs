This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSMContacts::Contact Stage

The `Stage` property type specifies a set amount of time that an escalation
plan or engagement plan engages the specified contacts or contact methods.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DurationInMinutes" : Integer,
  "RotationIds" : [ String, ... ],
  "Targets" : [ Targets, ... ]
}

```

### YAML

```yaml

  DurationInMinutes: Integer
  RotationIds:
    - String
  Targets:
    - Targets

```

## Properties

`DurationInMinutes`

The time to wait until beginning the next stage. The duration can only be set to 0 if a
target is specified.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RotationIds`

The Amazon Resource Names (ARNs) of the on-call rotations associated with the
plan.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Targets`

The contacts or contact methods that the escalation plan or engagement plan is
engaging.

_Required_: No

_Type_: [Array](aws-properties-ssmcontacts-contact-targets.md) of [Targets](aws-properties-ssmcontacts-contact-targets.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContactTargetInfo

Tag

All content copied from https://docs.aws.amazon.com/.
