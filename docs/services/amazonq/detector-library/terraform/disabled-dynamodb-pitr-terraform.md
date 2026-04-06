![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Terraform detectors(58/58)

[Unsecured encryption of SageMaker data at rest](unsecured-sagemaker-data-terraform.md) [Disabled AWS Glue security encryption](disabled-glue-sec-encrypt-terraform.md) [Restrict IAM asterisk action](restrict-iam-asterisk-action-critical-terraform.md) [Disabled AWS RDS Encryption](disabled-rds-encryption-terraform.md) [Exposed secrets in EC2 user data](exposed-ec2-user-data-secret-terraform.md) [Disabled block public acls](disabled-block-public-acls-terraform.md) [Disabled Glue Data Catalog encryption](disabled-glue-cat-encrypt-terraform.md) [S3 bucket restrict public bucket not true](s3-bucket-restrict-public-bucket-terraform.md) [nonhttps viewer protocol policy](nonhttps-viewer-protocol-policy-terraform.md) [Restrict log4j2 message lookup](restrict-log4j2-msg-lookup-terraform.md) [Restrict overly permissive VPC peering routes](restrict-vpc-peering-routes-terraform.md) [Restrict overly permissive access by AWS EKS to all traffic](restrict-eks-traffic-access-terraform.md) [Secure AWS Database Migration Service endpoints](secure-dms-endpoints-terraform.md) [Disabled logging for aws document db](disabled-logging-for-aws-document-db-terraform.md) [Unencrypted code build project](unencrypted-code-build-project-terraform.md) [Sns Topic Uses CMK](sns-topic-uses-cmk-terraform.md) [Enabled RDS public access](enabled-rds-public-access-terraform.md) [Unsecure encryption of DAX at rest](unsecure-encrypt-dax-terraform.md) [Public READ bucket ACL](public-read-bucket-acl-terraform.md) [disabled detailed monitoring for EC2](disabled-detailed-monitoring-for-ec2-terraform.md) [Disabled iam authentication](disabled-iam-authentication-terraform.md) [Unecrypted AWS Redshift using CMK](unencrypted-redshift-cmk-terraform.md) [Implicit SSH for AWS EKS node group](implicit-eks-ssh-access-terraform.md) [Restrict IAM asterisk action](restrict-iam-asterisk-action-terraform.md) [Disabled encryption on Aurora at rest](disabled-aurora-encryption-terraform.md) [Restrict assumed IAM role access](restrict-assumed-role-terraform.md) [Restrict AWS IAM policy with full administrative privileges](restrict-iam-admin-policy-terraform.md) [Restrict actions with any Principal for S3 buckets](restrict-s3-principal-action-terraform.md) [Disabled ALB drops HTTP headers](disabled-alb-drops-http-headers-terraform.md) [Restrict IAM policies with full 'asterisk-asterisk' administrative privileges](restrict-asterisk-iam-admin-policies-terraform.md) [Disabled athena database encryption](disabled-athena-database-encryption-terraform.md) [Use AWS certificate manager SSL certificate with Elastic Load Balancer](use-aws-certificate-ssl-elb-terraform.md) [Unencrypted backup vault](unencrypted-backup-vault-terraform.md) [Avoid hardcoded AWS access keys and secrets credentials](avoid-hardcoded-credentials-terraform.md) [Restrict IAM password reuse](restrict-iam-password-reuse-terraform.md) [Disabled document db encryption](disabled-document-db-encryption-terraform.md) [Configure TLS 1.2 in AWS Load balancer](configure-tls-elb-terraform.md) [Misconfigured data encryption at rest for AWS SageMaker instance](misconfigured-sagemaker-kms-encryption-terraform.md) [Disabled AWS S3 object versioning](disabled-s3-versioning-terraform.md) [Configure HTTPS for CloudFront distribution ViewerProtocolPolicy](conf-https-cloudfront-policy-terraform.md) [Unsecured Encryption in transit for EFS volumes](unsecure-encrypt-efs-terraform.md) [Unencrypted EBS Volumes](unencrypted-ebs-volumes-terraform.md) [Exposed secrets in Lambda function environment variables](exposed-lambda-env-secret-terraform.md) [RDS postgresql file read vulnerability](rds-postgresql-file-read-vulnerability-terraform.md) [Undefined lambda function urls authtype](undefined-lambda-function-urls-authtype-terraform.md) [Associate AWS Glue component with a security component](associate-glue-security-terraform.md) [Restrict public access on DMS replication instance](restrict-public-access-dms-terraform.md) [S3 bucket ignore public acls not true](s3-bucket-ignore-public-acls-not-true-terraform.md) [DynamoDB Table Autoscaling Enabled](dynamodb-table-autoscaling-enabled-terraform.md) [Restrict Neptune cluster instance public access](restrict-neptume-public-access-terraform.md) [Restrict the use of asterisk actions for SQS policy documents](restrict-sqs-asterisk-action-terraform.md) [Disabled Neptune logging](disabled-neptune-logging-terraform.md) [nonhttps load balancer terraform](nonhttps-load-balancer-terraform.md) [Unencryted Codebuild projects](unencryted-codebuild-projects-terraform.md) [Unencrypted Secrets Manager using CMK](unencrypted-secrets-manager-using-cmk-terraform.md) [AWS S3 public WRITE permission](public-write-s3-bucket-terraform.md) [Restrict public IP association on EC2 instance](restrict-public-ip-ec2-terraform.md) [Disabled DynamoDB Point-In-Time Recovery](https://docs.aws.amazon.com/amazonq/detector-library/terraform/disabled-dynamodb-pitr-terraform)

# Disabled DynamoDB Point-In-Time Recovery [High](https://docs.aws.amazon.com/amazonq/detector-library/terraform/severity/high)

Disabled DynamoDB Point-In-Time Recovery is detected. Make sure that DynamoDB Point-In-Time Recovery is enabled.

**Detector ID**

terraform/disabled-dynamodb-pitr-terraform@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/terraform/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-221](https://cwe.mitre.org/data/definitions/221.html)

**Tags**

[\# aws-terraform](https://docs.aws.amazon.com/amazonq/detector-library/terraform/tags/aws-terraform)

* * *

#### Noncompliant example

```terraform
1resource "aws_dynamodb_table" "basic-dynamodb-table" {
2  name           = "GameScores"
3  # Noncompliant: `point_in_time_recovery` is disabled.
4  attribute {
5    name = "UserId"
6    type = "S"
7  }
8  ttl {
9    attribute_name = "TimeToExist"
10    enabled        = false
11  }
12  tags = {
13    Name        = "dynamodb-table-1"
14    Environment = "production"
15  }
16  enabled = true
17  server_side_encryption {
18    enabled = true
19    kms_key_arn = "arn:aws:kms:us-west-2:123456789012:key/1234abcd-12ab-34cd-56ef-1234567890ab"
20  }
21}
22resource "aws_appautoscaling_target" "pass" {
23  resource_id        = "table/${aws_dynamodb_table.basic-dynamodb-table.name}"
24  scalable_dimension = "dynamodb:table:ReadCapacityUnits"
25  service_namespace  = "dynamodb"
26  min_capacity       = 1
27  max_capacity       = 15
28}
29
30resource "aws_appautoscaling_policy" "pass" {
31  name               = "rcu-auto-scaling"
32  service_namespace  = aws_appautoscaling_target.pass.service_namespace
33  scalable_dimension = aws_appautoscaling_target.pass.scalable_dimension
34  resource_id        = aws_appautoscaling_target.pass.resource_id
35  policy_type        = "TargetTrackingScaling"
36
37  target_tracking_scaling_policy_configuration {
38    predefined_metric_specification {
39      predefined_metric_type = "DynamoDBReadCapacityUtilization"
40    }
41
42    target_value       = 75
43    scale_in_cooldown  = 300
44    scale_out_cooldown = 300
45  }
46}

```

#### Compliant example

```terraform
1resource "aws_dynamodb_table" "basic-dynamodb-table" {
2  name           = "GameScores"
3  attribute {
4    name = "UserId"
5    type = "S"
6  }
7  ttl {
8    attribute_name = "TimeToExist"
9    enabled        = false
10  }
11  # Compliant: `point_in_time_recovery` is enabled.
12  point_in_time_recovery {
13    enabled = true
14  }
15  tags = {
16    Name        = "dynamodb-table-1"
17    Environment = "production"
18  }
19  enabled = true
20  server_side_encryption {
21    enabled = true
22    kms_key_arn = "arn:aws:kms:us-west-2:123456789012:key/1234abcd-12ab-34cd-56ef-1234567890ab"
23  }
24}
25resource "aws_appautoscaling_target" "pass" {
26  resource_id        = "table/${aws_dynamodb_table.basic-dynamodb-table.name}"
27  scalable_dimension = "dynamodb:table:ReadCapacityUnits"
28  service_namespace  = "dynamodb"
29  min_capacity       = 1
30  max_capacity       = 15
31}
32
33resource "aws_appautoscaling_policy" "pass" {
34  name               = "rcu-auto-scaling"
35  service_namespace  = aws_appautoscaling_target.pass.service_namespace
36  scalable_dimension = aws_appautoscaling_target.pass.scalable_dimension
37  resource_id        = aws_appautoscaling_target.pass.resource_id
38  policy_type        = "TargetTrackingScaling"
39
40  target_tracking_scaling_policy_configuration {
41    predefined_metric_specification {
42      predefined_metric_type = "DynamoDBReadCapacityUtilization"
43    }
44
45    target_value       = 75
46    scale_in_cooldown  = 300
47    scale_out_cooldown = 300
48  }
49}

```
