# Appendix: SOAP API

###### Note

SOAP APIs for Amazon S3 are not available for new customers, and are approaching End of Life (EOL) on October 31, 2025. We recommend that you use either the REST API or the AWS SDKs.

This section describes the SOAP API with respect to service, bucket, and object operations.
Note that SOAP requests, both authenticated and anonymous, must be sent to Amazon S3 using SSL.
Amazon S3 returns an error when you send a SOAP request over HTTP.

The latest Amazon S3 WSDL is available at [docs.aws.amazon.com/2006-03-01/AmazonS3.wsdl](https://doc.s3.amazonaws.com/2006-03-01/AmazonS3.wsdl).

###### Topics

- [Operations on the Service (SOAP API)](https://docs.aws.amazon.com/AmazonS3/latest/API/SOAPServiceOps.html)

- [Operations on Buckets (SOAP API)](https://docs.aws.amazon.com/AmazonS3/latest/API/SOAPBucketsOps.html)

- [Operations on Objects (SOAP API)](https://docs.aws.amazon.com/AmazonS3/latest/API/SOAPObjectsOps.html)

- [Authenticating SOAP requests](https://docs.aws.amazon.com/AmazonS3/latest/API/SOAPAuthentication.html)

- [Setting access policy with SOAP](https://docs.aws.amazon.com/AmazonS3/latest/API/SOAPAccessPolicy.html)

- [Common elements](#SOAPCommon)

- [SOAP Error Responses](https://docs.aws.amazon.com/AmazonS3/latest/API/SOAPErrorResponses.html)

## Common elements

You can include the following authorization-related elements with any SOAP
request:

- `AWSAccessKeyId:` The AWS Access Key ID of the
requester

- `Timestamp:` The current time on your system

- `Signature:` The signature for the request

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Appendix: OPTIONS object

Operations on the Service (SOAP API)
