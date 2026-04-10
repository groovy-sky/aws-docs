This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::Dataset

Creates a dataset to connect an external datasource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTSiteWise::Dataset",
  "Properties" : {
      "DatasetDescription" : String,
      "DatasetName" : String,
      "DatasetSource" : DatasetSource,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoTSiteWise::Dataset
Properties:
  DatasetDescription: String
  DatasetName: String
  DatasetSource:
    DatasetSource
  Tags:
    - Tag

```

## Properties

`DatasetDescription`

A description about the dataset, and its functionality.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatasetName`

The name of the dataset.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatasetSource`

The data source for the dataset.

_Required_: Yes

_Type_: [DatasetSource](aws-properties-iotsitewise-dataset-datasetsource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of key-value pairs that contain metadata for the access policy. For more
information, see [Tagging your\
AWS IoT SiteWise resources](../../../iot-sitewise/latest/userguide/tag-resources.md) in the _AWS IoT SiteWise User Guide_.

_Required_: No

_Type_: Array of [Tag](aws-properties-iotsitewise-dataset-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `DatasetId`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DatasetArn`

The ARN of the dataset, which has the following format.

`arn:${Partition}:iotsitewise:${Region}:${Account}:dataset/${DatasetId}`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

`DatasetId`

The ID of the dataset.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

DatasetSource

All content copied from https://docs.aws.amazon.com/.
