# ListCertificates

Retrieves a list of certificate ARNs and domain names. You can request that only
certificates that match a specific status be listed. You can also filter by specific
attributes of the certificate. Default filtering returns only `RSA_2048`
certificates. For more information, see [Filters](api-filters.md).

## Request Syntax

```nohighlight

{
   "CertificateStatuses": [ "string" ],
   "Includes": {
      "exportOption": "string",
      "extendedKeyUsage": [ "string" ],
      "keyTypes": [ "string" ],
      "keyUsage": [ "string" ],
      "managedBy": "string"
   },
   "MaxItems": number,
   "NextToken": "string",
   "SortBy": "string",
   "SortOrder": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[CertificateStatuses](#API_ListCertificates_RequestSyntax)**

Filter the certificate list by status value.

Type: Array of strings

Valid Values: `PENDING_VALIDATION | ISSUED | INACTIVE | EXPIRED | VALIDATION_TIMED_OUT | REVOKED | FAILED`

Required: No

**[Includes](#API_ListCertificates_RequestSyntax)**

Filter the certificate list. For more information, see the [Filters](api-filters.md)
structure.

Type: [Filters](api-filters.md) object

Required: No

**[MaxItems](#API_ListCertificates_RequestSyntax)**

Use this parameter when paginating results to specify the maximum number of items to
return in the response. If additional items exist beyond the number you specify, the
`NextToken` element is sent in the response. Use this
`NextToken` value in a subsequent request to retrieve additional
items.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 1000.

Required: No

**[NextToken](#API_ListCertificates_RequestSyntax)**

Use this parameter only when paginating results and only in a subsequent request after
you receive a response with truncated results. Set it to the value of
`NextToken` from the response you just received.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 10000.

Pattern: `[\u0009\u000A\u000D\u0020-\u00FF]*`

Required: No

**[SortBy](#API_ListCertificates_RequestSyntax)**

Specifies the field to sort results by. If you specify `SortBy`, you must
also specify `SortOrder`.

Type: String

Valid Values: `CREATED_AT`

Required: No

**[SortOrder](#API_ListCertificates_RequestSyntax)**

Specifies the order of sorted results. If you specify `SortOrder`, you must
also specify `SortBy`.

Type: String

Valid Values: `ASCENDING | DESCENDING`

Required: No

## Response Syntax

```nohighlight

{
   "CertificateSummaryList": [
      {
         "CertificateArn": "string",
         "CreatedAt": number,
         "DomainName": "string",
         "Exported": boolean,
         "ExportOption": "string",
         "ExtendedKeyUsages": [ "string" ],
         "HasAdditionalSubjectAlternativeNames": boolean,
         "ImportedAt": number,
         "InUse": boolean,
         "IssuedAt": number,
         "KeyAlgorithm": "string",
         "KeyUsages": [ "string" ],
         "ManagedBy": "string",
         "NotAfter": number,
         "NotBefore": number,
         "RenewalEligibility": "string",
         "RevokedAt": number,
         "Status": "string",
         "SubjectAlternativeNameSummaries": [ "string" ],
         "Type": "string"
      }
   ],
   "NextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[CertificateSummaryList](#API_ListCertificates_ResponseSyntax)**

A list of ACM certificates.

Type: Array of [CertificateSummary](api-certificatesummary.md) objects

**[NextToken](#API_ListCertificates_ResponseSyntax)**

When the list is truncated, this value is present and contains the value to use for
the `NextToken` parameter in a subsequent pagination request.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 10000.

Pattern: `[\u0009\u000A\u000D\u0020-\u00FF]*`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidArgsException**

One or more of request parameters specified is not valid.

HTTP Status Code: 400

**ValidationException**

The supplied input failed to satisfy constraints of an AWS service.

HTTP Status Code: 400

## Examples

### List Certificates

The following example lists certificates that you can use to create digital
signatures and to sign code.

#### Sample Request

```

POST / HTTP/1.1
Host: acm.us-east-1.amazonaws.com
Accept-Encoding: identity
Content-Length: 129
X-Amz-Target: CertificateManager.ListCertificates
X-Amz-Date: 20171118T204928Z
User-Agent: aws-cli/1.11.132 Python/2.7.9 Windows/8 botocore/1.5.95
Content-Type: application/x-amz-json-1.1
Authorization: AWS4-HMAC-SHA256 Credential=key_ID/20171118/us-east-1/acm/aws4_request,
SignedHeaders=content-type;host;x-amz-date;x-amz-target,
Signature=49a54...

{
	"MaxItems": 10,
	"Includes": {
		"keyUsage": ["DIGITAL_SIGNATURE"],
		"keyTypes": ["RSA_2048"],
		"extendedKeyUsage": ["CODE_SIGNING"]
	}
}

```

#### Sample Response

```nohighlight

HTTP/1.1 200 OK
x-amzn-RequestId: fa8ffa7f-cca1-11e7-80db-736b2201613a
Content-Type: application/x-amz-json-1.1
Content-Length: 164
Date: Sat, 18 Nov 2017 20:49:32 GMT
Connection: Keep-alive

{"CertificateSummaryList": [
  {
    "CertificateArn": "arn:aws:acm:region:account:certificate/12345678-1234-1234-1234-123456789012",
    "DomainName": "www.example.com"
  },
  {
    "CertificateArn": "arn:aws:acm:region:account:certificate/12345678-1234-1234-1234-123456789012",
    "DomainName": "www.corp.net"
  }
]
}

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/acm-2015-12-08/listcertificates.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/acm-2015-12-08/listcertificates.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/acm-2015-12-08/listcertificates.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/acm-2015-12-08/listcertificates.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/acm-2015-12-08/listcertificates.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/acm-2015-12-08/listcertificates.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/acm-2015-12-08/listcertificates.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/acm-2015-12-08/listcertificates.md)

- [AWS SDK for Python](../../../../services/goto/boto3/acm-2015-12-08/listcertificates.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/acm-2015-12-08/listcertificates.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ImportCertificate

ListTagsForCertificate
