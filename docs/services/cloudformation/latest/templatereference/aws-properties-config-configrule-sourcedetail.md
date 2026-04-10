This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Config::ConfigRule SourceDetail

Provides the source and the message types that trigger AWS Config to evaluate your AWS resources against a rule. It also
provides the frequency with which you want AWS Config to run
evaluations for the rule if the trigger type is periodic. You can
specify the parameter values for `SourceDetail` only for
custom rules.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EventSource" : String,
  "MaximumExecutionFrequency" : String,
  "MessageType" : String
}

```

### YAML

```yaml

  EventSource: String
  MaximumExecutionFrequency: String
  MessageType: String

```

## Properties

`EventSource`

The source of the event, such as an AWS service, that triggers
AWS Config to evaluate your AWS resources.

_Required_: Yes

_Type_: String

_Allowed values_: `aws.config`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumExecutionFrequency`

The frequency at which you want AWS Config to run evaluations
for a custom rule with a periodic trigger. If you specify a value
for `MaximumExecutionFrequency`, then
`MessageType` must use the
`ScheduledNotification` value.

###### Note

By default, rules with a periodic trigger are evaluated
every 24 hours. To change the frequency, specify a valid value
for the `MaximumExecutionFrequency`
parameter.

Based on the valid value you choose, AWS Config runs
evaluations once for each valid value. For example, if you
choose `Three_Hours`, AWS Config runs evaluations
once every three hours. In this case, `Three_Hours`
is the frequency of this rule.

_Required_: No

_Type_: String

_Allowed values_: `One_Hour | Three_Hours | Six_Hours | Twelve_Hours | TwentyFour_Hours`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MessageType`

The type of notification that triggers AWS Config to run an
evaluation for a rule. You can specify the following notification
types:

- `ConfigurationItemChangeNotification` \- Triggers
an evaluation when AWS Config delivers a configuration item
as a result of a resource change.

- `OversizedConfigurationItemChangeNotification`
\- Triggers an evaluation when AWS Config delivers an
oversized configuration item. AWS Config may generate this
notification type when a resource changes and the
notification exceeds the maximum size allowed by Amazon
SNS.

- `ScheduledNotification` \- Triggers a
periodic evaluation at the frequency specified for
`MaximumExecutionFrequency`.

- `ConfigurationSnapshotDeliveryCompleted` -
Triggers a periodic evaluation when AWS Config delivers a
configuration snapshot.

If you want your custom rule to be triggered by configuration
changes, specify two SourceDetail objects, one for
`ConfigurationItemChangeNotification` and one for
`OversizedConfigurationItemChangeNotification`.

_Required_: Yes

_Type_: String

_Allowed values_: `ConfigurationItemChangeNotification | ConfigurationSnapshotDeliveryCompleted | ScheduledNotification | OversizedConfigurationItemChangeNotification`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Source

AWS::Config::ConfigurationAggregator

All content copied from https://docs.aws.amazon.com/.
