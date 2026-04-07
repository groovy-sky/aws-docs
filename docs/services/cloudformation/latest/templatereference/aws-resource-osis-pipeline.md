This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OSIS::Pipeline

The AWS::OSIS::Pipeline resource creates an Amazon OpenSearch Ingestion pipeline.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::OSIS::Pipeline",
  "Properties" : {
      "BufferOptions" : BufferOptions,
      "EncryptionAtRestOptions" : EncryptionAtRestOptions,
      "LogPublishingOptions" : LogPublishingOptions,
      "MaxUnits" : Integer,
      "MinUnits" : Integer,
      "PipelineConfigurationBody" : String,
      "PipelineName" : String,
      "PipelineRoleArn" : String,
      "ResourcePolicy" : ResourcePolicy,
      "Tags" : [ Tag, ... ],
      "VpcOptions" : VpcOptions
    }
}

```

### YAML

```yaml

Type: AWS::OSIS::Pipeline
Properties:
  BufferOptions:
    BufferOptions
  EncryptionAtRestOptions:
    EncryptionAtRestOptions
  LogPublishingOptions:
    LogPublishingOptions
  MaxUnits: Integer
  MinUnits: Integer
  PipelineConfigurationBody: String
  PipelineName: String
  PipelineRoleArn: String
  ResourcePolicy:
    ResourcePolicy
  Tags:
    - Tag
  VpcOptions:
    VpcOptions

```

## Properties

`BufferOptions`

Options that specify the configuration of a persistent buffer.
To configure how OpenSearch Ingestion encrypts this data, set the `EncryptionAtRestOptions`. For more information, see [Persistent buffering](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/osis-features-overview.html#persistent-buffering).

_Required_: No

_Type_: [BufferOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-osis-pipeline-bufferoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionAtRestOptions`

Options to control how OpenSearch encrypts buffer data.

_Required_: No

_Type_: [EncryptionAtRestOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-osis-pipeline-encryptionatrestoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogPublishingOptions`

Key-value pairs that represent log publishing settings.

_Required_: No

_Type_: [LogPublishingOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-osis-pipeline-logpublishingoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxUnits`

The maximum pipeline capacity, in Ingestion Compute Units (ICUs).

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `384`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinUnits`

The minimum pipeline capacity, in Ingestion Compute Units (ICUs).

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `384`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PipelineConfigurationBody`

The Data Prepper pipeline configuration in YAML format.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PipelineName`

The name of the pipeline.

_Required_: Yes

_Type_: String

_Pattern_: `[a-z][a-z0-9\-]+`

_Minimum_: `3`

_Maximum_: `28`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PipelineRoleArn`

The Amazon Resource Name (ARN) of the IAM role that the pipeline uses to access AWS
resources.

_Required_: No

_Type_: String

_Pattern_: `^arn:(aws|aws\-cn|aws\-us\-gov|aws\-iso|aws\-iso\-b|aws\-iso\-e|aws\-iso\-f):iam::[0-9]+:role\/.*$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourcePolicy`

Property description not available.

_Required_: No

_Type_: [ResourcePolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-osis-pipeline-resourcepolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

List of tags to add to the pipeline upon creation.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-osis-pipeline-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcOptions`

Options that specify the subnets and security groups for an OpenSearch Ingestion
VPC endpoint.

_Required_: No

_Type_: [VpcOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-osis-pipeline-vpcoptions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When the logical ID of this resource is provided to the Ref intrinsic function, Ref
returns the resource name, such as `mystack-abc1d2efg3h4`. For more information
about using the Ref function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

### Fn::GetAtt

GetAtt returns a value for a specified attribute of this type. For more information, see
[Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html). The following are the available attributes and sample return
values.

`IngestEndpointUrls`

A list of the ingestion endpoints for the pipeline that you can send data to. Currently,
only a single ingestion endpoint is supported for a pipeline. For example,
`my-pipeline-123456789012.us-east-1.osis.amazonaws.com`.

`PipelineArn`

The Amazon Resource Name (ARN) of the pipeline.

`VpcEndpoints`

The VPC interface endpoints that have access to the pipeline.

`VpcEndpointService`

The VPC endpoint service name for the pipeline.

## Examples

### Create an OpenSearch Ingestion pipeline

The following example creates an OpenSearch Ingestion pipeline running version 2.x
of Data Prepper with log publishing enabled.

#### JSON

```json

{
   "Parameters": {
      "PipelineConfigurationBody": {
         "Type": "String"
      }
   },
   "Resources": {
      "OpenSearchIngestionPipeline": {
         "Type": "AWS::OSIS::Pipeline",
         "Properties": {
            "LogPublishingOptions": {
               "IsLoggingEnabled": true,
               "CloudWatchLogDestination": {
                  "LogGroup": "/aws/vendedlogs/OpenSearchService/pipelines"
               }
            },
            "MinUnits": 3,
            "MaxUnits": 9,
            "PipelineConfigurationBody": {
               "Value": {
                  "Ref": "PipelineConfigurationBody"
               }
            },
            "PipelineName": "test-pipeline"
         }
      }
   },
   "Outputs": {
      "PipelineArn": {
         "Value": {
            "Ref": "OpenSearchIngestionPipeline"
         }
      },
      "IngestEndpointUrls": {
         "Value": {
            "Fn::Select": [
               0,
               {
                  "Fn::GetAtt": [
                     "OpenSearchIngestionPipeline",
                     "IngestEndpointUrls"
                  ]
               }
            ]
         }
      }
   }
}
```

#### YAML

```yaml

Parameters:
  PipelineConfigurationBody:
    Type: String
Resources:
  OpenSearchIngestionPipeline:
    Type: 'AWS::OSIS::Pipeline'
    Properties:
      LogPublishingOptions:
        IsLoggingEnabled: true
        CloudWatchLogDestination:
          LogGroup: /aws/vendedlogs/test-pipeline
      MinUnits: 3
      MaxUnits: 9
      PipelineConfigurationBody:
        Ref: PipelineConfigurationBody
      PipelineName: test-pipeline
Outputs:
  PipelineArn:
    Value:
      Ref: OpenSearchIngestionPipeline
  IngestEndpointUrls:
    Value: !Select [0, !GetAtt OpenSearchIngestionPipeline.IngestEndpointUrls]
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon OpenSearch Ingestion

BufferOptions
