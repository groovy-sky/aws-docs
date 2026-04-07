This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::InspectorV2::CisScanConfiguration CisTargets

The CIS targets.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccountIds" : [ String, ... ],
  "TargetResourceTags" : [ String, ... ]
}

```

### YAML

```yaml

  AccountIds:
    - String
  TargetResourceTags:
    - String

```

## Properties

`AccountIds`

The CIS target account ids.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetResourceTags`

The CIS target resource tags.

_Required_: Yes

_Type_: Array of String

_Pattern_: `^.+$`

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::InspectorV2::CisScanConfiguration

DailySchedule
