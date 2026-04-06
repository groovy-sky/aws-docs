# Security in S3 on Outposts

Cloud security at AWS is the highest priority. As an AWS customer, you benefit from
data centers and network architectures that are built to meet the requirements of the most
security-sensitive organizations.

Security is a shared responsibility between AWS and you. The [shared responsibility\
model](https://aws.amazon.com/compliance/shared-responsibility-model) describes this as security _of_ the
cloud and security _in_ the cloud:

- **Security of the cloud** – AWS is
responsible for protecting the infrastructure that runs AWS services in the
AWS Cloud. AWS also provides you with services that you can use securely.
Third-party auditors regularly test and verify the effectiveness of our security as
part of the [AWS Compliance Programs](https://aws.amazon.com/compliance/programs). To learn about the compliance
programs that apply to Amazon S3 on Outposts, see [AWS Services in Scope\
by Compliance Program](https://aws.amazon.com/compliance/services-in-scope).

- **Security in the cloud** – Your responsibility
is determined by the AWS service that you use. You are also responsible for other
factors including the sensitivity of your data, your company's requirements, and
applicable laws and regulations.

This documentation helps you understand how to apply the shared responsibility model when
using S3 on Outposts. The following topics show you how to configure S3 on Outposts to meet your
security and compliance objectives. You also learn how to use other AWS services that help
you to monitor and secure your S3 on Outposts resources.

###### Topics

- [Setting up IAM with S3 on Outposts](s3outpostsiam.md)

- [Data encryption in S3 on Outposts](s3-outposts-data-encryption.md)

- [AWS PrivateLink for S3 on Outposts](https://docs.aws.amazon.com/AmazonS3/latest/s3-outposts/s3-outposts-privatelink-interface-endpoints.html)

- [AWS Signature Version 4 (SigV4) authentication-specific policy keys](https://docs.aws.amazon.com/AmazonS3/latest/s3-outposts/s3-outposts-bucket-policy-s3-sigv4-conditions.html)

- [AWS managed policies for Amazon S3 on Outposts](https://docs.aws.amazon.com/AmazonS3/latest/s3-outposts/s3-outposts-aws-manpol.html)

- [Using service-linked roles for Amazon S3 on Outposts](https://docs.aws.amazon.com/AmazonS3/latest/s3-outposts/S3OutpostsServiceLinkedRoles.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Authorization and authentication caching

Setting up IAM
