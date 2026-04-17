---
title: "MultiRegionAccessPointPolicyDocument"
---

# MultiRegionAccessPointPolicyDocument

The Multi-Region Access Point access control policy.

When you update the policy, the update is first listed as the proposed policy. After the
update is finished and all Regions have been updated, the proposed policy is listed as the
established policy. If both policies have the same version number, the proposed policy is
the established policy.

## Contents

**Established**

The last established policy for the Multi-Region Access Point.

Type: [EstablishedMultiRegionAccessPointPolicy](api-control-establishedmultiregionaccesspointpolicy.md) data type

Required: No

**Proposed**

The proposed policy for the Multi-Region Access Point.

Type: [ProposedMultiRegionAccessPointPolicy](api-control-proposedmultiregionaccesspointpolicy.md) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3control-2018-08-20/MultiRegionAccessPointPolicyDocument)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3control-2018-08-20/MultiRegionAccessPointPolicyDocument)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3control-2018-08-20/MultiRegionAccessPointPolicyDocument)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Metrics

MultiRegionAccessPointRegionalResponse

All content copied from https://docs.aws.amazon.com/.
