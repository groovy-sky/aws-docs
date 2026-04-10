# StackSetAutoDeploymentTargetSummary

One of the targets for the StackSet. Returned by the [ListStackSetAutoDeploymentTargets](api-liststacksetautodeploymenttargets.md) API operation.

## Contents

**OrganizationalUnitId**

The organization root ID or organizational unit (OU) IDs where the StackSet is
targeted.

Type: String

Pattern: `^(ou-[a-z0-9]{4,32}-[a-z0-9]{8,32}|r-[a-z0-9]{4,32})$`

Required: No

**Regions.member.N**

The list of Regions targeted for this organization or OU.

Type: Array of strings

Pattern: `^[a-zA-Z0-9-]{1,128}$`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/stacksetautodeploymenttargetsummary.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/stacksetautodeploymenttargetsummary.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/stacksetautodeploymenttargetsummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StackSet

StackSetDriftDetectionDetails

All content copied from https://docs.aws.amazon.com/.
