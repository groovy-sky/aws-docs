# StringListAttributeBoostingConfiguration

Provides information on boosting `STRING_LIST` type document
attributes.

In the current boosting implementation, boosting focuses primarily on `DATE` attributes for recency and `STRING` attributes for source prioritization. `STRING_LIST` attributes can serve as additional boosting factors when needed, but are not supported when using `NativeIndexConfiguration` version 2.

###### Note

For `STRING` and `STRING_LIST` type document attributes to
be used for boosting on the console and the API, they must be enabled for search
using the [DocumentAttributeConfiguration](api-documentattributeconfiguration.md) object of the [UpdateIndex](api-updateindex.md) API. If you haven't enabled searching on these attributes,
you can't boost attributes of these data types on either the console or the
API.

For more information on how boosting document attributes work in Amazon Q Business, see
[Boosting using document attributes](../business-use-dg/metadata-boosting.md).

## Contents

**boostingLevel**

Specifies the priority of boosted document attributes in relation to other boosted attributes. This parameter determines how strongly the attribute influences document ranking in search results. `STRING_LIST` attributes can serve as additional boosting factors when needed, but are not supported when using `NativeIndexConfiguration` version 2.

Type: String

Valid Values: `NONE | LOW | MEDIUM | HIGH | VERY_HIGH | ONE | TWO`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/stringlistattributeboostingconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/stringlistattributeboostingconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/stringlistattributeboostingconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StringAttributeBoostingConfiguration

Subscription

All content copied from https://docs.aws.amazon.com/.
