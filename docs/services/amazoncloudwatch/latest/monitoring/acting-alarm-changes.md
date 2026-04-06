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

For more information about alarm actions, see [Alarm actions](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/alarm-actions.html).

###### Topics

- [Notifying users on alarm changes](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Notify_Users_Alarm_Changes.html)

- [Invoke a Lambda function from an alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/alarms-and-actions-Lambda.html)

- [Start a CloudWatch investigations from an alarm](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Start-Investigation-Alarm.html)

- [Stop, terminate, reboot, or recover an EC2 instance](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/UsingAlarmActions.html)

- [Alarm events and EventBridge](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch-and-eventbridge.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Create a composite alarm

Notifying users on alarm changes
