This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DynamoDB::Table Csv

The options for imported source files in CSV format. The values are Delimiter and
HeaderList.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Delimiter" : String,
  "HeaderList" : [ String, ... ]
}

```

### YAML

```yaml

  Delimiter: String
  HeaderList:
    - String

```

## Properties

`Delimiter`

The delimiter used for separating items in the CSV file being imported.

_Required_: No

_Type_: String

_Pattern_: `[,;:|\t ]`

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`HeaderList`

List of the headers used to specify a common header for all source CSV files being
imported. If this field is specified then the first line of each CSV file is treated as
data instead of the header. If this field is not specified the the first line of each
CSV file is treated as the header.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContributorInsightsSpecification

GlobalSecondaryIndex

All content copied from https://docs.aws.amazon.com/.
