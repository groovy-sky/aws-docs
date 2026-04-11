# App

Represents the different branches of a repository for building, deploying, and hosting
an Amplify app.

## Contents

**appArn**

The Amazon Resource Name (ARN) of the Amplify app.

Type: String

Length Constraints: Maximum length of 1000.

Required: Yes

**appId**

The unique ID of the Amplify app.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 20.

Pattern: `d[a-z0-9]+`

Required: Yes

**createTime**

A timestamp of when Amplify created the application.

Type: Timestamp

Required: Yes

**defaultDomain**

The default domain for the Amplify app.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1000.

Required: Yes

**description**

The description for the Amplify app.

Type: String

Length Constraints: Maximum length of 1000.

Pattern: `(?s).*`

Required: Yes

**enableBasicAuth**

Enables basic authorization for the Amplify app's branches.

Type: Boolean

Required: Yes

**enableBranchAutoBuild**

Enables the auto-building of branches for the Amplify app.

Type: Boolean

Required: Yes

**environmentVariables**

The environment variables for the Amplify app.

For a list of the environment variables that are accessible to Amplify by default, see
[Amplify\
Environment variables](../../../../services/amplify/latest/userguide/amplify-console-environment-variables.md) in the _Amplify Hosting User_
_Guide_.

Type: String to string map

Key Length Constraints: Maximum length of 255.

Key Pattern: `(?s).*`

Value Length Constraints: Maximum length of 5500.

Value Pattern: `(?s).*`

Required: Yes

**name**

The name for the Amplify app.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `(?s).+`

Required: Yes

**platform**

The platform for the Amplify app. For a static app, set the platform type to
`WEB`. For a dynamic server-side rendered (SSR) app, set the platform
type to `WEB_COMPUTE`. For an app requiring Amplify Hosting's original SSR
support only, set the platform type to `WEB_DYNAMIC`.

If you are deploying an SSG only app with Next.js 14 or later, you must use the
platform type `WEB_COMPUTE`.

Type: String

Valid Values: `WEB | WEB_DYNAMIC | WEB_COMPUTE`

Required: Yes

**repository**

The Git repository for the Amplify app.

Type: String

Length Constraints: Maximum length of 1000.

Pattern: `(?s).*`

Required: Yes

**updateTime**

A timestamp of when Amplify updated the application.

Type: Timestamp

Required: Yes

**autoBranchCreationConfig**

Describes the automated branch creation configuration for the Amplify app.

Type: [AutoBranchCreationConfig](api-autobranchcreationconfig.md) object

Required: No

**autoBranchCreationPatterns**

Describes the automated branch creation glob patterns for the Amplify app.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `(?s).+`

Required: No

**basicAuthCredentials**

The basic authorization credentials for branches for the Amplify app. You must
base64-encode the authorization credentials and provide them in the format
`user:password`.

Type: String

Length Constraints: Maximum length of 2000.

Pattern: `(?s).*`

Required: No

**buildSpec**

Describes the content of the build specification (build spec) for the Amplify app.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 25000.

Pattern: `(?s).+`

Required: No

**cacheConfig**

The cache configuration for the Amplify app. If you don't specify the
cache configuration `type`, Amplify uses the default
`AMPLIFY_MANAGED` setting.

Type: [CacheConfig](api-cacheconfig.md) object

Required: No

**computeRoleArn**

The Amazon Resource Name (ARN) of the IAM role for an SSR app. The
Compute role allows the Amplify Hosting compute service to securely
access specific AWS resources based on the role's permissions. For more
information about the SSR Compute role, see [Adding an SSR Compute\
role](../../../../services/amplify/latest/userguide/amplify-ssr-compute-role.md) in the _Amplify User Guide_.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1000.

Pattern: `(?s).*`

Required: No

**customHeaders**

Describes the custom HTTP headers for the Amplify app.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 25000.

Pattern: `(?s).*`

Required: No

**customRules**

Describes the custom redirect and rewrite rules for the Amplify app.

Type: Array of [CustomRule](api-customrule.md) objects

Required: No

**enableAutoBranchCreation**

Enables automated branch creation for the Amplify app.

Type: Boolean

Required: No

**enableBranchAutoDeletion**

Automatically disconnect a branch in the Amplify console when you delete a branch from
your Git repository.

Type: Boolean

Required: No

**iamServiceRoleArn**

The Amazon Resource Name (ARN) of the IAM service role for the Amplify app.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1000.

Pattern: `(?s).*`

Required: No

**jobConfig**

The configuration details that apply to the jobs for an Amplify app.

Type: [JobConfig](api-jobconfig.md) object

Required: No

**productionBranch**

Describes the information about a production branch of the Amplify app.

Type: [ProductionBranch](api-productionbranch.md) object

Required: No

**repositoryCloneMethod**

###### Note

This is for internal use.

The Amplify service uses this parameter to specify the authentication protocol to use
to access the Git repository for an Amplify app. Amplify specifies `TOKEN`
for a GitHub repository, `SIGV4` for an AWS CodeCommit
repository, and `SSH` for GitLab and Bitbucket repositories.

Type: String

Valid Values: `SSH | TOKEN | SIGV4`

Required: No

**tags**

The tag for the Amplify app.

Type: String to string map

Map Entries: Minimum number of 0 items. Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `^(?!aws:)[a-zA-Z+-=._:/]+$`

Value Length Constraints: Maximum length of 256.

Value Pattern: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

Required: No

**wafConfiguration**

Describes the Firewall configuration for the Amplify app. Firewall
support enables you to protect your hosted applications with a direct integration with
AWS WAF.

Type: [WafConfiguration](api-wafconfiguration.md) object

Required: No

**webhookCreateTime**

A timestamp of when Amplify created the webhook in your Git
repository.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/amplify-2017-07-25/app.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/amplify-2017-07-25/app.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/amplify-2017-07-25/app.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Data Types

Artifact

All content copied from https://docs.aws.amazon.com/.
