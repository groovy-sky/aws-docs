This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaignsV2::Campaign RestrictedPeriod

Contains information about a restricted period.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EndDate" : String,
  "Name" : String,
  "StartDate" : String
}

```

### YAML

```yaml

  EndDate: String
  Name: String
  StartDate: String

```

## Properties

`EndDate`

The end date of the restricted period.

_Required_: Yes

_Type_: String

_Pattern_: `^\d{4}-\d{2}-\d{2}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the restricted period.

_Required_: No

_Type_: String

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartDate`

The start date of the restricted period.

_Required_: Yes

_Type_: String

_Pattern_: `^\d{4}-\d{2}-\d{2}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProgressiveConfig

RestrictedPeriods

All content copied from https://docs.aws.amazon.com/.
