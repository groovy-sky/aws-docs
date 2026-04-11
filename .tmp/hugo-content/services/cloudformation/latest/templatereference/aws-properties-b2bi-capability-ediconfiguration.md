This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Capability EdiConfiguration

Specifies the details for the EDI (electronic data interchange) transformation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CapabilityDirection" : String,
  "InputLocation" : S3Location,
  "OutputLocation" : S3Location,
  "TransformerId" : String,
  "Type" : EdiType
}

```

### YAML

```yaml

  CapabilityDirection: String
  InputLocation:
    S3Location
  OutputLocation:
    S3Location
  TransformerId: String
  Type:
    EdiType

```

## Properties

`CapabilityDirection`

Specifies whether this is capability is for inbound or outbound transformations.

_Required_: No

_Type_: String

_Allowed values_: `INBOUND | OUTBOUND`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputLocation`

Contains the Amazon S3 bucket and prefix for the location of the input file, which is
contained in an `S3Location` object.

_Required_: Yes

_Type_: [S3Location](aws-properties-b2bi-capability-s3location.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputLocation`

Contains the Amazon S3 bucket and prefix for the location of the output file, which is
contained in an `S3Location` object.

_Required_: Yes

_Type_: [S3Location](aws-properties-b2bi-capability-s3location.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransformerId`

Returns the system-assigned unique identifier for the transformer.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Returns the type of the capability. Currently, only `edi` is supported.

_Required_: Yes

_Type_: [EdiType](aws-properties-b2bi-capability-editype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CapabilityConfiguration

EdiType

All content copied from https://docs.aws.amazon.com/.
