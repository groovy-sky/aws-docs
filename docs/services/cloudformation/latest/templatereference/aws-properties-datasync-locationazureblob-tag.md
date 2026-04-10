This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationAzureBlob Tag

Specifies labels that help you categorize, filter, and search for your AWS
resources. We recommend creating at least a name tag for your transfer location.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

The key for an AWS resource tag.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9\s+=._:/-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value for an AWS resource tag.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9\s+=._:@/-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ManagedSecretConfig

AWS::DataSync::LocationEFS

All content copied from https://docs.aws.amazon.com/.
