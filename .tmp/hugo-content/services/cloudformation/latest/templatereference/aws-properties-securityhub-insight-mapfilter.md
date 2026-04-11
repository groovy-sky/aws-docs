This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::Insight MapFilter

A map filter for filtering AWS Security Hub CSPM findings. Each map filter provides the field to check for, the
value to check for, and the comparison operator.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Comparison" : String,
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Comparison: String
  Key: String
  Value: String

```

## Properties

`Comparison`

The condition to apply to the key value when filtering Security Hub CSPM findings with a map
filter.

To search for values that have the filter value, use one of the following comparison operators:

- To search for values that include the filter value, use `CONTAINS`. For example, for the
`ResourceTags` field, the filter `Department CONTAINS Security` matches findings that
include the value `Security` for the `Department` tag. In the same example, a finding with a value of
`Security team` for the `Department` tag is a match.

- To search for values that exactly match the filter value, use `EQUALS`. For example, for
the `ResourceTags` field, the filter `Department EQUALS Security` matches findings that
have the value `Security` for the `Department` tag.

`CONTAINS` and `EQUALS` filters on the same field are joined by `OR`. A
finding matches if it matches any one of those filters. For example, the filters `Department CONTAINS Security OR
               Department CONTAINS Finance` match a finding that includes either `Security`,
`Finance`, or both values.

To search for values that don't have the filter value, use one of the following comparison operators:

- To search for values that exclude the filter value, use `NOT_CONTAINS`. For example, for
the `ResourceTags` field, the filter `Department NOT_CONTAINS Finance` matches findings
that exclude the value `Finance` for the `Department` tag.

- To search for values other than the filter value, use `NOT_EQUALS`. For example, for the
`ResourceTags` field, the filter `Department NOT_EQUALS Finance` matches findings that
don’t have the value `Finance` for the `Department` tag.

`NOT_CONTAINS` and `NOT_EQUALS` filters on the same field are joined by `AND`.
A finding matches only if it matches all of those filters. For example, the filters `Department NOT_CONTAINS Security AND
               Department NOT_CONTAINS Finance` match a finding that excludes both the `Security` and
`Finance` values.

`CONTAINS` filters can only be used with other `CONTAINS` filters. `NOT_CONTAINS`
filters can only be used with other `NOT_CONTAINS` filters.

You can’t have both a `CONTAINS` filter and a `NOT_CONTAINS` filter on the same field.
Similarly, you can’t have both an `EQUALS` filter and a `NOT_EQUALS` filter on the same field.
Combining filters in this way returns an error.

`CONTAINS` and `NOT_CONTAINS` operators can be used only with automation rules. For more information,
see [Automation rules](../../../securityhub/latest/userguide/automation-rules.md) in the _AWS Security Hub CSPM User Guide_.

_Required_: Yes

_Type_: String

_Allowed values_: `EQUALS | NOT_EQUALS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Key`

The key of the map filter. For example, for `ResourceTags`, `Key`
identifies the name of the tag. For `UserDefinedFields`, `Key` is the
name of the field.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value for the key in the map filter. Filter values are case sensitive. For example,
one of the values for a tag called `Department` might be `Security`.
If you provide `security` as the filter value, then there's no match.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KeywordFilter

NumberFilter

All content copied from https://docs.aws.amazon.com/.
