This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Amplify::App AutoBranchCreationConfig

Use the AutoBranchCreationConfig property type to automatically create branches that
match a certain pattern.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoBranchCreationPatterns" : [ String, ... ],
  "BasicAuthConfig" : BasicAuthConfig,
  "BuildSpec" : String,
  "EnableAutoBranchCreation" : Boolean,
  "EnableAutoBuild" : Boolean,
  "EnablePerformanceMode" : Boolean,
  "EnablePullRequestPreview" : Boolean,
  "EnvironmentVariables" : [ EnvironmentVariable, ... ],
  "Framework" : String,
  "PullRequestEnvironmentName" : String,
  "Stage" : String
}

```

### YAML

```yaml

  AutoBranchCreationPatterns:
    - String
  BasicAuthConfig:
    BasicAuthConfig
  BuildSpec: String
  EnableAutoBranchCreation: Boolean
  EnableAutoBuild: Boolean
  EnablePerformanceMode: Boolean
  EnablePullRequestPreview: Boolean
  EnvironmentVariables:
    - EnvironmentVariable
  Framework: String
  PullRequestEnvironmentName: String
  Stage: String

```

## Properties

`AutoBranchCreationPatterns`

Automated branch creation glob patterns for the Amplify app.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BasicAuthConfig`

Sets password protection for your auto created branch.

_Required_: No

_Type_: [BasicAuthConfig](aws-properties-amplify-app-basicauthconfig.md)

_Pattern_: `(?s).*`

_Maximum_: `2000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BuildSpec`

The build specification (build spec) for the autocreated branch.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `25000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableAutoBranchCreation`

Enables automated branch creation for the Amplify app.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableAutoBuild`

Enables auto building for the auto created branch.

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

Sets whether pull request previews are enabled for each branch that Amplify Hosting
automatically creates for your app. Amplify creates previews by deploying your app to a unique
URL whenever a pull request is opened for the branch. Development and QA teams can use this
preview to test the pull request before it's merged into a production or integration
branch.

To provide backend support for your preview, Amplify Hosting automatically provisions a
temporary backend environment that it deletes when the pull request is closed. If you want to
specify a dedicated backend environment for your previews, use the
`PullRequestEnvironmentName` property.

For more information, see [Web Previews](../../../amplify/latest/userguide/pr-previews.md) in the
_AWS Amplify Hosting User Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentVariables`

The environment variables for the autocreated branch.

_Required_: No

_Type_: Array of [EnvironmentVariable](aws-properties-amplify-app-environmentvariable.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Framework`

The framework for the autocreated branch.

_Required_: No

_Type_: String

_Pattern_: `(?s).*`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PullRequestEnvironmentName`

If pull request previews are enabled, you can use this property to specify a dedicated
backend environment for your previews. For example, you could specify an environment named
`prod`, `test`, or `dev` that you initialized with the
Amplify CLI.

To enable pull request previews, set the `EnablePullRequestPreview` property
to `true`.

If you don't specify an environment, Amplify Hosting provides backend support for each
preview by automatically provisioning a temporary backend environment. Amplify deletes this
environment when the pull request is closed.

For more information about creating backend environments, see [Feature Branch\
Deployments and Team Workflows](../../../amplify/latest/userguide/multi-environments.md) in the _AWS Amplify Hosting_
_User Guide_.

_Required_: No

_Type_: String

_Pattern_: `(?s).*`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Stage`

Stage for the auto created branch.

_Required_: No

_Type_: String

_Allowed values_: `EXPERIMENTAL | BETA | PULL_REQUEST | PRODUCTION | DEVELOPMENT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Amplify::App

BasicAuthConfig

All content copied from https://docs.aws.amazon.com/.
