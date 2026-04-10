This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRoomsML::ConfiguredModelAlgorithm

Creates a configured model algorithm using a container image stored in an ECR
repository.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CleanRoomsML::ConfiguredModelAlgorithm",
  "Properties" : {
      "Description" : String,
      "InferenceContainerConfig" : InferenceContainerConfig,
      "KmsKeyArn" : String,
      "Name" : String,
      "RoleArn" : String,
      "Tags" : [ Tag, ... ],
      "TrainingContainerConfig" : ContainerConfig
    }
}

```

### YAML

```yaml

Type: AWS::CleanRoomsML::ConfiguredModelAlgorithm
Properties:
  Description: String
  InferenceContainerConfig:
    InferenceContainerConfig
  KmsKeyArn: String
  Name: String
  RoleArn: String
  Tags:
    - Tag
  TrainingContainerConfig:
    ContainerConfig

```

## Properties

`Description`

The description of the configured model algorithm.

_Required_: No

_Type_: String

_Pattern_: `^[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDBFF-\uDC00\uDFFF\t\r\n]*$`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InferenceContainerConfig`

Provides configuration information for the inference container that is used when you run an inference job on a configured model algorithm.

_Required_: No

_Type_: [InferenceContainerConfig](aws-properties-cleanroomsml-configuredmodelalgorithm-inferencecontainerconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKeyArn`

The Amazon Resource Name (ARN) of the KMS key. This key is used to encrypt and decrypt customer-owned data in the configured ML model algorithm and associated data.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[-a-z]*:kms:[-a-z0-9]+:[0-9]{12}:key/.+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the configured model algorithm.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!\s*$)[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDBFF-\uDC00\uDFFF\t]*$`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The Amazon Resource Name (ARN) of the role that is used to access the repository.

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

_Type_: Array of [Tag](aws-properties-cleanroomsml-configuredmodelalgorithm-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrainingContainerConfig`

Provides configuration information for the training container, including entrypoints and arguments.

_Required_: No

_Type_: [ContainerConfig](aws-properties-cleanroomsml-configuredmodelalgorithm-containerconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource ARN. For example:

`{ "Ref": "myConfiguredModelAlgorithm" }`

For the Clean Rooms ML configured model algorithm, `Ref` returns the ARN of the configured model algorithm.

Example: `arn:aws:cleanrooms-ml:us-east-1:123456789012:configured-model-algorithm/a1b2c3d4e5f6`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ConfiguredModelAlgorithmArn`

The Amazon Resource Name (ARN) of the configured model algorithm.

## Examples

### Create a configured model algorithm

The following example creates a configured model algorithm with training and inference container configurations.

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Resources": {
    "MyConfiguredModelAlgorithm": {
      "Type": "AWS::CleanRoomsML::ConfiguredModelAlgorithm",
      "Properties": {
        "Name": "MyMLAlgorithm",
        "Description": "A configured model algorithm for collaborative ML",
        "RoleArn": "arn:aws:iam::123456789012:role/CleanRoomsMLServiceRole",
        "TrainingContainerConfig": {
          "ImageUri": "123456789012.dkr.ecr.us-east-1.amazonaws.com/my-training-image:latest",
          "MetricDefinitions": [
            {
              "Name": "loss",
              "Regex": "Loss: ([0-9\\\\.]+)"
            }
          ]
        },
        "InferenceContainerConfig": {
          "ImageUri": "123456789012.dkr.ecr.us-east-1.amazonaws.com/my-inference-image:latest"
        },
        "Tags": [
          {
            "Key": "Environment",
            "Value": "Production"
          }
        ]
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Resources:
  MyConfiguredModelAlgorithm:
    Type: AWS::CleanRoomsML::ConfiguredModelAlgorithm
    Properties:
      Name: MyMLAlgorithm
      Description: A configured model algorithm for collaborative ML
      RoleArn: arn:aws:iam::123456789012:role/CleanRoomsMLServiceRole
      TrainingContainerConfig:
        ImageUri: 123456789012.dkr.ecr.us-east-1.amazonaws.com/my-training-image:latest
        MetricDefinitions:
          - Name: loss
            Regex: 'Loss: ([0-9\.]+)'
      InferenceContainerConfig:
        ImageUri: 123456789012.dkr.ecr.us-east-1.amazonaws.com/my-inference-image:latest
      Tags:
        - Key: Environment
          Value: Production
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CleanRoomsML

ContainerConfig

All content copied from https://docs.aws.amazon.com/.
