# Impact

The dollar value of the anomaly.

## Contents

**MaxImpact**

The maximum dollar value that's observed for an anomaly.

Type: Double

Required: Yes

**TotalActualSpend**

The cumulative dollar amount that was actually spent during the anomaly.

Type: Double

Valid Range: Minimum value of 0.0.

Required: No

**TotalExpectedSpend**

The cumulative dollar amount that was expected to be spent during the anomaly. It is
calculated using advanced machine learning models to determine the typical spending
pattern based on historical data for a customer.

Type: Double

Valid Range: Minimum value of 0.0.

Required: No

**TotalImpact**

The cumulative dollar difference between the total actual spend and total expected
spend. It is calculated as `TotalActualSpend - TotalExpectedSpend`.

Type: Double

Required: No

**TotalImpactPercentage**

The cumulative percentage difference between the total actual spend and total expected
spend. It is calculated as `(TotalImpact / TotalExpectedSpend) * 100`. When
`TotalExpectedSpend` is zero, this field is omitted. Expected spend can
be zero in situations such as when you start to use a service for the first time.

Type: Double

Valid Range: Minimum value of 0.0.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ce-2017-10-25/impact.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ce-2017-10-25/impact.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ce-2017-10-25/impact.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GroupDefinition

InstanceDetails
