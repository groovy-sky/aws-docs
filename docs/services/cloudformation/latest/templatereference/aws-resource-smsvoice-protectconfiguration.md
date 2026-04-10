This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SMSVOICE::ProtectConfiguration

Create a new protect configuration. By default all country rule sets for each capability are set to `ALLOW`.
A protect configurations name is stored as a Tag with the key set to `Name` and value as the name of the protect configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SMSVOICE::ProtectConfiguration",
  "Properties" : {
      "CountryRuleSet" : CountryRuleSet,
      "DeletionProtectionEnabled" : Boolean,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SMSVOICE::ProtectConfiguration
Properties:
  CountryRuleSet:
    CountryRuleSet
  DeletionProtectionEnabled: Boolean
  Tags:
    - Tag

```

## Properties

`CountryRuleSet`

The set of `CountryRules` you specify to control which countries AWS End User Messaging SMS can send your messages to.

_Required_: No

_Type_: [CountryRuleSet](aws-properties-smsvoice-protectconfiguration-countryruleset.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeletionProtectionEnabled`

The status of deletion protection for the protect configuration. When set to true deletion protection is enabled. By default this is set to false.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key and value pair tags that are associated with the resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-smsvoice-protectconfiguration-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns `ProtectConfigurationId`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the protect configuration.

`ProtectConfigurationId`

The unique identifier for the protect configuration.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TwoWay

CountryRule

All content copied from https://docs.aws.amazon.com/.
