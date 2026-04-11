This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Topic NamedEntityDefinition

A structure that represents a named entity.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldName" : String,
  "Metric" : NamedEntityDefinitionMetric,
  "PropertyName" : String,
  "PropertyRole" : String,
  "PropertyUsage" : String
}

```

### YAML

```yaml

  FieldName: String
  Metric:
    NamedEntityDefinitionMetric
  PropertyName: String
  PropertyRole: String
  PropertyUsage: String

```

## Properties

`FieldName`

The name of the entity.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Metric`

The definition of a metric.

_Required_: No

_Type_: [NamedEntityDefinitionMetric](aws-properties-quicksight-topic-namedentitydefinitionmetric.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropertyName`

The property name to be used for the named entity.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropertyRole`

The property role. Valid values for this structure are `PRIMARY` and `ID`.

_Required_: No

_Type_: String

_Allowed values_: `PRIMARY | ID`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PropertyUsage`

The property usage. Valid values for this structure are `INHERIT`,
`DIMENSION`,
and `MEASURE`.

_Required_: No

_Type_: String

_Allowed values_: `INHERIT | DIMENSION | MEASURE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DisplayFormatOptions

NamedEntityDefinitionMetric

All content copied from https://docs.aws.amazon.com/.
