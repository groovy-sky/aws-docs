This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::NotebookInstanceLifecycleConfig

The `AWS::SageMaker::NotebookInstanceLifecycleConfig` resource creates shell scripts that run when
you create and/or start a notebook instance. For information about notebook instance lifecycle configurations, see
[Customize a Notebook\
Instance](../../../sagemaker/latest/dg/notebook-lifecycle-config.md) in the _Amazon SageMaker Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SageMaker::NotebookInstanceLifecycleConfig",
  "Properties" : {
      "NotebookInstanceLifecycleConfigName" : String,
      "OnCreate" : [ NotebookInstanceLifecycleHook, ... ],
      "OnStart" : [ NotebookInstanceLifecycleHook, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SageMaker::NotebookInstanceLifecycleConfig
Properties:
  NotebookInstanceLifecycleConfigName: String
  OnCreate:
    - NotebookInstanceLifecycleHook
  OnStart:
    - NotebookInstanceLifecycleHook

```

## Properties

`NotebookInstanceLifecycleConfigName`

The name of the lifecycle configuration.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9](-*[a-zA-Z0-9])*`

_Minimum_: `0`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OnCreate`

A shell script that runs only once, when you create a notebook instance. The shell
script must be a base64-encoded string.

_Required_: No

_Type_: Array of [NotebookInstanceLifecycleHook](aws-properties-sagemaker-notebookinstancelifecycleconfig-notebookinstancelifecyclehook.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OnStart`

A shell script that runs every time you start a notebook instance, including when you
create the notebook instance. The shell script must be a base64-encoded string.

_Required_: No

_Type_: Array of [NotebookInstanceLifecycleHook](aws-properties-sagemaker-notebookinstancelifecycleconfig-notebookinstancelifecyclehook.md)

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the endpoint configuration, such as
`arn:aws:sagemaker:us-west-2:012345678901:notebook-instance-lifecycle-config/mylifecycleconfig`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The name of the notebook instance lifecycle configuration.

`NotebookInstanceLifecycleConfigName`

The name of the lifecycle configuration, such as `MyLifecycleConfig`.

## Examples

### SageMaker NotebookInstanceLifecycleConfig Example

The following example creates a notebook instance with an associated lifecycle configuration.

#### JSON

```json

{
  "Description": "Basic NotebookInstance test",
  "Resources": {
    "BasicNotebookInstance": {
      "Type": "AWS::SageMaker::NotebookInstance",
      "Properties": {
        "InstanceType": "ml.t2.medium",
        "RoleArn": { "Fn::GetAtt" : [ "ExecutionRole", "Arn" ] },
        "LifecycleConfigName": { "Fn::GetAtt" : [ "BasicNotebookInstanceLifecycleConfig", "NotebookInstanceLifecycleConfigName" ] }
    },
    "BasicNotebookInstanceLifecycleConfig": {
      "Type": "AWS::SageMaker::NotebookInstanceLifecycleConfig",
      "Properties": {
        "OnStart": [
          {
            "Content": {
              "Fn::Base64": "echo 'hello'"
            }
          }
        ]
      }
    },
    "ExecutionRole": {
      "Type": "AWS::IAM::Role",
      "Properties": {
        "AssumeRolePolicyDocument": {
          "Version": "2012-10-17",
          "Statement": [
            {
              "Effect": "Allow",
              "Principal": {
                "Service": [
                  "sagemaker.amazonaws.com"
                ]
              },
              "Action": [
                "sts:AssumeRole"
              ]
            }
          ]
        },
        "Path": "/",
        "Policies": [
          {
            "PolicyName": "root",
            "PolicyDocument": {
              "Version": "2012-10-17",
              "Statement": [
                {
                  "Effect": "Allow",
                  "Action": "*",
                  "Resource": "*"
                }
              ]
            }
          }
        ]
      }
    }
  },
  "Outputs": {
    "BasicNotebookInstanceId": {
      "Value": { "Ref" : "BasicNotebookInstance" }
    },
    "BasicNotebookInstanceLifecycleConfigId": {
      "Value": { "Ref" : "BasicNotebookInstanceLifecycleConfig" }
    }
  }
}
}
```

#### YAML

```yaml

Description: "Basic NotebookInstance test"
Resources:
  BasicNotebookInstance:
    Type: "AWS::SageMaker::NotebookInstance"
    Properties:
      InstanceType: "ml.t2.medium"
      RoleArn: !GetAtt ExecutionRole.Arn
      LifecycleConfigName: !GetAtt BasicNotebookInstanceLifecycleConfig.NotebookInstanceLifecycleConfigName
  BasicNotebookInstanceLifecycleConfig:
    Type: "AWS::SageMaker::NotebookInstanceLifecycleConfig"
    Properties:
      OnStart:
        - Content:
            Fn::Base64: "echo 'hello'"
  ExecutionRole:
    Type: "AWS::IAM::Role"
    Properties:
      AssumeRolePolicyDocument:
        Version: "2012-10-17"
        Statement:
          -
            Effect: "Allow"
            Principal:
              Service:
                - "sagemaker.amazonaws.com"
            Action:
              - "sts:AssumeRole"
      Path: "/"
      Policies:
        -
          PolicyName: "root"
          PolicyDocument:
            Version: "2012-10-17"
            Statement:
              -
                Effect: "Allow"
                Action: "*"
                Resource: "*"
Outputs:
  BasicNotebookInstanceId:
    Value: !Ref BasicNotebookInstance
  BasicNotebookInstanceLifecycleConfigId:
    Value: !Ref BasicNotebookInstanceLifecycleConfig
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

NotebookInstanceLifecycleHook

All content copied from https://docs.aws.amazon.com/.
