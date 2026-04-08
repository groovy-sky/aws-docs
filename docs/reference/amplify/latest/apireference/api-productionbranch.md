# ProductionBranch

Describes the information about a production branch for an Amplify app.

## Contents

**branchName**

The branch name for the production branch.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `(?s).+`

Required: No

**lastDeployTime**

The last deploy time of the production branch.

Type: Timestamp

Required: No

**status**

The status of the production branch.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 7.

Pattern: `.{3,7}`

Required: No

**thumbnailUrl**

The thumbnail URL for the production branch.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2000.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/amplify-2017-07-25/productionbranch.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/amplify-2017-07-25/productionbranch.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/amplify-2017-07-25/productionbranch.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

JobSummary

Step
