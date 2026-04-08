# Getting started with CloudWatch Logs

To collect logs from your Amazon EC2 instances and on-premises servers into CloudWatch Logs, use the unified CloudWatch agent. It
enables you to collect both logs and advanced metrics with one agent. It offers
support across operating systems, including servers running Windows Server. This
agent also provides better performance.

If you're using the unified CloudWatch agent to collect CloudWatch metrics, it enables the
collection of additional system metrics, for in-guest visibility. It also supports
collecting custom metrics using `StatsD` or `collectd`.

For more information, see [Installing the CloudWatch Agent](../monitoring/install-cloudwatch-agent-on-ec2-instance.md) in the
_Amazon CloudWatch User Guide_.

The older CloudWatch Logs agent, which
supports only the collection of logs from
servers running Linux, is deprecated and is no longer supported. For information about migrating from the older CloudWatch Logs agent to the unified agent, see
[Create the CloudWatch agent configuration file with the wizard](../monitoring/create-cloudwatch-agent-configuration-file-wizard.md).

###### Contents

- [Prerequisites](gettingsetup-cwl.md)

- [Using the unified CloudWatch agent](usecloudwatchunifiedagent.md)

- [Using the previous CloudWatch agent](usepreviouscloudwatchlogsagent.md)

- [Quick Start with CloudFormation](quickstartcloudformation.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Log classes

Prerequisites

All content copied from https://docs.aws.amazon.com/.
