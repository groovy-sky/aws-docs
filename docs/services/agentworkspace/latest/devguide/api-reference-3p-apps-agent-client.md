---
title: "Connect Customer agent workspace Agent API"
---

# Connect Customer agent workspace Agent API

The Amazon Connect SDK provides an `AgentClient` which serves as an interface that your app in the
Connect Customer agent workspace can use to subscribe to agent events and make agent data requests.

The `AgentClient` accepts an optional constructor argument, `
        ConnectClientConfig` which itself is defined as:

```typescript

export type ConnectClientConfig = {
    context?: ModuleContext;
    provider?: AmazonConnectProvider;
};

```

If you do not provide a value for this config, then the client will default to using the **AmazonConnectProvider** set in the global provider scope. You can also manually
configure this using **setGlobalProvider**.

You can instantiate the agent client as follows:

```typescript

import { AgentClient } from "@amazon-connect/contact";

const agentClient = new AgentClient({ provider });

```

###### Note

You must first instantiate the [AmazonConnectApp](../../../../reference/agentworkspace/latest/devguide/getting-started-initialize-sdk.md) which initializes the default AmazonConnectProvider and returns `
            { provider } `. This is the recommended option.

Alternatively, providing a constructor argument:

```typescript

import { AgentClient } from "@amazon-connect/contact";

const agentClient = new AgentClient({
    context: sampleContext,
    provider: sampleProvider
});

```

The following sections describe API calls for working with the Agent API.

###### Contents

- [getARN()](3p-apps-agent-requests-getarn.md)

- [getChannelConcurrency()](3p-apps-agent-requests-getchannelconcurrency.md)

- [getExtension()](3p-apps-agent-requests-getextension.md)

- [getName()](3p-apps-agent-requests-getname.md)

- [getRoutingProfile()](3p-apps-agent-requests-getroutingprofile.md)

- [getState() - Deprecated](3p-apps-agent-requests-getstate.md)

- [listAvailabilityStates()](3p-apps-agent-requests-listavailabilitystates.md)

- [listQuickConnects()](3p-apps-agent-requests-listquickconnects.md)

- [offEnabledChannelListChanged()](3p-apps-agent-requests-off-enabledchannellistchanged.md)

- [offRoutingProfileChanged()](3p-apps-agent-requests-off-routingprofilechanged.md)

- [onEnabledChannelListChanged()](3p-apps-agent-requests-on-enabledchannellistchanged.md)

- [onRoutingProfileChanged()](3p-apps-agent-requests-on-routingprofile-changed.md)

- [setAvailabilityState()](3p-apps-agent-requests-setavailabilitystate.md)

- [setAvailabilityStateByName()](3p-apps-agent-requests-setavailabilitystatebyname.md)

- [setOffline()](3p-apps-agent-requests-setoffline.md)

- [onStateChanged() - Deprecated](3p-apps-agent-events-statechanged-sub.md)

- [offStateChanged() - Deprecated](3p-apps-agent-events-statechanged-unsub.md)

- [getDefaultOutboundQueue()](3p-apps-agent-requests-getdefaultoutboundqueue.md)

- [getRoutingProfileQueues()](3p-apps-agent-requests-getroutingprofilequeues.md)

- [getAvailabilityState()](3p-apps-agent-requests-getavailabilitystate.md)

- [listSecurityProfilePermissions()](3p-apps-agent-requests-listsecurityprofilepermissions.md)

- [onAvailabilityStateChanged()](3p-apps-agent-events-availabilitystatechanged-sub.md)

- [offAvailabilityStateChanged()](3p-apps-agent-events-availabilitystatechanged-unsub.md)

- [onNextAvailabilityStateChanged()](3p-apps-agent-events-nextavailabilitystatechanged-sub.md)

- [getNetworkConnectionStatus()](3p-apps-agent-requests-getnetworkconnectionstatus.md)

- [onNetworkConnectionStatusChanged()](3p-apps-agent-events-networkconnectionstatuschanged-sub.md)

- [offNetworkConnectionStatusChanged()](3p-apps-agent-events-networkconnectionstatuschanged-unsub.md)

- [offNextAvailabilityStateChanged()](3p-apps-agent-events-nextavailabilitystatechanged-unsub.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

sendActivity()

getARN()

All content copied from https://docs.aws.amazon.com/.
