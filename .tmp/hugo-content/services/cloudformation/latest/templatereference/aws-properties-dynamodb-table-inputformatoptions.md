This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::Table InputFormatOptions

The format options for the data that was imported into the target table. There is one
value, CsvOption.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Csv" : Csv
}

```

### YAML

```yaml

  Csv:
    Csv

```

## Properties

`Csv`

The options for imported source files in CSV format. The values are Delimiter and
HeaderList.

_Required_: No

_Type_: [Csv](aws-properties-dynamodb-table-csv.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImportSourceSpecification

KeySchema

All content copied from https://docs.aws.amazon.com/.
