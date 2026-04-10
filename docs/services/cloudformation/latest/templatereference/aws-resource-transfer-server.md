This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::Server

Instantiates an auto-scaling virtual server based on the selected file transfer
protocol in AWS. When you make updates to your file transfer protocol-enabled server
or when you work with users, use the service-generated `ServerId` property
that is assigned to the newly created server.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Transfer::Server",
  "Properties" : {
      "Certificate" : String,
      "Domain" : String,
      "EndpointDetails" : EndpointDetails,
      "EndpointType" : String,
      "IdentityProviderDetails" : IdentityProviderDetails,
      "IdentityProviderType" : String,
      "IpAddressType" : String,
      "LoggingRole" : String,
      "PostAuthenticationLoginBanner" : String,
      "PreAuthenticationLoginBanner" : String,
      "ProtocolDetails" : ProtocolDetails,
      "Protocols" : [ String, ... ],
      "S3StorageOptions" : S3StorageOptions,
      "SecurityPolicyName" : String,
      "StructuredLogDestinations" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "WorkflowDetails" : WorkflowDetails
    }
}

```

### YAML

```yaml

Type: AWS::Transfer::Server
Properties:
  Certificate: String
  Domain: String
  EndpointDetails:
    EndpointDetails
  EndpointType: String
  IdentityProviderDetails:
    IdentityProviderDetails
  IdentityProviderType: String
  IpAddressType: String
  LoggingRole: String
  PostAuthenticationLoginBanner: String
  PreAuthenticationLoginBanner: String
  ProtocolDetails:
    ProtocolDetails
  Protocols:
    - String
  S3StorageOptions:
    S3StorageOptions
  SecurityPolicyName: String
  StructuredLogDestinations:
    - String
  Tags:
    - Tag
  WorkflowDetails:
    WorkflowDetails

```

## Properties

`Certificate`

The Amazon Resource Name (ARN) of the AWS Certificate Manager (ACM) certificate. Required when
`Protocols` is set to `FTPS`.

To request a new public certificate, see [Request a public\
certificate](../../../acm/latest/userguide/gs-acm-request-public.md) in the _AWS Certificate Manager User Guide_.

To import an existing certificate into ACM, see [Importing certificates into\
ACM](../../../acm/latest/userguide/import-certificate.md) in the _AWS Certificate Manager User Guide_.

To request a private certificate to use FTPS through private IP addresses, see [Request\
a private certificate](../../../acm/latest/userguide/gs-acm-request-private.md) in the _AWS Certificate Manager User_
_Guide_.

Certificates with the following cryptographic algorithms and key sizes are
supported:

- 2048-bit RSA (RSA\_2048)

- 4096-bit RSA (RSA\_4096)

- Elliptic Prime Curve 256 bit (EC\_prime256v1)

- Elliptic Prime Curve 384 bit (EC\_secp384r1)

- Elliptic Prime Curve 521 bit (EC\_secp521r1)

###### Note

The certificate must be a valid SSL/TLS X.509 version 3 certificate with FQDN or
IP address specified and information about the issuer.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Domain`

Specifies the domain of the storage system that is used for file transfers. There are
two domains available: Amazon Simple Storage Service (Amazon S3) and Amazon Elastic File
System (Amazon EFS). The default value is S3.

_Required_: No

_Type_: String

_Allowed values_: `S3 | EFS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EndpointDetails`

The virtual private cloud (VPC) endpoint settings that are configured for your server.
When you host your endpoint within your VPC, you can make your endpoint accessible only to resources
within your VPC, or you can attach Elastic IP addresses and make your endpoint accessible to clients over
the internet. Your VPC's default security groups are automatically assigned to your
endpoint.

_Required_: No

_Type_: [EndpointDetails](aws-properties-transfer-server-endpointdetails.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`EndpointType`

The type of endpoint that you want your server to use. You can choose to make your
server's endpoint publicly accessible (PUBLIC) or host it inside your VPC. With an
endpoint that is hosted in a VPC, you can restrict access to your server and resources
only within your VPC or choose to make it internet facing by attaching Elastic IP
addresses directly to it.

###### Note

After May 19, 2021, you won't be able to create a server using
`EndpointType=VPC_ENDPOINT` in your AWS account if
your account hasn't already done so before May 19, 2021. If you have already created
servers with `EndpointType=VPC_ENDPOINT` in your AWS account on or before May 19, 2021, you will not be affected. After this date, use
`EndpointType` = `VPC` .

For more information, see [Discontinuing the use of VPC\_ENDPOINT](../../../transfer/latest/userguide/create-server-in-vpc.md#deprecate-vpc-endpoint) .

It is recommended that you use `VPC` as the `EndpointType`
. With this endpoint type, you have the option to directly associate up to three
Elastic IPv4 addresses (BYO IP included) with your server's endpoint and use VPC
security groups to restrict traffic by the client's public IP address. This is not
possible with `EndpointType` set to `VPC_ENDPOINT` .

_Required_: No

_Type_: String

_Allowed values_: `PUBLIC | VPC | VPC_ENDPOINT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentityProviderDetails`

Required when `IdentityProviderType` is set to
`AWS_DIRECTORY_SERVICE`, `AWS_LAMBDA` or
`API_GATEWAY`. Accepts an array containing all of the information
required to use a directory in `AWS_DIRECTORY_SERVICE` or invoke a
customer-supplied authentication API, including the API Gateway URL. Cannot be specified
when `IdentityProviderType` is set to `SERVICE_MANAGED`.

_Required_: No

_Type_: [IdentityProviderDetails](aws-properties-transfer-server-identityproviderdetails.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdentityProviderType`

The mode of authentication for a server. The default value is
`SERVICE_MANAGED`, which allows you to store and access user credentials within
the AWS Transfer Family service.

Use `AWS_DIRECTORY_SERVICE` to provide access to
Active Directory groups in AWS Directory Service for Microsoft Active Directory or Microsoft Active Directory in your
on-premises environment or in AWS using AD Connector. This option also requires you to
provide a Directory ID by using the `IdentityProviderDetails` parameter.

Use the `API_GATEWAY` value to integrate with an identity provider of your choosing. The
`API_GATEWAY` setting requires you to provide an Amazon API Gateway endpoint URL to call
for authentication by using the `IdentityProviderDetails` parameter.

Use the `AWS_LAMBDA` value to directly use an AWS Lambda function as your identity provider.
If you choose this value, you must specify the ARN for the Lambda function in the `Function` parameter
for the `IdentityProviderDetails` data type.

_Required_: No

_Type_: String

_Allowed values_: `SERVICE_MANAGED | API_GATEWAY | AWS_DIRECTORY_SERVICE | AWS_LAMBDA`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpAddressType`

Specifies whether to use IPv4 only, or to use dual-stack (IPv4 and IPv6) for your AWS Transfer Family endpoint. The default value is `IPV4`.

###### Important

The `IpAddressType` parameter has the following limitations:

- It cannot be changed while the server is online. You must stop the server before modifying this parameter.

- It cannot be updated to `DUALSTACK` if the server has `AddressAllocationIds` specified.

###### Note

When using `DUALSTACK` as the `IpAddressType`, you cannot set the `AddressAllocationIds` parameter for the [EndpointDetails](../../../../reference/transfer/latest/apireference/api-endpointdetails.md) for the server.

_Required_: No

_Type_: String

_Allowed values_: `IPV4 | DUALSTACK`

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`LoggingRole`

The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that allows a server to turn
on Amazon CloudWatch logging for Amazon S3 or Amazon EFS events. When set, you can view user activity in
your CloudWatch logs.

_Required_: No

_Type_: String

_Pattern_: `^(|arn:.*role/\S+)$`

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PostAuthenticationLoginBanner`

Specifies a string to display when users connect to a server. This string is displayed after the user authenticates.

###### Note

The SFTP protocol does not support post-authentication display banners.

_Required_: No

_Type_: String

_Pattern_: `^[\x09-\x0D\x20-\x7E]*$`

_Minimum_: `0`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreAuthenticationLoginBanner`

Specifies a string to display when users connect to a server. This string is displayed before the user authenticates.
For example, the following banner displays details about using the system:

`This system is for the use of authorized users only. Individuals using this computer system without authority,
    or in excess of their authority, are subject to having all of their activities on this system monitored and recorded by
    system personnel.`

_Required_: No

_Type_: String

_Pattern_: `^[\x09-\x0D\x20-\x7E]*$`

_Minimum_: `0`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProtocolDetails`

The protocol settings that are configured for your server.

- To indicate passive mode (for FTP and FTPS protocols), use the
`PassiveIp` parameter. Enter a single dotted-quad IPv4 address,
such as the external IP address of a firewall, router, or load balancer.

- To ignore the error that is generated when the client attempts to use the
`SETSTAT` command on a file that you are uploading to an Amazon
S3 bucket, use the `SetStatOption` parameter. To have the AWS Transfer Family server ignore the `SETSTAT` command and upload
files without needing to make any changes to your SFTP client, set the value to
`ENABLE_NO_OP` . If you set the `SetStatOption`
parameter to `ENABLE_NO_OP` , Transfer Family generates a log entry
to Amazon CloudWatch Logs, so that you can determine when the client is making a
`SETSTAT` call.

- To determine whether your AWS Transfer Family server resumes recent,
negotiated sessions through a unique session ID, use the
`TlsSessionResumptionMode` parameter.

- `As2Transports` indicates the transport method for the AS2
messages. Currently, only HTTP is supported.

The `Protocols` parameter is an array of strings.

_Allowed values_ : One or more of `SFTP` ,
`FTPS` , `FTP` , `AS2`

_Required_: No

_Type_: [ProtocolDetails](aws-properties-transfer-server-protocoldetails.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocols`

Specifies the file transfer protocol or protocols over which your file transfer
protocol client can connect to your server's endpoint. The available protocols
are:

- `SFTP` (Secure Shell (SSH) File Transfer Protocol): File transfer
over SSH

- `FTPS` (File Transfer Protocol Secure): File transfer with TLS
encryption

- `FTP` (File Transfer Protocol): Unencrypted file transfer

- `AS2` (Applicability Statement 2): used for transporting structured
business-to-business data

###### Note

- If you select `FTPS` , you must choose a certificate stored in
AWS Certificate Manager (ACM) which is used to identify your server when
clients connect to it over FTPS.

- If `Protocol` includes either `FTP` or
`FTPS` , then the `EndpointType` must be
`VPC` and the `IdentityProviderType` must be
either `AWS_DIRECTORY_SERVICE` , `AWS_LAMBDA` , or
`API_GATEWAY` .

- If `Protocol` includes `FTP` , then
`AddressAllocationIds` cannot be associated.

- If `Protocol` is set only to `SFTP` , the
`EndpointType` can be set to `PUBLIC` and the
`IdentityProviderType` can be set any of the supported
identity types: `SERVICE_MANAGED` ,
`AWS_DIRECTORY_SERVICE` , `AWS_LAMBDA` , or
`API_GATEWAY` .

- If `Protocol` includes `AS2` , then the
`EndpointType` must be `VPC` , and domain must be
Amazon S3.

The `Protocols` parameter is an array of strings.

_Allowed values_ : One or more of `SFTP` ,
`FTPS` , `FTP` , `AS2`

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3StorageOptions`

Specifies whether or not performance for your Amazon S3 directories is optimized.

- If using the console, this is enabled by default.

- If using the API or CLI, this is disabled by default.

By default, home directory mappings have a `TYPE` of `DIRECTORY`. If you enable this option, you would then need to explicitly set the `HomeDirectoryMapEntry` `Type` to `FILE` if you want a mapping to have a file target.

_Required_: No

_Type_: [S3StorageOptions](aws-properties-transfer-server-s3storageoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityPolicyName`

Specifies the name of the security policy for the server.

_Required_: No

_Type_: String

_Pattern_: `^TransferSecurityPolicy-.+$`

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StructuredLogDestinations`

Specifies the log groups to which your server logs are sent.

To specify a log group, you must provide the ARN for an existing log group. In this case, the format of the log group is as follows:

`arn:aws:logs:region-name:amazon-account-id:log-group:log-group-name:*`

For example, `arn:aws:logs:us-east-1:111122223333:log-group:mytestgroup:*`

If you have previously specified a log group for a server, you can clear it, and in effect turn off structured logging, by providing an empty
value for this parameter in an `update-server` call. For example:

`update-server --server-id s-1234567890abcdef0 --structured-log-destinations`

_Required_: No

_Type_: Array of String

_Minimum_: `20 | 0`

_Maximum_: `1600 | 1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Key-value pairs that can be used to group and search for servers.

_Required_: No

_Type_: Array of [Tag](aws-properties-transfer-server-tag.md)

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkflowDetails`

Specifies the workflow ID for the workflow to assign and the execution role that's
used for executing the workflow.

In addition to a workflow to execute when a file is uploaded completely,
`WorkflowDetails` can also contain a workflow ID (and execution role) for
a workflow to execute on partial upload. A partial upload occurs when a file is open
when the session disconnects.

_Required_: No

_Type_: [WorkflowDetails](aws-properties-transfer-server-workflowdetails.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the server ARN, such as
`arn:aws:transfer:us-east-1:123456789012:server/s-01234567890abcdef` .

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

Specifies the unique Amazon Resource Name (ARN) of the server.

`As2ServiceManagedEgressIpAddresses`

The list of egress IP addresses of this server. These IP addresses are only relevant
for servers that use the AS2 protocol. They are used for sending asynchronous
MDNs.

These IP addresses are assigned automatically when you create an AS2 server.
Additionally, if you update an existing server and add the AS2 protocol, static IP
addresses are assigned as well.

`ServerId`

Specifies the unique system-assigned identifier for a server that you
instantiate.

`State`

The condition of the server that was described. A value of
`ONLINE` indicates that the server can accept jobs and transfer files. A
`State` value of `OFFLINE` means that the server cannot perform file
transfer operations.

The states of `STARTING` and `STOPPING` indicate that the server is
in an intermediate state, either not fully able to respond, or not fully offline. The values
of `START_FAILED` or `STOP_FAILED` can indicate an error
condition.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

EndpointDetails

All content copied from https://docs.aws.amazon.com/.
