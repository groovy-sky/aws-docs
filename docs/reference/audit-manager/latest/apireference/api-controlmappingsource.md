# ControlMappingSource

The data source that determines where AWS Audit Manager collects evidence from for
the control.

## Contents

**sourceDescription**

The description of the source.

Type: String

Length Constraints: Maximum length of 1000.

Pattern: `^[\w\W\s\S]*$`

Required: No

**sourceFrequency**

Specifies how often evidence is collected from the control mapping source.

Type: String

Valid Values: `DAILY | WEEKLY | MONTHLY`

Required: No

**sourceId**

The unique identifier for the source.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

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

- [AWS Config rules supported by AWS Audit Manager](https://docs.aws.amazon.com/audit-manager/latest/userguide/control-data-sources-config.html)

- [AWS Security Hub CSPM controls supported by AWS Audit Manager](https://docs.aws.amazon.com/audit-manager/latest/userguide/control-data-sources-ash.html)

- [API calls\
supported by AWS Audit Manager](https://docs.aws.amazon.com/audit-manager/latest/userguide/control-data-sources-api.html)

- [AWS CloudTrail event names supported by AWS Audit Manager](https://docs.aws.amazon.com/audit-manager/latest/userguide/control-data-sources-cloudtrail.html)

Type: [SourceKeyword](https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_SourceKeyword.html) object

Required: No

**sourceName**

The name of the source.

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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/auditmanager-2017-07-25/ControlMappingSource)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/auditmanager-2017-07-25/ControlMappingSource)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/auditmanager-2017-07-25/ControlMappingSource)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ControlInsightsMetadataItem

ControlMetadata
