# RenewalSummary

Contains information about the status of ACM's [managed renewal](../../../../services/acm/latest/userguide/acm-renewal.md) for the certificate. This
structure exists only when the certificate type is `AMAZON_ISSUED`.

## Contents

###### Note

In the following list, the required parameters are described first.

**DomainValidationOptions**

Contains information about the validation of each domain name in the certificate, as
it pertains to ACM's [managed renewal](../../../../services/acm/latest/userguide/acm-renewal.md). This is different from the initial validation that occurs
as a result of the [RequestCertificate](api-requestcertificate.md) request. This field exists only
when the certificate type is `AMAZON_ISSUED`.

Type: Array of [DomainValidation](api-domainvalidation.md) objects

Array Members: Minimum number of 1 item. Maximum number of 1000 items.

Required: Yes

**RenewalStatus**

The status of ACM's [managed renewal](../../../../services/acm/latest/userguide/acm-renewal.md) of the certificate.

Type: String

Valid Values: `PENDING_AUTO_RENEWAL | PENDING_VALIDATION | SUCCESS | FAILED`

Required: Yes

**UpdatedAt**

The time at which the renewal summary was last updated.

Type: Timestamp

Required: Yes

**RenewalStatusReason**

The reason that a renewal request was unsuccessful.

Type: String

Valid Values: `NO_AVAILABLE_CONTACTS | ADDITIONAL_VERIFICATION_REQUIRED | DOMAIN_NOT_ALLOWED | INVALID_PUBLIC_DOMAIN | DOMAIN_VALIDATION_DENIED | CAA_ERROR | PCA_LIMIT_EXCEEDED | PCA_INVALID_ARN | PCA_INVALID_STATE | PCA_REQUEST_FAILED | PCA_NAME_CONSTRAINTS_VALIDATION | PCA_RESOURCE_NOT_FOUND | PCA_INVALID_ARGS | PCA_INVALID_DURATION | PCA_ACCESS_DENIED | SLR_NOT_FOUND | OTHER`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/acm-2015-12-08/renewalsummary.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/acm-2015-12-08/renewalsummary.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/acm-2015-12-08/renewalsummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OtherName

ResourceRecord

All content copied from https://docs.aws.amazon.com/.
