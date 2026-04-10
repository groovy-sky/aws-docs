This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::AutomationRule SeverityUpdate

Updates to the severity information for a finding.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Label" : String,
  "Normalized" : Integer,
  "Product" : Number
}

```

### YAML

```yaml

  Label: String
  Normalized: Integer
  Product: Number

```

## Properties

`Label`

The severity value of the finding. The allowed values are the following.

- `INFORMATIONAL` \- No issue was found.

- `LOW` \- The issue does not require action on its own.

- `MEDIUM` \- The issue must be addressed but not urgently.

- `HIGH` \- The issue must be addressed as a priority.

- `CRITICAL` \- The issue must be remediated immediately to avoid it
escalating.

_Required_: No

_Type_: String

_Allowed values_: `INFORMATIONAL | LOW | MEDIUM | HIGH | CRITICAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Normalized`

The normalized severity for the finding. This attribute is to be deprecated in favor of
`Label`.

If you provide `Normalized` and don't provide `Label`,
`Label` is set automatically as follows.

- 0 - `INFORMATIONAL`

- 1–39 - `LOW`

- 40–69 - `MEDIUM`

- 70–89 - `HIGH`

- 90–100 - `CRITICAL`

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Product`

The native severity as defined by the AWS service or integrated partner product that
generated the finding.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RelatedFinding

StringFilter

All content copied from https://docs.aws.amazon.com/.
