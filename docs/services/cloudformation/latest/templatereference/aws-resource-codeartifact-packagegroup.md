This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeArtifact::PackageGroup

Creates a package group. For more information about creating package groups, including example CLI commands, see [Create a package group](../../../codeartifact/latest/ug/create-package-group.md) in the _CodeArtifact User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CodeArtifact::PackageGroup",
  "Properties" : {
      "ContactInfo" : String,
      "Description" : String,
      "DomainName" : String,
      "DomainOwner" : String,
      "OriginConfiguration" : OriginConfiguration,
      "Pattern" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::CodeArtifact::PackageGroup
Properties:
  ContactInfo: String
  Description: String
  DomainName: String
  DomainOwner: String
  OriginConfiguration:
    OriginConfiguration
  Pattern: String
  Tags:
    - Tag

```

## Properties

`ContactInfo`

The contact information of the package group.

_Required_: No

_Type_: String

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the package group.

_Required_: No

_Type_: String

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DomainName`

The domain that contains the package group.

_Required_: Yes

_Type_: String

_Pattern_: `^([a-z][a-z0-9\-]{0,48}[a-z0-9])$`

_Minimum_: `2`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DomainOwner`

The 12-digit account number of the AWS account that owns the domain. It does not include
dashes or spaces.

_Required_: No

_Type_: String

_Pattern_: `[0-9]{12}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OriginConfiguration`

Details about the package origin configuration of a package group.

_Required_: No

_Type_: [OriginConfiguration](aws-properties-codeartifact-packagegroup-originconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Pattern`

The pattern of the package group. The pattern determines which packages are associated with the package group.

_Required_: Yes

_Type_: String

_Minimum_: `2`

_Maximum_: `520`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-codeartifact-packagegroup-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The ARN of the package group.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

OriginConfiguration

All content copied from https://docs.aws.amazon.com/.
