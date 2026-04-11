# ControlInsightsMetadataItem

A summary of the latest analytics data for a specific control.

This data reflects the total counts for the specified control across all active
assessments. Control insights are grouped by control domain, and ranked by the highest
total count of non-compliant evidence.

## Contents

**evidenceInsights**

A breakdown of the compliance check status for the evidence that’s associated with the
control.

Type: [EvidenceInsights](api-evidenceinsights.md) object

Required: No

**id**

The unique identifier for the control.

Type: String

Length Constraints: Minimum length of 13. Maximum length of 2048.

Pattern: `^arn:.*:controlcatalog:.*:.*:domain/.*|UNCATEGORIZED|^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

**lastUpdated**

The time when the control insights were last updated.

Type: Timestamp

Required: No

**name**

The name of the control.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 2048.

Pattern: `.*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/controlinsightsmetadataitem.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/controlinsightsmetadataitem.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/controlinsightsmetadataitem.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ControlInsightsMetadataByAssessmentItem

ControlMappingSource

All content copied from https://docs.aws.amazon.com/.
