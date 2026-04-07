# MetricDataResult

Contains a single data point from a capacity metrics query, including the dimension values, timestamp, and metric values for that specific combination.

## Contents

**dimension**

The dimension values that identify this specific data point, such as account ID, region, and instance family.

Type: [CapacityManagerDimension](api-capacitymanagerdimension.md) object

Required: No

**MetricValueSet.N**

The metric values and statistics for this data point, containing the actual capacity usage numbers.

Type: Array of [MetricValue](api-metricvalue.md) objects

Required: No

**timestamp**

The timestamp for this data point, indicating when the capacity usage occurred.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/metricdataresult.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/metricdataresult.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/metricdataresult.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

MemoryMiBRequest

MetricPoint
