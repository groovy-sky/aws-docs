AWS Audit Manager will no longer be open to new customers starting
April 30, 2026. If you would like to use Audit Manager, sign up prior to that date. Existing customers
can continue to use the service as normal. For more information, see
[AWS Audit Manager availability change](audit-manager-availability-change.md).

# AWS API calls supported by AWS Audit Manager

You can use Audit Manager to capture snapshots of your AWS environment as evidence for audits.
When you create or edit a custom control, you can specify one or more AWS API calls as a
data source mapping for evidence collection. Audit Manager then makes API calls to the relevant
AWS services, and collects a snapshot of the configuration details for your AWS resources.

For every resource that's in the scope of an API call, Audit Manager captures a configuration
snapshot and converts it into evidence. This results in one piece of evidence per resource, as
opposed to one piece of evidence per API call.

For example, if the `ec2_DescribeRouteTables` API call captures configuration
snapshots from five route tables, then you'll get five pieces of evidence in total for the
single API call. Each piece of evidence is a snapshot of the configuration of an individual
route table.

###### Topics

- [Key points](#control-data-sources-api-key-points)

- [Supported API calls for custom control data sources](#apis-for-custom-control-data-sources)

- [API calls used in the AWS License Manager standard framework](#apis-in-license-manager-framework)

- [Additional resources](#using-api-calls-additional-resources)

## Key points

### Paginated API calls

Many AWS services collect and store a large amount of data. As a result, when a
`list`, `describe`, or `get` API call attempts to
return your data, there can be a lot of results. If the amount of data is too large to
return in a single response, the results can be broken into more manageable pieces through
the use of _pagination_. This divides the results into
"pages" of data, making the responses easier to handle.

Some of the [Supported API calls for custom control data sources](#apis-for-custom-control-data-sources) are paginated. This means that
they return partial results at first, and require subsequent requests to return the entire
result set. For example, the Amazon RDS [DescribeDBInstances](../../../../reference/amazonrds/latest/apireference/api-describedbinstances.md) operation returns up to 100 instances at a time, and
subsequent requests are needed to return the next page of results.

As of March 08, 2023, Audit Manager supports paginated API calls as a data source for evidence
collection. Previously, if a paginated API call was used as a data source, only a subset
of your resources was returned in the API response (up to 100 results). Now, Audit Manager calls
the paginated API operation multiple times, and gets each page of results until all
resources are returned. For each resource, Audit Manager then captures a configuration snapshot and
saves it as evidence. Because your complete set of resources is now captured in the API
response, it’s likely that you’ll notice an increase in the amount of evidence that’s
collected after March 08, 2023.

Audit Manager handles API call pagination for you automatically. If you create a custom control
that uses a paginated API call as a data source, you don’t need to specify any pagination
parameters.

## Supported API calls for custom control data sources

In your custom controls, you can use any of the following API calls as a data source.
Audit Manager can then use these API calls to collect evidence about your AWS usage.

Supported API callHow Audit Manager uses this API to collect evidence[acm\_GetAccountConfiguration](../../../../reference/acm/latest/apireference/api-getaccountconfiguration.md)Collect a snapshot of the account configuration options associated with your
AWS account.[acm\_ListCertificates](../../../../reference/acm/latest/apireference/api-listcertificates.md)Retrieve a list of certificate ARNs and domain names. [autoscaling\_DescribeAutoScalingGroups](../../../../reference/autoscaling/ec2/apireference/api-describeautoscalinggroups.md)Collect a snapshot about the Auto Scaling groups in your AWS account.[backup\_ListBackupPlans](../../../aws-backup/latest/devguide/api-listbackupplans.md)Retrieve a list of all active backup plans in your AWS account. [bedrock\_GetModelInvocationLoggingConfiguration](../../../../reference/bedrock/latest/apireference/api-getmodelinvocationloggingconfiguration.md)Collect a snapshot of the current configuration values for model invocation
logging for models in your AWS account.[cloudfront\_ListDistributions](../../../../reference/cloudfront/latest/apireference/api-listdistributions.md)

Retrieve a list of all distributions in your AWS account.

[cloudtrail\_DescribeTrails](../../../../reference/awscloudtrail/latest/apireference/api-describetrails.md)

Collect a snapshot of the settings for one or more trails associated with the
current Region for your AWS account.[cloudtrail\_ListTrails](../../../../reference/awscloudtrail/latest/apireference/api-listtrails.md)Retrieve a list of the trails that are in your AWS account.

[cloudwatch\_DescribeAlarms](../../../../reference/amazoncloudwatch/latest/apireference/api-describealarms.md)

Collect a configuration snapshot of the alarms that are used for your
AWS account.[config\_DescribeConfigRules](../../../../reference/config/latest/apireference/api-describeconfigrules.md)Retrieve details about your AWS Config rules.[config\_DescribeDeliveryChannels](../../../../reference/config/latest/apireference/api-describedeliverychannels.md)Collect a configuration snapshot for the delivery channels in your in your
AWS account.[directconnect\_DescribeDirectConnectGateways](../../../../reference/directconnect/latest/apireference/api-describedirectconnectgateways.md)Retrieve a list of all your Direct Connect gateways .[directconnect\_DescribeVirtualGateways](../../../../reference/directconnect/latest/apireference/api-describevirtualgateways.md)Retrieve a list of the virtual private gateways owned by your
AWS account.[docdb\_DescribeCertificates](../../../documentdb/latest/developerguide/api-describecertificates.md)Collect a list of certificates for your AWS account.[docdb\_DescribeDBClusterParameterGroups](../../../documentdb/latest/developerguide/api-describedbclusterparametergroups.md)Collect a list of `DBCLusterParameterGroup` descriptions for your
AWS account.[docdb\_DescribeDBInstances](../../../documentdb/latest/developerguide/api-describedbinstances.md)Collect information about provisioned Amazon DynamoDB instances for your
AWS account.

[cloudwatch\_DescribeAlarms](../../../../reference/amazoncloudwatch/latest/apireference/api-describealarms.md)

Collect information about the alarms in your AWS account.

[cloudtrail\_DescribeTrails](../../../../reference/awscloudtrail/latest/apireference/api-describetrails.md)

Collect a snapshot of the settings for one or more trails associated with your
AWS account.

[dynamodb\_DescribeTable](../../../../reference/amazondynamodb/latest/apireference/api-describetable.md)

Collect configuration snapshots for the DynamoDB tables in your
AWS account.

When you use this API as a data source, you don't need to provide the name of
a specific DynamoDB table. Instead, Audit Manager uses the `ListTables` operation
to list all of your tables. For every table that's listed, Audit Manager then performs the
`DescribeTable` operation to generate evidence for that
resource.

[dynamodb\_ListBackups](../../../../reference/amazondynamodb/latest/apireference/api-listbackups.md)Retrieve a list of the DynamoDB backups that are associated with your
AWS account.

[dynamodb\_ListTables](../../../../reference/amazondynamodb/latest/apireference/api-listtables.md)

Retrieve a list of all of the table names that are associated with your
AWS account and your current endpoint.[ec2\_DescribeAddresses](../../../../reference/awsec2/latest/apireference/api-describeaddresses.md)Collect a snapshot of your Elastic IP addresses.[ec2\_DescribeCustomerGateways](../../../../reference/awsec2/latest/apireference/api-describecustomergateways.md)Collect a snapshot of your VPN customer gateways.[ec2\_DescribeEgressOnlyInternetGateways](../../../../reference/awsec2/latest/apireference/api-describeegressonlyinternetgateways.md)Collect a snapshot of your egress-only internet gateways.

[ec2\_DescribeFlowLogs](../../../../reference/awsec2/latest/apireference/api-describeflowlogs.md)

Collect a snapshot of your flow logs.

[ec2\_DescribeInstances](../../../../reference/awsec2/latest/apireference/api-describeinstances.md)

Collect a snapshot of your instances.[ec2\_DescribeInternetGateways](../../../../reference/awsec2/latest/apireference/api-describeinternetgateways.md)Collect a snapshot of your internet gateways.[ec2\_DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations](../../../../reference/awsec2/latest/apireference/api-describelocalgatewayroutetablevirtualinterfacegroupassociations.md)Collect a description of the associations between the virtual interface groups
and the local gateway route tables in your AWS account.[ec2\_DescribeLocalGateways](../../../../reference/awsec2/latest/apireference/api-describelocalgateways.md)Collect a snapshot of your local gateways.[ec2\_DescribeLocalGatewayVirtualInterfaces](../../../../reference/awsec2/latest/apireference/api-describelocalgatewayvirtualinterfaces.md)Collect a snapshot of your local gateway virtual interfaces.[ec2\_DescribeNatGateways](../../../../reference/awsec2/latest/apireference/api-describenatgateways.md)Collect a snapshot of your NAT gateways.

[ec2\_DescribeNetworkAcls](../../../../reference/awsec2/latest/apireference/api-describenetworkacls.md)

Collect a snapshot of your network ACLs.

[ec2\_DescribeRouteTables](../../../../reference/awsec2/latest/apireference/api-describeroutetables.md)

Collect a snapshot of your route tables.

[ec2\_DescribeSecurityGroups](../../../../reference/awsec2/latest/apireference/api-describesecuritygroups.md)

Collect a snapshot of your security groups.[ec2\_DescribeSecurityGroupRules](../../../../reference/awsec2/latest/apireference/api-describesecuritygrouprules.md)Collect a snapshot of one or more of your security group rules.[ec2\_DescribeTransitGateways](../../../../reference/awsec2/latest/apireference/api-describetransitgateways.md)Collect a snapshot of your transit gateways.

[ec2\_DescribeVolumes](../../../../reference/awsec2/latest/apireference/api-describevolumes.md)

Collect a snapshot of your VPC endpoints.

[ec2\_DescribeVpcs](../../../../reference/awsec2/latest/apireference/api-describevpcs.md)

Collect a snapshot of your VPCs.

[ec2\_DescribeVpcEndpoints](../../../../reference/awsec2/latest/apireference/api-describevpcendpoints.md)

Collect a snapshot of your VPC endpoints.[ec2\_DescribeVpcEndpointConnections](../../../../reference/awsec2/latest/apireference/api-describevpcendpointconnections.md)Collect a snapshot of the VPC endpoint connections to your VPC endpoint
services, including any endpoints that are pending your acceptance.[ec2\_DescribeVpcEndpointServiceConfigurations](../../../../reference/awsec2/latest/apireference/api-describevpcendpointserviceconfigurations.md)Collect a snapshot of the VPC endpoint service configurations in your
AWS account.[ec2\_DescribeVpcPeeringConnections](../../../../reference/awsec2/latest/apireference/api-describevpcpeeringconnections.md)Collect a snapshot of your VPN connections.[ec2\_DescribeVpnConnections](../../../../reference/awsec2/latest/apireference/api-describevpnconnections.md)Collect a snapshot of your VPN connections.[ec2\_DescribeVpnGateways](../../../../reference/awsec2/latest/apireference/api-describevpngateways.md)Collect a snapshot of your virtual private gateways.[ec2\_GetEbsDefaultKmsKeyId](../../../../reference/awsec2/latest/apireference/api-getebsdefaultkmskeyid.md)Collect a snapshot of the default AWS KMS key for EBS encryption for your
AWS account in the current Region.[ec2\_GetEbsEncryptionByDefault](../../../../reference/awsec2/latest/apireference/api-getebsencryptionbydefault.md)Describe whether EBS encryption by default is enabled for your AWS account in
the current Region.[ecs\_DescribeClusters](../../../../reference/amazonecs/latest/apireference/api-describeclusters.md)Collect a snapshot of your ECS clusters.[eks\_DescribeAddonVersions](../../../../reference/eks/latest/apireference/api-describeaddonversions.md)Collect a snapshot of your add-on versions.[elasticache\_DescribeCacheClusters](../../../../reference/amazonelasticache/latest/apireference/api-describecacheclusters.md)Collect a snapshot of your provisioned clusters.[elasticache\_DescribeServiceUpdates](../../../../reference/amazonelasticache/latest/apireference/api-describeserviceupdates.md)Collect a snapshot of service updates for Amazon ElastiCache.[elasticfilesystem\_DescribeAccessPoints](../../../efs/latest/ug/api-describeaccesspoints.md)Collect a snapshot of the Amazon EFS access points in your AWS account.

[elasticfilesystem\_DescribeFileSystems](../../../efs/latest/ug/api-describefilesystems.md)

Collect a snapshot of your Amazon EFS file systems.[elasticloadbalancingv2\_DescribeLoadBalancers](../../../../reference/elasticloadbalancing/latest/apireference/api-describeloadbalancers.md)

Collect a snapshot of the load balancers in your AWS account.

[elasticloadbalancingv2\_DescribeSSLPolicies](../../../../reference/elasticloadbalancing/latest/apireference/api-describesslpolicies.md)Collect a snapshot of the policies that you use for SSL negotiation.[elasticloadbalancingv2\_DescribeTargetGroups](../../../../reference/elasticloadbalancing/latest/apireference/api-describetargetgroups.md)Collect a snapshot of your ELB target groups.[elasticmapreduce\_ListSecurityConfigurations](../../../../reference/emr/latest/apireference/api-listsecurityconfigurations.md)Retrieve a list of the security configurations that are visible to your
AWS account, along with their creation dates and times, and their names.[events\_ListConnections](../../../../reference/eventbridge/latest/apireference/api-listconnections.md)Retrieve a list of the Amazon EventBridge connections in your AWS account.[events\_ListEventBuses](../../../../reference/eventbridge/latest/apireference/api-listeventbuses.md)Retrieve a list of the Amazon EventBridge event buses in your AWS account, including
the default event bus, custom event buses, and partner event buses.[events\_ListEventSources](../../../../reference/eventbridge/latest/apireference/api-listeventsources.md)Retrieve a list of the partner event sources that have been shared with your
AWS account. [events\_ListRules](../../../../reference/eventbridge/latest/apireference/api-listrules.md)Retrieve a list of your Amazon EventBridge rules. [firehose\_ListDeliveryStreams](../../../../reference/firehose/latest/apireference/api-listdeliverystreams.md)Retrieve a list of your delivery streams.[fsx\_DescribeFileSystems](../../../efs/latest/ug/api-describefilesystems.md)Collect a snapshot of the file systems that are owned by your
AWS account.[guardduty\_ListDetectors](../../../../reference/guardduty/latest/apireference/api-listdetectors.md)

Retrieve a list of the `detectorIds` for your Amazon GuardDuty detector
resources.

[iam\_GenerateCredentialReport](../../../../reference/iam/latest/apireference/api-generatecredentialreport.md)

Generate a credential report for your AWS account.

[iam\_GetAccountPasswordPolicy](../../../../reference/iam/latest/apireference/api-getaccountpasswordpolicy.md)

Collect a snapshot of the password policy for your AWS account.

[iam\_GetAccountSummary](../../../../reference/iam/latest/apireference/api-getaccountsummary.md)

Collect a snapshot of the IAM entity usage and IAM quotas in your
AWS account.

[iam\_ListGroups](../../../../reference/iam/latest/apireference/api-listgroups.md)

Retrieve a list of the IAM groups that are associated with a path prefix
that's available in your AWS account.[iam\_ListOpenIDConnectProviders](../../../../reference/iam/latest/apireference/api-listopenidconnectproviders.md)Retrieve a list of the IAM OpenID Connect (OIDC) provider resource objects
that are defined in your AWS account.

[iam\_ListPolicies](../../../../reference/iam/latest/apireference/api-listpolicies.md)

Retrieve a list of all the managed policies that are available in your
AWS account, including your own customer-defined managed policies and all AWS
managed policies.

[iam\_ListRoles](../../../../reference/iam/latest/apireference/api-listroles.md)

Retrieve a list of the IAM roles that are associated with a path prefix
that's available in your AWS account.[iam\_ListSAMLProviders](../../../../reference/iam/latest/apireference/api-listsamlproviders.md)Retrieve a list of the SAML provider resource objects defined in IAM in your
AWS account.

[iam\_ListUsers](../../../../reference/iam/latest/apireference/api-listusers.md)

Retrieve a list of the IAM users in your AWS account.[iam\_ListVirtualMFADevices](../../../../reference/iam/latest/apireference/api-listvirtualmfadevices.md)Retrieve a list of the virtual MFA devices that are defined in your
AWS account.[kafka\_ListClusters](../../../../reference/msk/1-0/apireference/clusters.md#clustersget)Retrieve a list of the Amazon MSK clusters in your AWS account.[kafka\_ListKafkaVersions](../../../../reference/msk/1-0/apireference/kafka-versions.md#kafka-versionsget)Retrieve a list of the Apache Kafka version objects in your
AWS account.[kinesis\_ListStreams](../../../../reference/kinesis/latest/apireference/api-liststreams.md)Retrieve a list of your Kinesis data streams.

[kms\_GetKeyPolicy](../../../../reference/kms/latest/apireference/api-getkeypolicy.md)

Audit Manager uses this API to collect a snapshot of the key policies for the
AWS KMS keys in your AWS account.

When you use this API as a data source, you don't need to provide the name of
a specific AWS KMS key. Instead, Audit Manager uses the `ListKeys` operation
to list all of your KMS keys. For every KMS key that's listed, Audit Manager then
performs the `GetKeyPolicy` operation to generate evidence for that
resource.

[kms\_GetKeyRotationStatus](../../../../reference/kms/latest/apireference/api-getkeyrotationstatus.md)

Audit Manager uses this API to collect a snapshot of whether automatic rotation is
enabled for the AWS KMS keys in your AWS account.

When you use this API as a data source, you don't need to provide the name of
a specific AWS KMS key. Instead, Audit Manager uses the `ListKeys` operation
to list all of your KMS keys. For every KMS key that's listed, Audit Manager then
performs the `GetKeyRotationStatus` operation to generate evidence for
that resource.

[kms\_ListKeys](../../../../reference/kms/latest/apireference/api-listkeys.md)Retrieve a list of the AWS KMS keys in your AWS account.[lambda\_ListFunctions](../../../lambda/latest/dg/api-listfunctions.md)Retrieve a list of Lambda functions in your AWS account, with the
version-specific configuration of each. [rds\_DescribeDBClusters](../../../../reference/amazonrds/latest/apireference/api-describedbclusters.md)Collect a snapshot of the existing Amazon Aurora DB clusters and Multi-AZ DB
clusters in your AWS account.

[rds\_DescribeDBInstances](../../../../reference/amazonrds/latest/apireference/api-describedbinstances.md)

Collect a snapshot of the provisioned RDS instances in your
AWS account.[rds\_DescribeDbInstanceAutomatedBackups](../../../../reference/amazonrds/latest/apireference/api-describedbinstanceautomatedbackups.md)Collect a snapshot of the backups for both current and deleted instances in
your AWS account.[rds\_DescribeDbSecurityGroups](../../../../reference/amazonrds/latest/apireference/api-describedbsecuritygroups.md)Collect a snapshot of the DBSecurityGroups in your AWS account.

[redshift\_DescribeClusters](../../../../reference/redshift/latest/apireference/api-describeclusters.md)

Collect a snapshot of the provisioned Amazon Redshift clusters in your
AWS account.

[s3\_GetBucketEncryption](../../../s3/latest/api/api-getbucketencryption.md)

Collect a snapshot that shows the default encryption configuration for your S3
buckets.

When you use this API as a data source, you don't need to provide the name of
a specific S3 bucket. Instead, Audit Manager uses the `ListBuckets` operation to
list the buckets that were created in the same AWS Region as your assessment.
For every bucket that's listed, Audit Manager then performs the
`GetBucketEncryption` operation to generate evidence for that
resource.

Audit Manager can only provide the encryption status for buckets that were created in
the same AWS Region as your assessment. If you need to see the encryption status
of all your S3 buckets across multiple AWS Regions, we recommend that you create
an assessment in each AWS Region where you have an S3 bucket.

[s3\_ListBuckets](../../../s3/latest/api/api-listbuckets.md)

Retrieve a list of the S3 buckets in your AWS account. Audit Manager can only
list buckets that were created in the same AWS Region as your assessment. If you
need to see all your S3 buckets across multiple AWS Regions, we recommend that you
create an assessment in each AWS Region where you have an S3 bucket.[sagemaker\_ListAlgorithms](../../../../reference/sagemaker/latest/apireference/api-listalgorithms.md)Retrieve a list of the machine learning algorithms in your
AWS account.[sagemaker\_ListDomains](../../../../reference/sagemaker/latest/apireference/api-listdomains.md)Retrieve a list of the domains in your AWS account.[sagemaker\_ListEndpoints](../../../../reference/sagemaker/latest/apireference/api-listendpoints.md)Retrieve a list of the endpoints in your AWS account.[sagemaker\_ListEndpointConfigs](../../../../reference/sagemaker/latest/apireference/api-listendpointconfigs.md)Retrieve a list of the endpoint configurations in your AWS account.[sagemaker\_ListFlowDefinitions](../../../../reference/sagemaker/latest/apireference/api-listflowdefinitions.md)Retrieve a list of the flow definitions in your AWS account.[sagemaker\_ListHumanTaskUis](../../../../reference/sagemaker/latest/apireference/api-listhumantaskuis.md)Retrieve a list of the human task interfaces in your AWS account.[sagemaker\_ListLabelingJobs](../../../../reference/sagemaker/latest/apireference/api-listlabelingjobs.md)Retrieve a list of the labeling jobs in your AWS account.[sagemaker\_ListModels](../../../../reference/sagemaker/latest/apireference/api-listmodels.md)Retrieve a list of the models in your AWS account.[sagemaker\_ListModelBiasJobDefinitions](../../../../reference/sagemaker/latest/apireference/api-listmodelbiasjobdefinitions.md)Retrieve a list of the model bias job definitions in your
AWS account.[sagemaker\_ListModelCards](../../../../reference/sagemaker/latest/apireference/api-listmodelcards.md)Retrieve a list of the model cards in your AWS account.[sagemaker\_ListModelQualityJobDefinitions](../../../../reference/sagemaker/latest/apireference/api-listmodelqualityjobdefinitions.md)Retrieve a list of the model quality monitoring job definitions in your
AWS account.[sagemaker\_ListMonitoringAlerts](../../../../reference/sagemaker/latest/apireference/api-listmonitoringalerts.md)Retrieve a list of the alerts for a given monitoring schedule.[sagemaker\_ListMonitoringSchedules](../../../../reference/sagemaker/latest/apireference/api-listmonitoringschedules.md)Retrieve a list of all monitoring schedules in your AWS account.[sagemaker\_ListTrainingJobs](../../../../reference/sagemaker/latest/apireference/api-listtrainingjobs.md)Retrieve a list of training jobs in your AWS account.[sagemaker\_ListUserProfiles](../../../../reference/sagemaker/latest/apireference/api-listuserprofiles.md)Retrieve a list of user profiles in your AWS account.[secretsmanager\_ListSecrets](../../../../reference/secretsmanager/latest/apireference/api-listsecrets.md)Retrieve a list of the secrets that are stored in your AWS account, not
including secrets that are marked for deletion. [sns\_ListTopics](../../../sns/latest/api/api-listtopics.md)Retrieve a list of the SNS topics in your AWS account.[sqs\_ListQueues](../../../../reference/awssimplequeueservice/latest/apireference/api-listqueues.md)Retrieve a list of the SQS queues in your AWS account.[waf-regional\_ListWebAcls](../../../../reference/waf/latest/apireference/api-wafregional-listwebacls.md)Retrieve a list of the [WebACLSummary](../../../../reference/waf/latest/apireference/api-wafregional-webaclsummary.md) objects for your AWS account.[waf-regional\_ListRules](../../../../reference/waf/latest/apireference/api-wafregional-listrules.md)Retrieve a list of the [RuleSummary](../../../../reference/waf/latest/apireference/api-wafregional-rulesummary.md) objects for your AWS account.[waf\_ListRuleGroups](../../../../reference/waf/latest/apireference/api-listwebacls.md)Retrieve a list of the [RuleGroupSummary](../../../../reference/waf/latest/apireference/api-rulegroupsummary.md) objects for the rule groups in your
AWS account.[waf\_ListRules](../../../../reference/waf/latest/apireference/api-waf-listrules.md)Retrieve a list of the [RuleSummary](../../../../reference/waf/latest/apireference/api-waf-rulesummary.md) objects
for your AWS account.[waf\_ListWebAcls](../../../../reference/waf/latest/apireference/api-waf-listwebacls.md)Retrieve a list of the [WebACLSummary](../../../../reference/waf/latest/apireference/api-waf-webaclsummary.md)
objects for your AWS account.

## API calls used in the AWS License Manager standard framework

In the [AWS License Manager](licensemanager.md) standard
framework, Audit Manager uses a custom activity called `GetLicenseManagerSummary` to
collect evidence. This activity calls the following three License Manager APIs:

- [ListLicenseConfigurations](../../../../reference/license-manager/latest/apireference/api-listlicenseconfigurations.md)

- [ListAssociationsForLicenseConfiguration](../../../../reference/license-manager/latest/apireference/api-listassociationsforlicenseconfiguration.md)

- [ListUsageForLicenseConfiguration](../../../../reference/license-manager/latest/apireference/api-listusageforlicenseconfiguration.md)

The data that’s returned is then converted into evidence and attached to the relevant
controls in your assessment.

###### Example

Let's say that you use two licensed products ( _SQL Service_
_2017_ and _Oracle Database Enterprise_
_Edition_). First, the `GetLicenseManagerSummary` activity calls the
[ListLicenseConfigurations](../../../../reference/license-manager/latest/apireference/api-listlicenseconfigurations.md) API, which provides details of license configurations
in your account. Next, it adds additional contextual data for each license configuration
by calling [ListUsageForLicenseConfiguration](../../../../reference/license-manager/latest/apireference/api-listusageforlicenseconfiguration.md) and [ListAssociationsForLicenseConfiguration](../../../../reference/license-manager/latest/apireference/api-listassociationsforlicenseconfiguration.md). Finally, it converts the license
configuration data into evidence and attaches it to the respective controls in the
framework ( _4.5 - Customer managed license for SQL Server_
_2017_ and _3.0.4 - Customer managed license for Oracle_
_Database Enterprise Edition_).

If you’re using a licensed product that isn’t covered by any of the controls in the
framework, that license configuration data is attached as evidence to the following control:
_5.0 - Customer managed license for other_
_licenses_.

## Additional resources

- To find help with evidence collection issues for this data source type, see [My assessment isn’t collecting configuration data evidence for an AWS API call](evidence-collection-issues.md#no-evidence-from-aws-api-calls).

- To create a custom control using this data source type, see [Creating a custom control in AWS Audit Manager](create-controls.md).

- To create a custom framework that uses your custom control, see [Creating a custom framework in AWS Audit Manager](custom-frameworks.md).

- To add your custom control to an existing custom framework, see [Editing a custom framework in AWS Audit Manager](edit-custom-frameworks.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Security Hub CSPM

AWS CloudTrail

All content copied from https://docs.aws.amazon.com/.
