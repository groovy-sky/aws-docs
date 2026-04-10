This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Ruleset Rule

Represents a single data quality requirement that should be validated in the
scope of this dataset.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CheckExpression" : String,
  "ColumnSelectors" : [ ColumnSelector, ... ],
  "Disabled" : Boolean,
  "Name" : String,
  "SubstitutionMap" : [ SubstitutionValue, ... ],
  "Threshold" : Threshold
}

```

### YAML

```yaml

  CheckExpression: String
  ColumnSelectors:
    - ColumnSelector
  Disabled: Boolean
  Name: String
  SubstitutionMap:
    - SubstitutionValue
  Threshold:
    Threshold

```

## Properties

`CheckExpression`

The expression which includes column references, condition names followed by variable
references, possibly grouped and combined with other conditions. For example,
`(:col1 starts_with :prefix1 or :col1 starts_with :prefix2) and (:col1
                ends_with :suffix1 or :col1 ends_with :suffix2)`. Column and value references
are substitution variables that should start with the ':' symbol. Depending on the
context, substitution variables' values can be either an actual value or a column name.
These values are defined in the SubstitutionMap. If a CheckExpression starts with a
column reference, then ColumnSelectors in the rule should be null. If ColumnSelectors
has been defined, then there should be no columnn reference in the left side of a
condition, for example, `is_between :val1 and :val2`.

_Required_: Yes

_Type_: String

_Pattern_: `^[><0-9A-Za-z_.,:)(!= ]+$`

_Minimum_: `4`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColumnSelectors`

List of column selectors. Selectors can be used to select columns using a name or
regular expression from the dataset. Rule will be applied to selected columns.

_Required_: No

_Type_: Array of [ColumnSelector](aws-properties-databrew-ruleset-columnselector.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Disabled`

A value that specifies whether the rule is disabled. Once a rule is disabled, a
profile job will not validate it during a job run. Default value is false.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the rule.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubstitutionMap`

The map of substitution variable names to their values used in a check expression.
Variable names should start with a ':' (colon). Variable values can either be actual
values or column names. To differentiate between the two, column names should be
enclosed in backticks, for example, ``":col1": "`Column A`".``

_Required_: No

_Type_: Array of [SubstitutionValue](aws-properties-databrew-ruleset-substitutionvalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Threshold`

The threshold used with a non-aggregate check expression. Non-aggregate check
expressions will be applied to each row in a specific column, and the threshold will be
used to determine whether the validation succeeds.

_Required_: No

_Type_: [Threshold](aws-properties-databrew-ruleset-threshold.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ColumnSelector

SubstitutionValue

All content copied from https://docs.aws.amazon.com/.
