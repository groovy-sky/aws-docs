# CancelDomainTransferToAnotherAwsAccount

Cancels the transfer of a domain from the current AWS account to
another AWS account. You initiate a transfer betweenAWS accounts using [TransferDomainToAnotherAwsAccount](api-domains-transferdomaintoanotherawsaccount.md).

###### Important

You must cancel the transfer before the other AWS account accepts
the transfer using [AcceptDomainTransferFromAnotherAwsAccount](api-domains-acceptdomaintransferfromanotherawsaccount.md).

Use either [ListOperations](api-domains-listoperations.md) or [GetOperationDetail](api-domains-getoperationdetail.md) to determine whether the operation succeeded. [GetOperationDetail](api-domains-getoperationdetail.md) provides additional information, for example,
`Domain Transfer from Aws Account 111122223333 has been cancelled`.

## Request Syntax

```nohighlight

{
   "DomainName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[DomainName](#API_domains_CancelDomainTransferToAnotherAwsAccount_RequestSyntax)**

The name of the domain for which you want to cancel the transfer to another AWS account.

Type: String

Length Constraints: Maximum length of 255.

Required: Yes

## Response Syntax

```nohighlight

{
   "OperationId": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[OperationId](#API_domains_CancelDomainTransferToAnotherAwsAccount_ResponseSyntax)**

The identifier that `TransferDomainToAnotherAwsAccount` returned to track
the progress of the request. Because the transfer request was canceled, the value is no
longer valid, and you can't use `GetOperationDetail` to query the operation
status.

Type: String

Length Constraints: Maximum length of 255.

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

**OperationLimitExceeded**

The number of operations or jobs running exceeded the allowed threshold for the
account.

**message**

The number of operations or jobs running exceeded the allowed threshold for the
account.

HTTP Status Code: 400

**UnsupportedTLD**

Amazon Route 53 does not support this top-level domain (TLD).

**message**

Amazon Route 53 does not support this top-level domain (TLD).

HTTP Status Code: 400

## Examples

### CancelDomainTransferToAnotherAwsAccount Example

This example illustrates one usage of CancelDomainTransferToAnotherAwsAccount.

#### Sample Request

```

POST / HTTP/1.1
host:route53domains.us-east-1.amazonaws.com
x-amz-date:20200311T205230Z
authorization:AWS4-HMAC-SHA256
              Credential=AKIAIOSFODNN7EXAMPLE/20140711/us-east-1/route53domains/aws4_request,
              SignedHeaders=content-length;content-type;host;user-agent;x-amz-date;x-amz-target,
              Signature=[calculated-signature]
x-amz-target:Route53Domains_v20140515.CancelDomainTransferToAnotherAwsAccount
user-agent:aws-sdk-java/1.8.3 Linux/2.6.18-164.el5PAE Java_HotSpot (TM )_Server_VM/24.60-b09/1.7.0_60
content-type:application/x-amz-json-1.1
content-length:[number of characters in the JSON string]
{
    "DomainName":"example.com"
}
```

#### Sample Response

```

HTTP/1.1 200
Content-Length:[number of characters in the JSON string]
{
    "OperationId":"308c56712-faa4-40fe-94c8-b4230example"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53domains-2014-05-15/canceldomaintransfertoanotherawsaccount.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53domains-2014-05-15/canceldomaintransfertoanotherawsaccount.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53domains-2014-05-15/canceldomaintransfertoanotherawsaccount.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53domains-2014-05-15/canceldomaintransfertoanotherawsaccount.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53domains-2014-05-15/canceldomaintransfertoanotherawsaccount.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53domains-2014-05-15/canceldomaintransfertoanotherawsaccount.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53domains-2014-05-15/canceldomaintransfertoanotherawsaccount.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53domains-2014-05-15/canceldomaintransfertoanotherawsaccount.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53domains-2014-05-15/canceldomaintransfertoanotherawsaccount.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53domains-2014-05-15/canceldomaintransfertoanotherawsaccount.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AssociateDelegationSignerToDomain

CheckDomainAvailability
