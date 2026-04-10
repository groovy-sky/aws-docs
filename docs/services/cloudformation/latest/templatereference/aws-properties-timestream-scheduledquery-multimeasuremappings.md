This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Timestream::ScheduledQuery MultiMeasureMappings

Only one of MixedMeasureMappings or MultiMeasureMappings is to be provided.
MultiMeasureMappings can be used to ingest data as multi measures in the derived table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MultiMeasureAttributeMappings" : [ MultiMeasureAttributeMapping, ... ],
  "TargetMultiMeasureName" : String
}

```

### YAML

```yaml

  MultiMeasureAttributeMappings:
    - MultiMeasureAttributeMapping
  TargetMultiMeasureName: String

```

## Properties

`MultiMeasureAttributeMappings`

Required. Attribute mappings to be used for mapping query results to ingest data for
multi-measure attributes.

_Required_: Yes

_Type_: Array of [MultiMeasureAttributeMapping](aws-properties-timestream-scheduledquery-multimeasureattributemapping.md)

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TargetMultiMeasureName`

The name of the target multi-measure name in the derived table. This input is required
when measureNameColumn is not provided. If MeasureNameColumn is provided, then value from that
column will be used as multi-measure name.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MultiMeasureAttributeMapping

NotificationConfiguration

All content copied from https://docs.aws.amazon.com/.
