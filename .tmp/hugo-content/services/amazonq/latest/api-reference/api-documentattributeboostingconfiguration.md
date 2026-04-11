# DocumentAttributeBoostingConfiguration

Provides information on boosting supported Amazon Q Business document attribute types.
When an end user chat query matches document attributes that have been boosted,
Amazon Q Business prioritizes generating responses from content that matches the boosted
document attributes.

In version 2, boosting uses numeric values (ONE, TWO) to indicate priority tiers that establish clear hierarchical relationships between boosted attributes. This allows for more precise control over how different attributes influence search results.

###### Note

For `STRING` and `STRING_LIST` type document attributes to
be used for boosting on the console and the API, they must be enabled for search
using the [DocumentAttributeConfiguration](api-documentattributeconfiguration.md) object of the [UpdateIndex](api-updateindex.md) API. If you haven't enabled searching on these attributes,
you can't boost attributes of these data types on either the console or the
API.

For more information on how boosting document attributes work in Amazon Q Business, see
[Boosting using document attributes](../business-use-dg/metadata-boosting.md).

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**dateConfiguration**

Provides information on boosting `DATE` type document attributes.

Version 2 assigns priority tiers to `DATE` attributes, establishing clear hierarchical relationships with other boosted attributes.

Type: [DateAttributeBoostingConfiguration](api-dateattributeboostingconfiguration.md) object

Required: No

**numberConfiguration**

Provides information on boosting `NUMBER` type document attributes.

`NUMBER` attributes are not supported when using `NativeIndexConfiguration` version 2, which focuses on `DATE` attributes for recency and `STRING` attributes for source prioritization.

Type: [NumberAttributeBoostingConfiguration](api-numberattributeboostingconfiguration.md) object

Required: No

**stringConfiguration**

Provides information on boosting `STRING` type document attributes.

Version 2 assigns priority tiers to `STRING` attributes, establishing clear hierarchical relationships with other boosted attributes.

Type: [StringAttributeBoostingConfiguration](api-stringattributeboostingconfiguration.md) object

Required: No

**stringListConfiguration**

Provides information on boosting `STRING_LIST` type document attributes.

`STRING_LIST` attributes are not supported when using `NativeIndexConfiguration` version 2, which focuses on `DATE` attributes for recency and `STRING` attributes for source prioritization.

Type: [StringListAttributeBoostingConfiguration](api-stringlistattributeboostingconfiguration.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/documentattributeboostingconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/documentattributeboostingconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/documentattributeboostingconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentAttribute

DocumentAttributeCondition

All content copied from https://docs.aws.amazon.com/.
