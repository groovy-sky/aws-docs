This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::DataSource

Creates a data source connector for an Amazon Q Business application.

`CreateDataSource` is a synchronous operation. The operation returns 200 if
the data source was successfully created. Otherwise, an exception is raised.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::QBusiness::DataSource",
  "Properties" : {
      "ApplicationId" : String,
      "Configuration" : ,
      "Description" : String,
      "DisplayName" : String,
      "DocumentEnrichmentConfiguration" : DocumentEnrichmentConfiguration,
      "IndexId" : String,
      "MediaExtractionConfiguration" : MediaExtractionConfiguration,
      "RoleArn" : String,
      "SyncSchedule" : String,
      "Tags" : [ Tag, ... ],
      "VpcConfiguration" : DataSourceVpcConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::QBusiness::DataSource
Properties:
  ApplicationId: String
  Configuration:

  Description: String
  DisplayName: String
  DocumentEnrichmentConfiguration:
    DocumentEnrichmentConfiguration
  IndexId: String
  MediaExtractionConfiguration:
    MediaExtractionConfiguration
  RoleArn: String
  SyncSchedule: String
  Tags:
    - Tag
  VpcConfiguration:
    DataSourceVpcConfiguration

```

## Properties

`ApplicationId`

The identifier of the Amazon Q Business application the data source will be attached
to.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9-]{35}$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Configuration`

Use this property to specify a JSON or YAML schema with configuration properties
specific to your data source connector to connect your data source repository to Amazon Q Business. You must use the JSON or YAML schema provided by Amazon Q.

The following links have the configuration properties and schemas for AWS CloudFormation for the following connectors:

- [Amazon Simple Storage Service](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/s3-cfn.html)

- [Amazon Q\
Web Crawler](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/web-crawler-cfn.html)

Similarly, you can find configuration templates and properties for your specific data
source using the following steps:

1. Navigate to the [Supported\
    connectors](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/connectors-list.html) page in the Amazon Q Business User Guide, and select the
    data source connector of your choice.

2. Then, from that specific data source connector's page, choose the topic
    containing **Using CloudFormation** to find the
    schemas for your data source connector, including configuration parameter
    descriptions and examples.

_Required_: Yes

_Type_:

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description for the data source connector.

_Required_: No

_Type_: String

_Pattern_: `^[\s\S]*$`

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The name of the Amazon Q Business data source.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9_-]*$`

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentEnrichmentConfiguration`

Provides the configuration information for altering document metadata and content
during the document ingestion process.

For more information, see [Custom document\
enrichment](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/custom-document-enrichment.html).

_Required_: No

_Type_: [DocumentEnrichmentConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-datasource-documentenrichmentconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IndexId`

The identifier of the index the data source is attached to.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9-]{35}$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MediaExtractionConfiguration`

The configuration for extracting information from media in documents.

_Required_: No

_Type_: [MediaExtractionConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-datasource-mediaextractionconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of an IAM role with permission to access
the data source and required resources. This field is required for all connector types except custom connectors, where it is optional.

_Required_: No

_Type_: String

_Pattern_: `^arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}$`

_Minimum_: `0`

_Maximum_: `1284`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SyncSchedule`

Sets the frequency for Amazon Q Business to check the documents in your data source
repository and update your index. If you don't set a schedule, Amazon Q Business won't
periodically update the index.

Specify a `cron-` format schedule string or an empty string to indicate
that the index is updated on demand. You can't specify the `Schedule`
parameter when the `Type` parameter is set to `CUSTOM`. If you do,
you receive a `ValidationException` exception.

_Required_: No

_Type_: String

_Pattern_: `^[\s\S]*$`

_Maximum_: `998`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of key-value pairs that identify or categorize the data source connector. You
can also use tags to help control access to the data source connector. Tag keys and
values can consist of Unicode letters, digits, white space, and any of the following
symbols: \_ . : / = + - @.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-datasource-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcConfiguration`

Configuration information for an Amazon VPC (Virtual Private Cloud) to connect
to your data source. For more information, see [Using\
Amazon VPC with Amazon Q Business connectors](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/connector-vpc.html).

_Required_: No

_Type_: [DataSourceVpcConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-datasource-datasourcevpcconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the application ID, data source ID, and index ID. For
example:

`{"Ref": "ApplicationId|DataSourceId|IndexId"}`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The Unix timestamp when the Amazon Q Business data source was created.

`DataSourceArn`

The Amazon Resource Name (ARN) of a data source in an Amazon Q Business application.

`DataSourceId`

The identifier of the Amazon Q Business data source.

`Status`

The status of the Amazon Q Business data source.

`Type`

The type of the Amazon Q Business data source.

`UpdatedAt`

The Unix timestamp when the Amazon Q Business data source was last updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AudioExtractionConfiguration
