This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCAConnectorAD::DirectoryRegistration

Creates a directory registration that authorizes communication between AWS Private CA and an
Active Directory

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::PCAConnectorAD::DirectoryRegistration",
  "Properties" : {
      "DirectoryId" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::PCAConnectorAD::DirectoryRegistration
Properties:
  DirectoryId: String
  Tags:
    Key: Value

```

## Properties

`DirectoryId`

The identifier of the Active Directory.

_Required_: Yes

_Type_: String

_Pattern_: `^d-[0-9a-f]{10}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Metadata assigned to a directory registration consisting of a key-value pair.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DirectoryRegistrationArn`

The Amazon Resource Name (ARN) that was returned when you called [CreateDirectoryRegistration](https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateDirectoryRegistration.html) .

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VpcInformation

AWS::PCAConnectorAD::ServicePrincipalName
