This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchService::Domain ZoneAwarenessConfig

Specifies zone awareness configuration options. Only use if
`ZoneAwarenessEnabled` is `true`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AvailabilityZoneCount" : Integer
}

```

### YAML

```yaml

  AvailabilityZoneCount: Integer

```

## Properties

`AvailabilityZoneCount`

If you enabled multiple Availability Zones (AZs), the number of AZs that you want the
domain to use.

Valid values are `2` and `3`. Default is 2.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WindowStartTime

Next

All content copied from https://docs.aws.amazon.com/.
