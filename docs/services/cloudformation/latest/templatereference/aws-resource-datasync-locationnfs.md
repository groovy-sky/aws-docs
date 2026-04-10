This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationNFS

The `AWS::DataSync::LocationNFS` resource specifies a Network File System
(NFS) file server that AWS DataSync can use as a transfer source or
destination.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataSync::LocationNFS",
  "Properties" : {
      "MountOptions" : MountOptions,
      "OnPremConfig" : OnPremConfig,
      "ServerHostname" : String,
      "Subdirectory" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DataSync::LocationNFS
Properties:
  MountOptions:
    MountOptions
  OnPremConfig:
    OnPremConfig
  ServerHostname: String
  Subdirectory: String
  Tags:
    - Tag

```

## Properties

`MountOptions`

Specifies the options that DataSync can use to mount your NFS file
server.

_Required_: No

_Type_: [MountOptions](aws-properties-datasync-locationnfs-mountoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OnPremConfig`

Specifies the Amazon Resource Name (ARN) of the DataSync agent that can
connect to your NFS file server.

You can specify more than one agent. For more information, see [Using multiple DataSync agents](../../../datasync/latest/userguide/do-i-need-datasync-agent.md#multiple-agents).

_Required_: Yes

_Type_: [OnPremConfig](aws-properties-datasync-locationnfs-onpremconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerHostname`

Specifies the DNS name or IP address (IPv4 or IPv6) of the NFS file server that your DataSync agent connects to.

_Required_: No

_Type_: String

_Pattern_: `^(([a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9\-]*[A-Za-z0-9])$`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subdirectory`

Specifies the export path in your NFS file server that you want DataSync to
mount.

This path (or a subdirectory of the path) is where DataSync transfers data to
or from. For information on configuring an export for DataSync, see [Accessing NFS file servers](../../../datasync/latest/userguide/create-nfs-location.md#accessing-nfs).

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-\+\./\(\)\$\p{Zs}]+$`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Specifies labels that help you categorize, filter, and search for your AWS resources. We recommend creating at least a name tag for your location.

_Required_: No

_Type_: Array of [Tag](aws-properties-datasync-locationnfs-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref`
function, `Ref` returns the location resource ARN. For example:

`arn:aws:datasync:us-east-2:111222333444:location/loc-07db7abfc326c50s3`

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified
attribute of this type. The following are the available attributes and sample return
values.

For more information about using the `Fn::GetAtt` intrinsic function, see
[Fn::GetAtt](../userguide/intrinsic-function-reference-getatt.md).

`LocationArn`

The Amazon Resource Name (ARN) of the NFS location that you created.

`LocationUri`

The URI of the NFS location that you created.

## Examples

### Create an NFS location for DataSync

The following example specifies an NFS location for DataSync, using a source
and destination location. In this example, the server hostname is
`MyServer@example.com`, using NFS version 4.0, in the
subdirectory `/MySubdirectory`.

#### JSON

```json

{
    "Resources": {
        "LocationNFS": {
            "Type": "AWS::DataSync::LocationNFS",
            "Properties": {
                "MountOptions": {
                    "Version": "NFS4_0"
                },
                "OnPremConfig": {
                    "AgentArns": [
                        "arn:aws:datasync:us-east-2:111222333444:agent/agent-000addbcdf44bbnfs"
                    ]
                },
                "ServerHostname": "MyServer@example.com",
                "Subdirectory": "/MySubdirectory"
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  LocationNFS:
    Type: 'AWS::DataSync::LocationNFS'
    Properties:
      MountOptions:
        Version: NFS4_0
      OnPremConfig:
        AgentArns:
          - >-
            arn:aws:datasync:us-east-2:111222333444:agent/agent-000addbcdf44bbnfs
      ServerHostname: MyServer@example.com
      Subdirectory: /MySubdirectory
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

MountOptions

All content copied from https://docs.aws.amazon.com/.
