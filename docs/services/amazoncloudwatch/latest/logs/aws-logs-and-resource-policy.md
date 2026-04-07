# Enable logging from AWS services

While many services publish logs only to CloudWatch Logs, some AWS services can publish logs
directly to Amazon Simple Storage Service or Amazon Data Firehose. If your main requirement for logs is storage or
processing in one of these services, you can easily have the service that produces the logs
send them directly to Amazon S3 or Firehose without additional setup.

Even when you publish logs directly to Amazon S3 or Firehose, CloudWatch delivery charges
apply. If you send logs to Amazon S3, then
`AWS_REGION-S3-Egress-Bytes` charges appear in Cost
Explorer or on your bill. If you send logs to Firehose, then
`AWS_REGION-FH-Egress-Bytes` charges appear. For more
information about vended logs pricing, see the **Logs** tab at
[Amazon CloudWatch Pricing](https://aws.amazon.com/cloudwatch/pricing).

Some AWS services use a common infrastructure to send their logs. To enable logging from
these services, you must be logged in as a user that has certain permissions. Additionally,
you must grant permissions to AWS to enable the logs to be sent.

For services that require these permissions, there are two versions of the permissions
needed. The services that require these extra permissions are noted as **Supported**
**\[V1 Permissions\]** and **Supported \[V2 Permissions\]** in the
table. For information about these required permissions, see the sections after the
table.

Log sourceLog type[Logs sent to CloudWatch Logs](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-logs-infrastructure-CWL.html)[Logs sent to Amazon S3](aws-logs-infrastructure-s3.md)[Logs sent to Firehose](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-logs-infrastructure-Firehose.html)[Traces sent to X-Ray](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-logs-infrastructure-V2-XRayTraces.html)

[Amazon API Gateway access logs](../../../apigateway/latest/developerguide/set-up-logging.md)

Vended logs

[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[AWS AppSync logs](../../../appsync/latest/devguide/monitoring.md)

Custom logs

Supported

[Amazon Aurora MySQL logs](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/AuroraMySQL.Integrating.CloudWatch.html)

Custom logs

Supported

[Amazon Bedrock Knowledge bases logging](https://docs.aws.amazon.com/bedrock/latest/userguide/knowledge-bases-logging.html)

Vended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)

[Amazon Bedrock Agent logging](https://docs.aws.amazon.com/bedrock/latest/userguide/model-invocation-logging.html)

Vended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)

[Amazon Bedrock AgentCore Runtime](https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/agents-tools-runtime.html)

Vended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)

[Amazon Bedrock AgentCore Gateway](https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/gateway.html)

Vended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)

[Amazon Bedrock AgentCore Identity](https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/identity.html)

Vended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)

[Amazon Bedrock AgentCore Memory](https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/memory.html)

Vended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)

[Amazon Bedrock AgentCore Tools](https://docs.aws.amazon.com/bedrock-agentcore/latest/devguide/built-in-tools.html)

Vended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)

[Amazon Chime\
media quality metric logs and SIP message logs](https://docs.aws.amazon.com/chime/latest/ag/monitoring-cloudwatch.html#cw-logs)

Vended logs

[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[CloudFront:\
access logs](../../../amazoncloudfront/latest/developerguide/accesslogs.md)

Vended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)

[AWS CloudHSM audit logs](https://docs.aws.amazon.com/cloudhsm/latest/userguide/get-hsm-audit-logs-using-cloudwatch.html)

Custom logs

Supported

[CloudWatch Evidently evaluation event logs](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Evidently-datastorage.html#CloudWatch-Evidently-datastorage-logformat)

Vended logs[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[CloudWatch Internet Monitor logs](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-IM-view-cw-tools.S3_athena.html)

Vended logs[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[CloudTrail logs](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/monitor-cloudtrail-log-files-with-cloudwatch-logs.html)

Custom logs

Supported

[AWS CodeBuild logs](https://docs.aws.amazon.com/codebuild/latest/userguide/getting-started-build-log-console.html)

Custom logs

Supported

[Amazon CodeWhisperer event logs](https://docs.aws.amazon.com/eventbridge/latest/ref/events-ref-codewhisperer.html)

Vended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)

[Amazon Cognito logs](https://docs.aws.amazon.com/cognito/latest/developerguide/what-is-amazon-cognito.html)

Vended logs[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[Amazon Connect\
logs](https://docs.aws.amazon.com/connect/latest/adminguide/logging-and-monitoring.html)

Custom logs

Supported

[AWS DataSync logs](https://docs.aws.amazon.com/datasync/latest/userguide/monitor-datasync.html#cloudwatchlogs)

Custom logs

Supported

[AWS DevOps Agent logs](https://docs.aws.amazon.com/devopsagent/latest/userguide/configuring-capabilities-for-aws-devops-agent-vended-logs-and-metrics.html)

Vended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)

[Amazon ElastiCache (Redis OSS)\
logs](https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Log_Delivery.html)

Vended logs[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[AWS Elastic Beanstalk\
logs](https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/AWSHowTo.cloudwatchlogs.html)

Custom logs

Supported

[Amazon Elastic Container Service\
logs](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_cloudwatch_logs.html)

Custom logs

Supported

[Amazon Elastic Kubernetes Service Auto Mode\
logs](https://docs.aws.amazon.com/eks/latest/userguide/auto-managed-component-logs.html)

Vended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)

[Amazon Elastic Kubernetes Service control\
plane logs](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html)

Vended logs

Supported

[AWS Elemental MediaPackage access logs](https://docs.aws.amazon.com/mediapackage/latest/ug/access-logging.html)

Vended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)

[AWS Elemental MediaTailor logs](https://docs.aws.amazon.com/mediatailor/latest/ug/monitoring-cw-logs.html)

Vended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[AWS Entity Resolution\
logs](https://docs.aws.amazon.com/entityresolution/latest/userguide/what-is-service.html)Vended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)

[Amazon EventBridge Pipes logging](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-logs.html)

Vended logs[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[Amazon EventBridge event buses](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-pipes-logs.html)

Vended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)

[AWS Fargate\
logs](../../../amazonecs/latest/developerguide/using-awslogs.md)

Custom logs

Supported

[AWS Fault Injection Service\
experiment logs](https://docs.aws.amazon.com/fis/latest/userguide/monitoring-logging.html)

Vended logs[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[Amazon FinSpace](https://docs.aws.amazon.com/finspace/latest/userguide/finspace-what-is.html)

Vended logs[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[AWS Global Accelerator flow logs](https://docs.aws.amazon.com/global-accelerator/latest/dg/monitoring-global-accelerator.flow-logs.html)

Vended logs[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[AWS Glue\
job logs](https://docs.aws.amazon.com/glue/latest/dg/monitor-continuous-logging.html)

Custom logs

Supported

[IAM Identity Center error logs](https://docs.aws.amazon.com/singlesignon/latest/userguide/logging-ad-sync-errors.html)

Vended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)

[Amazon Interactive Video Service\
chat logs](https://docs.aws.amazon.com/ivs/latest/LowLatencyUserGuide/chat-logging.html)

Vended logs[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[AWS IoT\
logs](https://docs.aws.amazon.com/iot/latest/developerguide/cloud-watch-logs.html)

Custom logs

Supported

[AWS IoT FleetWise logs](https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/logging-cw.html)

Vended logs[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[AWS Lambda\
logs](https://docs.aws.amazon.com/lambda/latest/dg/monitoring-cloudwatchlogs.html)

Vended logs

Supported

Supported

Supported

[Amazon Macie logs](https://docs.aws.amazon.com/macie/latest/user/discovery-jobs-monitor-cw-logs.html)

Custom logs

Supported

[Amazon SES\
logs](https://docs.aws.amazon.com/ses/latest/dg/eb-logging.html)Vended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)

[AWS Mainframe Modernization](https://docs.aws.amazon.com/m2/latest/userguide/what-is-m2.html)

Vended logs[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[Amazon Managed Service for Prometheus\
logs](https://docs.aws.amazon.com/prometheus/latest/userguide/CW-logs.html)

Vended logs

[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[Amazon MSK broker\
logs](https://docs.aws.amazon.com/msk/latest/developerguide/msk-logging.html)

Vended logs

[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[Amazon MSK\
Connect logs](https://docs.aws.amazon.com/msk/latest/developerguide/msk-connect-logging.html)

Vended logs

[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[Amazon MQ logs](../../../amazon-mq/latest/developer-guide/configure-logging-monitoring-activemq.md)

Custom logs

Supported

[AWS Network Firewall logs](https://docs.aws.amazon.com/network-firewall/latest/developerguide/firewall-logging.html)

Vended logs

[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[AWS Network Firewall Proxy logs](https://docs.aws.amazon.com/network-firewall/latest/developerguide/proxy-logging-and-monitoring.html)

Vended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)

[Network Load Balancer access logs](https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-access-logs.html)

Vended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)

[OpenSearch logs](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createdomain-configure-slow-logs.html)

Custom logs

Supported

[Amazon OpenSearch Service ingestion logs](https://docs.aws.amazon.com/opensearch-service/latest/developerguide/monitoring-pipeline-logs.html)

Vended logs[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)[AWS PCS](https://docs.aws.amazon.com/pcs/latest/userguide/monitoring-overview.html)
logsVended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)

[Amazon Q Business connector logs](../../../amazonq/latest/qbusiness-ug/connectors-list.md)

Vended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)

[Amazon Q Business conversation logs](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/cw-logs-enable-logging.html)

Vended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)

[Amazon Quick chat and feedback logs](https://docs.aws.amazon.com/quicksuite/latest/userguide/monitoring-quicksuite-chat-feedback-cloudwatch.html)

Vended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)

[Amazon Relational Database ServicePostgreSQL logs](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_LogAccess.Concepts.PostgreSQL.html#USER_LogAccess.PostgreSQL.PublishtoCloudWatchLogs)

Custom logs

Supported

[AWS RTB Fabric](https://docs.aws.amazon.com/rtb-fabric/latest/userguide/what-is-rtb-fabric.html)
logsVended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)

[AWS Security Hub CSPM](https://docs.aws.amazon.com/securityhub/latest/userguide/what-is-securityhub.html)

Vended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)

[Amazon Route 53 public\
DNS query logs](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/logging-monitoring.html)

Vended logs

Supported

[Amazon Route 53 resolver query logs](../../../route53/latest/developerguide/resolver-query-logs-choosing-target-resource.md)

Vended logs

[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[Amazon SageMaker AI\
events](https://docs.aws.amazon.com/sagemaker/latest/dg/logging-cloudwatch.html)

Vended logs

[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[Amazon SageMaker AI\
worker events](https://docs.aws.amazon.com/sagemaker/latest/dg/workteam-private-tracking.html)

Vended logs

[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[AWS Site-to\_Site VPN\
logs](https://docs.aws.amazon.com/vpn/latest/s2svpn/monitoring-logs.html)

Vended logs

[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[Amazon Simple Email Service logs](https://docs.aws.amazon.com/ses/latest/dg/eb-logging.html)

Vended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)

[Amazon Simple Notification Service logs](https://docs.aws.amazon.com/sns/latest/dg/sms_stats_cloudwatch.html#sns-viewing-cloudwatch-logs)

Custom logs

Supported

[Amazon Simple Notification Service data protection policy logs](https://docs.aws.amazon.com/sns/latest/dg/sns-message-data-protection-operations.html)

Custom logs

Supported

[EC2 Spot\
Instance data feed files](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-data-feeds.html)

Vended logs

[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[AWS Step Functions Express\
Workflow and Standard Workflow logs](https://docs.aws.amazon.com/step-functions/latest/dg/cw-logs.html)

Vended logs

[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[Storage Gateway audit logs and health logs](https://docs.aws.amazon.com/storagegateway/latest/userguide/monitoring-file-gateway.html)

Vended logs

[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[AWS Transfer Family logs](https://docs.aws.amazon.com/transfer/latest/userguide/structured-logging.html)

Vended logs

[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[AWS Verified Access logs](https://docs.aws.amazon.com/verified-access/latest/ug/access-logs.html)

Vended logs

[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[Amazon Virtual Private Cloud flow\
logs](https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs-s3.html)

Vended logs

Supported

[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

[Amazon VPC\
Lattice access logs](https://docs.aws.amazon.com/vpc-lattice/latest/ug/monitoring-access-logs.html)

Vended logs[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)[Amazon VPC Route\
Server](https://docs.aws.amazon.com/vpc/latest/userguide/dynamic-routing-route-server.html)Vended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)

[AWS WAF\
logs](https://docs.aws.amazon.com/waf/latest/developerguide/logging-destinations.html)

Vended logs[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)[Supported \[V1\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions.html)

Supported

[Amazon WorkMail audit logs](https://docs.aws.amazon.com/workmail/latest/adminguide/monitoring-audit-logging.html)

Vended logs[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)[Supported \[V2\
Permissions\]](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AWS-vended-logs-permissions-V2.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Filter pattern syntax

Logging that requires additional permissions \[V1\]
