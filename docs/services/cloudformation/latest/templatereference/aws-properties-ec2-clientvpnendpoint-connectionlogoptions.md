This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::ClientVpnEndpoint ConnectionLogOptions

Describes the client connection logging options for the Client VPN endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudwatchLogGroup" : String,
  "CloudwatchLogStream" : String,
  "Enabled" : Boolean
}

```

### YAML

```yaml

  CloudwatchLogGroup: String
  CloudwatchLogStream: String
  Enabled: Boolean

```

## Properties

`CloudwatchLogGroup`

The name of the CloudWatch Logs log group. Required if connection logging is enabled.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CloudwatchLogStream`

The name of the CloudWatch Logs log stream to which the connection data is published.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

Indicates whether connection logging is enabled.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ClientRouteEnforcementOptions

DirectoryServiceAuthenticationRequest
