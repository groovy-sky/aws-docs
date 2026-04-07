This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::DataSource

###### Important

Properties with `__Update requires: Replacement__` can result in the
creation of a new data source and deletion of the old one. This can happen if you
also change the Name of the data source.

Specifies a data source as a resource in a top-level template. Minimally, you must
specify the following properties:

- Name – Specify a name for the data source.

- KnowledgeBaseId – Specify the ID of the knowledge base for the data source to
belong to.

- DataSourceConfiguration – Specify information about the Amazon S3
bucket containing the data source. The following sub-properties are
required:

- Type – Specify the value `S3`.

For more information about setting up data sources in Amazon Bedrock, see [Set up a\
data source for your knowledge base](https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-base-ds.html).

See the **Properties** section below for descriptions of
both the required and optional properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Bedrock::DataSource",
  "Properties" : {
      "DataDeletionPolicy" : String,
      "DataSourceConfiguration" : DataSourceConfiguration,
      "Description" : String,
      "KnowledgeBaseId" : String,
      "Name" : String,
      "ServerSideEncryptionConfiguration" : ServerSideEncryptionConfiguration,
      "VectorIngestionConfiguration" : VectorIngestionConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::Bedrock::DataSource
Properties:
  DataDeletionPolicy: String
  DataSourceConfiguration:
    DataSourceConfiguration
  Description: String
  KnowledgeBaseId: String
  Name: String
  ServerSideEncryptionConfiguration:
    ServerSideEncryptionConfiguration
  VectorIngestionConfiguration:
    VectorIngestionConfiguration

```

## Properties

`DataDeletionPolicy`

The data deletion policy for the data source.

_Required_: No

_Type_: String

_Allowed values_: `RETAIN | DELETE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataSourceConfiguration`

The connection configuration for the data source.

_Required_: Yes

_Type_: [DataSourceConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-datasource-datasourceconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the data source.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KnowledgeBaseId`

The unique identifier of the knowledge base to which the data source belongs.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-zA-Z]{10}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the data source.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9a-zA-Z][_-]?){1,100}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerSideEncryptionConfiguration`

Contains details about the configuration of the server-side encryption.

_Required_: No

_Type_: [ServerSideEncryptionConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-datasource-serversideencryptionconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VectorIngestionConfiguration`

Contains details about how to ingest the documents in the data source.

_Required_: No

_Type_: [VectorIngestionConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-datasource-vectoringestionconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the knowledge base ID and the data source ID, separated
by a pipe ( `|`).

For example, `{ "Ref": "myDataSource" }` could return the value
`"KB12345678|DS12345678"`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The time at which the data source was created.

`DataSourceConfiguration.WebConfiguration.CrawlerConfiguration.UserAgentHeader`

Property description not available.

`DataSourceId`

The unique identifier of the data source.

`DataSourceStatus`

The status of the data source. The following statuses are possible:

- Available – The data source has been created and is ready for ingestion into the knowledge base.

- Deleting – The data source is being deleted.

`FailureReasons`

The detailed reasons on the failure to delete a data source.

`UpdatedAt`

The time at which the data source was last updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VideoStandardOutputConfiguration

BedrockDataAutomationConfiguration
