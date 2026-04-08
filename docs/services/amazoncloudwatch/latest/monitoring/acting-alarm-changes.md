# Acting on alarm changes

CloudWatch can notify users on two types of alarm changes: when an alarm changes state, and when
the configuration of an alarm gets updated.

When an alarm evaluates, it might change from one state to another, such as ALARM or OK.
For Metrics Insights alarms that monitor multiple time series, each time series (contributor)
can only be in ALARM or OK state, never in INSUFFICIENT\_DATA state. This is because a time
series only exists when data is present.

Additionally, CloudWatch sends events to Amazon EventBridge whenever alarms change state, and when alarms
are created, deleted, or updated. You can write EventBridge rules to take actions or be notified when
EventBridge receives these events.

For more information about alarm actions, see [Alarm actions](alarm-actions.md).

###### Topics

- [Notifying users on alarm changes](notify-users-alarm-changes.md)

- [Invoke a Lambda function from an alarm](alarms-and-actions-lambda.md)

- [Start a CloudWatch investigations from an alarm](start-investigation-alarm.md)

- [Stop, terminate, reboot, or recover an EC2 instance](usingalarmactions.md)

- [Alarm events and EventBridge](cloudwatch-and-eventbridge.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Create a composite alarm

Notifying users on alarm changes

All content copied from https://docs.aws.amazon.com/.
