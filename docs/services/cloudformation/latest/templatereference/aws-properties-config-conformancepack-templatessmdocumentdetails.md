This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Config::ConformancePack TemplateSSMDocumentDetails

This API allows you to create a conformance pack template with an AWS Systems Manager document (SSM document).
To deploy a conformance pack using an SSM document, first create an SSM document with conformance pack content, and then provide the `DocumentName` in the [PutConformancePack API](https://docs.aws.amazon.com/config/latest/APIReference/API_PutConformancePack.html). You can also provide the `DocumentVersion`.

The `TemplateSSMDocumentDetails` object contains the name of the SSM document and the version of the SSM document.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DocumentName" : String,
  "DocumentVersion" : String
}

```

### YAML

```yaml

  DocumentName: String
  DocumentVersion: String

```

## Properties

`DocumentName`

The name or Amazon Resource Name (ARN) of the SSM document to use to create a conformance pack.
If you use the document name, AWS Config checks only your account and AWS Region for the SSM document.

_Required_: No

_Type_: String

_Minimum_: `3`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DocumentVersion`

The version of the SSM document to use to create a conformance pack. By default, AWS Config uses the latest version.

###### Note

This field is optional.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ConformancePackInputParameter

AWS::Config::DeliveryChannel
