This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::ImagePipeline

An image pipeline is the automation configuration for building secure OS images on AWS.
The Image Builder image pipeline is associated with an image recipe that defines the build,
validation, and test phases for an image build lifecycle. An image pipeline can be
associated with an infrastructure configuration that defines where your image is built. You
can define attributes, such as instance types, a subnet for your VPC, security groups, logging, and
other infrastructure-related configurations. You can also associate your image pipeline with a
distribution configuration to define how you would like to deploy your image.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ImageBuilder::ImagePipeline",
  "Properties" : {
      "ContainerRecipeArn" : String,
      "Description" : String,
      "DistributionConfigurationArn" : String,
      "EnhancedImageMetadataEnabled" : Boolean,
      "ExecutionRole" : String,
      "ImageRecipeArn" : String,
      "ImageScanningConfiguration" : ImageScanningConfiguration,
      "ImageTestsConfiguration" : ImageTestsConfiguration,
      "InfrastructureConfigurationArn" : String,
      "LoggingConfiguration" : PipelineLoggingConfiguration,
      "Name" : String,
      "Schedule" : Schedule,
      "Status" : String,
      "Tags" : {Key: Value, ...},
      "Workflows" : [ WorkflowConfiguration, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ImageBuilder::ImagePipeline
Properties:
  ContainerRecipeArn: String
  Description: String
  DistributionConfigurationArn: String
  EnhancedImageMetadataEnabled: Boolean
  ExecutionRole: String
  ImageRecipeArn: String
  ImageScanningConfiguration:
    ImageScanningConfiguration
  ImageTestsConfiguration:
    ImageTestsConfiguration
  InfrastructureConfigurationArn: String
  LoggingConfiguration:
    PipelineLoggingConfiguration
  Name: String
  Schedule:
    Schedule
  Status: String
  Tags:
    Key: Value
  Workflows:
    - WorkflowConfiguration

```

## Properties

`ContainerRecipeArn`

The Amazon Resource Name (ARN) of the container recipe that is used for this
pipeline.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of this image pipeline.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DistributionConfigurationArn`

The Amazon Resource Name (ARN) of the distribution configuration associated with this
image pipeline.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnhancedImageMetadataEnabled`

Collects additional information about the image being created, including the operating
system (OS) version and package list. This information is used to enhance the overall
experience of using EC2 Image Builder. Enabled by default.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionRole`

The name or Amazon Resource Name (ARN) for the IAM role you create that grants
Image Builder access to perform workflow actions.

_Required_: No

_Type_: String

_Pattern_: `^(?:arn:aws(?:-[a-z]+)*:iam::[0-9]{12}:role/)?[a-zA-Z_0-9+=,.@\-_/]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageRecipeArn`

The Amazon Resource Name (ARN) of the image recipe associated with this image pipeline.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageScanningConfiguration`

Contains settings for vulnerability scans.

_Required_: No

_Type_: [ImageScanningConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-imagepipeline-imagescanningconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ImageTestsConfiguration`

The configuration of the image tests that run after image creation to ensure the
quality of the image that was created.

_Required_: No

_Type_: [ImageTestsConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-imagepipeline-imagetestsconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InfrastructureConfigurationArn`

The Amazon Resource Name (ARN) of the infrastructure configuration associated with
this image pipeline.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoggingConfiguration`

Defines logging configuration for the output image.

_Required_: No

_Type_: [PipelineLoggingConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-imagepipeline-pipelineloggingconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the image pipeline.

_Required_: No

_Type_: String

_Pattern_: `^[-_A-Za-z-0-9][-_A-Za-z0-9 ]{1,126}[-_A-Za-z-0-9]$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Schedule`

The schedule of the image pipeline. A schedule configures how often and when a pipeline
automatically creates a new image.

_Required_: No

_Type_: [Schedule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-imagepipeline-schedule.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status of the image pipeline.

_Required_: No

_Type_: String

_Allowed values_: `DISABLED | ENABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags of this image pipeline.

_Required_: No

_Type_: Object of String

_Pattern_: `.{1,}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Workflows`

Contains the workflows that run for the image pipeline.

_Required_: No

_Type_: Array of [WorkflowConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-imagepipeline-workflowconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource ARN, such as
`arn:aws:imagebuilder:us-west-2:123456789012:image-pipeline/mywindows2016pipeline`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the Amazon Resource Name (ARN) of the image pipeline. For example,
`arn:aws:imagebuilder:us-west-2:123456789012:image-pipeline/mywindows2016pipeline`.

`DeploymentId`

Returns the deployment ID of the pipeline, used for resource create/update triggers.

`Name`

Returns the name of the image pipeline.

## Examples

### Create an image pipeline

The following example shows the schema for all of the parameters of the
ImagePipeline resource document in both YAML and JSON format .

#### YAML

```yaml

Resources:
  ImagePipelineAllParameters:
    Type: 'AWS::ImageBuilder::ImagePipeline'
    Properties:
      Name: 'image-pipeline-name'
      Description: 'description'
      ImageRecipeArn: !Ref ImageRecipeArn
      InfrastructureConfigurationArn: !Ref InfrastructureConfigurationArn
      DistributionConfigurationArn: !Ref DistributionConfigurationArn
      ImageTestsConfiguration:
        ImageTestsEnabled: false
        TimeoutMinutes: 90
      Schedule:
        ScheduleExpression: 'cron(0 0 * * ? *)'
        PipelineExecutionStartCondition: 'EXPRESSION_MATCH_ONLY'
      Status: 'DISABLED'
      Tags:
        CustomerImagePipelineTagKey1: 'CustomerImagePipelineTagValue1'
        CustomerImagePipelineTagKey2: 'CustomerImagePipelineTagValue2'
```

#### JSON

```json

{
    "Resources": {
        "ImagePipelineAllParameters": {
            "Type": "AWS::ImageBuilder::ImagePipeline",
            "Properties": {
                "Name": "image-pipeline-name",
                "Description": "description",
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
                    "TimeoutMinutes": 90
                },
                "Schedule": {
                "ScheduleExpression": "cron(0 0 * * ? *)",
                    "PipelineExecutionStartCondition": "EXPRESSION_MATCH_ONLY"
                },
                "Status": "DISABLED",
                "Tags": {
                    "CustomerImagePipelineTagKey1": "CustomerImagePipelineTagValue1",
                    "CustomerImagePipelineTagKey2": "CustomerImagePipelineTagValue2"
                }
            }
        }
    }
}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

WorkflowParameter

AutoDisablePolicy
