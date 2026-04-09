# NumberAttributeBoostingConfiguration

Provides information on boosting `NUMBER` type document attributes.

In the current boosting implementation, boosting focuses primarily on `DATE` attributes for recency and `STRING` attributes for source prioritization. `NUMBER` attributes can serve as additional boosting factors when needed, but are not supported when using `NativeIndexConfiguration` version 2.

For more information on how boosting document attributes work in Amazon Q Business, see
[Boosting using document attributes](../business-use-dg/metadata-boosting.md).

## Contents

**boostingLevel**

Specifies the priority of boosted document attributes in relation to other boosted attributes. This parameter determines how strongly the attribute influences document ranking in search results. `NUMBER` attributes can serve as additional boosting factors when needed, but are not supported when using `NativeIndexConfiguration` version 2.

Type: String

Valid Values: `NONE | LOW | MEDIUM | HIGH | VERY_HIGH | ONE | TWO`

Required: Yes

**boostingType**

Specifies whether higher or lower numeric values should be prioritized when boosting. Valid values are ASCENDING (higher numbers are more important) and DESCENDING (lower numbers are more important).

Type: String

Valid Values: `PRIORITIZE_LARGER_VALUES | PRIORITIZE_SMALLER_VALUES`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/numberattributeboostingconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/numberattributeboostingconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/numberattributeboostingconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NoAuthConfiguration

OAuth2ClientCredentialConfiguration

All content copied from https://docs.aws.amazon.com/.
