# Data Loss Prevention

We'll look at two kinds of data loss prevention.

## Client to AppStream 2.0 Instance Data Transfer Controls

_Table 9 — Guidance for controlling data ingress and egress_

**Setting****Options****Guidance****Clipboard**

- Copy and paste to remote session only

- Copy to local device only

- Disabled

Disabling this setting does not disable copy and paste within the session. If
copying data into the session is required, choose Paste to remote session only to
minimize the potential for data leakage. **File transfer**

- Upload and download

- Upload only

- Download only

- Disabled

Avoid enabling this setting to prevent data leakage. **Print to local device**

- Enabled

- Disabled

If printing is required, use network mapped printers that are controlled and
monitored by your organization.

Consider the advantages of the existing organizational data transfer solution over the
stack settings. These configurations are not designed to replace a comprehensive secure data
transfer solution.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Firewalls and Routing

Controlling egress traffic

All content copied from https://docs.aws.amazon.com/.
