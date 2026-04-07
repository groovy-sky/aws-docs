This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::Image

Creates a new image. This request will create a new image along with all of the
configured output resources defined in the distribution configuration. You must specify
exactly one recipe for your image, using either a ContainerRecipeArn or an
ImageRecipeArn.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ImageBuilder::Image",
  "Properties" : {
      "ContainerRecipeArn" : String,
      "DeletionSettings" : DeletionSettings,
      "DistributionConfigurationArn" : String,
      "EnhancedImageMetadataEnabled" : Boolean,
      "ExecutionRole" : String,
      "ImagePipelineExecutionSettings" : ImagePipelineExecutionSettings,
      "ImageRecipeArn" : String,
      "ImageScanningConfiguration" : ImageScanningConfiguration,
      "ImageTestsConfiguration" : ImageTestsConfiguration,
      "InfrastructureConfigurationArn" : String,
      "LoggingConfiguration" : ImageLoggingConfiguration,
      "Tags" : {Key: Value, ...},
      "Workflows" : [ WorkflowConfiguration, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ImageBuilder::Image
Properties:
  ContainerRecipeArn: String
  DeletionSettings:
    DeletionSettings
  DistributionConfigurationArn: String
  EnhancedImageMetadataEnabled: Boolean
  ExecutionRole: String
  ImagePipelineExecutionSettings:
    ImagePipelineExecutionSettings
  ImageRecipeArn: String
  ImageScanningConfiguration:
    ImageScanningConfiguration
  ImageTestsConfiguration:
    ImageTestsConfiguration
  InfrastructureConfigurationArn: String
  LoggingConfiguration:
    ImageLoggingConfiguration
  Tags:
    Key: Value
  Workflows:
    - WorkflowConfiguration

```

## Properties

`ContainerRecipeArn`

The Amazon Resource Name (ARN) of the container recipe that defines how images are configured and tested.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DeletionSettings`

Enables deletion of underlying resources of an image when it is replaced or deleted, including its Amazon Machine Images (AMIs), snapshots, or containers.

_Required_: No

_Type_: [DeletionSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-image-deletionsettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DistributionConfigurationArn`

The Amazon Resource Name (ARN) of the distribution configuration that defines and configures the outputs of your pipeline.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnhancedImageMetadataEnabled`

Collects additional information about the image being created, including the operating system (OS) version and package list. This information is used to enhance the overall experience of using EC2 Image Builder. Enabled by default.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExecutionRole`

The name or Amazon Resource Name (ARN) for the IAM role you create that grants
Image Builder access to perform workflow actions.

_Required_: No

_Type_: String

_Pattern_: `^(?:arn:aws(?:-[a-z]+)*:iam::[0-9]{12}:role/)?[a-zA-Z_0-9+=,.@\-_/]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImagePipelineExecutionSettings`

The image pipeline execution settings of the image.

_Required_: No

_Type_: [ImagePipelineExecutionSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-image-imagepipelineexecutionsettings.html)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`ImageRecipeArn`

The Amazon Resource Name (ARN) of the image recipe that defines how images are configured, tested, and assessed.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ImageScanningConfiguration`

Contains settings for vulnerability scans.

_Required_: No

_Type_: [ImageScanningConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-image-imagescanningconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ImageTestsConfiguration`

The image tests configuration of the image.

_Required_: No

_Type_: [ImageTestsConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-image-imagetestsconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`InfrastructureConfigurationArn`

The Amazon Resource Name (ARN) of the infrastructure configuration that defines the environment in which your image will be built and tested.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LoggingConfiguration`

The logging configuration that's defined for the image. Image Builder uses the defined settings
to direct execution log output during image creation.

_Required_: No

_Type_: [ImageLoggingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-image-imageloggingconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags of the image.

_Required_: No

_Type_: Object of String

_Pattern_: `.{1,}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Workflows`

Contains an array of workflow configuration objects.

_Required_: No

_Type_: Array of [WorkflowConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-image-workflowconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource ARN, such as
`arn:aws:imagebuilder:us-west-2:123456789012:image/my-example-image`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the Amazon Resource Name (ARN) of the image. For example,
`arn:aws:imagebuilder:us-west-2:123456789012:image/mybasicrecipe/2019.12.03/1`.

`ImageId`

Returns the AMI ID of the Amazon EC2 AMI in the Region in which you are using Image
Builder. Values are returned only for AMIs, and not for container images.

`ImageUri`

Returns the URI for a container image created in the context Region. Values are returned only for
container images, and not for AMIs.

`LatestVersion.Arn`

The latest version Amazon Resource Name (ARN) of the Image Builder resource.

`LatestVersion.Major`

The latest version Amazon Resource Name (ARN) with the same `major` version of the Image Builder resource.

`LatestVersion.Minor`

The latest version Amazon Resource Name (ARN) with the same `minor` version of the Image Builder resource.

`LatestVersion.Patch`

The latest version Amazon Resource Name (ARN) with the same `patch` version of the Image Builder resource.

`Name`

Returns the name of the image.

## Examples

### Create an image

The following example shows the schema for all of the parameters of the Image
resource document in both YAML and JSON format.

#### YAML

```yaml

Resources:
  ImageAllParameters:
    Type: 'AWS::ImageBuilder::Image'
    Properties:
      ImageRecipeArn: !Ref ImageRecipeArn
      InfrastructureConfigurationArn: !Ref InfrastructureConfigurationArn
      DistributionConfigurationArn: !Ref DistributionConfigurationArn
      ImageTestsConfiguration:
        ImageTestsEnabled: false
        TimeoutMinutes: 60
      Tags:
        CustomerImageTagKey1: 'CustomerImageTagValue1'
        CustomerImageTagKey2: 'CustomerImageTagValue2'
```

#### JSON

```json

{
    "Resources": {
        "ImageAllParameters": {
            "Type": "AWS::ImageBuilder::Image",
            "Properties": {
                "ImageRecipeArn": {
                    "Ref": "ImageRecipeArn"
                },
                "InfrastructureConfigurationArn": {
                    "Ref": "InfrastructureConfigurationArn"
                },
                "DistributionConfigurationArn": {
                    "Ref": "DistributionConfigurationArn"
                },
                "ImageTestsConfiguration": {
                    "ImageTestsEnabled": false,
                    "TimeoutMinutes": 60
                },
                "Tags": {
                    "CustomerImageTagKey1": "CustomerImageTagValue1",
                    "CustomerImageTagKey2": "CustomerImageTagValue2"
                }
            }
        }
    }
}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TargetContainerRepository

DeletionSettings
