This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3Outposts::Bucket LifecycleConfiguration

The container for the lifecycle configuration for the objects stored in an S3 on Outposts bucket.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Rules" : [ Rule, ... ]
}

```

### YAML

```yaml

  Rules:
    - Rule

```

## Properties

`Rules`

The container for the lifecycle configuration rules for the objects stored in the S3 on Outposts bucket.

_Required_: Yes

_Type_: Array of [Rule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-s3outposts-bucket-rule.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FilterTag

Rule
