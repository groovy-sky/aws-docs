AWS Audit Manager will no longer be open to new customers starting
April 30, 2026. If you would like to use Audit Manager, sign up prior to that date. Existing customers
can continue to use the service as normal. For more information, see
[AWS Audit Manager availability change](audit-manager-availability-change.md).

# AWS Config Rules supported by AWS Audit Manager

You can use Audit Manager to capture AWS Config evaluations as evidence for audits. When you create
or edit a custom control, you can specify one or more AWS Config rules as a data source mapping
for evidence collection. AWS Config performs compliance checks based on these rules, and Audit Manager
reports the results as compliance check evidence.

In addition to managed rules, you can also map your custom rules to a control data source.

###### Contents

- [Key points](control-data-sources-config.md#aws-config-key-points)

- [Supported AWS Config managed rules](control-data-sources-config.md#aws-config-managed-rules)

- [Using AWS Config custom rules with Audit Manager](control-data-sources-config.md#aws-config-custom-rules)

- [Additional resources](control-data-sources-config.md#aws-config-rules-troubleshoot)

## Key points

- Audit Manager doesn’t collect evidence from [service-linked\
AWS Config rules](../../../config/latest/developerguide/service-linked-awsconfig-rules.md), with the exception of service-linked rules from Conformance
Packs and from AWS Organizations.

- Audit Manager doesn't manage AWS Config rules for you. Before you start evidence collection, we
recommend that you review your current AWS Config rule parameters. Then, validate those
parameters against the requirements of your chosen framework. If needed, you can [update a rule's\
parameters in AWS Config](../../../config/latest/developerguide/evaluate-config-manage-rules.md) so that it aligns with framework requirements. This will
help to ensure that your assessments collect the correct compliance check evidence for
that framework.

For example, suppose that you’re creating an assessment for CIS v1.2.0. This
framework has a control named [Ensure IAM password policy\
requires a minimum length of 14 or greater](../../../securityhub/latest/userguide/iam-controls.md#iam-15). In AWS Config, the [iam-password-policy](../../../config/latest/developerguide/iam-password-policy.md) rule has a `MinimumPasswordLength` parameter
that checks password length. The default value for this parameter is 14 characters. As a
result, the rule aligns with the control requirements. If you aren’t using the default
parameter value, ensure that the value you’re using is equal to or greater than the 14
character requirement from CIS v1.2.0. You can find the default parameter details for
each managed rule in the [AWS Config\
documentation](../../../config/latest/developerguide/managed-rules-by-aws-config.md).

- If you need to verify if an AWS Config rule is a managed rule or a custom rule, you
can do this using the [AWS Config console](https://console.aws.amazon.com/config). From
the left navigation menu, choose **Rules** and look for the
rule in the table. If it's a managed rule, the **Type** column shows
**AWS managed**.

![A managed rule as shown in the AWS Config console.](https://docs.aws.amazon.com/images/audit-manager/latest/userguide/images/rules-managed-console.png)

## Supported AWS Config managed rules

The following AWS Config managed rules are supported by Audit Manager. You can use any of the
following managed rule identifier keywords when you set up a data source for a custom
control. For more information about any of the managed rules listed below, choose an item
from the list or see [AWS Config Managed\
Rules](../../../config/latest/developerguide/evaluate-config-use-managed-rules.md) in the _AWS Config User Guide_.

###### Tip

When you choose a managed rule in the Audit Manager console during custom control creation,
make sure that you look for one of the following rule identifier keywords, and not the
rule name. For information about the difference between the rule name and rule identifier,
and how to find the identifier for a managed rule, see the [Troubleshooting](control-issues.md#managed-rule-missing) section of this user guide.

Supported AWS Config managed rule keywords

- [ACCESS\_KEYS\_ROTATED](../../../config/latest/developerguide/access-keys-rotated.md)

- [ACCOUNT\_PART\_OF\_ORGANIZATIONS](../../../config/latest/developerguide/account-part-of-organizations.md)

- [ACM\_CERTIFICATE\_EXPIRATION\_CHECK](../../../config/latest/developerguide/acm-certificate-expiration-check.md)

- [ACM\_CERTIFICATE\_RSA\_CHECK](../../../config/latest/developerguide/acm-certificate-rsa-check.md)

- [ALB\_DESYNC\_MODE\_CHECK](../../../config/latest/developerguide/alb-desync-mode-check.md)

- [ALB\_HTTP\_DROP\_INVALID\_HEADER\_ENABLED](../../../config/latest/developerguide/alb-http-drop-invalid-header-enabled.md)

- [ALB\_HTTP\_TO\_HTTPS\_REDIRECTION\_CHECK](../../../config/latest/developerguide/alb-http-to-https-redirection-check.md)

- [ALB\_WAF\_ENABLED](../../../config/latest/developerguide/alb-waf-enabled.md)

- [API\_GW\_ASSOCIATED\_WITH\_WAF](../../../config/latest/developerguide/api-gw-associated-with-waf.md)

- [API\_GW\_CACHE\_ENABLED\_AND\_ENCRYPTED](../../../config/latest/developerguide/api-gw-cache-enabled-and-encrypted.md)

- [API\_GW\_ENDPOINT\_TYPE\_CHECK](../../../config/latest/developerguide/api-gw-endpoint-type-check.md)

- [API\_GW\_EXECUTION\_LOGGING\_ENABLED](../../../config/latest/developerguide/api-gw-execution-logging-enabled.md)

- [API\_GW\_SSL\_ENABLED](../../../config/latest/developerguide/api-gw-ssl-enabled.md)

- [API\_GW\_XRAY\_ENABLED](../../../config/latest/developerguide/api-gw-xray-enabled.md)

- [API\_GWV2\_ACCESS\_LOGS\_ENABLED](../../../config/latest/developerguide/api-gwv2-access-logs-enabled.md)

- [API\_GWV2\_AUTHORIZATION\_TYPE\_CONFIGURED](../../../config/latest/developerguide/api-gwv2-authorization-type-configured.md)

- [APPROVED\_AMIS\_BY\_ID](../../../config/latest/developerguide/approved-amis-by-id.md)

- [APPROVED\_AMIS\_BY\_TAG](../../../config/latest/developerguide/approved-amis-by-tag.md)

- [APPSYNC\_ASSOCIATED\_WITH\_WAF](../../../config/latest/developerguide/appsync-associated-with-waf.md)

- [APPSYNC\_CACHE\_ENCRYPTION\_AT\_REST](../../../config/latest/developerguide/appsync-cache-encryption-at-rest.md)

- [APPSYNC\_LOGGING\_ENABLED](../../../config/latest/developerguide/appsync-logging-enabled.md)

- [AURORA\_LAST\_BACKUP\_RECOVERY\_POINT\_CREATED](../../../config/latest/developerguide/aurora-last-backup-recovery-point-created.md)

- [AURORA\_MYSQL\_BACKTRACKING\_ENABLED](../../../config/latest/developerguide/aurora-mysql-backtracking-enabled.md)

- [AURORA\_RESOURCES\_PROTECTED\_BY\_BACKUP\_PLAN](../../../config/latest/developerguide/aurora-resources-protected-by-backup-plan.md)

- [AUTOSCALING\_CAPACITY\_REBALANCING](../../../config/latest/developerguide/autoscaling-capacity-rebalancing.md)

- [AUTOSCALING\_GROUP\_ELB\_HEALTHCHECK\_REQUIRED](../../../config/latest/developerguide/autoscaling-group-elb-healthcheck-required.md)

- [AUTOSCALING\_LAUNCH\_CONFIG\_HOP\_LIMIT](../../../config/latest/developerguide/autoscaling-launch-config-hop-limit.md)

- [AUTOSCALING\_LAUNCH\_CONFIG\_PUBLIC\_IP\_DISABLED](../../../config/latest/developerguide/autoscaling-launch-config-public-ip-disabled.md)

- [AUTOSCALING\_LAUNCHCONFIG\_REQUIRES\_IMDSV2](../../../config/latest/developerguide/autoscaling-launchconfig-requires-imdsv2.md)

- [AUTOSCALING\_LAUNCH\_TEMPLATE](../../../config/latest/developerguide/autoscaling-launch-template.md)

- [AUTOSCALING\_MULTIPLE\_AZ](../../../config/latest/developerguide/autoscaling-multiple-az.md)

- [AUTOSCALING\_MULTIPLE\_INSTANCE\_TYPES](../../../config/latest/developerguide/autoscaling-multiple-instance-types.md)

- [BACKUP\_PLAN\_MIN\_FREQUENCY\_AND\_MIN\_RETENTION\_CHECK](../../../config/latest/developerguide/backup-plan-min-frequency-and-min-retention-check.md)

- [BACKUP\_RECOVERY\_POINT\_ENCRYPTED](../../../config/latest/developerguide/backup-recovery-point-encrypted.md)

- [BACKUP\_RECOVERY\_POINT\_MANUAL\_DELETION\_DISABLED](../../../config/latest/developerguide/backup-recovery-point-manual-deletion-disabled.md)

- [BACKUP\_RECOVERY\_POINT\_MINIMUM\_RETENTION\_CHECK](../../../config/latest/developerguide/backup-recovery-point-minimum-retention-check.md)

- [BEANSTALK\_ENHANCED\_HEALTH\_REPORTING\_ENABLED](../../../config/latest/developerguide/beanstalk-enhanced-health-reporting-enabled.md)

- [CLB\_DESYNC\_MODE\_CHECK](../../../config/latest/developerguide/clb-desync-mode-check.md)

- [CLB\_MULTIPLE\_AZ](../../../config/latest/developerguide/clb-multiple-az.md)

- [CLOUD\_TRAIL\_CLOUD\_WATCH\_LOGS\_ENABLED](../../../config/latest/developerguide/cloud-trail-cloud-watch-logs-enabled.md)

- [CLOUD\_TRAIL\_ENABLED](../../../config/latest/developerguide/cloudtrail-enabled.md)

- [CLOUD\_TRAIL\_ENCRYPTION\_ENABLED](../../../config/latest/developerguide/cloud-trail-encryption-enabled.md)

- [CLOUD\_TRAIL\_LOG\_FILE\_VALIDATION\_ENABLED](../../../config/latest/developerguide/cloud-trail-log-file-validation-enabled.md)

- [CLOUDFORMATION\_STACK\_DRIFT\_DETECTION\_CHECK](../../../config/latest/developerguide/cloudformation-stack-drift-detection-check.md)

- [CLOUDFORMATION\_STACK\_NOTIFICATION\_CHECK](../../../config/latest/developerguide/cloudformation-stack-notification-check.md)

- [CLOUDFRONT\_ACCESSLOGS\_ENABLED](../../../config/latest/developerguide/cloudfront-accesslogs-enabled.md)

- [CLOUDFRONT\_ASSOCIATED\_WITH\_WAF](../../../config/latest/developerguide/cloudfront-associated-with-waf.md)

- [CLOUDFRONT\_CUSTOM\_SSL\_CERTIFICATE](../../../config/latest/developerguide/cloudfront-custom-ssl-certificate.md)

- [CLOUDFRONT\_DEFAULT\_ROOT\_OBJECT\_CONFIGURED](../../../config/latest/developerguide/cloudfront-default-root-object-configured.md)

- [CLOUDFRONT\_NO\_DEPRECATED\_SSL\_PROTOCOLS](../../../config/latest/developerguide/cloudfront-no-deprecated-ssl-protocols.md)

- [CLOUDFRONT\_ORIGIN\_ACCESS\_IDENTITY\_ENABLED](../../../config/latest/developerguide/cloudfront-origin-access-identity-enabled.md)

- [CLOUDFRONT\_ORIGIN\_FAILOVER\_ENABLED](../../../config/latest/developerguide/cloudfront-origin-failover-enabled.md)

- [CLOUDFRONT\_S3\_ORIGIN\_ACCESS\_CONTROL\_ENABLED](../../../config/latest/developerguide/cloudfront-s3-origin-access-control-enabled.md)

- [CLOUDFRONT\_S3\_ORIGIN\_NON\_EXISTENT\_BUCKET](../../../config/latest/developerguide/cloudfront-s3-origin-non-existent-bucket.md)

- [CLOUDFRONT\_SECURITY\_POLICY\_CHECK](../../../config/latest/developerguide/cloudfront-security-policy-check.md)

- [CLOUDFRONT\_SNI\_ENABLED](../../../config/latest/developerguide/cloudfront-sni-enabled.md)

- [CLOUDFRONT\_TRAFFIC\_TO\_ORIGIN\_ENCRYPTED](../../../config/latest/developerguide/cloudfront-traffic-to-origin-encrypted.md)

- [CLOUDFRONT\_VIEWER\_POLICY\_HTTPS](../../../config/latest/developerguide/cloudfront-viewer-policy-https.md)

- [CLOUDTRAIL\_S3\_DATAEVENTS\_ENABLED](../../../config/latest/developerguide/cloudtrail-s3-dataevents-enabled.md)

- [CLOUDTRAIL\_SECURITY\_TRAIL\_ENABLED](../../../config/latest/developerguide/cloudtrail-security-trail-enabled.md)

- [CLOUDWATCH\_ALARM\_ACTION\_CHECK](../../../config/latest/developerguide/cloudwatch-alarm-action-check.md)

- [CLOUDWATCH\_ALARM\_ACTION\_ENABLED\_CHECK](../../../config/latest/developerguide/cloudwatch-alarm-action-enabled-check.md)

- [CLOUDWATCH\_ALARM\_RESOURCE\_CHECK](../../../config/latest/developerguide/cloudwatch-alarm-resource-check.md)

- [CLOUDWATCH\_ALARM\_SETTINGS\_CHECK](../../../config/latest/developerguide/cloudwatch-alarm-settings-check.md)

- [CLOUDWATCH\_LOG\_GROUP\_ENCRYPTED](../../../config/latest/developerguide/cloudwatch-log-group-encrypted.md)

- [CMK\_BACKING\_KEY\_ROTATION\_ENABLED](../../../config/latest/developerguide/cmk-backing-key-rotation-enabled.md)

- [CODEBUILD\_PROJECT\_ARTIFACT\_ENCRYPTION](../../../config/latest/developerguide/codebuild-project-artifact-encryption.md)

- [CODEBUILD\_PROJECT\_ENVIRONMENT\_PRIVILEGED\_CHECK](../../../config/latest/developerguide/codebuild-project-environment-privileged-check.md)

- [CODEBUILD\_PROJECT\_ENVVAR\_AWSCRED\_CHECK](../../../config/latest/developerguide/codebuild-project-envvar-awscred-check.md)

- [CODEBUILD\_PROJECT\_LOGGING\_ENABLED](../../../config/latest/developerguide/codebuild-project-logging-enabled.md)

- [CODEBUILD\_PROJECT\_S3\_LOGS\_ENCRYPTED](../../../config/latest/developerguide/codebuild-project-s3-logs-encrypted.md)

- [CODEBUILD\_PROJECT\_SOURCE\_REPO\_URL\_CHECK](../../../config/latest/developerguide/codebuild-project-source-repo-url-check.md)

- [CODEDEPLOY\_AUTO\_ROLLBACK\_MONITOR\_ENABLED](../../../config/latest/developerguide/codedeploy-auto-rollback-monitor-enabled.md)

- [CODEDEPLOY\_EC2\_MINIMUM\_HEALTHY\_HOSTS\_CONFIGURED](../../../config/latest/developerguide/codedeploy-ec2-minimum-healthy-hosts-configured.md)

- [CODEDEPLOY\_LAMBDA\_ALLATONCE\_TRAFFIC\_SHIFT\_DISABLED](../../../config/latest/developerguide/codedeploy-lambda-allatonce-traffic-shift-disabled.md)

- [CODEPIPELINE\_DEPLOYMENT\_COUNT\_CHECK](../../../config/latest/developerguide/codepipeline-deployment-count-check.md)

- [CODEPIPELINE\_REGION\_FANOUT\_CHECK](../../../config/latest/developerguide/codepipeline-region-fanout-check.md)

- [CUSTOM\_SCHEMA\_REGISTRY\_POLICY\_ATTACHED](../../../config/latest/developerguide/custom-schema-registry-policy-attached.md)

- [CW\_LOGGROUP\_RETENTION\_PERIOD\_CHECK](../../../config/latest/developerguide/cw-loggroup-retention-period-check.md)

- [DAX\_ENCRYPTION\_ENABLED](../../../config/latest/developerguide/dax-encryption-enabled.md)

- [DB\_INSTANCE\_BACKUP\_ENABLED](../../../config/latest/developerguide/db-instance-backup-enabled.md)

- [DESIRED\_INSTANCE\_TENANCY](../../../config/latest/developerguide/desired-instance-tenancy.md)

- [DESIRED\_INSTANCE\_TYPE](../../../config/latest/developerguide/desired-instance-type.md)

- [DMS\_REPLICATION\_NOT\_PUBLIC](../../../config/latest/developerguide/dms-replication-not-public.md)

- [DYNAMODB\_AUTOSCALING\_ENABLED](../../../config/latest/developerguide/dynamodb-autoscaling-enabled.md)

- [DYNAMODB\_IN\_BACKUP\_PLAN](../../../config/latest/developerguide/dynamodb-in-backup-plan.md)

- [DYNAMODB\_LAST\_BACKUP\_RECOVERY\_POINT\_CREATED](../../../config/latest/developerguide/dynamodb-last-backup-recovery-point-created.md)

- [DYNAMODB\_PITR\_ENABLED](../../../config/latest/developerguide/dynamodb-pitr-enabled.md)

- [DYNAMODB\_RESOURCES\_PROTECTED\_BY\_BACKUP\_PLAN](../../../config/latest/developerguide/dynamodb-resources-protected-by-backup-plan.md)

- [DYNAMODB\_TABLE\_ENCRYPTED\_KMS](../../../config/latest/developerguide/dynamodb-table-encrypted-kms.md)

- [DYNAMODB\_TABLE\_ENCRYPTION\_ENABLED](../../../config/latest/developerguide/dynamodb-table-encryption-enabled.md)

- [DYNAMODB\_THROUGHPUT\_LIMIT\_CHECK](../../../config/latest/developerguide/dynamodb-throughput-limit-check.md)

- [EBS\_IN\_BACKUP\_PLAN](../../../config/latest/developerguide/ebs-in-backup-plan.md)

- [EBS\_LAST\_BACKUP\_RECOVERY\_POINT\_CREATED](../../../config/latest/developerguide/ebs-last-backup-recovery-point-created.md)

- [EBS\_OPTIMIZED\_INSTANCE](../../../config/latest/developerguide/ebs-optimized-instance.md)

- [EBS\_RESOURCES\_PROTECTED\_BY\_BACKUP\_PLAN](../../../config/latest/developerguide/ebs-resources-protected-by-backup-plan.md)

- [EBS\_SNAPSHOT\_PUBLIC\_RESTORABLE\_CHECK](../../../config/latest/developerguide/ebs-snapshot-public-restorable-check.md)

- [EC2\_CLIENT\_VPN\_NOT\_AUTHORIZE\_ALL](../../../config/latest/developerguide/ec2-client-vpn-not-authorize-all.md)

- [EC2\_EBS\_ENCRYPTION\_BY\_DEFAULT](../../../config/latest/developerguide/ec2-ebs-encryption-by-default.md)

- [EC2\_IMDSV2\_CHECK](../../../config/latest/developerguide/ec2-imdsv2-check.md)

- [EC2\_INSTANCE\_DETAILED\_MONITORING\_ENABLED](../../../config/latest/developerguide/ec2-instance-detailed-monitoring-enabled.md)

- [EC2\_INSTANCE\_MANAGED\_BY\_SSM](../../../config/latest/developerguide/ec2-instance-managed-by-systems-manager.md)

- [EC2\_INSTANCE\_MULTIPLE\_ENI\_CHECK](../../../config/latest/developerguide/ec2-instance-multiple-eni-check.md)

- [EC2\_INSTANCE\_NO\_PUBLIC\_IP](../../../config/latest/developerguide/ec2-instance-no-public-ip.md)

- [EC2\_INSTANCE\_PROFILE\_ATTACHED](../../../config/latest/developerguide/ec2-instance-profile-attached.md)

- [EC2\_LAST\_BACKUP\_RECOVERY\_POINT\_CREATED](../../../config/latest/developerguide/ec2-last-backup-recovery-point-created.md)

- [EC2\_LAUNCH\_TEMPLATE\_PUBLIC\_IP\_DISABLED](../../../config/latest/developerguide/ec2-launch-template-public-ip-disabled.md)

- [EC2\_MANAGEDINSTANCE\_APPLICATIONS\_BLACKLISTED](../../../config/latest/developerguide/ec2-managedinstance-applications-blacklisted.md)

- [EC2\_MANAGEDINSTANCE\_APPLICATIONS\_REQUIRED](../../../config/latest/developerguide/ec2-managedinstance-applications-required.md)

- [EC2\_MANAGEDINSTANCE\_ASSOCIATION\_COMPLIANCE\_STATUS\_CHECK](../../../config/latest/developerguide/ec2-managedinstance-association-compliance-status-check.md)

- [EC2\_MANAGEDINSTANCE\_INVENTORY\_BLACKLISTED](../../../config/latest/developerguide/ec2-managedinstance-inventory-blacklisted.md)

- [EC2\_MANAGEDINSTANCE\_PATCH\_COMPLIANCE\_STATUS\_CHECK](../../../config/latest/developerguide/ec2-managedinstance-patch-compliance-status-check.md)

- [EC2\_MANAGEDINSTANCE\_PLATFORM\_CHECK](../../../config/latest/developerguide/ec2-managedinstance-platform-check.md)

- [EC2\_NO\_AMAZON\_KEY\_PAIR](../../../config/latest/developerguide/ec2-no-amazon-key-pair.md)

- [EC2\_PARAVIRTUAL\_INSTANCE\_CHECK](../../../config/latest/developerguide/ec2-paravirtual-instance-check.md)

- [EC2\_RESOURCES\_PROTECTED\_BY\_BACKUP\_PLAN](../../../config/latest/developerguide/ec2-resources-protected-by-backup-plan.md)

- [EC2\_SECURITY\_GROUP\_ATTACHED\_TO\_ENI](../../../config/latest/developerguide/ec2-security-group-attached-to-eni.md)

- [EC2\_SECURITY\_GROUP\_ATTACHED\_TO\_ENI\_PERIODIC](../../../config/latest/developerguide/ec2-security-group-attached-to-eni-periodic.md)

- [EC2\_STOPPED\_INSTANCE](../../../config/latest/developerguide/ec2-stopped-instance.md)

- [EC2\_TOKEN\_HOP\_LIMIT\_CHECK](../../../config/latest/developerguide/ec2-token-hop-limit-check.md)

- [EC2\_TRANSIT\_GATEWAY\_AUTO\_VPC\_ATTACH\_DISABLED](../../../config/latest/developerguide/ec2-transit-gateway-auto-vpc-attach-disabled.md)

- [EC2\_VOLUME\_INUSE\_CHECK](../../../config/latest/developerguide/ec2-volume-inuse-check.md)

- [ECR\_PRIVATE\_IMAGE\_SCANNING\_ENABLED](../../../config/latest/developerguide/ecr-private-image-scanning-enabled.md)

- [ECR\_PRIVATE\_LIFECYCLE\_POLICY\_CONFIGURED](../../../config/latest/developerguide/ecr-private-lifecycle-policy-configured.md)

- [ECR\_PRIVATE\_TAG\_IMMUTABILITY\_ENABLED](../../../config/latest/developerguide/ecr-private-tag-immutability-enabled.md)

- [ECS\_AWSVPC\_NETWORKING\_ENABLED](../../../config/latest/developerguide/ecs-awsvpc-networking-enabled.md)

- [ECS\_CONTAINER\_INSIGHTS\_ENABLED](../../../config/latest/developerguide/ecs-container-insights-enabled.md)

- [ECS\_CONTAINERS\_NONPRIVILEGED](../../../config/latest/developerguide/ecs-containers-nonprivileged.md)

- [ECS\_CONTAINERS\_READONLY\_ACCESS](../../../config/latest/developerguide/ecs-containers-readonly-access.md)

- [ECS\_FARGATE\_LATEST\_PLATFORM\_VERSION](../../../config/latest/developerguide/ecs-fargate-latest-platform-version.md)

- [ECS\_NO\_ENVIRONMENT\_SECRETS](../../../config/latest/developerguide/ecs-no-environment-secrets.md)

- [ECS\_TASK\_DEFINITION\_LOG\_CONFIGURATION](../../../config/latest/developerguide/ecs-task-definition-log-configuration.md)

- [ECS\_TASK\_DEFINITION\_MEMORY\_HARD\_LIMIT](../../../config/latest/developerguide/ecs-task-definition-memory-hard-limit.md)

- [ECS\_TASK\_DEFINITION\_NONROOT\_USER](../../../config/latest/developerguide/ecs-task-definition-nonroot-user.md)

- [ECS\_TASK\_DEFINITION\_PID\_MODE\_CHECK](../../../config/latest/developerguide/ecs-task-definition-pid-mode-check.md)

- [ECS\_TASK\_DEFINITION\_USER\_FOR\_HOST\_MODE\_CHECK](../../../config/latest/developerguide/ecs-task-definition-user-for-host-mode-check.md)

- [EFS\_ACCESS\_POINT\_ENFORCE\_ROOT\_DIRECTORY](../../../config/latest/developerguide/efs-access-point-enforce-root-directory.md)

- [EFS\_ACCESS\_POINT\_ENFORCE\_USER\_IDENTITY](../../../config/latest/developerguide/efs-access-point-enforce-user-identity.md)

- [EFS\_ENCRYPTED\_CHECK](../../../config/latest/developerguide/efs-encrypted-check.md)

- [EFS\_IN\_BACKUP\_PLAN](../../../config/latest/developerguide/efs-in-backup-plan.md)

- [EFS\_LAST\_BACKUP\_RECOVERY\_POINT\_CREATED](../../../config/latest/developerguide/efs-last-backup-recovery-point-created.md)

- [EFS\_RESOURCES\_PROTECTED\_BY\_BACKUP\_PLAN](../../../config/latest/developerguide/efs-resources-protected-by-backup-plan.md)

- [EIP\_ATTACHED](../../../config/latest/developerguide/eip-attached.md)

- [EKS\_CLUSTER\_LOGGING\_ENABLED](../../../config/latest/developerguide/eks-cluster-logging-enabled.md)

- [EKS\_CLUSTER\_OLDEST\_SUPPORTED\_VERSION](../../../config/latest/developerguide/eks-cluster-oldest-supported-version.md)

- [EKS\_CLUSTER\_SUPPORTED\_VERSION](../../../config/latest/developerguide/eks-cluster-supported-version.md)

- [EKS\_ENDPOINT\_NO\_PUBLIC\_ACCESS](../../../config/latest/developerguide/eks-endpoint-no-public-access.md)

- [EKS\_SECRETS\_ENCRYPTED](../../../config/latest/developerguide/eks-secrets-encrypted.md)

- [ELASTIC\_BEANSTALK\_LOGS\_TO\_CLOUDWATCH](../../../config/latest/developerguide/elastic-beanstalk-logs-to-cloudwatch.md)

- [ELASTIC\_BEANSTALK\_MANAGED\_UPDATES\_ENABLED](../../../config/latest/developerguide/elastic-beanstalk-managed-updates-enabled.md)

- [ELASTICACHE\_AUTO\_MINOR\_VERSION\_UPGRADE\_CHECK](../../../config/latest/developerguide/elasticache-auto-minor-version-upgrade-check.md)

- [ELASTICACHE\_RBAC\_AUTH\_ENABLED](../../../config/latest/developerguide/elasticache-rbac-auth-enabled.md)

- [ELASTICACHE\_REDIS\_CLUSTER\_AUTOMATIC\_BACKUP\_CHECK](../../../config/latest/developerguide/elasticache-redis-cluster-automatic-backup-check.md)

- [ELASTICACHE\_REPL\_GRP\_AUTO\_FAILOVER\_ENABLED](../../../config/latest/developerguide/elasticache-repl-grp-auto-failover-enabled.md)

- [ELASTICACHE\_REPL\_GRP\_ENCRYPTED\_AT\_REST](../../../config/latest/developerguide/elasticache-repl-grp-encrypted-at-rest.md)

- [ELASTICACHE\_REPL\_GRP\_ENCRYPTED\_IN\_TRANSIT](../../../config/latest/developerguide/elasticache-repl-grp-encrypted-in-transit.md)

- [ELASTICACHE\_REPL\_GRP\_REDIS\_AUTH\_ENABLED](../../../config/latest/developerguide/elasticache-repl-grp-redis-auth-enabled.md)

- [ELASTICACHE\_SUBNET\_GROUP\_CHECK](../../../config/latest/developerguide/elasticache-subnet-group-check.md)

- [ELASTICACHE\_SUPPORTED\_ENGINE\_VERSION](../../../config/latest/developerguide/elasticache-supported-engine-version.md)

- [ELASTICSEARCH\_ENCRYPTED\_AT\_REST](../../../config/latest/developerguide/elasticsearch-encrypted-at-rest.md)

- [ELASTICSEARCH\_IN\_VPC\_ONLY](../../../config/latest/developerguide/elasticsearch-in-vpc-only.md)

- [ELASTICSEARCH\_LOGS\_TO\_CLOUDWATCH](../../../config/latest/developerguide/elasticsearch-logs-to-cloudwatch.md)

- [ELASTICSEARCH\_NODE\_TO\_NODE\_ENCRYPTION\_CHECK](../../../config/latest/developerguide/elasticsearch-node-to-node-encryption-check.md)

- [ELB\_ACM\_CERTIFICATE\_REQUIRED](../../../config/latest/developerguide/elb-acm-certificate-required.md)

- [ELB\_CROSS\_ZONE\_LOAD\_BALANCING\_ENABLED](../../../config/latest/developerguide/elb-cross-zone-load-balancing-enabled.md)

- [ELB\_CUSTOM\_SECURITY\_POLICY\_SSL\_CHECK](../../../config/latest/developerguide/elb-custom-security-policy-ssl-check.md)

- [ELB\_DELETION\_PROTECTION\_ENABLED](../../../config/latest/developerguide/elb-deletion-protection-enabled.md)

- [ELB\_LOGGING\_ENABLED](../../../config/latest/developerguide/elb-logging-enabled.md)

- [ELB\_PREDEFINED\_SECURITY\_POLICY\_SSL\_CHECK](../../../config/latest/developerguide/elb-predefined-security-policy-ssl-check.md)

- [ELB\_TLS\_HTTPS\_LISTENERS\_ONLY](../../../config/latest/developerguide/elb-tls-https-listeners-only.md)

- [ELBV2\_ACM\_CERTIFICATE\_REQUIRED](../../../config/latest/developerguide/elbv2-acm-certificate-required.md)

- [ELBV2\_MULTIPLE\_AZ](../../../config/latest/developerguide/elbv2-multiple-az.md)

- [EMR\_KERBEROS\_ENABLED](../../../config/latest/developerguide/emr-kerberos-enabled.md)

- [EMR\_MASTER\_NO\_PUBLIC\_IP](../../../config/latest/developerguide/emr-master-no-public-ip.md)

- [ENCRYPTED\_VOLUMES](../../../config/latest/developerguide/encrypted-volumes.md)

- [FMS\_SHIELD\_RESOURCE\_POLICY\_CHECK](../../../config/latest/developerguide/fms-shield-resource-policy-check.md)

- [FMS\_WEBACL\_RESOURCE\_POLICY\_CHECK](../../../config/latest/developerguide/fms-webacl-resource-policy-check.md)

- [FMS\_WEBACL\_RULEGROUP\_ASSOCIATION\_CHECK](../../../config/latest/developerguide/fms-webacl-rulegroup-association-check.md)

- [FSX\_LAST\_BACKUP\_RECOVERY\_POINT\_CREATED](../../../config/latest/developerguide/fsx-last-backup-recovery-point-created.md)

- [FSX\_RESOURCES\_PROTECTED\_BY\_BACKUP\_PLAN](../../../config/latest/developerguide/fsx-resources-protected-by-backup-plan.md)

- [GUARDDUTY\_ENABLED\_CENTRALIZED](../../../config/latest/developerguide/guardduty-enabled-centralized.md)

- [GUARDDUTY\_NON\_ARCHIVED\_FINDINGS](../../../config/latest/developerguide/guardduty-non-archived-findings.md)

- [IAM\_CUSTOMER\_POLICY\_BLOCKED\_KMS\_ACTIONS](../../../config/latest/developerguide/iam-inline-policy-blocked-kms-actions.md)

- [IAM\_GROUP\_HAS\_USERS\_CHECK](../../../config/latest/developerguide/iam-group-has-users-check.md)

- [IAM\_INLINE\_POLICY\_BLOCKED\_KMS\_ACTIONS](../../../config/latest/developerguide/iam-inline-policy-blocked-kms-actions.md)

- [IAM\_NO\_INLINE\_POLICY\_CHECK](../../../config/latest/developerguide/iam-no-inline-policy-check.md)

- [IAM\_PASSWORD\_POLICY](../../../config/latest/developerguide/iam-password-policy.md)

- [IAM\_POLICY\_BLACKLISTED\_CHECK](../../../config/latest/developerguide/iam-policy-blacklisted-check.md)

- [IAM\_POLICY\_IN\_USE](../../../config/latest/developerguide/iam-policy-in-use.md)

- [IAM\_POLICY\_NO\_STATEMENTS\_WITH\_ADMIN\_ACCESS](../../../config/latest/developerguide/iam-policy-no-statements-with-admin-access.md)

- [IAM\_POLICY\_NO\_STATEMENTS\_WITH\_FULL\_ACCESS](../../../config/latest/developerguide/iam-policy-no-statements-with-full-access.md)

- [IAM\_ROLE\_MANAGED\_POLICY\_CHECK](../../../config/latest/developerguide/iam-role-managed-policy-check.md)

- [IAM\_ROOT\_ACCESS\_KEY\_CHECK](../../../config/latest/developerguide/iam-root-access-key-check.md)

- [IAM\_USER\_GROUP\_MEMBERSHIP\_CHECK](../../../config/latest/developerguide/iam-user-group-membership-check.md)

- [IAM\_USER\_MFA\_ENABLED](../../../config/latest/developerguide/iam-user-mfa-enabled.md)

- [IAM\_USER\_NO\_POLICIES\_CHECK](../../../config/latest/developerguide/iam-user-no-policies-check.md)

- [IAM\_USER\_UNUSED\_CREDENTIALS\_CHECK](../../../config/latest/developerguide/iam-user-unused-credentials-check.md)

- [INCOMING\_SSH\_DISABLED](../../../config/latest/developerguide/restricted-ssh.md)

- [INSTANCES\_IN\_VPC](../../../config/latest/developerguide/ec2-instances-in-vpc.md)

- [KINESIS\_STREAM\_ENCRYPTED](../../../config/latest/developerguide/kinesis-stream-encrypted.md)

- [INTERNET\_GATEWAY\_AUTHORIZED\_VPC\_ONLY](../../../config/latest/developerguide/internet-gateway-authorized-vpc-only.md)

- [KMS\_CMK\_NOT\_SCHEDULED\_FOR\_DELETION](../../../config/latest/developerguide/kms-cmk-not-scheduled-for-deletion.md)

- [LAMBDA\_CONCURRENCY\_CHECK](../../../config/latest/developerguide/lambda-concurrency-check.md)

- [LAMBDA\_DLQ\_CHECK](../../../config/latest/developerguide/lambda-dlq-check.md)

- [LAMBDA\_FUNCTION\_PUBLIC\_ACCESS\_PROHIBITED](../../../config/latest/developerguide/lambda-function-public-access-prohibited.md)

- [LAMBDA\_FUNCTION\_SETTINGS\_CHECK](../../../config/latest/developerguide/lambda-function-settings-check.md)

- [LAMBDA\_INSIDE\_VPC](../../../config/latest/developerguide/lambda-inside-vpc.md)

- [LAMBDA\_VPC\_MULTI\_AZ\_CHECK](../../../config/latest/developerguide/lambda-vpc-multi-az-check.md)

- [MACIE\_STATUS\_CHECK](../../../config/latest/developerguide/macie-status-check.md)

- [MFA\_ENABLED\_FOR\_IAM\_CONSOLE\_ACCESS](../../../config/latest/developerguide/mfa-enabled-for-iam-console-access.md)

- [MQ\_AUTOMATIC\_MINOR\_VERSION\_UPGRADE\_ENABLED](../../../config/latest/developerguide/mq-automatic-minor-version-upgrade-enabled.md)

- [MQ\_CLOUDWATCH\_AUDIT\_LOGGING\_ENABLED](../../../config/latest/developerguide/mq-cloudwatch-audit-logging-enabled.md)

- [MQ\_NO\_PUBLIC\_ACCESS](../../../config/latest/developerguide/mq-no-public-access.md)

- [MULTI\_REGION\_CLOUD\_TRAIL\_ENABLED](../../../config/latest/developerguide/multi-region-cloudtrail-enabled.md)

- [NACL\_NO\_UNRESTRICTED\_SSH\_RDP](../../../config/latest/developerguide/nacl-no-unrestricted-ssh-rdp.md)

- [NETFW\_LOGGING\_ENABLED](../../../config/latest/developerguide/netfw-logging-enabled.md)

- [NETFW\_MULTI\_AZ\_ENABLED](../../../config/latest/developerguide/netfw-multi-az-enabled.md)

- [NETFW\_POLICY\_DEFAULT\_ACTION\_FRAGMENT\_PACKETS](../../../config/latest/developerguide/netfw-policy-default-action-fragment-packets.md)

- [NETFW\_POLICY\_DEFAULT\_ACTION\_FULL\_PACKETS](../../../config/latest/developerguide/netfw-policy-default-action-full-packets.md)

- [NETFW\_POLICY\_RULE\_GROUP\_ASSOCIATED](../../../config/latest/developerguide/netfw-policy-rule-group-associated.md)

- [NETFW\_STATELESS\_RULE\_GROUP\_NOT\_EMPTY](../../../config/latest/developerguide/netfw-stateless-rule-group-not-empty.md)

- [NLB\_CROSS\_ZONE\_LOAD\_BALANCING\_ENABLED](../../../config/latest/developerguide/nlb-cross-zone-load-balancing-enabled.md)

- [NO\_UNRESTRICTED\_ROUTE\_TO\_IGW](../../../config/latest/developerguide/no-unrestricted-route-to-igw.md)

- [OPENSEARCH\_ACCESS\_CONTROL\_ENABLED](../../../config/latest/developerguide/opensearch-access-control-enabled.md)

- [OPENSEARCH\_AUDIT\_LOGGING\_ENABLED](../../../config/latest/developerguide/opensearch-audit-logging-enabled.md)

- [OPENSEARCH\_DATA\_NODE\_FAULT\_TOLERANCE](../../../config/latest/developerguide/opensearch-data-node-fault-tolerance.md)

- [OPENSEARCH\_ENCRYPTED\_AT\_REST](../../../config/latest/developerguide/opensearch-encrypted-at-rest.md)

- [OPENSEARCH\_HTTPS\_REQUIRED](../../../config/latest/developerguide/opensearch-https-required.md)

- [OPENSEARCH\_IN\_VPC\_ONLY](../../../config/latest/developerguide/opensearch-in-vpc-only.md)

- [OPENSEARCH\_LOGS\_TO\_CLOUDWATCH](../../../config/latest/developerguide/opensearch-logs-to-cloudwatch.md)

- [OPENSEARCH\_NODE\_TO\_NODE\_ENCRYPTION\_CHECK](../../../config/latest/developerguide/opensearch-node-to-node-encryption-check.md)

- [RDS\_AUTOMATIC\_MINOR\_VERSION\_UPGRADE\_ENABLED](../../../config/latest/developerguide/rds-automatic-minor-version-upgrade-enabled.md)

- [RDS\_CLUSTER\_DEFAULT\_ADMIN\_CHECK](../../../config/latest/developerguide/rds-cluster-default-admin-check.md)

- [RDS\_CLUSTER\_DELETION\_PROTECTION\_ENABLED](../../../config/latest/developerguide/rds-cluster-deletion-protection-enabled.md)

- [RDS\_CLUSTER\_IAM\_AUTHENTICATION\_ENABLED](../../../config/latest/developerguide/rds-cluster-iam-authentication-enabled.md)

- [RDS\_CLUSTER\_MULTI\_AZ\_ENABLED](../../../config/latest/developerguide/rds-cluster-multi-az-enabled.md)

- [RDS\_DB\_SECURITY\_GROUP\_NOT\_ALLOWED](../../../config/latest/developerguide/rds-db-security-group-not-allowed.md)

- [RDS\_ENHANCED\_MONITORING\_ENABLED](../../../config/latest/developerguide/rds-enhanced-monitoring-enabled.md)

- [RDS\_IN\_BACKUP\_PLAN](../../../config/latest/developerguide/rds-in-backup-plan.md)

- [RDS\_INSTANCE\_DEFAULT\_ADMIN\_CHECK](../../../config/latest/developerguide/rds-instance-default-admin-check.md)

- [RDS\_INSTANCE\_DELETION\_PROTECTION\_ENABLED](../../../config/latest/developerguide/rds-instance-deletion-protection-enabled.md)

- [RDS\_INSTANCE\_IAM\_AUTHENTICATION\_ENABLED](../../../config/latest/developerguide/rds-instance-iam-authentication-enabled.md)

- [RDS\_INSTANCE\_PUBLIC\_ACCESS\_CHECK](../../../config/latest/developerguide/rds-instance-public-access-check.md)

- [RDS\_LAST\_BACKUP\_RECOVERY\_POINT\_CREATED](../../../config/latest/developerguide/rds-last-backup-recovery-point-created.md)

- [RDS\_LOGGING\_ENABLED](../../../config/latest/developerguide/rds-logging-enabled.md)

- [RDS\_MULTI\_AZ\_SUPPORT](../../../config/latest/developerguide/rds-multi-az-support.md)

- [RDS\_RESOURCES\_PROTECTED\_BY\_BACKUP\_PLAN](../../../config/latest/developerguide/rds-resources-protected-by-backup-plan.md)

- [RDS\_SNAPSHOT\_ENCRYPTED](../../../config/latest/developerguide/rds-snapshot-encrypted.md)

- [RDS\_SNAPSHOTS\_PUBLIC\_PROHIBITED](../../../config/latest/developerguide/rds-snapshots-public-prohibited.md)

- [RDS\_STORAGE\_ENCRYPTED](../../../config/latest/developerguide/rds-storage-encrypted.md)

- [REDSHIFT\_BACKUP\_ENABLED](../../../config/latest/developerguide/redshift-backup-enabled.md)

- [REDSHIFT\_REQUIRE\_TLS\_SSL](../../../config/latest/developerguide/redshift-require-tls-ssl.md)

- [REDSHIFT\_CLUSTER\_CONFIGURATION\_CHECK](../../../config/latest/developerguide/redshift-cluster-configuration-check.md)

- [REDSHIFT\_CLUSTER\_MAINTENANCESETTINGS\_CHECK](../../../config/latest/developerguide/redshift-cluster-maintenancesettings-check.md)

- [REDSHIFT\_CLUSTER\_PUBLIC\_ACCESS\_CHECK](../../../config/latest/developerguide/redshift-cluster-public-access-check.md)

- [REDSHIFT\_AUDIT\_LOGGING\_ENABLED](../../../config/latest/developerguide/redshift-audit-logging-enabled.md)

- [REDSHIFT\_CLUSTER\_KMS\_ENABLED](../../../config/latest/developerguide/redshift-cluster-kms-enabled.md)

- [REDSHIFT\_DEFAULT\_ADMIN\_CHECK](../../../config/latest/developerguide/redshift-default-admin-check.md)

- [REDSHIFT\_DEFAULT\_DB\_NAME\_CHECK](../../../config/latest/developerguide/redshift-default-db-name-check.md)

- [REDSHIFT\_ENHANCED\_VPC\_ROUTING\_ENABLED](../../../config/latest/developerguide/redshift-enhanced-vpc-routing-enabled.md)

- [REQUIRED\_TAGS](../../../config/latest/developerguide/required-tags.md)

- [RESTRICTED\_INCOMING\_TRAFFIC](../../../config/latest/developerguide/restricted-common-ports.md)

- [ROOT\_ACCOUNT\_HARDWARE\_MFA\_ENABLED](../../../config/latest/developerguide/root-account-hardware-mfa-enabled.md)

- [ROOT\_ACCOUNT\_MFA\_ENABLED](../../../config/latest/developerguide/root-account-mfa-enabled.md)

- [S3\_ACCOUNT\_LEVEL\_PUBLIC\_ACCESS\_BLOCKS\_PERIODIC](../../../config/latest/developerguide/s3-account-level-public-access-blocks.md)

- [S3\_ACCOUNT\_LEVEL\_PUBLIC\_ACCESS\_BLOCKS](../../../config/latest/developerguide/s3-account-level-public-access-blocks.md)

- [S3\_BUCKET\_ACL\_PROHIBITED](../../../config/latest/developerguide/s3-bucket-acl-prohibited.md)

- [S3\_BUCKET\_BLACKLISTED\_ACTIONS\_PROHIBITED](../../../config/latest/developerguide/s3-bucket-blacklisted-actions-prohibited.md)

- [S3\_BUCKET\_DEFAULT\_LOCK\_ENABLED](../../../config/latest/developerguide/s3-bucket-default-lock-enabled.md)

- [S3\_BUCKET\_LEVEL\_PUBLIC\_ACCESS\_PROHIBITED](../../../config/latest/developerguide/s3-bucket-level-public-access-prohibited.md)

- [S3\_BUCKET\_LOGGING\_ENABLED](../../../config/latest/developerguide/s3-bucket-logging-enabled.md)

- [S3\_BUCKET\_POLICY\_GRANTEE\_CHECK](../../../config/latest/developerguide/s3-bucket-policy-grantee-check.md)

- [S3\_BUCKET\_POLICY\_NOT\_MORE\_PERMISSIVE](../../../config/latest/developerguide/s3-bucket-policy-not-more-permissive.md)

- [S3\_BUCKET\_PUBLIC\_READ\_PROHIBITED](../../../config/latest/developerguide/s3-bucket-public-read-prohibited.md)

- [S3\_BUCKET\_PUBLIC\_WRITE\_PROHIBITED](../../../config/latest/developerguide/s3-bucket-public-write-prohibited.md)

- [S3\_BUCKET\_REPLICATION\_ENABLED](../../../config/latest/developerguide/s3-bucket-replication-enabled.md)

- [S3\_BUCKET\_SERVER\_SIDE\_ENCRYPTION\_ENABLED](../../../config/latest/developerguide/s3-bucket-server-side-encryption-enabled.md)

- [S3\_BUCKET\_SSL\_REQUESTS\_ONLY](../../../config/latest/developerguide/s3-bucket-ssl-requests-only.md)

- [S3\_BUCKET\_VERSIONING\_ENABLED](../../../config/latest/developerguide/s3-bucket-versioning-enabled.md)

- [S3\_DEFAULT\_ENCRYPTION\_KMS](../../../config/latest/developerguide/s3-default-encryption-kms.md)

- [S3\_EVENT\_NOTIFICATIONS\_ENABLED](../../../config/latest/developerguide/s3-event-notifications-enabled.md)

- [S3\_LAST\_BACKUP\_RECOVERY\_POINT\_CREATED](../../../config/latest/developerguide/s3-last-backup-recovery-point-created.md)

- [S3\_LIFECYCLE\_POLICY\_CHECK](../../../config/latest/developerguide/s3-lifecycle-policy-check.md)

- [S3\_RESOURCES\_PROTECTED\_BY\_BACKUP\_PLAN](../../../config/latest/developerguide/s3-resources-protected-by-backup-plan.md)

- [S3\_VERSION\_LIFECYCLE\_POLICY\_CHECK](../../../config/latest/developerguide/s3-version-lifecycle-policy-check.md)

- [SAGEMAKER\_ENDPOINT\_CONFIGURATION\_KMS\_KEY\_CONFIGURED](../../../config/latest/developerguide/sagemaker-endpoint-configuration-kms-key-configured.md)

- [SAGEMAKER\_NOTEBOOK\_INSTANCE\_INSIDE\_VPC](../../../config/latest/developerguide/sagemaker-notebook-instance-inside-vpc.md)

- [SAGEMAKER\_NOTEBOOK\_INSTANCE\_KMS\_KEY\_CONFIGURED](../../../config/latest/developerguide/sagemaker-notebook-instance-kms-key-configured.md)

- [SAGEMAKER\_NOTEBOOK\_INSTANCE\_ROOT\_ACCESS\_CHECK](../../../config/latest/developerguide/sagemaker-notebook-instance-root-access-check.md)

- [SAGEMAKER\_NOTEBOOK\_NO\_DIRECT\_INTERNET\_ACCESS](../../../config/latest/developerguide/sagemaker-notebook-no-direct-internet-access.md)

- [SECRETSMANAGER\_ROTATION\_ENABLED\_CHECK](../../../config/latest/developerguide/secretsmanager-rotation-enabled-check.md)

- [SECRETSMANAGER\_SCHEDULED\_ROTATION\_SUCCESS\_CHECK](../../../config/latest/developerguide/secretsmanager-scheduled-rotation-success-check.md)

- [SECRETSMANAGER\_SECRET\_PERIODIC\_ROTATION](../../../config/latest/developerguide/secretsmanager-secret-periodic-rotation.md)

- [SECRETSMANAGER\_SECRET\_UNUSED](../../../config/latest/developerguide/secretsmanager-secret-unused.md)

- [SECRETSMANAGER\_USING\_CMK](../../../config/latest/developerguide/secretsmanager-using-cmk.md)

- [SECURITY\_ACCOUNT\_INFORMATION\_PROVIDED](../../../config/latest/developerguide/security-account-information-provided.md)

- [SECURITYHUB\_ENABLED](../../../config/latest/developerguide/securityhub-enabled.md)

- [SERVICE\_VPC\_ENDPOINT\_ENABLED](../../../config/latest/developerguide/service-vpc-endpoint-enabled.md)

- [SES\_MALWARE\_SCANNING\_ENABLED](../../../config/latest/developerguide/ses-malware-scanning-enabled.md)

- [SHIELD\_ADVANCED\_ENABLED\_AUTORENEW](../../../config/latest/developerguide/shield-advanced-enabled-autorenew.md)

- [SHIELD\_DRT\_ACCESS](../../../config/latest/developerguide/shield-drt-access.md)

- [SNS\_ENCRYPTED\_KMS](../../../config/latest/developerguide/sns-encrypted-kms.md)

- [SNS\_TOPIC\_MESSAGE\_DELIVERY\_NOTIFICATION\_ENABLED](../../../config/latest/developerguide/sns-topic-message-delivery-notification-enabled.md)

- [SSM\_DOCUMENT\_NOT\_PUBLIC](../../../config/latest/developerguide/ssm-document-not-public.md)

- [STEP\_FUNCTIONS\_STATE\_MACHINE\_LOGGING\_ENABLED](../../../config/latest/developerguide/step-functions-state-machine-logging-enabled.md)

- [STORAGEGATEWAY\_LAST\_BACKUP\_RECOVERY\_POINT\_CREATED](../../../config/latest/developerguide/storagegateway-last-backup-recovery-point-created.md)

- [STORAGEGATEWAY\_RESOURCES\_PROTECTED\_BY\_BACKUP\_PLAN](../../../config/latest/developerguide/storagegateway-resources-protected-by-backup-plan.md)

- [SUBNET\_AUTO\_ASSIGN\_PUBLIC\_IP\_DISABLED](../../../config/latest/developerguide/subnet-auto-assign-public-ip-disabled.md)

- [VIRTUALMACHINE\_LAST\_BACKUP\_RECOVERY\_POINT\_CREATED](../../../config/latest/developerguide/virtualmachine-last-backup-recovery-point-created.md)

- [VIRTUALMACHINE\_RESOURCES\_PROTECTED\_BY\_BACKUP\_PLAN](../../../config/latest/developerguide/virtualmachine-resources-protected-by-backup-plan.md)

- [VPC\_DEFAULT\_SECURITY\_GROUP\_CLOSED](../../../config/latest/developerguide/vpc-default-security-group-closed.md)

- [VPC\_FLOW\_LOGS\_ENABLED](../../../config/latest/developerguide/vpc-flow-logs-enabled.md)

- [VPC\_NETWORK\_ACL\_UNUSED\_CHECK](../../../config/latest/developerguide/vpc-network-acl-unused-check.md)

- [VPC\_PEERING\_DNS\_RESOLUTION\_CHECK](../../../config/latest/developerguide/vpc-peering-dns-resolution-check.md)

- [VPC\_SG\_OPEN\_ONLY\_TO\_AUTHORIZED\_PORTS](../../../config/latest/developerguide/vpc-sg-open-only-to-authorized-ports.md)

- [VPC\_VPN\_2\_TUNNELS\_UP](../../../config/latest/developerguide/vpc-vpn-2-tunnels-up.md)

- [WAF\_CLASSIC\_LOGGING\_ENABLED](../../../config/latest/developerguide/waf-classic-logging-enabled.md)

- [WAF\_GLOBAL\_RULEGROUP\_NOT\_EMPTY](../../../config/latest/developerguide/waf-global-rulegroup-not-empty.md)

- [WAF\_GLOBAL\_RULE\_NOT\_EMPTY](../../../config/latest/developerguide/waf-global-rule-not-empty.md)

- [WAF\_GLOBAL\_WEBACL\_NOT\_EMPTY](../../../config/latest/developerguide/waf-global-webacl-not-empty.md)

- [WAF\_REGIONAL\_RULEGROUP\_NOT\_EMPTY](../../../config/latest/developerguide/waf-regional-rulegroup-not-empty.md)

- [WAF\_REGIONAL\_RULE\_NOT\_EMPTY](../../../config/latest/developerguide/waf-regional-rule-not-empty.md)

- [WAF\_REGIONAL\_WEBACL\_NOT\_EMPTY](../../../config/latest/developerguide/waf-regional-webacl-not-empty.md)

- [WAFV2\_LOGGING\_ENABLED](../../../config/latest/developerguide/wafv2-logging-enabled.md)

- [WAFV2\_RULEGROUP\_NOT\_EMPTY](../../../config/latest/developerguide/wafv2-rulegroup-not-empty.md)

- [WAFV2\_WEBACL\_NOT\_EMPTY](../../../config/latest/developerguide/wafv2-webacl-not-empty.md)

## Using AWS Config custom rules with Audit Manager

You can use AWS Config custom rules as a data source for audit reporting. When a control
has a data source that's mapped to an AWS Config rule, Audit Manager adds the evaluation that was
created by the AWS Config rule.

The custom rules that you can use depend on the AWS account that you sign in to Audit Manager
with. If you can access a custom rule in AWS Config, you can use it as a data source mapping
in Audit Manager.

- **For individual AWS accounts** – You can use
any of the custom rules that you created with your account.

- **For accounts that are part of an organization**
– Either, you can use any of your member-level custom rules. Or, you can use any
of the organization-level custom rules that are available to you in AWS Config.

After you map your custom rules as a data source for a control, you can add that control
to a custom framework in Audit Manager.

## Additional resources

- To find help with issues for this data source type, see [My assessment isn’t collecting compliance check evidence from AWS Config](evidence-collection-issues.md#no-evidence-from-config) and [AWS Config integration issues](control-issues.md#config-rule-integration.title).

- To create a custom control using this data source type, see [Creating a custom control in AWS Audit Manager](create-controls.md).

- To create a custom framework that uses your custom control, see [Creating a custom framework in AWS Audit Manager](custom-frameworks.md).

- To add your custom control to an existing custom framework, see [Editing a custom framework in AWS Audit Manager](edit-custom-frameworks.md).

- To create a custom rule in AWS Config, see [Developing a\
custom rule for AWS Config](../../../config/latest/developerguide/evaluate-config-develop-rules.md) in the _AWS Config Developer_
_Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Supported data sources

AWS Security Hub CSPM

All content copied from https://docs.aws.amazon.com/.
