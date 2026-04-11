This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::EventTrigger Period

Defines a limit and the time period during which it is enforced.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxInvocationsPerProfile" : Integer,
  "Unit" : String,
  "Unlimited" : Boolean,
  "Value" : Integer
}

```

### YAML

```yaml

  MaxInvocationsPerProfile: Integer
  Unit: String
  Unlimited: Boolean
  Value: Integer

```

## Properties

`MaxInvocationsPerProfile`

The maximum allowed number of destination invocations per profile.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Unit`

The unit of time.

_Required_: Yes

_Type_: String

_Allowed values_: `HOURS | DAYS | WEEKS | MONTHS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Unlimited`

If set to true, there is no limit on the number of destination invocations per
profile. The default is false.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The amount of time of the specified unit.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `24`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ObjectAttribute

Tag

All content copied from https://docs.aws.amazon.com/.
