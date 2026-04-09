# StringAttributeBoostingConfiguration

Provides information on boosting `STRING` type document attributes.

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

Specifies the priority tier ranking of boosting applied to document attributes. For version 2, this parameter indicates the relative ranking between boosted fields (ONE being highest priority, TWO being second highest, etc.) and determines the order in which attributes influence document ranking in search results. For version 1, this parameter specifies the boosting intensity. For version 2, boosting intensity (VERY HIGH, HIGH, MEDIUM, LOW, NONE) are not supported. Note that in version 2, you are not allowed to boost on only one field and make this value TWO.

Type: String

Valid Values: `NONE | LOW | MEDIUM | HIGH | VERY_HIGH | ONE | TWO`

Required: Yes

**attributeValueBoosting**

Specifies specific values of a `STRING` type document attribute being boosted. When using `NativeIndexConfiguration` version 2, you can specify up to five values in order of priority.

Type: String to string map

Map Entries: Maximum number of 10 items.

Key Length Constraints: Minimum length of 1. Maximum length of 2048.

Valid Values: `LOW | MEDIUM | HIGH | VERY_HIGH | ONE | TWO | THREE | FOUR | FIVE`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/stringattributeboostingconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/stringattributeboostingconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/stringattributeboostingconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SourceDetails

StringListAttributeBoostingConfiguration

All content copied from https://docs.aws.amazon.com/.
