This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Topic TopicCalculatedField

A structure that represents a calculated field.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Aggregation" : String,
  "AllowedAggregations" : [ String, ... ],
  "CalculatedFieldDescription" : String,
  "CalculatedFieldName" : String,
  "CalculatedFieldSynonyms" : [ String, ... ],
  "CellValueSynonyms" : [ CellValueSynonym, ... ],
  "ColumnDataRole" : String,
  "ComparativeOrder" : ComparativeOrder,
  "DefaultFormatting" : DefaultFormatting,
  "DisableIndexing" : Boolean,
  "Expression" : String,
  "IsIncludedInTopic" : Boolean,
  "NeverAggregateInFilter" : Boolean,
  "NonAdditive" : Boolean,
  "NotAllowedAggregations" : [ String, ... ],
  "SemanticType" : SemanticType,
  "TimeGranularity" : String
}

```

### YAML

```yaml

  Aggregation: String
  AllowedAggregations:
    - String
  CalculatedFieldDescription: String
  CalculatedFieldName: String
  CalculatedFieldSynonyms:
    - String
  CellValueSynonyms:
    - CellValueSynonym
  ColumnDataRole: String
  ComparativeOrder:
    ComparativeOrder
  DefaultFormatting:
    DefaultFormatting
  DisableIndexing: Boolean
  Expression: String
  IsIncludedInTopic: Boolean
  NeverAggregateInFilter: Boolean
  NonAdditive: Boolean
  NotAllowedAggregations:
    - String
  SemanticType:
    SemanticType
  TimeGranularity: String

```

## Properties

`Aggregation`

The default aggregation. Valid values for this structure are `SUM`,
`MAX`, `MIN`, `COUNT`,
`DISTINCT_COUNT`,
and `AVERAGE`.

_Required_: No

_Type_: String

_Allowed values_: `SUM | MAX | MIN | COUNT | DISTINCT_COUNT | AVERAGE | MEDIAN | STDEV | STDEVP | VAR | VARP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowedAggregations`

The list of aggregation types that are allowed for the calculated field. Valid values
for this structure are `COUNT`, `DISTINCT_COUNT`, `MIN`,
`MAX`, `MEDIAN`, `SUM`, `AVERAGE`,
`STDEV`, `STDEVP`, `VAR`,
`VARP`, and `PERCENTILE`.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CalculatedFieldDescription`

The calculated field description.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CalculatedFieldName`

The calculated field name.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CalculatedFieldSynonyms`

The other names or aliases for the calculated field.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CellValueSynonyms`

The other
names or aliases for the calculated field cell value.

_Required_: No

_Type_: Array of [CellValueSynonym](aws-properties-quicksight-topic-cellvaluesynonym.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColumnDataRole`

The column data role for a calculated field. Valid values for this structure are `DIMENSION` and `MEASURE`.

_Required_: No

_Type_: String

_Allowed values_: `DIMENSION | MEASURE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComparativeOrder`

The order in which data is displayed for the calculated field when
it's used in a comparative context.

_Required_: No

_Type_: [ComparativeOrder](aws-properties-quicksight-topic-comparativeorder.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultFormatting`

The default formatting definition.

_Required_: No

_Type_: [DefaultFormatting](aws-properties-quicksight-topic-defaultformatting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisableIndexing`

A Boolean value that indicates if a calculated field is visible in the autocomplete.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Expression`

The calculated field expression.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsIncludedInTopic`

A boolean value that indicates if a calculated field is included in the topic.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NeverAggregateInFilter`

A Boolean value that indicates whether to never aggregate calculated field in filters.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NonAdditive`

The non additive for the table style target.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotAllowedAggregations`

The list of aggregation types that are not allowed for the calculated field. Valid
values for this structure are `COUNT`, `DISTINCT_COUNT`,
`MIN`, `MAX`, `MEDIAN`, `SUM`,
`AVERAGE`, `STDEV`, `STDEVP`, `VAR`,
`VARP`, and `PERCENTILE`.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SemanticType`

The semantic type.

_Required_: No

_Type_: [SemanticType](aws-properties-quicksight-topic-semantictype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeGranularity`

The level of time precision that is used to aggregate `DateTime` values.

_Required_: No

_Type_: String

_Allowed values_: `SECOND | MINUTE | HOUR | DAY | WEEK | MONTH | QUARTER | YEAR`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

TopicCategoryFilter

All content copied from https://docs.aws.amazon.com/.
