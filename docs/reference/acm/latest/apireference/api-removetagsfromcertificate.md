# RemoveTagsFromCertificate

Remove one or more tags from an ACM certificate. A tag consists of a key-value pair.
If you do not specify the value portion of the tag when calling this function, the tag
will be removed regardless of value. If you specify a value, the tag is removed only if
it is associated with the specified value.

To add tags to a certificate, use the [AddTagsToCertificate](api-addtagstocertificate.md) action.
To view all of the tags that have been applied to a specific ACM certificate, use the
[ListTagsForCertificate](api-listtagsforcertificate.md) action.

## Request Syntax

```nohighlight

{
   "CertificateArn": "string",
   "Tags": [
      {
         "Key": "string",
         "Value": "string"
      }
   ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[CertificateArn](#API_RemoveTagsFromCertificate_RequestSyntax)**

String that contains the ARN of the ACM Certificate with one or more tags that you
want to remove. This must be of the form:

`arn:aws:acm:region:123456789012:certificate/12345678-1234-1234-1234-123456789012`

For more information about ARNs, see [Amazon Resource Names (ARNs)](../../../../general/general/latest/gr/aws-arns-and-namespaces.md).

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:[\w+=/,.@-]+:acm:[\w+=/,.@-]*:[0-9]+:[\w+=,.@-]+(/[\w+=,.@-]+)*`

Required: Yes

**[Tags](#API_RemoveTagsFromCertificate_RequestSyntax)**

The key-value pair that defines the tag to remove.

Type: Array of [Tag](api-tag.md) objects

Array Members: Minimum number of 1 item. Maximum number of 50 items.

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidArnException**

The requested Amazon Resource Name (ARN) does not refer to an existing
resource.

HTTP Status Code: 400

**InvalidParameterException**

An input parameter was invalid.

HTTP Status Code: 400

**InvalidTagException**

One or both of the values that make up the key-value pair is not valid. For example,
you cannot specify a tag value that begins with `aws:`.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified certificate cannot be found in the caller's account or the caller's
account cannot be found.

HTTP Status Code: 400

**TagPolicyException**

A specified tag did not comply with an existing tag policy and was rejected.

HTTP Status Code: 400

**ThrottlingException**

The request was denied because it exceeded a quota.

**throttlingReasons**

One or more reasons why the request was throttled.

HTTP Status Code: 400

## Examples

### Remove two tags from an ACM certificate

This example illustrates one usage of RemoveTagsFromCertificate.

#### Sample Request

```

POST / HTTP/1.1
Host: acm.us-east-1.amazonaws.com
X-Amz-Target: CertificateManager.RemoveTagsFromCertificate
X-Amz-Date: 20160414T163042Z
User-Agent: aws-cli/1.10.20 Python/2.7.3 Linux/3.13.0-83-generic botocore/1.4.11
Content-Type: application/x-amz-json-1.1
Authorization: AWS4-HMAC-SHA256 Credential=AKIAI44QH8DHBEXAMPLE/20160414/us-east-1/acm/aws4_request,
SignedHeaders=content-type;host;user-agent;x-amz-date;x-amz-target,
Signature=379429306c5e89b9b4be5b35e29c26cc1da38215d8055a5ed0bdda57bcc881cc

{
  "CertificateArn": "arn:aws:acm:us-east-1:111122223333:certificate/12345678-1234-1234-1234-123456789012",
  "Tags": [{
    "Key": "website",
    "Value": "example.com"
  },
  {
    "Key": "stack",
    "Value": "production"
  }]
}

```

#### Sample Response

```

HTTP/1.1 200 OK
x-amzn-RequestId: 3c8d676d-025e-11e6-8823-93164b47113c
Content-Type: application/x-amz-json-1.1
Content-Length: 0
Date: Thu, 14 Apr 2016 16:30:44 GMT

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/acm-2015-12-08/removetagsfromcertificate.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/acm-2015-12-08/removetagsfromcertificate.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/acm-2015-12-08/removetagsfromcertificate.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/acm-2015-12-08/removetagsfromcertificate.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/acm-2015-12-08/removetagsfromcertificate.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/acm-2015-12-08/removetagsfromcertificate.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/acm-2015-12-08/removetagsfromcertificate.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/acm-2015-12-08/removetagsfromcertificate.md)

- [AWS SDK for Python](../../../../services/goto/boto3/acm-2015-12-08/removetagsfromcertificate.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/acm-2015-12-08/removetagsfromcertificate.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

PutAccountConfiguration

RenewCertificate
