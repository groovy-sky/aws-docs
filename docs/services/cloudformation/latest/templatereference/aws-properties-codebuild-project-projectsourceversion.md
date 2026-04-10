This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeBuild::Project ProjectSourceVersion

A source identifier and its corresponding version.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SourceIdentifier" : String,
  "SourceVersion" : String
}

```

### YAML

```yaml

  SourceIdentifier: String
  SourceVersion: String

```

## Properties

`SourceIdentifier`

An identifier for a source in the build project. The identifier can only contain
alphanumeric characters and underscores, and must be less than 128 characters in length.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceVersion`

The source version for the corresponding source identifier. If specified, must be one
of:

- For CodeCommit: the commit ID, branch, or Git tag to use.

- For GitHub: the commit ID, pull request ID, branch name, or tag name that
corresponds to the version of the source code you want to build. If a pull
request ID is specified, it must use the format `pr/pull-request-ID`
(for example, `pr/25`). If a branch name is specified, the branch's
HEAD commit ID is used. If not specified, the default branch's HEAD commit ID is
used.

- For GitLab: the commit ID, branch, or Git tag to use.

- For Bitbucket: the commit ID, branch name, or tag name that corresponds to the
version of the source code you want to build. If a branch name is specified, the
branch's HEAD commit ID is used. If not specified, the default branch's HEAD
commit ID is used.

- For Amazon S3: the version ID of the object that represents the build input ZIP
file to use.

For more information, see [Source Version Sample\
with CodeBuild](../../../codebuild/latest/userguide/sample-source-version.md) in the _AWS CodeBuild User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProjectFleet

ProjectTriggers

All content copied from https://docs.aws.amazon.com/.
