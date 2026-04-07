Menu

- [Aws](namespace-aws.md)
- [Ec2](namespace-aws-ec2.md)

## Ec2Client     extends [AwsClient](class-aws-awsclient.md)   in package    - [Aws](package-aws.md)

Client used to interact with **Amazon EC2**.

## Supported API Versions

This class uses a _service description model_ that is associated at
runtime based on the `version` option given when constructing the
client. The `version` option will determine which API operations,
waiters, and paginators are available for a client. Creating a command or a
specific API operation can be done using magic methods (e.g.,
`$client->commandName(/** parameters */)`, or using the
`$client->getCommand` method of the client.

- [**2016-11-15 (latest)**](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html)

  - [AcceptAddressTransfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#acceptaddresstransfer)
  - [AcceptCapacityReservationBillingOwnership](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#acceptcapacityreservationbillingownership)
  - [AcceptReservedInstancesExchangeQuote](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#acceptreservedinstancesexchangequote)
  - [AcceptTransitGatewayMulticastDomainAssociations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#accepttransitgatewaymulticastdomainassociations)
  - [AcceptTransitGatewayPeeringAttachment](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#accepttransitgatewaypeeringattachment)
  - [AcceptTransitGatewayVpcAttachment](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#accepttransitgatewayvpcattachment)
  - [AcceptVpcEndpointConnections](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#acceptvpcendpointconnections)
  - [AcceptVpcPeeringConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#acceptvpcpeeringconnection)
  - [AdvertiseByoipCidr](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#advertisebyoipcidr)
  - [AllocateAddress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#allocateaddress)
  - [AllocateHosts](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#allocatehosts)
  - [AllocateIpamPoolCidr](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#allocateipampoolcidr)
  - [ApplySecurityGroupsToClientVpnTargetNetwork](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#applysecuritygroupstoclientvpntargetnetwork)
  - [AssignIpv6Addresses](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#assignipv6addresses)
  - [AssignPrivateIpAddresses](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#assignprivateipaddresses)
  - [AssignPrivateNatGatewayAddress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#assignprivatenatgatewayaddress)
  - [AssociateAddress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#associateaddress)
  - [AssociateCapacityReservationBillingOwner](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#associatecapacityreservationbillingowner)
  - [AssociateClientVpnTargetNetwork](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#associateclientvpntargetnetwork)
  - [AssociateDhcpOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#associatedhcpoptions)
  - [AssociateEnclaveCertificateIamRole](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#associateenclavecertificateiamrole)
  - [AssociateIamInstanceProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#associateiaminstanceprofile)
  - [AssociateInstanceEventWindow](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#associateinstanceeventwindow)
  - [AssociateIpamByoasn](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#associateipambyoasn)
  - [AssociateIpamResourceDiscovery](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#associateipamresourcediscovery)
  - [AssociateNatGatewayAddress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#associatenatgatewayaddress)
  - [AssociateRouteServer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#associaterouteserver)
  - [AssociateRouteTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#associateroutetable)
  - [AssociateSecurityGroupVpc](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#associatesecuritygroupvpc)
  - [AssociateSubnetCidrBlock](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#associatesubnetcidrblock)
  - [AssociateTransitGatewayMulticastDomain](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#associatetransitgatewaymulticastdomain)
  - [AssociateTransitGatewayPolicyTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#associatetransitgatewaypolicytable)
  - [AssociateTransitGatewayRouteTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#associatetransitgatewayroutetable)
  - [AssociateTrunkInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#associatetrunkinterface)
  - [AssociateVpcCidrBlock](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#associatevpccidrblock)
  - [AttachClassicLinkVpc](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#attachclassiclinkvpc)
  - [AttachInternetGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#attachinternetgateway)
  - [AttachNetworkInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#attachnetworkinterface)
  - [AttachVerifiedAccessTrustProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#attachverifiedaccesstrustprovider)
  - [AttachVolume](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#attachvolume)
  - [AttachVpnGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#attachvpngateway)
  - [AuthorizeClientVpnIngress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#authorizeclientvpningress)
  - [AuthorizeSecurityGroupEgress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#authorizesecuritygroupegress)
  - [AuthorizeSecurityGroupIngress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#authorizesecuritygroupingress)
  - [BundleInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#bundleinstance)
  - [CancelBundleTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#cancelbundletask)
  - [CancelCapacityReservation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#cancelcapacityreservation)
  - [CancelCapacityReservationFleets](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#cancelcapacityreservationfleets)
  - [CancelConversionTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#cancelconversiontask)
  - [CancelDeclarativePoliciesReport](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#canceldeclarativepoliciesreport)
  - [CancelExportTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#cancelexporttask)
  - [CancelImageLaunchPermission](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#cancelimagelaunchpermission)
  - [CancelImportTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#cancelimporttask)
  - [CancelReservedInstancesListing](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#cancelreservedinstanceslisting)
  - [CancelSpotFleetRequests](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#cancelspotfleetrequests)
  - [CancelSpotInstanceRequests](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#cancelspotinstancerequests)
  - [ConfirmProductInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#confirmproductinstance)
  - [CopyFpgaImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#copyfpgaimage)
  - [CopyImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#copyimage)
  - [CopySnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#copysnapshot)
  - [CopyVolumes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#copyvolumes)
  - [CreateCapacityManagerDataExport](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createcapacitymanagerdataexport)
  - [CreateCapacityReservation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createcapacityreservation)
  - [CreateCapacityReservationBySplitting](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createcapacityreservationbysplitting)
  - [CreateCapacityReservationFleet](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createcapacityreservationfleet)
  - [CreateCarrierGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createcarriergateway)
  - [CreateClientVpnEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createclientvpnendpoint)
  - [CreateClientVpnRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createclientvpnroute)
  - [CreateCoipCidr](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createcoipcidr)
  - [CreateCoipPool](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createcoippool)
  - [CreateCustomerGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createcustomergateway)
  - [CreateDefaultSubnet](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createdefaultsubnet)
  - [CreateDefaultVpc](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createdefaultvpc)
  - [CreateDelegateMacVolumeOwnershipTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createdelegatemacvolumeownershiptask)
  - [CreateDhcpOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createdhcpoptions)
  - [CreateEgressOnlyInternetGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createegressonlyinternetgateway)
  - [CreateFleet](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createfleet)
  - [CreateFlowLogs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createflowlogs)
  - [CreateFpgaImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createfpgaimage)
  - [CreateImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createimage)
  - [CreateImageUsageReport](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createimageusagereport)
  - [CreateInstanceConnectEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createinstanceconnectendpoint)
  - [CreateInstanceEventWindow](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createinstanceeventwindow)
  - [CreateInstanceExportTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createinstanceexporttask)
  - [CreateInternetGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createinternetgateway)
  - [CreateInterruptibleCapacityReservationAllocation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createinterruptiblecapacityreservationallocation)
  - [CreateIpam](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createipam)
  - [CreateIpamExternalResourceVerificationToken](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createipamexternalresourceverificationtoken)
  - [CreateIpamPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createipampolicy)
  - [CreateIpamPool](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createipampool)
  - [CreateIpamPrefixListResolver](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createipamprefixlistresolver)
  - [CreateIpamPrefixListResolverTarget](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createipamprefixlistresolvertarget)
  - [CreateIpamResourceDiscovery](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createipamresourcediscovery)
  - [CreateIpamScope](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createipamscope)
  - [CreateKeyPair](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createkeypair)
  - [CreateLaunchTemplate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createlaunchtemplate)
  - [CreateLaunchTemplateVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createlaunchtemplateversion)
  - [CreateLocalGatewayRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createlocalgatewayroute)
  - [CreateLocalGatewayRouteTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createlocalgatewayroutetable)
  - [CreateLocalGatewayRouteTableVirtualInterfaceGroupAssociation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createlocalgatewayroutetablevirtualinterfacegroupassociation)
  - [CreateLocalGatewayRouteTableVpcAssociation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createlocalgatewayroutetablevpcassociation)
  - [CreateLocalGatewayVirtualInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createlocalgatewayvirtualinterface)
  - [CreateLocalGatewayVirtualInterfaceGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createlocalgatewayvirtualinterfacegroup)
  - [CreateMacSystemIntegrityProtectionModificationTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createmacsystemintegrityprotectionmodificationtask)
  - [CreateManagedPrefixList](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createmanagedprefixlist)
  - [CreateNatGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createnatgateway)
  - [CreateNetworkAcl](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createnetworkacl)
  - [CreateNetworkAclEntry](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createnetworkaclentry)
  - [CreateNetworkInsightsAccessScope](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createnetworkinsightsaccessscope)
  - [CreateNetworkInsightsPath](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createnetworkinsightspath)
  - [CreateNetworkInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createnetworkinterface)
  - [CreateNetworkInterfacePermission](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createnetworkinterfacepermission)
  - [CreatePlacementGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createplacementgroup)
  - [CreatePublicIpv4Pool](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createpublicipv4pool)
  - [CreateReplaceRootVolumeTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createreplacerootvolumetask)
  - [CreateReservedInstancesListing](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createreservedinstanceslisting)
  - [CreateRestoreImageTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createrestoreimagetask)
  - [CreateRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createroute)
  - [CreateRouteServer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createrouteserver)
  - [CreateRouteServerEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createrouteserverendpoint)
  - [CreateRouteServerPeer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createrouteserverpeer)
  - [CreateRouteTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createroutetable)
  - [CreateSecondaryNetwork](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createsecondarynetwork)
  - [CreateSecondarySubnet](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createsecondarysubnet)
  - [CreateSecurityGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createsecuritygroup)
  - [CreateSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createsnapshot)
  - [CreateSnapshots](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createsnapshots)
  - [CreateSpotDatafeedSubscription](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createspotdatafeedsubscription)
  - [CreateStoreImageTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createstoreimagetask)
  - [CreateSubnet](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createsubnet)
  - [CreateSubnetCidrReservation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createsubnetcidrreservation)
  - [CreateTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createtags)
  - [CreateTrafficMirrorFilter](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createtrafficmirrorfilter)
  - [CreateTrafficMirrorFilterRule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createtrafficmirrorfilterrule)
  - [CreateTrafficMirrorSession](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createtrafficmirrorsession)
  - [CreateTrafficMirrorTarget](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createtrafficmirrortarget)
  - [CreateTransitGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createtransitgateway)
  - [CreateTransitGatewayConnect](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createtransitgatewayconnect)
  - [CreateTransitGatewayConnectPeer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createtransitgatewayconnectpeer)
  - [CreateTransitGatewayMeteringPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createtransitgatewaymeteringpolicy)
  - [CreateTransitGatewayMeteringPolicyEntry](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createtransitgatewaymeteringpolicyentry)
  - [CreateTransitGatewayMulticastDomain](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createtransitgatewaymulticastdomain)
  - [CreateTransitGatewayPeeringAttachment](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createtransitgatewaypeeringattachment)
  - [CreateTransitGatewayPolicyTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createtransitgatewaypolicytable)
  - [CreateTransitGatewayPrefixListReference](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createtransitgatewayprefixlistreference)
  - [CreateTransitGatewayRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createtransitgatewayroute)
  - [CreateTransitGatewayRouteTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createtransitgatewayroutetable)
  - [CreateTransitGatewayRouteTableAnnouncement](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createtransitgatewayroutetableannouncement)
  - [CreateTransitGatewayVpcAttachment](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createtransitgatewayvpcattachment)
  - [CreateVerifiedAccessEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createverifiedaccessendpoint)
  - [CreateVerifiedAccessGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createverifiedaccessgroup)
  - [CreateVerifiedAccessInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createverifiedaccessinstance)
  - [CreateVerifiedAccessTrustProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createverifiedaccesstrustprovider)
  - [CreateVolume](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createvolume)
  - [CreateVpc](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createvpc)
  - [CreateVpcBlockPublicAccessExclusion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createvpcblockpublicaccessexclusion)
  - [CreateVpcEncryptionControl](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createvpcencryptioncontrol)
  - [CreateVpcEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createvpcendpoint)
  - [CreateVpcEndpointConnectionNotification](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createvpcendpointconnectionnotification)
  - [CreateVpcEndpointServiceConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createvpcendpointserviceconfiguration)
  - [CreateVpcPeeringConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createvpcpeeringconnection)
  - [CreateVpnConcentrator](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createvpnconcentrator)
  - [CreateVpnConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createvpnconnection)
  - [CreateVpnConnectionRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createvpnconnectionroute)
  - [CreateVpnGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#createvpngateway)
  - [DeleteCapacityManagerDataExport](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletecapacitymanagerdataexport)
  - [DeleteCarrierGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletecarriergateway)
  - [DeleteClientVpnEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deleteclientvpnendpoint)
  - [DeleteClientVpnRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deleteclientvpnroute)
  - [DeleteCoipCidr](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletecoipcidr)
  - [DeleteCoipPool](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletecoippool)
  - [DeleteCustomerGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletecustomergateway)
  - [DeleteDhcpOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletedhcpoptions)
  - [DeleteEgressOnlyInternetGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deleteegressonlyinternetgateway)
  - [DeleteFleets](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletefleets)
  - [DeleteFlowLogs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deleteflowlogs)
  - [DeleteFpgaImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletefpgaimage)
  - [DeleteImageUsageReport](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deleteimageusagereport)
  - [DeleteInstanceConnectEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deleteinstanceconnectendpoint)
  - [DeleteInstanceEventWindow](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deleteinstanceeventwindow)
  - [DeleteInternetGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deleteinternetgateway)
  - [DeleteIpam](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deleteipam)
  - [DeleteIpamExternalResourceVerificationToken](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deleteipamexternalresourceverificationtoken)
  - [DeleteIpamPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deleteipampolicy)
  - [DeleteIpamPool](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deleteipampool)
  - [DeleteIpamPrefixListResolver](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deleteipamprefixlistresolver)
  - [DeleteIpamPrefixListResolverTarget](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deleteipamprefixlistresolvertarget)
  - [DeleteIpamResourceDiscovery](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deleteipamresourcediscovery)
  - [DeleteIpamScope](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deleteipamscope)
  - [DeleteKeyPair](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletekeypair)
  - [DeleteLaunchTemplate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletelaunchtemplate)
  - [DeleteLaunchTemplateVersions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletelaunchtemplateversions)
  - [DeleteLocalGatewayRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletelocalgatewayroute)
  - [DeleteLocalGatewayRouteTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletelocalgatewayroutetable)
  - [DeleteLocalGatewayRouteTableVirtualInterfaceGroupAssociation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletelocalgatewayroutetablevirtualinterfacegroupassociation)
  - [DeleteLocalGatewayRouteTableVpcAssociation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletelocalgatewayroutetablevpcassociation)
  - [DeleteLocalGatewayVirtualInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletelocalgatewayvirtualinterface)
  - [DeleteLocalGatewayVirtualInterfaceGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletelocalgatewayvirtualinterfacegroup)
  - [DeleteManagedPrefixList](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletemanagedprefixlist)
  - [DeleteNatGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletenatgateway)
  - [DeleteNetworkAcl](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletenetworkacl)
  - [DeleteNetworkAclEntry](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletenetworkaclentry)
  - [DeleteNetworkInsightsAccessScope](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletenetworkinsightsaccessscope)
  - [DeleteNetworkInsightsAccessScopeAnalysis](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletenetworkinsightsaccessscopeanalysis)
  - [DeleteNetworkInsightsAnalysis](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletenetworkinsightsanalysis)
  - [DeleteNetworkInsightsPath](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletenetworkinsightspath)
  - [DeleteNetworkInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletenetworkinterface)
  - [DeleteNetworkInterfacePermission](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletenetworkinterfacepermission)
  - [DeletePlacementGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deleteplacementgroup)
  - [DeletePublicIpv4Pool](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletepublicipv4pool)
  - [DeleteQueuedReservedInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletequeuedreservedinstances)
  - [DeleteRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deleteroute)
  - [DeleteRouteServer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deleterouteserver)
  - [DeleteRouteServerEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deleterouteserverendpoint)
  - [DeleteRouteServerPeer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deleterouteserverpeer)
  - [DeleteRouteTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deleteroutetable)
  - [DeleteSecondaryNetwork](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletesecondarynetwork)
  - [DeleteSecondarySubnet](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletesecondarysubnet)
  - [DeleteSecurityGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletesecuritygroup)
  - [DeleteSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletesnapshot)
  - [DeleteSpotDatafeedSubscription](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletespotdatafeedsubscription)
  - [DeleteSubnet](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletesubnet)
  - [DeleteSubnetCidrReservation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletesubnetcidrreservation)
  - [DeleteTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletetags)
  - [DeleteTrafficMirrorFilter](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletetrafficmirrorfilter)
  - [DeleteTrafficMirrorFilterRule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletetrafficmirrorfilterrule)
  - [DeleteTrafficMirrorSession](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletetrafficmirrorsession)
  - [DeleteTrafficMirrorTarget](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletetrafficmirrortarget)
  - [DeleteTransitGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletetransitgateway)
  - [DeleteTransitGatewayConnect](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletetransitgatewayconnect)
  - [DeleteTransitGatewayConnectPeer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletetransitgatewayconnectpeer)
  - [DeleteTransitGatewayMeteringPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletetransitgatewaymeteringpolicy)
  - [DeleteTransitGatewayMeteringPolicyEntry](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletetransitgatewaymeteringpolicyentry)
  - [DeleteTransitGatewayMulticastDomain](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletetransitgatewaymulticastdomain)
  - [DeleteTransitGatewayPeeringAttachment](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletetransitgatewaypeeringattachment)
  - [DeleteTransitGatewayPolicyTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletetransitgatewaypolicytable)
  - [DeleteTransitGatewayPrefixListReference](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletetransitgatewayprefixlistreference)
  - [DeleteTransitGatewayRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletetransitgatewayroute)
  - [DeleteTransitGatewayRouteTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletetransitgatewayroutetable)
  - [DeleteTransitGatewayRouteTableAnnouncement](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletetransitgatewayroutetableannouncement)
  - [DeleteTransitGatewayVpcAttachment](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletetransitgatewayvpcattachment)
  - [DeleteVerifiedAccessEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deleteverifiedaccessendpoint)
  - [DeleteVerifiedAccessGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deleteverifiedaccessgroup)
  - [DeleteVerifiedAccessInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deleteverifiedaccessinstance)
  - [DeleteVerifiedAccessTrustProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deleteverifiedaccesstrustprovider)
  - [DeleteVolume](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletevolume)
  - [DeleteVpc](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletevpc)
  - [DeleteVpcBlockPublicAccessExclusion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletevpcblockpublicaccessexclusion)
  - [DeleteVpcEncryptionControl](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletevpcencryptioncontrol)
  - [DeleteVpcEndpointConnectionNotifications](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletevpcendpointconnectionnotifications)
  - [DeleteVpcEndpointServiceConfigurations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletevpcendpointserviceconfigurations)
  - [DeleteVpcEndpoints](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletevpcendpoints)
  - [DeleteVpcPeeringConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletevpcpeeringconnection)
  - [DeleteVpnConcentrator](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletevpnconcentrator)
  - [DeleteVpnConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletevpnconnection)
  - [DeleteVpnConnectionRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletevpnconnectionroute)
  - [DeleteVpnGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deletevpngateway)
  - [DeprovisionByoipCidr](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deprovisionbyoipcidr)
  - [DeprovisionIpamByoasn](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deprovisionipambyoasn)
  - [DeprovisionIpamPoolCidr](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deprovisionipampoolcidr)
  - [DeprovisionPublicIpv4PoolCidr](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deprovisionpublicipv4poolcidr)
  - [DeregisterImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deregisterimage)
  - [DeregisterInstanceEventNotificationAttributes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deregisterinstanceeventnotificationattributes)
  - [DeregisterTransitGatewayMulticastGroupMembers](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deregistertransitgatewaymulticastgroupmembers)
  - [DeregisterTransitGatewayMulticastGroupSources](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#deregistertransitgatewaymulticastgroupsources)
  - [DescribeAccountAttributes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeaccountattributes)
  - [DescribeAddressTransfers](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeaddresstransfers)
  - [DescribeAddresses](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeaddresses)
  - [DescribeAddressesAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeaddressesattribute)
  - [DescribeAggregateIdFormat](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeaggregateidformat)
  - [DescribeAvailabilityZones](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeavailabilityzones)
  - [DescribeAwsNetworkPerformanceMetricSubscriptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeawsnetworkperformancemetricsubscriptions)
  - [DescribeBundleTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describebundletasks)
  - [DescribeByoipCidrs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describebyoipcidrs)
  - [DescribeCapacityBlockExtensionHistory](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describecapacityblockextensionhistory)
  - [DescribeCapacityBlockExtensionOfferings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describecapacityblockextensionofferings)
  - [DescribeCapacityBlockOfferings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describecapacityblockofferings)
  - [DescribeCapacityBlockStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describecapacityblockstatus)
  - [DescribeCapacityBlocks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describecapacityblocks)
  - [DescribeCapacityManagerDataExports](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describecapacitymanagerdataexports)
  - [DescribeCapacityReservationBillingRequests](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describecapacityreservationbillingrequests)
  - [DescribeCapacityReservationFleets](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describecapacityreservationfleets)
  - [DescribeCapacityReservationTopology](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describecapacityreservationtopology)
  - [DescribeCapacityReservations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describecapacityreservations)
  - [DescribeCarrierGateways](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describecarriergateways)
  - [DescribeClassicLinkInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeclassiclinkinstances)
  - [DescribeClientVpnAuthorizationRules](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeclientvpnauthorizationrules)
  - [DescribeClientVpnConnections](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeclientvpnconnections)
  - [DescribeClientVpnEndpoints](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeclientvpnendpoints)
  - [DescribeClientVpnRoutes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeclientvpnroutes)
  - [DescribeClientVpnTargetNetworks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeclientvpntargetnetworks)
  - [DescribeCoipPools](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describecoippools)
  - [DescribeConversionTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeconversiontasks)
  - [DescribeCustomerGateways](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describecustomergateways)
  - [DescribeDeclarativePoliciesReports](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describedeclarativepoliciesreports)
  - [DescribeDhcpOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describedhcpoptions)
  - [DescribeEgressOnlyInternetGateways](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeegressonlyinternetgateways)
  - [DescribeElasticGpus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeelasticgpus)
  - [DescribeExportImageTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeexportimagetasks)
  - [DescribeExportTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeexporttasks)
  - [DescribeFastLaunchImages](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describefastlaunchimages)
  - [DescribeFastSnapshotRestores](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describefastsnapshotrestores)
  - [DescribeFleetHistory](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describefleethistory)
  - [DescribeFleetInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describefleetinstances)
  - [DescribeFleets](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describefleets)
  - [DescribeFlowLogs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeflowlogs)
  - [DescribeFpgaImageAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describefpgaimageattribute)
  - [DescribeFpgaImages](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describefpgaimages)
  - [DescribeHostReservationOfferings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describehostreservationofferings)
  - [DescribeHostReservations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describehostreservations)
  - [DescribeHosts](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describehosts)
  - [DescribeIamInstanceProfileAssociations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeiaminstanceprofileassociations)
  - [DescribeIdFormat](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeidformat)
  - [DescribeIdentityIdFormat](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeidentityidformat)
  - [DescribeImageAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeimageattribute)
  - [DescribeImageReferences](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeimagereferences)
  - [DescribeImageUsageReportEntries](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeimageusagereportentries)
  - [DescribeImageUsageReports](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeimageusagereports)
  - [DescribeImages](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeimages)
  - [DescribeImportImageTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeimportimagetasks)
  - [DescribeImportSnapshotTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeimportsnapshottasks)
  - [DescribeInstanceAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeinstanceattribute)
  - [DescribeInstanceConnectEndpoints](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeinstanceconnectendpoints)
  - [DescribeInstanceCreditSpecifications](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeinstancecreditspecifications)
  - [DescribeInstanceEventNotificationAttributes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeinstanceeventnotificationattributes)
  - [DescribeInstanceEventWindows](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeinstanceeventwindows)
  - [DescribeInstanceImageMetadata](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeinstanceimagemetadata)
  - [DescribeInstanceSqlHaHistoryStates](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeinstancesqlhahistorystates)
  - [DescribeInstanceSqlHaStates](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeinstancesqlhastates)
  - [DescribeInstanceStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeinstancestatus)
  - [DescribeInstanceTopology](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeinstancetopology)
  - [DescribeInstanceTypeOfferings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeinstancetypeofferings)
  - [DescribeInstanceTypes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeinstancetypes)
  - [DescribeInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeinstances)
  - [DescribeInternetGateways](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeinternetgateways)
  - [DescribeIpamByoasn](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeipambyoasn)
  - [DescribeIpamExternalResourceVerificationTokens](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeipamexternalresourceverificationtokens)
  - [DescribeIpamPolicies](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeipampolicies)
  - [DescribeIpamPools](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeipampools)
  - [DescribeIpamPrefixListResolverTargets](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeipamprefixlistresolvertargets)
  - [DescribeIpamPrefixListResolvers](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeipamprefixlistresolvers)
  - [DescribeIpamResourceDiscoveries](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeipamresourcediscoveries)
  - [DescribeIpamResourceDiscoveryAssociations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeipamresourcediscoveryassociations)
  - [DescribeIpamScopes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeipamscopes)
  - [DescribeIpams](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeipams)
  - [DescribeIpv6Pools](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeipv6pools)
  - [DescribeKeyPairs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describekeypairs)
  - [DescribeLaunchTemplateVersions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describelaunchtemplateversions)
  - [DescribeLaunchTemplates](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describelaunchtemplates)
  - [DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describelocalgatewayroutetablevirtualinterfacegroupassociations)
  - [DescribeLocalGatewayRouteTableVpcAssociations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describelocalgatewayroutetablevpcassociations)
  - [DescribeLocalGatewayRouteTables](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describelocalgatewayroutetables)
  - [DescribeLocalGatewayVirtualInterfaceGroups](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describelocalgatewayvirtualinterfacegroups)
  - [DescribeLocalGatewayVirtualInterfaces](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describelocalgatewayvirtualinterfaces)
  - [DescribeLocalGateways](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describelocalgateways)
  - [DescribeLockedSnapshots](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describelockedsnapshots)
  - [DescribeMacHosts](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describemachosts)
  - [DescribeMacModificationTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describemacmodificationtasks)
  - [DescribeManagedPrefixLists](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describemanagedprefixlists)
  - [DescribeMovingAddresses](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describemovingaddresses)
  - [DescribeNatGateways](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describenatgateways)
  - [DescribeNetworkAcls](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describenetworkacls)
  - [DescribeNetworkInsightsAccessScopeAnalyses](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describenetworkinsightsaccessscopeanalyses)
  - [DescribeNetworkInsightsAccessScopes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describenetworkinsightsaccessscopes)
  - [DescribeNetworkInsightsAnalyses](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describenetworkinsightsanalyses)
  - [DescribeNetworkInsightsPaths](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describenetworkinsightspaths)
  - [DescribeNetworkInterfaceAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describenetworkinterfaceattribute)
  - [DescribeNetworkInterfacePermissions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describenetworkinterfacepermissions)
  - [DescribeNetworkInterfaces](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describenetworkinterfaces)
  - [DescribeOutpostLags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeoutpostlags)
  - [DescribePlacementGroups](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeplacementgroups)
  - [DescribePrefixLists](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeprefixlists)
  - [DescribePrincipalIdFormat](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeprincipalidformat)
  - [DescribePublicIpv4Pools](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describepublicipv4pools)
  - [DescribeRegions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeregions)
  - [DescribeReplaceRootVolumeTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describereplacerootvolumetasks)
  - [DescribeReservedInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describereservedinstances)
  - [DescribeReservedInstancesListings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describereservedinstanceslistings)
  - [DescribeReservedInstancesModifications](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describereservedinstancesmodifications)
  - [DescribeReservedInstancesOfferings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describereservedinstancesofferings)
  - [DescribeRouteServerEndpoints](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describerouteserverendpoints)
  - [DescribeRouteServerPeers](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describerouteserverpeers)
  - [DescribeRouteServers](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describerouteservers)
  - [DescribeRouteTables](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeroutetables)
  - [DescribeScheduledInstanceAvailability](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describescheduledinstanceavailability)
  - [DescribeScheduledInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describescheduledinstances)
  - [DescribeSecondaryInterfaces](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describesecondaryinterfaces)
  - [DescribeSecondaryNetworks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describesecondarynetworks)
  - [DescribeSecondarySubnets](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describesecondarysubnets)
  - [DescribeSecurityGroupReferences](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describesecuritygroupreferences)
  - [DescribeSecurityGroupRules](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describesecuritygrouprules)
  - [DescribeSecurityGroupVpcAssociations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describesecuritygroupvpcassociations)
  - [DescribeSecurityGroups](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describesecuritygroups)
  - [DescribeServiceLinkVirtualInterfaces](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeservicelinkvirtualinterfaces)
  - [DescribeSnapshotAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describesnapshotattribute)
  - [DescribeSnapshotTierStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describesnapshottierstatus)
  - [DescribeSnapshots](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describesnapshots)
  - [DescribeSpotDatafeedSubscription](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describespotdatafeedsubscription)
  - [DescribeSpotFleetInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describespotfleetinstances)
  - [DescribeSpotFleetRequestHistory](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describespotfleetrequesthistory)
  - [DescribeSpotFleetRequests](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describespotfleetrequests)
  - [DescribeSpotInstanceRequests](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describespotinstancerequests)
  - [DescribeSpotPriceHistory](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describespotpricehistory)
  - [DescribeStaleSecurityGroups](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describestalesecuritygroups)
  - [DescribeStoreImageTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describestoreimagetasks)
  - [DescribeSubnets](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describesubnets)
  - [DescribeTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describetags)
  - [DescribeTrafficMirrorFilterRules](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describetrafficmirrorfilterrules)
  - [DescribeTrafficMirrorFilters](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describetrafficmirrorfilters)
  - [DescribeTrafficMirrorSessions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describetrafficmirrorsessions)
  - [DescribeTrafficMirrorTargets](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describetrafficmirrortargets)
  - [DescribeTransitGatewayAttachments](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describetransitgatewayattachments)
  - [DescribeTransitGatewayConnectPeers](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describetransitgatewayconnectpeers)
  - [DescribeTransitGatewayConnects](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describetransitgatewayconnects)
  - [DescribeTransitGatewayMeteringPolicies](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describetransitgatewaymeteringpolicies)
  - [DescribeTransitGatewayMulticastDomains](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describetransitgatewaymulticastdomains)
  - [DescribeTransitGatewayPeeringAttachments](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describetransitgatewaypeeringattachments)
  - [DescribeTransitGatewayPolicyTables](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describetransitgatewaypolicytables)
  - [DescribeTransitGatewayRouteTableAnnouncements](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describetransitgatewayroutetableannouncements)
  - [DescribeTransitGatewayRouteTables](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describetransitgatewayroutetables)
  - [DescribeTransitGatewayVpcAttachments](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describetransitgatewayvpcattachments)
  - [DescribeTransitGateways](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describetransitgateways)
  - [DescribeTrunkInterfaceAssociations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describetrunkinterfaceassociations)
  - [DescribeVerifiedAccessEndpoints](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeverifiedaccessendpoints)
  - [DescribeVerifiedAccessGroups](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeverifiedaccessgroups)
  - [DescribeVerifiedAccessInstanceLoggingConfigurations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeverifiedaccessinstanceloggingconfigurations)
  - [DescribeVerifiedAccessInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeverifiedaccessinstances)
  - [DescribeVerifiedAccessTrustProviders](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describeverifiedaccesstrustproviders)
  - [DescribeVolumeAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describevolumeattribute)
  - [DescribeVolumeStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describevolumestatus)
  - [DescribeVolumes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describevolumes)
  - [DescribeVolumesModifications](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describevolumesmodifications)
  - [DescribeVpcAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describevpcattribute)
  - [DescribeVpcBlockPublicAccessExclusions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describevpcblockpublicaccessexclusions)
  - [DescribeVpcBlockPublicAccessOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describevpcblockpublicaccessoptions)
  - [DescribeVpcClassicLink](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describevpcclassiclink)
  - [DescribeVpcClassicLinkDnsSupport](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describevpcclassiclinkdnssupport)
  - [DescribeVpcEncryptionControls](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describevpcencryptioncontrols)
  - [DescribeVpcEndpointAssociations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describevpcendpointassociations)
  - [DescribeVpcEndpointConnectionNotifications](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describevpcendpointconnectionnotifications)
  - [DescribeVpcEndpointConnections](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describevpcendpointconnections)
  - [DescribeVpcEndpointServiceConfigurations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describevpcendpointserviceconfigurations)
  - [DescribeVpcEndpointServicePermissions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describevpcendpointservicepermissions)
  - [DescribeVpcEndpointServices](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describevpcendpointservices)
  - [DescribeVpcEndpoints](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describevpcendpoints)
  - [DescribeVpcPeeringConnections](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describevpcpeeringconnections)
  - [DescribeVpcs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describevpcs)
  - [DescribeVpnConcentrators](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describevpnconcentrators)
  - [DescribeVpnConnections](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describevpnconnections)
  - [DescribeVpnGateways](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#describevpngateways)
  - [DetachClassicLinkVpc](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#detachclassiclinkvpc)
  - [DetachInternetGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#detachinternetgateway)
  - [DetachNetworkInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#detachnetworkinterface)
  - [DetachVerifiedAccessTrustProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#detachverifiedaccesstrustprovider)
  - [DetachVolume](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#detachvolume)
  - [DetachVpnGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#detachvpngateway)
  - [DisableAddressTransfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disableaddresstransfer)
  - [DisableAllowedImagesSettings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disableallowedimagessettings)
  - [DisableAwsNetworkPerformanceMetricSubscription](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disableawsnetworkperformancemetricsubscription)
  - [DisableCapacityManager](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disablecapacitymanager)
  - [DisableEbsEncryptionByDefault](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disableebsencryptionbydefault)
  - [DisableFastLaunch](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disablefastlaunch)
  - [DisableFastSnapshotRestores](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disablefastsnapshotrestores)
  - [DisableImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disableimage)
  - [DisableImageBlockPublicAccess](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disableimageblockpublicaccess)
  - [DisableImageDeprecation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disableimagedeprecation)
  - [DisableImageDeregistrationProtection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disableimagederegistrationprotection)
  - [DisableInstanceSqlHaStandbyDetections](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disableinstancesqlhastandbydetections)
  - [DisableIpamOrganizationAdminAccount](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disableipamorganizationadminaccount)
  - [DisableIpamPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disableipampolicy)
  - [DisableRouteServerPropagation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disablerouteserverpropagation)
  - [DisableSerialConsoleAccess](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disableserialconsoleaccess)
  - [DisableSnapshotBlockPublicAccess](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disablesnapshotblockpublicaccess)
  - [DisableTransitGatewayRouteTablePropagation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disabletransitgatewayroutetablepropagation)
  - [DisableVgwRoutePropagation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disablevgwroutepropagation)
  - [DisableVpcClassicLink](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disablevpcclassiclink)
  - [DisableVpcClassicLinkDnsSupport](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disablevpcclassiclinkdnssupport)
  - [DisassociateAddress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disassociateaddress)
  - [DisassociateCapacityReservationBillingOwner](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disassociatecapacityreservationbillingowner)
  - [DisassociateClientVpnTargetNetwork](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disassociateclientvpntargetnetwork)
  - [DisassociateEnclaveCertificateIamRole](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disassociateenclavecertificateiamrole)
  - [DisassociateIamInstanceProfile](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disassociateiaminstanceprofile)
  - [DisassociateInstanceEventWindow](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disassociateinstanceeventwindow)
  - [DisassociateIpamByoasn](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disassociateipambyoasn)
  - [DisassociateIpamResourceDiscovery](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disassociateipamresourcediscovery)
  - [DisassociateNatGatewayAddress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disassociatenatgatewayaddress)
  - [DisassociateRouteServer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disassociaterouteserver)
  - [DisassociateRouteTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disassociateroutetable)
  - [DisassociateSecurityGroupVpc](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disassociatesecuritygroupvpc)
  - [DisassociateSubnetCidrBlock](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disassociatesubnetcidrblock)
  - [DisassociateTransitGatewayMulticastDomain](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disassociatetransitgatewaymulticastdomain)
  - [DisassociateTransitGatewayPolicyTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disassociatetransitgatewaypolicytable)
  - [DisassociateTransitGatewayRouteTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disassociatetransitgatewayroutetable)
  - [DisassociateTrunkInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disassociatetrunkinterface)
  - [DisassociateVpcCidrBlock](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#disassociatevpccidrblock)
  - [EnableAddressTransfer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#enableaddresstransfer)
  - [EnableAllowedImagesSettings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#enableallowedimagessettings)
  - [EnableAwsNetworkPerformanceMetricSubscription](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#enableawsnetworkperformancemetricsubscription)
  - [EnableCapacityManager](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#enablecapacitymanager)
  - [EnableEbsEncryptionByDefault](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#enableebsencryptionbydefault)
  - [EnableFastLaunch](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#enablefastlaunch)
  - [EnableFastSnapshotRestores](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#enablefastsnapshotrestores)
  - [EnableImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#enableimage)
  - [EnableImageBlockPublicAccess](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#enableimageblockpublicaccess)
  - [EnableImageDeprecation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#enableimagedeprecation)
  - [EnableImageDeregistrationProtection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#enableimagederegistrationprotection)
  - [EnableInstanceSqlHaStandbyDetections](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#enableinstancesqlhastandbydetections)
  - [EnableIpamOrganizationAdminAccount](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#enableipamorganizationadminaccount)
  - [EnableIpamPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#enableipampolicy)
  - [EnableReachabilityAnalyzerOrganizationSharing](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#enablereachabilityanalyzerorganizationsharing)
  - [EnableRouteServerPropagation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#enablerouteserverpropagation)
  - [EnableSerialConsoleAccess](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#enableserialconsoleaccess)
  - [EnableSnapshotBlockPublicAccess](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#enablesnapshotblockpublicaccess)
  - [EnableTransitGatewayRouteTablePropagation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#enabletransitgatewayroutetablepropagation)
  - [EnableVgwRoutePropagation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#enablevgwroutepropagation)
  - [EnableVolumeIO](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#enablevolumeio)
  - [EnableVpcClassicLink](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#enablevpcclassiclink)
  - [EnableVpcClassicLinkDnsSupport](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#enablevpcclassiclinkdnssupport)
  - [ExportClientVpnClientCertificateRevocationList](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#exportclientvpnclientcertificaterevocationlist)
  - [ExportClientVpnClientConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#exportclientvpnclientconfiguration)
  - [ExportImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#exportimage)
  - [ExportTransitGatewayRoutes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#exporttransitgatewayroutes)
  - [ExportVerifiedAccessInstanceClientConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#exportverifiedaccessinstanceclientconfiguration)
  - [GetActiveVpnTunnelStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getactivevpntunnelstatus)
  - [GetAllowedImagesSettings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getallowedimagessettings)
  - [GetAssociatedEnclaveCertificateIamRoles](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getassociatedenclavecertificateiamroles)
  - [GetAssociatedIpv6PoolCidrs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getassociatedipv6poolcidrs)
  - [GetAwsNetworkPerformanceData](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getawsnetworkperformancedata)
  - [GetCapacityManagerAttributes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getcapacitymanagerattributes)
  - [GetCapacityManagerMetricData](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getcapacitymanagermetricdata)
  - [GetCapacityManagerMetricDimensions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getcapacitymanagermetricdimensions)
  - [GetCapacityReservationUsage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getcapacityreservationusage)
  - [GetCoipPoolUsage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getcoippoolusage)
  - [GetConsoleOutput](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getconsoleoutput)
  - [GetConsoleScreenshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getconsolescreenshot)
  - [GetDeclarativePoliciesReportSummary](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getdeclarativepoliciesreportsummary)
  - [GetDefaultCreditSpecification](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getdefaultcreditspecification)
  - [GetEbsDefaultKmsKeyId](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getebsdefaultkmskeyid)
  - [GetEbsEncryptionByDefault](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getebsencryptionbydefault)
  - [GetEnabledIpamPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getenabledipampolicy)
  - [GetFlowLogsIntegrationTemplate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getflowlogsintegrationtemplate)
  - [GetGroupsForCapacityReservation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getgroupsforcapacityreservation)
  - [GetHostReservationPurchasePreview](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#gethostreservationpurchasepreview)
  - [GetImageAncestry](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getimageancestry)
  - [GetImageBlockPublicAccessState](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getimageblockpublicaccessstate)
  - [GetInstanceMetadataDefaults](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getinstancemetadatadefaults)
  - [GetInstanceTpmEkPub](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getinstancetpmekpub)
  - [GetInstanceTypesFromInstanceRequirements](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getinstancetypesfrominstancerequirements)
  - [GetInstanceUefiData](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getinstanceuefidata)
  - [GetIpamAddressHistory](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getipamaddresshistory)
  - [GetIpamDiscoveredAccounts](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getipamdiscoveredaccounts)
  - [GetIpamDiscoveredPublicAddresses](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getipamdiscoveredpublicaddresses)
  - [GetIpamDiscoveredResourceCidrs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getipamdiscoveredresourcecidrs)
  - [GetIpamPolicyAllocationRules](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getipampolicyallocationrules)
  - [GetIpamPolicyOrganizationTargets](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getipampolicyorganizationtargets)
  - [GetIpamPoolAllocations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getipampoolallocations)
  - [GetIpamPoolCidrs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getipampoolcidrs)
  - [GetIpamPrefixListResolverRules](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getipamprefixlistresolverrules)
  - [GetIpamPrefixListResolverVersionEntries](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getipamprefixlistresolverversionentries)
  - [GetIpamPrefixListResolverVersions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getipamprefixlistresolverversions)
  - [GetIpamResourceCidrs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getipamresourcecidrs)
  - [GetLaunchTemplateData](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getlaunchtemplatedata)
  - [GetManagedPrefixListAssociations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getmanagedprefixlistassociations)
  - [GetManagedPrefixListEntries](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getmanagedprefixlistentries)
  - [GetNetworkInsightsAccessScopeAnalysisFindings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getnetworkinsightsaccessscopeanalysisfindings)
  - [GetNetworkInsightsAccessScopeContent](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getnetworkinsightsaccessscopecontent)
  - [GetPasswordData](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getpassworddata)
  - [GetReservedInstancesExchangeQuote](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getreservedinstancesexchangequote)
  - [GetRouteServerAssociations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getrouteserverassociations)
  - [GetRouteServerPropagations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getrouteserverpropagations)
  - [GetRouteServerRoutingDatabase](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getrouteserverroutingdatabase)
  - [GetSecurityGroupsForVpc](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getsecuritygroupsforvpc)
  - [GetSerialConsoleAccessStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getserialconsoleaccessstatus)
  - [GetSnapshotBlockPublicAccessState](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getsnapshotblockpublicaccessstate)
  - [GetSpotPlacementScores](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getspotplacementscores)
  - [GetSubnetCidrReservations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getsubnetcidrreservations)
  - [GetTransitGatewayAttachmentPropagations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#gettransitgatewayattachmentpropagations)
  - [GetTransitGatewayMeteringPolicyEntries](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#gettransitgatewaymeteringpolicyentries)
  - [GetTransitGatewayMulticastDomainAssociations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#gettransitgatewaymulticastdomainassociations)
  - [GetTransitGatewayPolicyTableAssociations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#gettransitgatewaypolicytableassociations)
  - [GetTransitGatewayPolicyTableEntries](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#gettransitgatewaypolicytableentries)
  - [GetTransitGatewayPrefixListReferences](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#gettransitgatewayprefixlistreferences)
  - [GetTransitGatewayRouteTableAssociations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#gettransitgatewayroutetableassociations)
  - [GetTransitGatewayRouteTablePropagations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#gettransitgatewayroutetablepropagations)
  - [GetVerifiedAccessEndpointPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getverifiedaccessendpointpolicy)
  - [GetVerifiedAccessEndpointTargets](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getverifiedaccessendpointtargets)
  - [GetVerifiedAccessGroupPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getverifiedaccessgrouppolicy)
  - [GetVpcResourcesBlockingEncryptionEnforcement](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getvpcresourcesblockingencryptionenforcement)
  - [GetVpnConnectionDeviceSampleConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getvpnconnectiondevicesampleconfiguration)
  - [GetVpnConnectionDeviceTypes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getvpnconnectiondevicetypes)
  - [GetVpnTunnelReplacementStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#getvpntunnelreplacementstatus)
  - [ImportClientVpnClientCertificateRevocationList](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#importclientvpnclientcertificaterevocationlist)
  - [ImportImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#importimage)
  - [ImportInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#importinstance)
  - [ImportKeyPair](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#importkeypair)
  - [ImportSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#importsnapshot)
  - [ImportVolume](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#importvolume)
  - [ListImagesInRecycleBin](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#listimagesinrecyclebin)
  - [ListSnapshotsInRecycleBin](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#listsnapshotsinrecyclebin)
  - [ListVolumesInRecycleBin](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#listvolumesinrecyclebin)
  - [LockSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#locksnapshot)
  - [ModifyAddressAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyaddressattribute)
  - [ModifyAvailabilityZoneGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyavailabilityzonegroup)
  - [ModifyCapacityReservation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifycapacityreservation)
  - [ModifyCapacityReservationFleet](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifycapacityreservationfleet)
  - [ModifyClientVpnEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyclientvpnendpoint)
  - [ModifyDefaultCreditSpecification](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifydefaultcreditspecification)
  - [ModifyEbsDefaultKmsKeyId](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyebsdefaultkmskeyid)
  - [ModifyFleet](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyfleet)
  - [ModifyFpgaImageAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyfpgaimageattribute)
  - [ModifyHosts](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyhosts)
  - [ModifyIdFormat](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyidformat)
  - [ModifyIdentityIdFormat](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyidentityidformat)
  - [ModifyImageAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyimageattribute)
  - [ModifyInstanceAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyinstanceattribute)
  - [ModifyInstanceCapacityReservationAttributes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyinstancecapacityreservationattributes)
  - [ModifyInstanceConnectEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyinstanceconnectendpoint)
  - [ModifyInstanceCpuOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyinstancecpuoptions)
  - [ModifyInstanceCreditSpecification](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyinstancecreditspecification)
  - [ModifyInstanceEventStartTime](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyinstanceeventstarttime)
  - [ModifyInstanceEventWindow](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyinstanceeventwindow)
  - [ModifyInstanceMaintenanceOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyinstancemaintenanceoptions)
  - [ModifyInstanceMetadataDefaults](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyinstancemetadatadefaults)
  - [ModifyInstanceMetadataOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyinstancemetadataoptions)
  - [ModifyInstanceNetworkPerformanceOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyinstancenetworkperformanceoptions)
  - [ModifyInstancePlacement](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyinstanceplacement)
  - [ModifyIpam](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyipam)
  - [ModifyIpamPolicyAllocationRules](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyipampolicyallocationrules)
  - [ModifyIpamPool](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyipampool)
  - [ModifyIpamPrefixListResolver](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyipamprefixlistresolver)
  - [ModifyIpamPrefixListResolverTarget](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyipamprefixlistresolvertarget)
  - [ModifyIpamResourceCidr](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyipamresourcecidr)
  - [ModifyIpamResourceDiscovery](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyipamresourcediscovery)
  - [ModifyIpamScope](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyipamscope)
  - [ModifyLaunchTemplate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifylaunchtemplate)
  - [ModifyLocalGatewayRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifylocalgatewayroute)
  - [ModifyManagedPrefixList](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifymanagedprefixlist)
  - [ModifyNetworkInterfaceAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifynetworkinterfaceattribute)
  - [ModifyPrivateDnsNameOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyprivatednsnameoptions)
  - [ModifyPublicIpDnsNameOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifypublicipdnsnameoptions)
  - [ModifyReservedInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyreservedinstances)
  - [ModifyRouteServer](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyrouteserver)
  - [ModifySecurityGroupRules](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifysecuritygrouprules)
  - [ModifySnapshotAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifysnapshotattribute)
  - [ModifySnapshotTier](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifysnapshottier)
  - [ModifySpotFleetRequest](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyspotfleetrequest)
  - [ModifySubnetAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifysubnetattribute)
  - [ModifyTrafficMirrorFilterNetworkServices](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifytrafficmirrorfilternetworkservices)
  - [ModifyTrafficMirrorFilterRule](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifytrafficmirrorfilterrule)
  - [ModifyTrafficMirrorSession](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifytrafficmirrorsession)
  - [ModifyTransitGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifytransitgateway)
  - [ModifyTransitGatewayMeteringPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifytransitgatewaymeteringpolicy)
  - [ModifyTransitGatewayPrefixListReference](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifytransitgatewayprefixlistreference)
  - [ModifyTransitGatewayVpcAttachment](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifytransitgatewayvpcattachment)
  - [ModifyVerifiedAccessEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyverifiedaccessendpoint)
  - [ModifyVerifiedAccessEndpointPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyverifiedaccessendpointpolicy)
  - [ModifyVerifiedAccessGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyverifiedaccessgroup)
  - [ModifyVerifiedAccessGroupPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyverifiedaccessgrouppolicy)
  - [ModifyVerifiedAccessInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyverifiedaccessinstance)
  - [ModifyVerifiedAccessInstanceLoggingConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyverifiedaccessinstanceloggingconfiguration)
  - [ModifyVerifiedAccessTrustProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyverifiedaccesstrustprovider)
  - [ModifyVolume](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyvolume)
  - [ModifyVolumeAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyvolumeattribute)
  - [ModifyVpcAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyvpcattribute)
  - [ModifyVpcBlockPublicAccessExclusion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyvpcblockpublicaccessexclusion)
  - [ModifyVpcBlockPublicAccessOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyvpcblockpublicaccessoptions)
  - [ModifyVpcEncryptionControl](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyvpcencryptioncontrol)
  - [ModifyVpcEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyvpcendpoint)
  - [ModifyVpcEndpointConnectionNotification](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyvpcendpointconnectionnotification)
  - [ModifyVpcEndpointServiceConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyvpcendpointserviceconfiguration)
  - [ModifyVpcEndpointServicePayerResponsibility](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyvpcendpointservicepayerresponsibility)
  - [ModifyVpcEndpointServicePermissions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyvpcendpointservicepermissions)
  - [ModifyVpcPeeringConnectionOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyvpcpeeringconnectionoptions)
  - [ModifyVpcTenancy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyvpctenancy)
  - [ModifyVpnConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyvpnconnection)
  - [ModifyVpnConnectionOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyvpnconnectionoptions)
  - [ModifyVpnTunnelCertificate](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyvpntunnelcertificate)
  - [ModifyVpnTunnelOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#modifyvpntunneloptions)
  - [MonitorInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#monitorinstances)
  - [MoveAddressToVpc](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#moveaddresstovpc)
  - [MoveByoipCidrToIpam](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#movebyoipcidrtoipam)
  - [MoveCapacityReservationInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#movecapacityreservationinstances)
  - [ProvisionByoipCidr](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#provisionbyoipcidr)
  - [ProvisionIpamByoasn](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#provisionipambyoasn)
  - [ProvisionIpamPoolCidr](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#provisionipampoolcidr)
  - [ProvisionPublicIpv4PoolCidr](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#provisionpublicipv4poolcidr)
  - [PurchaseCapacityBlock](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#purchasecapacityblock)
  - [PurchaseCapacityBlockExtension](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#purchasecapacityblockextension)
  - [PurchaseHostReservation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#purchasehostreservation)
  - [PurchaseReservedInstancesOffering](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#purchasereservedinstancesoffering)
  - [PurchaseScheduledInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#purchasescheduledinstances)
  - [RebootInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#rebootinstances)
  - [RegisterImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#registerimage)
  - [RegisterInstanceEventNotificationAttributes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#registerinstanceeventnotificationattributes)
  - [RegisterTransitGatewayMulticastGroupMembers](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#registertransitgatewaymulticastgroupmembers)
  - [RegisterTransitGatewayMulticastGroupSources](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#registertransitgatewaymulticastgroupsources)
  - [RejectCapacityReservationBillingOwnership](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#rejectcapacityreservationbillingownership)
  - [RejectTransitGatewayMulticastDomainAssociations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#rejecttransitgatewaymulticastdomainassociations)
  - [RejectTransitGatewayPeeringAttachment](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#rejecttransitgatewaypeeringattachment)
  - [RejectTransitGatewayVpcAttachment](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#rejecttransitgatewayvpcattachment)
  - [RejectVpcEndpointConnections](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#rejectvpcendpointconnections)
  - [RejectVpcPeeringConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#rejectvpcpeeringconnection)
  - [ReleaseAddress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#releaseaddress)
  - [ReleaseHosts](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#releasehosts)
  - [ReleaseIpamPoolAllocation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#releaseipampoolallocation)
  - [ReplaceIamInstanceProfileAssociation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#replaceiaminstanceprofileassociation)
  - [ReplaceImageCriteriaInAllowedImagesSettings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#replaceimagecriteriainallowedimagessettings)
  - [ReplaceNetworkAclAssociation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#replacenetworkaclassociation)
  - [ReplaceNetworkAclEntry](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#replacenetworkaclentry)
  - [ReplaceRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#replaceroute)
  - [ReplaceRouteTableAssociation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#replaceroutetableassociation)
  - [ReplaceTransitGatewayRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#replacetransitgatewayroute)
  - [ReplaceVpnTunnel](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#replacevpntunnel)
  - [ReportInstanceStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#reportinstancestatus)
  - [RequestSpotFleet](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#requestspotfleet)
  - [RequestSpotInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#requestspotinstances)
  - [ResetAddressAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#resetaddressattribute)
  - [ResetEbsDefaultKmsKeyId](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#resetebsdefaultkmskeyid)
  - [ResetFpgaImageAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#resetfpgaimageattribute)
  - [ResetImageAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#resetimageattribute)
  - [ResetInstanceAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#resetinstanceattribute)
  - [ResetNetworkInterfaceAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#resetnetworkinterfaceattribute)
  - [ResetSnapshotAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#resetsnapshotattribute)
  - [RestoreAddressToClassic](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#restoreaddresstoclassic)
  - [RestoreImageFromRecycleBin](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#restoreimagefromrecyclebin)
  - [RestoreManagedPrefixListVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#restoremanagedprefixlistversion)
  - [RestoreSnapshotFromRecycleBin](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#restoresnapshotfromrecyclebin)
  - [RestoreSnapshotTier](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#restoresnapshottier)
  - [RestoreVolumeFromRecycleBin](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#restorevolumefromrecyclebin)
  - [RevokeClientVpnIngress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#revokeclientvpningress)
  - [RevokeSecurityGroupEgress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#revokesecuritygroupegress)
  - [RevokeSecurityGroupIngress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#revokesecuritygroupingress)
  - [RunInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#runinstances)
  - [RunScheduledInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#runscheduledinstances)
  - [SearchLocalGatewayRoutes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#searchlocalgatewayroutes)
  - [SearchTransitGatewayMulticastGroups](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#searchtransitgatewaymulticastgroups)
  - [SearchTransitGatewayRoutes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#searchtransitgatewayroutes)
  - [SendDiagnosticInterrupt](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#senddiagnosticinterrupt)
  - [StartDeclarativePoliciesReport](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#startdeclarativepoliciesreport)
  - [StartInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#startinstances)
  - [StartNetworkInsightsAccessScopeAnalysis](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#startnetworkinsightsaccessscopeanalysis)
  - [StartNetworkInsightsAnalysis](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#startnetworkinsightsanalysis)
  - [StartVpcEndpointServicePrivateDnsVerification](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#startvpcendpointserviceprivatednsverification)
  - [StopInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#stopinstances)
  - [TerminateClientVpnConnections](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#terminateclientvpnconnections)
  - [TerminateInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#terminateinstances)
  - [UnassignIpv6Addresses](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#unassignipv6addresses)
  - [UnassignPrivateIpAddresses](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#unassignprivateipaddresses)
  - [UnassignPrivateNatGatewayAddress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#unassignprivatenatgatewayaddress)
  - [UnlockSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#unlocksnapshot)
  - [UnmonitorInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#unmonitorinstances)
  - [UpdateCapacityManagerOrganizationsAccess](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#updatecapacitymanagerorganizationsaccess)
  - [UpdateInterruptibleCapacityReservationAllocation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#updateinterruptiblecapacityreservationallocation)
  - [UpdateSecurityGroupRuleDescriptionsEgress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#updatesecuritygroupruledescriptionsegress)
  - [UpdateSecurityGroupRuleDescriptionsIngress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#updatesecuritygroupruledescriptionsingress)
  - [WithdrawByoipCidr](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-11-15.html#withdrawbyoipcidr)
- [**2016-09-15**](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html)

  - [AcceptReservedInstancesExchangeQuote](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#acceptreservedinstancesexchangequote)
  - [AcceptVpcPeeringConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#acceptvpcpeeringconnection)
  - [AllocateAddress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#allocateaddress)
  - [AllocateHosts](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#allocatehosts)
  - [AssignPrivateIpAddresses](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#assignprivateipaddresses)
  - [AssociateAddress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#associateaddress)
  - [AssociateDhcpOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#associatedhcpoptions)
  - [AssociateRouteTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#associateroutetable)
  - [AttachClassicLinkVpc](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#attachclassiclinkvpc)
  - [AttachInternetGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#attachinternetgateway)
  - [AttachNetworkInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#attachnetworkinterface)
  - [AttachVolume](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#attachvolume)
  - [AttachVpnGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#attachvpngateway)
  - [AuthorizeSecurityGroupEgress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#authorizesecuritygroupegress)
  - [AuthorizeSecurityGroupIngress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#authorizesecuritygroupingress)
  - [BundleInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#bundleinstance)
  - [CancelBundleTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#cancelbundletask)
  - [CancelConversionTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#cancelconversiontask)
  - [CancelExportTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#cancelexporttask)
  - [CancelImportTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#cancelimporttask)
  - [CancelReservedInstancesListing](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#cancelreservedinstanceslisting)
  - [CancelSpotFleetRequests](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#cancelspotfleetrequests)
  - [CancelSpotInstanceRequests](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#cancelspotinstancerequests)
  - [ConfirmProductInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#confirmproductinstance)
  - [CopyImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#copyimage)
  - [CopySnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#copysnapshot)
  - [CreateCustomerGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#createcustomergateway)
  - [CreateDhcpOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#createdhcpoptions)
  - [CreateFlowLogs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#createflowlogs)
  - [CreateImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#createimage)
  - [CreateInstanceExportTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#createinstanceexporttask)
  - [CreateInternetGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#createinternetgateway)
  - [CreateKeyPair](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#createkeypair)
  - [CreateNatGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#createnatgateway)
  - [CreateNetworkAcl](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#createnetworkacl)
  - [CreateNetworkAclEntry](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#createnetworkaclentry)
  - [CreateNetworkInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#createnetworkinterface)
  - [CreatePlacementGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#createplacementgroup)
  - [CreateReservedInstancesListing](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#createreservedinstanceslisting)
  - [CreateRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#createroute)
  - [CreateRouteTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#createroutetable)
  - [CreateSecurityGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#createsecuritygroup)
  - [CreateSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#createsnapshot)
  - [CreateSpotDatafeedSubscription](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#createspotdatafeedsubscription)
  - [CreateSubnet](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#createsubnet)
  - [CreateTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#createtags)
  - [CreateVolume](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#createvolume)
  - [CreateVpc](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#createvpc)
  - [CreateVpcEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#createvpcendpoint)
  - [CreateVpcPeeringConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#createvpcpeeringconnection)
  - [CreateVpnConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#createvpnconnection)
  - [CreateVpnConnectionRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#createvpnconnectionroute)
  - [CreateVpnGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#createvpngateway)
  - [DeleteCustomerGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#deletecustomergateway)
  - [DeleteDhcpOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#deletedhcpoptions)
  - [DeleteFlowLogs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#deleteflowlogs)
  - [DeleteInternetGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#deleteinternetgateway)
  - [DeleteKeyPair](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#deletekeypair)
  - [DeleteNatGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#deletenatgateway)
  - [DeleteNetworkAcl](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#deletenetworkacl)
  - [DeleteNetworkAclEntry](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#deletenetworkaclentry)
  - [DeleteNetworkInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#deletenetworkinterface)
  - [DeletePlacementGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#deleteplacementgroup)
  - [DeleteRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#deleteroute)
  - [DeleteRouteTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#deleteroutetable)
  - [DeleteSecurityGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#deletesecuritygroup)
  - [DeleteSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#deletesnapshot)
  - [DeleteSpotDatafeedSubscription](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#deletespotdatafeedsubscription)
  - [DeleteSubnet](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#deletesubnet)
  - [DeleteTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#deletetags)
  - [DeleteVolume](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#deletevolume)
  - [DeleteVpc](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#deletevpc)
  - [DeleteVpcEndpoints](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#deletevpcendpoints)
  - [DeleteVpcPeeringConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#deletevpcpeeringconnection)
  - [DeleteVpnConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#deletevpnconnection)
  - [DeleteVpnConnectionRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#deletevpnconnectionroute)
  - [DeleteVpnGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#deletevpngateway)
  - [DeregisterImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#deregisterimage)
  - [DescribeAccountAttributes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describeaccountattributes)
  - [DescribeAddresses](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describeaddresses)
  - [DescribeAvailabilityZones](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describeavailabilityzones)
  - [DescribeBundleTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describebundletasks)
  - [DescribeClassicLinkInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describeclassiclinkinstances)
  - [DescribeConversionTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describeconversiontasks)
  - [DescribeCustomerGateways](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describecustomergateways)
  - [DescribeDhcpOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describedhcpoptions)
  - [DescribeExportTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describeexporttasks)
  - [DescribeFlowLogs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describeflowlogs)
  - [DescribeHostReservationOfferings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describehostreservationofferings)
  - [DescribeHostReservations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describehostreservations)
  - [DescribeHosts](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describehosts)
  - [DescribeIdFormat](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describeidformat)
  - [DescribeIdentityIdFormat](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describeidentityidformat)
  - [DescribeImageAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describeimageattribute)
  - [DescribeImages](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describeimages)
  - [DescribeImportImageTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describeimportimagetasks)
  - [DescribeImportSnapshotTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describeimportsnapshottasks)
  - [DescribeInstanceAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describeinstanceattribute)
  - [DescribeInstanceStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describeinstancestatus)
  - [DescribeInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describeinstances)
  - [DescribeInternetGateways](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describeinternetgateways)
  - [DescribeKeyPairs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describekeypairs)
  - [DescribeMovingAddresses](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describemovingaddresses)
  - [DescribeNatGateways](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describenatgateways)
  - [DescribeNetworkAcls](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describenetworkacls)
  - [DescribeNetworkInterfaceAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describenetworkinterfaceattribute)
  - [DescribeNetworkInterfaces](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describenetworkinterfaces)
  - [DescribePlacementGroups](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describeplacementgroups)
  - [DescribePrefixLists](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describeprefixlists)
  - [DescribeRegions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describeregions)
  - [DescribeReservedInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describereservedinstances)
  - [DescribeReservedInstancesListings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describereservedinstanceslistings)
  - [DescribeReservedInstancesModifications](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describereservedinstancesmodifications)
  - [DescribeReservedInstancesOfferings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describereservedinstancesofferings)
  - [DescribeRouteTables](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describeroutetables)
  - [DescribeScheduledInstanceAvailability](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describescheduledinstanceavailability)
  - [DescribeScheduledInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describescheduledinstances)
  - [DescribeSecurityGroupReferences](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describesecuritygroupreferences)
  - [DescribeSecurityGroups](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describesecuritygroups)
  - [DescribeSnapshotAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describesnapshotattribute)
  - [DescribeSnapshots](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describesnapshots)
  - [DescribeSpotDatafeedSubscription](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describespotdatafeedsubscription)
  - [DescribeSpotFleetInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describespotfleetinstances)
  - [DescribeSpotFleetRequestHistory](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describespotfleetrequesthistory)
  - [DescribeSpotFleetRequests](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describespotfleetrequests)
  - [DescribeSpotInstanceRequests](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describespotinstancerequests)
  - [DescribeSpotPriceHistory](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describespotpricehistory)
  - [DescribeStaleSecurityGroups](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describestalesecuritygroups)
  - [DescribeSubnets](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describesubnets)
  - [DescribeTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describetags)
  - [DescribeVolumeAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describevolumeattribute)
  - [DescribeVolumeStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describevolumestatus)
  - [DescribeVolumes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describevolumes)
  - [DescribeVpcAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describevpcattribute)
  - [DescribeVpcClassicLink](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describevpcclassiclink)
  - [DescribeVpcClassicLinkDnsSupport](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describevpcclassiclinkdnssupport)
  - [DescribeVpcEndpointServices](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describevpcendpointservices)
  - [DescribeVpcEndpoints](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describevpcendpoints)
  - [DescribeVpcPeeringConnections](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describevpcpeeringconnections)
  - [DescribeVpcs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describevpcs)
  - [DescribeVpnConnections](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describevpnconnections)
  - [DescribeVpnGateways](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#describevpngateways)
  - [DetachClassicLinkVpc](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#detachclassiclinkvpc)
  - [DetachInternetGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#detachinternetgateway)
  - [DetachNetworkInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#detachnetworkinterface)
  - [DetachVolume](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#detachvolume)
  - [DetachVpnGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#detachvpngateway)
  - [DisableVgwRoutePropagation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#disablevgwroutepropagation)
  - [DisableVpcClassicLink](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#disablevpcclassiclink)
  - [DisableVpcClassicLinkDnsSupport](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#disablevpcclassiclinkdnssupport)
  - [DisassociateAddress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#disassociateaddress)
  - [DisassociateRouteTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#disassociateroutetable)
  - [EnableVgwRoutePropagation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#enablevgwroutepropagation)
  - [EnableVolumeIO](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#enablevolumeio)
  - [EnableVpcClassicLink](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#enablevpcclassiclink)
  - [EnableVpcClassicLinkDnsSupport](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#enablevpcclassiclinkdnssupport)
  - [GetConsoleOutput](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#getconsoleoutput)
  - [GetConsoleScreenshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#getconsolescreenshot)
  - [GetHostReservationPurchasePreview](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#gethostreservationpurchasepreview)
  - [GetPasswordData](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#getpassworddata)
  - [GetReservedInstancesExchangeQuote](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#getreservedinstancesexchangequote)
  - [ImportImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#importimage)
  - [ImportInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#importinstance)
  - [ImportKeyPair](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#importkeypair)
  - [ImportSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#importsnapshot)
  - [ImportVolume](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#importvolume)
  - [ModifyHosts](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#modifyhosts)
  - [ModifyIdFormat](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#modifyidformat)
  - [ModifyIdentityIdFormat](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#modifyidentityidformat)
  - [ModifyImageAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#modifyimageattribute)
  - [ModifyInstanceAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#modifyinstanceattribute)
  - [ModifyInstancePlacement](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#modifyinstanceplacement)
  - [ModifyNetworkInterfaceAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#modifynetworkinterfaceattribute)
  - [ModifyReservedInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#modifyreservedinstances)
  - [ModifySnapshotAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#modifysnapshotattribute)
  - [ModifySpotFleetRequest](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#modifyspotfleetrequest)
  - [ModifySubnetAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#modifysubnetattribute)
  - [ModifyVolumeAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#modifyvolumeattribute)
  - [ModifyVpcAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#modifyvpcattribute)
  - [ModifyVpcEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#modifyvpcendpoint)
  - [ModifyVpcPeeringConnectionOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#modifyvpcpeeringconnectionoptions)
  - [MonitorInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#monitorinstances)
  - [MoveAddressToVpc](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#moveaddresstovpc)
  - [PurchaseHostReservation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#purchasehostreservation)
  - [PurchaseReservedInstancesOffering](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#purchasereservedinstancesoffering)
  - [PurchaseScheduledInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#purchasescheduledinstances)
  - [RebootInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#rebootinstances)
  - [RegisterImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#registerimage)
  - [RejectVpcPeeringConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#rejectvpcpeeringconnection)
  - [ReleaseAddress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#releaseaddress)
  - [ReleaseHosts](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#releasehosts)
  - [ReplaceNetworkAclAssociation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#replacenetworkaclassociation)
  - [ReplaceNetworkAclEntry](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#replacenetworkaclentry)
  - [ReplaceRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#replaceroute)
  - [ReplaceRouteTableAssociation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#replaceroutetableassociation)
  - [ReportInstanceStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#reportinstancestatus)
  - [RequestSpotFleet](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#requestspotfleet)
  - [RequestSpotInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#requestspotinstances)
  - [ResetImageAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#resetimageattribute)
  - [ResetInstanceAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#resetinstanceattribute)
  - [ResetNetworkInterfaceAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#resetnetworkinterfaceattribute)
  - [ResetSnapshotAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#resetsnapshotattribute)
  - [RestoreAddressToClassic](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#restoreaddresstoclassic)
  - [RevokeSecurityGroupEgress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#revokesecuritygroupegress)
  - [RevokeSecurityGroupIngress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#revokesecuritygroupingress)
  - [RunInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#runinstances)
  - [RunScheduledInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#runscheduledinstances)
  - [StartInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#startinstances)
  - [StopInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#stopinstances)
  - [TerminateInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#terminateinstances)
  - [UnassignPrivateIpAddresses](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#unassignprivateipaddresses)
  - [UnmonitorInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-09-15.html#unmonitorinstances)
- [**2016-04-01**](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html)

  - [AcceptVpcPeeringConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#acceptvpcpeeringconnection)
  - [AllocateAddress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#allocateaddress)
  - [AllocateHosts](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#allocatehosts)
  - [AssignPrivateIpAddresses](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#assignprivateipaddresses)
  - [AssociateAddress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#associateaddress)
  - [AssociateDhcpOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#associatedhcpoptions)
  - [AssociateRouteTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#associateroutetable)
  - [AttachClassicLinkVpc](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#attachclassiclinkvpc)
  - [AttachInternetGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#attachinternetgateway)
  - [AttachNetworkInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#attachnetworkinterface)
  - [AttachVolume](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#attachvolume)
  - [AttachVpnGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#attachvpngateway)
  - [AuthorizeSecurityGroupEgress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#authorizesecuritygroupegress)
  - [AuthorizeSecurityGroupIngress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#authorizesecuritygroupingress)
  - [BundleInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#bundleinstance)
  - [CancelBundleTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#cancelbundletask)
  - [CancelConversionTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#cancelconversiontask)
  - [CancelExportTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#cancelexporttask)
  - [CancelImportTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#cancelimporttask)
  - [CancelReservedInstancesListing](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#cancelreservedinstanceslisting)
  - [CancelSpotFleetRequests](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#cancelspotfleetrequests)
  - [CancelSpotInstanceRequests](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#cancelspotinstancerequests)
  - [ConfirmProductInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#confirmproductinstance)
  - [CopyImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#copyimage)
  - [CopySnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#copysnapshot)
  - [CreateCustomerGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#createcustomergateway)
  - [CreateDhcpOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#createdhcpoptions)
  - [CreateFlowLogs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#createflowlogs)
  - [CreateImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#createimage)
  - [CreateInstanceExportTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#createinstanceexporttask)
  - [CreateInternetGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#createinternetgateway)
  - [CreateKeyPair](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#createkeypair)
  - [CreateNatGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#createnatgateway)
  - [CreateNetworkAcl](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#createnetworkacl)
  - [CreateNetworkAclEntry](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#createnetworkaclentry)
  - [CreateNetworkInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#createnetworkinterface)
  - [CreatePlacementGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#createplacementgroup)
  - [CreateReservedInstancesListing](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#createreservedinstanceslisting)
  - [CreateRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#createroute)
  - [CreateRouteTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#createroutetable)
  - [CreateSecurityGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#createsecuritygroup)
  - [CreateSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#createsnapshot)
  - [CreateSpotDatafeedSubscription](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#createspotdatafeedsubscription)
  - [CreateSubnet](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#createsubnet)
  - [CreateTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#createtags)
  - [CreateVolume](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#createvolume)
  - [CreateVpc](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#createvpc)
  - [CreateVpcEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#createvpcendpoint)
  - [CreateVpcPeeringConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#createvpcpeeringconnection)
  - [CreateVpnConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#createvpnconnection)
  - [CreateVpnConnectionRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#createvpnconnectionroute)
  - [CreateVpnGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#createvpngateway)
  - [DeleteCustomerGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#deletecustomergateway)
  - [DeleteDhcpOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#deletedhcpoptions)
  - [DeleteFlowLogs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#deleteflowlogs)
  - [DeleteInternetGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#deleteinternetgateway)
  - [DeleteKeyPair](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#deletekeypair)
  - [DeleteNatGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#deletenatgateway)
  - [DeleteNetworkAcl](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#deletenetworkacl)
  - [DeleteNetworkAclEntry](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#deletenetworkaclentry)
  - [DeleteNetworkInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#deletenetworkinterface)
  - [DeletePlacementGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#deleteplacementgroup)
  - [DeleteRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#deleteroute)
  - [DeleteRouteTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#deleteroutetable)
  - [DeleteSecurityGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#deletesecuritygroup)
  - [DeleteSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#deletesnapshot)
  - [DeleteSpotDatafeedSubscription](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#deletespotdatafeedsubscription)
  - [DeleteSubnet](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#deletesubnet)
  - [DeleteTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#deletetags)
  - [DeleteVolume](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#deletevolume)
  - [DeleteVpc](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#deletevpc)
  - [DeleteVpcEndpoints](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#deletevpcendpoints)
  - [DeleteVpcPeeringConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#deletevpcpeeringconnection)
  - [DeleteVpnConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#deletevpnconnection)
  - [DeleteVpnConnectionRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#deletevpnconnectionroute)
  - [DeleteVpnGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#deletevpngateway)
  - [DeregisterImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#deregisterimage)
  - [DescribeAccountAttributes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describeaccountattributes)
  - [DescribeAddresses](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describeaddresses)
  - [DescribeAvailabilityZones](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describeavailabilityzones)
  - [DescribeBundleTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describebundletasks)
  - [DescribeClassicLinkInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describeclassiclinkinstances)
  - [DescribeConversionTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describeconversiontasks)
  - [DescribeCustomerGateways](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describecustomergateways)
  - [DescribeDhcpOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describedhcpoptions)
  - [DescribeExportTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describeexporttasks)
  - [DescribeFlowLogs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describeflowlogs)
  - [DescribeHostReservationOfferings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describehostreservationofferings)
  - [DescribeHostReservations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describehostreservations)
  - [DescribeHosts](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describehosts)
  - [DescribeIdFormat](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describeidformat)
  - [DescribeIdentityIdFormat](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describeidentityidformat)
  - [DescribeImageAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describeimageattribute)
  - [DescribeImages](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describeimages)
  - [DescribeImportImageTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describeimportimagetasks)
  - [DescribeImportSnapshotTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describeimportsnapshottasks)
  - [DescribeInstanceAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describeinstanceattribute)
  - [DescribeInstanceStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describeinstancestatus)
  - [DescribeInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describeinstances)
  - [DescribeInternetGateways](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describeinternetgateways)
  - [DescribeKeyPairs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describekeypairs)
  - [DescribeMovingAddresses](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describemovingaddresses)
  - [DescribeNatGateways](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describenatgateways)
  - [DescribeNetworkAcls](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describenetworkacls)
  - [DescribeNetworkInterfaceAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describenetworkinterfaceattribute)
  - [DescribeNetworkInterfaces](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describenetworkinterfaces)
  - [DescribePlacementGroups](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describeplacementgroups)
  - [DescribePrefixLists](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describeprefixlists)
  - [DescribeRegions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describeregions)
  - [DescribeReservedInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describereservedinstances)
  - [DescribeReservedInstancesListings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describereservedinstanceslistings)
  - [DescribeReservedInstancesModifications](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describereservedinstancesmodifications)
  - [DescribeReservedInstancesOfferings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describereservedinstancesofferings)
  - [DescribeRouteTables](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describeroutetables)
  - [DescribeScheduledInstanceAvailability](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describescheduledinstanceavailability)
  - [DescribeScheduledInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describescheduledinstances)
  - [DescribeSecurityGroupReferences](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describesecuritygroupreferences)
  - [DescribeSecurityGroups](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describesecuritygroups)
  - [DescribeSnapshotAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describesnapshotattribute)
  - [DescribeSnapshots](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describesnapshots)
  - [DescribeSpotDatafeedSubscription](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describespotdatafeedsubscription)
  - [DescribeSpotFleetInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describespotfleetinstances)
  - [DescribeSpotFleetRequestHistory](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describespotfleetrequesthistory)
  - [DescribeSpotFleetRequests](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describespotfleetrequests)
  - [DescribeSpotInstanceRequests](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describespotinstancerequests)
  - [DescribeSpotPriceHistory](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describespotpricehistory)
  - [DescribeStaleSecurityGroups](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describestalesecuritygroups)
  - [DescribeSubnets](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describesubnets)
  - [DescribeTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describetags)
  - [DescribeVolumeAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describevolumeattribute)
  - [DescribeVolumeStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describevolumestatus)
  - [DescribeVolumes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describevolumes)
  - [DescribeVpcAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describevpcattribute)
  - [DescribeVpcClassicLink](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describevpcclassiclink)
  - [DescribeVpcClassicLinkDnsSupport](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describevpcclassiclinkdnssupport)
  - [DescribeVpcEndpointServices](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describevpcendpointservices)
  - [DescribeVpcEndpoints](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describevpcendpoints)
  - [DescribeVpcPeeringConnections](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describevpcpeeringconnections)
  - [DescribeVpcs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describevpcs)
  - [DescribeVpnConnections](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describevpnconnections)
  - [DescribeVpnGateways](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#describevpngateways)
  - [DetachClassicLinkVpc](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#detachclassiclinkvpc)
  - [DetachInternetGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#detachinternetgateway)
  - [DetachNetworkInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#detachnetworkinterface)
  - [DetachVolume](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#detachvolume)
  - [DetachVpnGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#detachvpngateway)
  - [DisableVgwRoutePropagation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#disablevgwroutepropagation)
  - [DisableVpcClassicLink](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#disablevpcclassiclink)
  - [DisableVpcClassicLinkDnsSupport](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#disablevpcclassiclinkdnssupport)
  - [DisassociateAddress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#disassociateaddress)
  - [DisassociateRouteTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#disassociateroutetable)
  - [EnableVgwRoutePropagation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#enablevgwroutepropagation)
  - [EnableVolumeIO](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#enablevolumeio)
  - [EnableVpcClassicLink](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#enablevpcclassiclink)
  - [EnableVpcClassicLinkDnsSupport](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#enablevpcclassiclinkdnssupport)
  - [GetConsoleOutput](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#getconsoleoutput)
  - [GetConsoleScreenshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#getconsolescreenshot)
  - [GetHostReservationPurchasePreview](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#gethostreservationpurchasepreview)
  - [GetPasswordData](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#getpassworddata)
  - [ImportImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#importimage)
  - [ImportInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#importinstance)
  - [ImportKeyPair](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#importkeypair)
  - [ImportSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#importsnapshot)
  - [ImportVolume](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#importvolume)
  - [ModifyHosts](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#modifyhosts)
  - [ModifyIdFormat](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#modifyidformat)
  - [ModifyIdentityIdFormat](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#modifyidentityidformat)
  - [ModifyImageAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#modifyimageattribute)
  - [ModifyInstanceAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#modifyinstanceattribute)
  - [ModifyInstancePlacement](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#modifyinstanceplacement)
  - [ModifyNetworkInterfaceAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#modifynetworkinterfaceattribute)
  - [ModifyReservedInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#modifyreservedinstances)
  - [ModifySnapshotAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#modifysnapshotattribute)
  - [ModifySpotFleetRequest](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#modifyspotfleetrequest)
  - [ModifySubnetAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#modifysubnetattribute)
  - [ModifyVolumeAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#modifyvolumeattribute)
  - [ModifyVpcAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#modifyvpcattribute)
  - [ModifyVpcEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#modifyvpcendpoint)
  - [ModifyVpcPeeringConnectionOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#modifyvpcpeeringconnectionoptions)
  - [MonitorInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#monitorinstances)
  - [MoveAddressToVpc](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#moveaddresstovpc)
  - [PurchaseHostReservation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#purchasehostreservation)
  - [PurchaseReservedInstancesOffering](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#purchasereservedinstancesoffering)
  - [PurchaseScheduledInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#purchasescheduledinstances)
  - [RebootInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#rebootinstances)
  - [RegisterImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#registerimage)
  - [RejectVpcPeeringConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#rejectvpcpeeringconnection)
  - [ReleaseAddress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#releaseaddress)
  - [ReleaseHosts](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#releasehosts)
  - [ReplaceNetworkAclAssociation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#replacenetworkaclassociation)
  - [ReplaceNetworkAclEntry](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#replacenetworkaclentry)
  - [ReplaceRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#replaceroute)
  - [ReplaceRouteTableAssociation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#replaceroutetableassociation)
  - [ReportInstanceStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#reportinstancestatus)
  - [RequestSpotFleet](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#requestspotfleet)
  - [RequestSpotInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#requestspotinstances)
  - [ResetImageAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#resetimageattribute)
  - [ResetInstanceAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#resetinstanceattribute)
  - [ResetNetworkInterfaceAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#resetnetworkinterfaceattribute)
  - [ResetSnapshotAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#resetsnapshotattribute)
  - [RestoreAddressToClassic](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#restoreaddresstoclassic)
  - [RevokeSecurityGroupEgress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#revokesecuritygroupegress)
  - [RevokeSecurityGroupIngress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#revokesecuritygroupingress)
  - [RunInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#runinstances)
  - [RunScheduledInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#runscheduledinstances)
  - [StartInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#startinstances)
  - [StopInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#stopinstances)
  - [TerminateInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#terminateinstances)
  - [UnassignPrivateIpAddresses](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#unassignprivateipaddresses)
  - [UnmonitorInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2016-04-01.html#unmonitorinstances)
- [**2015-10-01**](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html)

  - [AcceptVpcPeeringConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#acceptvpcpeeringconnection)
  - [AllocateAddress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#allocateaddress)
  - [AllocateHosts](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#allocatehosts)
  - [AssignPrivateIpAddresses](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#assignprivateipaddresses)
  - [AssociateAddress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#associateaddress)
  - [AssociateDhcpOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#associatedhcpoptions)
  - [AssociateRouteTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#associateroutetable)
  - [AttachClassicLinkVpc](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#attachclassiclinkvpc)
  - [AttachInternetGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#attachinternetgateway)
  - [AttachNetworkInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#attachnetworkinterface)
  - [AttachVolume](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#attachvolume)
  - [AttachVpnGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#attachvpngateway)
  - [AuthorizeSecurityGroupEgress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#authorizesecuritygroupegress)
  - [AuthorizeSecurityGroupIngress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#authorizesecuritygroupingress)
  - [BundleInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#bundleinstance)
  - [CancelBundleTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#cancelbundletask)
  - [CancelConversionTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#cancelconversiontask)
  - [CancelExportTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#cancelexporttask)
  - [CancelImportTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#cancelimporttask)
  - [CancelReservedInstancesListing](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#cancelreservedinstanceslisting)
  - [CancelSpotFleetRequests](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#cancelspotfleetrequests)
  - [CancelSpotInstanceRequests](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#cancelspotinstancerequests)
  - [ConfirmProductInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#confirmproductinstance)
  - [CopyImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#copyimage)
  - [CopySnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#copysnapshot)
  - [CreateCustomerGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#createcustomergateway)
  - [CreateDhcpOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#createdhcpoptions)
  - [CreateFlowLogs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#createflowlogs)
  - [CreateImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#createimage)
  - [CreateInstanceExportTask](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#createinstanceexporttask)
  - [CreateInternetGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#createinternetgateway)
  - [CreateKeyPair](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#createkeypair)
  - [CreateNatGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#createnatgateway)
  - [CreateNetworkAcl](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#createnetworkacl)
  - [CreateNetworkAclEntry](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#createnetworkaclentry)
  - [CreateNetworkInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#createnetworkinterface)
  - [CreatePlacementGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#createplacementgroup)
  - [CreateReservedInstancesListing](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#createreservedinstanceslisting)
  - [CreateRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#createroute)
  - [CreateRouteTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#createroutetable)
  - [CreateSecurityGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#createsecuritygroup)
  - [CreateSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#createsnapshot)
  - [CreateSpotDatafeedSubscription](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#createspotdatafeedsubscription)
  - [CreateSubnet](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#createsubnet)
  - [CreateTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#createtags)
  - [CreateVolume](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#createvolume)
  - [CreateVpc](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#createvpc)
  - [CreateVpcEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#createvpcendpoint)
  - [CreateVpcPeeringConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#createvpcpeeringconnection)
  - [CreateVpnConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#createvpnconnection)
  - [CreateVpnConnectionRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#createvpnconnectionroute)
  - [CreateVpnGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#createvpngateway)
  - [DeleteCustomerGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#deletecustomergateway)
  - [DeleteDhcpOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#deletedhcpoptions)
  - [DeleteFlowLogs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#deleteflowlogs)
  - [DeleteInternetGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#deleteinternetgateway)
  - [DeleteKeyPair](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#deletekeypair)
  - [DeleteNatGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#deletenatgateway)
  - [DeleteNetworkAcl](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#deletenetworkacl)
  - [DeleteNetworkAclEntry](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#deletenetworkaclentry)
  - [DeleteNetworkInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#deletenetworkinterface)
  - [DeletePlacementGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#deleteplacementgroup)
  - [DeleteRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#deleteroute)
  - [DeleteRouteTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#deleteroutetable)
  - [DeleteSecurityGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#deletesecuritygroup)
  - [DeleteSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#deletesnapshot)
  - [DeleteSpotDatafeedSubscription](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#deletespotdatafeedsubscription)
  - [DeleteSubnet](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#deletesubnet)
  - [DeleteTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#deletetags)
  - [DeleteVolume](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#deletevolume)
  - [DeleteVpc](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#deletevpc)
  - [DeleteVpcEndpoints](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#deletevpcendpoints)
  - [DeleteVpcPeeringConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#deletevpcpeeringconnection)
  - [DeleteVpnConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#deletevpnconnection)
  - [DeleteVpnConnectionRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#deletevpnconnectionroute)
  - [DeleteVpnGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#deletevpngateway)
  - [DeregisterImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#deregisterimage)
  - [DescribeAccountAttributes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describeaccountattributes)
  - [DescribeAddresses](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describeaddresses)
  - [DescribeAvailabilityZones](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describeavailabilityzones)
  - [DescribeBundleTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describebundletasks)
  - [DescribeClassicLinkInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describeclassiclinkinstances)
  - [DescribeConversionTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describeconversiontasks)
  - [DescribeCustomerGateways](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describecustomergateways)
  - [DescribeDhcpOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describedhcpoptions)
  - [DescribeExportTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describeexporttasks)
  - [DescribeFlowLogs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describeflowlogs)
  - [DescribeHosts](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describehosts)
  - [DescribeIdFormat](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describeidformat)
  - [DescribeImageAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describeimageattribute)
  - [DescribeImages](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describeimages)
  - [DescribeImportImageTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describeimportimagetasks)
  - [DescribeImportSnapshotTasks](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describeimportsnapshottasks)
  - [DescribeInstanceAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describeinstanceattribute)
  - [DescribeInstanceStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describeinstancestatus)
  - [DescribeInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describeinstances)
  - [DescribeInternetGateways](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describeinternetgateways)
  - [DescribeKeyPairs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describekeypairs)
  - [DescribeMovingAddresses](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describemovingaddresses)
  - [DescribeNatGateways](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describenatgateways)
  - [DescribeNetworkAcls](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describenetworkacls)
  - [DescribeNetworkInterfaceAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describenetworkinterfaceattribute)
  - [DescribeNetworkInterfaces](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describenetworkinterfaces)
  - [DescribePlacementGroups](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describeplacementgroups)
  - [DescribePrefixLists](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describeprefixlists)
  - [DescribeRegions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describeregions)
  - [DescribeReservedInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describereservedinstances)
  - [DescribeReservedInstancesListings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describereservedinstanceslistings)
  - [DescribeReservedInstancesModifications](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describereservedinstancesmodifications)
  - [DescribeReservedInstancesOfferings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describereservedinstancesofferings)
  - [DescribeRouteTables](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describeroutetables)
  - [DescribeScheduledInstanceAvailability](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describescheduledinstanceavailability)
  - [DescribeScheduledInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describescheduledinstances)
  - [DescribeSecurityGroupReferences](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describesecuritygroupreferences)
  - [DescribeSecurityGroups](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describesecuritygroups)
  - [DescribeSnapshotAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describesnapshotattribute)
  - [DescribeSnapshots](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describesnapshots)
  - [DescribeSpotDatafeedSubscription](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describespotdatafeedsubscription)
  - [DescribeSpotFleetInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describespotfleetinstances)
  - [DescribeSpotFleetRequestHistory](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describespotfleetrequesthistory)
  - [DescribeSpotFleetRequests](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describespotfleetrequests)
  - [DescribeSpotInstanceRequests](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describespotinstancerequests)
  - [DescribeSpotPriceHistory](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describespotpricehistory)
  - [DescribeStaleSecurityGroups](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describestalesecuritygroups)
  - [DescribeSubnets](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describesubnets)
  - [DescribeTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describetags)
  - [DescribeVolumeAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describevolumeattribute)
  - [DescribeVolumeStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describevolumestatus)
  - [DescribeVolumes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describevolumes)
  - [DescribeVpcAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describevpcattribute)
  - [DescribeVpcClassicLink](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describevpcclassiclink)
  - [DescribeVpcClassicLinkDnsSupport](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describevpcclassiclinkdnssupport)
  - [DescribeVpcEndpointServices](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describevpcendpointservices)
  - [DescribeVpcEndpoints](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describevpcendpoints)
  - [DescribeVpcPeeringConnections](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describevpcpeeringconnections)
  - [DescribeVpcs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describevpcs)
  - [DescribeVpnConnections](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describevpnconnections)
  - [DescribeVpnGateways](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#describevpngateways)
  - [DetachClassicLinkVpc](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#detachclassiclinkvpc)
  - [DetachInternetGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#detachinternetgateway)
  - [DetachNetworkInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#detachnetworkinterface)
  - [DetachVolume](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#detachvolume)
  - [DetachVpnGateway](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#detachvpngateway)
  - [DisableVgwRoutePropagation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#disablevgwroutepropagation)
  - [DisableVpcClassicLink](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#disablevpcclassiclink)
  - [DisableVpcClassicLinkDnsSupport](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#disablevpcclassiclinkdnssupport)
  - [DisassociateAddress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#disassociateaddress)
  - [DisassociateRouteTable](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#disassociateroutetable)
  - [EnableVgwRoutePropagation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#enablevgwroutepropagation)
  - [EnableVolumeIO](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#enablevolumeio)
  - [EnableVpcClassicLink](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#enablevpcclassiclink)
  - [EnableVpcClassicLinkDnsSupport](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#enablevpcclassiclinkdnssupport)
  - [GetConsoleOutput](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#getconsoleoutput)
  - [GetConsoleScreenshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#getconsolescreenshot)
  - [GetPasswordData](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#getpassworddata)
  - [ImportImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#importimage)
  - [ImportInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#importinstance)
  - [ImportKeyPair](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#importkeypair)
  - [ImportSnapshot](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#importsnapshot)
  - [ImportVolume](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#importvolume)
  - [ModifyHosts](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#modifyhosts)
  - [ModifyIdFormat](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#modifyidformat)
  - [ModifyImageAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#modifyimageattribute)
  - [ModifyInstanceAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#modifyinstanceattribute)
  - [ModifyInstancePlacement](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#modifyinstanceplacement)
  - [ModifyNetworkInterfaceAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#modifynetworkinterfaceattribute)
  - [ModifyReservedInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#modifyreservedinstances)
  - [ModifySnapshotAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#modifysnapshotattribute)
  - [ModifySpotFleetRequest](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#modifyspotfleetrequest)
  - [ModifySubnetAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#modifysubnetattribute)
  - [ModifyVolumeAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#modifyvolumeattribute)
  - [ModifyVpcAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#modifyvpcattribute)
  - [ModifyVpcEndpoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#modifyvpcendpoint)
  - [ModifyVpcPeeringConnectionOptions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#modifyvpcpeeringconnectionoptions)
  - [MonitorInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#monitorinstances)
  - [MoveAddressToVpc](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#moveaddresstovpc)
  - [PurchaseReservedInstancesOffering](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#purchasereservedinstancesoffering)
  - [PurchaseScheduledInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#purchasescheduledinstances)
  - [RebootInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#rebootinstances)
  - [RegisterImage](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#registerimage)
  - [RejectVpcPeeringConnection](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#rejectvpcpeeringconnection)
  - [ReleaseAddress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#releaseaddress)
  - [ReleaseHosts](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#releasehosts)
  - [ReplaceNetworkAclAssociation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#replacenetworkaclassociation)
  - [ReplaceNetworkAclEntry](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#replacenetworkaclentry)
  - [ReplaceRoute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#replaceroute)
  - [ReplaceRouteTableAssociation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#replaceroutetableassociation)
  - [ReportInstanceStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#reportinstancestatus)
  - [RequestSpotFleet](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#requestspotfleet)
  - [RequestSpotInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#requestspotinstances)
  - [ResetImageAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#resetimageattribute)
  - [ResetInstanceAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#resetinstanceattribute)
  - [ResetNetworkInterfaceAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#resetnetworkinterfaceattribute)
  - [ResetSnapshotAttribute](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#resetsnapshotattribute)
  - [RestoreAddressToClassic](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#restoreaddresstoclassic)
  - [RevokeSecurityGroupEgress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#revokesecuritygroupegress)
  - [RevokeSecurityGroupIngress](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#revokesecuritygroupingress)
  - [RunInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#runinstances)
  - [RunScheduledInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#runscheduledinstances)
  - [StartInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#startinstances)
  - [StopInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#stopinstances)
  - [TerminateInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#terminateinstances)
  - [UnassignPrivateIpAddresses](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#unassignprivateipaddresses)
  - [UnmonitorInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-ec2-2015-10-01.html#unmonitorinstances)

## Examples

### Basics, Actions and Scenarios

The following code examples show you how to perform actions and implement common scenarios by using the AWS SDK for PHP with Amazon Elastic Compute Cloud.

- [See examples on AWS Docs](https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/php_ec2_code_examples.html)

### Legacy Code Examples With Guidance

The following examples demonstrate how to use this service with the AWS SDK for PHP. These code examples are available in the [AWS SDK for PHP Developer Guide](https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/ec2-examples.html).

- [Managing Amazon EC2 instances](https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/ec2-examples-managing-instances.html)
- [Using Elastic IP addresses](https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/ec2-examples-using-elastic-ip-addresses.html)
- [Using regions and availability zones](https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/ec2-examples-using-regions-and-zones.html)
- [Working with key pairs](https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/ec2-examples-working-with-key-pairs.html)
- [Working with security groups](https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/ec2-examples-using-security-groups.html)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Ec2.Ec2Client.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Ec2.Ec2Client.html\#toc-methods)

[\_\_call()](class-aws-awsclienttrait.md#method___call)
: mixed [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Ec2.Ec2Client.html#method___construct)
: mixed The client constructor accepts the following options:[\_\_sleep()](class-aws-awsclient.md#method___sleep)
: mixed [execute()](class-aws-awsclienttrait.md#method_execute)
: mixed [executeAsync()](class-aws-awsclienttrait.md#method_executeAsync)
: mixed [factory()](class-aws-awsclient.md#method_factory)
: static [getApi()](class-aws-awsclienttrait.md#method_getApi)
: [Service](class-aws-api-service.md)[getArguments()](class-aws-awsclient.md#method_getArguments)
: array<string\|int, mixed> Get an array of client constructor arguments used by the client.[getClientBuiltIns()](class-aws-awsclient.md#method_getClientBuiltIns)
: array<string\|int, mixed> Provides the set of built-in keys and values
used for endpoint resolution[getClientContextParams()](class-aws-awsclient.md#method_getClientContextParams)
: array<string\|int, mixed> Provides the set of service context parameter
key-value pairs used for endpoint resolution.[getCommand()](class-aws-awsclienttrait.md#method_getCommand)
: [CommandInterface](class-aws-commandinterface.md)[getConfig()](class-aws-awsclient.md#method_getConfig)
: mixed\|null Get a client configuration value.[getCredentials()](class-aws-awsclient.md#method_getCredentials)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.[getEndpoint()](class-aws-awsclient.md#method_getEndpoint)
: [UriInterface](class-psr-http-message-uriinterface.md)Gets the default endpoint, or base URL, used by the client.[getEndpointProvider()](class-aws-awsclient.md#method_getEndpointProvider)
: mixed [getEndpointProviderArgs()](class-aws-awsclient.md#method_getEndpointProviderArgs)
: array<string\|int, mixed> Retrieves arguments to be used in endpoint resolution.[getHandlerList()](class-aws-awsclient.md#method_getHandlerList)
: [HandlerList](class-aws-handlerlist.md)Get the handler list used to transfer commands.[getIterator()](class-aws-awsclienttrait.md#method_getIterator)
: mixed [getPaginator()](class-aws-awsclienttrait.md#method_getPaginator)
: mixed [getRegion()](class-aws-awsclient.md#method_getRegion)
: string Get the region to which the client is configured to send requests.[getSignatureProvider()](class-aws-awsclient.md#method_getSignatureProvider)
: callable Get the signature\_provider function of the client.[getToken()](class-aws-awsclient.md#method_getToken)
: mixed [getWaiter()](class-aws-awsclienttrait.md#method_getWaiter)
: mixed [waitUntil()](class-aws-awsclienttrait.md#method_waitUntil)
: mixed

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Ec2.Ec2Client.html\#methods)

#### \_\_call()  [header link](class-aws-awsclienttrait.md\#method___call)

`
    public
                    __call(mixed $name, array<string|int, mixed> $args) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Ec2.Ec2Client.html\#method___construct)

The client constructor accepts the following options:

`
    public
                    __construct(array<string|int, mixed> $args) : mixed`

- api\_provider: (callable) An optional PHP callable that accepts a
type, service, and version argument, and returns an array of
corresponding configuration data. The type value can be one of api,
waiter, or paginator.
- credentials:
(Aws\\Credentials\\CredentialsInterface\|array\|bool\|callable) Specifies
the credentials used to sign requests. Provide an
Aws\\Credentials\\CredentialsInterface object, an associative array of
"key", "secret", and an optional "token" key, `false` to use null
credentials, or a callable credentials provider used to create
credentials or return null. See Aws\\Credentials\\CredentialProvider for
a list of built-in credentials providers. If no credentials are
provided, the SDK will attempt to load them from the environment.
- token:
(Aws\\Token\\TokenInterface\|array\|bool\|callable) Specifies
the token used to authorize requests. Provide an
Aws\\Token\\TokenInterface object, an associative array of
"token" and an optional "expires" key, `false` to use no
token, or a callable token provider used to create a
token or return null. See Aws\\Token\\TokenProvider for
a list of built-in token providers. If no token is
provided, the SDK will attempt to load one from the environment.
- csm:
(Aws\\ClientSideMonitoring\\ConfigurationInterface\|array\|callable) Specifies
the credentials used to sign requests. Provide an
Aws\\ClientSideMonitoring\\ConfigurationInterface object, a callable
configuration provider used to create client-side monitoring configuration,
`false` to disable csm, or an associative array with the following keys:
enabled: (bool) Set to true to enable client-side monitoring, defaults
to false; host: (string) the host location to send monitoring events to,
defaults to 127.0.0.1; port: (int) The port used for the host connection,
defaults to 31000; client\_id: (string) An identifier for this project
- debug: (bool\|array) Set to true to display debug information when
sending requests. Alternatively, you can provide an associative array
with the following keys: logfn: (callable) Function that is invoked
with log messages; stream\_size: (int) When the size of a stream is
greater than this number, the stream data will not be logged (set to
"0" to not log any stream data); scrub\_auth: (bool) Set to false to
disable the scrubbing of auth data from the logged messages; http:
(bool) Set to false to disable the "debug" feature of lower level HTTP
adapters (e.g., verbose curl output).
- stats: (bool\|array) Set to true to gather transfer statistics on
requests sent. Alternatively, you can provide an associative array with
the following keys: retries: (bool) Set to false to disable reporting
on retries attempted; http: (bool) Set to true to enable collecting
statistics from lower level HTTP adapters (e.g., values returned in
GuzzleHttp\\TransferStats). HTTP handlers must support an
`http_stats_receiver` option for this to have an effect; timer: (bool)
Set to true to enable a command timer that reports the total wall clock
time spent on an operation in seconds.
- disable\_host\_prefix\_injection: (bool) Set to true to disable host prefix
injection logic for services that use it. This disables the entire
prefix injection, including the portions supplied by user-defined
parameters. Setting this flag will have no effect on services that do
not use host prefix injection.
- endpoint: (string) The full URI of the webservice. This is only
required when connecting to a custom endpoint (e.g., a local version
of S3).
- endpoint\_discovery: (Aws\\EndpointDiscovery\\ConfigurationInterface,
Aws\\CacheInterface, array, callable) Settings for endpoint discovery.
Provide an instance of Aws\\EndpointDiscovery\\ConfigurationInterface,
an instance Aws\\CacheInterface, a callable that provides a promise for
a Configuration object, or an associative array with the following
keys: enabled: (bool) Set to true to enable endpoint discovery, false
to explicitly disable it, defaults to false; cache\_limit: (int) The
maximum number of keys in the endpoints cache, defaults to 1000.
- endpoint\_provider: (callable) An optional PHP callable that
accepts a hash of options including a "service" and "region" key and
returns NULL or a hash of endpoint data, of which the "endpoint" key
is required. See Aws\\Endpoint\\EndpointProvider for a list of built-in
providers.
- handler: (callable) A handler that accepts a command object,
request object and returns a promise that is fulfilled with an
Aws\\ResultInterface object or rejected with an
Aws\\Exception\\AwsException. A handler does not accept a next handler
as it is terminal and expected to fulfill a command. If no handler is
provided, a default Guzzle handler will be utilized.
- http: (array, default=array(0)) Set to an array of SDK request
options to apply to each request (e.g., proxy, verify, etc.).
- http\_handler: (callable) An HTTP handler is a function that
accepts a PSR-7 request object and returns a promise that is fulfilled
with a PSR-7 response object or rejected with an array of exception
data. NOTE: This option supersedes any provided "handler" option.
- idempotency\_auto\_fill: (bool\|callable) Set to false to disable SDK to
populate parameters that enabled 'idempotencyToken' trait with a random
UUID v4 value on your behalf. Using default value 'true' still allows
parameter value to be overwritten when provided. Note: auto-fill only
works when cryptographically secure random bytes generator functions
(random\_bytes, openssl\_random\_pseudo\_bytes or mcrypt\_create\_iv) can be
found. You may also provide a callable source of random bytes.
- profile: (string) Allows you to specify which profile to use when
credentials are created from the AWS credentials file in your HOME
directory. This setting overrides the AWS\_PROFILE environment
variable. Note: Specifying "profile" will cause the "credentials" key
to be ignored.
- region: (string, required) Region to connect to. See
http://docs.aws.amazon.com/general/latest/gr/rande.html for a list of
available regions.
- retries: (int, Aws\\Retry\\ConfigurationInterface, Aws\\CacheInterface,
array, callable) Configures the retry mode and maximum number of
allowed retries for a client (pass 0 to disable retries). Provide an
integer for 'legacy' mode with the specified number of retries.
Otherwise provide an instance of Aws\\Retry\\ConfigurationInterface, an
instance of Aws\\CacheInterface, a callable function, or an array with
the following keys: mode: (string) Set to 'legacy', 'standard' (uses
retry quota management), or 'adapative' (an experimental mode that adds
client-side rate limiting to standard mode); max\_attempts (int) The
maximum number of attempts for a given request.
- scheme: (string, default=string(5) "https") URI scheme to use when
connecting connect. The SDK will utilize "https" endpoints (i.e.,
utilize SSL/TLS connections) by default. You can attempt to connect to
a service over an unencrypted "http" endpoint by setting `scheme` to
"http".
- signature\_provider: (callable) A callable that accepts a signature
version name (e.g., "v4"), a service name, and region, and
returns a SignatureInterface object or null. This provider is used to
create signers utilized by the client. See
Aws\\Signature\\SignatureProvider for a list of built-in providers
- signature\_version: (string) A string representing a custom
signature version to use with a service (e.g., v4). Note that
per/operation signature version MAY override this requested signature
version.
- use\_aws\_shared\_config\_files: (bool, default=bool(true)) Set to false to
disable checking for shared config file in '~/.aws/config' and
'~/.aws/credentials'. This will override the AWS\_CONFIG\_FILE
environment variable.
- validate: (bool, default=bool(true)) Set to false to disable
client-side parameter validation.
- version: (string, required) The version of the webservice to
utilize (e.g., 2006-03-01).
- account\_id\_endpoint\_mode: (string, default(preferred)) this option
decides whether credentials should resolve an accountId value,
which is going to be used as part of the endpoint resolution.
The valid values for this option are:
  - preferred: when this value is set then, a warning is logged when
    accountId is empty in the resolved identity.
  - required: when this value is set then, an exception is thrown when
    accountId is empty in the resolved identity.
  - disabled: when this value is set then, the validation for if accountId
    was resolved or not, is ignored.
- ua\_append: (string, array) To pass custom user agent parameters.
- app\_id: (string) an optional application specific identifier that can be set.
When set it will be appended to the User-Agent header of every request
in the form of App/{AppId}. This variable is sourced from environment
variable AWS\_SDK\_UA\_APP\_ID or the shared config profile attribute sdk\_ua\_app\_id.
See https://docs.aws.amazon.com/sdkref/latest/guide/settings-reference.html for
more information on environment variables and shared config settings.

##### Parameters

$args
: array<string\|int, mixed>

Client configuration arguments.

#### \_\_sleep()  [header link](class-aws-awsclient.md\#method___sleep)

`
    public
                    __sleep() : mixed`

#### execute()  [header link](class-aws-awsclienttrait.md\#method_execute)

`
    public
                    execute(CommandInterface $command) : mixed`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

#### executeAsync()  [header link](class-aws-awsclienttrait.md\#method_executeAsync)

`
    public
                    executeAsync(CommandInterface $command) : mixed`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

#### factory()  [header link](class-aws-awsclient.md\#method_factory)

`
    public
            static        factory([array<string|int, mixed> $config = [] ]) : static`

##### Parameters

$config
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-awsclient.md\#method_factory\#tags)

deprecated

##### Return values

static

#### getApi()  [header link](class-aws-awsclienttrait.md\#method_getApi)

`
    public
    abstract                getApi() : Service`

##### Return values

[Service](class-aws-api-service.md)

#### getArguments()  [header link](class-aws-awsclient.md\#method_getArguments)

Get an array of client constructor arguments used by the client.

`
    public
            static        getArguments() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getClientBuiltIns()  [header link](class-aws-awsclient.md\#method_getClientBuiltIns)

Provides the set of built-in keys and values
used for endpoint resolution

`
    public
                    getClientBuiltIns() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getClientContextParams()  [header link](class-aws-awsclient.md\#method_getClientContextParams)

Provides the set of service context parameter
key-value pairs used for endpoint resolution.

`
    public
                    getClientContextParams() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getCommand()  [header link](class-aws-awsclienttrait.md\#method_getCommand)

`
    public
    abstract                getCommand(string $name[, array<string|int, mixed> $args = [] ]) : CommandInterface`

##### Parameters

$name
: string$args
: array<string\|int, mixed>
= \[\]

##### Return values

[CommandInterface](class-aws-commandinterface.md)

#### getConfig()  [header link](class-aws-awsclient.md\#method_getConfig)

Get a client configuration value.

`
    public
                    getConfig([mixed $option = null ]) : mixed|null`

##### Parameters

$option
: mixed
= null

The option to retrieve. Pass null to retrieve
all options.

##### Return values

mixed\|null

#### getCredentials()  [header link](class-aws-awsclient.md\#method_getCredentials)

Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.

`
    public
                    getCredentials() : PromiseInterface`

If you need the credentials synchronously, then call the wait() method
on the returned promise.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### getEndpoint()  [header link](class-aws-awsclient.md\#method_getEndpoint)

Gets the default endpoint, or base URL, used by the client.

`
    public
                    getEndpoint() : UriInterface`

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)

#### getEndpointProvider()  [header link](class-aws-awsclient.md\#method_getEndpointProvider)

`
    public
                    getEndpointProvider() : mixed`

#### getEndpointProviderArgs()  [header link](class-aws-awsclient.md\#method_getEndpointProviderArgs)

Retrieves arguments to be used in endpoint resolution.

`
    public
                    getEndpointProviderArgs() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getHandlerList()  [header link](class-aws-awsclient.md\#method_getHandlerList)

Get the handler list used to transfer commands.

`
    public
                    getHandlerList() : HandlerList`

This list can be modified to add middleware or to change the underlying
handler used to send HTTP requests.

##### Return values

[HandlerList](class-aws-handlerlist.md)

#### getIterator()  [header link](class-aws-awsclienttrait.md\#method_getIterator)

`
    public
                    getIterator(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### getPaginator()  [header link](class-aws-awsclienttrait.md\#method_getPaginator)

`
    public
                    getPaginator(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### getRegion()  [header link](class-aws-awsclient.md\#method_getRegion)

Get the region to which the client is configured to send requests.

`
    public
                    getRegion() : string`

##### Return values

string

#### getSignatureProvider()  [header link](class-aws-awsclient.md\#method_getSignatureProvider)

Get the signature\_provider function of the client.

`
    public
        final            getSignatureProvider() : callable`

##### Return values

callable

#### getToken()  [header link](class-aws-awsclient.md\#method_getToken)

`
    public
                    getToken() : mixed`

#### getWaiter()  [header link](class-aws-awsclienttrait.md\#method_getWaiter)

`
    public
                    getWaiter(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### waitUntil()  [header link](class-aws-awsclienttrait.md\#method_waitUntil)

`
    public
                    waitUntil(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Ec2.Ec2Client.html#toc-methods)
- Methods
  - [\_\_call()](class-aws-awsclienttrait.md#method___call)
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Ec2.Ec2Client.html#method___construct)
  - [\_\_sleep()](class-aws-awsclient.md#method___sleep)
  - [execute()](class-aws-awsclienttrait.md#method_execute)
  - [executeAsync()](class-aws-awsclienttrait.md#method_executeAsync)
  - [factory()](class-aws-awsclient.md#method_factory)
  - [getApi()](class-aws-awsclienttrait.md#method_getApi)
  - [getArguments()](class-aws-awsclient.md#method_getArguments)
  - [getClientBuiltIns()](class-aws-awsclient.md#method_getClientBuiltIns)
  - [getClientContextParams()](class-aws-awsclient.md#method_getClientContextParams)
  - [getCommand()](class-aws-awsclienttrait.md#method_getCommand)
  - [getConfig()](class-aws-awsclient.md#method_getConfig)
  - [getCredentials()](class-aws-awsclient.md#method_getCredentials)
  - [getEndpoint()](class-aws-awsclient.md#method_getEndpoint)
  - [getEndpointProvider()](class-aws-awsclient.md#method_getEndpointProvider)
  - [getEndpointProviderArgs()](class-aws-awsclient.md#method_getEndpointProviderArgs)
  - [getHandlerList()](class-aws-awsclient.md#method_getHandlerList)
  - [getIterator()](class-aws-awsclienttrait.md#method_getIterator)
  - [getPaginator()](class-aws-awsclienttrait.md#method_getPaginator)
  - [getRegion()](class-aws-awsclient.md#method_getRegion)
  - [getSignatureProvider()](class-aws-awsclient.md#method_getSignatureProvider)
  - [getToken()](class-aws-awsclient.md#method_getToken)
  - [getWaiter()](class-aws-awsclienttrait.md#method_getWaiter)
  - [waitUntil()](class-aws-awsclienttrait.md#method_waitUntil)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Ec2.Ec2Client.html#top)
