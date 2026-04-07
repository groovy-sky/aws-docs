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

- [Key points](https://docs.aws.amazon.com/audit-manager/latest/userguide/control-data-sources-config.html#aws-config-key-points)

- [Supported AWS Config managed rules](https://docs.aws.amazon.com/audit-manager/latest/userguide/control-data-sources-config.html#aws-config-managed-rules)

- [Using AWS Config custom rules with Audit Manager](https://docs.aws.amazon.com/audit-manager/latest/userguide/control-data-sources-config.html#aws-config-custom-rules)

- [Additional resources](https://docs.aws.amazon.com/audit-manager/latest/userguide/control-data-sources-config.html#aws-config-rules-troubleshoot)

## Key points

- Audit Manager doesn’t collect evidence from [service-linked\
AWS Config rules](https://docs.aws.amazon.com/config/latest/developerguide/service-linked-awsconfig-rules.html), with the exception of service-linked rules from Conformance
Packs and from AWS Organizations.

- Audit Manager doesn't manage AWS Config rules for you. Before you start evidence collection, we
recommend that you review your current AWS Config rule parameters. Then, validate those
parameters against the requirements of your chosen framework. If needed, you can [update a rule's\
parameters in AWS Config](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_manage-rules.html) so that it aligns with framework requirements. This will
help to ensure that your assessments collect the correct compliance check evidence for
that framework.

For example, suppose that you’re creating an assessment for CIS v1.2.0. This
framework has a control named [Ensure IAM password policy\
requires a minimum length of 14 or greater](https://docs.aws.amazon.com/securityhub/latest/userguide/iam-controls.html#iam-15). In AWS Config, the [iam-password-policy](https://docs.aws.amazon.com/config/latest/developerguide/iam-password-policy.html) rule has a `MinimumPasswordLength` parameter
that checks password length. The default value for this parameter is 14 characters. As a
result, the rule aligns with the control requirements. If you aren’t using the default
parameter value, ensure that the value you’re using is equal to or greater than the 14
character requirement from CIS v1.2.0. You can find the default parameter details for
each managed rule in the [AWS Config\
documentation](https://docs.aws.amazon.com/config/latest/developerguide/managed-rules-by-aws-config.html).

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
Rules](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_use-managed-rules.html) in the _AWS Config User Guide_.

###### Tip

When you choose a managed rule in the Audit Manager console during custom control creation,
make sure that you look for one of the following rule identifier keywords, and not the
rule name. For information about the difference between the rule name and rule identifier,
and how to find the identifier for a managed rule, see the [Troubleshooting](https://docs.aws.amazon.com/audit-manager/latest/userguide/control-issues.html#managed-rule-missing) section of this user guide.

Supported AWS Config managed rule keywords

- [ACCESS\_KEYS\_ROTATED](https://docs.aws.amazon.com/config/latest/developerguide/access-keys-rotated.html)

- [ACCOUNT\_PART\_OF\_ORGANIZATIONS](https://docs.aws.amazon.com/config/latest/developerguide/account-part-of-organizations.html)

- [ACM\_CERTIFICATE\_EXPIRATION\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/acm-certificate-expiration-check.html)

- [ACM\_CERTIFICATE\_RSA\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/acm-certificate-rsa-check.html)

- [ALB\_DESYNC\_MODE\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/alb-desync-mode-check.html)

- [ALB\_HTTP\_DROP\_INVALID\_HEADER\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/alb-http-drop-invalid-header-enabled.html)

- [ALB\_HTTP\_TO\_HTTPS\_REDIRECTION\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/alb-http-to-https-redirection-check.html)

- [ALB\_WAF\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/alb-waf-enabled.html)

- [API\_GW\_ASSOCIATED\_WITH\_WAF](https://docs.aws.amazon.com/config/latest/developerguide/api-gw-associated-with-waf.html)

- [API\_GW\_CACHE\_ENABLED\_AND\_ENCRYPTED](https://docs.aws.amazon.com/config/latest/developerguide/api-gw-cache-enabled-and-encrypted.html)

- [API\_GW\_ENDPOINT\_TYPE\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/api-gw-endpoint-type-check.html)

- [API\_GW\_EXECUTION\_LOGGING\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/api-gw-execution-logging-enabled.html)

- [API\_GW\_SSL\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/api-gw-ssl-enabled.html)

- [API\_GW\_XRAY\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/api-gw-xray-enabled.html)

- [API\_GWV2\_ACCESS\_LOGS\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/api-gwv2-access-logs-enabled.html)

- [API\_GWV2\_AUTHORIZATION\_TYPE\_CONFIGURED](https://docs.aws.amazon.com/config/latest/developerguide/api-gwv2-authorization-type-configured.html)

- [APPROVED\_AMIS\_BY\_ID](https://docs.aws.amazon.com/config/latest/developerguide/approved-amis-by-id.html)

- [APPROVED\_AMIS\_BY\_TAG](https://docs.aws.amazon.com/config/latest/developerguide/approved-amis-by-tag.html)

- [APPSYNC\_ASSOCIATED\_WITH\_WAF](https://docs.aws.amazon.com/config/latest/developerguide/appsync-associated-with-waf.html)

- [APPSYNC\_CACHE\_ENCRYPTION\_AT\_REST](https://docs.aws.amazon.com/config/latest/developerguide/appsync-cache-encryption-at-rest.html)

- [APPSYNC\_LOGGING\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/appsync-logging-enabled.html)

- [AURORA\_LAST\_BACKUP\_RECOVERY\_POINT\_CREATED](https://docs.aws.amazon.com/config/latest/developerguide/aurora-last-backup-recovery-point-created.html)

- [AURORA\_MYSQL\_BACKTRACKING\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/aurora-mysql-backtracking-enabled.html)

- [AURORA\_RESOURCES\_PROTECTED\_BY\_BACKUP\_PLAN](https://docs.aws.amazon.com/config/latest/developerguide/aurora-resources-protected-by-backup-plan.html)

- [AUTOSCALING\_CAPACITY\_REBALANCING](https://docs.aws.amazon.com/config/latest/developerguide/autoscaling-capacity-rebalancing.html)

- [AUTOSCALING\_GROUP\_ELB\_HEALTHCHECK\_REQUIRED](https://docs.aws.amazon.com/config/latest/developerguide/autoscaling-group-elb-healthcheck-required.html)

- [AUTOSCALING\_LAUNCH\_CONFIG\_HOP\_LIMIT](https://docs.aws.amazon.com/config/latest/developerguide/autoscaling-launch-config-hop-limit.html)

- [AUTOSCALING\_LAUNCH\_CONFIG\_PUBLIC\_IP\_DISABLED](https://docs.aws.amazon.com/config/latest/developerguide/autoscaling-launch-config-public-ip-disabled.html)

- [AUTOSCALING\_LAUNCHCONFIG\_REQUIRES\_IMDSV2](https://docs.aws.amazon.com/config/latest/developerguide/autoscaling-launchconfig-requires-imdsv2.html)

- [AUTOSCALING\_LAUNCH\_TEMPLATE](https://docs.aws.amazon.com/config/latest/developerguide/autoscaling-launch-template.html)

- [AUTOSCALING\_MULTIPLE\_AZ](https://docs.aws.amazon.com/config/latest/developerguide/autoscaling-multiple-az.html)

- [AUTOSCALING\_MULTIPLE\_INSTANCE\_TYPES](https://docs.aws.amazon.com/config/latest/developerguide/autoscaling-multiple-instance-types.html)

- [BACKUP\_PLAN\_MIN\_FREQUENCY\_AND\_MIN\_RETENTION\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/backup-plan-min-frequency-and-min-retention-check.html)

- [BACKUP\_RECOVERY\_POINT\_ENCRYPTED](https://docs.aws.amazon.com/config/latest/developerguide/backup-recovery-point-encrypted.html)

- [BACKUP\_RECOVERY\_POINT\_MANUAL\_DELETION\_DISABLED](https://docs.aws.amazon.com/config/latest/developerguide/backup-recovery-point-manual-deletion-disabled.html)

- [BACKUP\_RECOVERY\_POINT\_MINIMUM\_RETENTION\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/backup-recovery-point-minimum-retention-check.html)

- [BEANSTALK\_ENHANCED\_HEALTH\_REPORTING\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/beanstalk-enhanced-health-reporting-enabled.html)

- [CLB\_DESYNC\_MODE\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/clb-desync-mode-check.html)

- [CLB\_MULTIPLE\_AZ](https://docs.aws.amazon.com/config/latest/developerguide/clb-multiple-az.html)

- [CLOUD\_TRAIL\_CLOUD\_WATCH\_LOGS\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/cloud-trail-cloud-watch-logs-enabled.html)

- [CLOUD\_TRAIL\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/cloudtrail-enabled.html)

- [CLOUD\_TRAIL\_ENCRYPTION\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/cloud-trail-encryption-enabled.html)

- [CLOUD\_TRAIL\_LOG\_FILE\_VALIDATION\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/cloud-trail-log-file-validation-enabled.html)

- [CLOUDFORMATION\_STACK\_DRIFT\_DETECTION\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/cloudformation-stack-drift-detection-check.html)

- [CLOUDFORMATION\_STACK\_NOTIFICATION\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/cloudformation-stack-notification-check.html)

- [CLOUDFRONT\_ACCESSLOGS\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/cloudfront-accesslogs-enabled.html)

- [CLOUDFRONT\_ASSOCIATED\_WITH\_WAF](https://docs.aws.amazon.com/config/latest/developerguide/cloudfront-associated-with-waf.html)

- [CLOUDFRONT\_CUSTOM\_SSL\_CERTIFICATE](https://docs.aws.amazon.com/config/latest/developerguide/cloudfront-custom-ssl-certificate.html)

- [CLOUDFRONT\_DEFAULT\_ROOT\_OBJECT\_CONFIGURED](https://docs.aws.amazon.com/config/latest/developerguide/cloudfront-default-root-object-configured.html)

- [CLOUDFRONT\_NO\_DEPRECATED\_SSL\_PROTOCOLS](https://docs.aws.amazon.com/config/latest/developerguide/cloudfront-no-deprecated-ssl-protocols.html)

- [CLOUDFRONT\_ORIGIN\_ACCESS\_IDENTITY\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/cloudfront-origin-access-identity-enabled.html)

- [CLOUDFRONT\_ORIGIN\_FAILOVER\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/cloudfront-origin-failover-enabled.html)

- [CLOUDFRONT\_S3\_ORIGIN\_ACCESS\_CONTROL\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/cloudfront-s3-origin-access-control-enabled.html)

- [CLOUDFRONT\_S3\_ORIGIN\_NON\_EXISTENT\_BUCKET](https://docs.aws.amazon.com/config/latest/developerguide/cloudfront-s3-origin-non-existent-bucket.html)

- [CLOUDFRONT\_SECURITY\_POLICY\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/cloudfront-security-policy-check.html)

- [CLOUDFRONT\_SNI\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/cloudfront-sni-enabled.html)

- [CLOUDFRONT\_TRAFFIC\_TO\_ORIGIN\_ENCRYPTED](https://docs.aws.amazon.com/config/latest/developerguide/cloudfront-traffic-to-origin-encrypted.html)

- [CLOUDFRONT\_VIEWER\_POLICY\_HTTPS](https://docs.aws.amazon.com/config/latest/developerguide/cloudfront-viewer-policy-https.html)

- [CLOUDTRAIL\_S3\_DATAEVENTS\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/cloudtrail-s3-dataevents-enabled.html)

- [CLOUDTRAIL\_SECURITY\_TRAIL\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/cloudtrail-security-trail-enabled.html)

- [CLOUDWATCH\_ALARM\_ACTION\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/cloudwatch-alarm-action-check.html)

- [CLOUDWATCH\_ALARM\_ACTION\_ENABLED\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/cloudwatch-alarm-action-enabled-check.html)

- [CLOUDWATCH\_ALARM\_RESOURCE\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/cloudwatch-alarm-resource-check.html)

- [CLOUDWATCH\_ALARM\_SETTINGS\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/cloudwatch-alarm-settings-check.html)

- [CLOUDWATCH\_LOG\_GROUP\_ENCRYPTED](https://docs.aws.amazon.com/config/latest/developerguide/cloudwatch-log-group-encrypted.html)

- [CMK\_BACKING\_KEY\_ROTATION\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/cmk-backing-key-rotation-enabled.html)

- [CODEBUILD\_PROJECT\_ARTIFACT\_ENCRYPTION](https://docs.aws.amazon.com/config/latest/developerguide/codebuild-project-artifact-encryption.html)

- [CODEBUILD\_PROJECT\_ENVIRONMENT\_PRIVILEGED\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/codebuild-project-environment-privileged-check.html)

- [CODEBUILD\_PROJECT\_ENVVAR\_AWSCRED\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/codebuild-project-envvar-awscred-check.html)

- [CODEBUILD\_PROJECT\_LOGGING\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/codebuild-project-logging-enabled.html)

- [CODEBUILD\_PROJECT\_S3\_LOGS\_ENCRYPTED](https://docs.aws.amazon.com/config/latest/developerguide/codebuild-project-s3-logs-encrypted.html)

- [CODEBUILD\_PROJECT\_SOURCE\_REPO\_URL\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/codebuild-project-source-repo-url-check.html)

- [CODEDEPLOY\_AUTO\_ROLLBACK\_MONITOR\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/codedeploy-auto-rollback-monitor-enabled.html)

- [CODEDEPLOY\_EC2\_MINIMUM\_HEALTHY\_HOSTS\_CONFIGURED](https://docs.aws.amazon.com/config/latest/developerguide/codedeploy-ec2-minimum-healthy-hosts-configured.html)

- [CODEDEPLOY\_LAMBDA\_ALLATONCE\_TRAFFIC\_SHIFT\_DISABLED](https://docs.aws.amazon.com/config/latest/developerguide/codedeploy-lambda-allatonce-traffic-shift-disabled.html)

- [CODEPIPELINE\_DEPLOYMENT\_COUNT\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/codepipeline-deployment-count-check.html)

- [CODEPIPELINE\_REGION\_FANOUT\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/codepipeline-region-fanout-check.html)

- [CUSTOM\_SCHEMA\_REGISTRY\_POLICY\_ATTACHED](https://docs.aws.amazon.com/config/latest/developerguide/custom-schema-registry-policy-attached.html)

- [CW\_LOGGROUP\_RETENTION\_PERIOD\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/cw-loggroup-retention-period-check.html)

- [DAX\_ENCRYPTION\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/dax-encryption-enabled.html)

- [DB\_INSTANCE\_BACKUP\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/db-instance-backup-enabled.html)

- [DESIRED\_INSTANCE\_TENANCY](https://docs.aws.amazon.com/config/latest/developerguide/desired-instance-tenancy.html)

- [DESIRED\_INSTANCE\_TYPE](https://docs.aws.amazon.com/config/latest/developerguide/desired-instance-type.html)

- [DMS\_REPLICATION\_NOT\_PUBLIC](https://docs.aws.amazon.com/config/latest/developerguide/dms-replication-not-public.html)

- [DYNAMODB\_AUTOSCALING\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/dynamodb-autoscaling-enabled.html)

- [DYNAMODB\_IN\_BACKUP\_PLAN](https://docs.aws.amazon.com/config/latest/developerguide/dynamodb-in-backup-plan.html)

- [DYNAMODB\_LAST\_BACKUP\_RECOVERY\_POINT\_CREATED](https://docs.aws.amazon.com/config/latest/developerguide/dynamodb-last-backup-recovery-point-created.html)

- [DYNAMODB\_PITR\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/dynamodb-pitr-enabled.html)

- [DYNAMODB\_RESOURCES\_PROTECTED\_BY\_BACKUP\_PLAN](https://docs.aws.amazon.com/config/latest/developerguide/dynamodb-resources-protected-by-backup-plan.html)

- [DYNAMODB\_TABLE\_ENCRYPTED\_KMS](https://docs.aws.amazon.com/config/latest/developerguide/dynamodb-table-encrypted-kms.html)

- [DYNAMODB\_TABLE\_ENCRYPTION\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/dynamodb-table-encryption-enabled.html)

- [DYNAMODB\_THROUGHPUT\_LIMIT\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/dynamodb-throughput-limit-check.html)

- [EBS\_IN\_BACKUP\_PLAN](https://docs.aws.amazon.com/config/latest/developerguide/ebs-in-backup-plan.html)

- [EBS\_LAST\_BACKUP\_RECOVERY\_POINT\_CREATED](https://docs.aws.amazon.com/config/latest/developerguide/ebs-last-backup-recovery-point-created.html)

- [EBS\_OPTIMIZED\_INSTANCE](https://docs.aws.amazon.com/config/latest/developerguide/ebs-optimized-instance.html)

- [EBS\_RESOURCES\_PROTECTED\_BY\_BACKUP\_PLAN](https://docs.aws.amazon.com/config/latest/developerguide/ebs-resources-protected-by-backup-plan.html)

- [EBS\_SNAPSHOT\_PUBLIC\_RESTORABLE\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/ebs-snapshot-public-restorable-check.html)

- [EC2\_CLIENT\_VPN\_NOT\_AUTHORIZE\_ALL](https://docs.aws.amazon.com/config/latest/developerguide/ec2-client-vpn-not-authorize-all.html)

- [EC2\_EBS\_ENCRYPTION\_BY\_DEFAULT](https://docs.aws.amazon.com/config/latest/developerguide/ec2-ebs-encryption-by-default.html)

- [EC2\_IMDSV2\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/ec2-imdsv2-check.html)

- [EC2\_INSTANCE\_DETAILED\_MONITORING\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/ec2-instance-detailed-monitoring-enabled.html)

- [EC2\_INSTANCE\_MANAGED\_BY\_SSM](https://docs.aws.amazon.com/config/latest/developerguide/ec2-instance-managed-by-systems-manager.html)

- [EC2\_INSTANCE\_MULTIPLE\_ENI\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/ec2-instance-multiple-eni-check.html)

- [EC2\_INSTANCE\_NO\_PUBLIC\_IP](https://docs.aws.amazon.com/config/latest/developerguide/ec2-instance-no-public-ip.html)

- [EC2\_INSTANCE\_PROFILE\_ATTACHED](https://docs.aws.amazon.com/config/latest/developerguide/ec2-instance-profile-attached.html)

- [EC2\_LAST\_BACKUP\_RECOVERY\_POINT\_CREATED](https://docs.aws.amazon.com/config/latest/developerguide/ec2-last-backup-recovery-point-created.html)

- [EC2\_LAUNCH\_TEMPLATE\_PUBLIC\_IP\_DISABLED](https://docs.aws.amazon.com/config/latest/developerguide/ec2-launch-template-public-ip-disabled.html)

- [EC2\_MANAGEDINSTANCE\_APPLICATIONS\_BLACKLISTED](https://docs.aws.amazon.com/config/latest/developerguide/ec2-managedinstance-applications-blacklisted.html)

- [EC2\_MANAGEDINSTANCE\_APPLICATIONS\_REQUIRED](https://docs.aws.amazon.com/config/latest/developerguide/ec2-managedinstance-applications-required.html)

- [EC2\_MANAGEDINSTANCE\_ASSOCIATION\_COMPLIANCE\_STATUS\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/ec2-managedinstance-association-compliance-status-check.html)

- [EC2\_MANAGEDINSTANCE\_INVENTORY\_BLACKLISTED](https://docs.aws.amazon.com/config/latest/developerguide/ec2-managedinstance-inventory-blacklisted.html)

- [EC2\_MANAGEDINSTANCE\_PATCH\_COMPLIANCE\_STATUS\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/ec2-managedinstance-patch-compliance-status-check.html)

- [EC2\_MANAGEDINSTANCE\_PLATFORM\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/ec2-managedinstance-platform-check.html)

- [EC2\_NO\_AMAZON\_KEY\_PAIR](https://docs.aws.amazon.com/config/latest/developerguide/ec2-no-amazon-key-pair.html)

- [EC2\_PARAVIRTUAL\_INSTANCE\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/ec2-paravirtual-instance-check.html)

- [EC2\_RESOURCES\_PROTECTED\_BY\_BACKUP\_PLAN](https://docs.aws.amazon.com/config/latest/developerguide/ec2-resources-protected-by-backup-plan.html)

- [EC2\_SECURITY\_GROUP\_ATTACHED\_TO\_ENI](https://docs.aws.amazon.com/config/latest/developerguide/ec2-security-group-attached-to-eni.html)

- [EC2\_SECURITY\_GROUP\_ATTACHED\_TO\_ENI\_PERIODIC](https://docs.aws.amazon.com/config/latest/developerguide/ec2-security-group-attached-to-eni-periodic.html)

- [EC2\_STOPPED\_INSTANCE](https://docs.aws.amazon.com/config/latest/developerguide/ec2-stopped-instance.html)

- [EC2\_TOKEN\_HOP\_LIMIT\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/ec2-token-hop-limit-check.html)

- [EC2\_TRANSIT\_GATEWAY\_AUTO\_VPC\_ATTACH\_DISABLED](https://docs.aws.amazon.com/config/latest/developerguide/ec2-transit-gateway-auto-vpc-attach-disabled.html)

- [EC2\_VOLUME\_INUSE\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/ec2-volume-inuse-check.html)

- [ECR\_PRIVATE\_IMAGE\_SCANNING\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/ecr-private-image-scanning-enabled.html)

- [ECR\_PRIVATE\_LIFECYCLE\_POLICY\_CONFIGURED](https://docs.aws.amazon.com/config/latest/developerguide/ecr-private-lifecycle-policy-configured.html)

- [ECR\_PRIVATE\_TAG\_IMMUTABILITY\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/ecr-private-tag-immutability-enabled.html)

- [ECS\_AWSVPC\_NETWORKING\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/ecs-awsvpc-networking-enabled.html)

- [ECS\_CONTAINER\_INSIGHTS\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/ecs-container-insights-enabled.html)

- [ECS\_CONTAINERS\_NONPRIVILEGED](https://docs.aws.amazon.com/config/latest/developerguide/ecs-containers-nonprivileged.html)

- [ECS\_CONTAINERS\_READONLY\_ACCESS](https://docs.aws.amazon.com/config/latest/developerguide/ecs-containers-readonly-access.html)

- [ECS\_FARGATE\_LATEST\_PLATFORM\_VERSION](https://docs.aws.amazon.com/config/latest/developerguide/ecs-fargate-latest-platform-version.html)

- [ECS\_NO\_ENVIRONMENT\_SECRETS](https://docs.aws.amazon.com/config/latest/developerguide/ecs-no-environment-secrets.html)

- [ECS\_TASK\_DEFINITION\_LOG\_CONFIGURATION](https://docs.aws.amazon.com/config/latest/developerguide/ecs-task-definition-log-configuration.html)

- [ECS\_TASK\_DEFINITION\_MEMORY\_HARD\_LIMIT](https://docs.aws.amazon.com/config/latest/developerguide/ecs-task-definition-memory-hard-limit.html)

- [ECS\_TASK\_DEFINITION\_NONROOT\_USER](https://docs.aws.amazon.com/config/latest/developerguide/ecs-task-definition-nonroot-user.html)

- [ECS\_TASK\_DEFINITION\_PID\_MODE\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/ecs-task-definition-pid-mode-check.html)

- [ECS\_TASK\_DEFINITION\_USER\_FOR\_HOST\_MODE\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/ecs-task-definition-user-for-host-mode-check.html)

- [EFS\_ACCESS\_POINT\_ENFORCE\_ROOT\_DIRECTORY](https://docs.aws.amazon.com/config/latest/developerguide/efs-access-point-enforce-root-directory.html)

- [EFS\_ACCESS\_POINT\_ENFORCE\_USER\_IDENTITY](https://docs.aws.amazon.com/config/latest/developerguide/efs-access-point-enforce-user-identity.html)

- [EFS\_ENCRYPTED\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/efs-encrypted-check.html)

- [EFS\_IN\_BACKUP\_PLAN](https://docs.aws.amazon.com/config/latest/developerguide/efs-in-backup-plan.html)

- [EFS\_LAST\_BACKUP\_RECOVERY\_POINT\_CREATED](https://docs.aws.amazon.com/config/latest/developerguide/efs-last-backup-recovery-point-created.html)

- [EFS\_RESOURCES\_PROTECTED\_BY\_BACKUP\_PLAN](https://docs.aws.amazon.com/config/latest/developerguide/efs-resources-protected-by-backup-plan.html)

- [EIP\_ATTACHED](https://docs.aws.amazon.com/config/latest/developerguide/eip-attached.html)

- [EKS\_CLUSTER\_LOGGING\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/eks-cluster-logging-enabled.html)

- [EKS\_CLUSTER\_OLDEST\_SUPPORTED\_VERSION](https://docs.aws.amazon.com/config/latest/developerguide/eks-cluster-oldest-supported-version.html)

- [EKS\_CLUSTER\_SUPPORTED\_VERSION](https://docs.aws.amazon.com/config/latest/developerguide/eks-cluster-supported-version.html)

- [EKS\_ENDPOINT\_NO\_PUBLIC\_ACCESS](https://docs.aws.amazon.com/config/latest/developerguide/eks-endpoint-no-public-access.html)

- [EKS\_SECRETS\_ENCRYPTED](https://docs.aws.amazon.com/config/latest/developerguide/eks-secrets-encrypted.html)

- [ELASTIC\_BEANSTALK\_LOGS\_TO\_CLOUDWATCH](https://docs.aws.amazon.com/config/latest/developerguide/elastic-beanstalk-logs-to-cloudwatch.html)

- [ELASTIC\_BEANSTALK\_MANAGED\_UPDATES\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/elastic-beanstalk-managed-updates-enabled.html)

- [ELASTICACHE\_AUTO\_MINOR\_VERSION\_UPGRADE\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/elasticache-auto-minor-version-upgrade-check.html)

- [ELASTICACHE\_RBAC\_AUTH\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/elasticache-rbac-auth-enabled.html)

- [ELASTICACHE\_REDIS\_CLUSTER\_AUTOMATIC\_BACKUP\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/elasticache-redis-cluster-automatic-backup-check.html)

- [ELASTICACHE\_REPL\_GRP\_AUTO\_FAILOVER\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/elasticache-repl-grp-auto-failover-enabled.html)

- [ELASTICACHE\_REPL\_GRP\_ENCRYPTED\_AT\_REST](https://docs.aws.amazon.com/config/latest/developerguide/elasticache-repl-grp-encrypted-at-rest.html)

- [ELASTICACHE\_REPL\_GRP\_ENCRYPTED\_IN\_TRANSIT](https://docs.aws.amazon.com/config/latest/developerguide/elasticache-repl-grp-encrypted-in-transit.html)

- [ELASTICACHE\_REPL\_GRP\_REDIS\_AUTH\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/elasticache-repl-grp-redis-auth-enabled.html)

- [ELASTICACHE\_SUBNET\_GROUP\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/elasticache-subnet-group-check.html)

- [ELASTICACHE\_SUPPORTED\_ENGINE\_VERSION](https://docs.aws.amazon.com/config/latest/developerguide/elasticache-supported-engine-version.html)

- [ELASTICSEARCH\_ENCRYPTED\_AT\_REST](https://docs.aws.amazon.com/config/latest/developerguide/elasticsearch-encrypted-at-rest.html)

- [ELASTICSEARCH\_IN\_VPC\_ONLY](https://docs.aws.amazon.com/config/latest/developerguide/elasticsearch-in-vpc-only.html)

- [ELASTICSEARCH\_LOGS\_TO\_CLOUDWATCH](https://docs.aws.amazon.com/config/latest/developerguide/elasticsearch-logs-to-cloudwatch.html)

- [ELASTICSEARCH\_NODE\_TO\_NODE\_ENCRYPTION\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/elasticsearch-node-to-node-encryption-check.html)

- [ELB\_ACM\_CERTIFICATE\_REQUIRED](https://docs.aws.amazon.com/config/latest/developerguide/elb-acm-certificate-required.html)

- [ELB\_CROSS\_ZONE\_LOAD\_BALANCING\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/elb-cross-zone-load-balancing-enabled.html)

- [ELB\_CUSTOM\_SECURITY\_POLICY\_SSL\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/elb-custom-security-policy-ssl-check.html)

- [ELB\_DELETION\_PROTECTION\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/elb-deletion-protection-enabled.html)

- [ELB\_LOGGING\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/elb-logging-enabled.html)

- [ELB\_PREDEFINED\_SECURITY\_POLICY\_SSL\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/elb-predefined-security-policy-ssl-check.html)

- [ELB\_TLS\_HTTPS\_LISTENERS\_ONLY](https://docs.aws.amazon.com/config/latest/developerguide/elb-tls-https-listeners-only.html)

- [ELBV2\_ACM\_CERTIFICATE\_REQUIRED](https://docs.aws.amazon.com/config/latest/developerguide/elbv2-acm-certificate-required.html)

- [ELBV2\_MULTIPLE\_AZ](https://docs.aws.amazon.com/config/latest/developerguide/elbv2-multiple-az.html)

- [EMR\_KERBEROS\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/emr-kerberos-enabled.html)

- [EMR\_MASTER\_NO\_PUBLIC\_IP](https://docs.aws.amazon.com/config/latest/developerguide/emr-master-no-public-ip.html)

- [ENCRYPTED\_VOLUMES](https://docs.aws.amazon.com/config/latest/developerguide/encrypted-volumes.html)

- [FMS\_SHIELD\_RESOURCE\_POLICY\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/fms-shield-resource-policy-check.html)

- [FMS\_WEBACL\_RESOURCE\_POLICY\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/fms-webacl-resource-policy-check.html)

- [FMS\_WEBACL\_RULEGROUP\_ASSOCIATION\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/fms-webacl-rulegroup-association-check.html)

- [FSX\_LAST\_BACKUP\_RECOVERY\_POINT\_CREATED](https://docs.aws.amazon.com/config/latest/developerguide/fsx-last-backup-recovery-point-created.html)

- [FSX\_RESOURCES\_PROTECTED\_BY\_BACKUP\_PLAN](https://docs.aws.amazon.com/config/latest/developerguide/fsx-resources-protected-by-backup-plan.html)

- [GUARDDUTY\_ENABLED\_CENTRALIZED](https://docs.aws.amazon.com/config/latest/developerguide/guardduty-enabled-centralized.html)

- [GUARDDUTY\_NON\_ARCHIVED\_FINDINGS](https://docs.aws.amazon.com/config/latest/developerguide/guardduty-non-archived-findings.html)

- [IAM\_CUSTOMER\_POLICY\_BLOCKED\_KMS\_ACTIONS](https://docs.aws.amazon.com/config/latest/developerguide/iam-inline-policy-blocked-kms-actions.html)

- [IAM\_GROUP\_HAS\_USERS\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/iam-group-has-users-check.html)

- [IAM\_INLINE\_POLICY\_BLOCKED\_KMS\_ACTIONS](https://docs.aws.amazon.com/config/latest/developerguide/iam-inline-policy-blocked-kms-actions.html)

- [IAM\_NO\_INLINE\_POLICY\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/iam-no-inline-policy-check.html)

- [IAM\_PASSWORD\_POLICY](https://docs.aws.amazon.com/config/latest/developerguide/iam-password-policy.html)

- [IAM\_POLICY\_BLACKLISTED\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/iam-policy-blacklisted-check.html)

- [IAM\_POLICY\_IN\_USE](https://docs.aws.amazon.com/config/latest/developerguide/iam-policy-in-use.html)

- [IAM\_POLICY\_NO\_STATEMENTS\_WITH\_ADMIN\_ACCESS](https://docs.aws.amazon.com/config/latest/developerguide/iam-policy-no-statements-with-admin-access.html)

- [IAM\_POLICY\_NO\_STATEMENTS\_WITH\_FULL\_ACCESS](https://docs.aws.amazon.com/config/latest/developerguide/iam-policy-no-statements-with-full-access.html)

- [IAM\_ROLE\_MANAGED\_POLICY\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/iam-role-managed-policy-check.html)

- [IAM\_ROOT\_ACCESS\_KEY\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/iam-root-access-key-check.html)

- [IAM\_USER\_GROUP\_MEMBERSHIP\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/iam-user-group-membership-check.html)

- [IAM\_USER\_MFA\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/iam-user-mfa-enabled.html)

- [IAM\_USER\_NO\_POLICIES\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/iam-user-no-policies-check.html)

- [IAM\_USER\_UNUSED\_CREDENTIALS\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/iam-user-unused-credentials-check.html)

- [INCOMING\_SSH\_DISABLED](https://docs.aws.amazon.com/config/latest/developerguide/restricted-ssh.html)

- [INSTANCES\_IN\_VPC](https://docs.aws.amazon.com/config/latest/developerguide/ec2-instances-in-vpc.html)

- [KINESIS\_STREAM\_ENCRYPTED](https://docs.aws.amazon.com/config/latest/developerguide/kinesis-stream-encrypted.html)

- [INTERNET\_GATEWAY\_AUTHORIZED\_VPC\_ONLY](https://docs.aws.amazon.com/config/latest/developerguide/internet-gateway-authorized-vpc-only.html)

- [KMS\_CMK\_NOT\_SCHEDULED\_FOR\_DELETION](https://docs.aws.amazon.com/config/latest/developerguide/kms-cmk-not-scheduled-for-deletion.html)

- [LAMBDA\_CONCURRENCY\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/lambda-concurrency-check.html)

- [LAMBDA\_DLQ\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/lambda-dlq-check.html)

- [LAMBDA\_FUNCTION\_PUBLIC\_ACCESS\_PROHIBITED](https://docs.aws.amazon.com/config/latest/developerguide/lambda-function-public-access-prohibited.html)

- [LAMBDA\_FUNCTION\_SETTINGS\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/lambda-function-settings-check.html)

- [LAMBDA\_INSIDE\_VPC](https://docs.aws.amazon.com/config/latest/developerguide/lambda-inside-vpc.html)

- [LAMBDA\_VPC\_MULTI\_AZ\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/lambda-vpc-multi-az-check.html)

- [MACIE\_STATUS\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/macie-status-check.html)

- [MFA\_ENABLED\_FOR\_IAM\_CONSOLE\_ACCESS](https://docs.aws.amazon.com/config/latest/developerguide/mfa-enabled-for-iam-console-access.html)

- [MQ\_AUTOMATIC\_MINOR\_VERSION\_UPGRADE\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/mq-automatic-minor-version-upgrade-enabled.html)

- [MQ\_CLOUDWATCH\_AUDIT\_LOGGING\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/mq-cloudwatch-audit-logging-enabled.html)

- [MQ\_NO\_PUBLIC\_ACCESS](https://docs.aws.amazon.com/config/latest/developerguide/mq-no-public-access.html)

- [MULTI\_REGION\_CLOUD\_TRAIL\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/multi-region-cloudtrail-enabled.html)

- [NACL\_NO\_UNRESTRICTED\_SSH\_RDP](https://docs.aws.amazon.com/config/latest/developerguide/nacl-no-unrestricted-ssh-rdp.html)

- [NETFW\_LOGGING\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/netfw-logging-enabled.html)

- [NETFW\_MULTI\_AZ\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/netfw-multi-az-enabled.html)

- [NETFW\_POLICY\_DEFAULT\_ACTION\_FRAGMENT\_PACKETS](https://docs.aws.amazon.com/config/latest/developerguide/netfw-policy-default-action-fragment-packets.html)

- [NETFW\_POLICY\_DEFAULT\_ACTION\_FULL\_PACKETS](https://docs.aws.amazon.com/config/latest/developerguide/netfw-policy-default-action-full-packets.html)

- [NETFW\_POLICY\_RULE\_GROUP\_ASSOCIATED](https://docs.aws.amazon.com/config/latest/developerguide/netfw-policy-rule-group-associated.html)

- [NETFW\_STATELESS\_RULE\_GROUP\_NOT\_EMPTY](https://docs.aws.amazon.com/config/latest/developerguide/netfw-stateless-rule-group-not-empty.html)

- [NLB\_CROSS\_ZONE\_LOAD\_BALANCING\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/nlb-cross-zone-load-balancing-enabled.html)

- [NO\_UNRESTRICTED\_ROUTE\_TO\_IGW](https://docs.aws.amazon.com/config/latest/developerguide/no-unrestricted-route-to-igw.html)

- [OPENSEARCH\_ACCESS\_CONTROL\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/opensearch-access-control-enabled.html)

- [OPENSEARCH\_AUDIT\_LOGGING\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/opensearch-audit-logging-enabled.html)

- [OPENSEARCH\_DATA\_NODE\_FAULT\_TOLERANCE](https://docs.aws.amazon.com/config/latest/developerguide/opensearch-data-node-fault-tolerance.html)

- [OPENSEARCH\_ENCRYPTED\_AT\_REST](https://docs.aws.amazon.com/config/latest/developerguide/opensearch-encrypted-at-rest.html)

- [OPENSEARCH\_HTTPS\_REQUIRED](https://docs.aws.amazon.com/config/latest/developerguide/opensearch-https-required.html)

- [OPENSEARCH\_IN\_VPC\_ONLY](https://docs.aws.amazon.com/config/latest/developerguide/opensearch-in-vpc-only.html)

- [OPENSEARCH\_LOGS\_TO\_CLOUDWATCH](https://docs.aws.amazon.com/config/latest/developerguide/opensearch-logs-to-cloudwatch.html)

- [OPENSEARCH\_NODE\_TO\_NODE\_ENCRYPTION\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/opensearch-node-to-node-encryption-check.html)

- [RDS\_AUTOMATIC\_MINOR\_VERSION\_UPGRADE\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/rds-automatic-minor-version-upgrade-enabled.html)

- [RDS\_CLUSTER\_DEFAULT\_ADMIN\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/rds-cluster-default-admin-check.html)

- [RDS\_CLUSTER\_DELETION\_PROTECTION\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/rds-cluster-deletion-protection-enabled.html)

- [RDS\_CLUSTER\_IAM\_AUTHENTICATION\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/rds-cluster-iam-authentication-enabled.html)

- [RDS\_CLUSTER\_MULTI\_AZ\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/rds-cluster-multi-az-enabled.html)

- [RDS\_DB\_SECURITY\_GROUP\_NOT\_ALLOWED](https://docs.aws.amazon.com/config/latest/developerguide/rds-db-security-group-not-allowed.html)

- [RDS\_ENHANCED\_MONITORING\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/rds-enhanced-monitoring-enabled.html)

- [RDS\_IN\_BACKUP\_PLAN](https://docs.aws.amazon.com/config/latest/developerguide/rds-in-backup-plan.html)

- [RDS\_INSTANCE\_DEFAULT\_ADMIN\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/rds-instance-default-admin-check.html)

- [RDS\_INSTANCE\_DELETION\_PROTECTION\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/rds-instance-deletion-protection-enabled.html)

- [RDS\_INSTANCE\_IAM\_AUTHENTICATION\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/rds-instance-iam-authentication-enabled.html)

- [RDS\_INSTANCE\_PUBLIC\_ACCESS\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/rds-instance-public-access-check.html)

- [RDS\_LAST\_BACKUP\_RECOVERY\_POINT\_CREATED](https://docs.aws.amazon.com/config/latest/developerguide/rds-last-backup-recovery-point-created.html)

- [RDS\_LOGGING\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/rds-logging-enabled.html)

- [RDS\_MULTI\_AZ\_SUPPORT](https://docs.aws.amazon.com/config/latest/developerguide/rds-multi-az-support.html)

- [RDS\_RESOURCES\_PROTECTED\_BY\_BACKUP\_PLAN](https://docs.aws.amazon.com/config/latest/developerguide/rds-resources-protected-by-backup-plan.html)

- [RDS\_SNAPSHOT\_ENCRYPTED](https://docs.aws.amazon.com/config/latest/developerguide/rds-snapshot-encrypted.html)

- [RDS\_SNAPSHOTS\_PUBLIC\_PROHIBITED](https://docs.aws.amazon.com/config/latest/developerguide/rds-snapshots-public-prohibited.html)

- [RDS\_STORAGE\_ENCRYPTED](https://docs.aws.amazon.com/config/latest/developerguide/rds-storage-encrypted.html)

- [REDSHIFT\_BACKUP\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/redshift-backup-enabled.html)

- [REDSHIFT\_REQUIRE\_TLS\_SSL](https://docs.aws.amazon.com/config/latest/developerguide/redshift-require-tls-ssl.html)

- [REDSHIFT\_CLUSTER\_CONFIGURATION\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/redshift-cluster-configuration-check.html)

- [REDSHIFT\_CLUSTER\_MAINTENANCESETTINGS\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/redshift-cluster-maintenancesettings-check.html)

- [REDSHIFT\_CLUSTER\_PUBLIC\_ACCESS\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/redshift-cluster-public-access-check.html)

- [REDSHIFT\_AUDIT\_LOGGING\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/redshift-audit-logging-enabled.html)

- [REDSHIFT\_CLUSTER\_KMS\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/redshift-cluster-kms-enabled.html)

- [REDSHIFT\_DEFAULT\_ADMIN\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/redshift-default-admin-check.html)

- [REDSHIFT\_DEFAULT\_DB\_NAME\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/redshift-default-db-name-check.html)

- [REDSHIFT\_ENHANCED\_VPC\_ROUTING\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/redshift-enhanced-vpc-routing-enabled.html)

- [REQUIRED\_TAGS](https://docs.aws.amazon.com/config/latest/developerguide/required-tags.html)

- [RESTRICTED\_INCOMING\_TRAFFIC](https://docs.aws.amazon.com/config/latest/developerguide/restricted-common-ports.html)

- [ROOT\_ACCOUNT\_HARDWARE\_MFA\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/root-account-hardware-mfa-enabled.html)

- [ROOT\_ACCOUNT\_MFA\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/root-account-mfa-enabled.html)

- [S3\_ACCOUNT\_LEVEL\_PUBLIC\_ACCESS\_BLOCKS\_PERIODIC](https://docs.aws.amazon.com/config/latest/developerguide/s3-account-level-public-access-blocks.html)

- [S3\_ACCOUNT\_LEVEL\_PUBLIC\_ACCESS\_BLOCKS](https://docs.aws.amazon.com/config/latest/developerguide/s3-account-level-public-access-blocks.html)

- [S3\_BUCKET\_ACL\_PROHIBITED](https://docs.aws.amazon.com/config/latest/developerguide/s3-bucket-acl-prohibited.html)

- [S3\_BUCKET\_BLACKLISTED\_ACTIONS\_PROHIBITED](https://docs.aws.amazon.com/config/latest/developerguide/s3-bucket-blacklisted-actions-prohibited.html)

- [S3\_BUCKET\_DEFAULT\_LOCK\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/s3-bucket-default-lock-enabled.html)

- [S3\_BUCKET\_LEVEL\_PUBLIC\_ACCESS\_PROHIBITED](https://docs.aws.amazon.com/config/latest/developerguide/s3-bucket-level-public-access-prohibited.html)

- [S3\_BUCKET\_LOGGING\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/s3-bucket-logging-enabled.html)

- [S3\_BUCKET\_POLICY\_GRANTEE\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/s3-bucket-policy-grantee-check.html)

- [S3\_BUCKET\_POLICY\_NOT\_MORE\_PERMISSIVE](https://docs.aws.amazon.com/config/latest/developerguide/s3-bucket-policy-not-more-permissive.html)

- [S3\_BUCKET\_PUBLIC\_READ\_PROHIBITED](https://docs.aws.amazon.com/config/latest/developerguide/s3-bucket-public-read-prohibited.html)

- [S3\_BUCKET\_PUBLIC\_WRITE\_PROHIBITED](https://docs.aws.amazon.com/config/latest/developerguide/s3-bucket-public-write-prohibited.html)

- [S3\_BUCKET\_REPLICATION\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/s3-bucket-replication-enabled.html)

- [S3\_BUCKET\_SERVER\_SIDE\_ENCRYPTION\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/s3-bucket-server-side-encryption-enabled.html)

- [S3\_BUCKET\_SSL\_REQUESTS\_ONLY](https://docs.aws.amazon.com/config/latest/developerguide/s3-bucket-ssl-requests-only.html)

- [S3\_BUCKET\_VERSIONING\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/s3-bucket-versioning-enabled.html)

- [S3\_DEFAULT\_ENCRYPTION\_KMS](https://docs.aws.amazon.com/config/latest/developerguide/s3-default-encryption-kms.html)

- [S3\_EVENT\_NOTIFICATIONS\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/s3-event-notifications-enabled.html)

- [S3\_LAST\_BACKUP\_RECOVERY\_POINT\_CREATED](https://docs.aws.amazon.com/config/latest/developerguide/s3-last-backup-recovery-point-created.html)

- [S3\_LIFECYCLE\_POLICY\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/s3-lifecycle-policy-check.html)

- [S3\_RESOURCES\_PROTECTED\_BY\_BACKUP\_PLAN](https://docs.aws.amazon.com/config/latest/developerguide/s3-resources-protected-by-backup-plan.html)

- [S3\_VERSION\_LIFECYCLE\_POLICY\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/s3-version-lifecycle-policy-check.html)

- [SAGEMAKER\_ENDPOINT\_CONFIGURATION\_KMS\_KEY\_CONFIGURED](https://docs.aws.amazon.com/config/latest/developerguide/sagemaker-endpoint-configuration-kms-key-configured.html)

- [SAGEMAKER\_NOTEBOOK\_INSTANCE\_INSIDE\_VPC](https://docs.aws.amazon.com/config/latest/developerguide/sagemaker-notebook-instance-inside-vpc.html)

- [SAGEMAKER\_NOTEBOOK\_INSTANCE\_KMS\_KEY\_CONFIGURED](https://docs.aws.amazon.com/config/latest/developerguide/sagemaker-notebook-instance-kms-key-configured.html)

- [SAGEMAKER\_NOTEBOOK\_INSTANCE\_ROOT\_ACCESS\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/sagemaker-notebook-instance-root-access-check.html)

- [SAGEMAKER\_NOTEBOOK\_NO\_DIRECT\_INTERNET\_ACCESS](https://docs.aws.amazon.com/config/latest/developerguide/sagemaker-notebook-no-direct-internet-access.html)

- [SECRETSMANAGER\_ROTATION\_ENABLED\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/secretsmanager-rotation-enabled-check.html)

- [SECRETSMANAGER\_SCHEDULED\_ROTATION\_SUCCESS\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/secretsmanager-scheduled-rotation-success-check.html)

- [SECRETSMANAGER\_SECRET\_PERIODIC\_ROTATION](https://docs.aws.amazon.com/config/latest/developerguide/secretsmanager-secret-periodic-rotation.html)

- [SECRETSMANAGER\_SECRET\_UNUSED](https://docs.aws.amazon.com/config/latest/developerguide/secretsmanager-secret-unused.html)

- [SECRETSMANAGER\_USING\_CMK](https://docs.aws.amazon.com/config/latest/developerguide/secretsmanager-using-cmk.html)

- [SECURITY\_ACCOUNT\_INFORMATION\_PROVIDED](https://docs.aws.amazon.com/config/latest/developerguide/security-account-information-provided.html)

- [SECURITYHUB\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/securityhub-enabled.html)

- [SERVICE\_VPC\_ENDPOINT\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/service-vpc-endpoint-enabled.html)

- [SES\_MALWARE\_SCANNING\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/ses-malware-scanning-enabled.html)

- [SHIELD\_ADVANCED\_ENABLED\_AUTORENEW](https://docs.aws.amazon.com/config/latest/developerguide/shield-advanced-enabled-autorenew.html)

- [SHIELD\_DRT\_ACCESS](https://docs.aws.amazon.com/config/latest/developerguide/shield-drt-access.html)

- [SNS\_ENCRYPTED\_KMS](https://docs.aws.amazon.com/config/latest/developerguide/sns-encrypted-kms.html)

- [SNS\_TOPIC\_MESSAGE\_DELIVERY\_NOTIFICATION\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/sns-topic-message-delivery-notification-enabled.html)

- [SSM\_DOCUMENT\_NOT\_PUBLIC](https://docs.aws.amazon.com/config/latest/developerguide/ssm-document-not-public.html)

- [STEP\_FUNCTIONS\_STATE\_MACHINE\_LOGGING\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/step-functions-state-machine-logging-enabled.html)

- [STORAGEGATEWAY\_LAST\_BACKUP\_RECOVERY\_POINT\_CREATED](https://docs.aws.amazon.com/config/latest/developerguide/storagegateway-last-backup-recovery-point-created.html)

- [STORAGEGATEWAY\_RESOURCES\_PROTECTED\_BY\_BACKUP\_PLAN](https://docs.aws.amazon.com/config/latest/developerguide/storagegateway-resources-protected-by-backup-plan.html)

- [SUBNET\_AUTO\_ASSIGN\_PUBLIC\_IP\_DISABLED](https://docs.aws.amazon.com/config/latest/developerguide/subnet-auto-assign-public-ip-disabled.html)

- [VIRTUALMACHINE\_LAST\_BACKUP\_RECOVERY\_POINT\_CREATED](https://docs.aws.amazon.com/config/latest/developerguide/virtualmachine-last-backup-recovery-point-created.html)

- [VIRTUALMACHINE\_RESOURCES\_PROTECTED\_BY\_BACKUP\_PLAN](https://docs.aws.amazon.com/config/latest/developerguide/virtualmachine-resources-protected-by-backup-plan.html)

- [VPC\_DEFAULT\_SECURITY\_GROUP\_CLOSED](https://docs.aws.amazon.com/config/latest/developerguide/vpc-default-security-group-closed.html)

- [VPC\_FLOW\_LOGS\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/vpc-flow-logs-enabled.html)

- [VPC\_NETWORK\_ACL\_UNUSED\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/vpc-network-acl-unused-check.html)

- [VPC\_PEERING\_DNS\_RESOLUTION\_CHECK](https://docs.aws.amazon.com/config/latest/developerguide/vpc-peering-dns-resolution-check.html)

- [VPC\_SG\_OPEN\_ONLY\_TO\_AUTHORIZED\_PORTS](https://docs.aws.amazon.com/config/latest/developerguide/vpc-sg-open-only-to-authorized-ports.html)

- [VPC\_VPN\_2\_TUNNELS\_UP](https://docs.aws.amazon.com/config/latest/developerguide/vpc-vpn-2-tunnels-up.html)

- [WAF\_CLASSIC\_LOGGING\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/waf-classic-logging-enabled.html)

- [WAF\_GLOBAL\_RULEGROUP\_NOT\_EMPTY](https://docs.aws.amazon.com/config/latest/developerguide/waf-global-rulegroup-not-empty.html)

- [WAF\_GLOBAL\_RULE\_NOT\_EMPTY](https://docs.aws.amazon.com/config/latest/developerguide/waf-global-rule-not-empty.html)

- [WAF\_GLOBAL\_WEBACL\_NOT\_EMPTY](https://docs.aws.amazon.com/config/latest/developerguide/waf-global-webacl-not-empty.html)

- [WAF\_REGIONAL\_RULEGROUP\_NOT\_EMPTY](https://docs.aws.amazon.com/config/latest/developerguide/waf-regional-rulegroup-not-empty.html)

- [WAF\_REGIONAL\_RULE\_NOT\_EMPTY](https://docs.aws.amazon.com/config/latest/developerguide/waf-regional-rule-not-empty.html)

- [WAF\_REGIONAL\_WEBACL\_NOT\_EMPTY](https://docs.aws.amazon.com/config/latest/developerguide/waf-regional-webacl-not-empty.html)

- [WAFV2\_LOGGING\_ENABLED](https://docs.aws.amazon.com/config/latest/developerguide/wafv2-logging-enabled.html)

- [WAFV2\_RULEGROUP\_NOT\_EMPTY](https://docs.aws.amazon.com/config/latest/developerguide/wafv2-rulegroup-not-empty.html)

- [WAFV2\_WEBACL\_NOT\_EMPTY](https://docs.aws.amazon.com/config/latest/developerguide/wafv2-webacl-not-empty.html)

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

- To find help with issues for this data source type, see [My assessment isn’t collecting compliance check evidence from AWS Config](evidence-collection-issues.md#no-evidence-from-config) and [AWS Config integration issues](https://docs.aws.amazon.com/audit-manager/latest/userguide/control-issues.html#config-rule-integration.title).

- To create a custom control using this data source type, see [Creating a custom control in AWS Audit Manager](https://docs.aws.amazon.com/audit-manager/latest/userguide/create-controls.html).

- To create a custom framework that uses your custom control, see [Creating a custom framework in AWS Audit Manager](https://docs.aws.amazon.com/audit-manager/latest/userguide/custom-frameworks.html).

- To add your custom control to an existing custom framework, see [Editing a custom framework in AWS Audit Manager](https://docs.aws.amazon.com/audit-manager/latest/userguide/edit-custom-frameworks.html).

- To create a custom rule in AWS Config, see [Developing a\
custom rule for AWS Config](https://docs.aws.amazon.com/config/latest/developerguide/evaluate-config_develop-rules.html) in the _AWS Config Developer_
_Guide_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Supported data sources

AWS Security Hub CSPM
