package app

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authVestTypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	authzTypes "github.com/cosmos/cosmos-sdk/x/authz"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisisTypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distTypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidenceTypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	feeGrantTypes "github.com/cosmos/cosmos-sdk/x/feegrant"

	paramTypes "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	slashingTypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradeTypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Codec             codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}

func MakeEncodingConfig() EncodingConfig {
	amino := codec.NewLegacyAmino()
	interfaceRegistry := types.NewInterfaceRegistry()
	codec := codec.NewProtoCodec(interfaceRegistry)
	txCfg := tx.NewTxConfig(codec, tx.DefaultSignModes)

	encodingConfig := EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             codec,
		TxConfig:          txCfg,
		Amino:             amino,
	}
	std.RegisterLegacyAminoCodec(encodingConfig.Amino)
	std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	bankTypes.RegisterLegacyAminoCodec(encodingConfig.Amino)
	bankTypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	stakingTypes.RegisterLegacyAminoCodec(encodingConfig.Amino)
	stakingTypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	distTypes.RegisterLegacyAminoCodec(encodingConfig.Amino)
	distTypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	//VScode might shows error as it doesn't compile the files with tags (tag)
	RegisterCodecForGov(&encodingConfig)
	authzTypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	paramTypes.RegisterLegacyAminoCodec(encodingConfig.Amino)
	paramTypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	upgradeTypes.RegisterLegacyAminoCodec(encodingConfig.Amino)
	upgradeTypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	authTypes.RegisterLegacyAminoCodec(encodingConfig.Amino)
	authTypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	feeGrantTypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	authVestTypes.RegisterLegacyAminoCodec(encodingConfig.Amino)
	authVestTypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	crisisTypes.RegisterLegacyAminoCodec(encodingConfig.Amino)
	crisisTypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	evidenceTypes.RegisterLegacyAminoCodec(encodingConfig.Amino)
	evidenceTypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	slashingTypes.RegisterLegacyAminoCodec(encodingConfig.Amino)
	slashingTypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	return encodingConfig
}
