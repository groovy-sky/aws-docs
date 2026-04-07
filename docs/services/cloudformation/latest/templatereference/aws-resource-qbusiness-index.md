This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::Index

Creates an Amazon Q Business index.

To determine if index creation has completed, check the `Status` field
returned from a call to `DescribeIndex`. The `Status` field is set
to `ACTIVE` when the index is ready to use.

Once the index is active, you can index your documents using the [`BatchPutDocument`](../../../amazonq/latest/api-reference/api-batchputdocument.md) API or the [`CreateDataSource`](../../../amazonq/latest/api-reference/api-createdatasource.md) API.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::QBusiness::Index",
  "Properties" : {
      "ApplicationId" : String,
      "CapacityConfiguration" : IndexCapacityConfiguration,
      "Description" : String,
      "DisplayName" : String,
      "DocumentAttributeConfigurations" : [ DocumentAttributeConfiguration, ... ],
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::QBusiness::Index
Properties:
  ApplicationId: String
  CapacityConfiguration:
    IndexCapacityConfiguration
  Description: String
  DisplayName: String
  DocumentAttributeConfigurations:
    - DocumentAttributeConfiguration
  Tags:
    - Tag
  Type: String

```

## Properties

`ApplicationId`

The identifier of the Amazon Q Business application using the index.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9-]{35}$`

_Minimum_: `36`

_Maximum_: `36`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CapacityConfiguration`

The capacity units you want to provision for your index. You can add and remove
capacity to fit your usage needs.

_Required_: No

_Type_: [IndexCapacityConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-index-indexcapacityconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description for the Amazon Q Business index.

_Required_: No

_Type_: String

_Pattern_: `^[\s\S]*$`

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DisplayName`

The name of the index.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9_-]*$`

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentAttributeConfigurations`

Configuration information for document attributes. Document attributes are metadata or
fields associated with your documents. For example, the company department name
associated with each document.

For more information, see [Understanding document attributes](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/doc-attributes.html).

_Required_: No

_Type_: Array of [DocumentAttributeConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-index-documentattributeconfiguration.html)

_Minimum_: `1`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of key-value pairs that identify or categorize the index. You can also use tags
to help control access to the index. Tag keys and values can consist of Unicode letters,
digits, white space, and any of the following symbols: \_ . : / = + - @.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-qbusiness-index-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The index type that's suitable for your needs. For more information on what's included
in each type of index, see [Amazon Q Business\
tiers](../../../amazonq/latest/qbusiness-ug/tiers.md#index-tiers).

_Required_: No

_Type_: String

_Allowed values_: `ENTERPRISE | STARTER`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the application ID and index ID. For example:

`{"Ref": "ApplicationId|IndexId"}`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The Unix timestamp when the index was created.

`IndexArn`

The Amazon Resource Name (ARN) of an Amazon Q Business index.

`IndexId`

The identifier for the index.

`Status`

The current status of the index. When the status is `ACTIVE`, the index is
ready.

`UpdatedAt`

The Unix timestamp when the index was last updated.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VideoExtractionConfiguration

DocumentAttributeConfiguration
