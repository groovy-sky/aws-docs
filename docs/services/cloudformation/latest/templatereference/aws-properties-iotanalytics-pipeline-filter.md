This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Pipeline Filter

An activity that filters a message based on its attributes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Filter" : String,
  "Name" : String,
  "Next" : String
}

```

### YAML

```yaml

  Filter: String
  Name: String
  Next: String

```

## Properties

`Filter`

An expression that looks like an SQL WHERE clause that must return a Boolean value.

_Required_: Yes

_Type_: [String](aws-properties-iotanalytics-pipeline-filter.md)

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the 'filter' activity.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Next`

The next activity in the pipeline.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeviceShadowEnrich

Lambda

All content copied from https://docs.aws.amazon.com/.
