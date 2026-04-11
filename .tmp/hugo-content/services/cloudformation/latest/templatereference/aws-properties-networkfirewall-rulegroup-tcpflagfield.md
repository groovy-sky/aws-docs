This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::NetworkFirewall::RuleGroup TCPFlagField

TCP flags and masks to inspect packets for. This is used in the match attributes
specification.

For example:

`"TCPFlags": [
        {
            "Flags": [
                "ECE",
                "SYN"
            ],
            "Masks": [
                "SYN",
                "ECE"
            ]
        }
             ]`

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Flags" : [ String, ... ],
  "Masks" : [ String, ... ]
}

```

### YAML

```yaml

  Flags:
    - String
  Masks:
    - String

```

## Properties

`Flags`

Used in conjunction with the `Masks` setting to define the flags that must be set and flags that must not be set in order for the packet to match. This setting can only specify values that are also specified in the `Masks` setting.

For the flags that are specified in the masks setting, the following must be true for the packet to match:

- The ones that are set in this flags setting must be set in the packet.

- The ones that are not set in this flags setting must also not be set in the packet.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Masks`

The set of flags to consider in the inspection. To inspect all flags in the valid values list, leave this with no setting.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::NetworkFirewall::TLSInspectionConfiguration

All content copied from https://docs.aws.amazon.com/.
