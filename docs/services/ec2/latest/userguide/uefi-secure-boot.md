# UEFI Secure Boot for Amazon EC2 instances

UEFI Secure Boot builds on the long-standing secure boot process of Amazon EC2, and provides
additional defense-in-depth that helps customers secure software from threats that
persist across reboots. It ensures that the instance only boots software that is signed
with cryptographic keys. The keys are stored in the key database of the [UEFI non-volatile variable store](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/uefi-variables.html). UEFI Secure Boot
prevents unauthorized modification of the instance boot flow.

###### Contents

- [How UEFI Secure Boot works with Amazon EC2 instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/how-uefi-secure-boot-works.html)

- [Requirements for UEFI Secure Boot on Amazon EC2](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/launch-instance-with-uefi-sb.html)

- [Verify whether an Amazon EC2 instance is enabled for UEFI Secure Boot](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/verify-uefi-secure-boot.html)

- [Create a Linux AMI with custom UEFI Secure Boot keys](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/create-ami-with-uefi-secure-boot.html)

- [Create the AWS binary blob for UEFI Secure Boot](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/aws-binary-blob-creation.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UEFI variables

How UEFI Secure Boot works
