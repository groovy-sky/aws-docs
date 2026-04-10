This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource

Creates a data source connector that you want to use with an Amazon Kendra
index.

You specify a name, data source connector type and description for your data source.
You also specify configuration information for the data source connector.

###### Important

`CreateDataSource` does _not_ support connectors
which [require a\
`TemplateConfiguration` object](../../../kendra/latest/dg/ds-schemas.md) for connecting to Amazon Kendra.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Kendra::DataSource",
  "Properties" : {
      "CustomDocumentEnrichmentConfiguration" : CustomDocumentEnrichmentConfiguration,
      "DataSourceConfiguration" : DataSourceConfiguration,
      "Description" : String,
      "IndexId" : String,
      "LanguageCode" : String,
      "Name" : String,
      "RoleArn" : String,
      "Schedule" : String,
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::Kendra::DataSource
Properties:
  CustomDocumentEnrichmentConfiguration:
    CustomDocumentEnrichmentConfiguration
  DataSourceConfiguration:
    DataSourceConfiguration
  Description: String
  IndexId: String
  LanguageCode: String
  Name: String
  RoleArn: String
  Schedule: String
  Tags:
    - Tag
  Type: String

```

## Properties

`CustomDocumentEnrichmentConfiguration`

Configuration information for altering document metadata and content during the
document ingestion process.

_Required_: No

_Type_: [CustomDocumentEnrichmentConfiguration](aws-properties-kendra-datasource-customdocumentenrichmentconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataSourceConfiguration`

Configuration information for an Amazon Kendra data source. The contents of the
configuration depend on the type of data source. You can only specify one type of data
source in the configuration.

You can't specify the `Configuration` parameter when the `Type`
parameter is set to `CUSTOM`.

The `Configuration` parameter is required for all other data
sources.

_Required_: No

_Type_: [DataSourceConfiguration](aws-properties-kendra-datasource-datasourceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description for the data source connector.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IndexId`

The identifier of the index you want to use with the data source connector.

_Required_: Yes

_Type_: String

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LanguageCode`

The code for a language. This shows a supported language for all documents
in the data source. English is supported by default.
For more information on supported languages, including their codes,
see [Adding \
documents in languages other than English](../../../kendra/latest/dg/in-adding-languages.md).

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z-]*`

_Minimum_: `2`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the data source.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of a role with permission to access the data
source.

You can't specify the `RoleArn` parameter when the `Type`
parameter is set to `CUSTOM`.

The `RoleArn` parameter is required for all other data sources.

_Required_: No

_Type_: String

_Pattern_: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

_Minimum_: `1`

_Maximum_: `1284`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Schedule`

Sets the frequency that Amazon Kendra checks the documents in your data source and
updates the index. If you don't set a schedule, Amazon Kendra doesn't periodically
update the index.

_Required_: No

_Type_: String

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-kendra-datasource-tag.md)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the data source.

_Required_: Yes

_Type_: String

_Allowed values_: `S3 | SHAREPOINT | SALESFORCE | ONEDRIVE | SERVICENOW | DATABASE | CUSTOM | CONFLUENCE | GOOGLEDRIVE | WEBCRAWLER | WORKDOCS | TEMPLATE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the data source ID. For example:

`{ "Ref": "<data source ID>|<index ID>" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the data source. For example:

`arn:aws:kendra:us-west-2:111122223333:index/335c3741-41df-46a6-b5d3-61f85b787884/data-source/b8cae438-6787-4091-8897-684a652bbb0a`

`Id`

The identifier for the data source. For example:

`b8cae438-6787-4091-8897-684a652bbb0a`.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Kendra

AccessControlListConfiguration

All content copied from https://docs.aws.amazon.com/.
