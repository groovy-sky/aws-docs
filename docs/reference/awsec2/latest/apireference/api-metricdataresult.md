# MetricDataResult

Contains a single data point from a capacity metrics query, including the dimension values, timestamp, and metric values for that specific combination.

## Contents

**dimension**

The dimension values that identify this specific data point, such as account ID, region, and instance family.

Type: [CapacityManagerDimension](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CapacityManagerDimension.html) object

Required: No

**MetricValueSet.N**

The metric values and statistics for this data point, containing the actual capacity usage numbers.

Type: Array of [MetricValue](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_MetricValue.html) objects

Required: No

**timestamp**

The timestamp for this data point, indicating when the capacity usage occurred.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/MetricDataResult)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/MetricDataResult)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/MetricDataResult)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MemoryMiBRequest

MetricPoint
