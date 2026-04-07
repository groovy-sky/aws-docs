This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EVS::Environment VcfHostnames

The DNS hostnames that Amazon EVS uses to install VMware vCenter Server, NSX, SDDC Manager, and Cloud Builder.
Each hostname must be unique, and resolve to a domain name that you've registered in your DNS service of choice.
Hostnames cannot be changed.

VMware VCF requires the deployment of two NSX Edge nodes, and three NSX Manager virtual machines.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CloudBuilder" : String,
  "Nsx" : String,
  "NsxEdge1" : String,
  "NsxEdge2" : String,
  "NsxManager1" : String,
  "NsxManager2" : String,
  "NsxManager3" : String,
  "SddcManager" : String,
  "VCenter" : String
}

```

### YAML

```yaml

  CloudBuilder: String
  Nsx: String
  NsxEdge1: String
  NsxEdge2: String
  NsxManager1: String
  NsxManager2: String
  NsxManager3: String
  SddcManager: String
  VCenter: String

```

## Properties

`CloudBuilder`

The hostname for VMware Cloud Builder.

_Required_: Yes

_Type_: String

_Pattern_: `^([a-zA-Z0-9\-]*)$`

_Update requires_: Updates are not supported.

`Nsx`

The VMware NSX hostname.

_Required_: Yes

_Type_: String

_Pattern_: `^([a-zA-Z0-9\-]*)$`

_Update requires_: Updates are not supported.

`NsxEdge1`

The hostname for the first NSX Edge node.

_Required_: Yes

_Type_: String

_Pattern_: `^([a-zA-Z0-9\-]*)$`

_Update requires_: Updates are not supported.

`NsxEdge2`

The hostname for the second NSX Edge node.

_Required_: Yes

_Type_: String

_Pattern_: `^([a-zA-Z0-9\-]*)$`

_Update requires_: Updates are not supported.

`NsxManager1`

The hostname for the first VMware NSX Manager virtual machine (VM).

_Required_: Yes

_Type_: String

_Pattern_: `^([a-zA-Z0-9\-]*)$`

_Update requires_: Updates are not supported.

`NsxManager2`

The hostname for the second VMware NSX Manager virtual machine (VM).

_Required_: Yes

_Type_: String

_Pattern_: `^([a-zA-Z0-9\-]*)$`

_Update requires_: Updates are not supported.

`NsxManager3`

The hostname for the third VMware NSX Manager virtual machine (VM).

_Required_: Yes

_Type_: String

_Pattern_: `^([a-zA-Z0-9\-]*)$`

_Update requires_: Updates are not supported.

`SddcManager`

The hostname for SDDC Manager.

_Required_: Yes

_Type_: String

_Pattern_: `^([a-zA-Z0-9\-]*)$`

_Update requires_: Updates are not supported.

`VCenter`

The VMware vCenter hostname.

_Required_: Yes

_Type_: String

_Pattern_: `^([a-zA-Z0-9\-]*)$`

_Update requires_: Updates are not supported.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Next
