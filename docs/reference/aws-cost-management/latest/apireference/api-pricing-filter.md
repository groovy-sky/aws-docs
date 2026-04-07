# Filter

The constraints that you want all returned products to match.

## Contents

**Field**

The product metadata field that you want to filter on. You can filter by just the
service code to see all products for a specific service, filter
by just the attribute name to see a specific attribute for multiple services, or use both a service code
and an attribute name to retrieve only products that match both fields.

Valid values include: `ServiceCode`, and all attribute names

For example, you can filter by the `AmazonEC2` service code and the
`volumeType` attribute name to get the prices for only Amazon EC2 volumes.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: Yes

**Type**

The type of filter that you want to use.

Valid values are:

- `TERM_MATCH`: Returns only
products that match both the given filter field and the given value.

- `EQUALS`: Returns products that have a field value exactly matching the provided value.

- `CONTAINS`: Returns products where the field value contains the provided value as a substring.

- `ANY_OF`: Returns products where the field value is any of the provided values.

- `NONE_OF`: Returns products where the field value is not any of the provided values.

Type: String

Valid Values: `TERM_MATCH | EQUALS | CONTAINS | ANY_OF | NONE_OF`

Required: Yes

**Value**

The service code or attribute value that you want to filter by. If you're filtering
by service code this is the actual service code, such as `AmazonEC2`. If you're
filtering by attribute name, this is the attribute value that you want the returned
products to match, such as a `Provisioned IOPS` volume.

For `ANY_OF` and `NONE_OF` filter types, you can provide multiple values as a comma-separated string. For example, `t2.micro,t2.small,t2.medium` or `Compute optimized, GPU instance, Micro instances`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/pricing-2017-10-15/Filter)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/pricing-2017-10-15/Filter)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/pricing-2017-10-15/Filter)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AttributeValue

PriceList
