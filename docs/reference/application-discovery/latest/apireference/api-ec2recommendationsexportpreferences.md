# Ec2RecommendationsExportPreferences

###### Important

AWS Application Discovery Service is no longer open to new customers. Existing customers can continue to use the service as normal. For more information, see
[AWS Application Discovery Service availability change](../../../../services/application-discovery/latest/userguide/application-discovery-service-availability-change.md).

Indicates that the exported data must include EC2 instance type matches for on-premises
servers that are discovered through AWS Application Discovery Service.

## Contents

**cpuPerformanceMetricBasis**

The recommended EC2 instance type that matches the CPU usage metric of server performance
data.

Type: [UsageMetricBasis](api-usagemetricbasis.md) object

Required: No

**enabled**

If set to true, the export [preferences](api-startexporttask.md#API_StartExportTask_RequestSyntax) is set to `Ec2RecommendationsExportPreferences`.

Type: Boolean

Required: No

**excludedInstanceTypes**

An array of instance types and families to exclude from recommendations. The following is an example: `["m5.4xlarge", "r3.large", "t3"]`

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 25.

Pattern: `[a-zA-Z0-9\d\.\-]+`

Required: No

**preferredRegion**

The target AWS Region for the recommendations. You can use any of the Region codes
available for the chosen service, as listed in [AWS service endpoints](../../../../general/latest/gr/rande.md) in the _AWS_
_General Reference_.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 30.

Pattern: `[a-z]{2}-[a-z\-]+-[0-9]+`

Required: No

**ramPerformanceMetricBasis**

The recommended EC2 instance type that matches the Memory usage metric of server
performance data.

Type: [UsageMetricBasis](api-usagemetricbasis.md) object

Required: No

**reservedInstanceOptions**

The contract type for a reserved instance. If blank, we assume an On-Demand instance is
preferred.

Type: [ReservedInstanceOptions](api-reservedinstanceoptions.md) object

Required: No

**tenancy**

The target tenancy to use for your recommended EC2 instances.

Type: String

Valid Values: `DEDICATED | SHARED`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/discovery-2015-11-01/ec2recommendationsexportpreferences.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/discovery-2015-11-01/ec2recommendationsexportpreferences.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/discovery-2015-11-01/ec2recommendationsexportpreferences.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeletionWarning

ExportFilter

All content copied from https://docs.aws.amazon.com/.
