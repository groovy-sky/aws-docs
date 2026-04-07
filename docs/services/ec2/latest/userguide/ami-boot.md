# Instance launch behavior with Amazon EC2 boot modes

When a computer boots, the first software that it runs is responsible for initializing the
platform and providing an interface for the operating system to perform platform-specific
operations.

In Amazon EC2, two variants of the boot mode software are supported: Unified Extensible Firmware Interface (UEFI) and Legacy BIOS.

###### Possible boot mode parameters on an AMI

An AMI can have one of the following boot mode parameter values: `uefi`,
`legacy-bios`, or `uefi-preferred`. The AMI boot mode
parameter is optional. For AMIs with no boot mode parameter, the instances launched from
these AMIs use the default boot mode value of the instance type.

###### Purpose of the AMI boot mode parameter

The AMI boot mode parameter signals to Amazon EC2 which boot mode to use when launching an
instance. When the boot mode parameter is set to `uefi`, EC2 attempts to
launch the instance on UEFI. If the operating system is not configured to support UEFI,
the instance launch will be unsuccessful.

###### UEFI Preferred boot mode parameter

You can create AMIs that support both UEFI and Legacy BIOS by using the `uefi-preferred` boot mode parameter.
When the boot mode parameter is set to `uefi-preferred`, and if the instance type supports UEFI, the instance
is launched on UEFI. If the instance type does not support UEFI, the instance is launched on Legacy BIOS.

###### Warning

Some features, like UEFI Secure Boot, are only available on instances that boot on UEFI.
When you use the `uefi-preferred` AMI boot mode parameter with an
instance type that does not support UEFI, the instance will launch as Legacy BIOS
and the UEFI-dependent feature will be disabled. If you rely on the availability of
a UEFI-dependent feature, set your AMI boot mode parameter to
`uefi`.

###### Default boot modes for instance types

- Graviton instance types: UEFI

- Intel and AMD instance types: Legacy BIOS

###### Zone support

UEFI boot is not supported in Wavelength Zones.

###### Boot mode topics

- [Requirements to launch an EC2 instance in UEFI boot mode](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/launch-instance-boot-mode.html)

- [Determine the boot mode parameter of an Amazon EC2 AMI](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-boot-mode.html)

- [Determine the supported boot modes of an EC2 instance type](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-type-boot-mode.html)

- [Determine the boot mode of an EC2 instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-boot-mode.html)

- [Determine the boot mode of the operating system for your EC2 instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/os-boot-mode.html)

- [Set the boot mode of an Amazon EC2 AMI](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/set-ami-boot-mode.html)

- [UEFI variables for Amazon EC2 instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/uefi-variables.html)

- [UEFI Secure Boot for Amazon EC2 instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/uefi-secure-boot.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Protect an AMI from deregistration

Requirements for UEFI boot mode
