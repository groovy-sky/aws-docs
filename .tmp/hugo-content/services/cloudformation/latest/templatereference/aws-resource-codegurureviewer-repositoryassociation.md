This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeGuruReviewer::RepositoryAssociation

This resource configures how Amazon CodeGuru Reviewer retrieves the source code to be reviewed. You can use an
AWS CloudFormation template to create an association with the following repository types:

- AWS CodeCommit - For more information, see
[Create an \
AWS CodeCommit repository association](../../../codeguru/latest/reviewer-ug/create-codecommit-association.md) in the _Amazon CodeGuru Reviewer User Guide_.

- Bitbucket - For more information, see
[Create a \
Bitbucket repository association](../../../codeguru/latest/reviewer-ug/create-bitbucket-association.md) in the _Amazon CodeGuru Reviewer User Guide_.

- GitHub Enterprise Server - For more information, see
[Create a \
GitHub Enterprise Server repository association](../../../codeguru/latest/reviewer-ug/create-github-enterprise-association.md) in the _Amazon CodeGuru Reviewer User Guide_.

- S3Bucket - For more information, see
[Create code reviews with GitHub Actions](../../../codeguru/latest/reviewer-ug/working-with-cicd.md) in the _Amazon CodeGuru Reviewer User Guide_.

###### Note

You cannot use a CloudFormation template to create an association with a GitHub repository.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CodeGuruReviewer::RepositoryAssociation",
  "Properties" : {
      "BucketName" : String,
      "ConnectionArn" : String,
      "Name" : String,
      "Owner" : String,
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::CodeGuruReviewer::RepositoryAssociation
Properties:
  BucketName: String
  ConnectionArn: String
  Name: String
  Owner: String
  Tags:
    - Tag
  Type: String

```

## Properties

`BucketName`

The name of the bucket. This is required for your S3Bucket repository. The name must start with the prefix `codeguru-reviewer-*`.

_Required_: No

_Type_: String

_Pattern_: `^\S(.*\S)?$`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConnectionArn`

The Amazon Resource Name (ARN) of an AWS CodeStar Connections connection. Its format is
`arn:aws:codestar-connections:region-id:aws-account_id:connection/connection-id`. For more information, see
[Connection](../../../../reference/codestar-connections/latest/apireference/api-connection.md) in
the _AWS CodeStar Connections API Reference_.

`ConnectionArn` must be specified for Bitbucket and GitHub Enterprise Server repositories. It has no effect if
it is specified for an AWS CodeCommit repository.

_Required_: No

_Type_: String

_Pattern_: `arn:aws(-[\w]+)*:.+:.+:[0-9]{12}:.+`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the repository.

_Required_: Yes

_Type_: String

_Pattern_: `^\S[\w.-]*$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Owner`

The owner of the repository. For a GitHub Enterprise Server or Bitbucket repository, this is the username
for the account that owns the repository.

`Owner` must be specified for Bitbucket and GitHub Enterprise Server repositories. It has no effect if
it is specified for an AWS CodeCommit repository.

_Required_: No

_Type_: String

_Pattern_: `^\S(.*\S)?$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs used to tag an associated repository. A tag is a custom attribute label with two parts:

- A _tag key_ (for example, `CostCenter`,
`Environment`, `Project`, or `Secret`). Tag
keys are case sensitive.

- An optional field known as a _tag value_ (for example,
`111122223333`, `Production`, or a team name).
Omitting the tag value is the same as using an empty string. Like tag keys, tag
values are case sensitive.

_Required_: No

_Type_: Array of [Tag](aws-properties-codegurureviewer-repositoryassociation-tag.md)

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The type of repository that contains the source code to be reviewed. The valid values are:

- `CodeCommit`

- `Bitbucket`

- `GitHubEnterpriseServer`

- `S3Bucket`

_Required_: Yes

_Type_: String

_Allowed values_: `CodeCommit | Bitbucket | GitHubEnterpriseServer | S3Bucket`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function,
`Ref` returns the Amazon Resource Name (ARN) of the AWS CodeGuru Reviewer
[`RepositoryAssociation`](../../../codeguru/latest/reviewer-api/api-repositoryassociation.md),
such as `arn:aws:codeguru-reviewer:region:123456789012:association/universally-unique-identifier`.

For more information about using the `Ref` function, see
[Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. The following is the available sample return values.
For more information about using `Fn::GetAtt`, see
[Fn::GetAtt](../userguide/intrinsic-function-reference-getatt.md).

`AssociationArn`

The Amazon Resource Name (ARN) of the
[`RepositoryAssociation`](../../../codeguru/latest/reviewer-api/api-repositoryassociation.md) object.
You can retrieve this ARN by calling `ListRepositories`.

## Examples

- [Create an AWS CodeCommit repository association using an existing CodeCommit repository](#aws-resource-codegurureviewer-repositoryassociation--examples--Create_an_CodeCommit_repository_association_using_an_existing_CodeCommit_repository)

- [Create an AWS CodeCommit repository association with a new CodeCommit repository](#aws-resource-codegurureviewer-repositoryassociation--examples--Create_an_CodeCommit_repository_association_with_a_new_CodeCommit_repository)

- [Create a Bitbucket repository association](#aws-resource-codegurureviewer-repositoryassociation--examples--Create_a_Bitbucket_repository_association)

- [Create a GitHub Enterprise Server repository association](#aws-resource-codegurureviewer-repositoryassociation--examples--Create_a_GitHub_Enterprise_Server_repository_association)

- [Create a repository association with tags](#aws-resource-codegurureviewer-repositoryassociation--examples--Create_a_repository_association_with_tags)

### Create an AWS CodeCommit repository association using an existing CodeCommit repository

The following example creates an Amazon CodeGuru Reviewer repository association using an existing
AWS CodeCommit repository named `MyRepository`.

#### YAML

```yaml

Resources:
  MyRepositoryAssociation:
    Type: AWS::CodeGuruReviewer::RepositoryAssociation
    Properties:
      Name: MyRepository
      Type: CodeCommit
```

#### JSON

```json

{
  "Resources": {
    "MyRepositoryAssociation": {
      "Type": "AWS::CodeGuruReviewer::RepositoryAssociation",
      "Properties": {
        "Name": "MyRepository",
        "Type": "CodeCommit"
      }
    }
  }
}
```

### Create an AWS CodeCommit repository association with a new CodeCommit repository

The following example creates an AWS CodeCommit repository named `MyRepositoryName`. Next, it creates an Amazon
CodeGuru Reviewer repository association using the CodeCommit repository. The `DependsOn` property in the
`RepositoryAssociation` specifies the CodeCommit repository. This is because the repository association
can be created only after the CodeCommit repository is created.

#### YAML

```yaml

Resources:
  MyRepository:
    Type: AWS::CodeCommit::Repository
    Properties:
      RepositoryName: MyRepositoryName
  MyRepositoryAssociation:
    Type: AWS::CodeGuruReviewer::RepositoryAssociation
    Properties:
      Name: MyRepositoryName
      Type: CodeCommit
```

#### JSON

```json

{
  "Resources": {
    "MyRepository": {
      "Type": "AWS::CodeCommit::Repository",
      "Properties": {
        "RepositoryName": "MyRepositoryName"
      }
    },
    "MyRepositoryAssociation": {
      "Type": "AWS::CodeGuruReviewer::RepositoryAssociation",
      "Properties": {
        "Name": "MyRepositoryName",
        "Type": "CodeCommit"
      }
    }
  }
}
```

### Create a Bitbucket repository association

The following example creates an Amazon CodeGuru Reviewer repository association using a
Bitbucket repository.

#### YAML

```yaml

Resources:
  MyRepositoryAssociation:
    Type: AWS::CodeGuruReviewer::RepositoryAssociation
    Properties:
      Name: MyBitbucketRepoName
      Type: Bitbucket
      ConnectionArn: arn:aws:codestar-connections:us-west-2:123456789012:connection/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee
      Owner: MyOwnerName
```

#### JSON

```json

{
  "Resources": {
    "MyRepositoryAssociation": {
      "Type": "AWS::CodeGuruReviewer::RepositoryAssociation",
      "Properties": {
        "Name": "MyBitbucketRepoName",
        "Type": "Bitbucket",
        "ConnectionArn": "arn:aws:codestar-connections:us-west-2:123456789012:connection/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
        "Owner": "MyOwnerName"
      }
    }
  }
}
```

### Create a GitHub Enterprise Server repository association

The following example creates an Amazon CodeGuru Reviewer repository association using a
GitHub Enterprise Server repository.

#### YAML

```yaml

Resources:
  MyRepositoryAssociation:
    Type: AWS::CodeGuruReviewer::RepositoryAssociation
    Properties:
      Name: MyGitHubEnterpriseRepoName
      Type: GitHubEnterpriseServer
      ConnectionArn: arn:aws:codestar-connections:us-west-2:123456789012:connection/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee
      Owner: MyOwnerName
```

#### JSON

```json

{
  "Resources": {
    "MyRepositoryAssociation": {
      "Type": "AWS::CodeGuruReviewer::RepositoryAssociation",
      "Properties": {
        "Name": "MyGitHubEnterpriseRepoName",
        "Type": "GitHubEnterpriseServer",
        "ConnectionArn": "arn:aws:codestar-connections:us-west-2:123456789012:connection/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
        "Owner": "MyOwnerName"
      }
    }
  }
}
```

### Create a repository association with tags

The following example creates an Amazon CodeGuru Reviewer repository association that has two tags. The repository is an AWS CodeCommit repository.
Use the `Tags` property the same way to add tags when you create a BitBucket or GitHub Enterprise repository association.

#### YAML

```yaml

Resources:
  MyRepositoryAssociation:
    Type: AWS::CodeGuruReviewer::RepositoryAssociation
    Properties:
      Name: MyRepository
      Type: CodeCommit
      Tags:
        - Key: tag1-key
          Value: tag1-value
        - Key: tag2-key
          Value: tag2-value
```

#### JSON

```json

{
  "Resources": {
    "MyRepositoryAssociation": {
      "Type": "AWS::CodeGuruReviewer::RepositoryAssociation",
      "Properties": {
        "Name": "MyRepository",
        "Type": "CodeCommit",
        "Tags": [
          {
            "Key": "tag1-key",
            "Value": "tag1-value"
          },
          {
            "Key": "tag2-key",
            "Value": "tag2-value"
          }
        ]
      }
    }
  }
}

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon CodeGuru Reviewer

Tag

All content copied from https://docs.aws.amazon.com/.
