# Use EventBridge to handle Auto Scaling events

Amazon EventBridge, formerly called CloudWatch Events, helps you set up event-driven rules that monitor
resources and initiate target actions that use other AWS services.

Events from Amazon EC2 Auto Scaling are delivered to EventBridge in near real time. You can establish EventBridge rules
that invoke programmatic actions and notifications in response to a variety of these events.
For example, while instances are in the process of launching or terminating, you can invoke
an AWS Lambda function to perform a preconfigured task.

Targets of EventBridge rules can include AWS Lambda functions, Amazon SNS topics, API destinations,
event buses in other AWS accounts, and many more. For information about supported targets,
see [Amazon EventBridge targets](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-targets.html) in the _Amazon EventBridge User_
_Guide_.

Get started by creating EventBridge rules with an example using an Amazon SNS topic and an EventBridge rule.
Then, when a user starts an instance refresh, Amazon SNS notifies you by email whenever a
checkpoint is reached. For more information, see [Create EventBridge rules for instance refresh events](https://docs.aws.amazon.com/autoscaling/ec2/userguide/monitor-events-eventbridge-sns.html).

###### Contents

- [Amazon EC2 Auto Scaling event reference](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-event-reference.html)

- [Instance refresh example events and patterns](https://docs.aws.amazon.com/autoscaling/ec2/userguide/instance-refresh-eventbridge-events.html)

- [Warm pool example events and patterns](https://docs.aws.amazon.com/autoscaling/ec2/userguide/warm-pools-eventbridge-events.html)

- [Use Amazon EventBridge rules to automate actions](https://docs.aws.amazon.com/autoscaling/ec2/userguide/create-eventbridge-rules.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Verify the
attachment status

Amazon EC2 Auto Scaling event reference
