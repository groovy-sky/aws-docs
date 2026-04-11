This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeArtifact::Repository

The `AWS::CodeArtifact::Repository` resource creates an AWS CodeArtifact repository.
CodeArtifact _repositories_ contain a set of package versions.
For more information about repositories, see the
[Repository concepts information](../../../codeartifact/latest/ug/codeartifact-concepts.md#welcome-concepts-repository)
in the _CodeArtifact User Guide_. For more information about the `CreateRepository` API, see
[CreateRepository](../../../../reference/codeartifact/latest/apireference/api-createrepository.md)
in the _CodeArtifact API Reference_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CodeArtifact::Repository",
  "Properties" : {
      "Description" : String,
      "DomainName" : String,
      "DomainOwner" : String,
      "ExternalConnections" : [ String, ... ],
      "PermissionsPolicyDocument" : Json,
      "RepositoryName" : String,
      "Tags" : [ Tag, ... ],
      "Upstreams" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CodeArtifact::Repository
Properties:
  Description: String
  DomainName: String
  DomainOwner: String
  ExternalConnections:
    - String
  PermissionsPolicyDocument: Json
  RepositoryName: String
  Tags:
    - Tag
  Upstreams:
    - String

```

## Properties

`Description`

A text description of the repository.

_Required_: No

_Type_: String

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainName`

The name of the domain that contains the repository.

_Required_: Yes

_Type_: String

_Pattern_: `^([a-z][a-z0-9\-]{0,48}[a-z0-9])$`

_Minimum_: `2`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DomainOwner`

The 12-digit account number of the AWS account that owns the domain that contains the repository. It does not include
dashes or spaces.

_Required_: No

_Type_: String

_Pattern_: `[0-9]{12}`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExternalConnections`

An array of external connections associated with the repository. For more information, see [Supported external connection repositories](../../../codeartifact/latest/ug/external-connection.md#supported-public-repositories) in the _CodeArtifact user guide_.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PermissionsPolicyDocument`

The document that defines the resource policy that is set on a repository.

_Required_: No

_Type_: Json

_Minimum_: `2`

_Maximum_: `5120`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RepositoryName`

The name of an upstream repository.

_Required_: Yes

_Type_: String

_Pattern_: `^([A-Za-z0-9][A-Za-z0-9._\-]{1,99})$`

_Minimum_: `2`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of tags to be applied to the repository.

_Required_: No

_Type_: Array of [Tag](aws-properties-codeartifact-repository-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Upstreams`

A list of upstream repositories to associate with the repository. The order of the upstream repositories
in the list determines their priority order when AWS CodeArtifact looks for a requested package version. For more
information, see [Working with upstream repositories](../../../codeartifact/latest/ug/repos-upstream.md).

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource arn.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

When you pass the logical ID of this resource, the function returns the Amazon Resource Name (ARN) of the repository.

`DomainName`

When you pass the logical ID of this resource, the function returns the domain name that contains the repository.

`DomainOwner`

When you pass the logical ID of this resource, the function returns the 12-digit account number of the AWS account that owns the domain that contains the repository.

`Name`

When you pass the logical ID of this resource, the function returns the name of the repository.

## Examples

The following examples can help you create CodeArtifact repositories using CloudFormation.

- [Create a domain and repository](#aws-resource-codeartifact-repository--examples--Create_a_domain_and_repository)

- [Create a repository with an upstream repository and external connection](#aws-resource-codeartifact-repository--examples--Create_a_repository_with_an_upstream_repository_and_external_connection)

- [Create a domain and repository with tags](#aws-resource-codeartifact-repository--examples--Create_a_domain_and_repository_with_tags)

### Create a domain and repository

The following example creates a CodeArtifact domain named _my-domain_ and a CodeArtifact
repository named _my-repo_ inside it.

#### YAML

```yaml

Resources:
  MyCodeArtifactDomain:
    Type: 'AWS::CodeArtifact::Domain'
    Properties:
      DomainName: "my-domain"
  MyCodeArtifactRepository:
    Type: 'AWS::CodeArtifact::Repository'
    Properties:
      RepositoryName: "my-repo"
      DomainName: !GetAtt MyCodeArtifactDomain.Name
```

#### JSON

```json

{
  "Resources": {
    "MyCodeArtifactDomain": {
      "Type": "AWS::CodeArtifact::Domain",
      "Properties": {
        "DomainName": "my-domain"
      }
    },
    "MyCodeArtifactRepository": {
      "Type": "AWS::CodeArtifact::Repository",
      "Properties": {
        "RepositoryName": "my-repo",
        "DomainName": {
          "Fn::GetAtt": [
            "MyCodeArtifactDomain",
            "Name"
          ]
        }
      }
    }
  }
}
```

### Create a repository with an upstream repository and external connection

The following example creates a CodeArtifact domain named _my-domain_ to store repositories. It also
creates two CodeArtifact repositories: _my-repo_ and _my-upstream-repo_
within the domain. _my-repo_ has _my-upstream-repo_ configured as an upstream
repository, and _my-upstream-repo_ has an external connection to the public repository,
**npmjs**.

#### YAML

```yaml

Resources:
  MyCodeArtifactDomain:
    Type: 'AWS::CodeArtifact::Domain'
    Properties:
      DomainName: "my-domain"
  MyCodeArtifactUpstreamRepository:
    Type: 'AWS::CodeArtifact::Repository'
    Properties:
      RepositoryName: "my-upstream-repo"
      DomainName: !GetAtt MyCodeArtifactDomain.Name
      ExternalConnections:
        - public:npmjs
  MyCodeArtifactRepository:
    Type: 'AWS::CodeArtifact::Repository'
    Properties:
      RepositoryName: "my-repo"
      DomainName: !GetAtt MyCodeArtifactDomain.Name
      Upstreams:
        - !GetAtt MyCodeArtifactUpstreamRepository.Name
```

#### JSON

```json

{
  "Resources": {
    "MyCodeArtifactDomain": {
      "Type": "AWS::CodeArtifact::Domain",
      "Properties": {
        "DomainName": "my-domain"
      }
    },
    "MyCodeArtifactUpstreamRepository": {
      "Type": "AWS::CodeArtifact::Repository",
      "Properties": {
        "RepositoryName": "my-upstream-repo",
        "DomainName": {
          "Fn::GetAtt": [
            "MyCodeArtifactDomain",
            "Name"
          ]
        },
        "ExternalConnections": [
          "public:npmjs"
        ]
      }
    },
    "MyCodeArtifactRepository": {
      "Type": "AWS::CodeArtifact::Repository",
      "Properties": {
        "RepositoryName": "my-repo",
        "DomainName": {
          "Fn::GetAtt": [
            "MyCodeArtifactDomain",
            "Name"
          ]
        },
        "Upstreams": [
          {
            "Fn::GetAtt": [
              "MyCodeArtifactUpstreamRepository",
              "Name"
            ]
          }
        ]
      }
    }
  }
}
```

### Create a domain and repository with tags

The following example creates a CodeArtifact domain named _my-domain_ and a CodeArtifact
repository named _my-repo_ inside it with two tags. One tag consists of a key
named `keyname1` and a value of `value1`. The other
consists of a key named `keyname2` and a value of `value2`.

#### YAML

```yaml

Resources:
  MyCodeArtifactDomain:
    Type: 'AWS::CodeArtifact::Domain'
    Properties:
      DomainName: "my-domain"
  MyCodeArtifactRepository:
    Type: 'AWS::CodeArtifact::Repository'
    Properties:
      RepositoryName: "my-repo"
      DomainName: !GetAtt MyCodeArtifactDomain.Name
      Tags:
        - Key: "keyname1"
          Value: "value1"
        - Key: "keyname2"
          Value: "value2"
```

#### JSON

```json

{
  "Resources": {
    "MyCodeArtifactDomain": {
      "Type": "AWS::CodeArtifact::Domain",
      "Properties": {
        "DomainName": "my-domain"
      }
    },
    "MyCodeArtifactRepository": {
      "Type": "AWS::CodeArtifact::Repository",
      "Properties": {
        "RepositoryName": "my-repo",
        "DomainName": {
          "Fn::GetAtt": [
            "MyCodeArtifactDomain",
            "Name"
          ]
        },
        "Tags" : [
          {
            "Key" : "keyname1",
            "Value" : "value1"
          },
          {
            "Key" : "keyname2",
            "Value" : "value2"
          }
        ]
      }
    }
  }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
