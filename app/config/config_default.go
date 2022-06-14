package config

var defaultCfg string = `
[App]
  channel_types = ["tendermint_and_eth"]
  env = "dev"
  log_level = "debug"
  metric_addr = "0.0.0.0:8083"
  bridge_status_api = "https://bridge.qa.davionlabs.com/bridge/bridge_status"
  bridge_enable = false

[Chain]

  [Chain.Dest]
    chain_type = "eth"
    enable = true

    [Chain.Dest.Cache]
      filename = "ethCache"
      start_height = 

    [Chain.Dest.Eth]
      chain_id = 3
      chain_name = "eth"
      comment_slot = 104
      gas_limit = 2000000
      max_gas_price = 150000000000
      tip_coefficient = 0.2
      uri = "https://ropsten.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161"
      update_client_frequency = 2
      query_filter = "ack"

      [Chain.Dest.Eth.Contracts]

        [Chain.Dest.Eth.Contracts.Ack_packet]
          addr = ""
          opt_priv_key = ""
          topic = "AckWritten((string,string,uint64,string,bytes,bytes,string,uint64),bytes)"

        [Chain.Dest.Eth.Contracts.Client]
          addr = ""
          opt_priv_key = ""
          topic = ""

        [Chain.Dest.Eth.Contracts.Packet]
          addr = ""
          opt_priv_key = ""
          topic = "PacketSent(bytes)"

  [Chain.Source]
    Chain_type = "tendermint"
    Enabled = true

    [Chain.Source.Cache]
      Filename = "sourcedata"
      Start_height = 10

    [Chain.Source.Tendermint]
      chain_id = ""
      chain_name = ""
      gas_limit =  
      gas_price = ""
      grpc_addr = "127.0.0.1:9090"
      simulation_addr = "127.0.0.1:9090"
      request_timeout = 0
      update_client_frequency = 10
      query_filter = ""

      [Chain.Source.Tendermint.Key]
        name = "node0"
        password = "1234567890"
        mnemonic = ""
`
