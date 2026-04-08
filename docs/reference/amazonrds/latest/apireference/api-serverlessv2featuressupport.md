# ServerlessV2FeaturesSupport

Specifies any Aurora Serverless v2 properties or limits that differ between Aurora engine versions and platform versions.
You can test the values of this attribute when deciding which Aurora version to use in a new or upgraded
DB cluster. You can also retrieve the version of an existing DB cluster and check whether that version
supports certain Aurora Serverless v2 features before you attempt to use those features.

## Contents

###### Note

In the following list, the required parameters are described first.

**MaxCapacity**

Specifies the upper Aurora Serverless v2 capacity limit for a particular engine version or platform version.
Depending on the engine version, the maximum capacity for an Aurora Serverless v2 cluster might be
`256` or `128`.

Type: Double

Required: No

**MinCapacity**

If the minimum capacity is 0 ACUs, the engine version or platform version supports the automatic pause/resume
feature of Aurora Serverless v2.

Type: Double

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/serverlessv2featuressupport.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/serverlessv2featuressupport.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/serverlessv2featuressupport.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ScalingConfigurationInfo

ServerlessV2ScalingConfiguration
