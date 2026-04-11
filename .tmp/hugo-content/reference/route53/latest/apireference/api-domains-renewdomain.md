# RenewDomain

This operation renews a domain for the specified number of years. The cost of renewing
your domain is billed to your AWS account.

We recommend that you renew your domain several weeks before the expiration date. Some
TLD registries delete domains before the expiration date if you haven't renewed far
enough in advance. For more information about renewing domain registration, see [Renewing\
Registration for a Domain](../../../../services/route53/latest/developerguide/domain-renew.md) in the _Amazon Route 53 Developer_
_Guide_.

## Request Syntax

```nohighlight

{
   "CurrentExpiryYear": number,
   "DomainName": "string",
   "DurationInYears": number
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[CurrentExpiryYear](#API_domains_RenewDomain_RequestSyntax)**

The year when the registration for the domain is set to expire. This value must match
the current expiration date for the domain.

Type: Integer

Required: Yes

**[DomainName](#API_domains_RenewDomain_RequestSyntax)**

The name of the domain that you want to renew.

Type: String

Length Constraints: Maximum length of 255.

Required: Yes

**[DurationInYears](#API_domains_RenewDomain_RequestSyntax)**

The number of years that you want to renew the domain for. The maximum number of years
depends on the top-level domain. For the range of valid values for your domain, see
[Domains that You Can\
Register with Amazon Route 53](../../../../services/route53/latest/developerguide/registrar-tld-list.md) in the _Amazon Route 53 Developer_
_Guide_.

Default: 1

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 10.

Required: No

## Response Syntax

```nohighlight

{
   "OperationId": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[OperationId](#API_domains_RenewDomain_ResponseSyntax)**

Identifier for tracking the progress of the request. To query the operation status,
use [GetOperationDetail](api-domains-getoperationdetail.md).

Type: String

Length Constraints: Maximum length of 255.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DuplicateRequest**

The request is already in progress for the domain.

**message**

The request is already in progress for the domain.

**requestId**

ID of the request operation.

HTTP Status Code: 400

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

**OperationLimitExceeded**

The number of operations or jobs running exceeded the allowed threshold for the
account.

**message**

The number of operations or jobs running exceeded the allowed threshold for the
account.

HTTP Status Code: 400

**TLDRulesViolation**

The top-level domain does not support this operation.

**message**

The top-level domain does not support this operation.

HTTP Status Code: 400

**UnsupportedTLD**

Amazon Route 53 does not support this top-level domain (TLD).

**message**

Amazon Route 53 does not support this top-level domain (TLD).

HTTP Status Code: 400

## Examples

### RenewDomain Example

This example illustrates one usage of RenewDomain.

#### Sample Request

```

POST / HTTP/1.1
host:route53domains.us-east-1.amazonaws.com
x-amz-date:20140711T205230Z
authorization:AWS4-HMAC-SHA256
              Credential=AKIAIOSFODNN7EXAMPLE/20140711/us-east-1/route53domains/aws4_request,
              SignedHeaders=content-length;content-type;host;user-agent;x-amz-date;x-amz-target,
              Signature=[calculated-signature]
              x-amz-target:Route53Domains_v20140515.RenewDomain
              user-agent:aws-sdk-java/1.8.3 Linux/2.6.18-164.el5PAE Java_HotSpot (TM )_Server_VM/24.60-b09/1.7.0_60
content-type:application/x-amz-json-1.1
content-length:[number of characters in the JSON string]
{
   "DomainName":"example.com",
   "DurationInYears":"6",
   "CurrentExpiryYear":"2017"
}
```

#### Sample Response

```

HTTP/1.1 200
Content-Length:[number of characters in the JSON string]{}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53domains-2014-05-15/renewdomain.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53domains-2014-05-15/renewdomain.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53domains-2014-05-15/renewdomain.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53domains-2014-05-15/renewdomain.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53domains-2014-05-15/renewdomain.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53domains-2014-05-15/renewdomain.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53domains-2014-05-15/renewdomain.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53domains-2014-05-15/renewdomain.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53domains-2014-05-15/renewdomain.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53domains-2014-05-15/renewdomain.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RejectDomainTransferFromAnotherAwsAccount

ResendContactReachabilityEmail
