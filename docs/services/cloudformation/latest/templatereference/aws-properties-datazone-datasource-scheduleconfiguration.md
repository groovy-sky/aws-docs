This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::DataSource ScheduleConfiguration

The details of the schedule of the data source runs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Schedule" : String,
  "Timezone" : String
}

```

### YAML

```yaml

  Schedule: String
  Timezone: String

```

## Properties

`Schedule`

The schedule of the data source runs.

_Required_: No

_Type_: String

_Pattern_: `cron\((\b[0-5]?[0-9]\b) (\b2[0-3]\b|\b[0-1]?[0-9]\b) (.*){1,5} (.*){1,5} (.*){1,5} (.*){1,5}\)`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Timezone`

The timezone of the data source run.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SageMakerRunConfigurationInput

AWS::DataZone::Domain

All content copied from https://docs.aws.amazon.com/.
