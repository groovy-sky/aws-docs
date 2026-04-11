# CreateControlMappingSource

The mapping attributes that determine the evidence source for a given control, along
with related parameters and metadata. This doesn't contain `mappingID`.

## Contents

**sourceDescription**

The description of the data source that determines where Audit Manager collects
evidence from for the control.

Type: String

Length Constraints: Maximum length of 1000.

Pattern: `^[\w\W\s\S]*$`

Required: No

**sourceFrequency**

Specifies how often evidence is collected from the control mapping source.

Type: String

Valid Values: `DAILY | WEEKLY | MONTHLY`

Required: No

**sourceKeyword**

A keyword that relates to the control data source.

For manual evidence, this keyword indicates if the manual evidence is a file or
text.

For automated evidence, this keyword identifies a specific CloudTrail event,
AWS Config rule, Security Hub CSPM control, or AWS API name.

To learn more about the supported keywords that you can use when mapping a control data
source, see the following pages in the _AWS Audit Manager User_
_Guide_:

- [AWS Config rules supported by AWS Audit Manager](../../../../services/audit-manager/latest/userguide/control-data-sources-config.md)

- [AWS Security Hub CSPM controls supported by AWS Audit Manager](../../../../services/audit-manager/latest/userguide/control-data-sources-ash.md)

- [API calls\
supported by AWS Audit Manager](../../../../services/audit-manager/latest/userguide/control-data-sources-api.md)

- [AWS CloudTrail event names supported by AWS Audit Manager](../../../../services/audit-manager/latest/userguide/control-data-sources-cloudtrail.md)

Type: [SourceKeyword](api-sourcekeyword.md) object

Required: No

**sourceName**

The name of the control mapping data source.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Required: No

**sourceSetUpOption**

The setup option for the data source. This option reflects if the evidence collection
method is automated or manual. If you don’t provide a value for
`sourceSetUpOption`, AWS Audit Manager automatically infers and populates
the correct value based on the `sourceType` that you specify.

Type: String

Valid Values: `System_Controls_Mapping | Procedural_Controls_Mapping`

Required: No

**sourceType**

Specifies which type of data source is used to collect evidence.

- The source can be an individual data source type, such as
`AWS_Cloudtrail`, `AWS_Config`,
`AWS_Security_Hub`, `AWS_API_Call`, or `MANUAL`.

- The source can also be a managed grouping of data sources, such as a
`Core_Control` or a `Common_Control`.

Type: String

Valid Values: `AWS_Cloudtrail | AWS_Config | AWS_Security_Hub | AWS_API_Call | MANUAL | Common_Control | Core_Control`

Required: No

**troubleshootingText**

The instructions for troubleshooting the control.

Type: String

Length Constraints: Maximum length of 1000.

Pattern: `^[\w\W\s\S]*$`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/createcontrolmappingsource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/createcontrolmappingsource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/createcontrolmappingsource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateAssessmentFrameworkControlSet

CreateDelegationRequest

All content copied from https://docs.aws.amazon.com/.
