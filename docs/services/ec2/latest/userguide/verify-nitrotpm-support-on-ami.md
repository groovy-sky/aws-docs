# Verify that an AMI is enabled for NitroTPM

To enable NitroTPM for an instance, you must launch the instance using an AMI
with NitroTPM enabled. You can describe an image to verify that it is enabled for
NitroTPM. If you are the AMI owner, you can describe the `tpmSupport`
image attribute.

The Amazon EC2 console does not display `TpmSupport`.

AWS CLI

###### To verify that NitroTPM is enabled

Use the [describe-images](../../../cli/latest/reference/ec2/describe-images.md)
command.

```nohighlight

aws ec2 describe-images \
    --image-ids ami-0abcdef1234567890 \
    --query Images[*].TpmSupport
```

If NitroTPM is enabled for the AMI, the output is as follows. If TPM is not enabled,
the output is empty.

```JSON

[
    "v2.0"
]
```

Alternatively, if you are the AMI owner, you can use the [describe-image-attribute](../../../cli/latest/reference/ec2/describe-image-attribute.md) command with the `tpmSupport` attribute.

```nohighlight

aws ec2 describe-image-attribute \
    --image-id ami-0abcdef1234567890 \
    --attribute tpmSupport
```

The following is example output.

```JSON

{
    "ImageId": "ami-0abcdef1234567890",
    "TpmSupport": {
        "Value": "v2.0"
    }
}
```

###### To find AMIs with NitroTPM enabled

The following example lists the IDs of the AMIs that you own with
NitroTPM enabled.

```nohighlight

aws ec2 describe-images \
    --owners self \
    --filters Name=tpm-support,Values=v2.0 \
    --query Images[].ImageId
```

PowerShell

###### To verify that NitroTPM is enabled

Use the [Get-EC2Image](../../../powershell/latest/reference/items/get-ec2image.md)
cmdlet.

```powershell

Get-EC2Image `
    -ImageId ami-0abcdef1234567890 | Select TpmSupport
```

If NitroTPM is enabled for the AMI, the output is as follows. If TPM is not enabled,
the output is empty.

```nohighlight

TpmSupport
----------
v2.0
```

Alternatively, if you are the AMI owner, you can use the [Get-EC2ImageAttribute](../../../powershell/latest/reference/items/get-ec2imageattribute.md) cmdlet with the `tpmSupport` attribute.

```powershell

Get-EC2ImageAttribute `
    -ImageId ami-0abcdef1234567890 `
    -Attribute tpmSupport
```

###### To find AMIs with NitroTPM enabled

The following example lists the IDs of the AMIs that you own with
NitroTPM enabled.

```powershell

Get-EC2Image `
    -Owner self `
    -Filter @{Name="tpm-support; Values="v2.0"} | Select ImageId
```

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Enable a Linux AMI for NitroTPM

Enable or stop using NitroTPM
