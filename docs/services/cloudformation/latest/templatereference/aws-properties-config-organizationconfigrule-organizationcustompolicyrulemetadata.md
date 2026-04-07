This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Config::OrganizationConfigRule OrganizationCustomPolicyRuleMetadata

An
object that specifies metadata for your organization's AWS Config Custom Policy rule. The metadata includes the runtime system in use, which accounts have
debug logging enabled, and other custom rule metadata, such as resource type, resource
ID of AWS resource, and organization trigger types that initiate AWS Config to evaluate AWS resources against a rule.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DebugLogDeliveryAccounts" : [ String, ... ],
  "Description" : String,
  "InputParameters" : String,
  "MaximumExecutionFrequency" : String,
  "OrganizationConfigRuleTriggerTypes" : [ String, ... ],
  "PolicyText" : String,
  "ResourceIdScope" : String,
  "ResourceTypesScope" : [ String, ... ],
  "Runtime" : String,
  "TagKeyScope" : String,
  "TagValueScope" : String
}

```

### YAML

```yaml

  DebugLogDeliveryAccounts:
    - String
  Description: String
  InputParameters: String
  MaximumExecutionFrequency: String
  OrganizationConfigRuleTriggerTypes:
    - String
  PolicyText: String
  ResourceIdScope: String
  ResourceTypesScope:
    - String
  Runtime: String
  TagKeyScope: String
  TagValueScope: String

```

## Properties

`DebugLogDeliveryAccounts`

A list of accounts that you can enable debug logging for your organization AWS Config Custom Policy rule. List is null when debug logging is enabled for all accounts.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description that you provide for your organization AWS Config Custom Policy rule.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputParameters`

A string, in JSON format, that is passed to your organization AWS Config Custom Policy rule.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumExecutionFrequency`

The maximum frequency with which AWS Config runs evaluations for a rule. Your
AWS Config Custom Policy rule is triggered when AWS Config delivers
the configuration snapshot. For more information, see [ConfigSnapshotDeliveryProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html#cfn-config-deliverychannel-configsnapshotdeliveryproperties).

_Required_: No

_Type_: String

_Allowed values_: `One_Hour | Three_Hours | Six_Hours | Twelve_Hours | TwentyFour_Hours`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrganizationConfigRuleTriggerTypes`

The type of notification that initiates AWS Config to run an evaluation for a rule.
For AWS Config Custom Policy rules, AWS Config supports change-initiated notification types:

- `ConfigurationItemChangeNotification` \- Initiates an evaluation when AWS Config delivers a configuration item as a result of a resource
change.

- `OversizedConfigurationItemChangeNotification` \- Initiates an evaluation when
AWS Config delivers an oversized configuration item. AWS Config may generate this notification type when a resource changes and the
notification exceeds the maximum size allowed by Amazon SNS.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyText`

The policy definition containing the logic for your organization AWS Config Custom Policy rule.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `10000`

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

`Runtime`

The runtime system for your organization AWS Config Custom Policy rules. Guard is a policy-as-code language that allows you to write policies that are enforced by AWS Config Custom Policy rules. For more information about Guard, see the [Guard GitHub\
Repository](https://github.com/aws-cloudformation/cloudformation-guard).

_Required_: Yes

_Type_: String

_Pattern_: `guard\-2\.x\.x`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagKeyScope`

One part of a key-value pair that make up a tag. A key is a general label that acts like a category for more specific tag values.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TagValueScope`

The optional part of a key-value pair that make up a tag. A value acts as a descriptor within a tag category (key).

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Config::OrganizationConfigRule

OrganizationCustomRuleMetadata
