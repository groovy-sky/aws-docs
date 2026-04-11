This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::ResourceDataSync AwsOrganizationsSource

Information about the `AwsOrganizationsSource` resource data sync source. A sync
source of this type can synchronize data from AWS Organizations or, if an AWS organization isn't
present, from multiple AWS Regions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OrganizationalUnits" : [ String, ... ],
  "OrganizationSourceType" : String
}

```

### YAML

```yaml

  OrganizationalUnits:
    - String
  OrganizationSourceType: String

```

## Properties

`OrganizationalUnits`

The AWS Organizations organization units included in the sync.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrganizationSourceType`

If an AWS organization is present, this is either `OrganizationalUnits` or
`EntireOrganization`. For `OrganizationalUnits`, the data is aggregated
from a set of organization units. For `EntireOrganization`, the data is aggregated
from the entire AWS organization.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SSM::ResourceDataSync

S3Destination

All content copied from https://docs.aws.amazon.com/.
