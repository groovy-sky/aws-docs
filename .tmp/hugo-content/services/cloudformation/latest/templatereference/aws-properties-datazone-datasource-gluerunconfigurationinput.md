This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::DataSource GlueRunConfigurationInput

The configuration details of the AWS Glue data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoImportDataQualityResult" : Boolean,
  "CatalogName" : String,
  "DataAccessRole" : String,
  "RelationalFilterConfigurations" : [ RelationalFilterConfiguration, ... ]
}

```

### YAML

```yaml

  AutoImportDataQualityResult: Boolean
  CatalogName: String
  DataAccessRole: String
  RelationalFilterConfigurations:
    - RelationalFilterConfiguration

```

## Properties

`AutoImportDataQualityResult`

Specifies whether to automatically import data quality metrics as part of the data
source run.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CatalogName`

The catalog name in the AWS Glue run configuration.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataAccessRole`

The data access role included in the configuration details of the AWS Glue data
source.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[^:]*:iam::\d{12}:role(/[a-zA-Z0-9+=,.@_-]+)*/[a-zA-Z0-9+=,.@_-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RelationalFilterConfigurations`

The relational filter configurations included in the configuration details of the AWS
Glue data source.

_Required_: Yes

_Type_: Array of [RelationalFilterConfiguration](aws-properties-datazone-datasource-relationalfilterconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FormInput

RecommendationConfiguration

All content copied from https://docs.aws.amazon.com/.
