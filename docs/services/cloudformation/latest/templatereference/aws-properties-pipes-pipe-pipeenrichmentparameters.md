This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe PipeEnrichmentParameters

The parameters required to set up enrichment on your pipe.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HttpParameters" : PipeEnrichmentHttpParameters,
  "InputTemplate" : String
}

```

### YAML

```yaml

  HttpParameters:
    PipeEnrichmentHttpParameters
  InputTemplate: String

```

## Properties

`HttpParameters`

Contains the HTTP parameters to use when the target is a API Gateway REST
endpoint or EventBridge ApiDestination.

If you specify an API Gateway REST API or EventBridge ApiDestination as a
target, you can use this parameter to specify headers, path parameters, and query string
keys/values as part of your target invoking request. If you're using ApiDestinations, the
corresponding Connection can also have these values configured. In case of any conflicting
keys, values from the Connection take precedence.

_Required_: No

_Type_: [PipeEnrichmentHttpParameters](aws-properties-pipes-pipe-pipeenrichmenthttpparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputTemplate`

Valid JSON text passed to the enrichment. In this case, nothing from the event itself is
passed to the enrichment. For more information, see [The JavaScript Object Notation (JSON)\
Data Interchange Format](http://www.rfc-editor.org/rfc/rfc7159.txt).

To remove an input template, specify an empty string.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `8192`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PipeEnrichmentHttpParameters

PipeLogConfiguration

All content copied from https://docs.aws.amazon.com/.
