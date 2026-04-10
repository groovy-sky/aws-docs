This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Topic TopicColumn

Represents a column in a dataset.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Aggregation" : String,
  "AllowedAggregations" : [ String, ... ],
  "CellValueSynonyms" : [ CellValueSynonym, ... ],
  "ColumnDataRole" : String,
  "ColumnDescription" : String,
  "ColumnFriendlyName" : String,
  "ColumnName" : String,
  "ColumnSynonyms" : [ String, ... ],
  "ComparativeOrder" : ComparativeOrder,
  "DefaultFormatting" : DefaultFormatting,
  "DisableIndexing" : Boolean,
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
  CellValueSynonyms:
    - CellValueSynonym
  ColumnDataRole: String
  ColumnDescription: String
  ColumnFriendlyName: String
  ColumnName: String
  ColumnSynonyms:
    - String
  ComparativeOrder:
    ComparativeOrder
  DefaultFormatting:
    DefaultFormatting
  DisableIndexing: Boolean
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

The type of aggregation that is performed on the column data when
it's queried.

_Required_: No

_Type_: String

_Allowed values_: `SUM | MAX | MIN | COUNT | DISTINCT_COUNT | AVERAGE | MEDIAN | STDEV | STDEVP | VAR | VARP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AllowedAggregations`

The list of aggregation types that are allowed for the column. Valid values for this
structure are `COUNT`, `DISTINCT_COUNT`, `MIN`,
`MAX`, `MEDIAN`, `SUM`, `AVERAGE`,
`STDEV`, `STDEVP`, `VAR`,
`VARP`,
and `PERCENTILE`.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CellValueSynonyms`

The other names or aliases for the column cell value.

_Required_: No

_Type_: Array of [CellValueSynonym](aws-properties-quicksight-topic-cellvaluesynonym.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColumnDataRole`

The role of the column in the data. Valid values are `DIMENSION` and `MEASURE`.

_Required_: No

_Type_: String

_Allowed values_: `DIMENSION | MEASURE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColumnDescription`

A description of the column and its contents.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColumnFriendlyName`

A user-friendly name for the column.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColumnName`

The name of the column.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColumnSynonyms`

The other names or aliases for the column.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComparativeOrder`

The order in which data is displayed for the column when
it's used in a comparative context.

_Required_: No

_Type_: [ComparativeOrder](aws-properties-quicksight-topic-comparativeorder.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultFormatting`

The default formatting used for values in the column.

_Required_: No

_Type_: [DefaultFormatting](aws-properties-quicksight-topic-defaultformatting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisableIndexing`

A Boolean value that indicates whether the column shows in the autocomplete functionality.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsIncludedInTopic`

A Boolean value that indicates whether the column is included in the query results.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NeverAggregateInFilter`

A Boolean
value that indicates whether to aggregate the column data when
it's used in a filter context.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NonAdditive`

The non additive value for the column.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotAllowedAggregations`

The list of aggregation types that are not allowed for the column. Valid values for this
structure are `COUNT`, `DISTINCT_COUNT`, `MIN`,
`MAX`, `MEDIAN`, `SUM`, `AVERAGE`,
`STDEV`, `STDEVP`, `VAR`,
`VARP`,
and `PERCENTILE`.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SemanticType`

The semantic type of data contained in the column.

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

TopicCategoryFilterConstant

TopicConfigOptions

All content copied from https://docs.aws.amazon.com/.
