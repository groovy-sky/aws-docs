# ENA queues

ENA queues are allocated to network interfaces with default static limits based on
the instance type and size. On supported instance types, you can dynamically
allocate these queues across Elastic Network Interfaces (ENIs). While the total
queue count per instance depends on its type and size, you can configure multiple
ENIs with ENA queues until you meet the maximum queue count for the ENI and the
instance.

Flexible ENA queue allocation optimizes resource distribution, enabling maximum
vCPU utilization. High network performance workloads typically require multiple ENA
queues. You can fine-tune network performance and packets per second (PPS) by
adjusting queue counts according to your specific workload needs. For example,
network-intensive applications may require more queues compared to CPU-intensive
applications.

###### Topics

- [Supported instances](#supported-instances)

- [Modify the number of queues](#modify)

## Supported instances

The following instances support dynamic allocation of multiple ENA queues.

### General purpose

Instance typeDefault ENA queues per interfaceMaximum ENA queues per interfaceMaximum ENA queues per instance**M6i**`m6i.large`226`m6i.xlarge`4416`m6i.2xlarge`8832`m6i.4xlarge`81664`m6i.8xlarge`83264`m6i.12xlarge`83264`m6i.16xlarge`832120`m6i.24xlarge`832120`m6i.32xlarge`832120**M6id**`m6id.large`226`m6id.xlarge`4416`m6id.2xlarge`8832`m6id.4xlarge`81664`m6id.8xlarge`83264`m6id.12xlarge`83264`m6id.16xlarge`832120`m6id.24xlarge`832120`m6id.32xlarge`832120**M6idn**`m6idn.large`226`m6idn.xlarge`4416`m6idn.2xlarge`8832`m6idn.4xlarge`81664`m6idn.8xlarge`1632128`m6idn.12xlarge`1632128`m6idn.16xlarge`1632240`m6idn.24xlarge`3232480`m6idn.32xlarge`3232512 \***M6in**`m6in.large`226`m6in.xlarge`4416`m6in.2xlarge`8832`m6in.4xlarge`81664`m6in.8xlarge`1632128`m6in.12xlarge`1632128`m6in.16xlarge`1632240`m6in.24xlarge`3232480`m6in.32xlarge`3232512 \***M8a**`m8a.medium`113`m8a.large`226`m8a.xlarge`4416`m8a.2xlarge`8832`m8a.4xlarge`81664`m8a.8xlarge`832128`m8a.12xlarge`1664192`m8a.16xlarge`1664256`m8a.24xlarge`16128384`m8a.48xlarge`32128768`m8a.metal-24xl`16128384`m8a.metal-48xl`32128768**M8azn**`m8azn.medium`113`m8azn.large`228`m8azn.xlarge`4416`m8azn.3xlarge`41648`m8azn.6xlarge`83296`m8azn.12xlarge`864192`m8azn.24xlarge`16128384`m8azn.metal-12xl`864192`m8azn.metal-24xl`16128384**M8gb**`m8gb.medium`112`m8gb.large`226`m8gb.xlarge`4416`m8gb.2xlarge`8832`m8gb.4xlarge`81664`m8gb.8xlarge`832128`m8gb.12xlarge`1664192`m8gb.16xlarge`1664256`m8gb.24xlarge`16128384`m8gb.48xlarge`32128768 \*`m8gb.metal-24xl`32128768`m8gb.metal-48xl`32128768 \***M8gn**`m8gn.medium`112`m8gn.large`226`m8gn.xlarge`4416`m8gn.2xlarge`8832`m8gn.4xlarge`81664`m8gn.8xlarge`832128`m8gn.12xlarge`1664192`m8gn.16xlarge`1664256`m8gn.24xlarge`16128384`m8gn.48xlarge`32128768 \*`m8gn.metal-24xl`32128768`m8gn.metal-48xl`32128768 \***M8i**`m8i.large`226`m8i.xlarge`4416`m8i.2xlarge`8832`m8i.4xlarge`81664`m8i.8xlarge`832128`m8i.12xlarge`1664192`m8i.16xlarge`1664256`m8i.24xlarge`16128384`m8i.32xlarge`16128512`m8i.48xlarge`32128768`m8i.96xlarge`321281536`m8i.metal-48xl`32128768`m8i.metal-96xl`321281536**M8id**`m8id.large`226`m8id.xlarge`4416`m8id.2xlarge`8832`m8id.4xlarge`81664`m8id.8xlarge`832128`m8id.12xlarge`1664192`m8id.16xlarge`1664256`m8id.24xlarge`16128384`m8id.32xlarge`16128512`m8id.48xlarge`32128768`m8id.96xlarge`321281536`m8id.metal-48xl`32128768`m8id.metal-96xl`321281536**M8i-flex**`m8i-flex.large`113`m8i-flex.xlarge`228`m8i-flex.2xlarge`4416`m8i-flex.4xlarge`4832`m8i-flex.8xlarge`41664`m8i-flex.12xlarge`83296`m8i-flex.16xlarge`832128

###### Note

\\* These instance types feature multiple network cards. Other instance types feature a
single network card. For more information, see [Network cards](using-eni.md#network-cards).

### Compute optimized

Instance typeDefault ENA queues per interfaceMaximum ENA queues per interfaceMaximum ENA queues per instance**C6i**`c6i.large`226`c6i.xlarge`4416`c6i.2xlarge`8832`c6i.4xlarge`81664`c6i.8xlarge`83264`c6i.12xlarge`83264`c6i.16xlarge`832120`c6i.24xlarge`832120`c6i.32xlarge`832120**C6id**`c6id.large`226`c6id.xlarge`4416`c6id.2xlarge`8832`c6id.4xlarge`81664`c6id.8xlarge`83264`c6id.12xlarge`83264`c6id.16xlarge`832120`c6id.24xlarge`832120`c6id.32xlarge`832120**C6in**`c6in.large`226`c6in.xlarge`4416`c6in.2xlarge`8832`c6in.4xlarge`81664`c6in.8xlarge`1632128`c6in.12xlarge`1632128`c6in.16xlarge`1632240`c6in.24xlarge`3232480`c6in.32xlarge`3232512 \***C8a**`c8a.medium`113`c8a.large`226`c8a.xlarge`4416`c8a.2xlarge`8832`c8a.4xlarge`81664`c8a.8xlarge`832128`c8a.12xlarge`1664192`c8a.16xlarge`1664256`c8a.24xlarge`16128384`c8a.48xlarge`32128768`c8a.metal-24xl`16128384`c8a.metal-48xl`32128768**C8gb**`c8gb.medium`112`c8gb.large`226`c8gb.xlarge`4416`c8gb.2xlarge`8832`c8gb.4xlarge`81664`c8gb.8xlarge`832128`c8gb.12xlarge`1664192`c8gb.16xlarge`1664256`c8gb.24xlarge`16128384`c8gb.48xlarge`32128768 \*`c8gb.metal-24xl`32128768`c8gb.metal-48xl`32128768 \***C8gn**`c8gn.medium`112`c8gn.large`226`c8gn.xlarge`4416`c8gn.2xlarge`8832`c8gn.4xlarge`81664`c8gn.8xlarge`832128`c8gn.12xlarge`1664192`c8gn.16xlarge`1664256`c8gn.24xlarge`16128384`c8gn.48xlarge`32128768 \*`c8gn.metal-24xl`32128768`c8gn.metal-48xl`32128768 \***C8i**`c8i.large`226`c8i.xlarge`4416`c8i.2xlarge`8832`c8i.4xlarge`81664`c8i.8xlarge`832128`c8i.12xlarge`1664192`c8i.16xlarge`1664256`c8i.24xlarge`16128384`c8i.32xlarge`16128512`c8i.48xlarge`32128768`c8i.96xlarge`321281536`c8i.metal-48xl`32128768`c8i.metal-96xl`321281536**C8id**`c8id.large`226`c8id.xlarge`4416`c8id.2xlarge`8832`c8id.4xlarge`81664`c8id.8xlarge`832128`c8id.12xlarge`1664192`c8id.16xlarge`1664256`c8id.24xlarge`16128384`c8id.32xlarge`16128512`c8id.48xlarge`32128768`c8id.96xlarge`321281536`c8id.metal-48xl`32128768`c8id.metal-96xl`321281536**C8i-flex**`c8i-flex.large`113`c8i-flex.xlarge`228`c8i-flex.2xlarge`4416`c8i-flex.4xlarge`4832`c8i-flex.8xlarge`41664`c8i-flex.12xlarge`83296`c8i-flex.16xlarge`832128

###### Note

\\* These instance types feature multiple network cards. Other instance types feature a
single network card. For more information, see [Network cards](using-eni.md#network-cards).

### Memory optimized

Instance typeDefault ENA queues per interfaceMaximum ENA queues per interfaceMaximum ENA queues per instance**R6i**`r6i.large`226`r6i.xlarge`4416`r6i.2xlarge`8832`r6i.4xlarge`81664`r6i.8xlarge`83264`r6i.12xlarge`83264`r6i.16xlarge`832120`r6i.24xlarge`832120`r6i.32xlarge`832120**R6id**`r6id.large`226`r6id.xlarge`4416`r6id.2xlarge`8832`r6id.4xlarge`81664`r6id.8xlarge`83264`r6id.12xlarge`83264`r6id.16xlarge`832120`r6id.24xlarge`832120`r6id.32xlarge`832120**R6idn**`r6idn.large`226`r6idn.xlarge`4416`r6idn.2xlarge`8832`r6idn.4xlarge`81664`r6idn.8xlarge`1632128`r6idn.12xlarge`1632128`r6idn.16xlarge`1632240`r6idn.24xlarge`3232480`r6idn.32xlarge`3232512 \***R6in**`r6in.large`226`r6in.xlarge`4416`r6in.2xlarge`8832`r6in.4xlarge`81664`r6in.8xlarge`1632128`r6in.12xlarge`1632128`r6in.16xlarge`1632240`r6in.24xlarge`3232480`r6in.32xlarge`3232512 \***R8a**`r8a.medium`113`r8a.large`226`r8a.xlarge`4416`r8a.2xlarge`8832`r8a.4xlarge`81664`r8a.8xlarge`832128`r8a.12xlarge`1664192`r8a.16xlarge`1664256`r8a.24xlarge`16128384`r8a.48xlarge`32128768`r8a.metal-24xl`16128384`r8a.metal-48xl`32128768**R8gb**`r8gb.medium`112`r8gb.large`226`r8gb.xlarge`4416`r8gb.2xlarge`8832`r8gb.4xlarge`81664`r8gb.8xlarge`832128`r8gb.12xlarge`1664192`r8gb.16xlarge`1664256`r8gb.24xlarge`16128384`r8gb.48xlarge`32128768 \*`r8gb.metal-24xl`32128768`r8gb.metal-48xl`32128768 \***R8gn**`r8gn.medium`112`r8gn.large`226`r8gn.xlarge`4416`r8gn.2xlarge`8832`r8gn.4xlarge`81664`r8gn.8xlarge`832128`r8gn.12xlarge`1664192`r8gn.16xlarge`1664256`r8gn.24xlarge`16128384`r8gn.48xlarge`32128768 \*`r8gn.metal-24xl`32128768`r8gn.metal-48xl`32128768 \***R8i**`r8i.large`226`r8i.xlarge`4416`r8i.2xlarge`8832`r8i.4xlarge`81664`r8i.8xlarge`832128`r8i.12xlarge`1664192`r8i.16xlarge`1664256`r8i.24xlarge`16128384`r8i.32xlarge`16128512`r8i.48xlarge`32128768`r8i.96xlarge`321281536`r8i.metal-48xl`32128768`r8i.metal-96xl`321281536**R8id**`r8id.large`226`r8id.xlarge`4416`r8id.2xlarge`8832`r8id.4xlarge`81664`r8id.8xlarge`832128`r8id.12xlarge`1664192`r8id.16xlarge`1664256`r8id.24xlarge`16128384`r8id.32xlarge`16128512`r8id.48xlarge`32128768`r8id.96xlarge`321281536`r8id.metal-48xl`32128768`r8id.metal-96xl`321281536**R8i-flex**`r8i-flex.large`113`r8i-flex.xlarge`228`r8i-flex.2xlarge`4416`r8i-flex.4xlarge`4832`r8i-flex.8xlarge`41664`r8i-flex.12xlarge`83296`r8i-flex.16xlarge`832128**X8aedz**`x8aedz.large`228`x8aedz.xlarge`4416`x8aedz.3xlarge`41648`x8aedz.6xlarge`83296`x8aedz.12xlarge`864192`x8aedz.24xlarge`16128384`x8aedz.metal-12xl`864192`x8aedz.metal-24xl`16128384**X8i**`x8i.large`226`x8i.xlarge`4416`x8i.2xlarge`8832`x8i.4xlarge`81664`x8i.8xlarge`832128`x8i.12xlarge`1664192`x8i.16xlarge`1664256`x8i.24xlarge`16128384`x8i.32xlarge`16128512`x8i.48xlarge`32128768`x8i.64xlarge`321281024`x8i.96xlarge`321281536`x8i.metal-48xl`32128768`x8i.metal-96xl`321281536

###### Note

\\* These instance types feature multiple network cards. Other instance types feature a
single network card. For more information, see [Network cards](using-eni.md#network-cards).

## Modify the number of queues

You can modify the number of ENA queues using AWS Management Console or AWS CLI. In the
AWS Management Console, the ENA queues configuration is available under each **Network interface** setting.

To modify the number of ENA queues using the AWS CLI, use either one of the
following commands. Before modifying the queue count, use the following command
to check your current queue count.

```json

aws ec2 describe-instances --instance-id i-1234567890abcdef0
```

###### Note

- Your instance must be stopped before modifying the number of ENA
queues.

- The value for ENA queues must be a power of 2, such as, 1, 2, 4,
8, 16, 32, etc.

- The number of queues allocated to any single ENI cannot exceed the
number of vCPUs available on your instance.

`attach-network-interface`

In the following example, 32 ENA queues are configured on an ENI.

```json

aws ec2 attach-network-interface \
  --network-interface-id eni-001aa1bb223cdd4e4 \
  --instance-id i-1234567890abcdef0 \
  --device-index 1 \
  --ena-queue-count 32
```

`run-instances`

In the following example, 2 ENA queues each are configured on 3 ENIs.

```json

aws ec2 run-instances \
  --image-id ami-12ab3c30 \
  --instance-type c6i.large \
  --min-count 1 \
  --max-count 1 \
  --network-interfaces \
    "[{\"DeviceIndex\":0,\"SubnetId\":\"subnet-123456789012a345a\",\"EnaQueueCount\":2},
      {\"DeviceIndex\":1,\"SubnetId\":\"subnet-123456789012a345a\",\"EnaQueueCount\":2},
      {\"DeviceIndex\":2,\"SubnetId\":\"subnet-123456789012a345a\",\"EnaQueueCount\":2}]"
```

`modify-network-interface-attribute`

In the following example, 32 ENA queues are configured on an ENI.

```json

aws ec2 modify-network-interface-attribute \
--network-interface-id eni-1234567890abcdef0 \
--attachment AttachmentId=eni-attach-12345678,EnaQueueCount=32
```

In the following example, the ENA count is reset to the default value.

```json

aws ec2 modify-network-interface-attribute \
--network-interface-id eni-1234567890abcdef0 \
--attachment AttachmentId=eni-attach-12345678,DefaultEnaQueueCount=true
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Enable ENA for an instance

Troubleshoot ENA on Linux
