This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DLM::LifecyclePolicy ShareRule

**\[Custom snapshot policies only\]** Specifies a rule for sharing snapshots across AWS accounts.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TargetAccounts" : [ String, ... ],
  "UnshareInterval" : Integer,
  "UnshareIntervalUnit" : String
}

```

### YAML

```yaml

  TargetAccounts:
    - String
  UnshareInterval: Integer
  UnshareIntervalUnit: String

```

## Properties

`TargetAccounts`

The IDs of the AWS accounts with which to share the snapshots.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UnshareInterval`

The period after which snapshots that are shared with other AWS accounts are automatically unshared.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UnshareIntervalUnit`

The unit of time for the automatic unsharing interval.

_Required_: No

_Type_: String

_Allowed values_: `DAYS | WEEKS | MONTHS | YEARS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Script

Tag

All content copied from https://docs.aws.amazon.com/.
