package webhooks

const (
	//possible keys
	BlockApplied        = Event("block.applied")
	BlockForged         = Event("block.forged")
	BlockReverted       = Event("block.reverted")
	DelegateRegistered  = Event("delegate.registered")
	DelegateResigned    = Event("delegate.resigned")
	ForgerFailed        = Event("forger.failed")
	ForgerMissing       = Event("forger.missing")
	ForgerStarted       = Event("forger.started")
	PeerAdded           = Event("peer.added")
	PeerRemoved         = Event("peer.removed")
	RoundCreated        = Event("round.created")
	TransactionApplied  = Event("transaction.applied")
	TransactionExpired  = Event("transaction.expired")
	TransactionForged   = Event("transaction.forged")
	TransactionReverted = Event("transaction.reverted")
	WalletSaved         = Event("wallet.saved")
	WalletCreatedCold   = Event("wallet.created.cold")

	Between    = Cond("between")
	Contains   = Cond("contains")
	Equal      = Cond("eq")
	Falsy      = Cond("falsy")
	Greater    = Cond("gt")
	GreaterEq  = Cond("gte")
	Lesser     = Cond("lt")
	LesserEq   = Cond("lte")
	NotEqual   = Cond("ne")
	NotBetween = Cond("not-between")
	Regexp     = Cond("regexp")
	Truthy     = Cond("truthy")
)

// Events are the possible events as described in https://docs.ark.io/api/webhooks/#create-a-webhook
type Event string

// Cond are the possible conditions for a webhook to trigger as described in https://docs.ark.io/api/webhooks/#create-a-webhook
type Cond string

// Condition is the total condition for a webhook to trigger
type Condition struct {
	Key       string `json:"key"`
	Condition string `json:"condition"`
	Value     string `json:"value"`
}

// NewCondition makes a new Condition
func NewCondition(key string, condition Cond, value string) Condition {
	return Condition{Key: key, Condition: string(condition), Value: value}
}
