This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeStarNotifications::NotificationRule

Creates a notification rule for a resource. The rule specifies the events you want
notifications about and the targets (such as Amazon Simple Notification Service topics or Amazon Q Developer in chat applications clients configured for Slack) where you want to receive
them.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CodeStarNotifications::NotificationRule",
  "Properties" : {
      "CreatedBy" : String,
      "DetailType" : String,
      "EventTypeId" : String,
      "EventTypeIds" : [ String, ... ],
      "Name" : String,
      "Resource" : String,
      "Status" : String,
      "Tags" : {Key: Value, ...},
      "TargetAddress" : String,
      "Targets" : [ Target, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CodeStarNotifications::NotificationRule
Properties:
  CreatedBy: String
  DetailType: String
  EventTypeId: String
  EventTypeIds:
    - String
  Name: String
  Resource: String
  Status: String
  Tags:
    Key: Value
  TargetAddress: String
  Targets:
    - Target

```

## Properties

`CreatedBy`

The name or email alias of the person who created the notification rule.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DetailType`

The level of detail to include in the notifications for this resource. `BASIC` will include only the
contents of the event as it would appear in Amazon CloudWatch. `FULL` will include any supplemental information
provided by AWS CodeStar Notifications and/or the service for the resource for which the notification is created.

_Required_: Yes

_Type_: String

_Allowed values_: `BASIC | FULL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventTypeId`

The event type associated with this notification rule. For a complete list of event types and IDs, see
[Notification concepts](../../../dtconsole/latest/userguide/concepts.md#concepts-api) in the
_Developer Tools Console User Guide_.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventTypeIds`

A list of event types associated with this notification rule. For a complete list of event types and IDs, see
[Notification concepts](../../../dtconsole/latest/userguide/concepts.md#concepts-api)
in the _Developer Tools Console User Guide_.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name for the notification rule. Notification rule names must be unique in your AWS account.

_Required_: Yes

_Type_: String

_Pattern_: `[A-Za-z0-9\-_ ]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Resource`

The Amazon Resource Name (ARN) of the resource to associate with the notification rule. Supported resources include pipelines in AWS CodePipeline,
repositories in AWS CodeCommit, and build projects in AWS CodeBuild.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[^:\s]*:[^:\s]*:[^:\s]*:[0-9]{12}:[^\s]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Status`

The status of the notification rule. The default value is `ENABLED`. If the status is
set to `DISABLED`, notifications aren't sent for the notification rule.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of tags to apply to this notification rule. Key names cannot start with " `aws`".

_Required_: No

_Type_: Object of String

_Pattern_: `.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetAddress`

The Amazon Resource Name (ARN) of the Amazon SNS topic or Amazon Q Developer in chat applications client.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Targets`

A list of Amazon Resource Names (ARNs) of Amazon SNS topics and Amazon Q Developer in chat applications clients to associate with the
notification rule.

_Required_: Yes

_Type_: Array of [Target](aws-properties-codestarnotifications-notificationrule-target.md)

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When the logical ID of this resource is provided to the Ref intrinsic function, `Ref` returns the notification rule ARN.

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the notification rule.

## Examples

### Example

The following example creates a notification rule with a name of My Notification Rule for Comments on Commits. The notification
rule is tagged with a key-value pair indicating what team owns the rule.

#### JSON

```json

{
    "Type": "AWS::CodeStarNotifications::NotificationRule",
    "Properties": {
        "Name": "My Notification Rule for Comments on Commits",
        "DetailType": "FULL",
        "Resource": "arn:aws:codecommit:us-east-2:123456789012:MyDemoRepo",
        "EventTypeIds": [
            "codecommit-repository-comments-on-commits"
        ],
        "Targets": [
            {
                "TargetType": "SNS",
                "TargetAddress": "arn:aws:sns:us-east-2:123456789012:MyNotificationTopic"
            }
        ],
        "Tags": {
            "Team": "Saanvi"
         }
     }
}
```

#### YAML

```yaml

Type: 'AWS::CodeStarNotifications::NotificationRule'
Properties:
        Name: 'My Notification Rule for Comments on Commits'
        DetailType: FULL
        Resource: 'arn:aws:codecommit:us-east-2:123456789012:MyDemoRepo'
        EventTypeIds:
            - codecommit-repository-comments-on-commits
        Targets:
            - TargetType: SNS
              TargetAddress: 'arn:aws:sns:us-east-2:123456789012:MyNotificationTopic'
        Tags:
             Team: Saanvi
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS CodeStar Notifications

Target

All content copied from https://docs.aws.amazon.com/.
