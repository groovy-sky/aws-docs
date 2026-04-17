---
title: "SearchCertificates"
---

# SearchCertificates

Retrieves a list of certificates matching search criteria. You can filter certificates by X.509 attributes and ACM specific properties like certificate status, type and renewal eligibility. This operation provides more flexible filtering than [ListCertificates](api-listcertificates.md) by supporting complex filter statements.

## Request Syntax

```nohighlight

{
   "FilterStatement": { ... },
   "MaxResults": number,
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

**[FilterStatement](#API_SearchCertificates_RequestSyntax)**

A filter statement that defines the search criteria. You can combine multiple filters
using AND, OR, and NOT logical operators to create complex queries.

Type: [CertificateFilterStatement](api-certificatefilterstatement.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

**[MaxResults](#API_SearchCertificates_RequestSyntax)**

The maximum number of results to return in the response. Default is 100.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 500.

Required: No

**[NextToken](#API_SearchCertificates_RequestSyntax)**

Use this parameter only when paginating results and only in a subsequent request after
you receive a response with truncated results. Set it to the value of `NextToken`
from the response you just received.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 10000.

Pattern: `[\u0009\u000A\u000D\u0020-\u00FF]*`

Required: No

**[SortBy](#API_SearchCertificates_RequestSyntax)**

Specifies the field to sort results by. Valid values are CREATED\_AT, NOT\_AFTER,
STATUS, RENEWAL\_STATUS, EXPORTED, IN\_USE, NOT\_BEFORE, KEY\_ALGORITHM, TYPE,
CERTIFICATE\_ARN, COMMON\_NAME, REVOKED\_AT, RENEWAL\_ELIGIBILITY, ISSUED\_AT,
MANAGED\_BY, EXPORT\_OPTION, VALIDATION\_METHOD, and IMPORTED\_AT.

Type: String

Valid Values: `CREATED_AT | NOT_AFTER | STATUS | RENEWAL_STATUS | EXPORTED | IN_USE | NOT_BEFORE | KEY_ALGORITHM | TYPE | CERTIFICATE_ARN | COMMON_NAME | REVOKED_AT | RENEWAL_ELIGIBILITY | ISSUED_AT | MANAGED_BY | EXPORT_OPTION | VALIDATION_METHOD | IMPORTED_AT`

Required: No

**[SortOrder](#API_SearchCertificates_RequestSyntax)**

Specifies the order of sorted results. Valid values are ASCENDING or DESCENDING.

Type: String

Valid Values: `ASCENDING | DESCENDING`

Required: No

## Response Syntax

```nohighlight

{
   "NextToken": "string",
   "Results": [
      {
         "CertificateArn": "string",
         "CertificateMetadata": { ... },
         "X509Attributes": {
            "ExtendedKeyUsages": [ "string" ],
            "Issuer": {
               "CommonName": "string",
               "Country": "string",
               "CustomAttributes": [
                  {
                     "ObjectIdentifier": "string",
                     "Value": "string"
                  }
               ],
               "DistinguishedNameQualifier": "string",
               "DomainComponents": [ "string" ],
               "GenerationQualifier": "string",
               "GivenName": "string",
               "Initials": "string",
               "Locality": "string",
               "Organization": "string",
               "OrganizationalUnit": "string",
               "Pseudonym": "string",
               "SerialNumber": "string",
               "State": "string",
               "Surname": "string",
               "Title": "string"
            },
            "KeyAlgorithm": "string",
            "KeyUsages": [ "string" ],
            "NotAfter": number,
            "NotBefore": number,
            "SerialNumber": "string",
            "Subject": {
               "CommonName": "string",
               "Country": "string",
               "CustomAttributes": [
                  {
                     "ObjectIdentifier": "string",
                     "Value": "string"
                  }
               ],
               "DistinguishedNameQualifier": "string",
               "DomainComponents": [ "string" ],
               "GenerationQualifier": "string",
               "GivenName": "string",
               "Initials": "string",
               "Locality": "string",
               "Organization": "string",
               "OrganizationalUnit": "string",
               "Pseudonym": "string",
               "SerialNumber": "string",
               "State": "string",
               "Surname": "string",
               "Title": "string"
            },
            "SubjectAlternativeNames": [
               { ... }
            ]
         }
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[NextToken](#API_SearchCertificates_ResponseSyntax)**

When the list is truncated, this value is present and contains the value to use for
the `NextToken` parameter in a subsequent pagination request.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 10000.

Pattern: `[\u0009\u000A\u000D\u0020-\u00FF]*`

**[Results](#API_SearchCertificates_ResponseSyntax)**

A list of certificate search results containing certificate ARNs, X.509 attributes,
and ACM metadata.

Type: Array of [CertificateSearchResult](api-certificatesearchresult.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You do not have access required to perform this action.

HTTP Status Code: 400

**ThrottlingException**

The request was denied because it exceeded a quota.

**throttlingReasons**

One or more reasons why the request was throttled.

HTTP Status Code: 400

**ValidationException**

The supplied input failed to satisfy constraints of an AWS service.

HTTP Status Code: 400

## Examples

### Search Certificates

The following example searches for exported, issued certificates that are either imported or private, excluding a specific domain name.

#### Sample Request

```

POST / HTTP/1.1
Host: acm.us-east-1.amazonaws.com
Accept-Encoding: identity
Content-Length: 450
X-Amz-Target: CertificateManager.SearchCertificates
X-Amz-Date: 20260213T034622Z
User-Agent: aws-cli/2.0.0 Python/3.9.0 Linux/5.10.0
Content-Type: application/x-amz-json-1.1
Authorization: AWS4-HMAC-SHA256 Credential=key_ID/20260213/us-east-1/acm/aws4_request,
SignedHeaders=content-type;host;x-amz-date;x-amz-target,
Signature=example...

{
    "FilterStatement": {
        "And": [
            {
                "Filter": {
                    "AcmCertificateMetadataFilter": {
                        "Status": "ISSUED"
                    }
                }
            },
            {
                "Or": [
                    {
                        "Filter": {
                            "AcmCertificateMetadataFilter": {
                                "Type": "IMPORTED"
                            }
                        }
                    },
                    {
                        "Filter": {
                            "AcmCertificateMetadataFilter": {
                                "Type": "PRIVATE"
                            }
                        }
                    }
                ]
            },
            {
                "Not": {
                    "Filter": {
                        "X509AttributeFilter": {
                            "SubjectAlternativeName": {
                                "DnsName": {
                                    "Value": "test.com",
                                    "ComparisonOperator": "CONTAINS"
                                }
                            }
                        }
                    }
                }
            },
            {
                "Filter": {
                    "AcmCertificateMetadataFilter": {
                        "Exported": true
                    }
                }
            }
        ]
    },
    "MaxResults": 10,
    "SortBy": "CREATED_AT",
    "SortOrder": "DESCENDING"
}

```

#### Sample Response

```nohighlight

HTTP/1.1 200 OK
x-amzn-RequestId: 12345678-1234-1234-1234-123456789012
Content-Type: application/x-amz-json-1.1
Content-Length: 500
Date: Fri, 13 Feb 2026 03:46:22 GMT
Connection: Keep-alive

{
    "Results": [
        {
            "CertificateArn": "arn:aws:acm:region:account:certificate/12345678-1234-1234-1234-123456789012",
            "X509Attributes": {
                "Issuer": {
                    "CommonName": "Example CA",
                    "Country": "US",
                    "Organization": "Example Corp"
                },
                "Subject": {
                    "CommonName": "www.example.com"
                },
                "ExtendedKeyUsages": [
                    "TLS_WEB_SERVER_AUTHENTICATION"
                ],
                "KeyAlgorithm": "RSA_2048",
                "KeyUsages": [
                    "DIGITAL_SIGNATURE"
                ],
                "SerialNumber": "e5:87:ef:34:7a:4a:0f:de",
                "NotAfter": "2028-12-31T23:59:59+00:00",
                "NotBefore": "2008-01-01T00:00:01+00:00"
            },
            "CertificateMetadata": {
                "AcmCertificateMetadata": {
                    "CreatedAt": "2020-06-15T18:47:09+00:00",
                    "Exported": true,
                    "ImportedAt": "2020-06-15T18:47:09+00:00",
                    "InUse": true,
                    "RenewalEligibility": "INELIGIBLE",
                    "Status": "ISSUED",
                    "Type": "IMPORTED",
                    "ExportOption": "DISABLED"
                }
            }
        }
    ],
    "NextToken": "nextToken"
}

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/acm-2015-12-08/SearchCertificates)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/acm-2015-12-08/SearchCertificates)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/acm-2015-12-08/SearchCertificates)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/acm-2015-12-08/SearchCertificates)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/acm-2015-12-08/SearchCertificates)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/acm-2015-12-08/SearchCertificates)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/acm-2015-12-08/SearchCertificates)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/acm-2015-12-08/SearchCertificates)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/acm-2015-12-08/SearchCertificates)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/acm-2015-12-08/SearchCertificates)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RevokeCertificate

UpdateCertificateOptions

All content copied from https://docs.aws.amazon.com/.
