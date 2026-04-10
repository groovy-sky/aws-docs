This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppIntegrations::DataIntegration ScheduleConfig

The name of the data and how often it should be pulled from the source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FirstExecutionFrom" : String,
  "Object" : String,
  "ScheduleExpression" : String
}

```

### YAML

```yaml

  FirstExecutionFrom: String
  Object: String
  ScheduleExpression: String

```

## Properties

`FirstExecutionFrom`

The start date for objects to import in the first flow run as an Unix/epoch timestamp in
milliseconds or in ISO-8601 format.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Object`

The name of the object to pull from the data source.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9/\._\-]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ScheduleExpression`

How often the data should be pulled from data source.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FileConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
