This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationSignals::GroupingConfiguration GroupingAttributeDefinition

A structure that defines how services should be grouped based on specific attributes. This includes the friendly name for the grouping, the source keys to derive values from, and an optional default value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultGroupingValue" : String,
  "GroupingName" : String,
  "GroupingSourceKeys" : [ String, ... ]
}

```

### YAML

```yaml

  DefaultGroupingValue: String
  GroupingName: String
  GroupingSourceKeys:
    - String

```

## Properties

`DefaultGroupingValue`

The default value to use for this grouping attribute when no value can be derived from the source keys. This ensures all services have a grouping value even if the source data is missing.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GroupingName`

The friendly name for this grouping attribute, such as `BusinessUnit` or `Environment`. This name is used to identify the grouping in the console and APIs.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GroupingSourceKeys`

An array of source keys used to derive the grouping attribute value from telemetry data, AWS tags, or other sources. For example, \["business\_unit", "team"\] would look for values in those fields.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ApplicationSignals::GroupingConfiguration

AWS::ApplicationSignals::ServiceLevelObjective
