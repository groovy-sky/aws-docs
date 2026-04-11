This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeStar::GitHubRepository

The `AWS::CodeStar::GitHubRepository` resource creates a GitHub repository
where users can store source code for use with AWS workflows. You must
provide a location for the source code ZIP file in the CloudFormation template, so
the code can be uploaded to the created repository. You must have created a personal access
token in GitHub to provide in the CloudFormation template. AWS uses
this token to connect to GitHub on your behalf.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CodeStar::GitHubRepository",
  "Properties" : {
      "Code" : Code,
      "ConnectionArn" : String,
      "EnableIssues" : Boolean,
      "IsPrivate" : Boolean,
      "RepositoryAccessToken" : String,
      "RepositoryDescription" : String,
      "RepositoryName" : String,
      "RepositoryOwner" : String
    }
}

```

### YAML

```yaml

Type: AWS::CodeStar::GitHubRepository
Properties:
  Code:
    Code
  ConnectionArn: String
  EnableIssues: Boolean
  IsPrivate: Boolean
  RepositoryAccessToken: String
  RepositoryDescription: String
  RepositoryName: String
  RepositoryOwner: String

```

## Properties

`Code`

Information about code to be committed to a repository after it is created in an
CloudFormation stack.

_Required_: No

_Type_: [Code](aws-properties-codestar-githubrepository-code.md)

_Update requires_: Updates are not supported.

`ConnectionArn`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableIssues`

Indicates whether to enable issues for the GitHub repository. You can use GitHub
issues to track information and bugs for your repository.

_Required_: No

_Type_: Boolean

_Update requires_: Updates are not supported.

`IsPrivate`

Indicates whether the GitHub repository is a private repository. If so, you choose
who can see and commit to this repository.

_Required_: No

_Type_: Boolean

_Update requires_: Updates are not supported.

`RepositoryAccessToken`

The GitHub user's personal access token for the GitHub repository.

_Required_: No

_Type_: String

_Update requires_: Updates are not supported.

`RepositoryDescription`

A comment or description about the new repository. This description is displayed in
GitHub after the repository is created.

_Required_: No

_Type_: String

_Update requires_: Updates are not supported.

`RepositoryName`

The name of the repository you want to create in GitHub with CloudFormation
stack creation.

_Required_: Yes

_Type_: String

_Update requires_: Updates are not supported.

`RepositoryOwner`

The GitHub user name for the owner of the GitHub repository to be created. If this
repository should be owned by a GitHub organization, provide its name.

_Required_: Yes

_Type_: String

_Update requires_: Updates are not supported.

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns a string combination of the repository owner and the
repository name, such as `my-github-account/my-github-repo`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The repository ID.

## Examples

### GitHub Repository Resource Configuration

The following example for the `AWS::CodeStar::GitHubRepository`
resource creates a private GitHub repository with issues enabled.

When passing secret parameters, do not enter the value directly into the template.
The value is rendered as plain text and is readable. Instead, enter secret parameters
using one of the following methods:

- Pass the value in as a `NoEcho` parameter. For more
information, see [Referencing a Parameter in a Template](../userguide/parameters-section-structure.md#parameters-section-structure-referencing).

- Store the GitHub token in AWS Secrets Manager and [retrieve it](../userguide/dynamic-references.md#dynamic-references-secretsmanager) through the resource property. The following example
shows the token ID as the parameter stored in AWS Secrets Manager with
this value:
`resolve:secretsmanager:your-secret-manager-name:SecretString:your-secret-manager-key`.

#### JSON

```json

{
    "MyRepo": {
        "Type": "AWS::CodeStar::GitHubRepository",
        "Properties": {
            "Code": {
                "S3": {
                    "Bucket": "amzn-s3-demo-bucket",
                    "Key": "sourcecode.zip",
                    "ObjectVersion": "1"
                }
            },
            "EnableIssues": true,
            "IsPrivate": true,
            "RepositoryAccessToken": "{{resolve:secretsmanager:your-secret-manager-name:SecretString:your-secret-manager-key}}",
            "RepositoryDescription": "a description",
            "RepositoryName": "my-github-repo",
            "RepositoryOwner": "my-github-account"
        }
    }
}
```

#### YAML

```yaml

MyRepo:
  Type: AWS::CodeStar::GitHubRepository
  Properties:
    Code:
      S3:
        Bucket: "amzn-s3-demo-bucket"
        Key: "sourcecode.zip"
        ObjectVersion: "1"
    EnableIssues: true
    IsPrivate: true
    RepositoryAccessToken: '{{resolve:secretsmanager:your-secret-manager-name:SecretString:your-secret-manager-key}}'
    RepositoryDescription: a description
    RepositoryName: my-github-repo
    RepositoryOwner: my-github-account
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS CodeStar

Code

All content copied from https://docs.aws.amazon.com/.
