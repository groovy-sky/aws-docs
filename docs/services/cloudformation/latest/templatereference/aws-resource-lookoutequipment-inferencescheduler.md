This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LookoutEquipment::InferenceScheduler

Creates a scheduled inference. Scheduling an inference is setting up a continuous
real-time inference plan to analyze new measurement data. When setting up the schedule, you
provide an Amazon S3 bucket location for the input data, assign it a delimiter
between separate entries in the data, set an offset delay if desired, and set the frequency
of inferencing. You must also provide an Amazon S3 bucket location for the output
data.

###### Note

Updating some properties below (for example,
InferenceSchedulerName and
ServerSideKmsKeyId) triggers a resource replacement, which requires a
new model. To replace such a property using CloudFormation, but without creating a
completely new stack,
you
must replace ModelName. If you need to replace the property, but want
to use the same model, delete the current stack and create a new one with the updated
properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::LookoutEquipment::InferenceScheduler",
  "Properties" : {
      "DataDelayOffsetInMinutes" : Integer,
      "DataInputConfiguration" : DataInputConfiguration,
      "DataOutputConfiguration" : DataOutputConfiguration,
      "DataUploadFrequency" : String,
      "InferenceSchedulerName" : String,
      "ModelName" : String,
      "RoleArn" : String,
      "ServerSideKmsKeyId" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::LookoutEquipment::InferenceScheduler
Properties:
  DataDelayOffsetInMinutes: Integer
  DataInputConfiguration:
    DataInputConfiguration
  DataOutputConfiguration:
    DataOutputConfiguration
  DataUploadFrequency: String
  InferenceSchedulerName: String
  ModelName: String
  RoleArn: String
  ServerSideKmsKeyId: String
  Tags:
    - Tag

```

## Properties

`DataDelayOffsetInMinutes`

A period of time (in minutes) by which inference on the data is delayed after the data
starts. For instance, if an offset delay time of five minutes was selected, inference will
not begin on the data until the first data measurement after the five minute mark. For
example, if five minutes is selected, the inference scheduler will wake up at the
configured frequency with the additional five minute delay time to check the customer S3
bucket. The customer can upload data at the same frequency and they don't need to stop and
restart the scheduler when uploading new data.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `60`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataInputConfiguration`

Specifies configuration information for the input data for the inference scheduler,
including delimiter, format, and dataset location.

_Required_: Yes

_Type_: [DataInputConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lookoutequipment-inferencescheduler-datainputconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataOutputConfiguration`

Specifies configuration information for the output results for the inference scheduler,
including the Amazon S3 location for the output.

_Required_: Yes

_Type_: [DataOutputConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lookoutequipment-inferencescheduler-dataoutputconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataUploadFrequency`

How often data is uploaded to the source S3 bucket for the input data. This value is the
length of time between data uploads. For instance, if you select 5 minutes, Amazon Lookout
for Equipment will upload the real-time data to the source bucket once every 5 minutes.
This frequency also determines how often Amazon Lookout for Equipment starts a scheduled inference on your
data. In this example, it starts once every 5 minutes.

_Required_: Yes

_Type_: String

_Allowed values_: `PT5M | PT10M | PT15M | PT30M | PT1H`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InferenceSchedulerName`

The name of the inference scheduler.

_Required_: No

_Type_: String

_Pattern_: `^[0-9a-zA-Z_-]{1,200}$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ModelName`

The name of the machine learning model used for the inference scheduler.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-zA-Z_-]{1,200}$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The Amazon Resource Name (ARN) of a role with permission to access the data source being
used for the inference.

_Required_: Yes

_Type_: String

_Pattern_: `arn:aws(-[^:]+)?:iam::[0-9]{12}:role/.+`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerSideKmsKeyId`

Provides the identifier of the AWS KMS key used to encrypt inference
scheduler data by Amazon Lookout for Equipment.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9][A-Za-z0-9:_/+=,@.-]{0,2048}$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Any tags associated with the inference scheduler.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lookoutequipment-inferencescheduler-tag.html)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`InferenceSchedulerArn`

The Amazon Resource Name (ARN) of the inference scheduler being created.

## Examples

### Inference Scheduler

Creates a schedule inference

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Parameters": {
    "Name": {
      "Type": "String"
    }
  },
  "Resources": {
    "MyInferenceScheduler": {
      "Type": "AWS::LookoutEquipment::InferenceScheduler",
      "Properties": {
        "InferenceSchedulerName": {"Ref": "Name"},
        "ModelName": "my-test-model",
        "DataUploadFrequency": "PT5M",
        "RoleArn": {"Fn::ImportValue": "MyExportRole"},
        "DataInputConfiguration": {
          "S3InputConfiguration": {
            "Bucket": {"Fn::ImportValue":"MyExportBucket"},
            "Prefix": "myTestData/"
          }
        },
        "DataOutputConfiguration": {
          "S3OutputConfiguration": {
            "Bucket": {"Fn::ImportValue": "MyExportBucket"},
            "Prefix": "myTestData/"
          }
        }
      }
    }
  },
  "Outputs": {
    "PrimaryId": {
      "Value": {
        "Ref": "MyInferenceScheduler"
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Parameters:
  Name:
    Type: String
Resources:
  MyInferenceScheduler:
    Type: 'AWS::LookoutEquipment::InferenceScheduler'
    Properties:
      InferenceSchedulerName: !Ref Name
      ModelName: my-test-model
      DataUploadFrequency: PT5M
      RoleArn: !ImportValue MyExportRole
      DataInputConfiguration:
        S3InputConfiguration:
          Bucket: !ImportValue MyExportBucket
          Prefix: myTestData/
      DataOutputConfiguration:
        S3OutputConfiguration:
          Bucket: !ImportValue MyExportBucket
          Prefix: myTestData/
Outputs:
  PrimaryId:
    Value: !Ref MyInferenceScheduler

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Lookout for Equipment

DataInputConfiguration
