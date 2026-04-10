This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Topic TopicFilter

A structure that represents a filter used to select items for a topic.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CategoryFilter" : TopicCategoryFilter,
  "DateRangeFilter" : TopicDateRangeFilter,
  "FilterClass" : String,
  "FilterDescription" : String,
  "FilterName" : String,
  "FilterSynonyms" : [ String, ... ],
  "FilterType" : String,
  "NumericEqualityFilter" : TopicNumericEqualityFilter,
  "NumericRangeFilter" : TopicNumericRangeFilter,
  "OperandFieldName" : String,
  "RelativeDateFilter" : TopicRelativeDateFilter
}

```

### YAML

```yaml

  CategoryFilter:
    TopicCategoryFilter
  DateRangeFilter:
    TopicDateRangeFilter
  FilterClass: String
  FilterDescription: String
  FilterName: String
  FilterSynonyms:
    - String
  FilterType: String
  NumericEqualityFilter:
    TopicNumericEqualityFilter
  NumericRangeFilter:
    TopicNumericRangeFilter
  OperandFieldName: String
  RelativeDateFilter:
    TopicRelativeDateFilter

```

## Properties

`CategoryFilter`

The category filter that is associated with this filter.

_Required_: No

_Type_: [TopicCategoryFilter](aws-properties-quicksight-topic-topiccategoryfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DateRangeFilter`

The date range filter.

_Required_: No

_Type_: [TopicDateRangeFilter](aws-properties-quicksight-topic-topicdaterangefilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterClass`

The class of the filter. Valid values for this structure are
`ENFORCED_VALUE_FILTER`,
`CONDITIONAL_VALUE_FILTER`,
and `NAMED_VALUE_FILTER`.

_Required_: No

_Type_: String

_Allowed values_: `ENFORCED_VALUE_FILTER | CONDITIONAL_VALUE_FILTER | NAMED_VALUE_FILTER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterDescription`

A description of the filter used to select items for a topic.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterName`

The name of the filter.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterSynonyms`

The other names or aliases for the filter.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterType`

The type of the filter. Valid values for this structure are
`CATEGORY_FILTER`, `NUMERIC_EQUALITY_FILTER`,
`NUMERIC_RANGE_FILTER`,
`DATE_RANGE_FILTER`,
and `RELATIVE_DATE_FILTER`.

_Required_: No

_Type_: String

_Allowed values_: `CATEGORY_FILTER | NUMERIC_EQUALITY_FILTER | NUMERIC_RANGE_FILTER | DATE_RANGE_FILTER | RELATIVE_DATE_FILTER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumericEqualityFilter`

The numeric equality filter.

_Required_: No

_Type_: [TopicNumericEqualityFilter](aws-properties-quicksight-topic-topicnumericequalityfilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumericRangeFilter`

The numeric range filter.

_Required_: No

_Type_: [TopicNumericRangeFilter](aws-properties-quicksight-topic-topicnumericrangefilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OperandFieldName`

The name of the field that the filter operates on.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RelativeDateFilter`

The relative date filter.

_Required_: No

_Type_: [TopicRelativeDateFilter](aws-properties-quicksight-topic-topicrelativedatefilter.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TopicDateRangeFilter

TopicNamedEntity

All content copied from https://docs.aws.amazon.com/.
