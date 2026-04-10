This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::ConfiguredTable AthenaTableReference

A reference to a table within Athena.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CatalogName" : String,
  "DatabaseName" : String,
  "OutputLocation" : String,
  "Region" : String,
  "TableName" : String,
  "WorkGroup" : String
}

```

### YAML

```yaml

  CatalogName: String
  DatabaseName: String
  OutputLocation: String
  Region: String
  TableName: String
  WorkGroup: String

```

## Properties

`CatalogName`

The catalog name.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseName`

The database name.

_Required_: Yes

_Type_: String

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputLocation`

The output location for the Athena table.

_Required_: No

_Type_: String

_Minimum_: `8`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Region`

The AWS Region where the Athena table is located. This parameter is required to uniquely identify and access tables across different Regions.

_Required_: No

_Type_: String

_Allowed values_: `us-west-1 | us-west-2 | us-east-1 | us-east-2 | af-south-1 | ap-east-1 | ap-south-2 | ap-southeast-1 | ap-southeast-2 | ap-southeast-5 | ap-southeast-4 | ap-southeast-7 | ap-south-1 | ap-northeast-3 | ap-northeast-1 | ap-northeast-2 | ca-central-1 | ca-west-1 | eu-south-1 | eu-west-3 | eu-south-2 | eu-central-2 | eu-central-1 | eu-north-1 | eu-west-1 | eu-west-2 | me-south-1 | me-central-1 | il-central-1 | sa-east-1 | mx-central-1 | ap-east-2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableName`

The table reference.

_Required_: Yes

_Type_: String

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkGroup`

The workgroup of the Athena table reference.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AnalysisRuleList

ConfiguredTableAnalysisRulePolicy

All content copied from https://docs.aws.amazon.com/.
