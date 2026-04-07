This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::ImageRecipe

Creates a new image recipe. Image recipes define how images are configured, tested,
and assessed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ImageBuilder::ImageRecipe",
  "Properties" : {
      "AdditionalInstanceConfiguration" : AdditionalInstanceConfiguration,
      "AmiTags" : {Key: Value, ...},
      "BlockDeviceMappings" : [ InstanceBlockDeviceMapping, ... ],
      "Components" : [ ComponentConfiguration, ... ],
      "Description" : String,
      "Name" : String,
      "ParentImage" : String,
      "Tags" : {Key: Value, ...},
      "Version" : String,
      "WorkingDirectory" : String
    }
}

```

### YAML

```yaml

Type: AWS::ImageBuilder::ImageRecipe
Properties:
  AdditionalInstanceConfiguration:
    AdditionalInstanceConfiguration
  AmiTags:
    Key: Value
  BlockDeviceMappings:
    - InstanceBlockDeviceMapping
  Components:
    - ComponentConfiguration
  Description: String
  Name: String
  ParentImage: String
  Tags:
    Key: Value
  Version: String
  WorkingDirectory: String

```

## Properties

`AdditionalInstanceConfiguration`

Before you create a new AMI, Image Builder launches temporary Amazon EC2 instances to build and test
your image configuration. Instance configuration adds a layer of control over those
instances. You can define settings and add scripts to run when an instance is launched
from your AMI.

_Required_: No

_Type_: [AdditionalInstanceConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-imagerecipe-additionalinstanceconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AmiTags`

Tags that are applied to the AMI that Image Builder creates during the Build phase
prior to image distribution.

_Required_: No

_Type_: Object of String

_Pattern_: `.{1,}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BlockDeviceMappings`

The block device mappings to apply when creating images from this recipe.

_Required_: No

_Type_: Array of [InstanceBlockDeviceMapping](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-imagerecipe-instanceblockdevicemapping.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Components`

The components that are included in the image recipe. Recipes require a minimum of one build component, and can
have a maximum of 20 build and test components in any combination.

_Required_: No

_Type_: Array of [ComponentConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-imagerecipe-componentconfiguration.html)

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of the image recipe.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the image recipe.

_Required_: Yes

_Type_: String

_Pattern_: `^[-_A-Za-z-0-9][-_A-Za-z0-9 ]{1,126}[-_A-Za-z-0-9]$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ParentImage`

The base image for customizations specified in the image recipe. You can specify the
parent image using one of the following options:

- AMI ID

- Image Builder image Amazon Resource Name (ARN)

- AWS Systems Manager (SSM) Parameter Store Parameter, prefixed by `ssm:`,
followed by the parameter name or ARN.

- AWS Marketplace product ID

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags of the image recipe.

_Required_: No

_Type_: Object of String

_Pattern_: `.{1,}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

The semantic version of the image recipe. This version follows the semantic version
syntax.

###### Note

The semantic version has four nodes: <major>.<minor>.<patch>/<build>.
You can assign values for the first three, and can filter on all of them.

**Assignment:** For the first three nodes you can assign any positive integer value, including
zero, with an upper limit of 2^30-1, or 1073741823 for each node. Image Builder automatically assigns the
build number to the fourth node.

**Patterns:** You can use any numeric pattern that adheres to the assignment requirements for
the nodes that you can assign. For example, you might choose a software version pattern, such as 1.0.0, or
a date, such as 2021.01.01.

_Required_: Yes

_Type_: String

_Pattern_: `^(?:[0-9]+|x)\.(?:[0-9]+|x)\.(?:[0-9]+|x)$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`WorkingDirectory`

The working directory to be used during build and test workflows.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource ARN, such as
`arn:aws:imagebuilder:us-east-1:123456789012:image-recipe/mybasicrecipe/2019.12.03`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the Amazon Resource Name (ARN) of the image recipe. For example,
`arn:aws:imagebuilder:us-east-1:123456789012:image-recipe/mybasicrecipe/2019.12.03`.

`LatestVersion.Arn`

The latest version Amazon Resource Name (ARN) of the Image Builder resource.

`LatestVersion.Major`

The latest version Amazon Resource Name (ARN) with the same `major` version of the Image Builder resource.

`LatestVersion.Minor`

The latest version Amazon Resource Name (ARN) with the same `minor` version of the Image Builder resource.

`LatestVersion.Patch`

The latest version Amazon Resource Name (ARN) with the same `patch` version of the Image Builder resource.

`Name`

The name of the image recipe.

## Examples

### Create an image recipe

The following example shows the schema for all of the parameters of the
ImageRecipe resource document in both YAML and JSON format .

#### YAML

```yaml

Resources:
  ImageRecipeAllParameters:
    Type: 'AWS::ImageBuilder::ImageRecipe'
    Properties:
      Name: 'image-recipe-name'
      Version: '1.0.0'
      ParentImage: !Ref ParentImage
      Description: 'description'
      Components:
        - ComponentArn: !Ref ComponentArn
        - ComponentArn: !Ref AnotherComponentArn
      BlockDeviceMappings:
        - DeviceName: "device-name"
          VirtualName: "virtual-name"
          Ebs:
            DeleteOnTermination: true
            Encrypted: true
            Iops: 100
            KmsKeyId: !Ref KmsKeyId
            SnapshotId: "snapshot-id"
            VolumeType: "gp2"
            VolumeSize: 100
      Tags:
        CustomerImageRecipeTagKey1: 'CustomerImageRecipeTagValue1'
        CustomerImageRecipeTagKey2: 'CustomerImageRecipeTagValue2'
```

#### JSON

```json

{
    "Resources": {
        "ImageRecipeAllParameters": {
            "Type": "AWS::ImageBuilder::ImageRecipe",
            "Properties": {
                "Name": "image-recipe-name",
                "Version": "1.0.0",
                "ParentImage": {
                    "Ref": "ParentImage"
                },
                "Description": "description",
                "Components": [
                    {
                        "ComponentArn": {
                            "Ref": "ComponentArn"
                        }
                    },
                    {
                        "ComponentArn": {
                            "Ref": "AnotherComponentArn"
                        }
                    }
                ],
                "BlockDeviceMappings": [
                    {
                        "DeviceName": "device-name",
                        "VirtualName": "virtual-name",
                        "Ebs": {
                            "DeleteOnTermination": true,
                            "Encrypted": true,
                            "Iops": 100,
                            "KmsKeyId": {
                                "Ref": "KmsKeyId"
                            },
                            "SnapshotId": "snapshot-id",
                            "VolumeType": "gp2",
                            "VolumeSize": 100
                        }
                    }
                ],
                "Tags": {
                    "CustomerImageRecipeTagKey1": "CustomerImageRecipeTagValue1",
                    "CustomerImageRecipeTagKey2": "CustomerImageRecipeTagValue2"
                }
            }
        }
    }
}
```

## See also

- [Create a basic image recipe](https://docs.aws.amazon.com/imagebuilder/latest/userguide/managing-image-builder-cli.html#image-builder-cli-create-recipe) in the
_Image Builder User Guide_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

WorkflowParameter

AdditionalInstanceConfiguration
