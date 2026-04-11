# DomainValidation

Contains information about the validation of each domain name in the
certificate.

## Contents

###### Note

In the following list, the required parameters are described first.

**DomainName**

A fully qualified domain name (FQDN) in the certificate. For example,
`www.example.com` or `example.com`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 253.

Pattern: `(\*\.)?(((?!-)[A-Za-z0-9-]{0,62}[A-Za-z0-9])\.)+((?!-)[A-Za-z0-9-]{1,62}[A-Za-z0-9])`

Required: Yes

**HttpRedirect**

Contains information for HTTP-based domain validation of certificates requested through Amazon CloudFront and issued by ACM.
This field exists only when the certificate type is `AMAZON_ISSUED` and the validation method is `HTTP`.

Type: [HttpRedirect](api-httpredirect.md) object

Required: No

**ResourceRecord**

Contains the CNAME record that you add to your DNS database for domain validation. For
more information, see [Use DNS to Validate Domain\
Ownership](../../../../services/acm/latest/userguide/gs-acm-validate-dns.md).

###### Note

The CNAME information that you need does not include the name of your domain. If
you include your domain name in the DNS database CNAME record, validation fails. For
example, if the name is
`_a79865eb4cd1a6ab990a45779b4e0b96.yourdomain.com`, only
`_a79865eb4cd1a6ab990a45779b4e0b96` must be used.

Type: [ResourceRecord](api-resourcerecord.md) object

Required: No

**ValidationDomain**

The domain name that ACM used to send domain validation emails.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 253.

Pattern: `(\*\.)?(((?!-)[A-Za-z0-9-]{0,62}[A-Za-z0-9])\.)+((?!-)[A-Za-z0-9-]{1,62}[A-Za-z0-9])`

Required: No

**ValidationEmails**

A list of email addresses that ACM used to send domain validation emails.

Type: Array of strings

Required: No

**ValidationMethod**

Specifies the domain validation method.

Type: String

Valid Values: `EMAIL | DNS | HTTP`

Required: No

**ValidationStatus**

The validation status of the domain name. This can be one of the following
values:

- `PENDING_VALIDATION`

- `` SUCCESS

- `` FAILED

Type: String

Valid Values: `PENDING_VALIDATION | SUCCESS | FAILED`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/acm-2015-12-08/domainvalidation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/acm-2015-12-08/domainvalidation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/acm-2015-12-08/domainvalidation.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DnsNameFilter

DomainValidationOption

All content copied from https://docs.aws.amazon.com/.
