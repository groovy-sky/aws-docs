This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::Rule

Creates a rule for the specified Amazon Connect instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::Rule",
  "Properties" : {
      "Actions" : Actions,
      "Function" : String,
      "InstanceArn" : String,
      "Name" : String,
      "PublishStatus" : String,
      "Tags" : [ Tag, ... ],
      "TriggerEventSource" : RuleTriggerEventSource
    }
}

```

### YAML

```yaml

Type: AWS::Connect::Rule
Properties:
  Actions:
    Actions
  Function: String
  InstanceArn: String
  Name: String
  PublishStatus: String
  Tags:
    - Tag
  TriggerEventSource:
    RuleTriggerEventSource

```

## Properties

`Actions`

A list of actions to be run when the rule is triggered.

_Required_: Yes

_Type_: [Actions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-rule-actions.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Function`

The conditions of the rule.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceArn`

The Amazon Resource Name (ARN) of the instance.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the rule.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9._-]{1,200}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PublishStatus`

The publish status of the rule.

_Allowed values_: `DRAFT` \|
`PUBLISHED`

_Required_: Yes

_Type_: String

_Allowed values_: `DRAFT | PUBLISHED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags used to organize, track, or control access for this resource. For example, {
"tags": {"key1":"value1", "key2":"value2"} }.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-rule-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TriggerEventSource`

The event source to trigger the rule.

_Required_: Yes

_Type_: [RuleTriggerEventSource](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connect-rule-ruletriggereventsource.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the quick rule name. For example:

`{ "Ref": "myRuleName" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`RuleArn`

The Amazon Resource Name (ARN) of the rule.

## Examples

### Specify a rule resource

The following example specifies a rule resource for an Amazon Connect
instance.

#### JSON

```json

{
    "Name": "ExampleRule",
    "InstanceArn": "arn:aws:connect:region-name:aws-account-id:instance/instance-arn",
    "TriggerEventSource": {
        "EventSourceName": "OnPostCallAnalysisAvailable"
    },
    "Function": "Example function using Amazon Connect Rules Function Language.",
    "Actions": {
        "AssignContactCategoryActions": [{}],
        "EventBridgeActions": [
            {
                "Name": "Name of the event bridge action"
            }
        ],
        "SendNotificationActions": [
            {
                "DeliveryMethod": "EMAIL",
                "Subject": "Email subject",
                "Content": "Email content",
                "ContentType": "PLAIN_TEXT",
                "Recipient": {
                    "UserArns": ["arn:aws:connect:region-name:aws-account-id:instance/instance-arn/agent/user-arn"]
                }
            }
        ],
        "TaskActions": [
            {
              "Name": "Name of the task action",
              "ContactFlowArn": "arn:aws:connect:region-name:aws-account-id:instance/instance-arn/contact-flow/contact-flow-arn",
              "References": {
                "reference1": {
                  "Type": "URL",
                  "Value": "URL of the reference"
                }
              },
              "Description": "Task description"
            }
          ]
    },
    "PublishStatus": "DRAFT"
}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Actions
