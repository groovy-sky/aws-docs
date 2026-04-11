This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ConnectionFunction ConnectionFunctionConfig

Contains configuration information about a CloudFront function.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Comment" : String,
  "KeyValueStoreAssociations" : [ KeyValueStoreAssociation, ... ],
  "Runtime" : String
}

```

### YAML

```yaml

  Comment: String
  KeyValueStoreAssociations:
    - KeyValueStoreAssociation
  Runtime: String

```

## Properties

`Comment`

A comment to describe the function.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyValueStoreAssociations`

The configuration for the key value store associations.

_Required_: No

_Type_: Array of [KeyValueStoreAssociation](aws-properties-cloudfront-connectionfunction-keyvaluestoreassociation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Runtime`

The function's runtime environment version.

_Required_: Yes

_Type_: String

_Allowed values_: `cloudfront-js-2.0`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CloudFront::ConnectionFunction

KeyValueStoreAssociation

All content copied from https://docs.aws.amazon.com/.
