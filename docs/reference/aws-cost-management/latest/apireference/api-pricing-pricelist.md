# PriceList

This is the type of price list references that match your request.

## Contents

**CurrencyCode**

The three alphabetical character ISO-4217 currency code the Price List files are
denominated in.

Type: String

Pattern: `[A-Z]{3}`

Required: No

**FileFormats**

The format you want to retrieve your Price List files. The `FileFormat` can
be obtained from the [`ListPriceList`](api-pricing-listpricelists.md) response.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: No

**PriceListArn**

The unique identifier that maps to where your Price List files are located.
`PriceListArn` can be obtained from the [`ListPriceList`](api-pricing-listpricelists.md) response.

Type: String

Length Constraints: Minimum length of 18. Maximum length of 2048.

Pattern: `arn:[A-Za-z0-9][-.A-Za-z0-9]{0,62}:pricing:::price-list/[A-Za-z0-9+_/.-]{1,1023}`

Required: No

**RegionCode**

This is used to filter the Price List by AWS Region. For example, to get
the price list only for the `US East (N. Virginia)` Region, use
`us-east-1`. If nothing is specified, you retrieve price lists for all
applicable Regions. The available `RegionCode` list can be retrieved from [`GetAttributeValues`](api-pricing-getattributevalues.md) API.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/pricing-2017-10-15/PriceList)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/pricing-2017-10-15/PriceList)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/pricing-2017-10-15/PriceList)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Filter

Service
