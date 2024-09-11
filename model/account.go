package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateAccountRequest struct {
	Contact        Contact         `json:"contact"`
	Currency       string          `json:"currency"`
	Identity       Identity        `json:"identity"`
	Agreements     []Agreement     `json:"agreements"`
	Disclosures    Disclosures     `json:"disclosures"`
	Documents      []Document      `json:"documents"`
	TrustedContact *TrustedContact `json:"trusted_contact"`
}

type Account struct {
	ID            uuid.UUID     `json:"id"`
	Status        AccountStatus `json:"status"`
	AccountNumber string        `json:"account_number"`
	AccountType   string        `json:"account_type"`

	Contact        Contact         `json:"contact"`
	Currency       string          `json:"currency"`
	Identity       Identity        `json:"identity"`
	Disclosures    Disclosures     `json:"disclosures"`
	Agreements     []Agreement     `json:"agreements"`
	Documents      []Document      `json:"documents"`
	TrustedContact *TrustedContact `json:"trusted_contact"`

	BalanceUSD AccountBalanceUSD `json:"usd"`
}

type AccountBalanceUSD struct {
	BuyingPower               decimal.Decimal `json:"buying_power"`
	RegtBuyingPower           decimal.Decimal `json:"regt_buying_power"`
	DaytradingBuyingPower     decimal.Decimal `json:"daytrading_buying_power"`
	OptionsBuyingPower        decimal.Decimal `json:"options_buying_power"`
	Cash                      decimal.Decimal `json:"cash"`
	CashWithdrawable          decimal.Decimal `json:"cash_withdrawable"`
	CashTransferable          decimal.Decimal `json:"cash_transferable"`
	PendingTransferOut        decimal.Decimal `json:"pending_transfer_out"`
	PortfolioValue            decimal.Decimal `json:"portfolio_value"`
	Equity                    decimal.Decimal `json:"equity"`
	LongMarketValue           decimal.Decimal `json:"long_market_value"`
	ShortMarketValue          decimal.Decimal `json:"short_market_value"`
	InitialMargin             decimal.Decimal `json:"initial_margin"`
	MaintenanceMargin         decimal.Decimal `json:"maintenance_margin"`
	LastMaintenanceMargin     decimal.Decimal `json:"last_maintenance_margin"`
	Sma                       decimal.Decimal `json:"sma"`
	LastEquity                decimal.Decimal `json:"last_equity"`
	LastLongMarketValue       decimal.Decimal `json:"last_long_market_value"`
	LastShortMarketValue      decimal.Decimal `json:"last_short_market_value"`
	LastCash                  decimal.Decimal `json:"last_cash"`
	LastBuyingPower           decimal.Decimal `json:"last_buying_power"`
	LastRegtBuyingPower       decimal.Decimal `json:"last_regt_buying_power"`
	LastDaytradingBuyingPower decimal.Decimal `json:"last_daytrading_buying_power"`
	LastOptionsBuyingPower    decimal.Decimal `json:"last_options_buying_power"`
}

type AccountStatus string

const (
	// AccountStatusInactive represents an account that is not set to trade given asset.
	AccountStatusInactive AccountStatus = "INACTIVE"
	// AccountStatusOnboarding represents an application that is expected for this user, but has not been submitted yet.
	AccountStatusOnboarding AccountStatus = "ONBOARDING"
	// AccountStatusSubmitted represents an application that has been submitted and in process.
	AccountStatusSubmitted AccountStatus = "SUBMITTED"
	// AccountStatusSubmissionFailed represents a failure on submission.
	AccountStatusSubmissionFailed AccountStatus = "SUBMISSION_FAILED"
	// AccountStatusActionRequired represents an application that requires manual action.
	AccountStatusActionRequired AccountStatus = "ACTION_REQUIRED"
	// AccountStatusAccountUpdated represents when an account has been modified by the user.
	AccountStatusAccountUpdated AccountStatus = "ACCOUNT_UPDATED"
	// AccountStatusApprovalPending represents the initial value. The application approval process is in process.
	AccountStatusApprovalPending AccountStatus = "APPROVAL_PENDING"
	// AccountStatusApproved represents an account application that has been approved, and waiting to be ACTIVE.
	AccountStatusApproved AccountStatus = "APPROVED"
	// AccountStatusRejected represents an account application that is rejected for some reason.
	AccountStatusRejected AccountStatus = "REJECTED"
	// AccountStatusActive represents an account that is fully active. Trading and funding are processed under this status.
	AccountStatusActive AccountStatus = "ACTIVE"
	// AccountStatusAccountClosed represents an account that is closed.
	AccountStatusAccountClosed AccountStatus = "ACCOUNT_CLOSED"
)

type AccountTradingDetails struct {
	AccountID                 uuid.UUID       `json:"id"`
	AccountNumber             string          `json:"account_number"`
	Status                    string          `json:"status"`
	Currency                  string          `json:"currency"`
	BuyingPower               decimal.Decimal `json:"buying_power"`
	RegtBuyingPower           decimal.Decimal `json:"regt_buying_power"`
	DaytradingBuyingPower     decimal.Decimal `json:"daytrading_buying_power"`
	OptionsBuyingPower        decimal.Decimal `json:"options_buying_power"`
	Cash                      decimal.Decimal `json:"cash"`
	CashWithdrawable          decimal.Decimal `json:"cash_withdrawable"`
	CashTransferable          decimal.Decimal `json:"cash_transferable"`
	PendingTransferOut        decimal.Decimal `json:"pending_transfer_out"`
	PortfolioValue            decimal.Decimal `json:"portfolio_value"`
	PatternDayTrader          bool            `json:"pattern_day_trader"`
	TradingBlocked            bool            `json:"trading_blocked"`
	TransfersBlocked          bool            `json:"transfers_blocked"`
	AccountBlocked            bool            `json:"account_blocked"`
	CreatedAt                 time.Time       `json:"created_at"`
	TradeSuspendedByUser      bool            `json:"trade_suspended_by_user"`
	Multiplier                string          `json:"multiplier"`
	ShortingEnabled           bool            `json:"shorting_enabled"`
	Equity                    decimal.Decimal `json:"equity"`
	LastEquity                decimal.Decimal `json:"last_equity"`
	LongMarketValue           decimal.Decimal `json:"long_market_value"`
	ShortMarketValue          decimal.Decimal `json:"short_market_value"`
	InitialMargin             decimal.Decimal `json:"initial_margin"`
	MaintenanceMargin         decimal.Decimal `json:"maintenance_margin"`
	LastMaintenanceMargin     decimal.Decimal `json:"last_maintenance_margin"`
	Sma                       decimal.Decimal `json:"sma"`
	DaytradeCount             int             `json:"daytrade_count"`
	BalanceAsof               string          `json:"balance_asof"`
	PreviousClose             time.Time       `json:"previous_close"`
	LastLongMarketValue       decimal.Decimal `json:"last_long_market_value"`
	LastShortMarketValue      decimal.Decimal `json:"last_short_market_value"`
	LastCash                  decimal.Decimal `json:"last_cash"`
	LastInitialMargin         decimal.Decimal `json:"last_initial_margin"`
	LastRegtBuyingPower       decimal.Decimal `json:"last_regt_buying_power"`
	LastDaytradingBuyingPower decimal.Decimal `json:"last_daytrading_buying_power"`
	LastOptionsBuyingPower    decimal.Decimal `json:"last_options_buying_power"`
	LastBuyingPower           decimal.Decimal `json:"last_buying_power"`
	LastDaytradeCount         int             `json:"last_daytrade_count"`
	ClearingBroker            string          `json:"clearing_broker"`
	OptionsApprovedLevel      int             `json:"options_approved_level"`
	OptionsTradingLevel       int             `json:"options_trading_level"`
	IntradayAdjustments       decimal.Decimal `json:"intraday_adjustments"`
	PendingRegTafFees         decimal.Decimal `json:"pending_reg_taf_fees"`
}

type Contact struct {
	EmailAddress  string   `json:"email_address"`
	PhoneNumber   string   `json:"phone_number"`
	StreetAddress []string `json:"street_address"`
	Unit          *string  `json:"unit"`
	City          *string  `json:"city"`
	State         *string  `json:"state"`
	PostalCode    *string  `json:"postal_code"`
	Country       *string  `json:"country"`
}

type TrustedContact struct {
	GivenName     string    `json:"given_name"`
	FamilyName    string    `json:"family_name"`
	EmailAddress  *string   `json:"email_address"`
	PhoneNumber   *string   `json:"phone_number"`
	StreetAddress *[]string `json:"street_address"`
	Unit          *string   `json:"unit"`
	City          *string   `json:"city"`
	State         *string   `json:"state"`
	PostalCode    *string   `json:"postal_code"`
	Country       *string   `json:"country"`
}

type Identity struct {
	GivenName             string  `json:"given_name"`
	FamilyName            string  `json:"family_name"`
	MiddleName            *string `json:"middle_name"`
	DateOfBirth           string  `json:"date_of_birth"`
	TaxID                 *string `json:"tax_id"`
	TaxIDType             *string `json:"tax_id_type"`
	CountryOfCitizenship  *string `json:"country_of_citizenship"`
	CountryOfBirth        *string `json:"country_of_birth"`
	CountryOfTaxResidence string  `json:"country_of_tax_residence"`

	FundingSource     []string `json:"funding_source"`
	AnnualIncomeMin   *int64   `json:"annual_income_min"`
	AnnualIncomeMax   *int64   `json:"annual_income_max"`
	LiquidNetWorthMin *int64   `json:"liguid_net_worth_min"`
	LiquidNetWorthMax *int64   `json:"liquid_net_worth_max"`
	TotalNetWorthMin  *int64   `json:"total_net_worth_min"`
	TotalNetWorthMax  *int64   `json:"total_net_worth_max"`

	VisaType               *string    `json:"visa_type"`
	VisaExpirationDate     *time.Time `json:"visa_expiration_date"`
	DateOfDepartureFromUsa *time.Time `json:"date_of_departure_from_usa"`
	PermanentResident      *bool      `json:"permanent_resident"`
}

type Agreement struct {
	Agreement string    `json:"agreement"`
	SignedAt  time.Time `json:"signed_at"`
	IPAddress string    `json:"ip_address"`
	Revision  *string   `json:"revision"`
}

type Disclosures struct {
	EmploymentStatus            *string             `json:"employment_status"`
	EmployerName                *string             `json:"empoyer_name"`
	EmployerAddress             *string             `json:"employer_address"`
	EmploymentPosition          *string             `json:"employment_position"`
	IsControlPerson             bool                `json:"is_control_person"`
	IsAffiliatedExchangeOrFinra bool                `json:"is_affiliated_exchange_or_finra"`
	IsAffiliatedExchangeOrIiroc bool                `json:"is_affiliated_exchange_or_iiroc"`
	IsPoliticallyExposed        bool                `json:"is_politically_exposed"`
	ImmediateFamilyExposed      bool                `json:"immediate_family_exposed"`
	Context                     []DisclosureContext `json:"context"`
}

type DisclosureContext struct {
	ContextType            string  `json:"context_type"`
	CompanyName            *string `json:"company_name"`
	CompanyStreetAddress   *string `json:"company_street_address"`
	CompanyCity            *string `json:"company_city"`
	CompanyState           *string `json:"company_state"`
	CompanyCountry         *string `json:"company_country"`
	CompanyComplianceEmail *string `json:"company_compliance_email"`
	GivenName              *string `json:"given_name"`
	FamilyName             *string `json:"family_name"`
}

type CreateAccountResponse struct {
	Account
}

type ListAccountEntity = string

var (
	ListAccountEntityContact        ListAccountEntity = "contact"
	ListAccountEntityIdentity       ListAccountEntity = "identity"
	ListAccountEntityAgreements     ListAccountEntity = "agreements"
	ListAccountEntityDisclosures    ListAccountEntity = "disclosures"
	ListAccountEntityDocuments      ListAccountEntity = "documents"
	ListAccountEntityTrustedContact ListAccountEntity = "trusted_contact"
)

type ListAccountsParams struct {
	Query    string `query:"query,omitempty"`
	Entities string `query:"entities,omitempty"`
}

type ListAccountsResponse = []Account

type GetAccountParams struct {
	AccountID string `path:"account_id,required"`
}

type GetAccountResponse struct {
	Account
}

type GetAccountTradingDetailsParams struct {
	AccountID string `path:"account_id,required"`
}

type GetAccountTradingDetailsResponse struct {
	AccountTradingDetails
}

type GetAccountHistoryParams struct {
	AccountID string  `path:"account_id,required"`
	Period    string  `query:"period,required"`
	Timeframe *string `query:"timeframe,omitempty"`
}

type GetAccountHistoryResponse struct {
	Timestamp     []int64   `json:"timestamp"`
	Equity        []float64 `json:"equity"`
	ProfitLoss    []float64 `json:"profit_loss"`
	ProfitLossPct []float64 `json:"profit_loss_pct"`
	BaseValue     float64   `json:"base_value"`
	Timeframe     string    `json:"timeframe"`
}

type KYCResults struct {
	Reject                string `json:"reject"`
	Accept                string `json:"accept"`
	Indeterminate         string `json:"indeterminate"`
	AdditionalInformation string `json:"additional_information"`
	Summary               string `json:"summary"`
}

type AccountAdminConfigurations struct {
	RestrictToLiquidationReasons struct {
		PatternDayTrading     bool `json:"pattern_day_trading"`
		AchReturn             bool `json:"ach_return"`
		PositionToEquityRatio bool `json:"position_to_equity_ratio"`
		Unspecified           bool `json:"unspecified"`
	} `json:"restrict_to_liquidation_reasons"`
	OutgoingTransfersBlocked bool   `json:"outgoing_transfers_blocked"`
	IncomingTransfersBlocked bool   `json:"incoming_transfers_blocked"`
	DisableShorting          bool   `json:"disable_shorting"`
	DisableFractional        bool   `json:"disable_fractional"`
	DisableCrypto            bool   `json:"disable_crypto"`
	DisableDayTrading        bool   `json:"disable_day_trading"`
	MaxMarginMultiplier      int    `json:"max_margin_multiplier"`
	AcctDailyTransferLimit   string `json:"acct_daily_transfer_limit"`
}

type GetOnfidoSDKTokenParams struct {
	AccountID string  `path:"account_id,required"`
	Referrer  *string `query:"referrer,omitempty"`
	Platform  *string `query:"platform,omitempty"`
}

type GetOnfidoSDKTokenResponse struct {
	Token string `json:"token"`
}

type UpdateOnfidoSDKOutcomeParams struct {
	AccountID string `path:"account_id,required"`
}

type UpdateOnfidoSDKOutcomeRequest struct {
	Token   string  `json:"token"`
	Outcome string  `json:"outcome"`
	Reason  *string `json:"reason"`
}
