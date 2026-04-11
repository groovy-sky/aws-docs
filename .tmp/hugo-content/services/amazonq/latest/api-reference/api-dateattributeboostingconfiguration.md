# DateAttributeBoostingConfiguration

Provides information on boosting `DATE` type document attributes.

For more information on how boosting document attributes work in Amazon Q Business, see
[Boosting using document attributes](../business-use-dg/metadata-boosting.md).

## Contents

**boostingLevel**

Specifies the priority tier ranking of boosting applied to document attributes. For version 2, this parameter indicates the relative ranking between boosted fields (ONE being highest priority, TWO being second highest, etc.) and determines the order in which attributes influence document ranking in search results. For version 1, this parameter specifies the boosting intensity. For version 2, boosting intensity (VERY HIGH, HIGH, MEDIUM, LOW, NONE) are not supported. Note that in version 2, you are not allowed to boost on only one field and make this value TWO.

Type: String

Valid Values: `NONE | LOW | MEDIUM | HIGH | VERY_HIGH | ONE | TWO`

Required: Yes

**boostingDurationInSeconds**

Specifies the duration, in seconds, of a boost applies to a `DATE` type
document attribute.

Type: Long

Valid Range: Minimum value of 0. Maximum value of 999999999.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/dateattributeboostingconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/dateattributeboostingconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/dateattributeboostingconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSourceVpcConfiguration

DeleteDocument

All content copied from https://docs.aws.amazon.com/.
