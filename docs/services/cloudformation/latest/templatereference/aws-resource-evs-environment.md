This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EVS::Environment

Creates an Amazon EVS environment that runs VCF software, such as SDDC Manager, NSX Manager, and vCenter Server.

During environment creation, Amazon EVS performs validations on DNS settings, provisions VLAN subnets and hosts, and deploys the supplied version of VCF.

It can take several hours to create an environment.
After the deployment completes, you can configure VCF in the vSphere user interface according to your needs.

###### Important

When creating a new environment, the default ESX version for the selected VCF version will be used, you cannot choose a specific ESX version in `CreateEnvironment` action. When a host has been added with a specific ESX version, it can only be upgraded using vCenter Lifecycle Manager.

###### Note

You cannot use the `dedicatedHostId` and `placementGroupId` parameters together in the same `CreateEnvironment` action.
This results in a `ValidationException` response.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EVS::Environment",
  "Properties" : {
      "ConnectivityInfo" : ConnectivityInfo,
      "EnvironmentName" : String,
      "Hosts" : [ HostInfoForCreate, ... ],
      "InitialVlans" : InitialVlans,
      "KmsKeyId" : String,
      "LicenseInfo" : LicenseInfo,
      "ServiceAccessSecurityGroups" : ServiceAccessSecurityGroups,
      "ServiceAccessSubnetId" : String,
      "SiteId" : String,
      "Tags" : [ Tag, ... ],
      "TermsAccepted" : Boolean,
      "VcfHostnames" : VcfHostnames,
      "VcfVersion" : String,
      "VpcId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EVS::Environment
Properties:
  ConnectivityInfo:
    ConnectivityInfo
  EnvironmentName: String
  Hosts:
    - HostInfoForCreate
  InitialVlans:
    InitialVlans
  KmsKeyId: String
  LicenseInfo:
    LicenseInfo
  ServiceAccessSecurityGroups:
    ServiceAccessSecurityGroups
  ServiceAccessSubnetId: String
  SiteId: String
  Tags:
    - Tag
  TermsAccepted: Boolean
  VcfHostnames:
    VcfHostnames
  VcfVersion: String
  VpcId: String

```

## Properties

`ConnectivityInfo`

The connectivity configuration for the environment.
Amazon EVS requires that you specify two route server peer IDs.
During environment creation, the route server endpoints peer with the NSX uplink VLAN for connectivity to the NSX overlay network.

_Required_: Yes

_Type_: [ConnectivityInfo](aws-properties-evs-environment-connectivityinfo.md)

_Update requires_: Updates are not supported.

`EnvironmentName`

The name of the environment.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]{1,100}$`

_Update requires_: Updates are not supported.

`Hosts`

Required for environment resource creation.

_Required_: Conditional

_Type_: Array of [HostInfoForCreate](aws-properties-evs-environment-hostinfoforcreate.md)

_Minimum_: `4`

_Maximum_: `4`

_Update requires_: Updates are not supported.

`InitialVlans`

The initial VLAN subnets for the environment.
Amazon EVS VLAN subnets have a minimum CIDR block size of /28 and a maximum size of /24.
Amazon EVS VLAN subnet CIDR blocks must not overlap with other subnets in the VPC.

Required for environment resource creation.

_Required_: Conditional

_Type_: [InitialVlans](aws-properties-evs-environment-initialvlans.md)

_Update requires_: Updates are not supported.

`KmsKeyId`

The AWS KMS key ID that AWS Secrets Manager uses to encrypt secrets
that are associated with the environment.
These secrets contain the VCF credentials that are needed to install vCenter Server, NSX, and SDDC Manager.

By default, Amazon EVS use the AWS Secrets Manager managed key `aws/secretsmanager`.
You can also specify a customer managed key.

_Required_: No

_Type_: String

_Update requires_: Updates are not supported.

`LicenseInfo`

The license information that Amazon EVS requires to create an environment.
Amazon EVS requires two license keys: a VCF solution key and a vSAN license key.
The VCF solution key must cover a minimum of 256 cores.
The vSAN license key must provide at least 110 TiB of vSAN capacity.

_Required_: Yes

_Type_: [LicenseInfo](aws-properties-evs-environment-licenseinfo.md)

_Update requires_: Updates are not supported.

`ServiceAccessSecurityGroups`

The security groups that allow traffic between the Amazon EVS control plane and your VPC for service access.
If a security group is not specified, Amazon EVS uses the default security group in your account for service access.

_Required_: No

_Type_: [ServiceAccessSecurityGroups](aws-properties-evs-environment-serviceaccesssecuritygroups.md)

_Update requires_: Updates are not supported.

`ServiceAccessSubnetId`

The subnet that is used to establish connectivity between the Amazon EVS control plane and VPC.
Amazon EVS uses this subnet to perform validations and create the environment.

_Required_: Yes

_Type_: String

_Pattern_: `^subnet-[a-f0-9]{8}([a-f0-9]{9})?$`

_Minimum_: `15`

_Maximum_: `24`

_Update requires_: Updates are not supported.

`SiteId`

The Broadcom Site ID that is associated with your Amazon EVS environment.
Amazon EVS uses the Broadcom Site ID that you provide to meet Broadcom VCF license usage reporting requirements for Amazon EVS.

_Required_: Yes

_Type_: String

_Update requires_: Updates are not supported.

`Tags`

Metadata that assists with categorization and organization.
Each tag consists of a key and an optional value.
You define both.
Tags don't propagate to any other cluster or AWS resources.

_Required_: No

_Type_: Array of [Tag](aws-properties-evs-environment-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TermsAccepted`

Customer confirmation that the customer has purchased and will continue to maintain the required number of VCF software licenses to cover all physical processor cores in the Amazon EVS environment.
Information about your VCF software in Amazon EVS will be shared with Broadcom to verify license compliance.
Amazon EVS does not validate license keys.
To validate license keys, visit the Broadcom support portal.

_Required_: Yes

_Type_: Boolean

_Update requires_: Updates are not supported.

`VcfHostnames`

The DNS hostnames to be used by the VCF management appliances in your environment.

For environment creation to be successful, each hostname entry must resolve to a domain name
that you've registered in your DNS service of choice and configured in the DHCP option set of your VPC.
DNS hostnames cannot be changed after environment creation has started.

_Required_: Yes

_Type_: [VcfHostnames](aws-properties-evs-environment-vcfhostnames.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VcfVersion`

The VCF version of the environment.

_Required_: Yes

_Type_: String

_Allowed values_: `VCF-5.2.1 | VCF-5.2.2`

_Update requires_: Updates are not supported.

`VpcId`

The VPC associated with the environment.

_Required_: Yes

_Type_: String

_Pattern_: `^vpc-[a-f0-9]{8}([a-f0-9]{9})?$`

_Minimum_: `12`

_Maximum_: `21`

_Update requires_: Updates are not supported.

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon EVS `EnvironmentId`. For example: `{ "Ref": "env-1234567890" }`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Checks`

A check on the environment to identify instance health and VMware VCF licensing issues. For example:

```

            {
                "checks": [
                    {
                        "type": "KEY_REUSE",
                        "result": "PASSED"
                    },
                    {
                        "type": "KEY_COVERAGE",
                        "result": "PASSED"
                    },
                    {
                        "type": "REACHABILITY",
                        "result": "PASSED"
                    },
                    {
                        "type": "HOST_COUNT",
                        "result": "PASSED"
                    }
                ]
            }

```

`CreatedAt`

The date and time that the environment was created. For example: `1749081600.000`.

`Credentials`

The VCF credentials that are stored as Amazon EVS managed secrets in AWS Secrets Manager.
Amazon EVS stores credentials that are needed to install vCenter Server, NSX, and SDDC Manager. For example:

```

            {
                [
                    {
                        "secretArn": "arn:aws:secretsmanager:us-east-1:000000000000:secret:evs!env-1234567890_vCenterAdmin-MnTMEi"
                    },
                    {
                        "secretArn": "arn:aws:secretsmanager:us-east-1:000000000000:secret:evs!env-1234567890_vCenterRoot-87VyCF"
                    },
                    {
                        "secretArn": "arn:aws:secretsmanager:us-east-1:000000000000:secret:evs!env-1234567890_NSXRoot-SR3k43"
                    },
                    {
                        "secretArn": "arn:aws:secretsmanager:us-east-1:000000000000:secret:evs!env-1234567890_NSXAdmin-L5LUiD"
                    },
                    {
                        "secretArn": "arn:aws:secretsmanager:us-east-1:000000000000:secret:evs!env-1234567890_NSXAudit-Q2oW46"
                    },
                    {
                        "secretArn": "arn:aws:secretsmanager:us-east-1:000000000000:secret:evs!env-1234567890_SDDCManagerRoot-bFulOq"
                    },
                    {
                        "secretArn": "arn:aws:secretsmanager:us-east-1:000000000000:secret:evs!env-1234567890_SDDCManagerVCF-Ec3gES"
                    },
                    {
                        "secretArn": "arn:aws:secretsmanager:us-east-1:000000000000:secret:evs!env-1234567890_SDDCManagerAdmin-JMTAAb"
                    }
                ]
            }

```

`EnvironmentArn`

The Amazon Resource Name (ARN) that is associated with the environment. For example: `arn:aws:evs:us-east-1:000000000000:environment/env-1234567890`.

`EnvironmentId`

The unique ID for the environment. For example: `env-1234567890`.

`EnvironmentState`

The state of an environment. For example: `CREATED`.

`ModifiedAt`

The date and time that the environment was modified. For example: `1749081600.000`.

`StateDetails`

A detailed description of the `environmentState` of an environment. For example: `Environment successfully created`.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Elastic VMware Service

Check

All content copied from https://docs.aws.amazon.com/.
