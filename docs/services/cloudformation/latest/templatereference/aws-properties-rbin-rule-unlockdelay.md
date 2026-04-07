This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Rbin::Rule UnlockDelay

Information about the retention rule unlock delay. The unlock delay is the period after which
a retention rule can be modified or edited after it has been unlocked by a user with the required
permissions. The retention rule can't be modified or deleted during the unlock delay.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "UnlockDelayUnit" : String,
  "UnlockDelayValue" : Integer
}

```

### YAML

```yaml

  UnlockDelayUnit: String
  UnlockDelayValue: Integer

```

## Properties

`UnlockDelayUnit`

The unit of time in which to measure the unlock delay. Currently, the unlock delay can
be measured only in days.

_Required_: No

_Type_: String

_Allowed values_: `DAYS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UnlockDelayValue`

The unlock delay period, measured in the unit specified for **UnlockDelayUnit**.

_Required_: No

_Type_: Integer

_Minimum_: `7`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Next
