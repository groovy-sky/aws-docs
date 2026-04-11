This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Comprehend::Flywheel

A flywheel is an AWS resource that orchestrates the ongoing training of a model for custom classification
or custom entity recognition. You can create a flywheel to start with an existing trained model, or
Comprehend can create and train a new model.

When you create the flywheel, Comprehend creates a data lake in your account. The data lake holds the training
data and test data for all versions of the model.

To use a flywheel with an existing trained model, you specify the active model version. Comprehend copies the model's
training data and test data into the flywheel's data lake.

To use the flywheel with a new model, you need to provide a dataset for training data (and optional test data)
when you create the flywheel.

For more information about flywheels, see [Flywheel overview](../../../comprehend/latest/dg/flywheels-about.md) in the _Amazon Comprehend Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Comprehend::Flywheel",
  "Properties" : {
      "ActiveModelArn" : String,
      "DataAccessRoleArn" : String,
      "DataLakeS3Uri" : String,
      "DataSecurityConfig" : DataSecurityConfig,
      "FlywheelName" : String,
      "ModelType" : String,
      "Tags" : [ Tag, ... ],
      "TaskConfig" : TaskConfig
    }
}

```

### YAML

```yaml

Type: AWS::Comprehend::Flywheel
Properties:
  ActiveModelArn: String
  DataAccessRoleArn: String
  DataLakeS3Uri: String
  DataSecurityConfig:
    DataSecurityConfig
  FlywheelName: String
  ModelType: String
  Tags:
    - Tag
  TaskConfig:
    TaskConfig

```

## Properties

`ActiveModelArn`

The Amazon Resource Number (ARN) of the active model version.

_Required_: No

_Type_: String

_Pattern_: `arn:aws(-[^:]+)?:comprehend:[a-zA-Z0-9-]*:[0-9]{12}:(document-classifier|entity-recognizer)/[a-zA-Z0-9](-*[a-zA-Z0-9])*(/version/[a-zA-Z0-9](-*[a-zA-Z0-9])*)?`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataAccessRoleArn`

The Amazon Resource Name (ARN) of the IAM role that
grants Amazon Comprehend permission to access the flywheel data.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws(-[^:]+)?:iam::[0-9]{12}:role/.+`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataLakeS3Uri`

Amazon S3 URI of the data lake location.

_Required_: Yes

_Type_: String

_Pattern_: `s3://[a-z0-9][\.\-a-z0-9]{1,61}[a-z0-9](/.*)?`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataSecurityConfig`

Data security configuration.

_Required_: No

_Type_: [DataSecurityConfig](aws-properties-comprehend-flywheel-datasecurityconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FlywheelName`

Name for the flywheel.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9])*$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModelType`

Model type of the flywheel's model.

_Required_: No

_Type_: String

_Allowed values_: `DOCUMENT_CLASSIFIER | ENTITY_RECOGNIZER`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Tags associated with the endpoint being created. A tag is a key-value pair that adds
metadata to the endpoint. For example, a tag with "Sales" as the key might be added to an
endpoint to indicate its use by the sales department.

_Required_: No

_Type_: Array of [Tag](aws-properties-comprehend-flywheel-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TaskConfig`

Configuration about the model associated with a flywheel.

_Required_: No

_Type_: [TaskConfig](aws-properties-comprehend-flywheel-taskconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the flywheel.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the flywheel.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VpcConfig

DataSecurityConfig

All content copied from https://docs.aws.amazon.com/.
