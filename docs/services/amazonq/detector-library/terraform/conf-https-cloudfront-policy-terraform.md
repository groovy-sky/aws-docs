![AWS logo](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/aws-logo.506636a4.svg)

[Amazon QDetector Library](../../detector-library.md) [Sign in to Amazon Q](https://console.aws.amazon.com/amazonq/home)

#### Amazon Q

###### Detector Library

#### Terraform detectors(58/58)

[Unsecured encryption of SageMaker data at rest](unsecured-sagemaker-data-terraform.md) [Disabled AWS Glue security encryption](disabled-glue-sec-encrypt-terraform.md) [Restrict IAM asterisk action](restrict-iam-asterisk-action-critical-terraform.md) [Disabled AWS RDS Encryption](disabled-rds-encryption-terraform.md) [Exposed secrets in EC2 user data](exposed-ec2-user-data-secret-terraform.md) [Disabled block public acls](disabled-block-public-acls-terraform.md) [Disabled Glue Data Catalog encryption](disabled-glue-cat-encrypt-terraform.md) [S3 bucket restrict public bucket not true](s3-bucket-restrict-public-bucket-terraform.md) [nonhttps viewer protocol policy](nonhttps-viewer-protocol-policy-terraform.md) [Restrict log4j2 message lookup](restrict-log4j2-msg-lookup-terraform.md) [Restrict overly permissive VPC peering routes](restrict-vpc-peering-routes-terraform.md) [Restrict overly permissive access by AWS EKS to all traffic](restrict-eks-traffic-access-terraform.md) [Secure AWS Database Migration Service endpoints](secure-dms-endpoints-terraform.md) [Disabled logging for aws document db](disabled-logging-for-aws-document-db-terraform.md) [Unencrypted code build project](unencrypted-code-build-project-terraform.md) [Sns Topic Uses CMK](sns-topic-uses-cmk-terraform.md) [Enabled RDS public access](enabled-rds-public-access-terraform.md) [Unsecure encryption of DAX at rest](unsecure-encrypt-dax-terraform.md) [Public READ bucket ACL](public-read-bucket-acl-terraform.md) [disabled detailed monitoring for EC2](disabled-detailed-monitoring-for-ec2-terraform.md) [Disabled iam authentication](disabled-iam-authentication-terraform.md) [Unecrypted AWS Redshift using CMK](unencrypted-redshift-cmk-terraform.md) [Implicit SSH for AWS EKS node group](implicit-eks-ssh-access-terraform.md) [Restrict IAM asterisk action](restrict-iam-asterisk-action-terraform.md) [Disabled encryption on Aurora at rest](disabled-aurora-encryption-terraform.md) [Restrict assumed IAM role access](restrict-assumed-role-terraform.md) [Restrict AWS IAM policy with full administrative privileges](restrict-iam-admin-policy-terraform.md) [Restrict actions with any Principal for S3 buckets](restrict-s3-principal-action-terraform.md) [Disabled ALB drops HTTP headers](disabled-alb-drops-http-headers-terraform.md) [Restrict IAM policies with full 'asterisk-asterisk' administrative privileges](restrict-asterisk-iam-admin-policies-terraform.md) [Disabled athena database encryption](disabled-athena-database-encryption-terraform.md) [Use AWS certificate manager SSL certificate with Elastic Load Balancer](use-aws-certificate-ssl-elb-terraform.md) [Unencrypted backup vault](unencrypted-backup-vault-terraform.md) [Avoid hardcoded AWS access keys and secrets credentials](avoid-hardcoded-credentials-terraform.md) [Restrict IAM password reuse](restrict-iam-password-reuse-terraform.md) [Disabled document db encryption](disabled-document-db-encryption-terraform.md) [Configure TLS 1.2 in AWS Load balancer](configure-tls-elb-terraform.md) [Misconfigured data encryption at rest for AWS SageMaker instance](misconfigured-sagemaker-kms-encryption-terraform.md) [Disabled AWS S3 object versioning](disabled-s3-versioning-terraform.md) [Configure HTTPS for CloudFront distribution ViewerProtocolPolicy](https://docs.aws.amazon.com/amazonq/detector-library/terraform/conf-https-cloudfront-policy-terraform) [Unsecured Encryption in transit for EFS volumes](https://docs.aws.amazon.com/amazonq/detector-library/terraform/unsecure-encrypt-efs-terraform) [Unencrypted EBS Volumes](https://docs.aws.amazon.com/amazonq/detector-library/terraform/unencrypted-ebs-volumes-terraform) [Exposed secrets in Lambda function environment variables](https://docs.aws.amazon.com/amazonq/detector-library/terraform/exposed-lambda-env-secret-terraform) [RDS postgresql file read vulnerability](https://docs.aws.amazon.com/amazonq/detector-library/terraform/rds-postgresql-file-read-vulnerability-terraform) [Undefined lambda function urls authtype](https://docs.aws.amazon.com/amazonq/detector-library/terraform/undefined-lambda-function-urls-authtype-terraform) [Associate AWS Glue component with a security component](https://docs.aws.amazon.com/amazonq/detector-library/terraform/associate-glue-security-terraform) [Restrict public access on DMS replication instance](https://docs.aws.amazon.com/amazonq/detector-library/terraform/restrict-public-access-dms-terraform) [S3 bucket ignore public acls not true](https://docs.aws.amazon.com/amazonq/detector-library/terraform/s3-bucket-ignore-public-acls-not-true-terraform) [DynamoDB Table Autoscaling Enabled](https://docs.aws.amazon.com/amazonq/detector-library/terraform/dynamodb-table-autoscaling-enabled-terraform) [Restrict Neptune cluster instance public access](https://docs.aws.amazon.com/amazonq/detector-library/terraform/restrict-neptume-public-access-terraform) [Restrict the use of asterisk actions for SQS policy documents](https://docs.aws.amazon.com/amazonq/detector-library/terraform/restrict-sqs-asterisk-action-terraform) [Disabled Neptune logging](https://docs.aws.amazon.com/amazonq/detector-library/terraform/disabled-neptune-logging-terraform) [nonhttps load balancer terraform](https://docs.aws.amazon.com/amazonq/detector-library/terraform/nonhttps-load-balancer-terraform) [Unencryted Codebuild projects](https://docs.aws.amazon.com/amazonq/detector-library/terraform/unencryted-codebuild-projects-terraform) [Unencrypted Secrets Manager using CMK](https://docs.aws.amazon.com/amazonq/detector-library/terraform/unencrypted-secrets-manager-using-cmk-terraform) [AWS S3 public WRITE permission](https://docs.aws.amazon.com/amazonq/detector-library/terraform/public-write-s3-bucket-terraform) [Restrict public IP association on EC2 instance](https://docs.aws.amazon.com/amazonq/detector-library/terraform/restrict-public-ip-ec2-terraform) [Disabled DynamoDB Point-In-Time Recovery](https://docs.aws.amazon.com/amazonq/detector-library/terraform/disabled-dynamodb-pitr-terraform)

# Configure HTTPS for CloudFront distribution ViewerProtocolPolicy [High](https://docs.aws.amazon.com/amazonq/detector-library/terraform/severity/high)

HTTPS is not configured in the ViewerProtocolPolicy of CloudFront distribution. Make sure that CloudFront distribution ViewerProtocolPolicy is configured to HTTPS

**Detector ID**

terraform/conf-https-cloudfront-policy-terraform@v1.0

**Category**

[Security](https://docs.aws.amazon.com/amazonq/detector-library/terraform/categories/security)

**Common Weakness Enumeration (CWE)![external icon](https://docs.aws.amazon.com/amazonq/detector-library/_next/static/media/external.cdeb5c57.svg)**

[CWE-319](https://cwe.mitre.org/data/definitions/319.html)

**Tags**

[\# aws-terraform](https://docs.aws.amazon.com/amazonq/detector-library/terraform/tags/aws-terraform)

* * *

#### Noncompliant example

```terraform
1resource "aws_cloudfront_distribution" "cf" {
2  origin_group {
3    origin_id = "groupS3"
4
5    failover_criteria {
6      status_codes = [403, 404, 500, 502]
7    }
8
9    member {
10      origin_id = "primaryS3"
11    }
12
13    member {
14      origin_id = "failoverS3"
15    }
16  }
17  origin {
18    domain_name = aws_s3_bucket.statics.bucket_regional_domain_name
19    origin_id   = local.s3_origin_id
20  }
21
22  enabled             = true
23  is_ipv6_enabled     = true
24  default_root_object = "index.html"
25  price_class         = "PriceClass_100"
26
27  aliases = [data.aws_route53_zone.selected.name]
28  default_cache_behavior {
29    allowed_methods        = []
30    cached_methods         = []
31    target_origin_id       = ""
32    # Noncompliant: `viewer_protocol_policy` is set to `allow-all`.
33    viewer_protocol_policy = "allow-all"
34  }
35  restrictions {
36    geo_restriction {
37      restriction_type = ""
38    }
39  }
40  viewer_certificate {
41    acm_certificate_arn = "aws_acm_certificate_arn"
42    minimum_protocol_version = "TLSv1.2_2018"
43  }
44  default_cache_behavior {
45    response_headers_policy_id = data.aws_cloudfront_response_headers_policy.simple_cors.id
46  }
47  web_acl_id = aws_wafv2_web_acl.waf_cloudfront
48  logging_config {
49    include_cookies = false
50    bucket          = "mylogs.s3.amazonaws.com"
51    prefix          = "myprefix"
52  }
53  origin {
54    domain_name = aws_s3_bucket.primary.bucket_regional_domain_name
55    origin_id   = "primaryS3"
56
57    s3_origin_config {
58      origin_access_identity = aws_cloudfront_origin_access_identity.default.cloudfront_access_identity_path
59    }
60  }
61}
62resource "aws_wafv2_web_acl_logging_configuration" "pike" {
63  resource_arn            = aws_wafv2_web_acl.waf_cloudfront.arn
64  log_destination_configs = ["arn:aws:logs:eu-west-2:680235478471:log-group:pike/"]
65}
66resource "aws_wafv2_web_acl" "waf_cloudfront" {
67  name        = "managed-rule-example"
68  description = "Example of a managed rule."
69  scope       = "REGIONAL"
70
71  default_action {
72    allow {}
73  }
74
75  rule {
76    name     = "rule-1"
77    priority = 1
78
79    override_action {
80      none {}
81    }
82
83    statement {
84      managed_rule_group_statement {
85        name        = "AWSManagedRulesAnonymousIpList"
86        vendor_name = "AWS"
87
88        excluded_rule {
89          name = "rule-1"
90        }
91
92        scope_down_statement {
93          geo_match_statement {
94            country_codes = ["US", "NL"]
95          }
96        }
97      }
98    }
99
100    visibility_config {
101      cloudwatch_metrics_enabled = false
102      metric_name                = "friendly-rule-metric-name"
103      sampled_requests_enabled   = false
104    }
105  }
106
107  rule {
108    name     = "rule-2"
109    priority = 2
110
111    override_action {
112      none {}
113    }
114
115    statement {
116      managed_rule_group_statement {
117        name        = "AWSManagedRulesKnownBadInputsRuleSet"
118        vendor_name = "AWS"
119
120        excluded_rule {
121          name = "SizeRestrictions_QUERYSTRING"
122        }
123
124        scope_down_statement {
125          geo_match_statement {
126            country_codes = ["US", "NL"]
127          }
128        }
129      }
130    }
131
132    visibility_config {
133      cloudwatch_metrics_enabled = false
134      metric_name                = "friendly-rule-metric-name"
135      sampled_requests_enabled   = false
136    }
137  }
138
139
140  tags = {
141    Tag1 = "Value1"
142    Tag2 = "Value2"
143  }
144
145  visibility_config {
146    cloudwatch_metrics_enabled = false
147    metric_name                = "friendly-metric-name"
148    sampled_requests_enabled   = false
149  }
150}
151data "aws_cloudfront_response_headers_policy" "simple_cors" {
152  name = "SimpleCORS"
153}

```

#### Compliant example

```terraform
1resource "aws_cloudfront_distribution" "cf" {
2  origin_group {
3    origin_id = "groupS3"
4
5    failover_criteria {
6      status_codes = [403, 404, 500, 502]
7    }
8
9    member {
10      origin_id = "primaryS3"
11    }
12
13    member {
14      origin_id = "failoverS3"
15    }
16  }
17  origin {
18    domain_name = aws_s3_bucket.statics.bucket_regional_domain_name
19    origin_id   = local.s3_origin_id
20  }
21
22  enabled             = true
23  is_ipv6_enabled     = true
24  default_root_object = "index.html"
25  price_class         = "PriceClass_100"
26
27  aliases = [data.aws_route53_zone.selected.name]
28  default_cache_behavior {
29    allowed_methods        = []
30    cached_methods         = []
31    target_origin_id       = ""
32    # Compliant: `viewer_protocol_policy` is set to `redirect-to-https`.
33    viewer_protocol_policy = "redirect-to-https"
34  }
35  restrictions {
36    geo_restriction {
37      restriction_type = ""
38    }
39  }
40  viewer_certificate {
41    acm_certificate_arn = "aws_acm_certificate_arn"
42    minimum_protocol_version = "TLSv1.2_2018"
43  }
44  default_cache_behavior {
45    response_headers_policy_id = data.aws_cloudfront_response_headers_policy.simple_cors.id
46  }
47  web_acl_id = aws_wafv2_web_acl.waf_cloudfront
48  logging_config {
49    include_cookies = false
50    bucket          = "mylogs.s3.amazonaws.com"
51    prefix          = "myprefix"
52  }
53  origin {
54    domain_name = aws_s3_bucket.primary.bucket_regional_domain_name
55    origin_id   = "primaryS3"
56
57    s3_origin_config {
58      origin_access_identity = aws_cloudfront_origin_access_identity.default.cloudfront_access_identity_path
59    }
60  }
61}
62resource "aws_wafv2_web_acl_logging_configuration" "pike" {
63  resource_arn            = aws_wafv2_web_acl.waf_cloudfront.arn
64  log_destination_configs = ["arn:aws:logs:eu-west-2:680235478471:log-group:pike/"]
65}
66resource "aws_wafv2_web_acl" "waf_cloudfront" {
67  name        = "managed-rule-example"
68  description = "Example of a managed rule."
69  scope       = "REGIONAL"
70
71  default_action {
72    allow {}
73  }
74
75  rule {
76    name     = "rule-1"
77    priority = 1
78
79    override_action {
80      none {}
81    }
82
83    statement {
84      managed_rule_group_statement {
85        name        = "AWSManagedRulesAnonymousIpList"
86        vendor_name = "AWS"
87
88        excluded_rule {
89          name = "rule-1"
90        }
91
92        scope_down_statement {
93          geo_match_statement {
94            country_codes = ["US", "NL"]
95          }
96        }
97      }
98    }
99
100    visibility_config {
101      cloudwatch_metrics_enabled = false
102      metric_name                = "friendly-rule-metric-name"
103      sampled_requests_enabled   = false
104    }
105  }
106
107  rule {
108    name     = "rule-2"
109    priority = 2
110
111    override_action {
112      none {}
113    }
114
115    statement {
116      managed_rule_group_statement {
117        name        = "AWSManagedRulesKnownBadInputsRuleSet"
118        vendor_name = "AWS"
119
120        excluded_rule {
121          name = "SizeRestrictions_QUERYSTRING"
122        }
123
124        scope_down_statement {
125          geo_match_statement {
126            country_codes = ["US", "NL"]
127          }
128        }
129      }
130    }
131
132    visibility_config {
133      cloudwatch_metrics_enabled = false
134      metric_name                = "friendly-rule-metric-name"
135      sampled_requests_enabled   = false
136    }
137  }
138
139
140  tags = {
141    Tag1 = "Value1"
142    Tag2 = "Value2"
143  }
144
145  visibility_config {
146    cloudwatch_metrics_enabled = false
147    metric_name                = "friendly-metric-name"
148    sampled_requests_enabled   = false
149  }
150}
151data "aws_cloudfront_response_headers_policy" "simple_cors" {
152  name = "SimpleCORS"
153}

```
