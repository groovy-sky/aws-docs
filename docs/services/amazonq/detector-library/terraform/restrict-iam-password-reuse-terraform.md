![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Terraform detectors(58/58)

[Unsecured encryption of SageMaker data at rest](unsecured-sagemaker-data-terraform.md) [Disabled AWS Glue security encryption](disabled-glue-sec-encrypt-terraform.md) [Restrict IAM asterisk action](restrict-iam-asterisk-action-critical-terraform.md) [Disabled AWS RDS Encryption](disabled-rds-encryption-terraform.md) [Exposed secrets in EC2 user data](exposed-ec2-user-data-secret-terraform.md) [Disabled block public acls](disabled-block-public-acls-terraform.md) [Disabled Glue Data Catalog encryption](disabled-glue-cat-encrypt-terraform.md) [S3 bucket restrict public bucket not true](s3-bucket-restrict-public-bucket-terraform.md) [nonhttps viewer protocol policy](nonhttps-viewer-protocol-policy-terraform.md) [Restrict log4j2 message lookup](restrict-log4j2-msg-lookup-terraform.md) [Restrict overly permissive VPC peering routes](restrict-vpc-peering-routes-terraform.md) [Restrict overly permissive access by AWS EKS to all traffic](restrict-eks-traffic-access-terraform.md) [Secure AWS Database Migration Service endpoints](secure-dms-endpoints-terraform.md) [Disabled logging for aws document db](disabled-logging-for-aws-document-db-terraform.md) [Unencrypted code build project](unencrypted-code-build-project-terraform.md) [Sns Topic Uses CMK](sns-topic-uses-cmk-terraform.md) [Enabled RDS public access](enabled-rds-public-access-terraform.md) [Unsecure encryption of DAX at rest](unsecure-encrypt-dax-terraform.md) [Public READ bucket ACL](public-read-bucket-acl-terraform.md) [disabled detailed monitoring for EC2](disabled-detailed-monitoring-for-ec2-terraform.md) [Disabled iam authentication](disabled-iam-authentication-terraform.md) [Unecrypted AWS Redshift using CMK](unencrypted-redshift-cmk-terraform.md) [Implicit SSH for AWS EKS node group](implicit-eks-ssh-access-terraform.md) [Restrict IAM asterisk action](restrict-iam-asterisk-action-terraform.md) [Disabled encryption on Aurora at rest](disabled-aurora-encryption-terraform.md) [Restrict assumed IAM role access](restrict-assumed-role-terraform.md) [Restrict AWS IAM policy with full administrative privileges](restrict-iam-admin-policy-terraform.md) [Restrict actions with any Principal for S3 buckets](restrict-s3-principal-action-terraform.md) [Disabled ALB drops HTTP headers](disabled-alb-drops-http-headers-terraform.md) [Restrict IAM policies with full 'asterisk-asterisk' administrative privileges](restrict-asterisk-iam-admin-policies-terraform.md) [Disabled athena database encryption](disabled-athena-database-encryption-terraform.md) [Use AWS certificate manager SSL certificate with Elastic Load Balancer](use-aws-certificate-ssl-elb-terraform.md) [Unencrypted backup vault](unencrypted-backup-vault-terraform.md) [Avoid hardcoded AWS access keys and secrets credentials](avoid-hardcoded-credentials-terraform.md) [Restrict IAM password reuse](https://docs.aws.amazon.com/amazonq/detector-library/terraform/restrict-iam-password-reuse-terraform) [Disabled document db encryption](https://docs.aws.amazon.com/amazonq/detector-library/terraform/disabled-document-db-encryption-terraform) [Configure TLS 1.2 in AWS Load balancer](https://docs.aws.amazon.com/amazonq/detector-library/terraform/configure-tls-elb-terraform) [Misconfigured data encryption at rest for AWS SageMaker instance](https://docs.aws.amazon.com/amazonq/detector-library/terraform/misconfigured-sagemaker-kms-encryption-terraform) [Disabled AWS S3 object versioning](https://docs.aws.amazon.com/amazonq/detector-library/terraform/disabled-s3-versioning-terraform) [Configure HTTPS for CloudFront distribution ViewerProtocolPolicy](https://docs.aws.amazon.com/amazonq/detector-library/terraform/conf-https-cloudfront-policy-terraform) [Unsecured Encryption in transit for EFS volumes](https://docs.aws.amazon.com/amazonq/detector-library/terraform/unsecure-encrypt-efs-terraform) [Unencrypted EBS Volumes](https://docs.aws.amazon.com/amazonq/detector-library/terraform/unencrypted-ebs-volumes-terraform) [Exposed secrets in Lambda function environment variables](https://docs.aws.amazon.com/amazonq/detector-library/terraform/exposed-lambda-env-secret-terraform) [RDS postgresql file read vulnerability](https://docs.aws.amazon.com/amazonq/detector-library/terraform/rds-postgresql-file-read-vulnerability-terraform) [Undefined lambda function urls authtype](https://docs.aws.amazon.com/amazonq/detector-library/terraform/undefined-lambda-function-urls-authtype-terraform) [Associate AWS Glue component with a security component](https://docs.aws.amazon.com/amazonq/detector-library/terraform/associate-glue-security-terraform) [Restrict public access on DMS replication instance](https://docs.aws.amazon.com/amazonq/detector-library/terraform/restrict-public-access-dms-terraform) [S3 bucket ignore public acls not true](https://docs.aws.amazon.com/amazonq/detector-library/terraform/s3-bucket-ignore-public-acls-not-true-terraform) [DynamoDB Table Autoscaling Enabled](https://docs.aws.amazon.com/amazonq/detector-library/terraform/dynamodb-table-autoscaling-enabled-terraform) [Restrict Neptune cluster instance public access](https://docs.aws.amazon.com/amazonq/detector-library/terraform/restrict-neptume-public-access-terraform) [Restrict the use of asterisk actions for SQS policy documents](https://docs.aws.amazon.com/amazonq/detector-library/terraform/restrict-sqs-asterisk-action-terraform) [Disabled Neptune logging](https://docs.aws.amazon.com/amazonq/detector-library/terraform/disabled-neptune-logging-terraform) [nonhttps load balancer terraform](https://docs.aws.amazon.com/amazonq/detector-library/terraform/nonhttps-load-balancer-terraform) [Unencryted Codebuild projects](https://docs.aws.amazon.com/amazonq/detector-library/terraform/unencryted-codebuild-projects-terraform) [Unencrypted Secrets Manager using CMK](https://docs.aws.amazon.com/amazonq/detector-library/terraform/unencrypted-secrets-manager-using-cmk-terraform) [AWS S3 public WRITE permission](https://docs.aws.amazon.com/amazonq/detector-library/terraform/public-write-s3-bucket-terraform) [Restrict public IP association on EC2 instance](https://docs.aws.amazon.com/amazonq/detector-library/terraform/restrict-public-ip-ec2-terraform) [Disabled DynamoDB Point-In-Time Recovery](https://docs.aws.amazon.com/amazonq/detector-library/terraform/disabled-dynamodb-pitr-terraform)

# Restrict IAM password reuse [High](https://docs.aws.amazon.com/amazonq/detector-library/terraform/severity/high)

The AWS IAM password policy permits the reuse of password. Make sure AWS IAM password policy restrict password reuse.

**Detector ID**

terraform/restrict-iam-password-reuse-terraform@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/terraform/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-521](https://cwe.mitre.org/data/definitions/521.html)

**Tags**

[\# aws-terraform](https://docs.aws.amazon.com/amazonq/detector-library/terraform/tags/aws-terraform)

* * *

#### Noncompliant example

```terraform
1resource "aws_iam_account_password_policy" "pike" {
2  allow_users_to_change_password = false
3  hard_expiry                    = true
4  max_password_age               = 90
5  minimum_password_length        = 14
6  # Noncompliant: `password_reuse_prevention` is less than 24.
7  password_reuse_prevention      = 10
8  require_lowercase_characters   = true
9  require_numbers                = true
10  require_symbols                = true
11  require_uppercase_characters   = true
12}

```

#### Compliant example

```terraform
1resource "aws_iam_account_password_policy" "pike" {
2  allow_users_to_change_password = true
3  hard_expiry                    = true
4  max_password_age               = 90
5  minimum_password_length        = 14
6  # Compliant: `password_reuse_prevention` is set to 24.
7  password_reuse_prevention      = 24
8  require_lowercase_characters   = true
9  require_numbers                = true
10  require_symbols                = true
11  require_uppercase_characters   = true
12}

```
