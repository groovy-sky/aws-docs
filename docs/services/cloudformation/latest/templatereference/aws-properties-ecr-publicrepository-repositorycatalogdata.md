This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECR::PublicRepository RepositoryCatalogData

The details about the repository that are publicly visible in the Amazon ECR Public
Gallery. For more information, see [Amazon ECR Public\
repository catalog data](https://docs.aws.amazon.com/AmazonECR/latest/public/public-repository-catalog-data.html) in the _Amazon ECR Public User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AboutText" : String,
  "Architectures" : [ String, ... ],
  "OperatingSystems" : [ String, ... ],
  "RepositoryDescription" : String,
  "UsageText" : String
}

```

### YAML

```yaml

  AboutText: String
  Architectures:
    - String
  OperatingSystems:
    - String
  RepositoryDescription: String
  UsageText: String

```

## Properties

`AboutText`

The longform description of the contents of the repository. This text appears in the
repository details on the Amazon ECR Public Gallery.

_Required_: No

_Type_: String

_Maximum_: `10240`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Architectures`

The architecture tags that are associated with the repository.

_Required_: No

_Type_: Array of String

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OperatingSystems`

The operating system tags that are associated with the repository.

_Required_: No

_Type_: Array of String

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RepositoryDescription`

The short description of the repository.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UsageText`

The longform usage details of the contents of the repository. The usage text provides
context for users of the repository.

_Required_: No

_Type_: String

_Maximum_: `10240`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ECR::PublicRepository

Tag
