This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Amplify::App

The AWS::Amplify::App resource specifies Apps in Amplify Hosting. An App is a
collection of branches.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Amplify::App",
  "Properties" : {
      "AccessToken" : String,
      "AutoBranchCreationConfig" : AutoBranchCreationConfig,
      "BasicAuthConfig" : BasicAuthConfig,
      "BuildSpec" : String,
      "CacheConfig" : CacheConfig,
      "ComputeRoleArn" : String,
      "CustomHeaders" : String,
      "CustomRules" : [ CustomRule, ... ],
      "Description" : String,
      "EnableBranchAutoDeletion" : Boolean,
      "EnvironmentVariables" : [ EnvironmentVariable, ... ],
      "IAMServiceRole" : String,
      "JobConfig" : JobConfig,
      "Name" : String,
      "OauthToken" : String,
      "Platform" : String,
      "Repository" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Amplify::App
Properties:
  AccessToken: String
  AutoBranchCreationConfig:
    AutoBranchCreationConfig
  BasicAuthConfig:
    BasicAuthConfig
  BuildSpec: String
  CacheConfig:
    CacheConfig
  ComputeRoleArn: String
  CustomHeaders: String
  CustomRules:
    - CustomRule
  Description: String
  EnableBranchAutoDeletion: Boolean
  EnvironmentVariables:
    - EnvironmentVariable
  IAMServiceRole: String
  JobConfig:
    JobConfig
  Name: String
  OauthToken: String
  Platform: String
  Repository: String
  Tags:
    - Tag

```

## Properties

`AccessToken`

The personal access token for a GitHub repository for an Amplify app. The personal access
token is used to authorize access to a GitHub repository using the Amplify GitHub App. The
token is not stored.

Use `AccessToken` for GitHub repositories only. To authorize access to a
repository provider such as Bitbucket or CodeCommit, use `OauthToken`.

You must specify either `AccessToken` or `OauthToken` when you
create a new app.

Existing Amplify apps deployed from a GitHub repository using OAuth continue to work with
CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub App. For
more information, see [Migrating an existing OAuth app to the Amplify GitHub App](../../../amplify/latest/userguide/setting-up-github-access.md#migrating-to-github-app-auth) in the _Amplify_
_User Guide_ .

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutoBranchCreationConfig`

Sets the configuration for your automatic branch creation.

_Required_: No

_Type_: [AutoBranchCreationConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplify-app-autobranchcreationconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BasicAuthConfig`

The credentials for basic authorization for an Amplify app. You must base64-encode the
authorization credentials and provide them in the format
`user:password`.

_Required_: No

_Type_: [BasicAuthConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplify-app-basicauthconfig.html)

_Pattern_: `(?s).*`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BuildSpec`

The build specification (build spec) for an Amplify app.

_Required_: No

_Type_: String

_Pattern_: `(?s).+`

_Minimum_: `1`

_Maximum_: `25000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CacheConfig`

The cache configuration for the Amplify app. If you don't specify the
cache configuration `type`, Amplify uses the default
`AMPLIFY_MANAGED` setting.

_Required_: No

_Type_: [CacheConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplify-app-cacheconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComputeRoleArn`

The Amazon Resource Name (ARN) of the IAM role for an SSR app. The
Compute role allows the Amplify Hosting compute service to securely
access specific AWS resources based on the role's permissions. For more
information about the SSR Compute role, see [Adding an SSR Compute\
role](../../../amplify/latest/userguide/amplify-ssr-compute-role.md) in the _Amplify User Guide_.

_Required_: No

_Type_: String

_Pattern_: `(?s).*`

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomHeaders`

The custom HTTP headers for an Amplify app.

_Required_: No

_Type_: String

_Pattern_: `(?s).*`

_Minimum_: `0`

_Maximum_: `25000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomRules`

The custom rewrite and redirect rules for an Amplify app.

_Required_: No

_Type_: Array of [CustomRule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplify-app-customrule.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the Amplify app.

_Required_: No

_Type_: String

_Pattern_: `(?s).*`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableBranchAutoDeletion`

Automatically disconnect a branch in Amplify Hosting when you delete a branch from your
Git repository.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentVariables`

The environment variables for the Amplify app.

For a list of the environment variables that are accessible to Amplify by default, see
[Amplify\
Environment variables](https://docs.aws.amazon.com/amplify/latest/userguide/amplify-console-environment-variables.html) in the _Amplify Hosting User_
_Guide_.

_Required_: No

_Type_: Array of [EnvironmentVariable](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplify-app-environmentvariable.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IAMServiceRole`

AWS Identity and Access Management (IAM) service role for the Amazon Resource Name
(ARN) of the Amplify app.

_Required_: No

_Type_: String

_Pattern_: `(?s).*`

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JobConfig`

The configuration details that apply to the jobs for an Amplify app.

_Required_: No

_Type_: [JobConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplify-app-jobconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the Amplify app.

_Required_: Yes

_Type_: String

_Pattern_: `(?s).+`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OauthToken`

The OAuth token for a third-party source control system for an Amplify app. The OAuth
token is used to create a webhook and a read-only deploy key using SSH cloning. The OAuth
token is not stored.

Use `OauthToken` for repository providers other than GitHub, such as Bitbucket
or CodeCommit. To authorize access to GitHub as your repository provider, use
`AccessToken`.

You must specify either `OauthToken` or `AccessToken` when you
create a new app.

Existing Amplify apps deployed from a GitHub repository using OAuth continue to work with
CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub App. For
more information, see [Migrating an existing OAuth app to the Amplify GitHub App](../../../amplify/latest/userguide/setting-up-github-access.md#migrating-to-github-app-auth) in the _Amplify_
_User Guide_ .

_Required_: No

_Type_: String

_Pattern_: `(?s).*`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Platform`

The platform for the Amplify app. For a static app, set the platform type to
`WEB`. For a dynamic server-side rendered (SSR) app, set the platform
type to `WEB_COMPUTE`. For an app requiring Amplify Hosting's original SSR
support only, set the platform type to `WEB_DYNAMIC`.

If you are deploying an SSG only app with Next.js version 14 or later, you must set
the platform type to `WEB_COMPUTE` and set the artifacts
`baseDirectory` to `.next` in the application's build
settings. For an example of the build specification settings, see [Amplify build settings for a Next.js 14 SSG application](https://docs.aws.amazon.com/amplify/latest/userguide/deploy-nextjs-app.html#build-setting-detection-ssg-14) in the
_Amplify Hosting User Guide_.

_Required_: No

_Type_: String

_Allowed values_: `WEB | WEB_DYNAMIC | WEB_COMPUTE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Repository`

The Git repository for the Amplify app.

_Required_: No

_Type_: String

_Pattern_: `(?s).*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tag for an Amplify app.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplify-app-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AppId`

Unique Id for the Amplify App.

`AppName`

Name for the Amplify App.

`Arn`

ARN for the Amplify App.

`DefaultDomain`

Default domain for the Amplify App.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Amplify Console

AutoBranchCreationConfig
