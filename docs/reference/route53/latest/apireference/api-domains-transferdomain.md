# TransferDomain

Transfers a domain from another registrar to Amazon Route 53.

For more information about transferring domains, see the following topics:

- For transfer requirements, a detailed procedure, and information about viewing
the status of a domain that you're transferring to Route 53, see [Transferring Registration for a Domain to Amazon Route 53](../../../../services/route53/latest/developerguide/domain-transfer-to-route-53.md) in the
_Amazon Route 53 Developer Guide_.

- For information about how to transfer a domain from one AWS account to another, see [TransferDomainToAnotherAwsAccount](https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html).

- For information about how to transfer a domain to another domain registrar,
see [Transferring a Domain from Amazon Route 53 to Another Registrar](../../../../services/route53/latest/developerguide/domain-transfer-from-route-53.md) in
the _Amazon Route 53 Developer Guide_.

###### Important

During the transfer of any country code top-level domains (ccTLDs) to Route 53, except for .cc and .tv,
updates to the owner contact are ignored and the owner contact data from the registry is used.
You can
update the owner contact after the transfer is complete. For more information, see
[UpdateDomainContact](api-domains-updatedomaincontact.md).

If the registrar for your domain is also the DNS service provider for the domain, we
highly recommend that you transfer your DNS service to Route 53 or to another DNS
service provider before you transfer your registration. Some registrars provide free DNS
service when you purchase a domain registration. When you transfer the registration, the
previous registrar will not renew your domain registration and could end your DNS
service at any time.

###### Important

If the registrar for your domain is also the DNS service provider for the domain
and you don't transfer DNS service to another provider, your website, email, and the
web applications associated with the domain might become unavailable.

If the transfer is successful, this method returns an operation ID that you can use to
track the progress and completion of the action. If the transfer doesn't complete
successfully, the domain registrant will be notified by email.

## Request Syntax

```nohighlight

{
   "AdminContact": {
      "AddressLine1": "string",
      "AddressLine2": "string",
      "City": "string",
      "ContactType": "string",
      "CountryCode": "string",
      "Email": "string",
      "ExtraParams": [
         {
            "Name": "string",
            "Value": "string"
         }
      ],
      "Fax": "string",
      "FirstName": "string",
      "LastName": "string",
      "OrganizationName": "string",
      "PhoneNumber": "string",
      "State": "string",
      "ZipCode": "string"
   },
   "AuthCode": "string",
   "AutoRenew": boolean,
   "BillingContact": {
      "AddressLine1": "string",
      "AddressLine2": "string",
      "City": "string",
      "ContactType": "string",
      "CountryCode": "string",
      "Email": "string",
      "ExtraParams": [
         {
            "Name": "string",
            "Value": "string"
         }
      ],
      "Fax": "string",
      "FirstName": "string",
      "LastName": "string",
      "OrganizationName": "string",
      "PhoneNumber": "string",
      "State": "string",
      "ZipCode": "string"
   },
   "DomainName": "string",
   "DurationInYears": number,
   "IdnLangCode": "string",
   "Nameservers": [
      {
         "GlueIps": [ "string" ],
         "Name": "string"
      }
   ],
   "PrivacyProtectAdminContact": boolean,
   "PrivacyProtectBillingContact": boolean,
   "PrivacyProtectRegistrantContact": boolean,
   "PrivacyProtectTechContact": boolean,
   "RegistrantContact": {
      "AddressLine1": "string",
      "AddressLine2": "string",
      "City": "string",
      "ContactType": "string",
      "CountryCode": "string",
      "Email": "string",
      "ExtraParams": [
         {
            "Name": "string",
            "Value": "string"
         }
      ],
      "Fax": "string",
      "FirstName": "string",
      "LastName": "string",
      "OrganizationName": "string",
      "PhoneNumber": "string",
      "State": "string",
      "ZipCode": "string"
   },
   "TechContact": {
      "AddressLine1": "string",
      "AddressLine2": "string",
      "City": "string",
      "ContactType": "string",
      "CountryCode": "string",
      "Email": "string",
      "ExtraParams": [
         {
            "Name": "string",
            "Value": "string"
         }
      ],
      "Fax": "string",
      "FirstName": "string",
      "LastName": "string",
      "OrganizationName": "string",
      "PhoneNumber": "string",
      "State": "string",
      "ZipCode": "string"
   }
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[AdminContact](#API_domains_TransferDomain_RequestSyntax)**

Provides detailed contact information.

Type: [ContactDetail](api-domains-contactdetail.md) object

Required: Yes

**[AuthCode](#API_domains_TransferDomain_RequestSyntax)**

The authorization code for the domain. You get this value from the current
registrar.

Type: String

Length Constraints: Maximum length of 1024.

Required: No

**[AutoRenew](#API_domains_TransferDomain_RequestSyntax)**

Indicates whether the domain will be automatically renewed (true) or not (false). Auto
renewal only takes effect after the account is charged.

Default: true

Type: Boolean

Required: No

**[BillingContact](#API_domains_TransferDomain_RequestSyntax)**

Provides detailed contact information.

Type: [ContactDetail](api-domains-contactdetail.md) object

Required: No

**[DomainName](#API_domains_TransferDomain_RequestSyntax)**

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

**[DurationInYears](#API_domains_TransferDomain_RequestSyntax)**

Reserved for future use.

Currently, the effect of a domain transfer on the registration period varies by TLD. For information about how transferring a domain affects the expiration date, see the Transfer Term column in the pricing information at [Amazon Route 53 Pricing](http://aws.amazon.com/route53/pricing).

Default: 1

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 10.

Required: No

**[IdnLangCode](#API_domains_TransferDomain_RequestSyntax)**

Reserved for future use.

Type: String

Pattern: `|[A-Za-z]{2,3}`

Required: No

**[Nameservers](#API_domains_TransferDomain_RequestSyntax)**

Contains details for the host and glue IP addresses.

Type: Array of [Nameserver](api-domains-nameserver.md) objects

Required: No

**[PrivacyProtectAdminContact](#API_domains_TransferDomain_RequestSyntax)**

Whether you want to conceal contact information from WHOIS queries. If you specify
`true`, WHOIS ("who is") queries return contact information for the
registrar, the phrase "REDACTED FOR PRIVACY", or "On behalf of <domain name>
owner.".

###### Note

While some domains may allow different privacy settings per contact, we recommend
specifying the same privacy setting for all contacts.

Default: `true`

Type: Boolean

Required: No

**[PrivacyProtectBillingContact](#API_domains_TransferDomain_RequestSyntax)**

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

**[PrivacyProtectRegistrantContact](#API_domains_TransferDomain_RequestSyntax)**

Whether you want to conceal contact information from WHOIS queries. If you specify
`true`, WHOIS ("who is") queries return contact information either for
Amazon Registrar or for our registrar associate,
Gandi. If you specify `false`, WHOIS queries return the
information that you entered for the registrant contact (domain owner).

###### Note

You must specify the same privacy setting for the administrative, billing, registrant, and
technical contacts.

Default: `true`

Type: Boolean

Required: No

**[PrivacyProtectTechContact](#API_domains_TransferDomain_RequestSyntax)**

Whether you want to conceal contact information from WHOIS queries. If you specify
`true`, WHOIS ("who is") queries return contact information either for
Amazon Registrar or for our registrar associate,
Gandi. If you specify `false`, WHOIS queries return the
information that you entered for the technical contact.

###### Note

You must specify the same privacy setting for the administrative, billing, registrant, and
technical contacts.

Default: `true`

Type: Boolean

Required: No

**[RegistrantContact](#API_domains_TransferDomain_RequestSyntax)**

Provides detailed contact information.

Type: [ContactDetail](api-domains-contactdetail.md) object

Required: Yes

**[TechContact](#API_domains_TransferDomain_RequestSyntax)**

Provides detailed contact information.

Type: [ContactDetail](api-domains-contactdetail.md) object

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

**[OperationId](#API_domains_TransferDomain_ResponseSyntax)**

Identifier for tracking the progress of the request. To query the operation status,
use [GetOperationDetail](https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html).

Type: String

Length Constraints: Maximum length of 255.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**DomainLimitExceeded**

The number of domains has exceeded the allowed threshold for the account.

**message**

The number of domains has exceeded the allowed threshold for the account.

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

## Examples

### TransferDomain Example

This example illustrates one usage of TransferDomain.

#### Sample Request

```

POST / HTTP/1.1
host:route53domains.us-east-1.amazonaws.com
x-amz-date:20140711T205230Z
authorization:AWS4-HMAC-SHA256
              Credential=AKIAIOSFODNN7EXAMPLE/20140711/us-east-1/route53domains/aws4_request,
              SignedHeaders=content-length;content-type;host;user-agent;x-amz-date;x-amz-target,
              Signature=[calculated-signature]
x-amz-target:Route53Domains_v20140515.TransferDomain
user-agent:aws-sdk-java/1.8.3 Linux/2.6.18-164.el5PAE Java_HotSpot (TM )_Server_VM/24.60-b09/1.7.0_60
content-type:application/x-amz-json-1.1
content-length:[number of characters in the JSON string]
{
   "DomainName":"example.com",
   "DurationInYears":1,
   "Nameservers":[
      {
         "Name":"ns-2048.awsdns-64.com",
         "GlueIps":[
         "192.0.2.11"
         ]
      },
      {
         "Name":"ns-2049.awsdns-65.net",
         "GlueIps":[
         "192.0.2.12"
         ]
      }
   ],
    "AuthCode":"a42qxjz1",
   "AutoRenew":true,
   "AdminContact":{
      "FirstName":"John",
      "MiddleName":"Richard",
      "LastName":"Doe",
      "ContactType":"PERSON",
      "OrganizationName":"",
      "AddressLine1":"123 Any Street",
      "AddressLine2":"",
      "City":"Any Town",
      "State":"WA",
      "CountryCode":"US",
      "ZipCode":"98101",
      "PhoneNumber":"+1.1234567890",
      "Email":"john@example.com",
      "Fax":"+1.1234567891"
   },
   "RegistrantContact":{
      "FirstName":"John",
      "MiddleName":"Richard",
      "LastName":"Doe",
      "ContactType":"PERSON",
      "OrganizationName":"",
      "AddressLine1":"123 Any Street",
      "AddressLine2":"",
      "City":"Any Town",
      "State":"WA",
      "CountryCode":"US",
      "ZipCode":"98101",
      "PhoneNumber":"+1.1234567890",
      "Email":"john@example.com",
      "Fax":"+1.1234567891"
   },
   "TechContact":{
      "FirstName":"John",
      "MiddleName":"Richard",
      "LastName":"Doe",
      "ContactType":"PERSON",
      "OrganizationName":"",
      "AddressLine1":"123 Any Street",
      "AddressLine2":"",
      "City":"Any Town",
      "State":"WA",
      "CountryCode":"US",
      "ZipCode":"98101",
      "PhoneNumber":"+1.1234567890",
      "Email":"john@example.com",
      "Fax":"+1.1234567891"
   },
   "PrivacyProtectAdminContact":true,
   "PrivacyProtectRegistrantContact":true,
   "PrivacyProtectTechContact":true,
}
```

#### Sample Response

```

HTTP/1.1 200
Content-Length:[number of characters in the JSON string]
{
"OperationId":"308c56712-faa4-40fe-94c8-b423069de3f6"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/route53domains-2014-05-15/TransferDomain)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/route53domains-2014-05-15/TransferDomain)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53domains-2014-05-15/TransferDomain)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/route53domains-2014-05-15/TransferDomain)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53domains-2014-05-15/TransferDomain)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/route53domains-2014-05-15/TransferDomain)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/route53domains-2014-05-15/TransferDomain)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/route53domains-2014-05-15/TransferDomain)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/route53domains-2014-05-15/TransferDomain)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53domains-2014-05-15/TransferDomain)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RetrieveDomainAuthCode

TransferDomainToAnotherAwsAccount
