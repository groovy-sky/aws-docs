# AssociateDelegationSignerToDomain

Creates a delegation signer (DS) record in the registry zone for this domain
name.

Note that creating DS record at the registry impacts DNSSEC validation of your DNS
records. This action may render your domain name unavailable on the internet if the
steps are completed in the wrong order, or with incorrect timing. For more information
about DNSSEC signing, see [Configuring DNSSEC\
signing](../../../../services/route53/latest/developerguide/dns-configuring-dnssec.md) in the _Route 53 developer_
_guide_.

## Request Syntax

```nohighlight

{
   "DomainName": "string",
   "SigningAttributes": {
      "Algorithm": number,
      "Flags": number,
      "PublicKey": "string"
   }
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[DomainName](#API_domains_AssociateDelegationSignerToDomain_RequestSyntax)**

The name of the domain.

Type: String

Length Constraints: Maximum length of 255.

Required: Yes

**[SigningAttributes](#API_domains_AssociateDelegationSignerToDomain_RequestSyntax)**

The information about a key, including the algorithm, public key-value, and
flags.

Type: [DnssecSigningAttributes](api-domains-dnssecsigningattributes.md) object

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

**[OperationId](#API_domains_AssociateDelegationSignerToDomain_ResponseSyntax)**

The identifier for tracking the progress of the request. To query the operation
status, use [GetOperationDetail](api-domains-getoperationdetail.md).

Type: String

Length Constraints: Maximum length of 255.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DnssecLimitExceeded**

This error is returned if you call `AssociateDelegationSignerToDomain`
when the specified domain has reached the maximum number of DS records. You can't add
any additional DS records unless you delete an existing one first.

HTTP Status Code: 400

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

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53domains-2014-05-15/associatedelegationsignertodomain.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53domains-2014-05-15/associatedelegationsignertodomain.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53domains-2014-05-15/associatedelegationsignertodomain.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53domains-2014-05-15/associatedelegationsignertodomain.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53domains-2014-05-15/associatedelegationsignertodomain.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53domains-2014-05-15/associatedelegationsignertodomain.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53domains-2014-05-15/associatedelegationsignertodomain.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53domains-2014-05-15/associatedelegationsignertodomain.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53domains-2014-05-15/associatedelegationsignertodomain.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53domains-2014-05-15/associatedelegationsignertodomain.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AcceptDomainTransferFromAnotherAwsAccount

CancelDomainTransferToAnotherAwsAccount
