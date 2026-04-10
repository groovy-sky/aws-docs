This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Config::OrganizationConfigRule OrganizationCustomRuleMetadata

An object that specifies organization custom rule metadata such as resource type, resource ID of AWS resource, Lambda function ARN,
and organization trigger types that trigger AWS Config to evaluate your AWS resources against a rule.
It also provides the frequency with which you want AWS Config to run evaluations for the rule if the trigger type is periodic.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Description" : String,
  "InputParameters" : String,
  "LambdaFunctionArn" : String,
  "MaximumExecutionFrequency" : String,
  "OrganizationConfigRuleTriggerTypes" : [ String, ... ],
  "ResourceIdScope" : String,
  "ResourceTypesScope" : [ String, ... ],
  "TagKeyScope" : String,
  "TagValueScope" : String
}

```

### YAML

```yaml

  Description: String
  InputParameters: String
  LambdaFunctionArn: String
  MaximumExecutionFrequency: String
  OrganizationConfigRuleTriggerTypes:
    - String
  ResourceIdScope: String
  ResourceTypesScope:
    - String
  TagKeyScope: String
  TagValueScope: String

```

## Properties

`Description`

The description that you provide for your organization AWS Config rule.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputParameters`

A string, in JSON format, that is passed to your organization AWS Config rule Lambda function.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaFunctionArn`

The lambda function ARN.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumExecutionFrequency`

The maximum frequency with which AWS Config runs evaluations for a rule.
Your custom rule is triggered when AWS Config delivers the configuration snapshot. For more information, see [ConfigSnapshotDeliveryProperties](../userguide/aws-resource-config-deliverychannel.md#cfn-config-deliverychannel-configsnapshotdeliveryproperties).

###### Note

By default, rules with a periodic trigger are evaluated every 24 hours. To change the frequency, specify a valid
value for the `MaximumExecutionFrequency` parameter.

_Required_: No

_Type_: String

_Allowed values_: `One_Hour | Three_Hours | Six_Hours | Twelve_Hours | TwentyFour_Hours`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrganizationConfigRuleTriggerTypes`

The type of notification that triggers AWS Config to run an evaluation for a rule. You can specify the following notification types:

- `ConfigurationItemChangeNotification` \- Triggers an evaluation when AWS Config delivers a configuration item as a result of a resource change.

- `OversizedConfigurationItemChangeNotification` \- Triggers an evaluation when AWS Config delivers an oversized configuration item.
AWS Config may generate this notification type when a resource changes and the notification exceeds the maximum size allowed by Amazon SNS.

- `ScheduledNotification` \- Triggers a periodic evaluation at the frequency specified for `MaximumExecutionFrequency`.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceIdScope`

The ID of the AWS resource that was evaluated.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `768`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceTypesScope`

The type of the AWS resource that was evaluated.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagKeyScope`

One part of a key-value pair that make up a tag.
A key is a general label that acts like a category for more specific tag values.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagValueScope`

The optional part of a key-value pair that make up a tag.
A value acts as a descriptor within a tag category (key).

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OrganizationCustomPolicyRuleMetadata

OrganizationManagedRuleMetadata

All content copied from https://docs.aws.amazon.com/.
