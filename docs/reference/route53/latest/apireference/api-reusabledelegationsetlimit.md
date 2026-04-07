# ReusableDelegationSetLimit

A complex type that contains the type of limit that you specified in the request and
the current value for that limit.

## Contents

**Type**

The limit that you requested: `MAX_ZONES_BY_REUSABLE_DELEGATION_SET`, the
maximum number of hosted zones that you can associate with the specified reusable
delegation set.

Type: String

Valid Values: `MAX_ZONES_BY_REUSABLE_DELEGATION_SET`

Required: Yes

**Value**

The current value for the `MAX_ZONES_BY_REUSABLE_DELEGATION_SET`
limit.

Type: Long

Valid Range: Minimum value of 1.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53-2013-04-01/ReusableDelegationSetLimit)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53-2013-04-01/ReusableDelegationSetLimit)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53-2013-04-01/ReusableDelegationSetLimit)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ResourceTagSet

StatusReport
