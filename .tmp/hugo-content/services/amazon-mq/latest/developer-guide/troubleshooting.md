# Troubleshooting Amazon MQ

This section describes common issues you might encounter when using Amazon MQ brokers, and the steps you can take to resolve them.
For general troubleshooting, see [Troubleshooting: General Amazon MQ](general.md). For troubleshooting your specific
engine version, see the following sections.

## Troubleshooting ActiveMQ on Amazon MQ

Troubleshooting topicDescription[General troubleshooting](troubleshooting-activemq.md)Use the information in this section to help you diagnose
and resolve common issues you might encounter when working with ActiveMQ on Amazon MQ brokers.[BROKER\_ENI\_DELETED](troubleshooting-action-required-codes-broker-eni-deleted.md)ActiveMQ on Amazon MQ will raise a `BROKER_ENI_DELETED` alarm when you delete a broker’s Elastic Network Interface (ENI). [BROKER\_OOM](troubleshooting-action-required-codes-broker-out-of-memory.md)ActiveMQ on Amazon MQ will raise a BROKER\_OOM alarm when the broker undergoes a restart loop due to
the insufficient memory capacity

## Troubleshooting RabbitMQ on Amazon MQ

Troubleshooting topicDescription[General troubleshooting](troubleshooting-rabbitmq.md)Diagnose common issues
you might encounter when working with RabbitMQ brokers.[RABBITMQ\_MEMORY\_ALARM](troubleshooting-action-required-codes-rabbitmq-memory-alarm.md)RabbitMQ will raise a high memory alarm when the broker's memory usage, identified by CloudWatch metric
`RabbitMQMemUsed`, exceeds the memory limit, identified by `RabbitMQMemLimit`.[RABBITMQ\_INVALID\_KMS\_KEY](troubleshooting-action-required-codes-invalid-kms-key.md)RabbitMQ on Amazon MQ will raise an INVALID\_KMS\_KEY critical action required code when a broker created
with a customer managed AWS KMS key(CMK) detects that the AWS Key Management Service (KMS) key is disabled. [RABBITMQ\_INVALID\_ASSUMEROLE](troubleshooting-action-required-codes-invalid-assumerole.md)RabbitMQ on Amazon MQ will raise an INVALID\_ASSUMEROLE critical action required code when the IAM role ARN specified in `aws.arns.assume_role_arn` cannot be assumed by Amazon MQ.[RABBITMQ\_INVALID\_ARN\_LDAP](troubleshooting-action-required-codes-invalid-arn-ldap.md)RabbitMQ on Amazon MQ will raise an INVALID\_ARN\_LDAP critical action required code when the LDAP service account password ARN is invalid or inaccessible.[RABBITMQ\_INVALID\_ARN\_HTTP](troubleshooting-action-required-codes-invalid-arn-http.md)RabbitMQ on Amazon MQ will raise an INVALID\_ARN\_HTTP critical action required code when one or more ARNs of SSL certificates or key file for HTTP auth\_backend are invalid or inaccessible.[RABBITMQ\_INVALID\_ARN\_SSL](troubleshooting-action-required-codes-invalid-arn-ssl.md)RabbitMQ on Amazon MQ will raise an INVALID\_ARN\_SSL critical action required code when one or more ARNs of CA certificate truststore for EXTERNAL auth\_mechanism are invalid or inaccessible.[RABBITMQ\_INVALID\_ARN](troubleshooting-action-required-codes-invalid-arn.md)RabbitMQ on Amazon MQ will raise an INVALID\_ARN critical action required code when one or more ARNs in the broker configuration are invalid or inaccessible.[RABBITMQ\_DISK\_ALARM](troubleshooting-action-required-codes-disk-limit-alarm.md)Disk limit alarm is an indication that the volume of disk used by a RabbitMQ node has decreased due to
a high number of messages not consumed while new messages were added.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Quotas

Troubleshooting: General Amazon MQ

All content copied from https://docs.aws.amazon.com/.
