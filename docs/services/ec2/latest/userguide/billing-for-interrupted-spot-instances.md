# Billing for interrupted Spot Instances

When a Spot Instance is interrupted, you're charged for instance and EBS volume usage, and you might
incur other charges, as follows.

## Instance usage

Who interrupts the Spot InstanceOperating systemInterrupted in the first hourInterrupted in any hour after the first hour

If **you** stop or terminate the Spot Instance

Windows and Linux (excluding SUSE)Charged for the seconds usedCharged for the seconds usedSUSECharged for the full hour even if you used a partial hourCharged for the full hours used, and charged a full hour for the
interrupted partial hour

If the **Amazon EC2** interrupts the Spot Instance

Windows and Linux (excluding SUSE)No chargeCharged for the seconds usedSUSENo charge

Charged for the full hours used, but no charge for the
interrupted partial hour

## EBS volume usage

While an interrupted Spot Instance is stopped, you are charged only for the EBS volumes, which are
preserved.

With EC2 Fleet and Spot Fleet, if you have many stopped instances, you can exceed the
limit on the number of EBS volumes for your account.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Determine whether Amazon EC2 terminated a Spot Instance

Rebalance recommendations
