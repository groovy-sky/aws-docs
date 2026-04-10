This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeCommit::Repository

Creates a new, empty repository.

###### Important

AWS CodeCommit is no longer available to new customers. Existing customers of
AWS CodeCommit can continue to use the service as normal.
[Learn more"](https://aws.amazon.com/blogs/devops/how-to-migrate-your-aws-codecommit-repository-to-another-git-provider)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CodeCommit::Repository",
  "Properties" : {
      "Code" : Code,
      "KmsKeyId" : String,
      "RepositoryDescription" : String,
      "RepositoryName" : String,
      "Tags" : [ Tag, ... ],
      "Triggers" : [ RepositoryTrigger, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CodeCommit::Repository
Properties:
  Code:
    Code
  KmsKeyId: String
  RepositoryDescription: String
  RepositoryName: String
  Tags:
    - Tag
  Triggers:
    - RepositoryTrigger

```

## Properties

`Code`

Information about code to be committed to a repository after it is created in an
AWS CloudFormation stack. Information about code is only used in resource
creation. Updates to a stack will not reflect changes made to code properties after
initial resource creation.

###### Note

You can only use this property to add code when creating a repository with a
CloudFormation template at creation time. This property cannot be used for
updating code to an existing repository.

_Required_: No

_Type_: [Code](aws-properties-codecommit-repository-code.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

The ID of the AWS Key Management Service encryption key used to encrypt and decrypt the repository.

###### Note

The input can be the full ARN, the key ID, or the key alias. For more information, see [Finding the key ID and key ARN](../../../kms/latest/developerguide/find-cmk-id-arn.md).

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9:/_-]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RepositoryDescription`

A comment or description about the new repository.

###### Note

The description field for a repository accepts all HTML characters and all valid
Unicode characters. Applications that do not HTML-encode the description and display
it in a webpage can expose users to potentially malicious code. Make sure that you
HTML-encode the description field in any application that uses this API to display
the repository description on a webpage.

_Required_: No

_Type_: String

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RepositoryName`

The name of the new repository to be created.

###### Note

The repository name must be unique across the calling AWS account. Repository names
are limited to 100 alphanumeric, dash, and underscore characters, and cannot include
certain characters. For more information about the limits on repository names, see
[Quotas](../../../codecommit/latest/userguide/limits.md) in the _AWS CodeCommit User Guide_. The
suffix .git is prohibited.

_Required_: Yes

_Type_: String

_Pattern_: `[\w\.-]+`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

One or more tag key-value pairs to use when tagging this repository.

_Required_: No

_Type_: Array of [Tag](aws-properties-codecommit-repository-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Triggers`

The JSON block of configuration information for each trigger.

_Required_: No

_Type_: Array of [RepositoryTrigger](aws-properties-codecommit-repository-repositorytrigger.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When the logical ID of this resource is provided to the Ref intrinsic function,
`Ref` returns the repository ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

When you pass the logical ID of this resource, the
function returns the Amazon Resource Name (ARN) of the repository.

`CloneUrlHttp`

When you pass the logical ID of this resource,
the function returns the URL to use for cloning the repository over HTTPS.

`CloneUrlSsh`

When you pass the logical ID of this resource, the
function returns the URL to use for cloning the repository over SSH.

`Name`

When you pass the logical ID of this resource, the
function returns the repository's name.

## Examples

The following examples can help you create CodeCommit repositories using
CloudFormation.

- [Example](#aws-resource-codecommit-repository--examples--Example)

- [Example](#aws-resource-codecommit-repository--examples--Example)

### Example

The following example creates a CodeCommit repository named
_MyDemoRepo_. The newly created repository is populated
with code stored in an Amazon S3 bucket named
_MySourceCodeBucket_ and placed in a branch named
_development_, which is the default branch for the
repository.

#### JSON

```json

{
    "MyRepo": {
        "Type": "AWS::CodeCommit::Repository",
        "Properties": {
            "RepositoryName": "MyDemoRepo",
            "RepositoryDescription": "This is a repository for my project with code from MySourceCodeBucket.",
            "Code": {
                "BranchName": "development",
                "S3": {
                    "Bucket": "MySourceCodeBucket",
                    "Key": "MyKey",
                    "ObjectVersion": "1"
                }
            }
        }
    }
}
```

#### YAML

```yaml

MyRepo:
  Type: AWS::CodeCommit::Repository
  Properties:
    RepositoryName: MyDemoRepo
    RepositoryDescription: This is a repository for my project with code from MySourceCodeBucket.
    Code:
      BranchName: development
      S3:
        Bucket: MySourceCodeBucket
        Key: MyKey
        ObjectVersion: 1

```

### Example

The following example creates a CodeCommit repository with a trigger
for all events in the _development_ branch.

#### JSON

```json

{
    "MyRepo": {
        "Type": "AWS: : CodeCommit: : Repository",
        "Properties": {
            "RepositoryName": "MyRepoName",
            "RepositoryDescription": "a description",
            "Triggers": [
                {
                    "Name": "MyTrigger",
                    "CustomData": "Project ID 12345",
                    "DestinationArn": {
                        "Ref": "SNSarn"
                    },
                    "Branches": [
                        "development"
                    ],
                    "Events": [
                        "all"
                    ]
                }
            ]
        }
    }
}
```

#### YAML

```yaml

MyRepo:
  Type: AWS::CodeCommit::Repository
  Properties:
    RepositoryName: MyRepoName
    RepositoryDescription: a description
    Triggers:
    - Name: MyTrigger
      CustomData: Project ID 12345
      DestinationArn:
        Ref: SNSarn
      Branches:
      - development
      Events:
      - all
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS CodeCommit

Code

All content copied from https://docs.aws.amazon.com/.
