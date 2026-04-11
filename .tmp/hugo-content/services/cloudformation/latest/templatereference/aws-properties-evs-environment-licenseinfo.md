This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EVS::Environment LicenseInfo

The license information that Amazon EVS requires to create an environment.
Amazon EVS requires two license keys: a VCF solution key and a vSAN license key.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SolutionKey" : String,
  "VsanKey" : String
}

```

### YAML

```yaml

  SolutionKey: String
  VsanKey: String

```

## Properties

`SolutionKey`

The VCF solution key.
This license unlocks VMware VCF product features, including vSphere, NSX, SDDC Manager, and vCenter Server.
The VCF solution key must cover a minimum of 256 cores.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9]{5}-[a-zA-Z0-9]{5}-[a-zA-Z0-9]{5}-[a-zA-Z0-9]{5}-[a-zA-Z0-9]{5}$`

_Update requires_: Updates are not supported.

`VsanKey`

The VSAN license key.
This license unlocks vSAN features.
The vSAN license key must provide at least 110 TiB of vSAN capacity.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9]{5}-[a-zA-Z0-9]{5}-[a-zA-Z0-9]{5}-[a-zA-Z0-9]{5}-[a-zA-Z0-9]{5}$`

_Update requires_: Updates are not supported.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InitialVlans

Secret

All content copied from https://docs.aws.amazon.com/.
