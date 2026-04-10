This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::LifecycleHook

The `AWS::AutoScaling::LifecycleHook` resource specifies lifecycle hooks for an
Auto Scaling group. These hooks let you create solutions that are aware of events in the Auto
Scaling instance lifecycle, and then perform a custom action on instances when the
corresponding lifecycle event occurs. A lifecycle hook provides a specified amount of time
(one hour by default) to wait for the action to complete before the instance transitions to
the next state.

Use lifecycle hooks to prepare new instances for use or to delay them from being
registered behind a load balancer before their configuration has been applied completely. You
can also use lifecycle hooks to prepare running instances to be terminated by, for example,
downloading logs or other data.

For more information, see [Amazon EC2 Auto Scaling lifecycle\
hooks](../../../autoscaling/ec2/userguide/lifecycle-hooks.md) in the _Amazon EC2 Auto Scaling User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AutoScaling::LifecycleHook",
  "Properties" : {
      "AutoScalingGroupName" : String,
      "DefaultResult" : String,
      "HeartbeatTimeout" : Integer,
      "LifecycleHookName" : String,
      "LifecycleTransition" : String,
      "NotificationMetadata" : String,
      "NotificationTargetARN" : String,
      "RoleARN" : String
    }
}

```

### YAML

```yaml

Type: AWS::AutoScaling::LifecycleHook
Properties:
  AutoScalingGroupName: String
  DefaultResult: String
  HeartbeatTimeout: Integer
  LifecycleHookName: String
  LifecycleTransition: String
  NotificationMetadata: String
  NotificationTargetARN: String
  RoleARN: String

```

## Properties

`AutoScalingGroupName`

The name of the Auto Scaling group.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DefaultResult`

The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an
unexpected failure occurs. The default value is `ABANDON`.

Valid values: `CONTINUE` \| `ABANDON`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HeartbeatTimeout`

The maximum time, in seconds, that can elapse before the lifecycle hook times out. The
range is from `30` to `7200` seconds. The default value is
`3600` seconds (1 hour).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LifecycleHookName`

The name of the lifecycle hook.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LifecycleTransition`

The lifecycle transition. For Auto Scaling groups, there are two major lifecycle
transitions.

- To create a lifecycle hook for scale-out events, specify
`autoscaling:EC2_INSTANCE_LAUNCHING`.

- To create a lifecycle hook for scale-in events, specify
`autoscaling:EC2_INSTANCE_TERMINATING`.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotificationMetadata`

Additional information that you want to include any time Amazon EC2 Auto Scaling sends a message to
the notification target.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1023`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NotificationTargetARN`

The Amazon Resource Name (ARN) of the notification target that Amazon EC2 Auto Scaling sends
notifications to when an instance is in a wait state for the lifecycle hook. You can
specify an Amazon SNS topic or an Amazon SQS queue.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleARN`

The ARN of the IAM role that allows the Auto Scaling group to publish to the specified
notification target. For information about creating this role, see [Prepare to\
add a lifecycle hook to your Auto Scaling group](../../../autoscaling/ec2/userguide/prepare-for-lifecycle-notifications.md) in the
_Amazon EC2 Auto Scaling User Guide_.

Valid only if the notification target is an Amazon SNS topic or an Amazon SQS queue.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:
`mylifecyclehook`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

The following examples specify lifecycle hooks.

For other examples of lifecycle hooks, see our [GitHub\
repository](https://github.com/aws-samples/amazon-ec2-auto-scaling-group-examples). Included are examples of lifecycle hooks that work with warm
pools.

- [Lifecycle hook for instance launch](#aws-resource-autoscaling-lifecyclehook--examples--Lifecycle_hook_for_instance_launch)

- [Lifecycle hook for instance termination](#aws-resource-autoscaling-lifecyclehook--examples--Lifecycle_hook_for_instance_termination)

- [Lifecycle hook that specifies a notification target ARN and role ARN](#aws-resource-autoscaling-lifecyclehook--examples--Lifecycle_hook_that_specifies_a_notification_target_ARN_and_role_ARN)

### Lifecycle hook for instance launch

The following example creates a launch template, an Auto Scaling group, and a
lifecycle hook that supports a custom action on your instances at launch. When you create
a launch template, you can include user data to perform configuration tasks and run
scripts when an instance launches. By doing so, you do not need to configure a
notification target. In this example, the hook keeps each instance in a wait state for 60
seconds to give the user data script time to complete before the instance enters the
`InService` state.

The example bash script in the user data customizes the hostname on instances to
include the launch template name and the instance ID. The script gets the ID of the
instance from the instance metadata. For more information, see [Retrieve instance\
metadata](../../../ec2/latest/userguide/instancedata-data-retrieval.md) in the _Amazon EC2 User Guide for Linux_
_Instances_.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Parameters": {
        "LatestAmiId": {
            "Description": "Region specific image from the Parameter Store",
            "Type": "AWS::SSM::Parameter::Value<AWS::EC2::Image::Id>",
            "Default": "/aws/service/ami-amazon-linux-latest/amzn2-ami-hvm-x86_64-gp2"
        },
        "InstanceType": {
            "Description": "Amazon EC2 instance type for the instances",
            "Type": "String",
            "AllowedValues": [
                "t3.micro",
                "t3.small",
                "t3.medium"
            ],
            "Default": "t3.micro"
        },
        "Subnets": {
            "Type": "CommaDelimitedList"
        }
    },
    "Resources": {
        "myAppLaunchTemplate": {
            "Type": "AWS::EC2::LaunchTemplate",
            "Properties": {
                "LaunchTemplateName": { "Fn::Sub": "${AWS::StackName}-launch-template" },
                "LaunchTemplateData": {
                    "ImageId": { "Ref": "LatestAmiId" },
                    "InstanceType": { "Ref": "InstanceType" },
                    "KeyName": "MyKeyPair",
                    "SecurityGroupIds": [  "sg-083cd3bfb8example" ],
                    "UserData": {"Fn::Base64": {"Fn::Join": [ "", [
                      "#!/usr/bin/env bash\n",
                      "set -e\n",
                      "export INSTANCE_ID=$(curl -sLf http://169.254.169.254/latest/meta-data/instance-id)", "\n",
                      "export NEW_HOSTNAME=\"${LaunchTemplateName}-$INSTANCE_ID\"", "\n",
                      "hostname $NEW_HOSTNAME", "\n"
                    ]]}}
                }
            }
        },
        "myASG": {
            "Type": "AWS::AutoScaling::AutoScalingGroup",
            "Properties": {
                "LaunchTemplate": {
                    "LaunchTemplateId": {
                        "Ref": "myAppLaunchTemplate"
                    },
                    "Version": {
                        "Fn::GetAtt": [
                            "myAppLaunchTemplate",
                            "DefaultVersionNumber"
                        ]
                    }
                },
                "MaxSize": 10,
                "MinSize": 1,
                "VPCZoneIdentifier": {
                    "Ref": "Subnets"
                },
                "LifecycleHookSpecificationList": [
                    {
                        "LifecycleTransition": "autoscaling:EC2_INSTANCE_LAUNCHING",
                        "LifecycleHookName": "myLaunchLifecycleHook",
                        "DefaultResult": "CONTINUE",
                        "HeartbeatTimeout": 60
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
Parameters:
  LatestAmiId:
    Description: Region specific image from the Parameter Store
    Type: 'AWS::SSM::Parameter::Value<AWS::EC2::Image::Id>'
    Default: '/aws/service/ami-amazon-linux-latest/amzn2-ami-hvm-x86_64-gp2'
  InstanceType:
    Description: Amazon EC2 instance type for the instances
    Type: String
    AllowedValues:
      - t3.micro
      - t3.small
      - t3.medium
    Default: t3.micro
  Subnets:
    Type: CommaDelimitedList
Resources:
  myAppLaunchTemplate:
    Type: AWS::EC2::LaunchTemplate
    Properties:
      LaunchTemplateName: !Sub ${AWS::StackName}-launch-template
      LaunchTemplateData:
        ImageId: !Ref LatestAmiId
        InstanceType: !Ref InstanceType
        KeyName: MyKeyPair
        SecurityGroupIds:
          - sg-083cd3bfb8example
        UserData:
          Fn::Base64: !Sub |
            #!/usr/bin/env bash
            set -e
            export INSTANCE_ID=$(curl -sLf http://169.254.169.254/latest/meta-data/instance-id)
            export NEW_HOSTNAME="${LaunchTemplateName}-$INSTANCE_ID"
            hostname $NEW_HOSTNAME
  myASG:
    Type: AWS::AutoScaling::AutoScalingGroup
    Properties:
      LaunchTemplate:
        LaunchTemplateId: !Ref myAppLaunchTemplate
        Version: !GetAtt myAppLaunchTemplate.DefaultVersionNumber
      MaxSize: '10'
      MinSize: '1'
      VPCZoneIdentifier: !Ref Subnets
      LifecycleHookSpecificationList:
        - LifecycleTransition: autoscaling:EC2_INSTANCE_LAUNCHING
          LifecycleHookName: myLaunchLifecycleHook
          DefaultResult: CONTINUE
          HeartbeatTimeout: '60'
```

### Lifecycle hook for instance termination

The following example specifies a lifecycle hook that supports a custom action at
instance termination. It uses the `Ref` intrinsic function to refer to an Auto
Scaling group (whose logical name is `myASG`) that is declared elsewhere in the
same template. It uses the `NotificationMetadata` property to specify
additional information to be sent with the notification, such as the name of the cluster
to which the instance belongs.

This snippet does not include the notification target that would need to exist or be
created to complete the configuration. For information about creating these resources, see
[Configure a notification target for lifecycle notifications](../../../autoscaling/ec2/userguide/prepare-for-lifecycle-notifications.md#lifecycle-hook-notification-target) in the
_Amazon EC2 Auto Scaling User Guide_.

#### JSON

```json

{
  "myTerminationLifecycleHook":{
    "Type":"AWS::AutoScaling::LifecycleHook",
    "Properties":{
      "AutoScalingGroupName":{
        "Ref":"myASG"
      },
      "LifecycleTransition":"autoscaling:EC2_INSTANCE_TERMINATING",
      "HeartbeatTimeout": 300,
      "DefaultResult": "CONTINUE",
      "NotificationMetadata": "optional metadata"
    }
  }
}
```

#### YAML

```yaml

---
myTerminationLifecycleHook:
  Type: AWS::AutoScaling::LifecycleHook
  Properties:
    AutoScalingGroupName: !Ref myASG
    LifecycleTransition: autoscaling:EC2_INSTANCE_TERMINATING
    HeartbeatTimeout: '300'
    DefaultResult: CONTINUE
    NotificationMetadata: optional metadata
```

### Lifecycle hook that specifies a notification target ARN and role ARN

The following example creates an Auto Scaling group with a lifecycle hook that
supports a custom action at instance launch. This example snippet uses the
`NotificationTargetARN` and `RoleARN` properties to specify the
Amazon SQS queue and IAM role to use to receive notification when a lifecycle action
occurs.

This snippet does not include the Amazon SQS queue and IAM role resources that would
need to exist or be created to complete the configuration. For information about creating
these resources, see [Configure a notification target for lifecycle notifications](../../../autoscaling/ec2/userguide/prepare-for-lifecycle-notifications.md#lifecycle-hook-notification-target) in the
_Amazon EC2 Auto Scaling User Guide_.

#### JSON

```json

{
  "AWSTemplateFormatVersion":"2010-09-09",
  "Parameters":{
    "Subnets":{
      "Type":"CommaDelimitedList"
    }
  },
  "Resources":{
    "myASG":{
      "Type":"AWS::AutoScaling::AutoScalingGroup",
      "Properties":{
        "MaxSize":"3",
        "MinSize":"1",
        "LaunchTemplate": {
          "LaunchTemplateId": {
            "Ref":"myLaunchTemplate"
          },
          "Version":{
            "Ref":"myLaunchTemplateVersionNumber"
          }
        },
        "VPCZoneIdentifier":{
          "Ref":"Subnets"
        },
        "LifecycleHookSpecificationList":[
          {
            "LifecycleTransition":"autoscaling:EC2_INSTANCE_LAUNCHING",
            "LifecycleHookName":"myLaunchLifecycleHook",
            "HeartbeatTimeout":4800,
            "NotificationTargetARN":{
              "Fn::GetAtt":[
                "SQS",
                "Arn"
              ]
            },
            "RoleArn":{
              "Fn::Join":[
                ":",
                [
                  "arn:aws:iam:",
                  {
                    "Ref":"AWS::AccountId"
                  },
                  "role/role-name"
                ]
              ]
            }
          }
        ]
      }
    },
    "SQS":{
      "Type":"AWS::SQS::Queue"
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Parameters:
  Subnets:
    Type: CommaDelimitedList
Resources:
  myASG:
    Type: AWS::AutoScaling::AutoScalingGroup
    Properties:
      MaxSize: '3'
      MinSize: '1'
      LaunchTemplate:
        LaunchTemplateId: !Ref myLaunchTemplate
        Version: !Ref myLaunchTemplateVersionNumber
      VPCZoneIdentifier: !Ref Subnets
      LifecycleHookSpecificationList:
        - LifecycleTransition: autoscaling:EC2_INSTANCE_LAUNCHING
          LifecycleHookName: myLaunchLifecycleHook
          HeartbeatTimeout: '4800'
          NotificationTargetARN: !GetAtt SQS.Arn
          RoleARN: !Join
            - ':'
            - - 'arn:aws:iam:'
              - !Ref 'AWS::AccountId'
              - role/role-name
  SQS:
    Type: AWS::SQS::Queue
```

## See also

- You can find additional useful snippets in the following sections of the _AWS CloudFormation User Guide_:

- For examples of Auto Scaling groups, see [Configure\
Amazon EC2 Auto Scaling resources](../userguide/quickref-ec2-auto-scaling.md).

- For examples of launch templates, see [Create\
launch templates](../userguide/quickref-ec2-launch-templates.md).

- [PutLifecycleHook](../../../../reference/autoscaling/ec2/apireference/api-putlifecyclehook.md) in the _Amazon EC2 Auto Scaling API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetadataOptions

AWS::AutoScaling::ScalingPolicy

All content copied from https://docs.aws.amazon.com/.
