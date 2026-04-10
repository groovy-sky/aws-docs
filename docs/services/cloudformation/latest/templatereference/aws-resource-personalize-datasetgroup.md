This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Personalize::DatasetGroup

A dataset group is a collection of related datasets (Item interactions,
Users, Items, Actions, Action interactions). You create a dataset group by calling [CreateDatasetGroup](../../../personalize/latest/dg/api-createdatasetgroup.md). You then create a dataset and add it to a
dataset group by calling [CreateDataset](../../../personalize/latest/dg/api-createdataset.md). The dataset group is used to create and train a
solution by calling [CreateSolution](../../../personalize/latest/dg/api-createsolution.md). A dataset group can contain only one of each
type of dataset.

You can specify an AWS Key Management Service (KMS) key to encrypt the datasets in
the group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Personalize::DatasetGroup",
  "Properties" : {
      "Domain" : String,
      "KmsKeyArn" : String,
      "Name" : String,
      "RoleArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::Personalize::DatasetGroup
Properties:
  Domain: String
  KmsKeyArn: String
  Name: String
  RoleArn: String

```

## Properties

`Domain`

The domain of a Domain dataset group.

_Required_: No

_Type_: String

_Allowed values_: `ECOMMERCE | VIDEO_ON_DEMAND`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKeyArn`

The Amazon Resource Name (ARN) of the AWS Key Management Service (KMS) key used to
encrypt the datasets.

_Required_: No

_Type_: String

_Pattern_: `arn:aws.*:kms:.*:[0-9]{12}:key/.*`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the dataset group.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9\-_]*`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The ARN of the AWS Identity and Access Management (IAM) role that has permissions to access
the AWS Key Management Service (KMS) key. Supplying an IAM role is only valid when also
specifying a KMS key.

_Required_: No

_Type_: String

_Pattern_: `arn:([a-z\d-]+):iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the resource.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DatasetGroupArn`

The Amazon Resource Name (ARN) of the dataset group.

## Examples

### Creating a dataset group

The following example creates an Amazon Personalize dataset group.

#### JSON

```json

{
   "AWSTemplateFormatVersion":"2010-09-09",
   "Resources":{
      "MyDatasetGroup": {
            "Type": "AWS::Personalize::DatasetGroup",
            "Properties": {
               "Name": "my-dataset-group-name"
            }
      }
   }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  MyDatasetGroup:
    Type: 'AWS::Personalize::DatasetGroup'
    Properties:
      Name: my-dataset-group-name
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSource

AWS::Personalize::Schema

All content copied from https://docs.aws.amazon.com/.
