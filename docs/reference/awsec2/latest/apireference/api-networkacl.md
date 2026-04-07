# NetworkAcl

Describes a network ACL.

## Contents

**AssociationSet.N**

Any associations between the network ACL and your subnets

Type: Array of [NetworkAclAssociation](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_NetworkAclAssociation.html) objects

Required: No

**default**

Indicates whether this is the default network ACL for the VPC.

Type: Boolean

Required: No

**EntrySet.N**

The entries (rules) in the network ACL.

Type: Array of [NetworkAclEntry](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_NetworkAclEntry.html) objects

Required: No

**networkAclId**

The ID of the network ACL.

Type: String

Required: No

**ownerId**

The ID of the AWS account that owns the network ACL.

Type: String

Required: No

**TagSet.N**

Any tags assigned to the network ACL.

Type: Array of [Tag](api-tag.md) objects

Required: No

**vpcId**

The ID of the VPC for the network ACL.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/NetworkAcl)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/NetworkAcl)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/NetworkAcl)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

NativeApplicationOidcOptions

NetworkAclAssociation
