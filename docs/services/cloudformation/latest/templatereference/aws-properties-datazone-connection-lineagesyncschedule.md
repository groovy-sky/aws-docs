This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::Connection LineageSyncSchedule

The lineage sync schedule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Schedule" : String
}

```

### YAML

```yaml

  Schedule: String

```

## Properties

`Schedule`

The lineage sync schedule.

_Required_: No

_Type_: String

_Pattern_: `^cron\((\b[0-5]?[0-9]\b) (\b2[0-3]\b|\b[0-1]?[0-9]\b) ([-?*,/\dLW]){1,83} ([-*,/\d]|[a-zA-Z]{3}){1,23} ([-?#*,/\dL]|[a-zA-Z]{3}){1,13} ([^\)]+)\)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IamPropertiesInput

MlflowPropertiesInput

All content copied from https://docs.aws.amazon.com/.
