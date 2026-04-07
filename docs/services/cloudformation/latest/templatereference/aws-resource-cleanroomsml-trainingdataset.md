This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRoomsML::TrainingDataset

Defines the information necessary to create a training dataset. In Clean Rooms ML, the
`TrainingDataset` is metadata that points to a Glue table, which is read only
during `AudienceModel` creation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CleanRoomsML::TrainingDataset",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "RoleArn" : String,
      "Tags" : [ Tag, ... ],
      "TrainingData" : [ Dataset, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CleanRoomsML::TrainingDataset
Properties:
  Description: String
  Name: String
  RoleArn: String
  Tags:
    - Tag
  TrainingData:
    - Dataset

```

## Properties

`Description`

The description of the training dataset.

_Required_: No

_Type_: String

_Pattern_: `^[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDBFF-\uDC00\uDFFF\t\r\n]*$`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the training dataset.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!\s*$)[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDBFF-\uDC00\uDFFF\t]*$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The ARN of the IAM role that Clean Rooms ML can assume to read the data referred to in the `dataSource` field of each dataset.

Passing a role across accounts is not allowed. If you pass a role that isn't in your account, you get an `AccessDeniedException` error.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z]*:iam::[0-9]{12}:role/.+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The optional metadata that you apply to the resource to help you categorize and organize them. Each tag consists of a key and an optional value, both of which you define.

The following basic restrictions apply to tags:

- Maximum number of tags per resource - 50.

- For each resource, each tag key must be unique, and each tag key can have only one value.

- Maximum key length - 128 Unicode characters in UTF-8.

- Maximum value length - 256 Unicode characters in UTF-8.

- If your tagging schema is used across multiple services and resources, remember that other services may have restrictions on allowed characters. Generally allowed characters are: letters, numbers, and spaces representable in UTF-8, and the following characters: + - = . \_ : / @.

- Tag keys and values are case sensitive.

- Do not use `aws:`, `AWS:`, or any upper or lowercase combination of such as a prefix for keys as it is reserved. You cannot edit or delete tag keys with this prefix. Values can have this prefix. If a tag value has `aws` as its prefix but the key does not, then Clean Rooms ML considers it to be a user tag and will count against the limit of 50 tags. Tags with only the key prefix of `aws` do not count against your tags per resource limit.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanroomsml-trainingdataset-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrainingData`

An array of information that lists the Dataset objects, which specifies the dataset type
and details on its location and schema. You must provide a role that has read access to
these tables.

_Required_: Yes

_Type_: Array of [Dataset](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanroomsml-trainingdataset-dataset.html)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource ARN. For example:

`{ "Ref": "myTrainingDataset" }`

For the Clean Rooms ML training dataset, `Ref` returns the ARN of the training dataset.

Example: `arn:aws:cleanrooms-ml:ap-northeast-1:891377082322:training-dataset/fR8doOMxlv5q5HD5qB0f68`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Status`

The status of the training dataset.

`TrainingDatasetArn`

The Amazon Resource Name (ARN) of the training dataset.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TrainedModelsConfigurationPolicy

ColumnSchema
