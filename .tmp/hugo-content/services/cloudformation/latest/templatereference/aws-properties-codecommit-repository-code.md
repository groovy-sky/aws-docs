This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeCommit::Repository Code

Information about code to be committed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BranchName" : String,
  "S3" : S3
}

```

### YAML

```yaml

  BranchName: String
  S3:
    S3

```

## Properties

`BranchName`

Optional. Specifies a branch name to be used as the default branch when importing code
into a repository on initial creation. If this property is not set, the name
_main_ will be used for the default branch for the repository.
Changes to this property are ignored after initial resource creation. We recommend using
this parameter to set the name to _main_ to align with the default
behavior of CodeCommit unless another name is needed.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3`

Information about the Amazon S3 bucket that contains a ZIP file of code to
be committed to the repository. Changes to this property are ignored after initial
resource creation.

_Required_: Yes

_Type_: [S3](aws-properties-codecommit-repository-s3.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CodeCommit::Repository

RepositoryTrigger

All content copied from https://docs.aws.amazon.com/.
