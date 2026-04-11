This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Forecast::Dataset

Creates an Amazon Forecast dataset.

###### Important

Amazon Forecast is no longer available to new customers. Existing customers of
Amazon Forecast can continue to use the service as normal.
[Learn more"](https://aws.amazon.com/blogs/machine-learning/transition-your-amazon-forecast-usage-to-amazon-sagemaker-canvas)

The information about the dataset that you provide helps
Forecast understand how to consume the data for model training. This includes the
following:

- _`DataFrequency`_ \- How frequently your historical
time-series data is collected.

- _`Domain`_ and
_`DatasetType`_ \- Each dataset has an associated dataset
domain and a type within the domain. Amazon Forecast provides a list of predefined domains and
types within each domain. For each unique dataset domain and type within the domain,
Amazon Forecast requires your data to include a minimum set of predefined fields.

- _`Schema`_ \- A schema specifies the fields in the dataset,
including the field name and data type.

After creating a dataset, you import your training data into it and add the dataset to a
dataset group. You use the dataset group to create a predictor. For more information, see
[Importing datasets](../../../forecast/latest/dg/howitworks-datasets-groups.md).

To get a list of all your datasets, use the [ListDatasets](../../../forecast/latest/dg/api-listdatasets.md) operation.

For example Forecast datasets, see the [Amazon Forecast Sample GitHub\
repository](https://github.com/aws-samples/amazon-forecast-samples).

###### Note

The `Status` of a dataset must be `ACTIVE` before you can import
training data. Use the [DescribeDataset](../../../forecast/latest/dg/api-describedataset.md) operation to get
the status.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Forecast::Dataset",
  "Properties" : {
      "DataFrequency" : String,
      "DatasetName" : String,
      "DatasetType" : String,
      "Domain" : String,
      "EncryptionConfig" : EncryptionConfig,
      "Schema" : Schema,
      "Tags" : [ TagsItems, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Forecast::Dataset
Properties:
  DataFrequency: String
  DatasetName: String
  DatasetType: String
  Domain: String
  EncryptionConfig:
    EncryptionConfig
  Schema:
    Schema
  Tags:
    - TagsItems

```

## Properties

`DataFrequency`

The frequency of data collection. This parameter is required for RELATED\_TIME\_SERIES
datasets.

Valid intervals are an integer followed by Y (Year), M (Month), W (Week), D (Day), H (Hour), and min (Minute). For example,
"1D" indicates every day and "15min" indicates every 15 minutes. You cannot specify a value that would overlap with the next larger frequency. That means, for example, you cannot specify a frequency of 60 minutes, because that is equivalent to 1 hour. The valid values for each frequency are the following:

- Minute - 1-59

- Hour - 1-23

- Day - 1-6

- Week - 1-4

- Month - 1-11

- Year - 1

Thus, if you want every other week forecasts, specify "2W". Or, if you want quarterly forecasts, you specify "3M".

_Required_: No

_Type_: String

_Pattern_: `^Y|M|W|D|H|30min|15min|10min|5min|1min$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatasetName`

The name of the dataset.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z][a-zA-Z0-9_]*`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DatasetType`

The dataset type.

_Required_: Yes

_Type_: String

_Allowed values_: `TARGET_TIME_SERIES | RELATED_TIME_SERIES | ITEM_METADATA`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Domain`

The domain associated with the dataset.

_Required_: Yes

_Type_: String

_Allowed values_: `RETAIL | CUSTOM | INVENTORY_PLANNING | EC2_CAPACITY | WORK_FORCE | WEB_TRAFFIC | METRICS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionConfig`

A Key Management Service (KMS) key and the Identity and Access Management (IAM) role that Amazon Forecast can assume to access
the key.

_Required_: No

_Type_: [EncryptionConfig](aws-properties-forecast-dataset-encryptionconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Schema`

The schema for the dataset. The schema attributes and their order must match the fields in
your data. The dataset `Domain` and `DatasetType` that you choose
determine the minimum required fields in your training data. For information about the
required fields for a specific dataset domain and type, see [Dataset Domains and Dataset\
Types](../../../forecast/latest/dg/howitworks-domains-ds-types.md).

_Required_: Yes

_Type_: [Schema](aws-properties-forecast-dataset-schema.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [TagsItems](aws-properties-forecast-dataset-tagsitems.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the dataset.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Forecast

AttributesItems

All content copied from https://docs.aws.amazon.com/.
