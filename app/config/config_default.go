package config

var defaultCfg string = `
[App]
  channel_types = ["tendermint_and_tendermint"]
  Env = "dev"
  log_level = "debug"
  metric_addr = "0.0.0.0:8083"

[Chain]

  [Chain.Dest]
    chain_type = "tendermint"
    enabled = true

    [Chain.Dest.Cache]
      filename = "destdata"
      start_height = 1

    [chain.dest.eth]
      chain_id = 0
      chain_name = "testA"
      comment_slot = 0
      gas_limit = 0
      max_gas_price = 0
      tip_coefficient = 0.0
      uri = ""
      UpdateClientFrequency = 0

      [chain.dest.eth.contracts]

        [Chain.Dest.Eth.Contracts.Ack_packet]
          Addr = ""
          OptPrivKey = ""
          Topic = ""

        [Chain.Dest.Eth.Contracts.CleanPacket]
          Addr = ""
          OptPrivKey = ""
          Topic = ""

        [Chain.Dest.Eth.Contracts.Client]
          Addr = ""
          OptPrivKey = ""
          Topic = ""

        [Chain.Dest.Eth.Contracts.Packet]
          Addr = ""
          OptPrivKey = ""
          Topic = ""

    [Chain.Dest.Tendermint]
      Algo = ""
      chain_id = "teleport_9000-2"
      chain_name = "testB"
      clean_packet_enabled = false
      Gas = 0
      grpc_addr = "127.0.0.1:19090"
      request_timeout = 0
      update_client_frequency = 10

      [Chain.Dest.Tendermint.Fee]
        Amount = 0
        Denom = ""

      [Chain.Dest.Tendermint.Key]
        Name = "validator"
        Password = "1234567890"
        mnemonic = "small pretty lock logic loud beef please boring space picnic essence also opera come roast pepper pumpkin vivid topple asset upon dismiss debris awful"

  [Chain.Source]
    Chain_type = "tendermint"
    Enabled = true

    [Chain.Source.Cache]
      Filename = "sourcedata"
      start_height = 1

    [Chain.Source.Eth]
      ChainID = 0
      ChainName = ""
      CommentSlot = 0
      GasLimit = 0
      MaxGasPrice = 0
      TipCoefficient = 0.0
      URI = ""
      UpdateClientFrequency = 0

      [Chain.Source.Eth.Contracts]

        [Chain.Source.Eth.Contracts.AckPacket]
          Addr = ""
          OptPrivKey = ""
          Topic = ""

        [Chain.Source.Eth.Contracts.CleanPacket]
          Addr = ""
          OptPrivKey = ""
          Topic = ""

        [Chain.Source.Eth.Contracts.Client]
          Addr = ""
          OptPrivKey = ""
          Topic = ""

        [Chain.Source.Eth.Contracts.Packet]
          Addr = ""
          OptPrivKey = ""
          Topic = ""

    [Chain.Source.Tendermint]
      Chain_id = "teleport_9000-1"
      Chain_name = "testA"
      Clean_packet_enabled = true
      Gas = 0
      grpc_addr = "127.0.0.1:9090"
      request_timeout = 100
      update_client_frequency = 10

      [Chain.Source.Tendermint.Fee]
        Amount = 0
        Denom = ""

      [Chain.Source.Tendermint.Key]
        Name = "validator"
        Password = "1234567890"
        mnemonic = "small pretty lock logic loud beef please boring space picnic essence also opera come roast pepper pumpkin vivid topple asset upon dismiss debris awful"
`
