This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECR::SigningConfiguration

The signing configuration for a registry, which specifies rules
for automatically signing images when pushed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ECR::SigningConfiguration",
  "Properties" : {
      "Rules" : [ Rule, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ECR::SigningConfiguration
Properties:
  Rules:
    - Rule

```

## Properties

`Rules`

A list of signing rules. Each rule defines a signing profile and optional repository
filters that determine which images are automatically signed.

_Required_: Yes

_Type_: Array of [Rule](aws-properties-ecr-signingconfiguration-rule.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`RegistryId`

The account ID of the destination registry.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

RepositoryFilter

All content copied from https://docs.aws.amazon.com/.
