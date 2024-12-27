(window.webpackJsonp=window.webpackJsonp||[]).push([[165],{476:function(e,t,s){"use strict";s.r(t);var a=s(14),r=Object(a.a)({},(function(){var e=this,t=e._self._c;return t("ContentSlotsDistributor",{attrs:{"slot-key":e.$parent.slotKey}},[t("h1",{attrs:{id:"government"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#government"}},[e._v("#")]),e._v(" Government")]),e._v(" "),t("h2",{attrs:{id:"abstract"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#abstract"}},[e._v("#")]),e._v(" Abstract")]),e._v(" "),t("p",[e._v("In order to allow some specific operations to be performed only by a small set of individuals,\ninside Commercio.network we introduced the "),t("code",[e._v("government")]),e._v(" module.")]),e._v(" "),t("p",[e._v("This module allows for two simple operations:")]),e._v(" "),t("ol",[t("li",[e._v("Set a government address that will later be used as a sort of on-chain authentication method.")]),e._v(" "),t("li",[e._v("Read the government address that has been set.")])]),e._v(" "),t("h3",{attrs:{id:"setting-a-government-address"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#setting-a-government-address"}},[e._v("#")]),e._v(" Setting a government address")]),e._v(" "),t("p",[e._v("The address identified as the "),t("code",[e._v("government")]),e._v(" can be set "),t("strong",[e._v("only during the genesis")]),e._v(".\nThis operation can be performed using the following command:")]),e._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[e._v("commercionetworkd set-genesis-government-address "),t("span",{pre:!0,attrs:{class:"token operator"}},[e._v("<")]),e._v("ADDRESS-TO-USE"),t("span",{pre:!0,attrs:{class:"token operator"}},[e._v(">")]),e._v("\n")])])]),t("div",{staticClass:"custom-block danger"},[t("p",{staticClass:"custom-block-title"},[e._v("DANGER")]),e._v(" "),t("p",[t("strong",[e._v("Note")]),e._v(": you can run this command only once.\nRunning it several times after the first value has been set will result in an error been thrown inside the console.")])]),e._v(" "),t("h3",{attrs:{id:"query-government-address"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#query-government-address"}},[e._v("#")]),e._v(" Query government address")]),e._v(" "),t("p",[e._v("The government address can be get:")]),e._v(" "),t("ul",[t("li",[e._v("via "),t("strong",[e._v("CLI")]),e._v(", "),t("code",[e._v("commercionetworkd query government gov-address")])]),e._v(" "),t("li",[e._v("via "),t("strong",[e._v("REST")]),e._v(", by making a GET request to the "),t("code",[e._v("/commercionetwork/government/governmentAddress")]),e._v(" endpoint")]),e._v(" "),t("li",[e._v("via "),t("strong",[e._v("gRPC")]),e._v(", by making a Query to the "),t("code",[e._v("commercionetwork.commercionetwork.government.Query")]),e._v(" method")])]),e._v(" "),t("h4",{attrs:{id:"grpc"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#grpc"}},[e._v("#")]),e._v(" gRPC")]),e._v(" "),t("p",[e._v("Endpoint:")]),e._v(" "),t("div",{staticClass:"language- extra-class"},[t("pre",{pre:!0,attrs:{class:"language-text"}},[t("code",[e._v("commercionetwork.commercionetwork.government.Query/GovernmentAddr\n")])])]),t("h5",{attrs:{id:"example"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#example"}},[e._v("#")]),e._v(" Example")]),e._v(" "),t("div",{staticClass:"language-bash extra-class"},[t("pre",{pre:!0,attrs:{class:"language-bash"}},[t("code",[e._v("grpcurl "),t("span",{pre:!0,attrs:{class:"token parameter variable"}},[e._v("-plaintext")]),e._v(" "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v("\\")]),e._v("\n    localhost:9090 "),t("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v("\\")]),e._v("\n    commercionetwork.commercionetwork.government.Query/GovernmentAddr\n")])])]),t("h5",{attrs:{id:"response"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#response"}},[e._v("#")]),e._v(" Response")]),e._v(" "),t("div",{staticClass:"language-json extra-class"},[t("pre",{pre:!0,attrs:{class:"language-json"}},[t("code",[t("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v("{")]),e._v("\n  "),t("span",{pre:!0,attrs:{class:"token property"}},[e._v('"governmentAddress"')]),t("span",{pre:!0,attrs:{class:"token operator"}},[e._v(":")]),e._v(" "),t("span",{pre:!0,attrs:{class:"token string"}},[e._v('"did:com:1mj9h87yqjel0fsvkq55v345kxk0n09krtfvtyx"')]),e._v("\n"),t("span",{pre:!0,attrs:{class:"token punctuation"}},[e._v("}")]),e._v("\n")])])]),t("h2",{attrs:{id:"contents"}},[t("a",{staticClass:"header-anchor",attrs:{href:"#contents"}},[e._v("#")]),e._v(" Contents")]),e._v(" "),t("ol",[t("li",[t("strong",[t("RouterLink",{attrs:{to:"/modules/government/01_state.html"}},[e._v("State")])],1)])])])}),[],!1,null,null,null);t.default=r.exports}}]);