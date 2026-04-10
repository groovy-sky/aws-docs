This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeBuild::Project BuildStatusConfig

Contains information that defines how the AWS CodeBuild build project reports the build status
to the source provider.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Context" : String,
  "TargetUrl" : String
}

```

### YAML

```yaml

  Context: String
  TargetUrl: String

```

## Properties

`Context`

Specifies the context of the build status CodeBuild sends to the source provider. The
usage of this parameter depends on the source provider.

Bitbucket

This parameter is used for the `name` parameter in the
Bitbucket commit status. For more information, see [build](https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Bworkspace%7D/%7Brepo_slug%7D/commit/%7Bnode%7D/statuses/build) in the Bitbucket API documentation.

GitHub/GitHub Enterprise Server

This parameter is used for the `context` parameter in the
GitHub commit status. For more information, see [Create a commit status](https://developer.github.com/v3/repos/statuses) in the GitHub developer guide.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetUrl`

Specifies the target url of the build status CodeBuild sends to the source provider. The
usage of this parameter depends on the source provider.

Bitbucket

This parameter is used for the `url` parameter in the Bitbucket
commit status. For more information, see [build](https://developer.atlassian.com/bitbucket/api/2/reference/resource/repositories/%7Bworkspace%7D/%7Brepo_slug%7D/commit/%7Bnode%7D/statuses/build) in the Bitbucket API documentation.

GitHub/GitHub Enterprise Server

This parameter is used for the `target_url` parameter in the
GitHub commit status. For more information, see [Create a commit status](https://developer.github.com/v3/repos/statuses) in the GitHub developer guide.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BatchRestrictions

CloudWatchLogsConfig

All content copied from https://docs.aws.amazon.com/.
