This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Crawler LakeFormationConfiguration

Specifies AWS Lake Formation configuration settings for the crawler.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccountId" : String,
  "UseLakeFormationCredentials" : Boolean
}

```

### YAML

```yaml

  AccountId: String
  UseLakeFormationCredentials: Boolean

```

## Properties

`AccountId`

Required for cross account crawls. For same account crawls as the target data, this can be left as null.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `12`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseLakeFormationCredentials`

Specifies whether to use AWS Lake Formation credentials for the crawler instead of the IAM role credentials.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

JdbcTarget

MongoDBTarget
