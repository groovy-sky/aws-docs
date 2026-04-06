This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# `Ref`

The intrinsic function `Ref` returns the value of a specified parameter,
resource, or another intrinsic function. This function is commonly used to create references
between resources within a CloudFormation template.

## Declaration

### JSON

```json

{ "Ref" : "logicalName" }
```

```json

{ "Ref" : "IntrinsicFunction" }
```

### YAML

Syntax for the full function name:

```yaml

Ref: logicalName
```

```yaml

Ref:
   IntrinsicFunction
```

Syntax for the short form:

```yaml

!Ref logicalName
```

```yaml

!Ref
   IntrinsicFunction
```

## Parameters

logicalName

The logical name of the resource or parameter you want to
reference.

IntrinsicFunction

The intrinsic function that resolves to a valid string. It should contain
references to parameters or identifiers, and should not contain resource
logical identifiers.

## Return value

The return value of `Ref` depends on the type of entity being
referenced:

- When you specify a parameter's logical name, it returns the value of the
parameter. For more information, see [CloudFormation template\
Parameters syntax](../userguide/parameters-section-structure.md).

- When you specify a resource's logical name, it returns a value that you use to
identify that resource. Usually, that's the name of the resource. However, for
some resources, an identifier is returned that has another significant meaning
in the context of the resource. For example, the `AWS::EC2::EIP`
resource returns the IP address, and the `AWS::EC2::Instance` returns
the instance ID. For more information about `Ref` return values for a
resource, see the documentation for that resource in the [Resource and property reference](aws-template-resource-type-ref.md).

- When you specify an intrinsic function, it returns the output of that
function.

## Examples

### Create references between resources

The following resource declaration for an Elastic IP address needs the instance ID
of an EC2 instance. It uses the `Ref` function to specify the instance ID
of the `MyEC2Instance` resource declared elsewhere in the
template.

#### JSON

```json

{
  "AWSTemplateFormatVersion":"2010-09-09",
  "Resources":{

      ...

    "MyEIP":{
      "Type":"AWS::EC2::EIP",
      "Properties":{
        "InstanceId":{
          "Ref":"MyEC2Instance"
        }
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:

  ...

  MyEIP:
    Type: AWS::EC2::EIP
    Properties:
      InstanceId: !Ref MyEC2Instance
```

### Return a resource identifier as stack output

The following examples show how to use the `Ref` function to return the
name of an Amazon S3 bucket with the logical name `MyBucket` as stack output.

#### JSON

```json

{
  "AWSTemplateFormatVersion":"2010-09-09",
  "Resources":{
    "MyBucket":{
      "Type":"AWS::S3::Bucket",
      "Properties":{
        "BucketName":{
          "Fn::Sub": "${AWS::StackName}-mybucket"
        }
      }
    }
  },
  "Outputs":{
    "BucketNameOutput":{
      "Description":"The name of the S3 bucket",
      "Value":{
        "Ref":"MyBucket"
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  MyBucket:
    Type: AWS::S3::Bucket
    Properties:
      BucketName: !Sub ${AWS::StackName}-mybucket

Outputs:
  BucketNameOutput:
    Description: The name of the S3 bucket
    Value: !Ref MyBucket
```

### Use `Fn::Join` intrinsic function inside `Ref` function

###### Note

When you use the `AWS::LanguageExtensions` transform, you can use
`Ref` in combination with other intrinsic functions. For
supported functions, see [Supported functions](#ref-supported-functions).

The following examples show how to set identifiers of resources using the
`Fn::Sub` intrinsic function, conditions, and the input for the
`Stage` parameter. The `Ref` and the
`Fn::GetAtt` functions then reference the appropriate values, based
on the stage. `Fn::Sub` is first used with `Fn::GetAtt` to
obtain the ARN of the appropriate Amazon SQS queue to set the dimensions of the Amazon CloudWatch
alarm. Next, [Fn::Join](intrinsic-function-reference-join.md) is used with `Ref` to
create the name of the SNS topic for the `AlarmActions` property.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Transform": "AWS::LanguageExtensions",
    "Parameters": {
        "Stage": {
            "Type": "String",
            "Default": "Dev",
            "AllowedValues": [
                "Dev",
                "Prod"
            ]
        }
    },
    "Conditions": {
        "isProd": {
            "Fn::Equals": [
                {"Ref": "Stage"},
                "Prod"
            ]
        },
        "isDev": {
            "Fn::Equals": [
                {"Ref": "Stage"},
                "Dev"
            ]
        }
    },
    "Resources": {
        "DevQueue": {
            "Type": "AWS::SQS::Queue",
            "Condition": "isDev",
            "Properties": {
                "QueueName": {"Fn::Sub": "My${Stage}Queue"}
            }
        },
        "ProdQueue": {
            "Type": "AWS::SQS::Queue",
            "Condition": "isProd",
            "Properties": {
                "QueueName": {"Fn::Sub": "My${Stage}Queue"}
            }
        },
        "DevTopic": {
            "Condition": "isDev",
            "Type": "AWS::SNS::Topic"
        },
        "ProdTopic": {
            "Condition": "isProd",
            "Type": "AWS::SNS::Topic"
        },
        "MyAlarm": {
            "Type": "AWS::CloudWatch::Alarm",
            "Properties": {
                "AlarmDescription": "Alarm if queue depth grows beyond 10 messages",
                "Namespace": "AWS/SQS",
                "MetricName": "ApproximateNumberOfMessagesVisible",
                "Dimensions":[
                    {
                        "Name": {"Fn::Sub": "${Stage}Queue"},
                        "Value": {"Fn::GetAtt": [{"Fn::Sub": "${Stage}Queue"}, "QueueName"]}
                    }
                ],
                "Statistic": "Sum",
                "Period": 300,
                "EvaluationPeriods": 1,
                "Threshold": 10,
                "ComparisonOperator": "GreaterThanThreshold",
                "AlarmActions": [
                    {
                        "Ref": {"Fn::Join": ["", [{"Ref": "Stage"}, "Topic"]]}
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Transform: AWS::LanguageExtensions
Parameters:
  Stage:
    Type: String
    Default: Dev
    AllowedValues:
      - Dev
      - Prod
Conditions:
  isProd: !Equals
    - !Ref Stage
    - Prod
  isDev: !Equals
    - !Ref Stage
    - Dev
Resources:
  DevQueue:
    Type: AWS::SQS::Queue
    Condition: isDev
    Properties:
      QueueName: !Sub My${Stage}Queue
  ProdQueu:
    Type: AWS::SQS::Queue
    Condition: isProd
    Properties:
      QueueName: !Sub My${Stage}Queue
  DevTopic:
    Condition: isDev
    Type: AWS::SNS::Topic
  ProdTopic:
    Condition: isProd
    Type: AWS::SNS::Topic
  MyAlarm:
    Type: AWS::CloudWatch::Alarm
    Properties:
      AlarmDescription: Alarm if queue depth grows beyond 10 messages
      Namespace: AWS/SQS
      MetricName: ApproximateNumberOfMessagesVisible
      Dimensions:
        - Name: !Sub '${Stage}Queue'
          Value: !GetAtt
            - !Sub '${Stage}Queue'
            - QueueName
      Statistic: Sum
      Period: 300
      EvaluationPeriods: 1
      Threshold: 10
      ComparisonOperator: GreaterThanThreshold
      AlarmActions:
        - !Ref
          'Fn::Join':
            - ''
            - - !Ref Stage
              - Topic
```

## Supported functions

When you use the [AWS::LanguageExtensions transform](transform-aws-languageextensions.md), you can use the
following functions within the `Ref` function.

- [Fn::Base64](intrinsic-function-reference-base64.md)

- [Fn::FindInMap](intrinsic-function-reference-findinmap.md)

- [Fn::If](intrinsic-function-reference-conditions.md#intrinsic-function-reference-conditions-if)

- [Fn::ImportValue](intrinsic-function-reference-importvalue.md)

- [Fn::Join](intrinsic-function-reference-join.md)

- [Fn::Sub](intrinsic-function-reference-sub.md)

- [Fn::ToJsonString](intrinsic-function-reference-tojsonstring.md)

- `Ref`

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Fn::Transform

Rule functions
