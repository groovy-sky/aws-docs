This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe MultiMeasureMapping

Maps multiple measures from the source event to the same Timestream for
LiveAnalytics record.

For more information, see [Amazon Timestream for LiveAnalytics concepts](../../../timestream/latest/developerguide/concepts.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MultiMeasureAttributeMappings" : [ MultiMeasureAttributeMapping, ... ],
  "MultiMeasureName" : String
}

```

### YAML

```yaml

  MultiMeasureAttributeMappings:
    - MultiMeasureAttributeMapping
  MultiMeasureName: String

```

## Properties

`MultiMeasureAttributeMappings`

Mappings that represent multiple source event fields mapped to measures in the same
Timestream for LiveAnalytics record.

_Required_: Yes

_Type_: Array of [MultiMeasureAttributeMapping](aws-properties-pipes-pipe-multimeasureattributemapping.md)

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MultiMeasureName`

The name of the multiple measurements per record (multi-measure).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MultiMeasureAttributeMapping

NetworkConfiguration

All content copied from https://docs.aws.amazon.com/.
