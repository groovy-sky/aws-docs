This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::DataSource RedshiftRunConfigurationInput

The relational filter configurations included in the configuration details of the Amazon Redshift data
source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataAccessRole" : String,
  "RedshiftCredentialConfiguration" : RedshiftCredentialConfiguration,
  "RedshiftStorage" : RedshiftStorage,
  "RelationalFilterConfigurations" : [ RelationalFilterConfiguration, ... ]
}

```

### YAML

```yaml

  DataAccessRole: String
  RedshiftCredentialConfiguration:
    RedshiftCredentialConfiguration
  RedshiftStorage:
    RedshiftStorage
  RelationalFilterConfigurations:
    - RelationalFilterConfiguration

```

## Properties

`DataAccessRole`

The data access role included in the configuration details of the Amazon Redshift data
source.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[^:]*:iam::\d{12}:role(/[a-zA-Z0-9+=,.@_-]+)*/[a-zA-Z0-9+=,.@_-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RedshiftCredentialConfiguration`

The details of the credentials required to access an Amazon Redshift cluster.

_Required_: No

_Type_: [RedshiftCredentialConfiguration](aws-properties-datazone-datasource-redshiftcredentialconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RedshiftStorage`

The details of the Amazon Redshift storage as part of the configuration of an Amazon
Redshift data source run.

_Required_: No

_Type_: [RedshiftStorage](aws-properties-datazone-datasource-redshiftstorage.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RelationalFilterConfigurations`

The relational filter configurations included in the configuration details of the AWS
Glue data source.

_Required_: Yes

_Type_: Array of [RelationalFilterConfiguration](aws-properties-datazone-datasource-relationalfilterconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RedshiftCredentialConfiguration

RedshiftServerlessStorage

All content copied from https://docs.aws.amazon.com/.
