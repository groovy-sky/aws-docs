This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel DvbSdtSettings

A DVB Service Description Table (SDT).

The parent of this entity is M2tsSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OutputSdt" : String,
  "RepInterval" : Integer,
  "ServiceName" : String,
  "ServiceProviderName" : String
}

```

### YAML

```yaml

  OutputSdt: String
  RepInterval: Integer
  ServiceName: String
  ServiceProviderName: String

```

## Properties

`OutputSdt`

Selects a method of inserting SDT information into an output stream. The sdtFollow setting copies SDT
information from input stream to output stream. The sdtFollowIfPresent setting copies SDT information from input
stream to output stream if SDT information is present in the input. Otherwise, it falls back on the user-defined
values. The sdtManual setting means that the user will enter the SDT information. The sdtNone setting means that the
output stream will not contain SDT information.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RepInterval`

The number of milliseconds between instances of this table in the output transport stream.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceName`

The service name placed in the serviceDescriptor in the Service Description Table (SDT). The maximum length is
256 characters.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceProviderName`

The service provider name placed in the serviceDescriptor in the Service Description Table (SDT). The maximum
length is 256 characters.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DvbNitSettings

DvbSubDestinationSettings

All content copied from https://docs.aws.amazon.com/.
