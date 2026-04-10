This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeBuild::Project Source

`Source` is a property of the [AWS::CodeBuild::Project](../userguide/aws-resource-codebuild-project.md) resource that specifies
the source code settings for the project, such as the source code's repository type and location.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Auth" : SourceAuth,
  "BuildSpec" : String,
  "BuildStatusConfig" : BuildStatusConfig,
  "GitCloneDepth" : Integer,
  "GitSubmodulesConfig" : GitSubmodulesConfig,
  "InsecureSsl" : Boolean,
  "Location" : String,
  "ReportBuildStatus" : Boolean,
  "SourceIdentifier" : String,
  "Type" : String
}

```

### YAML

```yaml

  Auth:
    SourceAuth
  BuildSpec: String
  BuildStatusConfig:
    BuildStatusConfig
  GitCloneDepth: Integer
  GitSubmodulesConfig:
    GitSubmodulesConfig
  InsecureSsl: Boolean
  Location: String
  ReportBuildStatus: Boolean
  SourceIdentifier: String
  Type: String

```

## Properties

`Auth`

Information about the authorization settings for AWS CodeBuild to access the source code to be
built.

_Required_: No

_Type_: [SourceAuth](aws-properties-codebuild-project-sourceauth.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BuildSpec`

The build specification for the project. If this value is not provided, then the source code must contain a
buildspec file named `buildspec.yml` at the root level. If this value is provided, it can be either
a single string containing the entire build specification, or the path to an alternate buildspec file relative
to the value of the built-in environment variable `CODEBUILD_SRC_DIR`. The alternate buildspec file
can have a name other than `buildspec.yml`, for example `myspec.yml` or
`build_spec_qa.yml` or similar. For more information, see the [Build Spec Reference](../../../codebuild/latest/userguide/build-spec-ref.md#build-spec-ref-example) in the
_AWS CodeBuild User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BuildStatusConfig`

Contains information that defines how the build project reports the build status to
the source provider. This option is only used when the source provider is
`GITHUB`, `GITHUB_ENTERPRISE`, or
`BITBUCKET`.

_Required_: No

_Type_: [BuildStatusConfig](aws-properties-codebuild-project-buildstatusconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GitCloneDepth`

The depth of history to download. Minimum value is 0. If this value is 0,
greater than 25, or not provided, then the full history is downloaded with each build project.
If your source type is Amazon S3, this value is not supported.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GitSubmodulesConfig`

Information about the Git submodules configuration for the build project.

_Required_: No

_Type_: [GitSubmodulesConfig](aws-properties-codebuild-project-gitsubmodulesconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InsecureSsl`

This is used with GitHub Enterprise only. Set to true to ignore SSL warnings while connecting to your GitHub Enterprise project repository.
The default value is `false`. `InsecureSsl` should be used for testing purposes only. It should not be used in a production environment.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Location`

Information about the location of the source code to be built. Valid values
include:

- For source code settings that are specified in the source action of a pipeline
in CodePipeline, `location` should not be specified. If it is specified,
CodePipeline ignores it. This is because CodePipeline uses the settings in a pipeline's source
action instead of this value.

- For source code in an CodeCommit repository, the HTTPS clone URL to the repository
that contains the source code and the buildspec file (for example,
`https://git-codecommit.<region-ID>.amazonaws.com/v1/repos/<repo-name>`).

- For source code in an Amazon S3 input bucket, one of the following.

- The path to the ZIP file that contains the source code (for example,
`<bucket-name>/<path>/<object-name>.zip`).

- The path to the folder that contains the source code (for example,
`<bucket-name>/<path-to-source-code>/<folder>/`).

- For source code in a GitHub repository, the HTTPS clone URL to the repository
that contains the source and the buildspec file. You must connect your AWS account
to your GitHub account. Use the AWS CodeBuild console to start creating a build
project. When you use the console to connect (or reconnect) with GitHub, on the
GitHub **Authorize application** page, for
**Organization access**, choose **Request access** next to each repository you want to
allow AWS CodeBuild to have access to, and then choose **Authorize**
**application**. (After you have connected to your GitHub account,
you do not need to finish creating the build project. You can leave the AWS CodeBuild
console.) To instruct AWS CodeBuild to use this connection, in the `source`
object, set the `auth` object's `type` value to
`OAUTH`.

- For source code in an GitLab or self-managed GitLab repository, the HTTPS clone URL to the repository
that contains the source and the buildspec file. You must connect your AWS account
to your GitLab account. Use the AWS CodeBuild console to start creating a build
project. When you use the console to connect (or reconnect) with GitLab, on the
Connections **Authorize application** page, choose **Authorize**. Then on the AWS CodeConnections **Create GitLab connection** page,
choose **Connect to GitLab**. (After you have connected to your GitLab account,
you do not need to finish creating the build project. You can leave the AWS CodeBuild
console.) To instruct AWS CodeBuild to override the default connection and use this connection instead,
set the `auth` object's `type` value to
`CODECONNECTIONS` in the `source` object.

- For source code in a Bitbucket repository, the HTTPS clone URL to the
repository that contains the source and the buildspec file. You must connect
your AWS account to your Bitbucket account. Use the AWS CodeBuild console to start
creating a build project. When you use the console to connect (or reconnect)
with Bitbucket, on the Bitbucket **Confirm access to your**
**account** page, choose **Grant**
**access**. (After you have connected to your Bitbucket account, you
do not need to finish creating the build project. You can leave the AWS CodeBuild
console.) To instruct AWS CodeBuild to use this connection, in the `source`
object, set the `auth` object's `type` value to
`OAUTH`.

If you specify `CODEPIPELINE` for the `Type` property, don't specify this
property. For all of the other types, you must specify `Location`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReportBuildStatus`

Set to true to report the status of a build's start and finish to your source provider.
This option is valid only when your source provider is GitHub, GitHub Enterprise, GitLab, GitLab Self Managed, or
Bitbucket. If this is set and you use a different source provider, an `invalidInputException`
is thrown.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceIdentifier`

An identifier for this project source. The identifier can only contain
alphanumeric characters and underscores, and must be less than 128 characters in length.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of repository that contains the source code to be built. Valid values
include:

- `BITBUCKET`: The source code is in a Bitbucket repository.

- `CODECOMMIT`: The source code is in an CodeCommit repository.

- `CODEPIPELINE`: The source code settings are specified in the
source action of a pipeline in CodePipeline.

- `GITHUB`: The source code is in a GitHub repository.

- `GITHUB_ENTERPRISE`: The source code is in a GitHub Enterprise
Server repository.

- `GITLAB`: The source code is in a GitLab repository.

- `GITLAB_SELF_MANAGED`: The source code is in a self-managed GitLab repository.

- `NO_SOURCE`: The project does not have input source code.

- `S3`: The source code is in an Amazon S3 bucket.

_Required_: Yes

_Type_: String

_Allowed values_: `CODECOMMIT | CODEPIPELINE | GITHUB | GITLAB | GITLAB_SELF_MANAGED | S3 | BITBUCKET | GITHUB_ENTERPRISE | NO_SOURCE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScopeConfiguration

SourceAuth

All content copied from https://docs.aws.amazon.com/.
