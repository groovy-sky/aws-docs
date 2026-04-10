This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Timestream::ScheduledQuery MultiMeasureAttributeMapping

Attribute mapping for MULTI value measures.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MeasureValueType" : String,
  "SourceColumn" : String,
  "TargetMultiMeasureAttributeName" : String
}

```

### YAML

```yaml

  MeasureValueType: String
  SourceColumn: String
  TargetMultiMeasureAttributeName: String

```

## Properties

`MeasureValueType`

Type of the attribute to be read from the source column.

_Required_: Yes

_Type_: String

_Allowed values_: `BIGINT | BOOLEAN | DOUBLE | VARCHAR | TIMESTAMP`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceColumn`

Source column from where the attribute value is to be read.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TargetMultiMeasureAttributeName`

Custom name to be used for attribute name in derived table. If not provided, source column
name would be used.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MixedMeasureMapping

MultiMeasureMappings

All content copied from https://docs.aws.amazon.com/.
