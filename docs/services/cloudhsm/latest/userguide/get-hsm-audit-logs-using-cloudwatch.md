# Working with Amazon CloudWatch Logs and AWS CloudHSM Audit Logs

When an HSM in your account receives a command from the AWS CloudHSM [command line tools](https://docs.aws.amazon.com/cloudhsm/latest/userguide/command-line-tools.html) or [software\
libraries](https://docs.aws.amazon.com/cloudhsm/latest/userguide/use-hsm.html), it records its execution of the command in audit log form. The HSM audit logs
include all client-initiated [management\
commands](https://docs.aws.amazon.com/cloudhsm/latest/userguide/cloudhsm-audit-log-reference.html), including those that create and delete the HSM, log into and out of the HSM,
and manage users and keys. These logs provide a reliable record of actions that have changed the
state of the HSM.

AWS CloudHSM collects your HSM audit logs and sends them to [Amazon CloudWatch Logs](../../../amazoncloudwatch/latest/logs/whatiscloudwatchlogs.md) on your behalf. You can use
the features of CloudWatch Logs to manage your AWS CloudHSM Audit Logs, including searching and filtering the
logs and exporting log data to Amazon S3. You can work with your HSM audit logs in the [Amazon CloudWatch console](https://console.aws.amazon.com/cloudwatch) or use the CloudWatch Logs commands in
the [AWS CLI](https://docs.aws.amazon.com/cli/latest/reference/logs/index.html) and [CloudWatch Logs\
SDKs](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference).

###### Topics

- [How HSM audit logging works](https://docs.aws.amazon.com/cloudhsm/latest/userguide/get-audit-logs-from-cloudwatch.html)

- [Viewing AWS CloudHSM audit logs in CloudWatch Logs](https://docs.aws.amazon.com/cloudhsm/latest/userguide/understand-audit-logs.html)

- [Interpreting AWS CloudHSM audit logs](https://docs.aws.amazon.com/cloudhsm/latest/userguide/interpreting-audit-logs.html)

- [AWS CloudHSM audit log reference](https://docs.aws.amazon.com/cloudhsm/latest/userguide/cloudhsm-audit-log-reference.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS CloudTrail

How logging works
