---
title: "Restrict log4j2 message lookup Critical"
---

![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Terraform detectors(58/58)

[Unsecured encryption of SageMaker data at rest](unsecured-sagemaker-data-terraform.md) [Disabled AWS Glue security encryption](disabled-glue-sec-encrypt-terraform.md) [Restrict IAM asterisk action](restrict-iam-asterisk-action-critical-terraform.md) [Disabled AWS RDS Encryption](disabled-rds-encryption-terraform.md) [Exposed secrets in EC2 user data](exposed-ec2-user-data-secret-terraform.md) [Disabled block public acls](disabled-block-public-acls-terraform.md) [Disabled Glue Data Catalog encryption](disabled-glue-cat-encrypt-terraform.md) [S3 bucket restrict public bucket not true](s3-bucket-restrict-public-bucket-terraform.md) [nonhttps viewer protocol policy](nonhttps-viewer-protocol-policy-terraform.md) [Restrict log4j2 message lookup](restrict-log4j2-msg-lookup-terraform.md) [Restrict overly permissive VPC peering routes](restrict-vpc-peering-routes-terraform.md) [Restrict overly permissive access by AWS EKS to all traffic](restrict-eks-traffic-access-terraform.md) [Secure AWS Database Migration Service endpoints](secure-dms-endpoints-terraform.md) [Disabled logging for aws document db](disabled-logging-for-aws-document-db-terraform.md) [Unencrypted code build project](unencrypted-code-build-project-terraform.md) [Sns Topic Uses CMK](sns-topic-uses-cmk-terraform.md) [Enabled RDS public access](enabled-rds-public-access-terraform.md) [Unsecure encryption of DAX at rest](unsecure-encrypt-dax-terraform.md) [Public READ bucket ACL](public-read-bucket-acl-terraform.md) [disabled detailed monitoring for EC2](disabled-detailed-monitoring-for-ec2-terraform.md) [Disabled iam authentication](disabled-iam-authentication-terraform.md) [Unecrypted AWS Redshift using CMK](unencrypted-redshift-cmk-terraform.md) [Implicit SSH for AWS EKS node group](implicit-eks-ssh-access-terraform.md) [Restrict IAM asterisk action](restrict-iam-asterisk-action-terraform.md) [Disabled encryption on Aurora at rest](disabled-aurora-encryption-terraform.md) [Restrict assumed IAM role access](restrict-assumed-role-terraform.md) [Restrict AWS IAM policy with full administrative privileges](restrict-iam-admin-policy-terraform.md) [Restrict actions with any Principal for S3 buckets](restrict-s3-principal-action-terraform.md) [Disabled ALB drops HTTP headers](disabled-alb-drops-http-headers-terraform.md) [Restrict IAM policies with full 'asterisk-asterisk' administrative privileges](restrict-asterisk-iam-admin-policies-terraform.md) [Disabled athena database encryption](disabled-athena-database-encryption-terraform.md) [Use AWS certificate manager SSL certificate with Elastic Load Balancer](use-aws-certificate-ssl-elb-terraform.md) [Unencrypted backup vault](unencrypted-backup-vault-terraform.md) [Avoid hardcoded AWS access keys and secrets credentials](avoid-hardcoded-credentials-terraform.md) [Restrict IAM password reuse](restrict-iam-password-reuse-terraform.md) [Disabled document db encryption](disabled-document-db-encryption-terraform.md) [Configure TLS 1.2 in AWS Load balancer](configure-tls-elb-terraform.md) [Misconfigured data encryption at rest for AWS SageMaker instance](misconfigured-sagemaker-kms-encryption-terraform.md) [Disabled AWS S3 object versioning](disabled-s3-versioning-terraform.md) [Configure HTTPS for CloudFront distribution ViewerProtocolPolicy](conf-https-cloudfront-policy-terraform.md) [Unsecured Encryption in transit for EFS volumes](unsecure-encrypt-efs-terraform.md) [Unencrypted EBS Volumes](unencrypted-ebs-volumes-terraform.md) [Exposed secrets in Lambda function environment variables](exposed-lambda-env-secret-terraform.md) [RDS postgresql file read vulnerability](rds-postgresql-file-read-vulnerability-terraform.md) [Undefined lambda function urls authtype](undefined-lambda-function-urls-authtype-terraform.md) [Associate AWS Glue component with a security component](associate-glue-security-terraform.md) [Restrict public access on DMS replication instance](restrict-public-access-dms-terraform.md) [S3 bucket ignore public acls not true](s3-bucket-ignore-public-acls-not-true-terraform.md) [DynamoDB Table Autoscaling Enabled](dynamodb-table-autoscaling-enabled-terraform.md) [Restrict Neptune cluster instance public access](restrict-neptume-public-access-terraform.md) [Restrict the use of asterisk actions for SQS policy documents](restrict-sqs-asterisk-action-terraform.md) [Disabled Neptune logging](disabled-neptune-logging-terraform.md) [nonhttps load balancer terraform](nonhttps-load-balancer-terraform.md) [Unencryted Codebuild projects](unencryted-codebuild-projects-terraform.md) [Unencrypted Secrets Manager using CMK](unencrypted-secrets-manager-using-cmk-terraform.md) [AWS S3 public WRITE permission](public-write-s3-bucket-terraform.md) [Restrict public IP association on EC2 instance](restrict-public-ip-ec2-terraform.md) [Disabled DynamoDB Point-In-Time Recovery](disabled-dynamodb-pitr-terraform.md)

# Restrict log4j2 message lookup [Critical](severity/critical.md)

Allowance of message lookup in Log4j2 by WAF is detected. Make Sure WAF disallow message lookup in Log4j2.

**Detector ID**

terraform/restrict-log4j2-msg-lookup-terraform@v1.0

**Category**

[Security](categories/security.md)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-117](https://cwe.mitre.org/data/definitions/117.html)

**Tags**

[\# aws-terraform](tags/aws-terraform.md)

* * *

#### Noncompliant example

```terraform
1resource "aws_wafv2_web_acl" "default" {
2  name  = "x-always-block_web_acl"
3  scope = "REGIONAL"
4
5  default_action {
6    allow {}
7  }
8
9  rule {
10    name     = "x-always-block_web_acl_rule"
11    priority = 1
12
13    override_action {
14      none {}
15    }
16    # Noncompliant: `visibility_config` is not defined.
17    statement {
18      rule_group_reference_statement {
19        arn = aws_wafv2_rule_group.x_always_block.arn
20      }
21    }
22    visibility_config {
23      cloudwatch_metrics_enabled = false
24      metric_name                = "friendly-rule-metric-name"
25      sampled_requests_enabled   = false
26    }
27  }
28
29  visibility_config {
30    cloudwatch_metrics_enabled = false
31    metric_name                = ""
32    sampled_requests_enabled   = false
33  }
34}
35resource "aws_wafv2_web_acl_logging_configuration" "example" {
36  log_destination_configs = [aws_kinesis_firehose_delivery_stream.example.arn]
37  resource_arn            = aws_wafv2_web_acl.default.arn
38  redacted_fields {
39    single_header {
40      name = "user-agent"
41    }
42  }
43}

```

#### Compliant example

```terraform
1resource "aws_wafv2_web_acl" "default" {
2  name  = "x-always-block_web_acl"
3  scope = "REGIONAL"
4
5  default_action {
6    allow {}
7  }
8
9  rule {
10    name     = "x-always-block_web_acl_rule"
11    priority = 1
12
13    override_action {
14      none {}
15    }
16    statement {
17      managed_rule_group_statement {
18        name        = "AWSManagedRulesKnownBadInputsRuleSet"
19        vendor_name = "AWS"
20      }
21    }
22    # Compliant: WAF prevents message lookup in Log4j2.
23    visibility_config {
24      cloudwatch_metrics_enabled = false
25      metric_name                = ""
26      sampled_requests_enabled   = false
27    }
28  }
29  visibility_config {
30    cloudwatch_metrics_enabled = false
31    metric_name                = ""
32    sampled_requests_enabled   = false
33  }
34  visibility_config {
35    cloudwatch_metrics_enabled = false
36    metric_name                = "friendly-rule-metric-name"
37    sampled_requests_enabled   = false
38  }
39}
40resource "aws_wafv2_web_acl_logging_configuration" "example" {
41  log_destination_configs = [aws_kinesis_firehose_delivery_stream.example.arn]
42  resource_arn            = aws_wafv2_web_acl.default.arn
43  redacted_fields {
44    single_header {
45      name = "user-agent"
46    }
47  }
48}

```

All content copied from https://docs.aws.amazon.com/.
