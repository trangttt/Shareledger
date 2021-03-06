package message

import (
	"encoding/json"

	sdk "bitbucket.org/shareringvn/cosmos-sdk/types"
	"github.com/sharering/shareledger/constants"
	types "github.com/sharering/shareledger/types"
	posTypes "github.com/sharering/shareledger/x/pos/type"
)

// MsgDelegate - struct for bonding transactions
type MsgDelegate struct {
	DelegatorAddr sdk.Address `json:"delegatorAddress"`
	ValidatorAddr sdk.Address `json:"validatorAddress"`
	Delegation    types.Coin  `json:"delegation"`
}

func NewMsgDelegate(delAddr sdk.Address, valAddr sdk.Address, delegation types.Coin) MsgDelegate {
	return MsgDelegate{
		DelegatorAddr: delAddr,
		ValidatorAddr: valAddr,
		Delegation:    delegation,
	}
}

var _ sdk.Msg = MsgDelegate{}

//nolint
func (msg MsgDelegate) Type() string { return constants.MESSAGE_POS }

func (msg MsgDelegate) GetSigners() []sdk.Address {
	return []sdk.Address{msg.DelegatorAddr}
}

// get the bytes for the message signer to sign on
func (msg MsgDelegate) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return b //sdk.MustSortJSON(b)
}

// quick validity check
func (msg MsgDelegate) ValidateBasic() sdk.Error {
	if msg.DelegatorAddr == nil {
		return posTypes.ErrNilDelegatorAddr(posTypes.DefaultCodespace)
	}
	if msg.ValidatorAddr == nil {
		return posTypes.ErrNilValidatorAddr(posTypes.DefaultCodespace)
	}
	if !(msg.Delegation.Amount.IsPositive()) {
		return posTypes.ErrBadDelegationAmount(posTypes.DefaultCodespace)
	}
	return nil
}
