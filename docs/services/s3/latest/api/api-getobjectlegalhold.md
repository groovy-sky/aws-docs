# GetObjectLegalHold

###### Note

This operation is not supported for directory buckets.

Gets an object's current legal hold status. For more information, see [Locking Objects](../dev/object-lock.md).

This functionality is not supported for Amazon S3 on Outposts.

The following action is related to `GetObjectLegalHold`:

- [GetObjectAttributes](api-getobjectattributes.md)

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

## Request Syntax

```nohighlight

GET /{Key+}?legal-hold&versionId=VersionId HTTP/1.1
Host: Bucket.s3.amazonaws.com
x-amz-request-payer: RequestPayer
x-amz-expected-bucket-owner: ExpectedBucketOwner

```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_GetObjectLegalHold_RequestSyntax)**

The bucket name containing the object whose legal hold status you want to retrieve.

**Access points** \- When you use this action with an access point for general purpose buckets, you must provide the alias of the access point in place of the bucket name or specify the access point ARN. When you use this action with an access point for directory buckets, you must provide the access point name in place of the bucket name. When using the access point ARN, you must direct requests to the access point hostname. The access point hostname takes the form _AccessPointName_- _AccountId_.s3-accesspoint. _Region_.amazonaws.com. When using this action with an access point through the AWS SDKs, you provide the access point ARN in place of the bucket name. For more information about access point ARNs, see [Using access points](../userguide/using-access-points.md) in the _Amazon S3 User Guide_.

Required: Yes

**[Key](#API_GetObjectLegalHold_RequestSyntax)**

The key name for the object whose legal hold status you want to retrieve.

Length Constraints: Minimum length of 1.

Required: Yes

**[versionId](#API_GetObjectLegalHold_RequestSyntax)**

The version ID of the object whose legal hold status you want to retrieve.

**[x-amz-expected-bucket-owner](#API_GetObjectLegalHold_RequestSyntax)**

The account ID of the expected bucket owner. If the account ID that you provide does not match the actual owner of the bucket, the request fails with the HTTP status code `403 Forbidden` (access denied).

**[x-amz-request-payer](#API_GetObjectLegalHold_RequestSyntax)**

Confirms that the requester knows that they will be charged for the request. Bucket owners need not
specify this parameter in their requests. If either the source or destination S3 bucket has Requester
Pays enabled, the requester will pay for the corresponding charges. For information about
downloading objects from Requester Pays buckets, see [Downloading Objects in Requester Pays\
Buckets](../dev/objectsinrequesterpaysbuckets.md) in the _Amazon S3 User Guide_.

###### Note

This functionality is not supported for directory buckets.

Valid Values: `requester`

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<LegalHold>
   <Status>string</Status>
</LegalHold>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[LegalHold](#API_GetObjectLegalHold_ResponseSyntax)**

Root level tag for the LegalHold parameters.

Required: Yes

**[Status](#API_GetObjectLegalHold_ResponseSyntax)**

Indicates whether the specified object has a legal hold in place.

Type: String

Valid Values: `ON | OFF`

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3-2006-03-01/getobjectlegalhold.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3-2006-03-01/getobjectlegalhold.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/getobjectlegalhold.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3-2006-03-01/getobjectlegalhold.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/getobjectlegalhold.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3-2006-03-01/getobjectlegalhold.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3-2006-03-01/getobjectlegalhold.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3-2006-03-01/getobjectlegalhold.md)

- [AWS SDK for Python](../../../goto/boto3/s3-2006-03-01/getobjectlegalhold.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/getobjectlegalhold.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetObjectAttributes

GetObjectLockConfiguration

All content copied from https://docs.aws.amazon.com/.
