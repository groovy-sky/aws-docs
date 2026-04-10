This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECR::PublicRepository

The `AWS::ECR::PublicRepository` resource specifies an Amazon Elastic
Container Registry Public (Amazon ECR Public) repository, where users can push and pull
Docker images, Open Container Initiative (OCI) images, and OCI compatible artifacts. For
more information, see [Amazon ECR public\
repositories](../../../amazonecr/latest/public/public-repositories.md) in the _Amazon ECR Public User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ECR::PublicRepository",
  "Properties" : {
      "RepositoryCatalogData" : RepositoryCatalogData,
      "RepositoryName" : String,
      "RepositoryPolicyText" : Json,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ECR::PublicRepository
Properties:
  RepositoryCatalogData:
    RepositoryCatalogData
  RepositoryName: String
  RepositoryPolicyText: Json
  Tags:
    - Tag

```

## Properties

`RepositoryCatalogData`

The details about the repository that are publicly visible in the Amazon ECR Public
Gallery. For more information, see [Amazon ECR Public\
repository catalog data](../../../amazonecr/latest/public/public-repository-catalog-data.md) in the _Amazon ECR Public User_
_Guide_.

_Required_: No

_Type_: [RepositoryCatalogData](aws-properties-ecr-publicrepository-repositorycatalogdata.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RepositoryName`

The name to use for the public repository. The repository name may be specified on its
own (such as `nginx-web-app`) or it can be prepended with a namespace to
group the repository into a category (such as `project-a/nginx-web-app`). If
you don't specify a name, AWS CloudFormation generates a unique physical ID and
uses that ID for the repository name. For more information, see [Name\
Type](../userguide/aws-properties-name.md).

###### Note

If you specify a name, you cannot perform updates that require replacement of this
resource. You can perform updates that require no or some interruption. If you must
replace the resource, specify a new name.

_Required_: No

_Type_: String

_Pattern_: `^(?=.{2,256}$)((?:[a-z0-9]+(?:[._-][a-z0-9]+)*/)*[a-z0-9]+(?:[._-][a-z0-9]+)*)$`

_Minimum_: `2`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RepositoryPolicyText`

The JSON repository policy text to apply to the public repository. For more
information, see [Amazon ECR Public\
repository policies](../../../amazonecr/latest/public/public-repository-policies.md) in the _Amazon ECR Public User_
_Guide_.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-ecr-publicrepository-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name, such as
`test-repository`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Returns the Amazon Resource Name (ARN) for the specified
`AWS::ECR::PublicRepository` resource. For example,
`arn:aws:ecr-public::123456789012:repository/test-repository`.

## Examples

### Specify a public repository

The following example specifies a public repository named
`test-repository`. The repository catalog data appears publicly
on the Amazon ECR Public Gallery.

#### JSON

```json

"MyPublicRepository": {
  "Type": "AWS::ECR::PublicRepository",
  "Properties": {
    "RepositoryName" : "test-repository",
    "RepositoryCatalogData" : {
      "UsageText": "This is a sample usage text.",
      "AboutText": "This is a sample about text.",
      "OperatingSystems": [
           "Linux",
           "Windows"
       ],
       "Architectures": [
           "x86",
           "ARM"
       ],
      "RepositoryDescription": "This is a sample repository description."
    }
  }
}
```

#### YAML

```yaml

Resources:
  MyPublicRepositry:
    Type: AWS::ECR::PublicRepository
    Properties:
      RepositoryName: "test-repository"
      RepositoryCatalogData:
        UsageText: "This is a sample usage text."
        AboutText: "This is a sample about text."
        OperatingSystems:
          - "Linux"
          - "Windows"
        Architectures:
          - "x86"
          - "ARM"
        RepositoryDescription: "This is a sample repository description."
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon ECR

RepositoryCatalogData

All content copied from https://docs.aws.amazon.com/.
