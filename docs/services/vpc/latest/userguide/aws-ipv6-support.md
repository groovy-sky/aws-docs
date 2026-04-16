---
title: "AWS services that support IPv6"
---

# AWS services that support IPv6

Computers and smart devices use IP addresses to communicate with each other over the
internet and other networks. As the internet continues to grow, so does the need for IP
addresses. The most common format for IP addresses is IPv4. The new format for IP addresses
is IPv6, which provides a larger address space than IPv4.

AWS services support for IPv6 includes support for dual stack configuration (IPv4 and
IPv6) or IPv6 only configurations. For example, a virtual private cloud (VPC) is a logically
isolated section of the AWS Cloud where you can launch AWS resources. Within a VPC, you
can create subnets that are IPv4 only, dual stack, or IPv6 only.

AWS services support access through public endpoints. Some AWS services also support
access using private endpoints powered by AWS PrivateLink. AWS services can support IPv6
through their private endpoints even if they do not support IPv6 through their public
endpoints. Endpoints that support IPv6 can respond to DNS queries with AAAA records.

## Services that support IPv6

The following table lists the AWS services that provide dual stack support, IPv6
only support, and endpoints that support IPv6. We will update this table as we release
additional support for IPv6. For the specifics about how a service supports IPv6, refer
to the documentation for the service.

Service nameDual stack supportIPv6 only supportPublic endpoints support IPv6Private endpoints support IPv6 1AWS AmplifyYesNoYesAmazon API GatewayYesNoYesYesAWS App MeshYesYesYesNoAWS App RunnerYesNoYesYesAWS AppConfig[Yes](../../../appconfig/latest/userguide/setting-up-ipv6.md)NoYesYesApplication Auto ScalingNoNoYesYesAWS Application Discovery ServiceYesNoYesYesApplication Recovery Controller (ARC)YesNoYesAmazon WorkSpaces ApplicationsYesNoNoNoAWS AppSync2PartialNoPartialYesAmazon AthenaYesNoYes[Yes](https://aws.amazon.com/about-aws/whats-new/2023/05/amazon-athena-ipv6-endpoints-inbound-connections)AWS Audit ManagerNoNoYesYesAmazon Aurora[Yes](../../../amazonrds/latest/aurorauserguide/user-vpc-workingwithrdsinstanceinavpc.md#USER_VPC.IP_addressing)NoYesNoAmazon Aurora DSQLNoNoYesYesAWS Auto ScalingNoNoYesYesAWS B2B Data InterchangeYesNoYesYesAWS BackupYesNo[Yes](../../../../general/latest/gr/bk.md)[Yes](../../../aws-backup/latest/devguide/backup-network.md#backup-privatelink)AWS Batch[Yes](../../../batch/latest/userguide/vpc-interface-endpoints.md)NoYesYesAmazon BedrockNoNoYesYesAWS Billing and Cost Management Data ExportsYesNoYesYesAWS Billing and Cost Management Pricing CalculatorYesNoYesYesAWS Billing ConductorYesNoYesYesAWS BudgetsYesNoYesAmazon BraketYesYesYesYesAWS Certificate ManagerYesNoYesNoAmazon Chime SDKYesNoYesAmazon ComprehendYesYesYesYesAWS Clean RoomsYesYesYesYesAWS Clean Rooms MLYesYesYesYesAWS Cloud9[Yes](../../../cloud9/latest/user-guide/vpc-settings.md)NoYesAWS Cloud Control APIYesNoYesYesCloudFormationNoNoYesYesAmazon CloudFront[Yes](../../../amazoncloudfront/latest/developerguide/cloudfront-enable-ipv6.md)[Yes](../../../amazoncloudfront/latest/developerguide/cloudfront-enable-ipv6.md)[Yes](../../../../general/latest/gr/cf-region.md)AWS CloudHSMYesNo[Yes](../../../cloudhsm/latest/userguide/ip-access.md)[Yes](../../../cloudhsm/latest/userguide/cloudhsm-vpc-endpoint.md)AWS CloudTrailYesNoYesYesAmazon CloudWatchYesYesYesYesAmazon CloudWatch Application InsightsNoNoYesYesAmazon CloudWatch Internet MonitorNoNoYesYesAmazon CloudWatch LogsYesYesYesYesAmazon CloudWatch Observability Access ManagerYesYesYesAmazon CloudWatch SyntheticsYesNo[Yes](https://aws.amazon.com/about-aws/whats-new/2025/01/amazon-cloudwatch-synthetics-ipv6-support)YesAWS Cloud Map[Yes](../../../cloud-map/latest/dg/registering-instances.md)YesYesYesAWS Cloud WANYesNoYesNoAWS CodeArtifactYesNoYesYesAmazon Connect Customer ProfilesYesNoYesYesAWS CodeBuildNoNoYesYesAWS CodeCommitNoNoYesYesAWS CodeDeployNoNoYesYesAWS Compute OptimizerNoNoYesYesAmazon Comprehend MedicalNoNoYesYesAmazon CodeGuru ProfilerYesNoYesYesAmazon CognitoYesNoYesAWS ConfigNoNoYesYesAWS Control TowerNoNoYesYesAWS Cost ExplorerYesNoYesYesAWS Cost Optimization HubYesNoYesYesAWS Data ExchangeNoNoYesYesAmazon Data FirehoseNoNoYesYesAmazon Data Lifecycle ManagerYesNoYesYesAWS Database Migration Service[Yes](../../../dms/latest/userguide/chap-replicationinstance-ipaddressing.md)NoNoYesAWS DataSync[Yes](../../../datasync/latest/userguide/datasync-network.md#ipv6-support)YesYesYesAmazon DataZoneNoNoYesYesAWS Deadline CloudYesNoYes[Yes](../../../deadline-cloud/latest/userguide/vpc-interface-endpoints.md)Amazon DetectiveYesYes[Yes](../../../../general/latest/gr/detective.md)Direct ConnectYesYesNoYesDirectory ServiceNoNoYesYesAmazon EBS direct APIsYesNoYesYesAmazon EC2[Yes](../../../ec2/latest/userguide/using-instance-addressing.md#ipv6-addressing)Yes[Yes](../../../ec2/latest/devguide/ec2-endpoints.md)NoEC2 Image BuilderYesYesYesYesAmazon ECRYesNo[Yes](https://aws.amazon.com/about-aws/whats-new/2025/05/amazon-ecr-support-ipv6)NoAmazon ECS[Yes](../../../amazonecs/latest/developerguide/task-networking-awsvpc.md)YesYesYesAmazon EFS[Yes](../../../efs/latest/ug/manage-fs-access-create-delete-mount-targets.md)Yes[Yes](https://aws.amazon.com/about-aws/whats-new/2025/06/amazon-efs-internet-protocol-version-6)YesAmazon EKS[Partial](../../../eks/latest/userguide/network-reqs.md#network-requirements-ip-table)[Partial](../../../eks/latest/userguide/network-reqs.md#network-requirements-ip-table)YesYesAmazon EMRNoNoYesYesAWS Elastic Beanstalk[Yes](../../../elasticbeanstalk/latest/dg/environments-cfg-elbv2-ipv6-dualstack.md)No[Yes](../../../elasticbeanstalk/latest/dg/vpc-vpce.md#vpc-vpce.ipv6)[Yes](../../../elasticbeanstalk/latest/dg/vpc-vpce.md#vpc-vpce.ipv6)AWS Elastic Disaster RecoveryNoNoYesYesElastic Load Balancing[Partial](../../../elasticloadbalancing/latest/userguide/how-elastic-load-balancing-works.md#ip-address-types)[Partial](../../../elasticloadbalancing/latest/userguide/how-elastic-load-balancing-works.md#ip-address-types)YesNoAmazon ElastiCache[Yes](../../../amazonelasticache/latest/dg/network-type.md)YesNoYesAWS Elemental MediaConvertNoNoYesYesAWS Elemental MediaConnectYesYesYesPartialAWS End User Messaging SocialYesNoYesNoAWS Entity ResolutionYesNoYesYesAmazon EventBridgeNoNoYesYesAWS Fargate[Yes](../../../amazonecs/latest/developerguide/fargate-task-networking.md)NoYesYesAmazon FSxNoNo[Yes](../../../../general/latest/gr/fsxn.md)[Yes](../../../fsx/latest/ontapguide/fsx-vpc-endpoints.md)Amazon GameLift StreamsYesNo[Yes](../../../../general/latest/gr/gameliftstreams.md)YesAWS Global Accelerator[Yes](https://aws.amazon.com/about-aws/whats-new/2022/07/aws-global-accelerator-announces-ipv6-support)NoYesAWS GlueYesNoNoYesAWS Glue DataBrewNoNoYesYesAmazon Managed Grafana 3YesNoYesYesAWS Ground Station 4YesNoYesYesAmazon GuardDutyNoNoYesYesAWS HealthImagingNoNoYesYesAWS HealthLakeNoNoYesYesAWS HealthOmicsNoNoYesYesAWS Identity and Access Management (IAM)[Yes](../../../iam/latest/userguide/reference-dual-stack-endpoint-support.md)YesYesNoAWS IAM Access Analyzer[Yes](https://aws.amazon.com/about-aws/whats-new/2025/03/iam-access-analyzer-supports-ipv6)NoYesYesAWS IAM Identity CenterYesNoYesAWS IAM Roles AnywhereNoNoYesYesAmazon InspectorYesYesYesYesAmazon Interactive Video Service (IVS) 5YesNoYesYesAWS IoT CoreYesNo[Yes](../../../iot/latest/developerguide/protocols.md)YesAWS IoT Device DefenderYesNoYesNoAWS IoT Device ManagementYesNoYesNoAWS IoT FleetWiseYesNo[Yes](../../../iot-fleetwise/latest/developerguide/fleetwise-ipv6-access.md)YesAWS IoT GreengrassYesNoYesNoAWS IoT SiteWiseYesNoYesYesAWS IoT TwinMakerYesNoYesYesAWS IoT WirelessYesNo[Yes](../../../iot-wireless/latest/developerguide/wireless-ipv6-access.md)[Yes](../../../iot-wireless/latest/developerguide/vpc-interface-endpoints.md)Amazon KendraNoNoYesNoAWS Key Management Service[Yes](../../../kms/latest/developerguide/ipv6-kms.md)[Partial](../../../kms/latest/developerguide/ipv6-kms.md)YesYesAmazon Keyspaces[Yes](../../../keyspaces/latest/devguide/ipv6-support.md)YesYesYesAmazon Keyspaces CDC streamsYesYesYesYesAmazon Kinesis Data StreamsYesNoYesYesAWS Lake FormationNoNoNoYesAWS Lambda[Yes](../../../lambda/latest/dg/configuration-vpc.md#configuration-vpc-ipv6)No[Yes](https://aws.amazon.com/about-aws/whats-new/2021/12/aws-lambda-ipv6-endpoints-inbound-connections)YesAWS Launch WizardNoNoYesYesAWS License ManagerNoNoYesYesAmazon Lightsail[Yes](https://aws.amazon.com/about-aws/whats-new/2021/01/amazon-lightsail-supports-ipv6)[Yes](https://aws.amazon.com/about-aws/whats-new/2024/01/ipv6-instance-bundles-amazon-lightsail)[Yes](https://aws.amazon.com/about-aws/whats-new/2024/12/amazon-lightsail-api-endpoints-connectivity-ipv6)YesAmazon Location ServiceNoNoYesYesAmazon MQNoNoYesYesAmazon MWAANoNoYesYesAmazon MacieYesNoYesYesAWS Mainframe ModernizationYesNoYesYesAmazon Managed GrafanaNoNoYesYesAmazon Managed Service for PrometheusYesNoYesYesAWS Migration Hub OrchestratorNoNoYesYesAWS Network Firewall[Yes](https://aws.amazon.com/about-aws/whats-new/2023/01/aws-network-firewall-ipv6-support)[Yes](https://aws.amazon.com/about-aws/whats-new/2023/04/aws-network-firewall-ipv6-only-subnets)NoYesAWS Network ManagerYesNoYesNoAmazon OpenSearch Service[Yes](https://aws.amazon.com/about-aws/whats-new/2023/10/amazon-opensearch-service-ipv6)NoYesAWS OrganizationsYesNoYesYesAWS OutpostsNoNoYesYesAmazon PersonalizeYesNoYesYesAmazon PinpointYesNoYesYesAmazon PollyYesNoYesYesAWS Price ListNoNoYesNoAWS Private Certificate AuthorityYesNoYesYesAWS Private CA Connector for Active DirectoryYesNoYesYesAWS Private CA Connector for SCEPYesNoYesYesAWS PrivateLinkYesYesYesAmazon Q BusinessNoNoYesNoAmazon RDS[Yes](../../../amazonrds/latest/userguide/user-vpc-workingwithrdsinstanceinavpc.md#USER_VPC.IP_addressing)NoYesNoAmazon RDS Data APINoNoYesYesAmazon RDS Performance InsightsNoNoYesYesAmazon RedshiftYesNoYesAmazon RekognitionNoNoYesYesRecycle BinYesNoYesYesAWS re:Post PrivateYesNoYesYesAWS Resource Access ManagerYesNoYesYesAWS Resource ExplorerYesNoYesNoAWS Resource GroupsYesYesYesYesAWS Resource Groups Tagging APIYesYesYesYesAmazon Route 53YesYesYesYesAmazon S3[Yes](../../../s3/latest/api/ipv6-access.md)No[Yes](../../../s3/latest/api/ipv6-access.md)NoAmazon S3 on OutpostsNoNoYesNoAmazon SageMakerNoNoYesYesAWS Secrets ManagerYesNo[Yes](../../../secretsmanager/latest/userguide/asm-access.md#endpoints)YesAWS Security HubNoNoYesYesAmazon Security Lake[Yes](https://aws.amazon.com/about-aws/whats-new/2025/04/amazon-security-lake-internet-protocol-version-6)No[Yes](../../../../general/latest/gr/securitylake.md)[Yes](../../../security-lake/latest/userguide/security-vpc-endpoints.md)AWS Security Token ServiceYesNoYesYesAWS Service CatalogNoNoYesYesAWS ShieldYesYesNoYesAmazon Simple Email ServiceYesNoYesYesAmazon Simple Notification ServiceYesNoYesYesAmazon Simple Queue ServiceYesNoYesYesAmazon Simple Workflow ServiceYesNoYesYesAWS Site-to-Site VPN[Yes](../../../vpn/latest/s2svpn/ipv4-ipv6.md)No[Yes](../../../ec2/latest/devguide/ec2-endpoints.md)NoAWS Snow FamilyNoNoYesAWS Step FunctionsYesNoYesYesAWS Storage GatewayYesYesYesYesAWS Systems ManagerNoNoYesYesAWS Systems Manager Incident ManagerNoNoYesYesAWS Systems Manager for SAPNoNoYesYesAmazon TextractNoNoYesYesAmazon TimestreamNoNoYesYesAmazon TranscribeYesYesYesYesAWS Transfer Family 6[Yes](../../../transfer/latest/userguide/ipv6-support.md)NoYesYesAWS Transit GatewayYesNoYesNoAmazon TranslateYesYesYesYesAWS Trusted AdvisorNoNoYesYesAWS User NotificationsNoNoYesYesAmazon Verified PermissionsYesNoYesYesVMware Cloud on AWSNoNoYesYesAmazon VPC[Yes](how-it-works.md#vpc-ip-addressing)Yes[Yes](../../../ec2/latest/devguide/ec2-endpoints.md)NoAmazon VPC LatticeNoNoYesYesAWS WAF[Yes](https://aws.amazon.com/about-aws/whats-new/2016/10/ipv6-support-for-cloudfront-waf-and-s3-transfer-acceleration)YesNoAWS WAFV2NoNoYesYesAWS Well-Architected ToolNoNoYesYesAmazon WorkMailNoNoYesYesAmazon WorkSpaces[Yes](../../../workspaces/latest/adminguide/amazon-workspaces-vpc.md)NoYesYesAWS X-RayYesNoYesYes

1 An empty cell indicates that the service does not [integrate with\
AWS PrivateLink](../privatelink/aws-services-privatelink-support.md).

2 This entry represents IPv6 support for AWS AppSync
GraphlQL and Event API configuration operations, through the [AWS AppSync SDK API](../../../../reference/appsync/latest/apireference/api-operations.md).
IPv6 is not supported for client connections to customer managed AWS AppSync GraphQL and
Event APIs.

3 This entry represents IPv6 support for Grafana _workspace management_ operations, such as updating
workspaces and workspace permissions. There is no IPv6 support for general Grafana
workspace operations, such as creating and editing dashboards or querying data
sources.

4 This entry represents IPv6 support for AWS Ground Station
_control plane_ operations, such as calling the
[AWS Ground Station API](../../../../reference/ground-station/latest/apireference/api-operations.md).
IPv6 is not supported by the AWS Ground Station _data plane_, so
make sure the resources you are delivering data to (such as Amazon EC2 instances) are
accessible over IPv4.

5 This entry represents IPv6 support for Amazon IVS _control plane_
operations, such as calling an [IVS endpoint](../../../../general/latest/gr/ivs.md).

6 For more details on IPv6 support in AWS Transfer Family, visit [IPv6 limitations](../../../transfer/latest/userguide/ipv6-support.md#ipv6-limitations).

## Additional IPv6 support

###### Compute

- Amazon EC2 supports launching instances based on the Nitro System into IPv6-only
subnets.

- Amazon EC2 provides IPv6 endpoints for Instance Metadata Service (IMDS) and Amazon
Time Sync Service.

###### Game Development

- Amazon GameLift Streams supports streaming over IPv6 on the Microsoft Windows Server 2022 Base runtime.

###### Networking and Content Delivery

- Amazon VPC supports creating IPv6-only subnets.

- Amazon VPC helps IPv6 AWS resources communicate with IPv4 resources by supporting
DNS64 on your subnets and NAT64 on your NAT gateways.

###### Security, Identity, and Compliance

- Amazon Detective supports IPv6 addresses in its network-related findings and entity
profiles.

- AWS Identity and Access Management (IAM) supports IPv6 addresses in IAM identity-based
policies.

- Amazon Macie supports IPv6 addresses in personally identifiable information
(PII).

- Amazon Security Lake supports IPv6 addresses across all operations on log sources and
subscribers.

###### Management and Governance

- AWS CloudTrail records include source IPv6 information.

- AWS CLI v2 supports download over IPv6 connections for IPv6-only clients.

## Learn more

- [IPv6 on\
AWS](../../../whitepapers/latest/ipv6-on-aws/ipv6-on-aws.md)

- [Dual Stack and IPv6-only Amazon VPC Reference Architectures](https://d1.awsstatic.com/architecture-diagrams/ArchitectureDiagrams/IPv6-reference-architectures-for-AWS-and-hybrid-networks-ra.pdf) (PDF)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Example dual-stack VPC

Virtual private clouds

All content copied from https://docs.aws.amazon.com/.
