# ControlDomainInsights

A summary of the latest analytics data for a specific control domain.

Control domain insights are grouped by control domain, and ranked by the highest total
count of non-compliant evidence.

## Contents

**controlsCountByNoncompliantEvidence**

The number of controls in the control domain that collected non-compliant evidence on
the `lastUpdated` date.

Type: Integer

Required: No

**evidenceInsights**

A breakdown of the compliance check status for the evidence that’s associated with the
control domain.

Type: [EvidenceInsights](api-evidenceinsights.md) object

Required: No

**id**

The unique identifier for the control domain. Audit Manager supports the control
domains that are provided by AWS Control Catalog. For information about how
to find a list of available control domains, see [`ListDomains`](../../../controlcatalog/latest/apireference/api-listdomains.md) in the AWS Control Catalog API
Reference.

Type: String

Length Constraints: Minimum length of 13. Maximum length of 2048.

Pattern: `^arn:.*:controlcatalog:.*:.*:domain/.*|UNCATEGORIZED|^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

**lastUpdated**

The time when the control domain insights were last updated.

Type: Timestamp

Required: No

**name**

The name of the control domain.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 2048.

Pattern: `.*`

Required: No

**totalControlsCount**

The total number of controls in the control domain.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/controldomaininsights.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/controldomaininsights.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/controldomaininsights.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ControlComment

ControlInsightsMetadataByAssessmentItem

All content copied from https://docs.aws.amazon.com/.
