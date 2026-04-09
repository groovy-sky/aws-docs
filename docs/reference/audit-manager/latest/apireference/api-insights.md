# Insights

A summary of the latest analytics data for all your active assessments.

This summary is a snapshot of the data that your active assessments collected on the
`lastUpdated` date. It’s important to understand that the following totals
are daily counts based on this date — they aren’t a total sum to date.

The `Insights` data is eventually consistent. This means that, when you read
data from `Insights`, the response might not instantly reflect the results of a
recently completed write or update operation. If you repeat your read request after a few
hours, the response should return the latest data.

###### Note

If you delete an assessment or change its status to inactive,
`InsightsByAssessment` includes data for that assessment as
follows.

- **Inactive assessments** \- If Audit Manager
collected evidence for your assessment before you changed it inactive, that
evidence is included in the `InsightsByAssessment` counts for that
day.

- **Deleted assessments** \- If Audit Manager
collected evidence for your assessment before you deleted it, that evidence isn't
included in the `InsightsByAssessment` counts for that day.

## Contents

**activeAssessmentsCount**

The number of active assessments in Audit Manager.

Type: Integer

Required: No

**assessmentControlsCountByNoncompliantEvidence**

The number of assessment controls that collected non-compliant evidence on the
`lastUpdated` date.

Type: Integer

Required: No

**compliantEvidenceCount**

The number of compliance check evidence that Audit Manager classified as compliant
on the `lastUpdated` date. This includes evidence that was collected from
AWS Security Hub CSPM with a _Pass_ ruling, or collected from
AWS Config with a _Compliant_ ruling.

Type: Integer

Required: No

**inconclusiveEvidenceCount**

The number of evidence without a compliance check ruling. Evidence is inconclusive when
the associated control uses AWS Security Hub CSPM or AWS Config as a data
source but you didn't enable those services. This is also the case when a control uses a
data source that doesn’t support compliance checks (for example: manual evidence, API
calls, or AWS CloudTrail).

###### Note

If evidence has a compliance check status of _not applicable_,
it's classed as _inconclusive_ in `Insights` data.

Type: Integer

Required: No

**lastUpdated**

The time when the cross-assessment insights were last updated.

Type: Timestamp

Required: No

**noncompliantEvidenceCount**

The number of compliance check evidence that Audit Manager classified as
non-compliant on the `lastUpdated` date. This includes evidence that was
collected from AWS Security Hub CSPM with a _Fail_ ruling, or collected
from AWS Config with a _Non-compliant_ ruling.

Type: Integer

Required: No

**totalAssessmentControlsCount**

The total number of controls across all active assessments.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/insights.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/insights.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/insights.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FrameworkMetadata

InsightsByAssessment

All content copied from https://docs.aws.amazon.com/.
