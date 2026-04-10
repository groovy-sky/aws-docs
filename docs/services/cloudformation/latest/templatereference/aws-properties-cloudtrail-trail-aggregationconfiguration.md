This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudTrail::Trail AggregationConfiguration

An object that contains configuration settings for aggregating events.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EventCategory" : String,
  "Templates" : [ String, ... ]
}

```

### YAML

```yaml

  EventCategory: String
  Templates:
    - String

```

## Properties

`EventCategory`

Specifies the event category for which aggregation should be performed.

_Required_: Yes

_Type_: String

_Allowed values_: `Data`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Templates`

A list of aggregation templates that can be used to configure event aggregation.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AdvancedFieldSelector

DataResource

All content copied from https://docs.aws.amazon.com/.
