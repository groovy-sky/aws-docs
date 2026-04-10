This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeArtifact::Domain

The `AWS::CodeArtifact::Domain` resource creates an AWS CodeArtifact domain.
CodeArtifact _domains_ make it easier to manage multiple repositories across an
organization. You can use a domain to apply permissions across many repositories owned by different
AWS accounts. For more information about domains, see the
[Domain concepts information](../../../codeartifact/latest/ug/codeartifact-concepts.md#welcome-concepts-domain)
in the _CodeArtifact User Guide_. For more information about the `CreateDomain` API, see
[CreateDomain](../../../../reference/codeartifact/latest/apireference/api-createdomain.md)
in the _CodeArtifact API Reference_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CodeArtifact::Domain",
  "Properties" : {
      "DomainName" : String,
      "EncryptionKey" : String,
      "PermissionsPolicyDocument" : Json,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CodeArtifact::Domain
Properties:
  DomainName: String
  EncryptionKey: String
  PermissionsPolicyDocument: Json
  Tags:
    - Tag

```

## Properties

`DomainName`

A string that specifies the name of the requested domain.

_Required_: Yes

_Type_: String

_Pattern_: `^([a-z][a-z0-9\-]{0,48}[a-z0-9])$`

_Minimum_: `2`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EncryptionKey`

The key used to encrypt the domain.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Minimum_: `1`

_Maximum_: `1011`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PermissionsPolicyDocument`

The document that defines the resource policy that is set on a domain.

_Required_: No

_Type_: Json

_Minimum_: `2`

_Maximum_: `5120`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of tags to be applied to the domain.

_Required_: No

_Type_: Array of [Tag](aws-properties-codeartifact-domain-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource arn.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

When you pass the logical ID of this resource, the function returns the Amazon Resource Name (ARN) of the domain.

`EncryptionKey`

When you pass the logical ID of this resource, the function returns the key used to encrypt the domain.

`Name`

When you pass the logical ID of this resource, the function returns the name of the domain.

`Owner`

When you pass the logical ID of this resource, the function returns the 12-digit account number of the AWS account that owns the domain.

## Examples

The following examples can help you create CodeArtifact domains using CloudFormation.

- [Create a domain](#aws-resource-codeartifact-domain--examples--Create_a_domain)

- [Create a domain with an AWS Key Management Service encryption key and IAM resource-based policy](#aws-resource-codeartifact-domain--examples--Create_a_domain_with_an_encryption_key_and_IAM_resource-based_policy)

- [Create a domain with tags](#aws-resource-codeartifact-domain--examples--Create_a_domain_with_tags)

### Create a domain

The following example creates a CodeArtifact domain named _my-domain_.

#### YAML

```yaml

Resources:
  MyCodeArtifactDomain:
    Type: 'AWS::CodeArtifact::Domain'
    Properties:
      DomainName: "my-domain"
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
    }
  }
}
```

### Create a domain with an AWS Key Management Service encryption key and IAM resource-based policy

The following example creates a CodeArtifact domain named _my-domain_
with an AWS Key Management Service encryption key and attaches an IAM resource-based policy.

#### YAML

```yaml

Resources:
  MyCodeArtifactDomain:
    Type: 'AWS::CodeArtifact::Domain'
    Properties:
      DomainName: "my-domain"
      EncryptionKey: arn:aws:kms:us-west-2:123456789012:key/12345678-9abc-def1-2345-6789abcdef12
      PermissionsPolicyDocument:
          Version: 2012-10-17
          Statement:
            - Action:
                - codeartifact:ReadFromRepository
                - codeartifact:DescribePackageVersion
                - codeartifact:DescribeRepository
                - codeartifact:GetPackageVersionReadme
                - codeartifact:GetRepositoryEndpoint
                - codeartifact:ListPackageVersionAssets
                - codeartifact:ListPackageVersionDependencies
                - codeartifact:ListPackageVersions
                - codeartifact:ListPackages
                - codeartifact:ReadFromRepository
              Effect: Allow
              Principal:
                AWS: "arn:aws:iam::123456789012:root"
              Resource: "*"
```

#### JSON

```json

{
  "Resources": {
    "MyCodeArtifactDomain": {
      "Type": "AWS::CodeArtifact::Domain",
      "Properties": {
        "DomainName": "my-domain",
        "EncryptionKey": "arn:aws:kms:us-west-2:123456789012:key/12345678-9abc-def1-2345-6789abcdef12",
        "PermissionsPolicyDocument": {
          "Version": "2012-10-17",
          "Statement": [
            {
              "Action": [
                "codeartifact:ReadFromRepository",
                "codeartifact:DescribePackageVersion",
                "codeartifact:DescribeRepository",
                "codeartifact:GetPackageVersionReadme",
                "codeartifact:GetRepositoryEndpoint",
                "codeartifact:ListPackageVersionAssets",
                "codeartifact:ListPackageVersionDependencies",
                "codeartifact:ListPackageVersions",
                "codeartifact:ListPackages",
                "codeartifact:ReadFromRepository"
              ],
              "Effect": "Allow",
              "Principal": {
                "AWS": "arn:aws:iam::123456789012:root"
              },
              "Resource": "*"
            }
          ]
        }
      }
    }
  }
}
```

### Create a domain with tags

The following example creates a CodeArtifact domain named _my-domain_ with two tags.
One tag consists of a key named `keyname1` and a value of `value1`. The other
consists of a key named `keyname2` and a value of `value2`.

#### YAML

```yaml

Resources:
  MyCodeArtifactDomain:
    Type: 'AWS::CodeArtifact::Domain'
    Properties:
      DomainName: "my-domain"
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
        "DomainName": "my-domain",
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

AWS CodeArtifact

Tag

All content copied from https://docs.aws.amazon.com/.
