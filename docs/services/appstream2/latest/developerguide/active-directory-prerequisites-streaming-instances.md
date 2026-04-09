# Domain-Joined WorkSpaces Applications Streaming Instances

SAML 2.0-based user federation is required for application streaming from
domain-joined Always-On and On-Demand fleets. You cannot launch sessions to
domain-joined instances by using [CreateStreamingURL](../../../../reference/appstream2/latest/apireference/api-createstreamingurl.md) or the WorkSpaces Applications
user pool.

Also, you must use an image that supports joining image builders and fleets to an
Active Directory domain. All public images published on or after July 24, 2017
support joining an Active Directory domain. For more information, see [WorkSpaces Applications Base Image and Managed Image Update Release Notes](base-image-version-history.md) and [Tutorial: Setting Up Active Directory](active-directory-directory-setup.md).

###### Note

You can join Always-On and On-Demand fleet streaming instances to an
Active Directory domain. Supported operating systems include Windows, Red Hat
Enterprise Linux, and Rocky Linux.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Active Directory Domain Environment

Group Policy Settings

All content copied from https://docs.aws.amazon.com/.
