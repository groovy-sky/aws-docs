![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Terraform detectors(58/58)

[Unsecured encryption of SageMaker data at rest](unsecured-sagemaker-data-terraform.md) [Disabled AWS Glue security encryption](disabled-glue-sec-encrypt-terraform.md) [Restrict IAM asterisk action](restrict-iam-asterisk-action-critical-terraform.md) [Disabled AWS RDS Encryption](disabled-rds-encryption-terraform.md) [Exposed secrets in EC2 user data](exposed-ec2-user-data-secret-terraform.md) [Disabled block public acls](disabled-block-public-acls-terraform.md) [Disabled Glue Data Catalog encryption](disabled-glue-cat-encrypt-terraform.md) [S3 bucket restrict public bucket not true](s3-bucket-restrict-public-bucket-terraform.md) [nonhttps viewer protocol policy](nonhttps-viewer-protocol-policy-terraform.md) [Restrict log4j2 message lookup](restrict-log4j2-msg-lookup-terraform.md) [Restrict overly permissive VPC peering routes](restrict-vpc-peering-routes-terraform.md) [Restrict overly permissive access by AWS EKS to all traffic](restrict-eks-traffic-access-terraform.md) [Secure AWS Database Migration Service endpoints](secure-dms-endpoints-terraform.md) [Disabled logging for aws document db](disabled-logging-for-aws-document-db-terraform.md) [Unencrypted code build project](unencrypted-code-build-project-terraform.md) [Sns Topic Uses CMK](sns-topic-uses-cmk-terraform.md) [Enabled RDS public access](enabled-rds-public-access-terraform.md) [Unsecure encryption of DAX at rest](unsecure-encrypt-dax-terraform.md) [Public READ bucket ACL](public-read-bucket-acl-terraform.md) [disabled detailed monitoring for EC2](disabled-detailed-monitoring-for-ec2-terraform.md) [Disabled iam authentication](disabled-iam-authentication-terraform.md) [Unecrypted AWS Redshift using CMK](unencrypted-redshift-cmk-terraform.md) [Implicit SSH for AWS EKS node group](implicit-eks-ssh-access-terraform.md) [Restrict IAM asterisk action](restrict-iam-asterisk-action-terraform.md) [Disabled encryption on Aurora at rest](disabled-aurora-encryption-terraform.md) [Restrict assumed IAM role access](restrict-assumed-role-terraform.md) [Restrict AWS IAM policy with full administrative privileges](restrict-iam-admin-policy-terraform.md) [Restrict actions with any Principal for S3 buckets](restrict-s3-principal-action-terraform.md) [Disabled ALB drops HTTP headers](disabled-alb-drops-http-headers-terraform.md) [Restrict IAM policies with full 'asterisk-asterisk' administrative privileges](restrict-asterisk-iam-admin-policies-terraform.md) [Disabled athena database encryption](disabled-athena-database-encryption-terraform.md) [Use AWS certificate manager SSL certificate with Elastic Load Balancer](use-aws-certificate-ssl-elb-terraform.md) [Unencrypted backup vault](unencrypted-backup-vault-terraform.md) [Avoid hardcoded AWS access keys and secrets credentials](avoid-hardcoded-credentials-terraform.md) [Restrict IAM password reuse](restrict-iam-password-reuse-terraform.md) [Disabled document db encryption](disabled-document-db-encryption-terraform.md) [Configure TLS 1.2 in AWS Load balancer](configure-tls-elb-terraform.md) [Misconfigured data encryption at rest for AWS SageMaker instance](misconfigured-sagemaker-kms-encryption-terraform.md) [Disabled AWS S3 object versioning](disabled-s3-versioning-terraform.md) [Configure HTTPS for CloudFront distribution ViewerProtocolPolicy](conf-https-cloudfront-policy-terraform.md) [Unsecured Encryption in transit for EFS volumes](unsecure-encrypt-efs-terraform.md) [Unencrypted EBS Volumes](unencrypted-ebs-volumes-terraform.md) [Exposed secrets in Lambda function environment variables](https://docs.aws.amazon.com/amazonq/detector-library/terraform/exposed-lambda-env-secret-terraform) [RDS postgresql file read vulnerability](https://docs.aws.amazon.com/amazonq/detector-library/terraform/rds-postgresql-file-read-vulnerability-terraform) [Undefined lambda function urls authtype](https://docs.aws.amazon.com/amazonq/detector-library/terraform/undefined-lambda-function-urls-authtype-terraform) [Associate AWS Glue component with a security component](https://docs.aws.amazon.com/amazonq/detector-library/terraform/associate-glue-security-terraform) [Restrict public access on DMS replication instance](https://docs.aws.amazon.com/amazonq/detector-library/terraform/restrict-public-access-dms-terraform) [S3 bucket ignore public acls not true](https://docs.aws.amazon.com/amazonq/detector-library/terraform/s3-bucket-ignore-public-acls-not-true-terraform) [DynamoDB Table Autoscaling Enabled](https://docs.aws.amazon.com/amazonq/detector-library/terraform/dynamodb-table-autoscaling-enabled-terraform) [Restrict Neptune cluster instance public access](https://docs.aws.amazon.com/amazonq/detector-library/terraform/restrict-neptume-public-access-terraform) [Restrict the use of asterisk actions for SQS policy documents](https://docs.aws.amazon.com/amazonq/detector-library/terraform/restrict-sqs-asterisk-action-terraform) [Disabled Neptune logging](https://docs.aws.amazon.com/amazonq/detector-library/terraform/disabled-neptune-logging-terraform) [nonhttps load balancer terraform](https://docs.aws.amazon.com/amazonq/detector-library/terraform/nonhttps-load-balancer-terraform) [Unencryted Codebuild projects](https://docs.aws.amazon.com/amazonq/detector-library/terraform/unencryted-codebuild-projects-terraform) [Unencrypted Secrets Manager using CMK](https://docs.aws.amazon.com/amazonq/detector-library/terraform/unencrypted-secrets-manager-using-cmk-terraform) [AWS S3 public WRITE permission](https://docs.aws.amazon.com/amazonq/detector-library/terraform/public-write-s3-bucket-terraform) [Restrict public IP association on EC2 instance](https://docs.aws.amazon.com/amazonq/detector-library/terraform/restrict-public-ip-ec2-terraform) [Disabled DynamoDB Point-In-Time Recovery](https://docs.aws.amazon.com/amazonq/detector-library/terraform/disabled-dynamodb-pitr-terraform)

# Exposed secrets in Lambda function environment variables [High](https://docs.aws.amazon.com/amazonq/detector-library/terraform/severity/high)

The exposure of secrets through Lambda function's environment variables is detected. Make sure that secrets are not exposed by environment variables of Lambda function.

**Detector ID**

terraform/exposed-lambda-env-secret-terraform@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/terraform/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-200](https://cwe.mitre.org/data/definitions/200.html)

**Tags**

[\# aws-terraform](https://docs.aws.amazon.com/amazonq/detector-library/terraform/tags/aws-terraform)

* * *

#### Noncompliant example

```terraform
1resource "aws_lambda_function" "LambdaFunction" {
2  function_name = "IdempotencyFunction"
3  role          = aws_iam_role.IdempotencyFunctionRole.arn
4  runtime       = "python3.11"
5  handler       = "app.lambda_handler"
6  filename      = "lambda.zip"
7  # Noncompliant: Hard-coded secrets exist in lambda environment.
8  environment {
9    variables = {
10      AWS_ACCESS_KEY_ID     = "AKIAIOSFODNN7EXAMPLE",
11      AWS_SECRET_ACCESS_KEY = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
12      AWS_DEFAULT_REGION    = "us-west-2"
13    }
14  }
15  kms_key_arn = aws_kms_key.anyoldguff.arn
16  reserved_concurrent_executions = 100
17  code_signing_config_arn = "123123123"
18  tracing_config {
19    mode = "Active"
20  }
21  vpc_config {
22    subnet_ids         = [aws_subnet.subnet_for_lambda.id]
23    security_group_ids = [aws_security_group.sg_for_lambda.id]
24  }
25  dead_letter_config {
26    target_arn = "test"
27  }
28}

```

#### Compliant example

```terraform
1resource "aws_lambda_function" "IdempotencyFunction" {
2  function_name = "IdempotencyFunction"
3  role          = aws_iam_role.IdempotencyFunctionRole.arn
4  runtime       = "python3.11"
5  handler       = "app.lambda_handler"
6  # Compliant: Hard-coded secrets does not exist in lambda environment.
7  filename      = "lambda.zip"
8  environment {}
9  kms_key_arn = aws_kms_key.anyoldguff.arn
10  code_signing_config_arn = "123123123"
11  reserved_concurrent_executions = 100
12  tracing_config {
13    mode = "Active"
14  }
15  vpc_config {
16    subnet_ids         = [aws_subnet.subnet_for_lambda.id]
17    security_group_ids = [aws_security_group.sg_for_lambda.id]
18  }
19  dead_letter_config {
20    target_arn = "test"
21  }
22}

```
