This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GroundStation::Config TrackingConfig

Provides information about how AWS Ground Station should track the satellite through the sky during a contact.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Autotrack" : String
}

```

### YAML

```yaml

  Autotrack: String

```

## Properties

`Autotrack`

Specifies whether or not to use autotrack.
`REMOVED` specifies that program track should only be used during the contact.
`PREFERRED` specifies that autotracking is preferred during the contact but fallback to program track if the signal is lost.
`REQUIRED` specifies that autotracking is required during the contact and not to use program track if the signal is lost.

_Required_: No

_Type_: String

_Allowed values_: `REQUIRED | PREFERRED | REMOVED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Create a TrackingConfig

The following example creates a Ground Station `TrackingConfig`

#### JSON

```json

{
  "TrackingConfig": {
    "Autotrack": "PREFERRED"
  }
}
```

#### YAML

```yaml

TrackingConfig:
  Autotrack: "PREFERRED"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TelemetrySinkData

UplinkEchoConfig

All content copied from https://docs.aws.amazon.com/.
