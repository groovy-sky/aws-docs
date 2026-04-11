This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Classifier CsvClassifier

A classifier for custom `CSV` content.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowSingleColumn" : Boolean,
  "ContainsCustomDatatype" : [ String, ... ],
  "ContainsHeader" : String,
  "CustomDatatypeConfigured" : Boolean,
  "Delimiter" : String,
  "DisableValueTrimming" : Boolean,
  "Header" : [ String, ... ],
  "Name" : String,
  "QuoteSymbol" : String
}

```

### YAML

```yaml

  AllowSingleColumn: Boolean
  ContainsCustomDatatype:
    - String
  ContainsHeader: String
  CustomDatatypeConfigured: Boolean
  Delimiter: String
  DisableValueTrimming: Boolean
  Header:
    - String
  Name: String
  QuoteSymbol: String

```

## Properties

`AllowSingleColumn`

Enables the processing of files that contain only one column.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContainsCustomDatatype`

Indicates whether the CSV file contains custom data types.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContainsHeader`

Indicates whether the CSV file contains a header.

A value of `UNKNOWN` specifies that the classifier will detect whether the CSV file contains headings.

A value of `PRESENT` specifies that the CSV file contains headings.

A value of `ABSENT` specifies that the CSV file does not contain headings.

_Required_: No

_Type_: String

_Allowed values_: `UNKNOWN | PRESENT | ABSENT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomDatatypeConfigured`

Enables the configuration of custom data types.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Delimiter`

A custom symbol to denote what separates each column entry in the row.

_Required_: No

_Type_: String

_Pattern_: `[^\r\n]`

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisableValueTrimming`

Specifies not to trim values before identifying the type of column values. The default
value is `true`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Header`

A list of strings representing column names.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the classifier.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QuoteSymbol`

A custom symbol to denote what combines content into a single column value. It must be
different from the column delimiter.

_Required_: No

_Type_: String

_Pattern_: `[^\r\n]`

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Create a custom classifier test

With `AWS::Glue::Classifier` you can create a custom classifier test.

#### JSON

```json

{
    "Description": "AWS Glue custom classifier test",
    "Resources": {
        "MyCSVclassifier": {
            "Type": "AWS::Glue::Classifier",
            "Properties": {
                "CsvClassifier": {
                    "AllowSingleColumn": true,
                    "ContainsHeader": "PRESENT",
                    "Delimiter": ",",
                    "Header": [
                        "id",
                        "name"
                    ],
                    "Name": "csvclassify",
                    "QuoteSymbol": "\""
                }
            }
        }
    }
}
```

#### YAML

```yaml

Description: AWS Glue custom classifier test
Resources:
  MyCSVclassifier:
    Type: 'AWS::Glue::Classifier'
    Properties:
      CsvClassifier:
        AllowSingleColumn: true
        ContainsHeader: PRESENT
        Delimiter: ','
        Header:
          - id
          - name
        Name: csvclassify
        QuoteSymbol: '"'
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Glue::Classifier

GrokClassifier

All content copied from https://docs.aws.amazon.com/.
