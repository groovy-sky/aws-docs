This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::B2BI::Transformer

Creates a transformer. AWS B2B Data Interchange currently supports two scenarios:

- _Inbound EDI_: the AWS customer receives an EDI
file from their trading partner. AWS B2B Data Interchange converts this EDI file into a JSON or
XML file with a service-defined structure. A mapping template provided by the
customer, in JSONata or XSLT format, is optionally applied to this file to produce a
JSON or XML file with the structure the customer requires.

- _Outbound EDI_: the AWS customer has a JSON or
XML file containing data that they wish to use in an EDI file. A mapping template,
provided by the customer (in either JSONata or XSLT format) is applied to this file
to generate a JSON or XML file in the service-defined structure. This file is then
converted to an EDI file.

###### Note

The following fields are provided for backwards compatibility only:
`fileFormat`, `mappingTemplate`, `ediType`, and
`sampleDocument`.

- Use the `mapping` data type in place of `mappingTemplate`
and `fileFormat`

- Use the `sampleDocuments` data type in place of
`sampleDocument`

- Use either the `inputConversion` or `outputConversion` in
place of `ediType`

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::B2BI::Transformer",
  "Properties" : {
      "InputConversion" : InputConversion,
      "Mapping" : Mapping,
      "Name" : String,
      "OutputConversion" : OutputConversion,
      "SampleDocuments" : SampleDocuments,
      "Status" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::B2BI::Transformer
Properties:
  InputConversion:
    InputConversion
  Mapping:
    Mapping
  Name: String
  OutputConversion:
    OutputConversion
  SampleDocuments:
    SampleDocuments
  Status: String
  Tags:
    - Tag

```

## Properties

`InputConversion`

Returns a structure that contains the format options for the transformation.

_Required_: No

_Type_: [InputConversion](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-b2bi-transformer-inputconversion.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Mapping`

Returns the structure that contains the mapping template and its language (either XSLT or JSONATA).

_Required_: No

_Type_: [Mapping](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-b2bi-transformer-mapping.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

Returns the descriptive name for the transformer.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]{1,512}$`

_Minimum_: `1`

_Maximum_: `254`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutputConversion`

Returns the `OutputConversion` object, which contains the format options for the outbound transformation.

_Required_: No

_Type_: [OutputConversion](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-b2bi-transformer-outputconversion.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SampleDocuments`

Returns a structure that contains the Amazon S3 bucket and an array of the corresponding keys used to identify the location for your sample documents.

_Required_: No

_Type_: [SampleDocuments](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-b2bi-transformer-sampledocuments.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

Returns the state of the newly created transformer. The transformer can be either
`active` or `inactive`. For the transformer to be used in a
capability, its status must `active`.

_Required_: Yes

_Type_: String

_Allowed values_: `active | inactive`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A key-value pair for a specific transformer. Tags are metadata that you can use to
search for and group capabilities for various purposes.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-b2bi-transformer-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`CreatedAt`

Returns a timestamp indicating when the transformer was created. For example,
`2023-07-20T19:58:44.624Z`.

`ModifiedAt`

Returns a timestamp representing the date and time for the most recent change for the
transformer object.

`TransformerArn`

Returns an Amazon Resource Name (ARN) for a specific transformer.

`TransformerId`

The system-assigned unique identifier for the transformer.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AdvancedOptions
