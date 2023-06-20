"use strict";(self.webpackChunkcosmos_sdk_docs=self.webpackChunkcosmos_sdk_docs||[]).push([[4096],{3905:(e,t,n)=>{n.d(t,{Zo:()=>m,kt:()=>u});var r=n(67294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function s(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?s(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):s(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function o(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},s=Object.keys(e);for(r=0;r<s.length;r++)n=s[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);for(r=0;r<s.length;r++)n=s[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var l=r.createContext({}),p=function(e){var t=r.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},m=function(e){var t=p(e.components);return r.createElement(l.Provider,{value:t},e.children)},d={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},c=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,s=e.originalType,l=e.parentName,m=o(e,["components","mdxType","originalType","parentName"]),c=p(n),u=a,g=c["".concat(l,".").concat(u)]||c[u]||d[u]||s;return n?r.createElement(g,i(i({ref:t},m),{},{components:n})):r.createElement(g,i({ref:t},m))}));function u(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var s=n.length,i=new Array(s);i[0]=c;var o={};for(var l in t)hasOwnProperty.call(t,l)&&(o[l]=t[l]);o.originalType=e,o.mdxType="string"==typeof e?e:a,i[1]=o;for(var p=2;p<s;p++)i[p]=n[p];return r.createElement.apply(null,i)}return r.createElement.apply(null,n)}c.displayName="MDXCreateElement"},28831:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>l,contentTitle:()=>i,default:()=>d,frontMatter:()=>s,metadata:()=>o,toc:()=>p});var r=n(87462),a=(n(67294),n(3905));const s={sidebar_position:1},i="Msg Services",o={unversionedId:"building-modules/msg-services",id:"version-v0.50/building-modules/msg-services",title:"Msg Services",description:"A Protobuf Msg service processes messages. Protobuf Msg services are specific to the module in which they are defined, and only process messages defined within the said module. They are called from BaseApp during DeliverTx.",source:"@site/versioned_docs/version-v0.50/building-modules/03-msg-services.md",sourceDirName:"building-modules",slug:"/building-modules/msg-services",permalink:"/v0.50/building-modules/msg-services",draft:!1,tags:[],version:"v0.50",sidebarPosition:1,frontMatter:{sidebar_position:1},sidebar:"tutorialSidebar",previous:{title:"Messages and Queries",permalink:"/v0.50/building-modules/messages-and-queries"},next:{title:"Query Services",permalink:"/v0.50/building-modules/query-services"}},l={},p=[{value:"Implementation of a module <code>Msg</code> service",id:"implementation-of-a-module-msg-service",level:2},{value:"Validation",id:"validation",level:3},{value:"State Transition",id:"state-transition",level:3},{value:"Events",id:"events",level:3},{value:"Telemetry",id:"telemetry",level:2}],m={toc:p};function d(e){let{components:t,...n}=e;return(0,a.kt)("wrapper",(0,r.Z)({},m,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"msg-services"},(0,a.kt)("inlineCode",{parentName:"h1"},"Msg")," Services"),(0,a.kt)("admonition",{title:"Synopsis",type:"note"},(0,a.kt)("p",{parentName:"admonition"},"A Protobuf ",(0,a.kt)("inlineCode",{parentName:"p"},"Msg")," service processes ",(0,a.kt)("a",{parentName:"p",href:"/v0.50/building-modules/messages-and-queries#messages"},"messages"),". Protobuf ",(0,a.kt)("inlineCode",{parentName:"p"},"Msg")," services are specific to the module in which they are defined, and only process messages defined within the said module. They are called from ",(0,a.kt)("inlineCode",{parentName:"p"},"BaseApp")," during ",(0,a.kt)("a",{parentName:"p",href:"/v0.50/core/baseapp#delivertx"},(0,a.kt)("inlineCode",{parentName:"a"},"DeliverTx")),".")),(0,a.kt)("admonition",{type:"note"},(0,a.kt)("h3",{parentName:"admonition",id:"pre-requisite-readings"},"Pre-requisite Readings"),(0,a.kt)("ul",{parentName:"admonition"},(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"/v0.50/building-modules/module-manager"},"Module Manager")),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"/v0.50/building-modules/messages-and-queries"},"Messages and Queries")))),(0,a.kt)("h2",{id:"implementation-of-a-module-msg-service"},"Implementation of a module ",(0,a.kt)("inlineCode",{parentName:"h2"},"Msg")," service"),(0,a.kt)("p",null,"Each module should define a Protobuf ",(0,a.kt)("inlineCode",{parentName:"p"},"Msg")," service, which will be responsible for processing requests (implementing ",(0,a.kt)("inlineCode",{parentName:"p"},"sdk.Msg"),") and returning responses."),(0,a.kt)("p",null,"As further described in ",(0,a.kt)("a",{parentName:"p",href:"/v0.50/architecture/adr-031-msg-service"},"ADR 031"),", this approach has the advantage of clearly specifying return types and generating server and client code."),(0,a.kt)("p",null,"Protobuf generates a ",(0,a.kt)("inlineCode",{parentName:"p"},"MsgServer")," interface based on a definition of ",(0,a.kt)("inlineCode",{parentName:"p"},"Msg")," service. It is the role of the module developer to implement this interface, by implementing the state transition logic that should happen upon receival of each ",(0,a.kt)("inlineCode",{parentName:"p"},"sdk.Msg"),". As an example, here is the generated ",(0,a.kt)("inlineCode",{parentName:"p"},"MsgServer")," interface for ",(0,a.kt)("inlineCode",{parentName:"p"},"x/bank"),", which exposes two ",(0,a.kt)("inlineCode",{parentName:"p"},"sdk.Msg"),"s:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/x/bank/types/tx.pb.go#L550-L568\n")),(0,a.kt)("p",null,"When possible, the existing module's ",(0,a.kt)("a",{parentName:"p",href:"/v0.50/building-modules/keeper"},(0,a.kt)("inlineCode",{parentName:"a"},"Keeper"))," should implement ",(0,a.kt)("inlineCode",{parentName:"p"},"MsgServer"),", otherwise a ",(0,a.kt)("inlineCode",{parentName:"p"},"msgServer")," struct that embeds the ",(0,a.kt)("inlineCode",{parentName:"p"},"Keeper")," can be created, typically in ",(0,a.kt)("inlineCode",{parentName:"p"},"./keeper/msg_server.go"),":"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/x/bank/keeper/msg_server.go#L15-L17\n")),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"msgServer")," methods can retrieve the ",(0,a.kt)("inlineCode",{parentName:"p"},"sdk.Context")," from the ",(0,a.kt)("inlineCode",{parentName:"p"},"context.Context")," parameter method using the ",(0,a.kt)("inlineCode",{parentName:"p"},"sdk.UnwrapSDKContext"),":"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/x/bank/keeper/msg_server.go#L28\n")),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"sdk.Msg")," processing usually follows these 3 steps:"),(0,a.kt)("h3",{id:"validation"},"Validation"),(0,a.kt)("p",null,"The message server must perform all validation required (both ",(0,a.kt)("em",{parentName:"p"},"stateful")," and ",(0,a.kt)("em",{parentName:"p"},"stateless"),") to make sure the ",(0,a.kt)("inlineCode",{parentName:"p"},"message")," is valid.\nThe ",(0,a.kt)("inlineCode",{parentName:"p"},"signer")," is charged for the gas cost of this validation."),(0,a.kt)("p",null,"For example, a ",(0,a.kt)("inlineCode",{parentName:"p"},"msgServer")," method for a ",(0,a.kt)("inlineCode",{parentName:"p"},"transfer")," message should check that the sending account has enough funds to actually perform the transfer. "),(0,a.kt)("p",null,"It is recommended to implement all validation checks in a separate function that passes state values as arguments. This implementation simplifies testing. As expected, expensive validation functions charge additional gas. Example:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'ValidateMsgA(msg MsgA, now Time, gm GasMeter) error {\n    if now.Before(msg.Expire) {\n        return sdkerrrors.ErrInvalidRequest.Wrap("msg expired")\n    }\n    gm.ConsumeGas(1000, "signature verification")\n    return signatureVerificaton(msg.Prover, msg.Data)\n}\n')),(0,a.kt)("admonition",{type:"warning"},(0,a.kt)("p",{parentName:"admonition"},"Previously, the ",(0,a.kt)("inlineCode",{parentName:"p"},"ValidateBasic")," method was used to perform simple and stateless validation checks.\nThis way of validating is deprecated, this means the ",(0,a.kt)("inlineCode",{parentName:"p"},"msgServer")," must perform all validation checks.")),(0,a.kt)("h3",{id:"state-transition"},"State Transition"),(0,a.kt)("p",null,"After the validation is successful, the ",(0,a.kt)("inlineCode",{parentName:"p"},"msgServer")," method uses the ",(0,a.kt)("a",{parentName:"p",href:"/v0.50/building-modules/keeper"},(0,a.kt)("inlineCode",{parentName:"a"},"keeper"))," functions to access the state and perform a state transition."),(0,a.kt)("h3",{id:"events"},"Events"),(0,a.kt)("p",null,"Before returning, ",(0,a.kt)("inlineCode",{parentName:"p"},"msgServer")," methods generally emit one or more ",(0,a.kt)("a",{parentName:"p",href:"/v0.50/core/events"},"events")," by using the ",(0,a.kt)("inlineCode",{parentName:"p"},"EventManager")," held in the ",(0,a.kt)("inlineCode",{parentName:"p"},"ctx"),". Use the new ",(0,a.kt)("inlineCode",{parentName:"p"},"EmitTypedEvent")," function that uses protobuf-based event types:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"ctx.EventManager().EmitTypedEvent(\n    &group.EventABC{Key1: Value1,  Key2, Value2})\n")),(0,a.kt)("p",null,"or the older ",(0,a.kt)("inlineCode",{parentName:"p"},"EmitEvent")," function: "),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"ctx.EventManager().EmitEvent(\n    sdk.NewEvent(\n        eventType,  // e.g. sdk.EventTypeMessage for a message, types.CustomEventType for a custom event defined in the module\n        sdk.NewAttribute(key1, value1),\n        sdk.NewAttribute(key2, value2),\n    ),\n)\n")),(0,a.kt)("p",null,"These events are relayed back to the underlying consensus engine and can be used by service providers to implement services around the application. Click ",(0,a.kt)("a",{parentName:"p",href:"/v0.50/core/events"},"here")," to learn more about events."),(0,a.kt)("p",null,"The invoked ",(0,a.kt)("inlineCode",{parentName:"p"},"msgServer")," method returns a ",(0,a.kt)("inlineCode",{parentName:"p"},"proto.Message")," response and an ",(0,a.kt)("inlineCode",{parentName:"p"},"error"),". These return values are then wrapped into an ",(0,a.kt)("inlineCode",{parentName:"p"},"*sdk.Result")," or an ",(0,a.kt)("inlineCode",{parentName:"p"},"error")," using ",(0,a.kt)("inlineCode",{parentName:"p"},"sdk.WrapServiceResult(ctx sdk.Context, res proto.Message, err error)"),":"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/baseapp/msg_service_router.go#L131\n")),(0,a.kt)("p",null,"This method takes care of marshaling the ",(0,a.kt)("inlineCode",{parentName:"p"},"res")," parameter to protobuf and attaching any events on the ",(0,a.kt)("inlineCode",{parentName:"p"},"ctx.EventManager()")," to the ",(0,a.kt)("inlineCode",{parentName:"p"},"sdk.Result"),"."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-protobuf",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/proto/cosmos/base/abci/v1beta1/abci.proto#L88-L109\n")),(0,a.kt)("p",null,"This diagram shows a typical structure of a Protobuf ",(0,a.kt)("inlineCode",{parentName:"p"},"Msg")," service, and how the message propagates through the module."),(0,a.kt)("p",null,(0,a.kt)("img",{parentName:"p",src:"https://raw.githubusercontent.com/cosmos/cosmos-sdk/release/v0.46.x/docs/uml/svg/transaction_flow.svg",alt:"Transaction flow"})),(0,a.kt)("h2",{id:"telemetry"},"Telemetry"),(0,a.kt)("p",null,"New ",(0,a.kt)("a",{parentName:"p",href:"/v0.50/core/telemetry"},"telemetry metrics")," can be created from ",(0,a.kt)("inlineCode",{parentName:"p"},"msgServer")," methods when handling messages."),(0,a.kt)("p",null,"This is an example from the ",(0,a.kt)("inlineCode",{parentName:"p"},"x/auth/vesting")," module:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/x/auth/vesting/msg_server.go#L68-L80\n")))}d.isMDXComponent=!0}}]);