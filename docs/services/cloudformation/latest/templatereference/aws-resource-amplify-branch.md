This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Amplify::Branch

The AWS::Amplify::Branch resource specifies a new branch within an app.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Amplify::Branch",
  "Properties" : {
      "AppId" : String,
      "Backend" : Backend,
      "BasicAuthConfig" : BasicAuthConfig,
      "BranchName" : String,
      "BuildSpec" : String,
      "ComputeRoleArn" : String,
      "Description" : String,
      "EnableAutoBuild" : Boolean,
      "EnablePerformanceMode" : Boolean,
      "EnablePullRequestPreview" : Boolean,
      "EnableSkewProtection" : Boolean,
      "EnvironmentVariables" : [ EnvironmentVariable, ... ],
      "Framework" : String,
      "PullRequestEnvironmentName" : String,
      "Stage" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Amplify::Branch
Properties:
  AppId: String
  Backend:
    Backend
  BasicAuthConfig:
    BasicAuthConfig
  BranchName: String
  BuildSpec: String
  ComputeRoleArn: String
  Description: String
  EnableAutoBuild: Boolean
  EnablePerformanceMode: Boolean
  EnablePullRequestPreview: Boolean
  EnableSkewProtection: Boolean
  EnvironmentVariables:
    - EnvironmentVariable
  Framework: String
  PullRequestEnvironmentName: String
  Stage: String
  Tags:
    - Tag

```

## Properties

`AppId`

The unique ID for an Amplify app.

_Required_: Yes

_Type_: String

_Pattern_: `d[a-z0-9]+`

_Minimum_: `1`

_Maximum_: `20`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Backend`

The backend for a `Branch` of an Amplify app. Use for a
backend created from an CloudFormation stack.

This field is available to Amplify Gen 2 apps only. When you deploy an
application with Amplify Gen 2, you provision the app's backend infrastructure using
Typescript code.

_Required_: No

_Type_: [Backend](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplify-branch-backend.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BasicAuthConfig`

The basic authorization credentials for a branch of an Amplify app. You must
base64-encode the authorization credentials and provide them in the format
`user:password`.

_Required_: No

_Type_: [BasicAuthConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplify-branch-basicauthconfig.html)

_Pattern_: `(?s).*`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BranchName`

The name for the branch.

_Required_: Yes

_Type_: String

_Pattern_: `(?s).+`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BuildSpec`

The build specification (build spec) for the branch.

_Required_: No

_Type_: String

_Pattern_: `(?s).+`

_Minimum_: `1`

_Maximum_: `25000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComputeRoleArn`

The Amazon Resource Name (ARN) of the IAM role to assign to a branch of
an SSR app. The SSR Compute role allows the Amplify Hosting compute
service to securely access specific AWS resources based on the role's
permissions. For more information about the SSR Compute role, see [Adding an SSR Compute role](../../../amplify/latest/userguide/amplify-ssr-compute-role.md) in the _Amplify User_
_Guide_.

_Required_: No

_Type_: String

_Pattern_: `(?s).*`

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description for the branch that is part of an Amplify app.

_Required_: No

_Type_: String

_Pattern_: `(?s).*`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableAutoBuild`

Enables auto building for the branch.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnablePerformanceMode`

Enables performance mode for the branch.

Performance mode optimizes for faster hosting performance by keeping content cached at
the edge for a longer interval. When performance mode is enabled, hosting configuration
or code changes can take up to 10 minutes to roll out.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnablePullRequestPreview`

Specifies whether Amplify Hosting creates a preview for each pull request that is made
for this branch. If this property is enabled, Amplify deploys your app to a unique preview URL
after each pull request is opened. Development and QA teams can use this preview to test the
pull request before it's merged into a production or integration branch.

To provide backend support for your preview, Amplify automatically provisions a
temporary backend environment that it deletes when the pull request is closed. If you want to
specify a dedicated backend environment for your previews, use the
`PullRequestEnvironmentName` property.

For more information, see [Web Previews](../../../amplify/latest/userguide/pr-previews.md) in the
_AWS Amplify Hosting User Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableSkewProtection`

Specifies whether the skew protection feature is enabled for the branch.

Deployment skew protection is available to Amplify applications to
eliminate version skew issues between client and servers in web applications. When you
apply skew protection to a branch, you can ensure that your clients always interact with
the correct version of server-side assets, regardless of when a deployment occurs. For
more information about skew protection, see [Skew protection for Amplify deployments](https://docs.aws.amazon.com/amplify/latest/userguide/skew-protection.html) in the _Amplify User_
_Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentVariables`

The environment variables for the branch.

_Required_: No

_Type_: Array of [EnvironmentVariable](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplify-branch-environmentvariable.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Framework`

The framework for the branch.

_Required_: No

_Type_: String

_Pattern_: `(?s).*`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PullRequestEnvironmentName`

If pull request previews are enabled for this branch, you can use this property to
specify a dedicated backend environment for your previews. For example, you could specify an
environment named `prod`, `test`, or `dev` that you
initialized with the Amplify CLI and mapped to this branch.

To enable pull request previews, set the `EnablePullRequestPreview` property
to `true`.

If you don't specify an environment, Amplify Hosting provides backend support for each
preview by automatically provisioning a temporary backend environment. Amplify Hosting deletes
this environment when the pull request is closed.

For more information about creating backend environments, see [Feature Branch\
Deployments and Team Workflows](../../../amplify/latest/userguide/multi-environments.md) in the _AWS Amplify Hosting_
_User Guide_.

_Required_: No

_Type_: String

_Pattern_: `(?s).*`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Stage`

Describes the current stage for the branch.

_Required_: No

_Type_: String

_Allowed values_: `EXPERIMENTAL | BETA | PULL_REQUEST | PRODUCTION | DEVELOPMENT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tag for the branch.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-amplify-branch-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

ARN for a branch, part of an Amplify App.

`BranchName`

Name for a branch, part of an Amplify App.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Backend
