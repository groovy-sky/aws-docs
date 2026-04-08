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

Log sourceLog type[Logs sent to CloudWatch Logs](aws-logs-infrastructure-cwl.md)[Logs sent to Amazon S3](aws-logs-infrastructure-s3.md)[Logs sent to Firehose](aws-logs-infrastructure-firehose.md)[Traces sent to X-Ray](aws-logs-infrastructure-v2-xraytraces.md)

[Amazon API Gateway access logs](../../../apigateway/latest/developerguide/set-up-logging.md)

Vended logs

[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[AWS AppSync logs](../../../appsync/latest/devguide/monitoring.md)

Custom logs

Supported

[Amazon Aurora MySQL logs](../../../amazonrds/latest/aurorauserguide/auroramysql-integrating-cloudwatch.md)

Custom logs

Supported

[Amazon Bedrock Knowledge bases logging](../../../bedrock/latest/userguide/knowledge-bases-logging.md)

Vended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)

[Amazon Bedrock Agent logging](../../../bedrock/latest/userguide/model-invocation-logging.md)

Vended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)

[Amazon Bedrock AgentCore Runtime](../../../bedrock-agentcore/latest/devguide/agents-tools-runtime.md)

Vended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)

[Amazon Bedrock AgentCore Gateway](../../../bedrock-agentcore/latest/devguide/gateway.md)

Vended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)

[Amazon Bedrock AgentCore Identity](../../../bedrock-agentcore/latest/devguide/identity.md)

Vended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)

[Amazon Bedrock AgentCore Memory](../../../bedrock-agentcore/latest/devguide/memory.md)

Vended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)

[Amazon Bedrock AgentCore Tools](../../../bedrock-agentcore/latest/devguide/built-in-tools.md)

Vended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)

[Amazon Chime\
media quality metric logs and SIP message logs](../../../chime/latest/ag/monitoring-cloudwatch.md#cw-logs)

Vended logs

[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[CloudFront:\
access logs](../../../amazoncloudfront/latest/developerguide/accesslogs.md)

Vended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)

[AWS CloudHSM audit logs](../../../cloudhsm/latest/userguide/get-hsm-audit-logs-using-cloudwatch.md)

Custom logs

Supported

[CloudWatch Evidently evaluation event logs](../monitoring/cloudwatch-evidently-datastorage.md#CloudWatch-Evidently-datastorage-logformat)

Vended logs[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[CloudWatch Internet Monitor logs](../monitoring/cloudwatch-im-view-cw-tools-s3-athena.md)

Vended logs[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[CloudTrail logs](../../../awscloudtrail/latest/userguide/monitor-cloudtrail-log-files-with-cloudwatch-logs.md)

Custom logs

Supported

[AWS CodeBuild logs](../../../codebuild/latest/userguide/getting-started-build-log-console.md)

Custom logs

Supported

[Amazon CodeWhisperer event logs](../../../eventbridge/latest/ref/events-ref-codewhisperer.md)

Vended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)

[Amazon Cognito logs](../../../cognito/latest/developerguide/what-is-amazon-cognito.md)

Vended logs[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[Amazon Connect\
logs](../../../connect/latest/adminguide/logging-and-monitoring.md)

Custom logs

Supported

[AWS DataSync logs](../../../datasync/latest/userguide/monitor-datasync.md#cloudwatchlogs)

Custom logs

Supported

[AWS DevOps Agent logs](../../../devopsagent/latest/userguide/configuring-capabilities-for-aws-devops-agent-vended-logs-and-metrics.md)

Vended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)

[Amazon ElastiCache (Redis OSS)\
logs](../../../amazonelasticache/latest/red-ug/log-delivery.md)

Vended logs[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[AWS Elastic Beanstalk\
logs](../../../elasticbeanstalk/latest/dg/awshowto-cloudwatchlogs.md)

Custom logs

Supported

[Amazon Elastic Container Service\
logs](../../../amazonecs/latest/developerguide/using-cloudwatch-logs.md)

Custom logs

Supported

[Amazon Elastic Kubernetes Service Auto Mode\
logs](../../../eks/latest/userguide/auto-managed-component-logs.md)

Vended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)

[Amazon Elastic Kubernetes Service control\
plane logs](../../../eks/latest/userguide/control-plane-logs.md)

Vended logs

Supported

[AWS Elemental MediaPackage access logs](../../../mediapackage/latest/ug/access-logging.md)

Vended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)

[AWS Elemental MediaTailor logs](../../../mediatailor/latest/ug/monitoring-cw-logs.md)

Vended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[AWS Entity Resolution\
logs](../../../entityresolution/latest/userguide/what-is-service.md)Vended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)

[Amazon EventBridge Pipes logging](../../../eventbridge/latest/userguide/eb-pipes-logs.md)

Vended logs[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[Amazon EventBridge event buses](../../../eventbridge/latest/userguide/eb-pipes-logs.md)

Vended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)

[AWS Fargate\
logs](../../../amazonecs/latest/developerguide/using-awslogs.md)

Custom logs

Supported

[AWS Fault Injection Service\
experiment logs](../../../fis/latest/userguide/monitoring-logging.md)

Vended logs[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[Amazon FinSpace](../../../finspace/latest/userguide/finspace-what-is.md)

Vended logs[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[AWS Global Accelerator flow logs](../../../global-accelerator/latest/dg/monitoring-global-accelerator-flow-logs.md)

Vended logs[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[AWS Glue\
job logs](../../../glue/latest/dg/monitor-continuous-logging.md)

Custom logs

Supported

[IAM Identity Center error logs](../../../singlesignon/latest/userguide/logging-ad-sync-errors.md)

Vended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)

[Amazon Interactive Video Service\
chat logs](../../../ivs/latest/lowlatencyuserguide/chat-logging.md)

Vended logs[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[AWS IoT\
logs](../../../iot/latest/developerguide/cloud-watch-logs.md)

Custom logs

Supported

[AWS IoT FleetWise logs](../../../iot-fleetwise/latest/developerguide/logging-cw.md)

Vended logs[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[AWS Lambda\
logs](../../../lambda/latest/dg/monitoring-cloudwatchlogs.md)

Vended logs

Supported

Supported

Supported

[Amazon Macie logs](../../../macie/latest/user/discovery-jobs-monitor-cw-logs.md)

Custom logs

Supported

[Amazon SES\
logs](../../../ses/latest/dg/eb-logging.md)Vended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)

[AWS Mainframe Modernization](../../../m2/latest/userguide/what-is-m2.md)

Vended logs[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[Amazon Managed Service for Prometheus\
logs](../../../prometheus/latest/userguide/cw-logs.md)

Vended logs

[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[Amazon MSK broker\
logs](../../../msk/latest/developerguide/msk-logging.md)

Vended logs

[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[Amazon MSK\
Connect logs](../../../msk/latest/developerguide/msk-connect-logging.md)

Vended logs

[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[Amazon MQ logs](../../../amazon-mq/latest/developer-guide/configure-logging-monitoring-activemq.md)

Custom logs

Supported

[AWS Network Firewall logs](../../../network-firewall/latest/developerguide/firewall-logging.md)

Vended logs

[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[AWS Network Firewall Proxy logs](../../../network-firewall/latest/developerguide/proxy-logging-and-monitoring.md)

Vended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)

[Network Load Balancer access logs](../../../elasticloadbalancing/latest/network/load-balancer-access-logs.md)

Vended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)

[OpenSearch logs](../../../opensearch-service/latest/developerguide/createdomain-configure-slow-logs.md)

Custom logs

Supported

[Amazon OpenSearch Service ingestion logs](../../../opensearch-service/latest/developerguide/monitoring-pipeline-logs.md)

Vended logs[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)[AWS PCS](../../../pcs/latest/userguide/monitoring-overview.md)
logsVended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)

[Amazon Q Business connector logs](../../../amazonq/latest/qbusiness-ug/connectors-list.md)

Vended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)

[Amazon Q Business conversation logs](../../../amazonq/latest/qbusiness-ug/cw-logs-enable-logging.md)

Vended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)

[Amazon Quick chat and feedback logs](../../../quicksuite/latest/userguide/monitoring-quicksuite-chat-feedback-cloudwatch.md)

Vended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)

[Amazon Relational Database ServicePostgreSQL logs](../../../amazonrds/latest/userguide/user-logaccess-concepts-postgresql.md#USER_LogAccess.PostgreSQL.PublishtoCloudWatchLogs)

Custom logs

Supported

[AWS RTB Fabric](../../../rtb-fabric/latest/userguide/what-is-rtb-fabric.md)
logsVended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)

[AWS Security Hub CSPM](../../../securityhub/latest/userguide/what-is-securityhub.md)

Vended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)

[Amazon Route 53 public\
DNS query logs](../../../route53/latest/developerguide/logging-monitoring.md)

Vended logs

Supported

[Amazon Route 53 resolver query logs](../../../route53/latest/developerguide/resolver-query-logs-choosing-target-resource.md)

Vended logs

[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[Amazon SageMaker AI\
events](../../../sagemaker/latest/dg/logging-cloudwatch.md)

Vended logs

[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[Amazon SageMaker AI\
worker events](../../../sagemaker/latest/dg/workteam-private-tracking.md)

Vended logs

[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[AWS Site-to\_Site VPN\
logs](../../../vpn/latest/s2svpn/monitoring-logs.md)

Vended logs

[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[Amazon Simple Email Service logs](../../../ses/latest/dg/eb-logging.md)

Vended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)

[Amazon Simple Notification Service logs](../../../sns/latest/dg/sms-stats-cloudwatch.md#sns-viewing-cloudwatch-logs)

Custom logs

Supported

[Amazon Simple Notification Service data protection policy logs](../../../sns/latest/dg/sns-message-data-protection-operations.md)

Custom logs

Supported

[EC2 Spot\
Instance data feed files](../../../ec2/latest/userguide/spot-data-feeds.md)

Vended logs

[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[AWS Step Functions Express\
Workflow and Standard Workflow logs](../../../step-functions/latest/dg/cw-logs.md)

Vended logs

[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[Storage Gateway audit logs and health logs](../../../storagegateway/latest/userguide/monitoring-file-gateway.md)

Vended logs

[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[AWS Transfer Family logs](../../../transfer/latest/userguide/structured-logging.md)

Vended logs

[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[AWS Verified Access logs](../../../verified-access/latest/ug/access-logs.md)

Vended logs

[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[Amazon Virtual Private Cloud flow\
logs](../../../vpc/latest/userguide/flow-logs-s3.md)

Vended logs

Supported

[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

[Amazon VPC\
Lattice access logs](../../../vpc-lattice/latest/ug/monitoring-access-logs.md)

Vended logs[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)[Amazon VPC Route\
Server](../../../vpc/latest/userguide/dynamic-routing-route-server.md)Vended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)

[AWS WAF\
logs](../../../waf/latest/developerguide/logging-destinations.md)

Vended logs[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)[Supported \[V1\
Permissions\]](aws-vended-logs-permissions.md)

Supported

[Amazon WorkMail audit logs](../../../workmail/latest/adminguide/monitoring-audit-logging.md)

Vended logs[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)[Supported \[V2\
Permissions\]](aws-vended-logs-permissions-v2.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Filter pattern syntax

Logging that requires additional permissions \[V1\]

All content copied from https://docs.aws.amazon.com/.
