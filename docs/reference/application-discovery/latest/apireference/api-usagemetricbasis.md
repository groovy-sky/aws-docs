# UsageMetricBasis

###### Important

AWS Application Discovery Service is no longer open to new customers. Existing customers can continue to use the service as normal. For more information, see
[AWS Application Discovery Service availability change](../../../../services/application-discovery/latest/userguide/application-discovery-service-availability-change.md).

Specifies the performance metrics to use for the server that is used for recommendations.

## Contents

**name**

A utilization metric that is used by the recommendations.

Type: String

Pattern: `^(p(\d{1,2}|100)|AVG|SPEC|MAX)$`

Required: No

**percentageAdjust**

Specifies the percentage of the specified utilization metric that is used by the
recommendations.

Type: Double

Valid Range: Minimum value of 0.0. Maximum value of 100.0.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/discovery-2015-11-01/usagemetricbasis.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/discovery-2015-11-01/usagemetricbasis.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/discovery-2015-11-01/usagemetricbasis.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagFilter

Common Parameters

All content copied from https://docs.aws.amazon.com/.
