This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Forecast::DatasetGroup

Creates a dataset group, which holds a collection of related datasets. You can add
datasets to the dataset group when you create the dataset group, or later by using the [UpdateDatasetGroup](../../../forecast/latest/dg/api-updatedatasetgroup.md) operation.

###### Important

Amazon Forecast is no longer available to new customers. Existing customers of
Amazon Forecast can continue to use the service as normal.
[Learn more"](https://aws.amazon.com/blogs/machine-learning/transition-your-amazon-forecast-usage-to-amazon-sagemaker-canvas)

After creating a dataset group and adding datasets, you use the dataset group when you
create a predictor. For more information, see [Dataset groups](../../../forecast/latest/dg/howitworks-datasets-groups.md).

To get a list of all your datasets groups, use the [ListDatasetGroups](../../../forecast/latest/dg/api-listdatasetgroups.md)
operation.

###### Note

The `Status` of a dataset group must be `ACTIVE` before you can
use the dataset group to create a predictor. To get the status, use the [DescribeDatasetGroup](../../../forecast/latest/dg/api-describedatasetgroup.md) operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Forecast::DatasetGroup",
  "Properties" : {
      "DatasetArns" : [ String, ... ],
      "DatasetGroupName" : String,
      "Domain" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Forecast::DatasetGroup
Properties:
  DatasetArns:
    - String
  DatasetGroupName: String
  Domain: String
  Tags:
    - Tag

```

## Properties

`DatasetArns`

An array of Amazon Resource Names (ARNs) of the datasets that you want to include in the
dataset group.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatasetGroupName`

The name of the dataset group.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z][a-zA-Z0-9_]*`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Domain`

The domain associated with the dataset group. When you add a dataset to a dataset group,
this value and the value specified for the `Domain` parameter of the [CreateDataset](../../../forecast/latest/dg/api-createdataset.md)
operation must match.

The `Domain` and `DatasetType` that you choose determine the fields
that must be present in training data that you import to a dataset. For example, if you choose
the `RETAIL` domain and `TARGET_TIME_SERIES` as the
`DatasetType`, Amazon Forecast requires that `item_id`,
`timestamp`, and `demand` fields are present in your data. For more
information, see [Dataset groups](../../../forecast/latest/dg/howitworks-datasets-groups.md).

_Required_: Yes

_Type_: String

_Allowed values_: `RETAIL | CUSTOM | INVENTORY_PLANNING | EC2_CAPACITY | WORK_FORCE | WEB_TRAFFIC | METRICS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-forecast-datasetgroup-tag.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DatasetGroupArn`

The Amazon Resource Name (ARN) of the dataset group.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagsItems

Tag

All content copied from https://docs.aws.amazon.com/.
