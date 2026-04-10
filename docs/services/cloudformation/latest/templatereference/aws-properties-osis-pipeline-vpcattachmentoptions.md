This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OSIS::Pipeline VpcAttachmentOptions

Options for attaching a VPC to pipeline.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttachToVpc" : Boolean,
  "CidrBlock" : String
}

```

### YAML

```yaml

  AttachToVpc: Boolean
  CidrBlock: String

```

## Properties

`AttachToVpc`

Whether a VPC is attached to the pipeline.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CidrBlock`

The CIDR block to be reserved for OpenSearch Ingestion to create elastic network interfaces (ENIs).

_Required_: Yes

_Type_: String

_Pattern_: `^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)/(3[0-2]|[12]?[0-9])$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

VpcEndpoint

All content copied from https://docs.aws.amazon.com/.
