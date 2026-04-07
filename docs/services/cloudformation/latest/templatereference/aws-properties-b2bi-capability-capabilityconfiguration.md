This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Capability CapabilityConfiguration

A capability object. Currently, only EDI (electronic data interchange) capabilities are supported.
A trading capability contains the information required to transform incoming EDI documents into JSON or XML outputs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Edi" : EdiConfiguration
}

```

### YAML

```yaml

  Edi:
    EdiConfiguration

```

## Properties

`Edi`

An EDI (electronic data interchange) configuration object.

_Required_: Yes

_Type_: [EdiConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-b2bi-capability-ediconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::B2BI::Capability

EdiConfiguration
