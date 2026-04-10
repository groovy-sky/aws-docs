This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Comprehend::DocumentClassifier Tag

A key-value pair that adds as a metadata to a resource used by Amazon Comprehend. For
example, a tag with the key-value pair ‘Department’:’Sales’ might be added to a resource to
indicate its use by a particular department.

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

The initial part of a key-value pair that forms a tag associated with a given resource.
For instance, if you want to show which resources are used by which departments, you might use
“Department” as the key portion of the pair, with multiple possible values such as “sales,”
“legal,” and “administration.”

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The second part of a key-value pair that forms a tag associated with a given resource.
For instance, if you want to show which resources are used by which departments, you might use
“Department” as the initial (key) portion of the pair, with a value of “sales” to indicate the
sales department.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentReaderConfig

VpcConfig

All content copied from https://docs.aws.amazon.com/.
