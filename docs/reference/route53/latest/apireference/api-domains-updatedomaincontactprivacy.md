# UpdateDomainContactPrivacy

This operation updates the specified domain contact's privacy setting. When privacy
protection is enabled, your contact information is replaced with contact information for
the registrar or with the phrase "REDACTED FOR PRIVACY", or "On behalf of <domain
name> owner."

###### Note

While some domains may allow different privacy settings per contact, we recommend
specifying the same privacy setting for all contacts.

This operation affects only the contact information for the specified contact type
(administrative, registrant, or technical). If the request succeeds, Amazon Route 53
returns an operation ID that you can use with [GetOperationDetail](api-domains-getoperationdetail.md) to track the progress and completion of the action. If
the request doesn't complete successfully, the domain registrant will be notified by
email.

###### Important

By disabling the privacy service via API, you consent to the publication of the
contact information provided for this domain via the public WHOIS database. You
certify that you are the registrant of this domain name and have the authority to
make this decision. You may withdraw your consent at any time by enabling privacy
protection using either `UpdateDomainContactPrivacy` or the Route 53
console. Enabling privacy protection removes the contact information provided for
this domain from the WHOIS database. For more information on our privacy practices,
see [https://aws.amazon.com/privacy/](https://aws.amazon.com/privacy).

## Request Syntax

```nohighlight

{
   "AdminPrivacy": boolean,
   "BillingPrivacy": boolean,
   "DomainName": "string",
   "RegistrantPrivacy": boolean,
   "TechPrivacy": boolean
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[AdminPrivacy](#API_domains_UpdateDomainContactPrivacy_RequestSyntax)**

Whether you want to conceal contact information from WHOIS queries. If you specify
`true`, WHOIS ("who is") queries return contact information either for
Amazon Registrar or for our registrar associate,
Gandi. If you specify `false`, WHOIS queries return the
information that you entered for the admin contact.

###### Note

You must specify the same privacy setting for the administrative, billing, registrant, and
technical contacts.

Type: Boolean

Required: No

**[BillingPrivacy](#API_domains_UpdateDomainContactPrivacy_RequestSyntax)**

Whether you want to conceal contact information from WHOIS queries. If you specify
`true`, WHOIS ("who is") queries return contact information either for
Amazon Registrar or for our registrar associate,
Gandi. If you specify `false`, WHOIS queries return the
information that you entered for the billing contact.

###### Note

You must specify the same privacy setting for the administrative, billing, registrant, and
technical contacts.

Type: Boolean

Required: No

**[DomainName](#API_domains_UpdateDomainContactPrivacy_RequestSyntax)**

The name of the domain that you want to update the privacy setting for.

Type: String

Length Constraints: Maximum length of 255.

Required: Yes

**[RegistrantPrivacy](#API_domains_UpdateDomainContactPrivacy_RequestSyntax)**

Whether you want to conceal contact information from WHOIS queries. If you specify
`true`, WHOIS ("who is") queries return contact information either for
Amazon Registrar or for our registrar associate,
Gandi. If you specify `false`, WHOIS queries return the
information that you entered for the registrant contact (domain owner).

###### Note

You must specify the same privacy setting for the administrative, billing, registrant, and
technical contacts.

Type: Boolean

Required: No

**[TechPrivacy](#API_domains_UpdateDomainContactPrivacy_RequestSyntax)**

Whether you want to conceal contact information from WHOIS queries. If you specify
`true`, WHOIS ("who is") queries return contact information either for
Amazon Registrar or for our registrar associate,
Gandi. If you specify `false`, WHOIS queries return the
information that you entered for the technical contact.

###### Note

You must specify the same privacy setting for the administrative, billing, registrant, and
technical contacts.

Type: Boolean

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

**[OperationId](#API_domains_UpdateDomainContactPrivacy_ResponseSyntax)**

Identifier for tracking the progress of the request. To use this ID to query the
operation status, use GetOperationDetail.

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

### UpdateDomainContactPrivacy Example

This example illustrates one usage of UpdateDomainContactPrivacy.

#### Sample Request

```

POST / HTTP/1.1
host:route53domains.us-east-1.amazonaws.com
x-amz-date:20140711T205230Z
authorization:AWS4-HMAC-SHA256
              Credential=AKIAIOSFODNN7EXAMPLE/20140711/us-east-1/route53domains/aws4_request,
              SignedHeaders=content-length;content-type;host;user-agent;x-amz-date;x-amz-target,
              Signature=[calculated-signature]
x-amz-target:Route53Domains_v20140515.UpdateDomainContactPrivacy
user-agent:aws-sdk-java/1.8.3 Linux/2.6.18-164.el5PAE Java_HotSpot (TM )_Server_VM/24.60-b09/1.7.0_60
content-type:application/x-amz-json-1.1
content-length:[number of characters in the JSON string]
{
"DomainName":"example.com",
"AdminPrivacy":true,
"RegistrantPrivacy":true,
"TechPrivacy":true,
}
```

#### Sample Response

```

HTTP/1.1 200
Content-Length:[number of characters in the JSON string]
{
"OperationId":"777bc5da-fbf7-482c-b2ba-8946884a7dd6"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/route53domains-2014-05-15/updatedomaincontactprivacy.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/route53domains-2014-05-15/updatedomaincontactprivacy.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/route53domains-2014-05-15/updatedomaincontactprivacy.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/route53domains-2014-05-15/updatedomaincontactprivacy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53domains-2014-05-15/updatedomaincontactprivacy.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/route53domains-2014-05-15/updatedomaincontactprivacy.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/route53domains-2014-05-15/updatedomaincontactprivacy.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/route53domains-2014-05-15/updatedomaincontactprivacy.md)

- [AWS SDK for Python](../../../../services/goto/boto3/route53domains-2014-05-15/updatedomaincontactprivacy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53domains-2014-05-15/updatedomaincontactprivacy.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

UpdateDomainContact

UpdateDomainNameservers
