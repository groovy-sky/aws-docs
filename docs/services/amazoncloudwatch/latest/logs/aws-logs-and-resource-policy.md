---
title: "Enable logging from AWS services"
---

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
[Supported log destinations](aws-logs-destinations-table.md). For information about these required
permissions, see the sections after the table.

- [![Amazon API Gateway logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/api-gateway.png)Amazon API Gateway](../../../apigateway/latest/developerguide/set-up-logging.md)
- [![AWS AppSync logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/appsync.png)AWS AppSync](../../../appsync/latest/devguide/monitoring.md)
- [![Amazon Aurora MySQL logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/aurora.png)Amazon Aurora MySQL](../../../amazonrds/latest/aurorauserguide/auroramysql-integrating-cloudwatch.md)
- [![Amazon Bedrock Knowledge Bases logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/bedrock.png)Amazon Bedrock Knowledge Bases](../../../bedrock/latest/userguide/knowledge-bases-logging.md)
- [![Amazon Bedrock Agents logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/bedrock.png)Amazon Bedrock Agents](../../../bedrock/latest/userguide/model-invocation-logging.md)
- [![Amazon Bedrock AgentCore Runtime logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/bedrock-agentcore.png)Amazon Bedrock AgentCore Runtime](../../../bedrock-agentcore/latest/devguide/agents-tools-runtime.md)
- [![Amazon Bedrock AgentCore Gateway logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/bedrock-agentcore.png)Amazon Bedrock AgentCore Gateway](../../../bedrock-agentcore/latest/devguide/gateway.md)
- [![Amazon Bedrock AgentCore Identity logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/bedrock-agentcore.png)Amazon Bedrock AgentCore Identity](../../../bedrock-agentcore/latest/devguide/identity.md)
- [![Amazon Bedrock AgentCore Memory logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/bedrock-agentcore.png)Amazon Bedrock AgentCore Memory](../../../bedrock-agentcore/latest/devguide/memory.md)
- [![Amazon Bedrock AgentCore Payments logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/bedrock-agentcore.png)Amazon Bedrock AgentCore Payments](../../../bedrock-agentcore/latest/devguide/payments.md)
- [![Amazon Bedrock AgentCore Tools logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/bedrock-agentcore.png)Amazon Bedrock AgentCore Tools](../../../bedrock-agentcore/latest/devguide/built-in-tools.md)
- [![Amazon Chime logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/chime.png)Amazon Chime](../../../chime/latest/ag/monitoring-cloudwatch.md#cw-logs)
- [![Amazon CloudFront logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/cloudfront.png)Amazon CloudFront](../../../amazoncloudfront/latest/developerguide/accesslogs.md)
- [![AWS CloudHSM logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/cloudhsm.png)AWS CloudHSM](../../../cloudhsm/latest/userguide/get-hsm-audit-logs-using-cloudwatch.md)
- [![CloudWatch Evidently logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/cloudwatch.png)CloudWatch Evidently](../monitoring/cloudwatch-evidently-datastorage.md#CloudWatch-Evidently-datastorage-logformat)
- [![CloudWatch Internet Monitor logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/cloudwatch.png)CloudWatch Internet Monitor](../monitoring/cloudwatch-im-view-cw-tools-s3-athena.md)
- [![AWS CloudTrail logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/cloudtrail.png)AWS CloudTrail](../../../awscloudtrail/latest/userguide/monitor-cloudtrail-log-files-with-cloudwatch-logs.md)
- [![AWS CodeBuild logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/codebuild.png)AWS CodeBuild](../../../codebuild/latest/userguide/getting-started-build-log-console.md)
- [![Amazon CodeWhisperer logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/codewhisperer.png)Amazon CodeWhisperer](../../../eventbridge/latest/ref/events-ref-codewhisperer.md)
- [![Amazon Cognito logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/cognito.png)Amazon Cognito](../../../cognito/latest/developerguide/what-is-amazon-cognito.md)
- [![Amazon Connect logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/connect.png)Amazon Connect](../../../connect/latest/adminguide/logging-and-monitoring.md)
- [![AWS DataSync logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/datasync.png)AWS DataSync](../../../datasync/latest/userguide/monitor-datasync.md#cloudwatchlogs)
- [![AWS DevOps Agent logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/devops-agent.png)AWS DevOps Agent](../../../devopsagent/latest/userguide/configuring-capabilities-for-aws-devops-agent-vended-logs-and-metrics.md)
- [![Amazon ElastiCache (Redis OSS) logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/elasticache.png)Amazon ElastiCache (Redis OSS)](../../../amazonelasticache/latest/red-ug/log-delivery.md)
- [![AWS Elastic Beanstalk logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/elastic-beanstalk.png)AWS Elastic Beanstalk](../../../elasticbeanstalk/latest/dg/awshowto-cloudwatchlogs.md)
- [![Amazon ECS logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/ecs.png)Amazon ECS](../../../amazonecs/latest/developerguide/using-cloudwatch-logs.md)
- [![Amazon EKS Auto Mode logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/eks.png)Amazon EKS Auto Mode](../../../eks/latest/userguide/auto-managed-component-logs.md)
- [![Amazon EKS Control Plane logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/eks.png)Amazon EKS Control Plane](../../../eks/latest/userguide/control-plane-logs.md)
- [![AWS Elemental MediaPackage logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/mediapackage.png)AWS Elemental MediaPackage](../../../mediapackage/latest/ug/access-logging.md)
- [![AWS Elemental MediaTailor logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/mediatailor.png)AWS Elemental MediaTailor](../../../mediatailor/latest/ug/monitoring-cw-logs.md)
- [![AWS Entity Resolution logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/entity-resolution.png)AWS Entity Resolution](../../../entityresolution/latest/userguide/what-is-service.md)
- [![Amazon EventBridge Pipes logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/eventbridge.png)Amazon EventBridge Pipes](../../../eventbridge/latest/userguide/eb-pipes-logs.md)
- [![Amazon EventBridge Event Buses logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/eventbridge.png)Amazon EventBridge Event Buses](../../../eventbridge/latest/userguide/eb-pipes-logs.md)
- [![AWS Fargate logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/fargate.png)AWS Fargate](../../../amazonecs/latest/developerguide/using-awslogs.md)
- [![AWS Fault Injection Service logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/fis.png)AWS Fault Injection Service](../../../fis/latest/userguide/monitoring-logging.md)
- [![Amazon FinSpace logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/finspace.png)Amazon FinSpace](../../../finspace/latest/userguide/finspace-what-is.md)
- [![AWS Global Accelerator logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/global-accelerator.png)AWS Global Accelerator](../../../global-accelerator/latest/dg/monitoring-global-accelerator-flow-logs.md)
- [![AWS Glue logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/glue.png)AWS Glue](../../../glue/latest/dg/monitor-continuous-logging.md)
- [![IAM Identity Center logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/iam-identity-center.png)IAM Identity Center](../../../singlesignon/latest/userguide/logging-ad-sync-errors.md)
- [![Amazon IVS Chat logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/ivs.png)Amazon IVS Chat](../../../ivs/latest/lowlatencyuserguide/chat-logging.md)
- [![AWS IoT logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/iot.png)AWS IoT](../../../iot/latest/developerguide/cloud-watch-logs.md)
- [![AWS IoT FleetWise logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/iot-fleetwise.png)AWS IoT FleetWise](../../../iot-fleetwise/latest/developerguide/logging-cw.md)
- [![AWS Lambda logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/lambda.png)AWS Lambda](../../../lambda/latest/dg/monitoring-cloudwatchlogs.md)
- [![Amazon Macie logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/macie.png)Amazon Macie](../../../macie/latest/user/discovery-jobs-monitor-cw-logs.md)
- [![Amazon SES logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/ses.png)Amazon SES](../../../ses/latest/dg/eb-logging.md)
- [![AWS Mainframe Modernization logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/mainframe.png)AWS Mainframe Modernization](../../../m2/latest/userguide/what-is-m2.md)
- [![Amazon Managed Service for Prometheus logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/prometheus.png)Amazon Managed Service for Prometheus](../../../prometheus/latest/userguide/cw-logs.md)
- [![Amazon MSK logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/msk.png)Amazon MSK](../../../msk/latest/developerguide/msk-logging.md)
- [![Amazon MSK Connect logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/msk-connect.png)Amazon MSK Connect](../../../msk/latest/developerguide/msk-connect-logging.md)
- [![Amazon MQ logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/mq.png)Amazon MQ](../../../amazon-mq/latest/developer-guide/configure-logging-monitoring-activemq.md)
- [![AWS Network Firewall logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/network-firewall.png)AWS Network Firewall](../../../network-firewall/latest/developerguide/firewall-logging.md)
- [![AWS Network Firewall Proxy logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/network-firewall.png)AWS Network Firewall Proxy](../../../network-firewall/latest/developerguide/proxy-logging-and-monitoring.md)
- [![Network Load Balancer logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/nlb.png)Network Load Balancer](../../../elasticloadbalancing/latest/network/load-balancer-access-logs.md)
- [![Amazon OpenSearch Service logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/opensearch.png)Amazon OpenSearch Service](../../../opensearch-service/latest/developerguide/createdomain-configure-slow-logs.md)
- [![Amazon OpenSearch Ingestion logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/opensearch-ingestion.png)Amazon OpenSearch Ingestion](../../../opensearch-service/latest/developerguide/monitoring-pipeline-logs.md)
- [![AWS PCS logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/pcs.png)AWS PCS](../../../pcs/latest/userguide/monitoring-overview.md)
- [![Amazon Q Business Connectors logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/q-business.png)Amazon Q Business Connectors](../../../amazonq/latest/qbusiness-ug/connectors-list.md)
- [![Amazon Q Business Conversations logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/q-business.png)Amazon Q Business Conversations](../../../amazonq/latest/qbusiness-ug/cw-logs-enable-logging.md)
- [![Amazon Quick Chat and Feedback logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/quick-chat.png)Amazon Quick Chat and Feedback](../../../quicksuite/latest/userguide/monitoring-quicksuite-chat-feedback-cloudwatch.md)
- [![Amazon RDS PostgreSQL logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/rds.png)Amazon RDS PostgreSQL](../../../amazonrds/latest/userguide/user-logaccess-concepts-postgresql.md#USER_LogAccess.PostgreSQL.PublishtoCloudWatchLogs)
- [![AWS RTB Fabric logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/rtb-fabric.png)AWS RTB Fabric](../../../rtb-fabric/latest/userguide/what-is-rtb-fabric.md)
- [![AWS Security Hub CSPM logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/security-hub.png)AWS Security Hub CSPM](../../../securityhub/latest/userguide/what-is-securityhub.md)
- [![AWS Security Hub logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/security-hub.png)AWS Security Hub](../../../securityhub/latest/userguide/what-is-securityhub-v2.md)
- [![Amazon Route 53 Public DNS logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/route53.png)Amazon Route 53 Public DNS](../../../route53/latest/developerguide/logging-monitoring.md)
- [![Amazon Route 53 Resolver logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/route53.png)Amazon Route 53 Resolver](../../../route53/latest/developerguide/resolver-query-logs-choosing-target-resource.md)
- [![Amazon SageMaker AI Events logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/sagemaker.png)Amazon SageMaker AI Events](../../../sagemaker/latest/dg/logging-cloudwatch.md)
- [![Amazon SageMaker AI Worker Events logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/sagemaker.png)Amazon SageMaker AI Worker Events](../../../sagemaker/latest/dg/workteam-private-tracking.md)
- [![AWS Site-to-Site VPN logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/site-to-site-vpn.png)AWS Site-to-Site VPN](../../../vpn/latest/s2svpn/monitoring-logs.md)
- [![Amazon SES logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/ses.png)Amazon SES](../../../ses/latest/dg/eb-logging.md)
- [![Amazon SNS logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/sns.png)Amazon SNS](../../../sns/latest/dg/sms-stats-cloudwatch.md#sns-viewing-cloudwatch-logs)
- [![Amazon SNS Data Protection logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/sns.png)Amazon SNS Data Protection](../../../sns/latest/dg/sns-message-data-protection-operations.md)
- [![EC2 Spot Instance logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/ec2.png)EC2 Spot Instance](../../../ec2/latest/userguide/spot-data-feeds.md)
- [![AWS Step Functions logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/step-functions.png)AWS Step Functions](../../../step-functions/latest/dg/cw-logs.md)
- [![AWS Storage Gateway logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/storage-gateway.png)AWS Storage Gateway](../../../storagegateway/latest/userguide/monitoring-file-gateway.md)
- [![AWS Transfer Family logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/transfer-family.png)AWS Transfer Family](../../../transfer/latest/userguide/structured-logging.md)
- [![AWS Verified Access logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/verified-access.png)AWS Verified Access](../../../verified-access/latest/ug/access-logs.md)
- [![Amazon VPC Flow Logs logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/vpc.png)Amazon VPC Flow Logs](../../../vpc/latest/userguide/flow-logs-s3.md)
- [![Amazon VPC Lattice logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/vpc-lattice.png)Amazon VPC Lattice](../../../vpc-lattice/latest/ug/monitoring-access-logs.md)
- [![Amazon VPC Route Server logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/vpc-route-server.png)Amazon VPC Route Server](../../../vpc/latest/userguide/dynamic-routing-route-server.md)
- [![AWS WAF logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/waf.png)AWS WAF](../../../waf/latest/developerguide/logging-destinations.md)
- [![Amazon WorkMail logo](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/images/integration-icons/workmail.png)Amazon WorkMail](../../../workmail/latest/adminguide/monitoring-audit-logging.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Filter pattern syntax

Supported log destinations

All content copied from https://docs.aws.amazon.com/.
