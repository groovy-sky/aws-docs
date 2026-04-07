This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow SAPODataSourceProperties

The properties that are applied when using SAPOData as a flow source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ObjectPath" : String,
  "paginationConfig" : SAPODataPaginationConfig,
  "parallelismConfig" : SAPODataParallelismConfig
}

```

### YAML

```yaml

  ObjectPath: String
  paginationConfig:
    SAPODataPaginationConfig
  parallelismConfig:
    SAPODataParallelismConfig

```

## Properties

`ObjectPath`

The object path specified in the SAPOData flow source.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`paginationConfig`

Sets the page size for each concurrent process that transfers OData records from your SAP
instance.

_Required_: No

_Type_: [SAPODataPaginationConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-sapodatapaginationconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`parallelismConfig`

Sets the number of concurrent processes that transfers OData records from your SAP
instance.

_Required_: No

_Type_: [SAPODataParallelismConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-flow-sapodataparallelismconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SAPODataParallelismConfig

ScheduledTriggerProperties
