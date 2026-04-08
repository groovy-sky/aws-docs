# ExportPreferences

###### Important

AWS Application Discovery Service is no longer open to new customers. Existing customers can continue to use the service as normal. For more information, see
[AWS Application Discovery Service availability change](../../../../services/application-discovery/latest/userguide/application-discovery-service-availability-change.md).

Indicates the type of data that is being exported. Only one
`ExportPreferences` can be enabled for a [StartExportTask](api-startexporttask.md) action.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**ec2RecommendationsPreferences**

If enabled, exported data includes EC2 instance type matches for on-premises servers
discovered through AWS Application Discovery Service.

Type: [Ec2RecommendationsExportPreferences](api-ec2recommendationsexportpreferences.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/discovery-2015-11-01/exportpreferences.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/discovery-2015-11-01/exportpreferences.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/discovery-2015-11-01/exportpreferences.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ExportInfo

FailedConfiguration
