# Elastic Fabric Adapter release notes

The following table describes the version history and changelog for the Elastic Fabric Adapter software.

VersionChangesRelease date1.47.0

- **Upgrade to [libfabric 2.4.0amzn1.0](https://github.com/aws/libfabric/releases/tag/v2.4.0amzn1.0)**

- Refactor peer management: move endpoint→peer hashmap to AV level

- Add packet entry generation counter to prevent ABA problem

- Implement blocking CQ read support (fi\_cq\_sread) with wait objects

- Fix double-free in eager/mulreq/longread RTM packet processing

- Fix MR registration IBV access flags based on device RDMA capabilities

- Disable FI\_OPT\_EFA\_SENDRECV\_IN\_ORDER\_ALIGNED\_128\_BYTES (returns -FI\_EOPNOTSUPP)

- Optimize packet entry to fit in two x86 cache lines (128 bytes)

- Always succeed efa\_rdm\_srx\_start and slide recv window on RTM/RTA errors

- Unlink RX packet entries before releasing during endpoint cleanup

- **Upgrade to [libnccl-ofi 1.18.0](https://github.com/aws/aws-ofi-nccl/releases/tag/v1.18.0)**

- P6-B300 support: added custom tuner decisions for P6-B300

- Improved performance of PAT on P6-B200 by reducing channel count for smaller message sizes

- Changed protocol defaults, default to RDMA protocol on Trn1 and default to SENDRECV protocol on g7e.8xlarge

- Dynamic platform selection: added feature to enable AWS optimizations at runtime based on presence of AWS NICs. This allows a single plugin binary to be used for both AWS and non-AWS platforms. AWS optimizations can still be disabled at compile-time.

- Fixed support for non-FI\_MR\_VIRT\_ADDR providers in RDMA protocol

- Improved NIC PCIe link speed and width reporting to NCCL

- Redesigned threading model to support multi-threaded applications without requiring a separate Libfabric domain for each thread.

- Fixed support for FI\_MR\_ENDPOINT providers (supports SENDRECV protocol only) by cleaning up resources in correct order

- **Upgrade to [rdma-core 61.0](https://github.com/linux-rdma/rdma-core/releases/tag/v61.0)**

- Use single threaded CQ if Thread Domain was provided

- Verify QP number on CQE processes

- **Upgrade to [efa driver 3.0.0](https://github.com/amzn/amzn-drivers/releases/tag/efa_linux_3.0.0)**

- Improve admin error handling

- Check for QP number correctness on completion poll

- Remove unconditional backport of best page size finding

- Print errno strings for error pointers

- Simplify the code in CQ creation flows

- Backport upstream changes in CQ with umem creation flow

- Adjust DKMS configuration for new DKMS versions

- **Upgrade to [Open MPI 5.0.9amzn1](https://github.com/open-mpi/ompi/tree/v5.0.9amzn1)**

- Bugfix: Set domain threading level based on MPI thread support

- Bugfix: Request FI\_COMPLETION flag to ensure completions are generated for all data transfer events

- Share domains between BTL and MTL to reduce the total number of domains created to fix resource exhaustion on systems with high core counts

January 29, 20261.46.0

- **Add Support for Debian 13**

- **Upgrade to [efa-nv-peermem 1.2.3](https://github.com/amzn/amzn-drivers/releases/tag/efa_nv_peermem_linux_1.2.3)**

- Fix build errors and warnings with new kernels and DKMS

- **Upgrade to [libfabric 2.3.1amzn4.0](https://github.com/aws/libfabric/releases/tag/v2.3.1amzn4.0)**

- Enable data path direct for efa-rdm protocol path

- Enable ROCr HMEM support

- New FI\_OPT\_EFA\_USE\_UNSOLICITED\_WRITE\_RECV option to disable unsolicited write recv

- Optimize connection establishment time

- **Add libnccl-ofi on RHEL 9, Rocky Linux 9, and Debian 13**

December 12, 20251.45.1

- **Upgrade to [libfabric 2.3.1amzn3.0](https://github.com/aws/libfabric/releases/tag/v2.3.1amzn3.0)**

- Fix bug that truncated cq\_data to 2 bytes while provider advertised support for 4 bytes

November 26, 20251.45.0

- **Upgrade to [rdma-core 60.amzn0](https://github.com/linux-rdma/rdma-core/releases/tag/v60.0)**

- Extend DV query CQ to return doorbell

- **Upgrade to [libfabric 2.3.1amzn2.0](https://github.com/aws/libfabric/releases/tag/v2.3.1amzn2.0)**

- Support FI\_RX\_CQ\_DATA mode for efa-direct

- Optimize WQE post in data path direct path

- Fix the race condition in cntr progress

- Improve the logging for warning messages

- Handles RMA completion correctly for removed peers

- Evict AH entries from implicit AV when AH limit is reached

- **Upgrade to [libnccl-ofi 1.17.2](https://github.com/aws/aws-ofi-nccl/releases/tag/v1.17.2)**

- Fixed a crash with NCCL v2.28.x when Libfabric initialization failed

- Added support for g7e instance family

- Fixed an issue where NCCL could erroneously attempt to use a GPUDirect RDMA path on platforms that support DMA-BUF

- Fixed shutdown ordering issue on NICs that require per-endpoint memory registration (Cray Slingshot)

November 17, 20251.44.0

- **Upgrade to [rdma-core 59.amzn0](https://github.com/linux-rdma/rdma-core/releases/tag/v59.0)**

- Add support to bind QPs and CQs to thread domains

- Add support to allocate a parent domain

- Add support to allocate thread domains

- Add direct verbs query QP and CQ

- Add option to create CQ with external memory\`

- Fix CQ poll after QP destroy

- **Upgrade to [libfabric 2.3.1amzn1.0](https://github.com/aws/libfabric/releases/tag/v2.3.1amzn1.0)**

- **Upgrade to [Open MPI 5.0.8amzn1](https://docs.open-mpi.org/en/v5.0.8amzn1/index.html)**

- Functionality enhancements

- Bug fixes

- **Upgrade to [libnccl-ofi 1.17.1](https://github.com/aws/aws-ofi-nccl/releases/tag/v1.17.1)**

- Supports NCCL v2.28.3-1 while maintaining backward compatibility with NCCL v2.17.1 and later.

- Building with platform-aws requires Libfabric v1.22.0amzn4.0 or greater. Tested with versions up
to Libfabric v2.3.1amzn1.0.

- Fixes to support compatibility across CUDA12 and CUDA13.

October 29, 20251.43.3

- **Upgrade to [efa driver 2.17.3](https://github.com/amzn/amzn-drivers/releases/tag/efa_linux_2.17.3)**

- Support P2P with NVIDIA 580 drivers

- Adjust MR registration for 6.17 and on kernels

October 01, 20251.43.2

- **Upgrade to [libnccl-ofi 1.16.3](https://github.com/aws/aws-ofi-nccl/releases/tag/v1.16.3)**

- Supports NCCL v2.27.7-1 while maintaining backward compatibility with
NCCL v2.17.1 and later.

- Building with platform-aws requires Libfabric v1.22.0amzn4.0 or greater.
Tested with versions up to Libfabric 2.1.0amzn5.0.

- Enable domain-per-thread by default on all Amazon EC2 instance types for improved performance for some applications where NCCL creates multiple proxy threads

August 15, 20251.43.1

- **Upgrade to [libnccl-ofi 1.16.2](https://github.com/aws/aws-ofi-nccl/releases/tag/v1.16.2)**

- Supports NCCL v2.27.6-1 while maintaining backward compatibility with
NCCL v2.17.1 and later.

- Building with platform-aws requires Libfabric v1.22.0amzn4.0 or greater.
Tested with versions up to Libfabric 2.1.0amzn4.0.

- Added a new platform configuration to support using the OFI NCCL plugin on the Amazon EC2 p5.4xlarge instance type

- **Upgrade to [libfabric \**
**2.1.0amzn5.0](https://github.com/aws/libfabric/releases/tag/v2.1.0amzn5.0)**

July 31, 20251.43.0

- **Upgrade to [libnccl-ofi 1.16.1](https://github.com/aws/aws-ofi-nccl/releases/tag/v1.16.1)**

- Supports NCCL v2.27.6-1 while maintaining backward compatibility with NCCL v2.17.1 and later.

- Building with platform-aws requires Libfabric v1.22.0amzn4.0 or greater.
Tested with versions up to Libfabric 2.1.0amzn3.

- Update the PCI link speed format reported in the topology file to match kernel 5.7+

- Added SKIP\_NICS\_WITHOUT\_ACCEL\_AT\_SAME\_PCI\_LEVEL to skip libfabric nics that do not have an accelerator at the same pci level

- **Upgrade to [efa driver 2.17.2](https://github.com/amzn/amzn-drivers/releases/tag/efa_linux_2.17.2)**

- **Upgrade to [efa-nv-peermem 1.2.2](https://github.com/amzn/amzn-drivers/releases/tag/efa_nv_peermem_linux_1.2.2)**

- **Upgrade to [libfabric \**
**2.1.0amzn4.0](https://github.com/aws/libfabric/releases/tag/v2.1.0amzn4.0)**

- **Upgrade to [rdma-core 58.amzn0](https://github.com/linux-rdma/rdma-core/releases/tag/v58.0)**

- Fix work request index double use

- Add WQE length to post\_send tracepoint

- **Enable optimizations for Graviton platforms**

- **Deprecate support for Ubuntu 20.04**

July 25, 20251.42.0

- **Upgrade to [efa driver 2.15.3](https://github.com/amzn/amzn-drivers/releases/tag/efa_linux_2.15.3)**

- **Upgrade to [efa-nv-peermem 1.2.1](https://github.com/amzn/amzn-drivers/releases/tag/efa_nv_peermem_linux_1.2.1)**

- **Upgrade to [rdma-core 57.amzn1](https://github.com/linux-rdma/rdma-core/releases/tag/v57.1)**

- Fix work request index double use

- **Upgrade to [libfabric 2.1.0amzn3.0](https://github.com/aws/libfabric/releases/tag/v2.1.0amzn3.0)**

- **Upgrade to [libnccl-ofi 1.15.0](https://github.com/aws/aws-ofi-nccl/releases/tag/v1.15.0)**

- Supports NCCL 2.26.6-1 while maintaining backward compatibility with (NCCL v2.17.1 and later.

- Building with platform-aws requires Libfabric v1.22.0amzn4.0 or greater. Tested with versions up to Libfabric 2.1.0amzn3.

- Build system and platform support

- Added Amazon EC2 P6-B200 platform support

- Changed default plugin library name to libnccl-net-ofi.so, and by default create symlink from libnccl-net-ofi.so to libnccl-net.so to maintain backward compatibility. This allows users to set NCCL\_NET\_PLUGIN=ofi to force NCCL to use the OFI plugin for communication. Specifying --disable-nccl-net-symlink to configure will skip the symlink, allowing multiple plugins to be installed in the same container.

- Tuning and performance improvements

- Added tuner support on P6-B200 for AllReduce, AllGather, and ReduceScatter regions for 0x0 and 0x7 bitmask

- Updated default latency for P5en and P6-B200 platforms based on empirical results and analysis

- Update to use NCCL v10 API with trafficClass parameter support for future traffic prioritization

- Migrated plugin code base from C to C++

- Added support for jobs where the number of NICs per GPU is different across systems. See the OFI\_NCCL\_FORCE\_NUM\_RAILS runtime environment variable documentation for more information.

June 6, 20251.41.0

- Upgrade to rdma-core 57.amzn0

May 16, 20251.40.0

- Upgrade to libfabric 2.1.0amzn2.0

- Upgrade to rdma-core 56.0

- Upgrade to efa-config 1.18.

- Deprecate support for Debian 10

- Upgrade to libnccl-ofi 1.14.2

- Enable CUDA support in Libfabric on ARM for Ubuntu and Amazon Linux 2023

- Add libnccl-ofi on ARM for Ubuntu and Amazon Linux 2023

May 2, 20251.39.0

- Upgrade to libfabric 2.1.0

- Upgrade to Open MPI 5.0.6

- Upgrade to libnccl-ofi 1.14.1

- Upgrade to efa driver 2.15.0

April 16, 20251.38.1

- Update to `Libfabric 1.22.0amzn5.0`

March 3, 20251.38.0

- Add `libnccl-ofi 1.13.2-1`

January 8, 20251.37.0

- Upgrade to `libfabric 1.22.0amzn4.0`

- Upgrade to `rdma-core 54.amzn0`

November 18, 20241.36.0

- Add support for Debian 12

- Distribute `efa_test.sh` as a utility script at `/opt/amazon/efa/bin`

- Upgrade to efa driver 2.13.0

- Upgrade to `libfabric 1.22.0amzn3.0`

- Upgrade to `rdma-core 54.0`

- Upgrade to open mpi 4.1.7

November 7, 20241.35.0

- Upgrade to Open MPI 5.0.5

- Upgrade to PRRTE 3.0.6 and rename the RPM package to prrte-aws

- Rename the OpenPMIx RPM package to pmix-aws

- Configure build options for dpkg and RPM builds

- Upgrade to `Libfabric 1.22.0amzn2.0`

- Upgrade to efa driver 2.12.1

October 14, 20241.34.0

- Drop support for CentOS 7 and RHEL 7

- Ingest `Libfabric 1.22.0amzn1.0`

- Update efa-config package to 1.17

August 6, 20241.33.0

- Upgrade to efa driver 2.10.0

- Upgrade to `rdma-core 52.0`

- Upgrade to PRRTE 3.0.5

- Upgrade to Open MPI 5.0.3

- Add support for Ubuntu 24.04 LTS

- Remove OpenPMIx, PRRTE and Open MPI HTML documentation. Users should refer to the official website.

- Drop support for openSUSE 15.4 or older

June 20, 20241.32.0

- Upgrade to efa driver 2.8.0

- Upgrade to efa-nv-peermem driver 1.1.1

- Upgrade efa-config package to 1.16

- Upgrade efa-profile package to 1.7

- Upgrade to `rdma-core 50.0`

- Upgrade to `libfabric 1.21.0amzn1.0`

- Enhance `efa_test.sh` with `fi_pingpong` port selection

- Install newer rdma-core from system when required

April 18, 20241.31.0

- Upgrade to OpenPMIx 4.2.8

- Upgrade to PRRTE 3.0.3

- Upgrade to Open MPI 5.0.2, and convert MCA components to DSO.

- Upgrade to `Libfabric 1.20.1amzn1.0`

March 7, 20241.30.0

- Include OpenPMIx 4.2.7, installed at `/opt/amazon/pmix` by default

- Include PRRTE 3.0.2, installed at `/opt/amazon/prrte` by default

- Include Open MPI 5.0.0, installed at `/opt/amazon/openmpi5` by default

- Update efa-profile to 1.6

December 20231.29.1

- Ingest `libfabric 1.19.0amzn4.0`

December 20231.29.0

- Ingest efa kernel driver 2.6.0

- Ingest `libfabric 1.19.0amzn3.0`

November 20231.28.0

- Upgrade Open MPI to 4.1.6

- Ingest `libfabric 1.19.0amzn2.0`

October 20231.27.0

- Add `Libfabric 1.19.0amzn1.0` to installer

- Enable GCC builtin atomics for Open MPI 4

September 20231.26.1

- `efa_test.sh`: fix infinite retry bug

September 20231.26.0

- Add support for RHEL 9

- Fix debug build on multiple distributions

- Ingest `libfabric 1.18.2amzn1.0`

- `efa_test.sh`: Add retry for `fi_pingpong`

September 20231.25.1

- Ingest `libfabric 1.18.1amzn1.0`

- Update timeouts in EFA installer scripts

September 20231.25.0

- Add support for Amazon Linux 2023

- Bugfix in post-installation test

July 20231.24.1

- Upgrade libfabric to 1.18.1

- Upgrade efa driver to 2.5.0

July 20231.24.0

- Ingest `rdma-core 46.0`

- Ingest efa driver 2.4.1

- Support Debian 11

June 20231.23.1

- Ingest `libfabric 1.18.0amzn2.0`

June 20231.23.0

- Add support for Debian 10

- Drop support for Ubuntu 18.04 LTS

- Upgrade efa-config package to 1.14

- Ingest `libfabric 1.18.0amzn1.0`

May 20231.22.1

- Upgrade libfabric to 1.17.1

March 20231.22.0

- Upgrade Open MPI to 4.1.5

- Upgrade libfabric to 1.17.0

- Upgrade efa-config package to 1.13

February 20231.21.0

- Add support for Rocky Linux 9 OS

- Ingest efa driver 2.1.1

- Ingest `libfabric 1.16.1amzn3.0`

- Upgrade efa-config package to 1.12

December 20221.20.0

- Add support for Rocky Linux 8 OS.

- Ingest efa driver 2.1.0.

- Ingest `rdma-core 43.0`.

- Ingest `libfabric 1.16.1amzn1.0`

November 20221.19.0

- Ingest `libfabric 1.16.0`

- Build Open MPI with `--enable-orterun-prefix-by-default`

October 20221.18.0

- Add support for Ubuntu22.04

August 20221.17.3

- Update libfabric to `1.16.0~amzn4.0`. The `~` indicates it is a pre-release version of libfabric 1.16.0.

- Extend post-installation pingpong test timeout to 20 seconds

August 20221.17.2

- Update libfabric to `1.16.0~amzn3.0`. The `~` indicates it is a pre-release version of libfabric 1.16.0.

July 20221.17.1

- Update libfabric to `1.16.0~amzn2.0`. The `~` indicates it is a pre-release version of libfabric 1.16.0.

- Disable the experimental net provider when building libfabric

July 20221.17.0

- Update `rdma-core` to v41.0

- Update Open MPI to 4.1.4

- Update libfabric to `1.16.0~amzn1.0`. The `~` indicates it is a pre-release version of libfabric 1.16.0.

July 20221.16.0

- Update libfabric to `1.15.1amzn1.0`, contains neuron library name change

- Upgrade efa-config to 1.10

- Exclude opx and rxd providers in the libfabric build

June 20221.15.2

- Update libfabric to 1.14.1

May 20221.15.1

- Update libfabric to `1.14.0amzn1.0`

March 20221.15.0

- Fix a bug that cause installation fail on Open SuSE 15.3

- Drop support of Open SuSE 15.2 (as Open SuSE 15.2 reached end of life)

- Drop Support of CentOS 8 (as CentOS 8 reached end of life)

- Update libfabric to 1.14.0

- Update efa kernel driver to 1.16.0

- Update `rdma-core` to v39.0

- Update Open MPI to version 4.1.2.

Feburary 20221.14.1

- Update libfabric to `1.13.2amzn1.0`.

October 20211.14.0

- Ingest efa kernel driver 1.14.2.

- Make `-g, --enable-gdr` in `efa_installer.sh` as a no-op option as the latest efa kernel driver enables GDR support by default.

- Ingest `rdma-core v37.0`.

- Ingest `libfabric 1.13.2`.

- Add packages list and compare RPM/DEB to list during installation to prevent unknown package installations.

- Add sleep in installer script to wait for udev rule to apply after EFA driver reload.

October 20211.13.0

- Update `rdma-core` to v35.0.

- Update `libfabric` to v1.13.0amzn1.0.

- Add EFA support for CentOS/RHEL 8 on Gravition2 platform.

- Add version comparison logic in installer script to skip the local package installation when there is one installed in system with higher version.

August 20211.12.3

- Update EFA kernel module to 1.13.0.

- Update efa-config package to version 1.9. Improve the calculation of huge page reservation to handle large defaulted huge page size.

July 20211.12.2

- Update EFA kernel module to 1.12.3.

- Build Open MPI debian packages with `--with-libevent=external` and `--with-hwloc=external`.

- Bump Open MPI rpm build ID to 2 to fix backward compatibility issue of HWLOC on CentOS 8.

- Remove the installation of kernel-devel and kernel-source packages on SLES15SP2 and openSUSE 15.2.

June 20211.12.1

- Update Libfabric to version 1.11.2amzon1.1.

- Update EFA kernel module to version 1.12.1.

May 20211.12.0

- Update Open MPI to version 4.1.1.

- Update Libfabric to version 1.11.2amzn1.0.

- Build rdma-core for Amazon Linux 2 using the same packaging configuration as the AL2 rdma-core.

- Do not force `-Wl,--enable-new-dtags` when building Open MPI RPMs.

- Build Open MPI with system libraries for hwloc and libevent.

- Update EFA kernel module to version 1.12.0

- Update efa-config package to version 1.8. Improve the calculation of huge page reservation for long-lived instances.

- Update efa-profile package to version 1.5. Remove the open mpi collective tuning file that worked as a workaround to fix Open MPI 4.1.0 hang on P4d.

- Update `rdma-core` to v32.1.

- Drop support for Amazon Linux 1 and Ubuntu 16.04.

May 20211.11.2

- Fix Open MPI hang when using Open MPI on P4d by changing the default algorithm used to implement MPI\_BARRIER via a configuration file.

- Disable use of builtin atomics in Open MPI on ARM via `--disable-builtin-atomics` to work around compiler issue.

February 20211.11.1

- Update Open MPI to version 4.1.0.

- Update efa-config package to version 1.7. Improve calculation of huge page reservation count.

- Update efa-profile package to version 1.3. Removes unneeded collectives decision file now that Open MPI 4.1.0 is used.

December 20201.11.0

- Add support for Gravition2 platform.

- Update `rdma-core` to version 31.2amzn.

- Update Libfabric to version 1.11.1amzn1.0.

- Update efa-config to version 1.6.

- Update efa-profile to version 1.2.

December 20201.10.1

- Add support for CentOS / RHEL 8.

- Add support for Ubuntu 20.04.

- Add support for SUSE Linux Enterprise 15.

November 20201.10.0

- Add GPUDirect RDMA support for P4d platform. Use `--enable-gdr` installer option to instal GDR-aware kernel module and userspace.

- Update EFA kernel module to version 1.10.2.

- Update `rdma-core` to version 31.amzn0.

- Update Libfabric to version 1.11.1.

- Update Open MPI to version 4.0.5.

- Update efa-config to version 1.5.

- Update efa-profile to version 1.1. Includes improved Open MPI collectives decision file.

October 20201.9.5

- Update efa-config to version 1.4. Fixes bug in Open MPI collective decision file.

September 20201.9.4

- Update Open MPI to version 4.0.3.

- Update Libfabric to version 1.10.1amazon1.1.

- Update `rdma-core` to version 28.amzn0.

July 20201.9.3

- Update EFA kernel module to version 1.6.0.

- Update `rdma-core` to version 28.amzn0.

- Update Libfabric to version 1.10.1amzn1.1.

- Update efa-config to version 1.3. Adds collectives tuning file for Open MPI.

- Skip dkms installation if it is already installed.

- Fix `--skip-kmod` installation mode to actually work.

June 20201.8.4

- Move configuration files into efa-config and efa-profile packages so that they are tracked by the operating system package manager.

- Update Open MPI to version 4.0.3.

April 20201.8.3

- Update EFA kernel module to version 1.5.1.

- Distributed DKMS on some platforms rather than relying on EPEL repositories for added installation reliability.

- On RHEL 7, install RPMs built on CentOS 7 instead of RPMs built on Amazon Linux 2.

February 20201.8.2

- Revert `rdma-core` to version 25 due to a mismatch in device naming between kernel module and `rdma-core`.

January 20201.8.1

- Update Libfabric to version 1.9.0amzn1.1.

January 20201.8.0

- Update `rdma-core` to version 27.0.

- Update EFA kernel module to version 1.5.0.

- Update Libfabric to version 1.9.0amzn1.0.

- Add option `--minimal` to install just the EFA kernel module and `rdma-core`.

December 20191.7.1

- Update Libfabric to version 1.8.1amzn1.3.

December 20191.7.0

- Add Libfabric module file.

- Update Libfabric to version 1.8.1amzn1.1.

November 20191.6.2

- Update Open MPI to version 1.6.2.

October 20191.6.1

- Update Libfabric to version 1.8.1amzn1.0.

- Update Open MPI to version 4.0.1.

- Update `rdma-core` to version 26.0.

October 20191.5.4

- Update EFA kernel module to version 1.4.1.

September 20191.5.3

- Update EFA kernel module to version 1.3.1.

- Avoid installing `kernel-devel` or `linux-headers` packages unless installing the kernel driver.

September 20191.5.1

- Only configure huge pages when EFA device is present.

August 20191.5.0

- Update Libfabric to version 1.8.0amzn1.1.

- Update `rmda-core` to version 25.0.

August 20191.4.1

- Add Libfabric and Open MPI library paths ( `/opt/amazon/efa/lib64` and `/opt/amazon/efa/openmpi/lib`) to `/etc/ld.so.conf.d/efa.conf` to ensure the Open MPI and Libfabric shared libraries are properly located.

July 20191.4.0

- Update EFA kernel module to version 1.3.0.

- Update Libfabric to version 1.8.0amzn1.0

- First release to support Intel MPI 2019 Update 4.

July 2019

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Verify the EFA installer

EC2 topology
