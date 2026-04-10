This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Recipe RecipeParameters

Parameters that are used as inputs for various recipe actions. The parameters are
specific to the context in which they're used.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AggregateFunction" : String,
  "Base" : String,
  "CaseStatement" : String,
  "CategoryMap" : String,
  "CharsToRemove" : String,
  "CollapseConsecutiveWhitespace" : String,
  "ColumnDataType" : String,
  "ColumnRange" : String,
  "Count" : String,
  "CustomCharacters" : String,
  "CustomStopWords" : String,
  "CustomValue" : String,
  "DatasetsColumns" : String,
  "DateAddValue" : String,
  "DateTimeFormat" : String,
  "DateTimeParameters" : String,
  "DeleteOtherRows" : String,
  "Delimiter" : String,
  "EndPattern" : String,
  "EndPosition" : String,
  "EndValue" : String,
  "ExpandContractions" : String,
  "Exponent" : String,
  "FalseString" : String,
  "GroupByAggFunctionOptions" : String,
  "GroupByColumns" : String,
  "HiddenColumns" : String,
  "IgnoreCase" : String,
  "IncludeInSplit" : String,
  "Input" : Input,
  "Interval" : String,
  "IsText" : String,
  "JoinKeys" : String,
  "JoinType" : String,
  "LeftColumns" : String,
  "Limit" : String,
  "LowerBound" : String,
  "MapType" : String,
  "ModeType" : String,
  "MultiLine" : Boolean,
  "NumRows" : String,
  "NumRowsAfter" : String,
  "NumRowsBefore" : String,
  "OrderByColumn" : String,
  "OrderByColumns" : String,
  "Other" : String,
  "Pattern" : String,
  "PatternOption1" : String,
  "PatternOption2" : String,
  "PatternOptions" : String,
  "Period" : String,
  "Position" : String,
  "RemoveAllPunctuation" : String,
  "RemoveAllQuotes" : String,
  "RemoveAllWhitespace" : String,
  "RemoveCustomCharacters" : String,
  "RemoveCustomValue" : String,
  "RemoveLeadingAndTrailingPunctuation" : String,
  "RemoveLeadingAndTrailingQuotes" : String,
  "RemoveLeadingAndTrailingWhitespace" : String,
  "RemoveLetters" : String,
  "RemoveNumbers" : String,
  "RemoveSourceColumn" : String,
  "RemoveSpecialCharacters" : String,
  "RightColumns" : String,
  "SampleSize" : String,
  "SampleType" : String,
  "SecondaryInputs" : [ SecondaryInput, ... ],
  "SecondInput" : String,
  "SheetIndexes" : [ Integer, ... ],
  "SheetNames" : [ String, ... ],
  "SourceColumn" : String,
  "SourceColumn1" : String,
  "SourceColumn2" : String,
  "SourceColumns" : String,
  "StartColumnIndex" : String,
  "StartPattern" : String,
  "StartPosition" : String,
  "StartValue" : String,
  "StemmingMode" : String,
  "StepCount" : String,
  "StepIndex" : String,
  "StopWordsMode" : String,
  "Strategy" : String,
  "TargetColumn" : String,
  "TargetColumnNames" : String,
  "TargetDateFormat" : String,
  "TargetIndex" : String,
  "TimeZone" : String,
  "TokenizerPattern" : String,
  "TrueString" : String,
  "UdfLang" : String,
  "Units" : String,
  "UnpivotColumn" : String,
  "UpperBound" : String,
  "UseNewDataFrame" : String,
  "Value" : String,
  "Value1" : String,
  "Value2" : String,
  "ValueColumn" : String,
  "ViewFrame" : String
}

```

### YAML

```yaml

  AggregateFunction: String
  Base: String
  CaseStatement: String
  CategoryMap: String
  CharsToRemove: String
  CollapseConsecutiveWhitespace: String
  ColumnDataType: String
  ColumnRange: String
  Count: String
  CustomCharacters: String
  CustomStopWords: String
  CustomValue: String
  DatasetsColumns: String
  DateAddValue: String
  DateTimeFormat: String
  DateTimeParameters: String
  DeleteOtherRows: String
  Delimiter: String
  EndPattern: String
  EndPosition: String
  EndValue: String
  ExpandContractions: String
  Exponent: String
  FalseString:
    String
  GroupByAggFunctionOptions: String
  GroupByColumns: String
  HiddenColumns: String
  IgnoreCase: String
  IncludeInSplit: String
  Input:
    Input
  Interval: String
  IsText: String
  JoinKeys: String
  JoinType: String
  LeftColumns: String
  Limit: String
  LowerBound: String
  MapType: String
  ModeType: String
  MultiLine: Boolean
  NumRows: String
  NumRowsAfter: String
  NumRowsBefore: String
  OrderByColumn: String
  OrderByColumns: String
  Other: String
  Pattern: String
  PatternOption1: String
  PatternOption2: String
  PatternOptions: String
  Period: String
  Position: String
  RemoveAllPunctuation: String
  RemoveAllQuotes: String
  RemoveAllWhitespace: String
  RemoveCustomCharacters: String
  RemoveCustomValue: String
  RemoveLeadingAndTrailingPunctuation: String
  RemoveLeadingAndTrailingQuotes: String
  RemoveLeadingAndTrailingWhitespace: String
  RemoveLetters: String
  RemoveNumbers: String
  RemoveSourceColumn: String
  RemoveSpecialCharacters: String
  RightColumns: String
  SampleSize: String
  SampleType: String
  SecondaryInputs:
    - SecondaryInput
  SecondInput: String
  SheetIndexes:
    - Integer
  SheetNames:
    - String
  SourceColumn: String
  SourceColumn1: String
  SourceColumn2: String
  SourceColumns: String
  StartColumnIndex: String
  StartPattern: String
  StartPosition: String
  StartValue: String
  StemmingMode: String
  StepCount: String
  StepIndex: String
  StopWordsMode: String
  Strategy: String
  TargetColumn: String
  TargetColumnNames: String
  TargetDateFormat: String
  TargetIndex: String
  TimeZone: String
  TokenizerPattern: String
  TrueString:
    String
  UdfLang: String
  Units: String
  UnpivotColumn: String
  UpperBound: String
  UseNewDataFrame: String
  Value: String
  Value1: String
  Value2: String
  ValueColumn: String
  ViewFrame: String

```

## Properties

`AggregateFunction`

The name of an aggregation function to apply.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Base`

The number of digits used in a counting system.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CaseStatement`

A case statement associated with a recipe.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CategoryMap`

A category map used for one-hot encoding.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CharsToRemove`

Characters to remove from a step that applies one-hot encoding or tokenization.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CollapseConsecutiveWhitespace`

Remove any non-word non-punctuation character.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColumnDataType`

The data type of the column.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColumnRange`

A range of columns to which a step is applied.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Count`

The number of times a string needs to be repeated.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomCharacters`

One or more characters that can be substituted or removed, depending on the
context.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomStopWords`

A list of words to ignore in a step that applies word tokenization.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomValue`

A list of custom values to use in a step that requires that you provide a value to
finish the operation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatasetsColumns`

A list of the dataset columns included in a project.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DateAddValue`

A value that specifies how many units of time to add or subtract for a date math
operation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DateTimeFormat`

A date format to apply to a date.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DateTimeParameters`

A set of parameters associated with a datetime.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeleteOtherRows`

Determines whether unmapped rows in a categorical mapping should be deleted

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Delimiter`

The delimiter to use when parsing separated values in a text file.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndPattern`

The end pattern to locate.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndPosition`

The end position to locate.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EndValue`

The end value to locate.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExpandContractions`

A list of word contractions and what they expand to. For eample:
_can't_; _cannot_; _can_
_not_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Exponent`

The exponent to apply in an exponential operation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FalseString`

A value that represents `FALSE`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GroupByAggFunctionOptions`

Specifies options to apply to the `GROUP BY` used in an aggregation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GroupByColumns`

The columns to use in the `GROUP BY` clause.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HiddenColumns`

A list of columns to hide.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IgnoreCase`

Indicates that lower and upper case letters are treated equally.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeInSplit`

Indicates if this column is participating in a split transform.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Input`

The input location to load the dataset from - Amazon S3 or AWS Glue Data Catalog.

_Required_: No

_Type_: [Input](aws-properties-databrew-recipe-input.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Interval`

The number of characters to split by.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IsText`

Indicates if the content is text.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JoinKeys`

The keys or columns involved in a join.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JoinType`

The type of join to use, for example, `INNER JOIN`, `OUTER
            JOIN`, and so on.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LeftColumns`

The columns on the left side of the join.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Limit`

The number of times to perform `split` or `replaceBy` in a
string

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LowerBound`

The lower boundary for a value.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MapType`

The type of mappings to apply to construct a new dynamic frame.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModeType`

Determines the manner in which mode value is calculated, in case there is more than
one mode value. Valid values: `NONE` \| `AVERAGE` \|
`MINIMUM` \| `MAXIMUM`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MultiLine`

Specifies whether JSON input contains embedded new line characters.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumRows`

The number of rows to consider in a window.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumRowsAfter`

The number of rows to consider after the current row in a window

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumRowsBefore`

The number of rows to consider before the current row in a window

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrderByColumn`

A column to sort the results by.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrderByColumns`

The columns to sort the results by.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Other`

The value to assign to unmapped cells, in categorical mapping

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Pattern`

The pattern to locate.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PatternOption1`

The starting pattern to split between.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PatternOption2`

The ending pattern to split between.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PatternOptions`

For splitting by multiple delimiters: A JSON-encoded string that lists the patterns in
the format. For example:
`[{\"pattern\":\"1\",\"includeInSplit\":true}]`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Period`

The size of the rolling window.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Position`

The character index within a string

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RemoveAllPunctuation`

If `true`, removes all of the following characters: `.` `.!` `.,` `.?`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RemoveAllQuotes`

If `true`, removes all single quotes and double quotes.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RemoveAllWhitespace`

If `true`, removes all whitespaces from the value.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RemoveCustomCharacters`

If `true`, removes all chraracters specified by
`CustomCharacters`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RemoveCustomValue`

If `true`, removes all chraracters specified by `CustomValue`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RemoveLeadingAndTrailingPunctuation`

If `true`, removes the following characters if they occur at the start or
end of the value: `.` `!` `,` `?`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RemoveLeadingAndTrailingQuotes`

If `true`, removes single quotes and double quotes from the beginning and
end of the value.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RemoveLeadingAndTrailingWhitespace`

If `true`, removes all whitespaces from the beginning and end of the
value.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RemoveLetters`

If `true`, removes all uppercase and lowercase alphabetic characters (A
through Z; a through z).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RemoveNumbers`

If `true`, removes all numeric characters (0 through 9).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RemoveSourceColumn`

If `true`, the source column will be removed after un-nesting that column.
(Used with nested column types, such as Map, Struct, or Array.)

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RemoveSpecialCharacters`

If `true`, removes all of the following characters: ``! " # $ % & '
                ( ) * + , - . / : ; < = > ? @ [ \ ] ^ _ ` { | } ~``

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RightColumns`

The columns on the right side of a join.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SampleSize`

The number of rows in the sample.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SampleType`

The sampling type to apply to the dataset. Valid values: `FIRST_N` \|
`LAST_N` \| `RANDOM`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecondaryInputs`

A list of secondary inputs in a UNION transform

_Required_: No

_Type_: Array of [SecondaryInput](aws-properties-databrew-recipe-secondaryinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecondInput`

A object value to indicate the second dataset used in a join.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SheetIndexes`

One or more sheet numbers in the Excel file, which will be included in a
dataset.

_Required_: No

_Type_: Array of Integer

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SheetNames`

Oone or more named sheets in the Excel file, which will be included in a dataset.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceColumn`

A source column needed for an operation, step, or transform.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceColumn1`

A source column needed for an operation, step, or transform.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceColumn2`

A source column needed for an operation, step, or transform.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceColumns`

A list of source columns needed for an operation, step, or transform.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartColumnIndex`

The index number of the first column used by an operation, step, or transform.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartPattern`

The starting pattern to locate.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartPosition`

The starting position to locate.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartValue`

The starting value to locate.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StemmingMode`

Indicates this operation uses stems and lemmas (base words) for word tokenization.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StepCount`

The total number of transforms in this recipe.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StepIndex`

The index ID of a step.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StopWordsMode`

Indicates this operation uses stop words as part of word tokenization.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Strategy`

The resolution strategy to apply in resolving ambiguities.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetColumn`

The column targeted by this operation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetColumnNames`

The names to give columns altered by this operation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetDateFormat`

The date format to convert to.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetIndex`

The index number of an object that is targeted by this operation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeZone`

The current timezone that you want to use for dates.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TokenizerPattern`

A regex expression to use when splitting text into terms, also called words or
tokens.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrueString`

A value to use to represent `TRUE`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UdfLang`

The language that's used in the user-defined function.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Units`

Specifies a unit of time. For example: `MINUTES`; `SECONDS`;
`HOURS`; etc.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UnpivotColumn`

Cast columns as rows, so that each value is a different row in a single column.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UpperBound`

The upper boundary for a value.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseNewDataFrame`

Create a new container to hold a dataset.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

A static value that can be used in a comparison, a substitution, or in another
context-specific way. A `Value` can be a number, string, or other datatype,
depending on the recipe action in which it's used.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value1`

A value that's used by this operation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value2`

A value that's used by this operation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueColumn`

The column that is provided as a value that's used by this operation.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ViewFrame`

The subset of rows currently available for viewing.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Parameters

RecipeStep

All content copied from https://docs.aws.amazon.com/.
