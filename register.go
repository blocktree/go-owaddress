package owaddress

import (
	"github.com/star001007/go-owaddress/address"
	"github.com/star001007/go-owaddress/coins/ada"
	"github.com/star001007/go-owaddress/coins/ae"
	"github.com/star001007/go-owaddress/coins/alc"
	"github.com/star001007/go-owaddress/coins/ark"
	"github.com/star001007/go-owaddress/coins/atom"
	"github.com/star001007/go-owaddress/coins/avax"
	"github.com/star001007/go-owaddress/coins/bbc"
	"github.com/star001007/go-owaddress/coins/bch"
	"github.com/star001007/go-owaddress/coins/beth"
	"github.com/star001007/go-owaddress/coins/bnb"
	"github.com/star001007/go-owaddress/coins/bsc"
	"github.com/star001007/go-owaddress/coins/bsv"
	"github.com/star001007/go-owaddress/coins/btc"
	"github.com/star001007/go-owaddress/coins/btx"
	"github.com/star001007/go-owaddress/coins/cxc"
	"github.com/star001007/go-owaddress/coins/dgb"
	"github.com/star001007/go-owaddress/coins/dsc"
	"github.com/star001007/go-owaddress/coins/ela"
	"github.com/star001007/go-owaddress/coins/eos"
	"github.com/star001007/go-owaddress/coins/eth"
	"github.com/star001007/go-owaddress/coins/etp"
	"github.com/star001007/go-owaddress/coins/eva"
	"github.com/star001007/go-owaddress/coins/fac"
	"github.com/star001007/go-owaddress/coins/fiii"
	"github.com/star001007/go-owaddress/coins/ftm"
	"github.com/star001007/go-owaddress/coins/g50"
	"github.com/star001007/go-owaddress/coins/hc"
	"github.com/star001007/go-owaddress/coins/hns"
	"github.com/star001007/go-owaddress/coins/hss"
	"github.com/star001007/go-owaddress/coins/ilc"
	"github.com/star001007/go-owaddress/coins/kpg"
	"github.com/star001007/go-owaddress/coins/ltc"
	"github.com/star001007/go-owaddress/coins/macc"
	"github.com/star001007/go-owaddress/coins/moac"
	"github.com/star001007/go-owaddress/coins/nas"
	"github.com/star001007/go-owaddress/coins/nhss"
	"github.com/star001007/go-owaddress/coins/ntn"
	"github.com/star001007/go-owaddress/coins/nuls2"
	"github.com/star001007/go-owaddress/coins/ont"
	"github.com/star001007/go-owaddress/coins/pb"
	"github.com/star001007/go-owaddress/coins/pess"
	"github.com/star001007/go-owaddress/coins/qtum"
	"github.com/star001007/go-owaddress/coins/rcp"
	"github.com/star001007/go-owaddress/coins/sgu"
	"github.com/star001007/go-owaddress/coins/sinoc"
	"github.com/star001007/go-owaddress/coins/sol"
	"github.com/star001007/go-owaddress/coins/tron"
	truechain "github.com/star001007/go-owaddress/coins/true"
	"github.com/star001007/go-owaddress/coins/trx"
	"github.com/star001007/go-owaddress/coins/tv"
	"github.com/star001007/go-owaddress/coins/vas"
	"github.com/star001007/go-owaddress/coins/vcc"
	"github.com/star001007/go-owaddress/coins/vds"
	"github.com/star001007/go-owaddress/coins/vlx"
	"github.com/star001007/go-owaddress/coins/vsys"
	"github.com/star001007/go-owaddress/coins/wicc"
	"github.com/star001007/go-owaddress/coins/xif"
	"github.com/star001007/go-owaddress/coins/xrp"
	"github.com/star001007/go-owaddress/coins/xvg"
	"github.com/star001007/go-owaddress/coins/xwc"
	"github.com/star001007/go-owaddress/coins/zen"
)

var AddressVerifierRegistry = make(map[string]address.AddressVerifier)

func RegisterAddressVerify(verifier address.AddressVerifier, coin string) {
	AddressVerifierRegistry[coin] = verifier
}

func init() {
	RegisterAddressVerify(btc.DefaultStruct, btc.CoinName)
	RegisterAddressVerify(ltc.DefaultStruct, ltc.CoinName)
	RegisterAddressVerify(qtum.DefaultStruct, qtum.CoinName)
	RegisterAddressVerify(ont.DefaultStruct, ont.CoinName)
	RegisterAddressVerify(atom.DefaultStruct, atom.CoinName)
	RegisterAddressVerify(xrp.DefaultStruct, xrp.CoinName)
	RegisterAddressVerify(trx.DefaultStruct, trx.CoinName)
	RegisterAddressVerify(wicc.DefaultStruct, wicc.CoinName)
	RegisterAddressVerify(hc.DefaultStruct, hc.CoinName)
	RegisterAddressVerify(bnb.DefaultStruct, bnb.CoinName)
	RegisterAddressVerify(vsys.DefaultStruct, vsys.CoinName)
	RegisterAddressVerify(nas.DefaultStruct, nas.CoinName)
	RegisterAddressVerify(ela.DefaultStruct, ela.CoinName)
	RegisterAddressVerify(tv.DefaultStruct, tv.CoinName)
	RegisterAddressVerify(moac.DefaultStruct, moac.CoinName)
	RegisterAddressVerify(dsc.DefaultStruct, dsc.CoinName)
	RegisterAddressVerify(fiii.DefaultStruct, fiii.CoinName)
	RegisterAddressVerify(vds.DefaultStruct, vds.CoinName)
	RegisterAddressVerify(hss.DefaultStruct, hss.CoinName)
	RegisterAddressVerify(vlx.DefaultStruct, vlx.CoinName)
	RegisterAddressVerify(btx.DefaultStruct, btx.CoinName)
	RegisterAddressVerify(cxc.DefaultStruct, cxc.CoinName)
	RegisterAddressVerify(xvg.DefaultStruct, xvg.CoinName)
	RegisterAddressVerify(zen.DefaultStruct, zen.CoinName)
	RegisterAddressVerify(dgb.DefaultStruct, dgb.CoinName)
	RegisterAddressVerify(alc.DefaultStruct, alc.CoinName)
	RegisterAddressVerify(eth.DefaultStruct, eth.CoinName)
	RegisterAddressVerify(pess.DefaultStruct, pess.CoinName)
	RegisterAddressVerify(vcc.DefaultStruct, vcc.CoinName)
	RegisterAddressVerify(truechain.DefaultStruct, truechain.CoinName)
	RegisterAddressVerify(g50.DefaultStruct, g50.CoinName)
	RegisterAddressVerify(sinoc.DefaultStruct, sinoc.CoinName)
	RegisterAddressVerify(ae.DefaultStruct, ae.CoinName)
	RegisterAddressVerify(macc.DefaultStruct, macc.CoinName)
	RegisterAddressVerify(ntn.DefaultStruct, ntn.CoinName)
	RegisterAddressVerify(beth.DefaultStruct, beth.CoinName)
	RegisterAddressVerify(xwc.DefaultStruct, xwc.CoinName)
	RegisterAddressVerify(vas.DefaultStruct, vas.CoinName)
	RegisterAddressVerify(etp.DefaultStruct, etp.CoinName)
	RegisterAddressVerify(nuls2.DefaultStruct, nuls2.CoinName)
	RegisterAddressVerify(bch.DefaultStruct, bch.CoinName)
	RegisterAddressVerify(bsv.DefaultStruct, bsv.CoinName)
	RegisterAddressVerify(bbc.DefaultStruct, bbc.CoinName)
	RegisterAddressVerify(sgu.DefaultStruct, sgu.CoinName)
	RegisterAddressVerify(eva.DefaultStruct, eva.CoinName)
	RegisterAddressVerify(rcp.DefaultStruct, rcp.CoinName)
	RegisterAddressVerify(pb.DefaultStruct, pb.CoinName)
	RegisterAddressVerify(ark.DefaultStruct, ark.CoinName)
	RegisterAddressVerify(fac.DefaultStruct, fac.CoinName)
	RegisterAddressVerify(nhss.DefaultStruct, nhss.CoinName)
	RegisterAddressVerify(ilc.DefaultStruct, ilc.CoinName)
	RegisterAddressVerify(xif.DefaultStruct, xif.CoinNameXIF)
	RegisterAddressVerify(xif.DefaultStruct, xif.CoinNameAUSD)
	RegisterAddressVerify(hns.DefaultStruct, hns.CoinName)
	RegisterAddressVerify(kpg.DefaultStruct, kpg.CoinName)
	RegisterAddressVerify(eos.DefaultStruct, eos.CoinName)
	RegisterAddressVerify(tron.DefaultStruct, tron.CoinName)
	RegisterAddressVerify(bsc.DefaultStruct, bsc.CoinName)
	RegisterAddressVerify(ftm.DefaultStruct, ftm.CoinName)
	RegisterAddressVerify(avax.DefaultStruct, avax.CoinName)
	RegisterAddressVerify(ada.DefaultStruct, ada.CoinName)
	RegisterAddressVerify(sol.DefaultStruct, sol.CoinName)
}
