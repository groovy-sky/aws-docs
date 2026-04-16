---
title: "Process overview"
---

# Process overview

This section explains how to integrate your IPAM with AWS accounts outside of your organization. It refers to topics that are covered in other sections of this guide. Keep this page visible, and open the topics linked below in a new window so that you can return to this page for guidance.

When you integrate IPAM with AWS accounts outside of your organization, there are 4 AWS accounts involved in the process:

- **Primary Org Owner** \- The AWS Organizations management account for
organization 1.

- **Primary Org IPAM Account** \- The IPAM delegated administrator
account for organization 1.

- **Secondary Org Owner** \- The AWS Organizations management account for
organization 2.

- **Secondary Org Admin Account** \- The IPAM delegated
administrator account for organization 2.

###### Steps

01. Primary Org Owner delegates a member of their organization as the Primary Org IPAM Account (see [Integrate IPAM with accounts in an AWS Organization](enable-integ-ipam.md)).

02. Primary Org IPAM Account creates an IPAM (see [Create an IPAM](create-ipam.md)).

03. Secondary Org Owner delegates a member of their organization as the Secondary Org Admin
     Account (see [Integrate IPAM with accounts in an AWS Organization](enable-integ-ipam.md)).

04. Secondary Org Admin Account creates a resource discovery and shares it with the Primary Org
     IPAM Account using AWS RAM (see [Create a resource discovery to integrate with another IPAM](res-disc-work-with-create.md) and [Share a resource discovery with another AWS account](res-disc-work-with-share.md)). The resource discovery must be created in the same home Region as the
     Primary Org IPAM.

05. Primary Org IPAM Account accepts the resource share invitation using AWS RAM (see [Accepting and rejecting resource share invitations](../../../ram/latest/userguide/working-with-shared-invitations.md) in the _AWS RAM User Guide_).

06. Primary Org IPAM Account associates the resource discovery with their IPAM (see [Associate a resource discovery with an IPAM](res-disc-work-with-associate.md)).

07. Primary Org IPAM Account can now monitor and/or manage IPAM resources created by the accounts in Secondary Org.

08. (Optional) Primary Org IPAM Account shares IPAM pools with member accounts in Secondary Org (see [Share an IPAM pool using AWS RAM](share-pool-ipam.md)).

09. (Optional) If Primary Org IPAM Account wants to stop discovering resources in Secondary Org, it can disassociate the resource discovery from the IPAM (see [Disassociate a resource discovery](res-disc-work-with-disassociate.md)).

10. (Optional) If the Secondary Org Admin Account wants to stop participating in the Primary Org’s
     IPAM, they can unshare the shared resource discovery (see [Update a\
     resource share in AWS RAM](../../../ram/latest/userguide/working-with-sharing-update.md) in the _AWS RAM User_
    _Guide_) or delete the resource discovery (see [Delete a resource discovery](res-disc-work-with-delete.md)).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Considerations and limitations

Use IPAM with a single account

All content copied from https://docs.aws.amazon.com/.
