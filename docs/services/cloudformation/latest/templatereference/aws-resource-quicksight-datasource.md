This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSource

Creates a data source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::QuickSight::DataSource",
  "Properties" : {
      "AlternateDataSourceParameters" : [ DataSourceParameters, ... ],
      "AwsAccountId" : String,
      "Credentials" : DataSourceCredentials,
      "DataSourceId" : String,
      "DataSourceParameters" : DataSourceParameters,
      "ErrorInfo" : DataSourceErrorInfo,
      "FolderArns" : [ String, ... ],
      "Name" : String,
      "Permissions" : [ ResourcePermission, ... ],
      "SslProperties" : SslProperties,
      "Tags" : [ Tag, ... ],
      "Type" : String,
      "VpcConnectionProperties" : VpcConnectionProperties
    }
}

```

### YAML

```yaml

Type: AWS::QuickSight::DataSource
Properties:
  AlternateDataSourceParameters:
    - DataSourceParameters
  AwsAccountId: String
  Credentials:
    DataSourceCredentials
  DataSourceId: String
  DataSourceParameters:
    DataSourceParameters
  ErrorInfo:
    DataSourceErrorInfo
  FolderArns:
    - String
  Name: String
  Permissions:
    - ResourcePermission
  SslProperties:
    SslProperties
  Tags:
    - Tag
  Type: String
  VpcConnectionProperties:
    VpcConnectionProperties

```

## Properties

`AlternateDataSourceParameters`

A set of alternate data source parameters that you want to share for the credentials
stored with this data source. The credentials are applied in tandem with the data source
parameters when you copy a data source by using a create or update request. The API
operation compares the `DataSourceParameters` structure that's in the request
with the structures in the `AlternateDataSourceParameters` allow list. If the
structures are an exact match, the request is allowed to use the credentials from this
existing data source. If the `AlternateDataSourceParameters` list is null,
the `Credentials` originally used with this `DataSourceParameters`
are automatically allowed.

_Required_: No

_Type_: Array of [DataSourceParameters](aws-properties-quicksight-datasource-datasourceparameters.md)

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AwsAccountId`

The AWS account ID.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Credentials`

The credentials Amazon Quick Sight that uses to connect to your underlying source.
Currently, only credentials based on user name and password are supported.

_Required_: No

_Type_: [DataSourceCredentials](aws-properties-quicksight-datasource-datasourcecredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataSourceId`

An ID for the data source. This ID is unique per AWS Region for each
AWS account.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataSourceParameters`

The parameters that Amazon Quick Sight uses to connect to your underlying
source.

_Required_: No

_Type_: [DataSourceParameters](aws-properties-quicksight-datasource-datasourceparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ErrorInfo`

Error information from the last update or the creation of the data source.

_Required_: No

_Type_: [DataSourceErrorInfo](aws-properties-quicksight-datasource-datasourceerrorinfo.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FolderArns`

Property description not available.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A display name for the data source.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Permissions`

A list of resource permissions on the data source.

_Required_: No

_Type_: Array of [ResourcePermission](aws-properties-quicksight-datasource-resourcepermission.md)

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SslProperties`

Secure Socket Layer (SSL) properties that apply when Amazon Quick Sight connects to
your underlying source.

_Required_: No

_Type_: [SslProperties](aws-properties-quicksight-datasource-sslproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Contains a map of the key-value pairs for the resource tag or tags assigned to the
data source.

_Required_: No

_Type_: Array of [Tag](aws-properties-quicksight-datasource-tag.md)

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the data source. To return a list of all data sources, use
`ListDataSources`.

Use `AMAZON_ELASTICSEARCH` for Amazon OpenSearch Service.

_Required_: Yes

_Type_: String

_Allowed values_: `ADOBE_ANALYTICS | AMAZON_ELASTICSEARCH | AMAZON_OPENSEARCH | ATHENA | AURORA | AURORA_POSTGRESQL | AWS_IOT_ANALYTICS | DATABRICKS | DENODO | DREMIO | DYNAMODB | SAPHANA | DB2_AS400 | EXASOL | FILE | GITHUB | INTERNATIONAL_DATA_CORPORATION | JIRA | MARIADB | MYSQL | ORACLE | POSTGRESQL | PRESTO | QBUSINESS | REDSHIFT | S3 | S3_TABLES | S3_KNOWLEDGE_BASE | SALESFORCE | SERVICENOW | SNOWFLAKE | SPARK | SPICE | SQLSERVER | TERADATA | TIMESTREAM | TWITTER | BIGQUERY | GOOGLE_ANALYTICS | TRINO | STARBURST | MONGO | MONGO_ATLAS | DOCUMENTDB | APPFLOW | IMPALA | GLUE | GOOGLE_DRIVE | CONFLUENCE | SHAREPOINT | ONE_DRIVE | WEB_CRAWLER | BOX`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcConnectionProperties`

Use this parameter only when you want Amazon Quick Sight to use a VPC connection when
connecting to your underlying source.

_Required_: No

_Type_: [VpcConnectionProperties](aws-properties-quicksight-datasource-vpcconnectionproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the dataset.

`CreatedTime`

The time that this data source was created.

`LastUpdatedTime`

The last time that this data source was updated.

`Status`

The HTTP status of the request.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ValueColumnConfiguration

AmazonElasticsearchParameters

All content copied from https://docs.aws.amazon.com/.
