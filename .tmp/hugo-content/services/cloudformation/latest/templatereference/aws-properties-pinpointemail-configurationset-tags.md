This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PinpointEmail::ConfigurationSet Tags

An object that defines the tags (keys and values) that you want to associate with the
configuration set.

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

One part of a key-value pair that defines a tag. The maximum length of a tag key is 128
characters. The minimum length is 1 character.

If you specify tags for the configuration set, then this value is required.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The optional part of a key-value pair that defines a tag. The maximum length of a tag
value is 256 characters. The minimum length is 0 characters. If you don’t want a resource to
have a specific tag value, don’t specify a value for this parameter. Amazon Pinpoint will set
the value to an empty string.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SendingOptions

TrackingOptions

All content copied from https://docs.aws.amazon.com/.
