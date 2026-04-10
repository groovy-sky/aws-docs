This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::ConfiguredTable

Creates a new configured table resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CleanRooms::ConfiguredTable",
  "Properties" : {
      "AllowedColumns" : [ String, ... ],
      "AnalysisMethod" : String,
      "AnalysisRules" : [ AnalysisRule, ... ],
      "Description" : String,
      "Name" : String,
      "SelectedAnalysisMethods" : [ String, ... ],
      "TableReference" : TableReference,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CleanRooms::ConfiguredTable
Properties:
  AllowedColumns:
    - String
  AnalysisMethod: String
  AnalysisRules:
    - AnalysisRule
  Description: String
  Name: String
  SelectedAnalysisMethods:
    - String
  TableReference:
    TableReference
  Tags:
    - Tag

```

## Properties

`AllowedColumns`

The columns within the underlying AWS Glue table that can be used within
collaborations.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `128 | 100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AnalysisMethod`

The analysis method for the configured table.

`DIRECT_QUERY` allows SQL queries to be run directly on this table.

`DIRECT_JOB` allows PySpark jobs to be run directly on this table.

`MULTIPLE` allows both SQL queries and PySpark jobs to be run directly on this table.

_Required_: Yes

_Type_: String

_Allowed values_: `DIRECT_QUERY | DIRECT_JOB | MULTIPLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AnalysisRules`

The analysis rule that was created for the configured table.

_Required_: No

_Type_: Array of [AnalysisRule](aws-properties-cleanrooms-configuredtable-analysisrule.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description for the configured table.

_Required_: No

_Type_: String

_Pattern_: `^[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDBFF-\uDC00\uDFFF\t\r\n]*$`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A name for the configured table.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!\s*$)[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDBFF-\uDC00\uDFFF\t]*$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectedAnalysisMethods`

The selected analysis methods for the configured table.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableReference`

The table that this configured table represents.

_Required_: Yes

_Type_: [TableReference](aws-properties-cleanrooms-configuredtable-tablereference.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An optional label that you can assign to a resource when you create it. Each tag
consists of a key and an optional value, both of which you define. When you use tagging,
you can also use tag-based access control in IAM policies to control access
to this resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-cleanrooms-configuredtable-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:

`{"Ref": "MyConfiguredTable"}`

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the Amazon Resource Name (ARN) of the specified configured table.

Example: `arn:aws:cleanrooms:us-east-1:111122223333:configuredtable/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`

`ConfiguredTableIdentifier`

Returns the unique identifier of the specified configured table.

Example: `a1b2c3d4-5678-90ab-cdef-EXAMPLE33333`

## Examples

### A configured table using a list analysis rule

The following is an example of a configured table with a list analysis rule
applied.

#### JSON

```json

"ListConfiguredTable": {
  {
    "Type" : "AWS::CleanRooms::ConfiguredTable",
    "Properties" : {
        "Name" : "List Table",
        "Description" : "Example configured table with list AR",
        "AllowedColumns" : ["column1", "column2", "column4"],
        "AnalysisMethod" : "DIRECT_QUERY",
        "AnalysisRules" : [
          "Type": "LIST",
          "Policy": {
            "V1": {
              "List": {
                "JoinColumns": [
                  "column1"
                ],
                "ListColumns": [
                  "column2"
                ]
              }
            }
          }
        ],
        "TableReference" : {
          "Glue": {
            "DatabaseName": "ExampleDB",
            "TableName": "ExampleTable"
          }
        }
      }
  }
}
```

#### YAML

```yaml

ListConfiguredTable:
  Type: AWS::CleanRooms::ConfiguredTable
  Properties:
    Name: List Table
    Description: Example configured table with list AR
    AllowedColumns:
      - column1
      - column2
      - column4
    AnalysisMethod: DIRECT_QUERY
    AnalysisRules:
      - Type: LIST
        Policy:
          V1:
            List:
              JoinColumns:
                - column1
              ListColumns:
                - column2
    TableReference:
      Glue:
        DatabaseName: ExampleDB
        TableName: ExampleTable
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AggregateColumn

All content copied from https://docs.aws.amazon.com/.
