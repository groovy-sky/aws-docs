# Methods for terminating an instance

###### Warning

**Terminating an instance is permanent and irreversible.**

After you terminate an instance, you can no longer connect to it, and it can't be recovered.
All attached Amazon EBS volumes that are configured to be deleted on termination are also permanently
deleted and can't be recovered. All data stored on instance store volumes is permanently lost.
For more information, see [How instance termination works](how-ec2-instance-termination-works.md).

Before you terminate an instance, ensure that you have backed up all data that you need to
retain after the termination to persistent storage.

There are four ways to perform a user-initiated instance termination: default terminate,
terminate with skip OS shutdown, force terminate, and force terminate with skip OS
shutdown. The following table compares the key differences between the termination
methods:

###### Note

You can't terminate an instance if termination protection is turned
on. For more information, see [Change instance termination protection](using-changingdisableapitermination.md).

Termination methodKey purposeUse caseCLI commandDefault terminateNormal instance shutdown with attempted graceful OS shutdown.Typical instance termination.

```nohighlight

aws ec2 terminate-instances \
--instance-id i-1234567890abcdef0
```

Terminate with skip OS shutdownBypasses the graceful OS shutdown when terminating an instance.When bypassing graceful OS shutdown is required.

```nohighlight

aws ec2 terminate-instances \
--instance-id i-1234567890abcdef0 \
--skip-os-shutdown
```

Force terminateHandles stuck instances. Attempts a default termination first; if
instance fails to terminate, then forcibly terminates the
instance.When instance is stuck in `shutting-down` state.

```nohighlight

aws ec2 terminate-instances \
--instance-id i-1234567890abcdef0 \
--force
```

Force terminate with skip OS shutdownForce terminates and bypasses the graceful OS shutdown when terminating an
instance.When force termination and bypassing graceful OS shutdown is
required.

```nohighlight

aws ec2 terminate-instances \
--instance-id i-1234567890abcdef0 \
--force \
--skip-os-shutdown
```

For instructions on how to use each method, see the following:

- [Terminate an instance with a graceful OS shutdown](terminating-instances.md#terminating-instances-console)

- [Terminate an instance and bypass the graceful OS shutdown](terminating-instances.md#terminating-instances-bypass-graceful-os-shutdown)

- [Force terminate an instance](troubleshootinginstancesshuttingdown.md#force-terminate-ec2-instance)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

How it
works

Change termination
protection
