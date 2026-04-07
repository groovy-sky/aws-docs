This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeStar::GitHubRepository Code

The `Code` property type specifies information about code to be
committed.

`Code` is a property of the `AWS::CodeStar::GitHubRepository`
resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3" : S3
}

```

### YAML

```yaml

  S3:
    S3

```

## Properties

`S3`

Information about the Amazon S3 bucket that contains a ZIP file of code to be
committed to the repository.

_Required_: Yes

_Type_: [S3](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-codestar-githubrepository-s3.html)

_Update requires_: Updates are not supported.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::CodeStar::GitHubRepository

S3
