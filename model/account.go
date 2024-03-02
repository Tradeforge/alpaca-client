package model

import (
	"time"
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
	ID            string `json:"id"`
	Status        string `json:"status"`
	AccountNumber string `json:"account_number"`
	AccountType   string `json:"account_type"`

	Contact        Contact         `json:"contact"`
	Currency       string          `json:"currency"`
	Identity       Identity        `json:"identity"`
	Disclosures    Disclosures     `json:"disclosures"`
	Agreements     []Agreement     `json:"agreements"`
	Documents      []Document      `json:"documents"`
	TrustedContact *TrustedContact `json:"trusted_contact"`
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
	AnnualIncomeMin   *int     `json:"annual_income_min"`
	AnnualIncomeMax   *int     `json:"annual_income_max"`
	LiquidNetWorthMin *int     `json:"liguid_net_worth_min"`
	LiquidNetWorthMax *int     `json:"liquid_net_worth_max"`
	TotalNetWorthMin  *int     `json:"total_net_worth_min"`
	TotalNetWorthMax  *int     `json:"total_net_worth_max"`

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
	ID string `path:"id,required"`
}

type GetAccountResponse struct {
	Account
}

type GetAccountHistoryParams struct {
	ID        string  `path:"id,required"`
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
