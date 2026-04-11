# ContributorInsightsSummary

Represents a Contributor Insights summary entry.

## Contents

###### Note

In the following list, the required parameters are described first.

**ContributorInsightsMode**

Indicates the current mode of CloudWatch Contributor Insights, specifying whether it
tracks all access and throttled events or throttled events only for the DynamoDB
table or index.

Type: String

Valid Values: `ACCESSED_AND_THROTTLED_KEYS | THROTTLED_KEYS`

Required: No

**ContributorInsightsStatus**

Describes the current status for contributor insights for the given table and index,
if applicable.

Type: String

Valid Values: `ENABLING | ENABLED | DISABLING | DISABLED | FAILED`

Required: No

**IndexName**

Name of the index associated with the summary, if any.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: No

**TableName**

Name of the table associated with the summary.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/contributorinsightssummary.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/contributorinsightssummary.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/contributorinsightssummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContinuousBackupsDescription

CreateGlobalSecondaryIndexAction

All content copied from https://docs.aws.amazon.com/.
