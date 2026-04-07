This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::XRay::Group

Use the `AWS::XRay::Group` resource to specify a group with a name and a filter expression.
Groups enable the collection of traces that match the filter expression, can be used to filter service graphs and traces, and to supply Amazon CloudWatch metrics.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::XRay::Group",
  "Properties" : {
      "FilterExpression" : String,
      "GroupName" : String,
      "InsightsConfiguration" : InsightsConfiguration,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::XRay::Group
Properties:
  FilterExpression: String
  GroupName: String
  InsightsConfiguration:
    InsightsConfiguration
  Tags:
    - Tag

```

## Properties

`FilterExpression`

The filter expression defining the parameters to include traces.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GroupName`

The unique case-sensitive name of the group.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InsightsConfiguration`

The structure containing configurations related to insights.

- The InsightsEnabled boolean can be set to true to enable insights for the
group or false to disable insights for the group.

- The NotificationsEnabled boolean can be set to true to enable insights
notifications through Amazon EventBridge for the group.

_Required_: No

_Type_: [InsightsConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-xray-group-insightsconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-xray-group-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the group.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`GroupARN`

The group ARN that was created or updated.

## Examples

### Create group

This example creates a new group called MyGroup.

#### JSON

```json

{
   "AWSTemplateFormatVersion": "2010-09-09",
   "Resources": {
      "MyGroupResource": {
         "Type": "AWS::XRay::Group",
         "Properties": {
            "GroupName": "MyGroup",
            "FilterExpression": "duration > 10",
            "InsightsConfiguration": {
               "InsightsEnabled": "false",
               "NotificationsEnabled": "false"
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
   MyGroupResource:
      Type: AWS::XRay::Group
      Properties:
         GroupName: MyGroup
         FilterExpression: duration > 10
         InsightsConfiguration:
            InsightsEnabled: false
            NotificationsEnabled: false

```

## See also

- [Configuring groups in the X-Ray console](https://docs.aws.amazon.com/xray/latest/devguide/xray-console-groups.html)

- [Configuring groups with the X-Ray API](https://docs.aws.amazon.com/xray/latest/devguide/xray-api-configuration.html#xray-api-configuration-groups)

- [CreateGroup](https://docs.aws.amazon.com/xray/latest/api/API_CreateGroup.html) action in the X-Ray API Reference

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS X-Ray

InsightsConfiguration
