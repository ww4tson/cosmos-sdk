(window.webpackJsonp=window.webpackJsonp||[]).push([[291],{822:function(t,e,s){"use strict";s.r(e);var a=s(1),o=Object(a.a)({},(function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("ContentSlotsDistributor",{attrs:{"slot-key":t.$parent.slotKey}},[s("h1",{attrs:{id:"running-a-testnet"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#running-a-testnet"}},[t._v("#")]),t._v(" Running a Testnet")]),t._v(" "),s("p",{attrs:{synopsis:""}},[t._v("The "),s("code",[t._v("simd testnet")]),t._v(" subcommand makes it easy to initialize and start a simulated test network for testing purposes.")]),t._v(" "),s("p",[t._v("In addition to the commands for "),s("RouterLink",{attrs:{to:"/run-node/run-node.html"}},[t._v("running a node")]),t._v(", the "),s("code",[t._v("simd")]),t._v(" binary also includes a "),s("code",[t._v("testnet")]),t._v(" command that allows you to start a simulated test network in-process or to initialize files for a simulated test network that runs in a separate process.")],1),t._v(" "),s("h2",{attrs:{id:"initialize-files"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#initialize-files"}},[t._v("#")]),t._v(" Initialize Files")]),t._v(" "),s("p",[t._v("First, let's take a look at the "),s("code",[t._v("init-files")]),t._v(" subcommand.")]),t._v(" "),s("p",[t._v("This is similar to the "),s("code",[t._v("init")]),t._v(" command when initializing a single node, but in this case we are initializing multiple nodes, generating the genesis transactions for each node, and then collecting those transactions.")]),t._v(" "),s("p",[t._v("The "),s("code",[t._v("init-files")]),t._v(" subcommand initializes the necessary files to run a test network in a separate process (i.e. using a Docker container). Running this command is not a prerequisite for the "),s("code",[t._v("start")]),t._v(" subcommand ("),s("a",{attrs:{href:"#start-testnet"}},[t._v("see below")]),t._v(").")]),t._v(" "),s("p",[t._v("In order to initialize the files for a test network, run the following command:")]),t._v(" "),s("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"c2ltZCB0ZXN0bmV0IGluaXQtZmlsZXMK"}}),t._v(" "),s("p",[t._v("You should see the following output in your terminal:")]),t._v(" "),s("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"U3VjY2Vzc2Z1bGx5IGluaXRpYWxpemVkIDQgbm9kZSBkaXJlY3Rvcmllcwo="}}),t._v(" "),s("p",[t._v("The default output directory is a relative "),s("code",[t._v(".testnets")]),t._v(" directory. Let's take a look at the files created within the "),s("code",[t._v(".testnets")]),t._v(" directory.")]),t._v(" "),s("h3",{attrs:{id:"gentxs"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#gentxs"}},[t._v("#")]),t._v(" gentxs")]),t._v(" "),s("p",[t._v("The "),s("code",[t._v("gentxs")]),t._v(" directory includes a genesis transaction for each validator node. Each file includes a JSON encoded genesis transaction used to register a validator node at the time of genesis. The genesis transactions are added to the "),s("code",[t._v("genesis.json")]),t._v(" file within each node directory during the initilization process.")]),t._v(" "),s("h3",{attrs:{id:"nodes"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#nodes"}},[t._v("#")]),t._v(" nodes")]),t._v(" "),s("p",[t._v("A node directory is created for each validator node. Within each node directory is a "),s("code",[t._v("simd")]),t._v(" directory. The "),s("code",[t._v("simd")]),t._v(" directory is the home directory for each node, which includes the configuration and data files for that node (i.e. the same files included in the default "),s("code",[t._v("~/.simapp")]),t._v(" directory when running a single node).")]),t._v(" "),s("h2",{attrs:{id:"start-testnet"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#start-testnet"}},[t._v("#")]),t._v(" Start Testnet")]),t._v(" "),s("p",[t._v("Now, let's take a look at the "),s("code",[t._v("start")]),t._v(" subcommand.")]),t._v(" "),s("p",[t._v("The "),s("code",[t._v("start")]),t._v(" subcommand both initializes and starts an in-process test network. This is the fastest way to spin up a local test network for testing purposes.")]),t._v(" "),s("p",[t._v("You can start the local test network by running the following command:")]),t._v(" "),s("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"c2ltZCB0ZXN0bmV0IHN0YXJ0Cg=="}}),t._v(" "),s("p",[t._v("You should see something similar to the following:")]),t._v(" "),s("tm-code-block",{staticClass:"codeblock",attrs:{language:"bash",base64:"YWNxdWlyaW5nIHRlc3QgbmV0d29yayBsb2NrCnByZXBhcmluZyB0ZXN0IG5ldHdvcmsgd2l0aCBjaGFpbi1pZCAmcXVvdDtjaGFpbi1tdG9EOXYmcXVvdDsKCgorKysrKysrKysrKysrKysrKysrKysrKysrKysrKysrKysrKysrKysrKysrKysrKysrKysrKysrKysrKysrCisrICAgICAgIFRISVMgTU5FTU9OSUMgSVMgRk9SIFRFU1RJTkcgUFVSUE9TRVMgT05MWSAgICAgICAgKysKKysgICAgICAgICAgICAgICAgRE8gTk9UIFVTRSBJTiBQUk9EVUNUSU9OICAgICAgICAgICAgICAgICArKworKyAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICsrCisrICBzdXN0YWluIGtub3cgZGVicmlzIG1pbnV0ZSBnYXRlIGh5YnJpZCBzdGVyZW8gY3VzdG9tICAgKysKKysgIGRpdm9yY2UgY3Jvc3Mgc3Bvb24gbWFjaGluZSBsYXRpbiB2aWJyYW50IHRlcm0gb2JsaWdlICArKworKyAgIG1vbWVudCBiZWF1dHkgbGF1bmRyeSByZXBlYXQgZ3JhYiBnYW1lIGJyb256ZSB0cnVseSAgICsrCisrKysrKysrKysrKysrKysrKysrKysrKysrKysrKysrKysrKysrKysrKysrKysrKysrKysrKysrKysrKysKCgpzdGFydGluZyB0ZXN0IG5ldHdvcmsuLi4Kc3RhcnRlZCB0ZXN0IG5ldHdvcmsKcHJlc3MgdGhlIEVudGVyIEtleSB0byB0ZXJtaW5hdGUK"}}),t._v(" "),s("p",[t._v("The first validator node is now running in-process, which means the test network will terminate once you either close the terminal window or you press the Enter key. In the output, the mnemonic phrase for the first validator node is provided for testing purposes. The validator node is using the same default addresses being used when initializing and starting a single node (no need to provide a "),s("code",[t._v("--node")]),t._v(" flag).")]),t._v(" "),s("p",[t._v("Check the status of the first validator node:")]),t._v(" "),s("tm-code-block",{staticClass:"codeblock",attrs:{language:"",base64:"c2ltZCBzdGF0dXMK"}}),t._v(" "),s("p",[t._v("Import the key from the provided mnemonic:")]),t._v(" "),s("tm-code-block",{staticClass:"codeblock",attrs:{language:"",base64:"c2ltZCBrZXlzIGFkZCB0ZXN0IC0tcmVjb3ZlciAtLWtleXJpbmctYmFja2VuZCB0ZXN0Cg=="}}),t._v(" "),s("p",[t._v("Check the balance of the account address:")]),t._v(" "),s("tm-code-block",{staticClass:"codeblock",attrs:{language:"",base64:"c2ltZCBxIGJhbmsgYmFsYW5jZXMgW2FkZHJlc3NdCg=="}}),t._v(" "),s("p",[t._v("Use this test account to manually test against the test network.")]),t._v(" "),s("h2",{attrs:{id:"testnet-options"}},[s("a",{staticClass:"header-anchor",attrs:{href:"#testnet-options"}},[t._v("#")]),t._v(" Testnet Options")]),t._v(" "),s("p",[t._v("You can customize the configuration of the test network with flags. In order to see all flag options, append the "),s("code",[t._v("--help")]),t._v(" flag to each command.")])],1)}),[],!1,null,null,null);e.default=o.exports}}]);