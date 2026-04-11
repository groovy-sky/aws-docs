This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElementalInference::Feed OutputConfig

Contains one output, of a specific type, that you want in this feed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Clipping" : ClippingConfig,
  "Cropping" : Json
}

```

### YAML

```yaml

  Clipping:
    ClippingConfig
  Cropping: Json

```

## Properties

`Clipping`

Specifies that this is a clipping output.

_Required_: No

_Type_: [ClippingConfig](aws-properties-elementalinference-feed-clippingconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Cropping`

Specifies that this is a cropping output. Enter empty brace brackets {} as the value.

You might be creating this feed in order to pass it to AWS Elemental MediaLive, because you are using MediaLive to
implement cropping. In this case, MediaLive will insert the cropping output in the feed that you pass.
Don't include the output yourself when you create the feed.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetOutput

Next

All content copied from https://docs.aws.amazon.com/.
