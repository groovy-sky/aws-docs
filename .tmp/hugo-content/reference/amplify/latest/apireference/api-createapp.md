# CreateApp

Creates a new Amplify app.

## Request Syntax

```nohighlight

POST /apps HTTP/1.1
Content-type: application/json

{
   "accessToken": "string",
   "autoBranchCreationConfig": {
      "basicAuthCredentials": "string",
      "buildSpec": "string",
      "enableAutoBuild": boolean,
      "enableBasicAuth": boolean,
      "enablePerformanceMode": boolean,
      "enablePullRequestPreview": boolean,
      "environmentVariables": {
         "string" : "string"
      },
      "framework": "string",
      "pullRequestEnvironmentName": "string",
      "stage": "string"
   },
   "autoBranchCreationPatterns": [ "string" ],
   "basicAuthCredentials": "string",
   "buildSpec": "string",
   "cacheConfig": {
      "type": "string"
   },
   "computeRoleArn": "string",
   "customHeaders": "string",
   "customRules": [
      {
         "condition": "string",
         "source": "string",
         "status": "string",
         "target": "string"
      }
   ],
   "description": "string",
   "enableAutoBranchCreation": boolean,
   "enableBasicAuth": boolean,
   "enableBranchAutoBuild": boolean,
   "enableBranchAutoDeletion": boolean,
   "environmentVariables": {
      "string" : "string"
   },
   "iamServiceRoleArn": "string",
   "jobConfig": {
      "buildComputeType": "string"
   },
   "name": "string",
   "oauthToken": "string",
   "platform": "string",
   "repository": "string",
   "tags": {
      "string" : "string"
   }
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[accessToken](#API_CreateApp_RequestSyntax)**

The personal access token for a GitHub repository for an Amplify app. The personal
access token is used to authorize access to a GitHub repository using the Amplify GitHub
App. The token is not stored.

Use `accessToken` for GitHub repositories only. To authorize access to a
repository provider such as Bitbucket or CodeCommit, use `oauthToken`.

You must specify either `accessToken` or `oauthToken` when you
create a new app.

Existing Amplify apps deployed from a GitHub repository using OAuth continue to work
with CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub
App. For more information, see [Migrating an existing OAuth app to the Amplify GitHub App](../../../../services/amplify/latest/userguide/setting-up-github-access.md#migrating-to-github-app-auth) in the
_Amplify User Guide_ .

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `(?s).+`

Required: No

**[autoBranchCreationConfig](#API_CreateApp_RequestSyntax)**

The automated branch creation configuration for an Amplify app.

Type: [AutoBranchCreationConfig](api-autobranchcreationconfig.md) object

Required: No

**[autoBranchCreationPatterns](#API_CreateApp_RequestSyntax)**

The automated branch creation glob patterns for an Amplify app.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `(?s).+`

Required: No

**[basicAuthCredentials](#API_CreateApp_RequestSyntax)**

The credentials for basic authorization for an Amplify app. You must base64-encode the
authorization credentials and provide them in the format
`user:password`.

Type: String

Length Constraints: Maximum length of 2000.

Pattern: `(?s).*`

Required: No

**[buildSpec](#API_CreateApp_RequestSyntax)**

The build specification (build spec) for an Amplify app.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 25000.

Pattern: `(?s).+`

Required: No

**[cacheConfig](#API_CreateApp_RequestSyntax)**

The cache configuration for the Amplify app.

Type: [CacheConfig](api-cacheconfig.md) object

Required: No

**[computeRoleArn](#API_CreateApp_RequestSyntax)**

The Amazon Resource Name (ARN) of the IAM role to assign to an SSR app.
The SSR Compute role allows the Amplify Hosting compute service to
securely access specific AWS resources based on the role's permissions.
For more information about the SSR Compute role, see [Adding an SSR Compute\
role](../../../../services/amplify/latest/userguide/amplify-ssr-compute-role.md) in the _Amplify User Guide_.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1000.

Pattern: `(?s).*`

Required: No

**[customHeaders](#API_CreateApp_RequestSyntax)**

The custom HTTP headers for an Amplify app.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 25000.

Pattern: `(?s).*`

Required: No

**[customRules](#API_CreateApp_RequestSyntax)**

The custom rewrite and redirect rules for an Amplify app.

Type: Array of [CustomRule](api-customrule.md) objects

Required: No

**[description](#API_CreateApp_RequestSyntax)**

The description of the Amplify app.

Type: String

Length Constraints: Maximum length of 1000.

Pattern: `(?s).*`

Required: No

**[enableAutoBranchCreation](#API_CreateApp_RequestSyntax)**

Enables automated branch creation for an Amplify app.

Type: Boolean

Required: No

**[enableBasicAuth](#API_CreateApp_RequestSyntax)**

Enables basic authorization for an Amplify app. This will apply to all branches that
are part of this app.

Type: Boolean

Required: No

**[enableBranchAutoBuild](#API_CreateApp_RequestSyntax)**

Enables the auto building of branches for an Amplify app.

Type: Boolean

Required: No

**[enableBranchAutoDeletion](#API_CreateApp_RequestSyntax)**

Automatically disconnects a branch in the Amplify console when you delete a branch
from your Git repository.

Type: Boolean

Required: No

**[environmentVariables](#API_CreateApp_RequestSyntax)**

The environment variables map for an Amplify app.

For a list of the environment variables that are accessible to Amplify by default, see
[Amplify\
Environment variables](../../../../services/amplify/latest/userguide/amplify-console-environment-variables.md) in the _Amplify Hosting User_
_Guide_.

Type: String to string map

Key Length Constraints: Maximum length of 255.

Key Pattern: `(?s).*`

Value Length Constraints: Maximum length of 5500.

Value Pattern: `(?s).*`

Required: No

**[iamServiceRoleArn](#API_CreateApp_RequestSyntax)**

The Amazon Resource Name (ARN) of the IAM service role for the Amplify app.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1000.

Pattern: `(?s).*`

Required: No

**[jobConfig](#API_CreateApp_RequestSyntax)**

Describes the configuration details that apply to the jobs for an Amplify app.

Type: [JobConfig](api-jobconfig.md) object

Required: No

**[name](#API_CreateApp_RequestSyntax)**

The name of the Amplify app.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `(?s).+`

Required: Yes

**[oauthToken](#API_CreateApp_RequestSyntax)**

The OAuth token for a third-party source control system for an Amplify app. The OAuth
token is used to create a webhook and a read-only deploy key using SSH cloning. The
OAuth token is not stored.

Use `oauthToken` for repository providers other than GitHub, such as
Bitbucket or CodeCommit. To authorize access to GitHub as your repository provider, use
`accessToken`.

You must specify either `oauthToken` or `accessToken` when you
create a new app.

Existing Amplify apps deployed from a GitHub repository using OAuth continue to work
with CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub
App. For more information, see [Migrating an existing OAuth app to the Amplify GitHub App](../../../../services/amplify/latest/userguide/setting-up-github-access.md#migrating-to-github-app-auth) in the
_Amplify User Guide_ .

Type: String

Length Constraints: Maximum length of 1000.

Pattern: `(?s).*`

Required: No

**[platform](#API_CreateApp_RequestSyntax)**

The platform for the Amplify app. For a static app, set the platform type to
`WEB`. For a dynamic server-side rendered (SSR) app, set the platform
type to `WEB_COMPUTE`. For an app requiring Amplify Hosting's original SSR
support only, set the platform type to `WEB_DYNAMIC`.

If you are deploying an SSG only app with Next.js version 14 or later, you must set
the platform type to `WEB_COMPUTE` and set the artifacts
`baseDirectory` to `.next` in the application's build
settings. For an example of the build specification settings, see [Amplify build settings for a Next.js 14 SSG application](../../../../services/amplify/latest/userguide/deploy-nextjs-app.md#build-setting-detection-ssg-14) in the
_Amplify Hosting User Guide_.

Type: String

Valid Values: `WEB | WEB_DYNAMIC | WEB_COMPUTE`

Required: No

**[repository](#API_CreateApp_RequestSyntax)**

The Git repository for the Amplify app.

Type: String

Length Constraints: Maximum length of 1000.

Pattern: `(?s).*`

Required: No

**[tags](#API_CreateApp_RequestSyntax)**

The tag for an Amplify app.

Type: String to string map

Map Entries: Minimum number of 0 items. Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `^(?!aws:)[a-zA-Z+-=._:/]+$`

Value Length Constraints: Maximum length of 256.

Value Pattern: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "app": {
      "appArn": "string",
      "appId": "string",
      "autoBranchCreationConfig": {
         "basicAuthCredentials": "string",
         "buildSpec": "string",
         "enableAutoBuild": boolean,
         "enableBasicAuth": boolean,
         "enablePerformanceMode": boolean,
         "enablePullRequestPreview": boolean,
         "environmentVariables": {
            "string" : "string"
         },
         "framework": "string",
         "pullRequestEnvironmentName": "string",
         "stage": "string"
      },
      "autoBranchCreationPatterns": [ "string" ],
      "basicAuthCredentials": "string",
      "buildSpec": "string",
      "cacheConfig": {
         "type": "string"
      },
      "computeRoleArn": "string",
      "createTime": number,
      "customHeaders": "string",
      "customRules": [
         {
            "condition": "string",
            "source": "string",
            "status": "string",
            "target": "string"
         }
      ],
      "defaultDomain": "string",
      "description": "string",
      "enableAutoBranchCreation": boolean,
      "enableBasicAuth": boolean,
      "enableBranchAutoBuild": boolean,
      "enableBranchAutoDeletion": boolean,
      "environmentVariables": {
         "string" : "string"
      },
      "iamServiceRoleArn": "string",
      "jobConfig": {
         "buildComputeType": "string"
      },
      "name": "string",
      "platform": "string",
      "productionBranch": {
         "branchName": "string",
         "lastDeployTime": number,
         "status": "string",
         "thumbnailUrl": "string"
      },
      "repository": "string",
      "repositoryCloneMethod": "string",
      "tags": {
         "string" : "string"
      },
      "updateTime": number,
      "wafConfiguration": {
         "statusReason": "string",
         "wafStatus": "string",
         "webAclArn": "string"
      },
      "webhookCreateTime": number
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[app](#API_CreateApp_ResponseSyntax)**

Represents the different branches of a repository for building, deploying, and hosting
an Amplify app.

Type: [App](api-app.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**BadRequestException**

A request contains unexpected data.

HTTP Status Code: 400

**DependentServiceFailureException**

An operation failed because a dependent service threw an exception.

HTTP Status Code: 503

**InternalFailureException**

The service failed to perform an operation due to an internal issue.

HTTP Status Code: 500

**LimitExceededException**

A resource could not be created because service quotas were exceeded.

HTTP Status Code: 429

**UnauthorizedException**

An operation failed due to a lack of access.

HTTP Status Code: 401

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/amplify-2017-07-25/createapp.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/amplify-2017-07-25/createapp.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/amplify-2017-07-25/createapp.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/amplify-2017-07-25/createapp.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/amplify-2017-07-25/createapp.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/amplify-2017-07-25/createapp.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/amplify-2017-07-25/createapp.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/amplify-2017-07-25/createapp.md)

- [AWS SDK for Python](../../../../services/goto/boto3/amplify-2017-07-25/createapp.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/amplify-2017-07-25/createapp.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Actions

CreateBackendEnvironment

All content copied from https://docs.aws.amazon.com/.
