This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationObjectStorage

The `AWS::DataSync::LocationObjectStorage` resource specifies an endpoint
for a self-managed object storage bucket. For more information about self-managed object
storage locations, see [Creating a Location for\
Object Storage](../../../datasync/latest/userguide/create-object-location.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DataSync::LocationObjectStorage",
  "Properties" : {
      "AccessKey" : String,
      "AgentArns" : [ String, ... ],
      "BucketName" : String,
      "CmkSecretConfig" : CmkSecretConfig,
      "CustomSecretConfig" : CustomSecretConfig,
      "SecretKey" : String,
      "ServerCertificate" : String,
      "ServerHostname" : String,
      "ServerPort" : Integer,
      "ServerProtocol" : String,
      "Subdirectory" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DataSync::LocationObjectStorage
Properties:
  AccessKey: String
  AgentArns:
    - String
  BucketName: String
  CmkSecretConfig:
    CmkSecretConfig
  CustomSecretConfig:
    CustomSecretConfig
  SecretKey: String
  ServerCertificate: String
  ServerHostname: String
  ServerPort: Integer
  ServerProtocol: String
  Subdirectory: String
  Tags:
    - Tag

```

## Properties

`AccessKey`

Specifies the access key (for example, a user name) if credentials are required to
authenticate with the object storage server.

_Required_: No

_Type_: String

_Pattern_: `^.+$`

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AgentArns`

(Optional) Specifies the Amazon Resource Names (ARNs) of the DataSync agents
that can connect with your object storage system. If you are setting up an agentless
cross-cloud transfer, you do not need to specify a value for this parameter.

###### Note

Make sure you configure this parameter correctly when you first create your storage
location. You cannot add or remove agents from a storage location after you create
it.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `128 | 4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketName`

Specifies the name of the object storage bucket involved in the transfer.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-\+\./\(\)\$\p{Zs}]+$`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CmkSecretConfig`

Specifies configuration information for a DataSync-managed secret, which
includes the `SecretKey` that DataSync uses to access a specific object
storage location, with a customer-managed AWS KMS key.

When you include this parameter as part of a `CreateLocationObjectStorage`
request, you provide only the KMS key ARN. DataSync uses this KMS key together with the value you specify for the `SecretKey` parameter
to create a DataSync-managed secret to store the location access credentials.

Make sure that DataSync has permission to access the KMS key that
you specify. For more information, see [Using a service-managed secret encrypted with a custom AWS KMS key](../../../datasync/latest/userguide/location-credentials.md#service-secret-custom-key).

###### Note

You can use either `CmkSecretConfig` (with `SecretKey`) or
`CustomSecretConfig` (without `SecretKey`) to provide credentials
for a `CreateLocationObjectStorage` request. Do not provide both parameters for
the same request.

_Required_: No

_Type_: [CmkSecretConfig](aws-properties-datasync-locationobjectstorage-cmksecretconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CustomSecretConfig`

Specifies configuration information for a customer-managed Secrets Manager secret where
the secret key for a specific object storage location is stored in plain text, in Secrets Manager.
This configuration includes the secret ARN, and the ARN for an IAM role that
provides access to the secret. For more information, see [Using a secret that you manage](../../../datasync/latest/userguide/location-credentials.md#custom-secret-custom-key).

###### Note

You can use either `CmkSecretConfig` (with `SecretKey`) or
`CustomSecretConfig` (without `SecretKey`) to provide credentials
for a `CreateLocationObjectStorage` request. Do not provide both parameters for
the same request.

_Required_: No

_Type_: [CustomSecretConfig](aws-properties-datasync-locationobjectstorage-customsecretconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretKey`

Specifies the secret key (for example, a password) if credentials are required to
authenticate with the object storage server.

###### Note

If you provide a secret using `SecretKey`, but do not provide secret
configuration details using `CmkSecretConfig` or `CustomSecretConfig`,
then DataSync stores the token using your AWS account's Secrets Manager secret.

_Required_: No

_Type_: String

_Pattern_: `^.+$`

_Minimum_: `8`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerCertificate`

Specifies a certificate chain for DataSync to authenticate with your object
storage system if the system uses a private or self-signed certificate authority (CA). You
must specify a single `.pem` file with a full certificate chain (for example,
`file:///home/user/.ssh/object_storage_certificates.pem`).

The certificate chain might include:

- The object storage system's certificate

- All intermediate certificates (if there are any)

- The root certificate of the signing CA

You can concatenate your certificates into a `.pem` file (which can be up to
32768 bytes before base64 encoding). The following example `cat` command creates an
`object_storage_certificates.pem` file that includes three certificates:

`cat object_server_certificate.pem intermediate_certificate.pem
        ca_root_certificate.pem > object_storage_certificates.pem`

To use this parameter, configure `ServerProtocol` to `HTTPS`.

_Required_: No

_Type_: String

_Maximum_: `32768`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerHostname`

Specifies the domain name or IP address (IPv4 or IPv6) of the object storage server that
your DataSync agent connects to.

_Required_: No

_Type_: String

_Pattern_: `^(([a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9\-]*[A-Za-z0-9])$`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerPort`

Specifies the port that your object storage server accepts inbound network traffic on (for
example, port 443).

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `65536`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerProtocol`

Specifies the protocol that your object storage server uses to communicate. If not specified, the default
value is `HTTPS`.

_Required_: No

_Type_: String

_Allowed values_: `HTTPS | HTTP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subdirectory`

Specifies the object prefix for your object storage server. If this is a source location,
DataSync only copies objects with this prefix. If this is a destination location,
DataSync writes all objects with this prefix.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-\+\./\(\)\p{Zs}]*$`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Specifies the key-value pair that represents a tag that you want to add to the resource.
Tags can help you manage, filter, and search for your resources. We recommend creating a name
tag for your location.

_Required_: No

_Type_: Array of [Tag](aws-properties-datasync-locationobjectstorage-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref`
function, `Ref` returns the location resource Amazon Resource Name (ARN). For
example:

`arn:aws:datasync:us-east-2:111222333444:location/loc-07db7abfc326c50s3`

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified
attribute of this type. The following are the available attributes and sample return
values.

For more information about using the `Fn::GetAtt` intrinsic function, see
[Fn::GetAtt](../userguide/intrinsic-function-reference-getatt.md).

`CmkSecretConfig.SecretArn`

Property description not available.

`LocationArn`

The Amazon Resource Name (ARN) of the specified object storage location.

`LocationUri`

The URI of the specified object storage location.

## Examples

### Create an object storage location for DataSync

The following example specifies an object storage location for DataSync. In
this example, the object storage location uses the bucket named
`MyBucket`, on the server named
`MyServer@example.com`. This example also specifies the server
protocol `HTTPS` and the subdirectory `/Subdirectory`.

#### JSON

```json

{
"AWSTemplateFormatVersion": "2010-09-09",
"Description": "Specifies an object storage location for DataSync",
"Resources":
{
  "LocationObjectStorage": {
    "Type": "AWS::DataSync::LocationObjectStorage",
    "Properties": {
      "AgentArns": [
        "arn:aws:datasync:us-east-2:111222333444:agent/agent-0b0addbeef44b3nfs"
      ],
      "BucketName": "MyBucket",
      "ServerHostname": "MyServer@example.com",
      "ServerProtocol": "HTTPS",
      "Subdirectory": "/MySubdirectory"
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: Specifies an object storage location for DataSync
Resources:
  LocationObjectStorage:
    Type: AWS::DataSync::LocationObjectStorage
    Properties:
      AgentArns:
        - arn:aws:datasync:us-east-2:111222333444:agent/agent-0b0addbeef44b3nfs
      BucketName: MyBucket
      ServerHostname: MyServer@example.com
      ServerProtocol: HTTPS
      Subdirectory: /MySubdirectory

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

CmkSecretConfig

All content copied from https://docs.aws.amazon.com/.
