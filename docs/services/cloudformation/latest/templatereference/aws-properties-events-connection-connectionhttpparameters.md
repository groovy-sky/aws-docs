This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Connection ConnectionHttpParameters

Any additional parameters for the connection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BodyParameters" : [ Parameter, ... ],
  "HeaderParameters" : [ Parameter, ... ],
  "QueryStringParameters" : [ Parameter, ... ]
}

```

### YAML

```yaml

  BodyParameters:
    - Parameter
  HeaderParameters:
    - Parameter
  QueryStringParameters:
    - Parameter

```

## Properties

`BodyParameters`

Any additional body string parameters for the connection.

_Required_: No

_Type_: Array of [Parameter](aws-properties-events-connection-parameter.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HeaderParameters`

Any additional header parameters for the connection.

_Required_: No

_Type_: Array of [Parameter](aws-properties-events-connection-parameter.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryStringParameters`

Any additional query string parameters for the connection.

_Required_: No

_Type_: Array of [Parameter](aws-properties-events-connection-parameter.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClientParameters

ConnectivityParameters

All content copied from https://docs.aws.amazon.com/.
