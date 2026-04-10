This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::MultiRegionEndpoint Details

An object that contains configuration details of multi-region endpoint (global-endpoint).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RouteDetails" : [ RouteDetailsItems, ... ]
}

```

### YAML

```yaml

  RouteDetails:
    - RouteDetailsItems

```

## Properties

`RouteDetails`

A list of route configuration details. Must contain exactly one route configuration.

_Required_: Yes

_Type_: Array of [RouteDetailsItems](aws-properties-ses-multiregionendpoint-routedetailsitems.md)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SES::MultiRegionEndpoint

RouteDetailsItems

All content copied from https://docs.aws.amazon.com/.
