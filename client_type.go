package client

type Keys struct {
	PKey string
	AKey string
	OKey string
	MKey string
}

type BResp struct {
	ID       string
	BlockNum uint32
	TrxNum   uint32
	Expired  bool
}

type PC_Options struct {
	Percent uint16
}

type PC_Vote struct {
	Weight int
}

type ArrTransfer struct {
	To      string
	Memo    string
	Ammount string
}

type ArrVote struct {
	User   string
	Weight int
}
