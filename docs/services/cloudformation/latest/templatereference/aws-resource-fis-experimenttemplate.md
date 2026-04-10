This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FIS::ExperimentTemplate

Specifies an experiment template.

An experiment template includes the following components:

- **Targets**: A target can be a specific resource
in your AWS environment, or one or more resources that match criteria that you
specify, for example, resources that have specific tags.

- **Actions**: The actions to carry out on the
target. You can specify multiple actions, the duration of each action, and when
to start each action during an experiment.

- **Stop conditions**: If a stop condition is
triggered while an experiment is running, the experiment is automatically
stopped. You can define a stop condition as a CloudWatch alarm.

For more information, see [Experiment templates](../../../fis/latest/userguide/experiment-templates.md)
in the _AWS Fault Injection Service User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::FIS::ExperimentTemplate",
  "Properties" : {
      "Actions" : {Key: Value, ...},
      "Description" : String,
      "ExperimentOptions" : ExperimentTemplateExperimentOptions,
      "ExperimentReportConfiguration" : ExperimentTemplateExperimentReportConfiguration,
      "LogConfiguration" : ExperimentTemplateLogConfiguration,
      "RoleArn" : String,
      "StopConditions" : [ ExperimentTemplateStopCondition, ... ],
      "Tags" : {Key: Value, ...},
      "Targets" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::FIS::ExperimentTemplate
Properties:
  Actions:
    Key: Value
  Description: String
  ExperimentOptions:
    ExperimentTemplateExperimentOptions
  ExperimentReportConfiguration:
    ExperimentTemplateExperimentReportConfiguration
  LogConfiguration:
    ExperimentTemplateLogConfiguration
  RoleArn: String
  StopConditions:
    - ExperimentTemplateStopCondition
  Tags:
    Key: Value
  Targets:
    Key: Value

```

## Properties

`Actions`

The actions for the experiment.

_Required_: No

_Type_: Object of [ExperimentTemplateAction](aws-properties-fis-experimenttemplate-experimenttemplateaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description for the experiment template.

_Required_: Yes

_Type_: String

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExperimentOptions`

The experiment options for an experiment template.

_Required_: No

_Type_: [ExperimentTemplateExperimentOptions](aws-properties-fis-experimenttemplate-experimenttemplateexperimentoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExperimentReportConfiguration`

Describes the report configuration for the experiment template.

_Required_: No

_Type_: [ExperimentTemplateExperimentReportConfiguration](aws-properties-fis-experimenttemplate-experimenttemplateexperimentreportconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogConfiguration`

The configuration for experiment logging.

_Required_: No

_Type_: [ExperimentTemplateLogConfiguration](aws-properties-fis-experimenttemplate-experimenttemplatelogconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of an IAM role.

_Required_: Yes

_Type_: String

_Maximum_: `1224`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StopConditions`

The stop conditions for the experiment.

_Required_: Yes

_Type_: Array of [ExperimentTemplateStopCondition](aws-properties-fis-experimenttemplate-experimenttemplatestopcondition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags for the experiment template.

_Required_: No

_Type_: Object of String

_Pattern_: `.{1,128}`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Targets`

The targets for the experiment.

_Required_: Yes

_Type_: Object of [ExperimentTemplateTarget](aws-properties-fis-experimenttemplate-experimenttemplatetarget.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the experiment template ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The ID of the experiment template.

## Examples

### Stop and start an instance based on a tag

The following example creates an experiment template that stops and starts one instance
with the tag env=prod, chosen at random.

#### YAML

```yaml

Resources:
  ExperimentTemplate:
    Type: 'AWS::FIS::ExperimentTemplate'
    Properties:
      Description: 'stop an instance based on a tag'
      Actions:
        stopInstances:
          ActionId: 'aws:ec2:stop-instances'
          Parameters:
            startInstancesAfterDuration: 'PT2M'
          Targets:
            Instances: oneRandomInstance
      Targets:
        oneRandomInstance:
          ResourceTags:
            'env': 'prod'
          ResourceType: 'aws:ec2:instance'
          SelectionMode: 'COUNT(1)'
      StopConditions:
        - Source: 'none'
      Tags:
        Name: 'fisStopInstances'
      RoleArn: !GetAtt FISRole.Arn
  FISRole:
    Type: 'AWS::IAM::Role'
    Properties:
      AssumeRolePolicyDocument:
        Version: '2012-10-17'
        Statement:
          - Effect: Allow
            Principal:
              Service: 'fis.amazonaws.com'
            Action: 'sts:AssumeRole'
      Policies:
        - PolicyName: 'FISRoleEC2Actions'
          PolicyDocument:
            Version: '2012-10-17'
            Statement:
              - Effect: Allow
                Action:
                  - 'ec2:RebootInstances'
                  - 'ec2:StopInstances'
                  - 'ec2:StartInstances'
                  - 'ec2:TerminateInstances'
                Resource: 'arn:aws:ec2:*:*:instance/*'
```

#### JSON

```json

{
  "Resources": {
    "ExperimentTemplate": {
      "Type": "AWS::FIS::ExperimentTemplate",
      "DeletionPolicy": "Retain",
      "Properties": {
        "Description": "stop an instance based on a tag",
        "Actions": {
          "stopInstances": {
            "ActionId": "aws:ec2:stop-instances",
            "Parameters": {
              "startInstancesAfterDuration": "PT2M"
            },
            "Targets": {
              "Instances": "oneRandomInstance"
            }
          }
        },
        "Targets": {
          "oneRandomInstance": {
            "ResourceTags": {
              "env": "prod"
            },
            "ResourceType": "aws:ec2:instance",
            "SelectionMode": "COUNT(1)"
          }
        },
        "StopConditions": [
          {
            "Source": "none"
          }
        ],
        "Tags": {
          "Name": "fisStopInstancesJson"
        },
        "RoleArn": {
          "Fn::GetAtt": ["FISRole", "Arn"]
        }
      }
    },
    "FISRole": {
      "Type": "AWS::IAM::Role",
      "Properties": {
        "AssumeRolePolicyDocument": {
          "Version": "2012-10-17",
          "Statement": [
            {
                "Effect": "Allow",
                "Principal": {
                  "Service": "fis.amazonaws.com"
                },
                "Action": "sts:AssumeRole"
            }
          ]
        },
        "Policies": [
          {
            "PolicyName": "FISRoleEC2Actions",
            "PolicyDocument": {
              "Version": "2012-10-17",
              "Statement": [
                {
                  "Effect": "Allow",
                  "Action": [
                    "ec2:RebootInstances",
                    "ec2:StopInstances",
                    "ec2:StartInstances",
                    "ec2:TerminateInstances"
                  ],
                  "Resource": "arn:aws:ec2:*:*:instance/*"
                }
              ]
            }
          }
        ]
      }
    }
  }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Fault Injection Service

CloudWatchDashboard

All content copied from https://docs.aws.amazon.com/.
