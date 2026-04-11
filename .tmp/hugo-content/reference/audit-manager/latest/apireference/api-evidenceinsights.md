# EvidenceInsights

A breakdown of the latest compliance check status for the evidence in your Audit Manager assessments.

## Contents

**compliantEvidenceCount**

The number of compliance check evidence that Audit Manager classified as compliant.
This includes evidence that was collected from AWS Security Hub CSPM with a
_Pass_ ruling, or collected from AWS Config with a
_Compliant_ ruling.

Type: Integer

Required: No

**inconclusiveEvidenceCount**

The number of evidence that a compliance check ruling isn't available for. Evidence is
inconclusive when the associated control uses AWS Security Hub CSPM or AWS Config as a data source but you didn't enable those services. This is also the case when a
control uses a data source that doesn’t support compliance checks (for example, manual
evidence, API calls, or AWS CloudTrail).

###### Note

If evidence has a compliance check status of _not applicable_ in
the console, it's classified as _inconclusive_ in
`EvidenceInsights` data.

Type: Integer

Required: No

**noncompliantEvidenceCount**

The number of compliance check evidence that Audit Manager classified as
non-compliant. This includes evidence that was collected from AWS Security Hub CSPM with a
_Fail_ ruling, or collected from AWS Config with a
_Non-compliant_ ruling.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/evidenceinsights.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/evidenceinsights.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/evidenceinsights.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EvidenceFinderEnablement

Framework

All content copied from https://docs.aws.amazon.com/.
