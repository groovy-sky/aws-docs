This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::MitigationAction

Defines an action that can be applied to audit findings by using
StartAuditMitigationActionsTask. For API reference, see [CreateMitigationAction](https://docs.aws.amazon.com/iot/latest/apireference/API_CreateMitigationAction.html) and for general information, see [Mitigation actions](https://docs.aws.amazon.com/iot/latest/developerguide/dd-mitigation-actions.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::MitigationAction",
  "Properties" : {
      "ActionName" : String,
      "ActionParams" : ActionParams,
      "RoleArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::IoT::MitigationAction
Properties:
  ActionName: String
  ActionParams:
    ActionParams
  RoleArn: String
  Tags:
    - Tag

```

## Properties

`ActionName`

The friendly name of the mitigation action.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9:_-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ActionParams`

The set of parameters for this mitigation action. The parameters vary, depending on the kind of action you apply.

_Required_: Yes

_Type_: [ActionParams](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-mitigationaction-actionparams.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The IAM role ARN used to apply this mitigation action.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata that can be used to manage the mitigation action.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-mitigationaction-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the mitigation action name.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`MitigationActionArn`

The Amazon Resource Name (ARN) of the mitigation action.

`MitigationActionId`

The ID of the mitigation action.

## Examples

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Description": "Amazon Web Services IoT MitigationAction Sample Template",
  "Resources": {
    "PublishToSnsMitigationAction": {
      "Type": "AWS::IoT::MitigationAction",
      "Properties": {
        "ActionName": "PublishToSns",
        "RoleArn": "arn:aws:us-east-1:123456789012:iam:role/RoleForIoTMitigationActions",
        "ActionParams": {
          "PublishFindingToSnsParams": {
            "TopicArn": "arn:aws:sns:us-east-1:123456789012:IoTFindingNotifications"
          }
        }
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Description: Amazon Web Services IoT MitigationAction Sample Template
Resources:
  PublishToSnsMitigationAction:
    Type: AWS::IoT::MitigationAction
    Properties:
      ActionName: PublishToSns
      RoleArn: arn:aws:us-east-1:123456789012:iam:role/RoleForIoTMitigationActions
      ActionParams:
        PublishFindingToSnsParams:
          TopicArn: arn:aws:sns:us-east-1:123456789012:IoTFindingNotifications
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EventConfiguration

ActionParams
