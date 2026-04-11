This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe PipeEnrichmentHttpParameters

These are custom parameter to be used when the target is an API Gateway REST APIs
or EventBridge ApiDestinations. In the latter case, these are merged with any
InvocationParameters specified on the Connection, with any values from the Connection
taking precedence.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HeaderParameters" : {Key: Value, ...},
  "PathParameterValues" : [ String, ... ],
  "QueryStringParameters" : {Key: Value, ...}
}

```

### YAML

```yaml

  HeaderParameters:
    Key: Value
  PathParameterValues:
    - String
  QueryStringParameters:
    Key: Value

```

## Properties

`HeaderParameters`

The headers that need to be sent as part of request invoking the API Gateway REST
API or EventBridge ApiDestination.

_Required_: No

_Type_: Object of String

_Pattern_: ``^[!#$%&'*+-.^_`|~0-9a-zA-Z]+|(\$(\.[\w/_-]+(\[(\d+|\*)\])*)*)$``

_Minimum_: `0`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PathParameterValues`

The path parameter values to be used to populate API Gateway REST API or EventBridge ApiDestination path wildcards ("\*").

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryStringParameters`

The query string keys/values that need to be sent as part of request invoking the
API Gateway REST API or EventBridge ApiDestination.

_Required_: No

_Type_: Object of String

_Pattern_: `^[^\x00-\x1F\x7F]+|(\$(\.[\w/_-]+(\[(\d+|\*)\])*)*)$`

_Minimum_: `0`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NetworkConfiguration

PipeEnrichmentParameters

All content copied from https://docs.aws.amazon.com/.
