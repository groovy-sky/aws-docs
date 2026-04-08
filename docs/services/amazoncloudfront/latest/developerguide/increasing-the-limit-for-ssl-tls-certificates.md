# Increase the quotas for SSL/TLS certificates

There are quotas on the number of SSL/TLS certificates
that you can import into AWS Certificate Manager (ACM) or upload to AWS Identity and Access Management (IAM). There also is a quota on the number of SSL/TLS certificates that
you can use with an AWS account when you configure CloudFront to serve HTTPS requests by
using dedicated IP addresses. However, you can request higher quotas.

###### Topics

- [Increase quota on certificates imported into ACM](#certificates-to-import-into-acm)

- [Increase quota on certificates uploaded to IAM](#certificates-to-upload-into-iam)

- [Increase quota on certificates used with dedicated IP addresses](#certificates-using-dedicated-ip-address)

## Increase quota on certificates imported into ACM

For the quota on the number of certificates that you can import into ACM,
see [Quotas](../../../acm/latest/userguide/acm-limits.md) in the _AWS Certificate Manager User Guide_.

To request a higher quota, use the Service Quotas console. For more information, see [Requesting a quota increase](../../../servicequotas/latest/userguide/request-quota-increase.md) in the _Service Quotas User Guide_.

## Increase quota on certificates uploaded to IAM

For the quota (formerly known as limit) on the number of certificates that you
can upload to IAM, see [IAM\
and STS Limits](../../../iam/latest/userguide/reference-iam-limits.md) in the _IAM User Guide_.

To request a higher quota, use the Service Quotas console. For more information, see [Requesting a quota increase](../../../servicequotas/latest/userguide/request-quota-increase.md) in the _Service Quotas User Guide_.

## Increase quota on certificates used with dedicated IP addresses

For the quota on the number of SSL certificates that
you can use for each AWS account when serving HTTPS requests using dedicated
IP addresses, see [Quotas on SSL certificates](cloudfront-limits.md#limits-ssl-certificates).

To request a higher quota, use the Service Quotas console. For more information, see [Requesting a quota increase](../../../servicequotas/latest/userguide/request-quota-increase.md) in the _Service Quotas User Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Determine the size of the public key in an SSL/TLS RSA certificate

Rotate SSL/TLS certificates

All content copied from https://docs.aws.amazon.com/.
