This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LakeFormation::DataCellsFilter

A structure that represents a data cell filter with column-level, row-level, and/or cell-level security. Data cell filters belong to a specific table in a Data Catalog. During a stack operation,
AWS CloudFormation calls the AWS Lake Formation `CreateDataCellsFilter` API operation to create
a `DataCellsFilter` resource, and calls the `DeleteDataCellsFilter` API operation to delete it.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::LakeFormation::DataCellsFilter",
  "Properties" : {
      "ColumnNames" : [ String, ... ],
      "ColumnWildcard" : ColumnWildcard,
      "DatabaseName" : String,
      "Name" : String,
      "RowFilter" : RowFilter,
      "TableCatalogId" : String,
      "TableName" : String
    }
}

```

### YAML

```yaml

Type: AWS::LakeFormation::DataCellsFilter
Properties:
  ColumnNames:
    - String
  ColumnWildcard:
    ColumnWildcard
  DatabaseName: String
  Name: String
  RowFilter:
    RowFilter
  TableCatalogId: String
  TableName: String

```

## Properties

`ColumnNames`

An array of UTF-8 strings. A list of column names.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ColumnWildcard`

A wildcard with exclusions. You must specify either a `ColumnNames` list or the `ColumnWildCard`.

_Required_: No

_Type_: [ColumnWildcard](aws-properties-lakeformation-datacellsfilter-columnwildcard.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DatabaseName`

UTF-8 string, not less than 1 or more than 255 bytes long, matching the [single-line string pattern](../../../lake-formation/latest/dg/aws-lake-formation-api-aws-lake-formation-api-common.md).

A database in the Data Catalog.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

UTF-8 string, not less than 1 or more than 255 bytes long, matching the [single-line string pattern](../../../lake-formation/latest/dg/aws-lake-formation-api-aws-lake-formation-api-common.md).

The name given by the user to the data filter cell.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RowFilter`

A PartiQL predicate.

_Required_: No

_Type_: [RowFilter](aws-properties-lakeformation-datacellsfilter-rowfilter.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TableCatalogId`

Catalog id string, not less than 1 or more than 255 bytes long, matching the [single-line string pattern](../../../lake-formation/latest/dg/aws-lake-formation-api-aws-lake-formation-api-common.md).

The ID of the catalog to which the table belongs.

_Required_: Yes

_Type_: String

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TableName`

UTF-8 string, not less than 1 or more than 255 bytes long, matching the [single-line string pattern](../../../lake-formation/latest/dg/aws-lake-formation-api-aws-lake-formation-api-common.md).

A table in the database.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource properties such as TableCatalogId, DatabaseName, TableName, and FilterName.
For example: 123456789012\|ExampleDbName\|ExampleTableName\|ExampleFilterName

## Remarks

The level of filtering that you get depends on how you populate the data filter.

- When you specify the "all columns" wildcard and provide a row filter expression, you are establishing row-level security (row filtering) only.

- When you include or exclude specific columns and specify _all rows_ using the all-rows wildcard, you are establishing column-level security (column filtering) only.

- When you include or exclude specific columns and also provide a row filter expression, you are
establishing cell-level security (cell filtering).

_Specify the following to create a valid data cells filter:_

- `ColumnWildcard` or `ColumnNames`

- `RowFilter.AllRowsWildcard` or `RowFilter.FilterExpression`

## Examples

- [Creating a DataCellsFilter using row and column wildcards](#aws-resource-lakeformation-datacellsfilter--examples--Creating_a_DataCellsFilter_using_row_and_column_wildcards)

- [Creating a DataCellsFilter using a row wild card and specified columns](#aws-resource-lakeformation-datacellsfilter--examples--Creating_a_DataCellsFilter_using_a_row_wild_card_and_specified_columns)

- [Creating a DataCellsFilter using a row filter expression and a column wildcard](#aws-resource-lakeformation-datacellsfilter--examples--Creating_a_DataCellsFilter_using_a_row_filter_expression_and_a_column_wildcard)

- [Creating a DataCellsFilter using a row filter and specified columns](#aws-resource-lakeformation-datacellsfilter--examples--Creating_a_DataCellsFilter_using_a_row_filter_and_specified_columns)

### Creating a DataCellsFilter using row and column wildcards

The following example demonstrates how to create a `DataCellsFilter` resource using row and column wildcards:

#### JSON

```json

{
  "SampleDataCellsFilter": {
    "Type": "AWS::LakeFormation::DataCellsFilter",
    "Properties": {
      "TableCatalogId": "12345678910",
      "DatabaseName": "sample_db",
      "TableName": "sample_tbl",
      "Name": "sample_data_cells_filter",
      "RowFilter": {
          "AllRowsWildcard": {}
       },
          "ColumnWildcard": {}
     }
   }
}
```

#### YAML

```yaml

SampleDataCellsFilter:
    Type: AWS::LakeFormation::DataCellsFilter
    Properties:
        TableCatalogId: "12345678910"
        DatabaseName: "sample_db"
        TableName: "sample_tbl"
        Name: "sample_data_cells_filter"
        RowFilter:
            AllRowsWildcard: {}
        ColumnWildcard: {}

```

### Creating a DataCellsFilter using a row wild card and specified columns

The following example demonstrates how to create a `DataCellsFilter`
resource using a row wild card and specified columns:

#### JSON

```json

{
    "SampleDataCellsFilter": {
        "Type": "AWS::LakeFormation::DataCellsFilter",
        "Properties": {
            "TableCatalogId": "12345678910",
            "DatabaseName": "sample_db",
            "TableName": "sample_tbl",
            "Name": "sample_data_cells_filter",
            "RowFilter": {
                "AllRowsWildcard": {}
            },
            "ColumnNames": ["sample_column_1", "sample_column_2"]
        }
    }
}
```

#### YAML

```yaml

SampleDataCellsFilter:
    Type: AWS::LakeFormation::DataCellsFilter
    Properties:
        TableCatalogId: "12345678910"
        DatabaseName: "sample_db"
        TableName: "sample_tbl"
        Name: "sample_data_cells_filter"
        RowFilter:
            AllRowsWildcard: {}
        ColumnNames: ["sample_column_1", "sample_column_2"]
```

### Creating a DataCellsFilter using a row filter expression and a column wildcard

The following example demonstrates how to create a `DataCellsFilter` using a row filter expression and a column wildcard:

#### JSON

```json

{
    "SampleDataCellsFilter": {
        "Type": "AWS::LakeFormation::DataCellsFilter",
        "Properties": {
            "TableCatalogId": "12345678910",
            "DatabaseName": "sample_db",
            "TableName": "sample_tbl",
            "Name": "sample_data_cells_filter",
            "RowFilter": {
                "FilterExpression": "sample_column_1 > 0"
            },
            "ColumnWildcard": {}
        }
    }
}
```

#### YAML

```yaml

SampleDataCellsFilter:
    Type: AWS::LakeFormation::DataCellsFilter
    Properties:
        TableCatalogId: "12345678910"
        DatabaseName: "sample_db"
        TableName: "sample_tbl"
        Name: "sample_data_cells_filter"
        RowFilter:
            FilterExpression: "sample_column_1 > 0"
        ColumnWildcard: {}
```

### Creating a DataCellsFilter using a row filter and specified columns

The following example demonstrates how to create a `DataCellsFilter`
resource using a row filter and specified columns:

#### JSON

```json

{
    "SampleDataCellsFilter": {
        "Type": "AWS::LakeFormation::DataCellsFilter",
        "Properties": {
            "TableCatalogId": "12345678910",
            "DatabaseName": "sample_db",
            "TableName": "sample_tbl",
            "Name": "sample_data_cells_filter",
            "RowFilter": {
                "FilterExpression": "sample_column_1 > 0"
            },
            "ColumnNames": ["sample_column_1", "sample_column_2"]
        }
    }
}

```

#### YAML

```yaml

SampleDataCellsFilter:
    Type: AWS::LakeFormation::DataCellsFilter
    Properties:
        TableCatalogId: "12345678910"
        DatabaseName: "sample_db"
        TableName: "sample_tbl"
        Name: "sample_data_cells_filter"
        RowFilter:
            FilterExpression: "sample_column_1 > 0"
        ColumnNames: ["sample_column_1", "sample_column_2"]

```

## See also

[Data filtering and cell-level security in AWS Lake Formation.](../../../lake-formation/latest/dg/data-filters-about.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Lake Formation

ColumnWildcard

All content copied from https://docs.aws.amazon.com/.
