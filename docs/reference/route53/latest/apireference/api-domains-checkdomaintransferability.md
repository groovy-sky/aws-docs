# CheckDomainTransferability

Checks whether a domain name can be transferred to Amazon Route 53.

## Request Syntax

```nohighlight

{
   "AuthCode": "string",
   "DomainName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[AuthCode](#API_domains_CheckDomainTransferability_RequestSyntax)**

If the registrar for the top-level domain (TLD) requires an authorization code to
transfer the domain, the code that you got from the current registrar for the
domain.

Type: String

Length Constraints: Maximum length of 1024.

Required: No

**[DomainName](#API_domains_CheckDomainTransferability_RequestSyntax)**

The name of the domain that you want to transfer to Route 53. The top-level domain
(TLD), such as .com, must be a TLD that Route 53 supports. For a list of supported TLDs,
see [Domains that You Can\
Register with Amazon Route 53](../../../../services/route53/latest/developerguide/registrar-tld-list.md) in the _Amazon Route 53 Developer_
_Guide_.

The domain name can contain only the following characters:

- Letters a through z. Domain names are not case sensitive.

- Numbers 0 through 9.

- Hyphen (-). You can't specify a hyphen at the beginning or end of a label.

- Period (.) to separate the labels in the name, such as the `.` in
`example.com`.

Type: String

Length Constraints: Maximum length of 255.

Required: Yes

## Response Syntax

```nohighlight

{
   "Message": "string",
   "Transferability": {
      "Transferable": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Message](#API_domains_CheckDomainTransferability_ResponseSyntax)**

Provides an explanation for when a domain can't be transferred.

Type: String

**[Transferability](#API_domains_CheckDomainTransferability_ResponseSyntax)**

A complex type that contains information about whether the specified domain can be
transferred to Route 53.

Type: [DomainTransferability](https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_DomainTransferability.html) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidInput**

The requested item is not acceptable. For example, for APIs that accept a domain name,
the request might specify a domain name that doesn't belong to the account that
submitted the request. For `AcceptDomainTransferFromAnotherAwsAccount`, the
password might be invalid.

**message**

The requested item is not acceptable. For example, for an OperationId it might refer
to the ID of an operation that is already completed. For a domain name, it might not be
a valid domain name or belong to the requester account.

HTTP Status Code: 400

**UnsupportedTLD**

Amazon Route 53 does not support this top-level domain (TLD).

**message**

Amazon Route 53 does not support this top-level domain (TLD).

HTTP Status Code: 400

## Examples

### CheckDomainTransferability Example

This example illustrates one usage of CheckDomainTransferability.

#### Sample Request

```

POST / HTTP/1.1
host:route53domains.us-east-1.amazonaws.com
x-amz-date:20140711T205225Z
authorization:AWS4-HMAC-SHA256
              Credential=AKIAIOSFODNN7EXAMPLE/20140711/us-east-1/route53domains/aws4_request,
              SignedHeaders=content-length;content-type;host;user-agent;x-amz-date;x-amz-target,
              Signature=[calculated-signature]
x-amz-target:Route53Domains_v20140515.CheckDomainTransferability
user-agent:aws-sdk-java/1.8.3 Linux/2.6.18-164.el5PAE Java_HotSpot (TM )_Server_VM/24.60-b09/1.7.0_60
content-type:application/x-amz-json-1.1
content-length:[number of characters in the JSON string]
connections:Keep-Alive
{
   "DomainName": "example.com",
   "AuthCode": "T92XJ38"
}
```

#### Sample Response

```

HTTP/1.1 200
Content-Length:[number of characters in the JSON string]
{
   "Transferability":
      {"Transferable":"TRANSFERABLE"}
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/route53domains-2014-05-15/CheckDomainTransferability)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/route53domains-2014-05-15/CheckDomainTransferability)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53domains-2014-05-15/CheckDomainTransferability)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/route53domains-2014-05-15/CheckDomainTransferability)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53domains-2014-05-15/CheckDomainTransferability)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/route53domains-2014-05-15/CheckDomainTransferability)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/route53domains-2014-05-15/CheckDomainTransferability)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/route53domains-2014-05-15/CheckDomainTransferability)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/route53domains-2014-05-15/CheckDomainTransferability)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53domains-2014-05-15/CheckDomainTransferability)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CheckDomainAvailability

DeleteDomain
