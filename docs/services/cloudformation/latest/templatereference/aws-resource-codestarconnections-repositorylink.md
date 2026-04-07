This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeStarConnections::RepositoryLink

Information about the repository link resource, such as the repository link ARN, the associated connection ARN, encryption key ARN, and owner ID.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CodeStarConnections::RepositoryLink",
  "Properties" : {
      "ConnectionArn" : String,
      "EncryptionKeyArn" : String,
      "OwnerId" : String,
      "RepositoryName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CodeStarConnections::RepositoryLink
Properties:
  ConnectionArn: String
  EncryptionKeyArn: String
  OwnerId: String
  RepositoryName: String
  Tags:
    - Tag

```

## Properties

`ConnectionArn`

The Amazon Resource Name (ARN) of the connection associated with the repository link.

_Required_: Yes

_Type_: String

_Pattern_: `arn:(aws|aws-us-gov|aws-cn):.+:.+:[0-9]{12}:.+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionKeyArn`

The Amazon Resource Name (ARN) of the encryption key for the repository associated with the repository link.

_Required_: No

_Type_: String

_Pattern_: `arn:(aws|aws-us-gov|aws-cn):.+:.+:[0-9]{12}:.+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OwnerId`

The owner ID for the repository associated with the repository link, such as the owner
ID in GitHub.

_Required_: Yes

_Type_: String

_Pattern_: `[a-za-z0-9_\.-]+`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RepositoryName`

The name of the repository associated with the repository link.

_Required_: Yes

_Type_: String

_Pattern_: `[a-za-z0-9_\.-]+`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags for the repository to be associated with the repository link.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-codestarconnections-repositorylink-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the repository link. For
example:

`arn:aws:codestar-connections:region:account-id:repository-link/repository-link-id`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ProviderType`

Property description not available.

`RepositoryLinkArn`

Property description not available.

`RepositoryLinkId`

Property description not available.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
