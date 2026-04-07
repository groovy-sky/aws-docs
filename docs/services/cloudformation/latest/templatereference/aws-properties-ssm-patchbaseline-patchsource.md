This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::PatchBaseline PatchSource

`PatchSource` is the property type for the `Sources` resource of
the [AWS::SSM::PatchBaseline](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html) resource.

The AWS CloudFormation `AWS::SSM::PatchSource` resource is used to
provide information about the patches to use to update target instances, including
target operating systems and source repository. Applies to Linux managed nodes
only.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Configuration" : String,
  "Name" : String,
  "Products" : [ String, ... ]
}

```

### YAML

```yaml

  Configuration: String
  Name: String
  Products:
    - String

```

## Properties

`Configuration`

The value of the repo configuration.

**Example for yum repositories**

`[main]`

`name=MyCustomRepository`

`baseurl=https://my-custom-repository`

`enabled=1`

For information about other options available for your yum repository configuration, see
[dnf.conf(5)](https://man7.org/linux/man-pages/man5/dnf.conf.5.html) on the
_man7.org_ website.

**Examples for Ubuntu Server and Debian Server**

`deb http://security.ubuntu.com/ubuntu jammy main`

`deb https://site.example.com/debian distribution component1 component2 component3`

Repo information for Ubuntu Server repositories must be specifed in a single line. For more
examples and information, see [jammy (5)\
sources.list.5.gz](https://manpages.ubuntu.com/manpages/jammy/man5/sources.list.5.html) on the _Ubuntu Server Manuals_ website and [sources.list format](https://wiki.debian.org/SourcesList) on the
_Debian Wiki_.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name specified to identify the patch source.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-.]{3,50}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Products`

The specific operating system versions a patch repository applies to, such as
"Ubuntu16.04", "RedhatEnterpriseLinux7.2" or "Suse12.7". For lists of supported product
values, see [PatchFilter](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PatchFilter.html)
in the _AWS Systems Manager API Reference_.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `128 | 20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PatchFilterGroup

Rule
