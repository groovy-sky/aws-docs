This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ResourceGroups::Group Tag

Adds tags to a resource group with the specified Amazon resource name (ARN). Existing tags on a resource
group are not changed if they are not specified in the request parameters.

###### Important

Do not store personally identifiable information (PII) or other confidential or
sensitive information in tags. We use tags to provide you with billing and
administration services. Tags are not intended to be used for private or sensitive
data.

**Minimum permissions**

To run this command, you must have the following permissions:

- `resource-groups:Tag`

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

The tag key.

_Required_: No

_Type_: String

_Pattern_: `^(?!aws:).+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The tag value.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ResourceQuery

TagFilter
