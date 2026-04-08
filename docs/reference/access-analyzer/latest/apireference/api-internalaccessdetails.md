# InternalAccessDetails

Contains information about an internal access finding. This includes details about the
access that was identified within your AWS organization or account.

## Contents

**accessType**

The type of internal access identified in the finding. This indicates how the access is
granted within your AWS environment.

Type: String

Valid Values: `INTRA_ACCOUNT | INTRA_ORG`

Required: No

**action**

The action in the analyzed policy statement that has internal access permission to
use.

Type: Array of strings

Required: No

**condition**

The condition in the analyzed policy statement that resulted in an internal access
finding.

Type: String to string map

Required: No

**principal**

The principal that has access to a resource within the internal environment.

Type: String to string map

Required: No

**principalOwnerAccount**

The AWS account ID that owns the principal identified in the internal access
finding.

Type: String

Required: No

**principalType**

The type of principal identified in the internal access finding, such as IAM role or
IAM user.

Type: String

Valid Values: `IAM_ROLE | IAM_USER`

Required: No

**resourceControlPolicyRestriction**

The type of restriction applied to the finding by the resource owner with an AWS Organizations
resource control policy (RCP).

- `APPLICABLE`: There is an RCP present in the organization but
IAM Access Analyzer does not include it in the evaluation of effective permissions. For
example, if `s3:DeleteObject` is blocked by the RCP and the restriction is
`APPLICABLE`, then `s3:DeleteObject` would still be included
in the list of actions for the finding. Only applicable to internal access findings
with the account as the zone of trust.

- `FAILED_TO_EVALUATE_RCP`: There was an error evaluating the RCP.

- `NOT_APPLICABLE`: There was no RCP present in the organization. For
internal access findings with the account as the zone of trust,
`NOT_APPLICABLE` could also indicate that there was no RCP applicable
to the resource.

- `APPLIED`: An RCP is present in the organization and IAM Access Analyzer
included it in the evaluation of effective permissions. For example, if
`s3:DeleteObject` is blocked by the RCP and the restriction is
`APPLIED`, then `s3:DeleteObject` would not be included in
the list of actions for the finding. Only applicable to internal access findings with
the organization as the zone of trust.

Type: String

Valid Values: `APPLICABLE | FAILED_TO_EVALUATE_RCP | NOT_APPLICABLE | APPLIED`

Required: No

**serviceControlPolicyRestriction**

The type of restriction applied to the finding by an AWS Organizations service control policy
(SCP).

- `APPLICABLE`: There is an SCP present in the organization but
IAM Access Analyzer does not include it in the evaluation of effective permissions. Only
applicable to internal access findings with the account as the zone of trust.

- `FAILED_TO_EVALUATE_SCP`: There was an error evaluating the SCP.

- `NOT_APPLICABLE`: There was no SCP present in the organization. For
internal access findings with the account as the zone of trust,
`NOT_APPLICABLE` could also indicate that there was no SCP applicable
to the principal.

- `APPLIED`: An SCP is present in the organization and IAM Access Analyzer
included it in the evaluation of effective permissions. Only applicable to internal
access findings with the organization as the zone of trust.

Type: String

Valid Values: `APPLICABLE | FAILED_TO_EVALUATE_SCP | NOT_APPLICABLE | APPLIED`

Required: No

**sources**

The sources of the internal access finding. This indicates how the access that generated
the finding is granted within your AWS environment.

Type: Array of [FindingSource](api-findingsource.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/internalaccessdetails.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/internalaccessdetails.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/internalaccessdetails.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

InternalAccessConfiguration

InternalAccessFindingsStatistics
