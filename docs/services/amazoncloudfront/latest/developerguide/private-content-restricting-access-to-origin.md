# Restrict access to an AWS origin

You can configure CloudFront and some AWS origins in a way that provides the following
benefits:

- Restricts access to the AWS origin so that it's not publicly accessible.

- Makes sure that viewers (users) can access the content in the AWS origin only
through the specified CloudFront distribution. This prevents viewers from accessing the
content directly from the origin, or through an unintended CloudFront distribution.

To do this, configure CloudFront to send authenticated requests to your AWS origin, and
configure the AWS origin to only allow access to authenticated requests from CloudFront. For
more information, see following topics for compatible types of AWS origins.

###### Topics

- [Restrict access to an AWS Elemental MediaPackage v2 origin](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-mediapackage.html)

- [Restrict access to an AWS Elemental MediaStore origin](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-mediastore.html)

- [Restrict access to an AWS Lambda function URL origin](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-lambda.html)

- [Restrict access to an Amazon S3 origin](private-content-restricting-access-to-s3.md)

- [Restrict access with VPC origins](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-vpc-origins.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Create a URL signature using Java

Restrict access to an AWS Elemental MediaPackage v2 origin
