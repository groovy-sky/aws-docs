This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Config::ConfigurationRecorder

The `AWS::Config::ConfigurationRecorder` resource type describes the AWS resource types that AWS Config records for configuration changes.

The configuration recorder stores the configuration changes of the specified resources in your account as configuration items.

###### Note

To enable AWS Config, you must create a configuration recorder and a delivery channel.

AWS Config uses the delivery channel to deliver the configuration changes to your Amazon S3 bucket or Amazon SNS topic.
For more information, see [AWS::Config::DeliveryChannel](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html).

AWS CloudFormation starts the recorder as soon as the delivery channel is available.

To stop the recorder and delete it, delete the configuration recorder from your stack. To stop the recorder without deleting it,
call the [StopConfigurationRecorder](https://docs.aws.amazon.com/config/latest/APIReference/API_StopConfigurationRecorder.html) action of the AWS Config API directly.

For more information, see [Configuration Recorder](https://docs.aws.amazon.com/config/latest/developerguide/config-concepts.html#config-recorder) in the AWS Config Developer Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Config::ConfigurationRecorder",
  "Properties" : {
      "Name" : String,
      "RecordingGroup" : RecordingGroup,
      "RecordingMode" : RecordingMode,
      "RoleARN" : String
    }
}

```

### YAML

```yaml

Type: AWS::Config::ConfigurationRecorder
Properties:
  Name: String
  RecordingGroup:
    RecordingGroup
  RecordingMode:
    RecordingMode
  RoleARN: String

```

## Properties

`Name`

The name of the configuration recorder. AWS Config automatically assigns the name of "default" when creating the configuration recorder.

You cannot change the name of the configuration recorder after it has been created. To change the configuration recorder name, you must delete it and create a new configuration recorder with a new name.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RecordingGroup`

Specifies which resource types AWS Config
records for configuration changes.

###### Note

**High Number of AWS Config Evaluations**

You may notice increased activity in your account during your initial month recording with AWS Config when compared to subsequent months. During the
initial bootstrapping process, AWS Config runs evaluations on all the resources in your account that you have selected
for AWS Config to record.

If you are running ephemeral workloads, you may see increased activity from AWS Config as it records configuration changes associated with creating and deleting these
temporary resources. An _ephemeral workload_ is a temporary use of computing resources that are loaded
and run when needed. Examples include Amazon Elastic Compute Cloud (Amazon EC2)
Spot Instances, Amazon EMR jobs, and AWS Auto Scaling. If you want
to avoid the increased activity from running ephemeral workloads, you can run these
types of workloads in a separate account with AWS Config turned off to avoid
increased configuration recording and rule evaluations.

_Required_: No

_Type_: [RecordingGroup](aws-properties-config-configurationrecorder-recordinggroup.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordingMode`

Specifies the default recording frequency for the configuration recorder.

AWS Config supports _Continuous recording_ and _Daily recording_.

- Continuous recording allows you to record configuration changes continuously whenever a change occurs.

- Daily recording allows you to receive a configuration item (CI) representing the most recent state of your resources over the last 24-hour period, only if it’s different from the previous CI recorded.

###### Note

**Some resource types require continuous recording**

AWS Firewall Manager depends on continuous recording to monitor your resources. If you are using Firewall Manager,
it is recommended that you set the recording frequency to Continuous.

You can also override the recording frequency for specific resource types.

_Required_: No

_Type_: [RecordingMode](aws-properties-config-configurationrecorder-recordingmode.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleARN`

Amazon Resource Name (ARN) of the IAM role assumed by AWS Config and used by the configuration recorder.
For more information, see [Permissions for the IAM Role Assigned](https://docs.aws.amazon.com/config/latest/developerguide/iamrole-permissions.html) to AWS Config in the AWS Config Developer Guide.

###### Note

**Pre-existing AWS Config role**

If you have used an AWS service that uses AWS Config, such as AWS Security Hub CSPM or
AWS Control Tower, and an AWS Config role has already been created, make sure that the
IAM role that you use when setting up AWS Config keeps the same minimum
permissions as the already created AWS Config role. You must do this so that the
other AWS service continues to run as expected.

For example, if AWS Control Tower has an IAM role that allows AWS Config to read
Amazon Simple Storage Service (Amazon S3) objects, make sure that the same permissions are granted
within the IAM role you use when setting up AWS Config. Otherwise, it may
interfere with how AWS Control Tower operates. For more information about IAM
roles for AWS Config,
see [**Identity and Access Management for AWS Config**](https://docs.aws.amazon.com/config/latest/developerguide/security-iam.html) in the _AWS Config Developer Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the configuration recorder name, such as default.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

### Configuration Recorder

The following example creates a configuration recorder for EC2 volumes.

#### JSON

```json

"ConfigRecorder": {
  "Type": "AWS::Config::ConfigurationRecorder",
  "Properties": {
    "Name": "default",
    "RecordingGroup": {
      "ResourceTypes": ["AWS::EC2::Volume"]
    },
    "RoleARN": {"Fn::GetAtt": ["ConfigRole", "Arn"]}
  }
}
```

#### YAML

```yaml

ConfigRecorder:
  Type: AWS::Config::ConfigurationRecorder
  Properties:
    Name: default
    RecordingGroup:
      ResourceTypes:
        - "AWS::EC2::Volume"
    RoleARN:
      Fn::GetAtt:
        - ConfigRole
        - Arn
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

ExclusionByResourceTypes
