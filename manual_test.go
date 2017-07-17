package exact_test

import (
	"encoding/json"
	"log"
	"net/url"
	"testing"
	"time"

	exact "github.com/tim-online/go-exactonline"
	"github.com/tim-online/go-exactonline/edm"
	"github.com/tim-online/go-exactonline/financial"
	"github.com/tim-online/go-exactonline/vat"
	"golang.org/x/oauth2"
)

func TestDing(t *testing.T) {
	// create oauth configuration
	oauthConfig := exact.NewOauth2Config()
	redirectURL, _ := url.Parse("https://mews-exact.omniboost.io/oauth2/auth?app=test")
	oauthConfig.SetRedirectURL(redirectURL)
	oauthConfig.ClientID = "72b356ea-3a4e-4541-b01b-6d7cc303a72c"
	oauthConfig.ClientSecret = "DTxWEEnu4FEz"

	token := &oauth2.Token{
		RefreshToken: "rrkN!IAAAADZ6zxTrheoEv1wMCqeQjuQ5iW0hGTYuvJhQDY7wFy6ysQAAAAFwmEQB0qvtQQkq2ubk1ht4LPEKaclUqVUfWU_s_SlcN9_9R5UmSMLfJbM6Yx5zEHFGfl0x4t3QNmvBzn59FS6IDiqd-UtnkVl8IhpDKCgVkTzs31XNSYE5B1GgPEeSu4aAR7yfLLRv176ayqPM30x6w5H7tiRL0dhNCDw78m8Ljz7t8VrQcQFpZdzJODxBldVlmPWWPwYgRsEYPTJ25WFZztkwZjA_7UaVwQBWcgvwUg",
	}

	refresh := false
	if refresh == true {
		var err error
		// baseURL, _ := url.Parse("https://mews-exact.omniboost.io/oauth2")
		// oauthConfig.SetBaseURL(baseURL)

		// webookSecret = KcjK6SLbNJJCBwee

		linkCallback := func(URL *url.URL) error {
			log.Println(URL)
			return nil
		}

		authorizationCallback := func(code string) error {
			log.Println(code)
			return nil
		}

		tokenCallback := func(token *oauth2.Token) error {
			log.Println(token)
			return nil
		}

		token, err = exact.GetNewOauth2Token(oauthConfig, linkCallback, authorizationCallback, tokenCallback)
		if err != nil {
			panic(err)
		}
	}

	// get http client with automatic oauth logic
	httpClient := oauthConfig.Client(oauth2.NoContext, token)

	divisionID := 974700
	client := exact.NewClient(httpClient, divisionID)
	client.SetDebug(true)

	func() {
		params := client.System.NewMeGetParams()
		resp, err := client.System.MeGet(params, nil)
		if err != nil {
			log.Fatal(err)
		}

		log.Fatal(resp)
	}()

	// params := client.FinancialTransaction.NewTransactionsGetParams()
	// params.Top.Set(2)
	// params.Select.Add("Type")
	// params.Select.Add("Status")
	// params.Select.Add("TransactionLines")
	// params.Expand.Add("TransactionLines")
	// resp, err := client.FinancialTransaction.TransactionsGet(1665186, params, nil)
	// if err != nil {
	// 	panic(err)
	// }

	// params := client.FinancialTransaction.NewBankEntriesGetParams()
	// params.Top.Set(1)
	// // params.Select.Add("BankEntryLines")
	// // params.Expand.Add("BankEntryLines")
	// resp, err := client.FinancialTransaction.BankEntriesGet(1665186, params, nil)
	// if err != nil {
	// 	panic(err)
	// }

	// params := client.Logistics.NewItemsGetParams()
	// params.Top.Set(1)
	// params.Select.Add("Description")
	// params.Expand.Add("BankEntryLines")
	// params.Filter.Set("startswith(Code, '8') eq true")
	// params.Filter.Set("Code ge 8000 and Code lt 9000")
	// params.Filter.Set("Type eq 24")
	// params.Filter.Set("ItemGroupDescription eq 'Diensten'")
	// resp, err := client.Logistics.ItemsGet(params, nil)
	// if err != nil {
	// 	panic(err)
	// }

	// params := client.SalesOrder.NewSalesOrderLinesGetParams()
	// // params.Top.Set(1)
	// params.Filter.Set("ItemDescription eq 'Trampoline'")
	// resp, err := client.SalesOrder.SalesOrderLinesGet(params, nil)
	// if err != nil {
	// 	panic(err)
	// }

	// account := func() crm.Account {
	// 	params := client.CRM.NewAccountsGetParams()
	// 	params.Top.Set(1)
	// 	params.Select.Add("ID")
	// 	params.Filter.Set("IsSupplier eq false")
	// 	resp, _ := client.CRM.AccountsGet(params, nil)
	// 	account := resp.Results[0]
	// 	return account
	// }()

	// body := client.SalesInvoice.NewSalesInvoicesPostBody()
	// body.DueDate.Time = time.Now()
	// body.OrderDate.Time = time.Now()
	// body.DeliverTo = account.ID
	// body.InvoiceTo = account.ID
	// body.OrderedBy = account.ID
	// body.Type = salesinvoice.InvoiceTypeInvoice

	// item := func() logistics.Item {
	// 	params := client.Logistics.NewItemsGetParams()
	// 	params.Top.Set(1)
	// 	resp, _ := client.Logistics.ItemsGet(params, nil)
	// 	item := resp.Results[0]
	// 	return item
	// }()

	// vatCode := func() vat.VatCode {
	// 	params := client.VAT.NewVatCodesGetParams()
	// 	params.Top.Set(1)
	// 	params.Filter.Set("Percentage eq 0.21")
	// 	params.Expand.Add("VATPercentages")
	// 	resp, _ := client.VAT.VatCodesGet(params, nil)
	// 	item := resp.Results[0]
	// 	return item
	// }()

	// invoiceLine := client.SalesInvoice.NewSalesInvoiceLine()
	// invoiceLine.Item = item.ID
	// invoiceLine.VATCode = vatCode.Code
	// invoiceLine.VATPercentage = vatCode.Percentage

	// body.SalesInvoiceLines = append(body.SalesInvoiceLines, *invoiceLine)
	// // params.Filter.Set("ItemDescription eq 'Trampoline'")
	// resp, err := client.SalesInvoice.SalesInvoicesPost(body, nil)
	// if err != nil {
	// 	panic(err)
	// }

	// params := client.SalesInvoice.NewSalesInvoiceLinesGetParams()
	// // params.Top.Set(1)
	// // params.Expand.Add("SalesInvoiceLines")
	// params.Filter.Set("ItemDescription eq 'Energy reep'")
	// resp, err := client.SalesInvoice.SalesInvoiceLinesGet(params, nil)
	// if err != nil {
	// 	panic(err)
	// }

	// params := client.CRM.NewAccountsGetParams()
	// params.Top.Set(1)
	// params.Select.Add("ID")
	// params.Expand.Add("BankAccounts")
	// params.Filter.Set("IsSupplier eq false")
	// resp, err := client.CRM.AccountsGet(params, nil)
	// if err != nil {
	// 	panic(err)
	// }

	// params := client.GeneralJournalEntry.NewGeneralJournalEntriesGetParams()
	// params.Top.Set(1)
	// // params.Select.Add("EntryID")
	// // params.Select.Add("EntryNumber")
	// // params.Select.Add("GeneralJournalEntryLines")
	// params.Filter.Set("EntryNumber eq 17900001 and JournalCode eq '90'")
	// params.Expand.Add("GeneralJournalEntryLines")
	// resp, err := client.GeneralJournalEntry.GeneralJournalEntriesGet(params, nil)
	// if err != nil {
	// 	panic(err)
	// }

	// client.SalesInvoice.SalesInvoice.List()
	// client.SalesInvoice.SalesInvoice.Create(newSalesInvoice)
	// client.SalesInvoice.SalesInvoice.Update(newSalesInvoice)
	// client.SalesInvoice.SalesInvoice.Delete()

	vatCode := func() vat.VatCode {
		params := client.VAT.NewVatCodesGetParams()
		params.Filter.Set("Percentage eq 0.21")
		params.Select.Add("GeneralJournalEntryLines")
		// params.Expand.Add("VATPercentages")
		resp, err := client.VAT.VatCodesGet(params, nil)
		if err != nil {
			log.Fatal(err)
		}

		item := resp.Results[0]
		return item
	}()

	glaccount := func() financial.GLAccount {
		params := client.Financial.NewGLAccountsGetParams()
		params.Top.Set(1)
		// params.Filter.Set("Percentage eq 0.21")
		// params.Expand.Add("VATPercentages")
		resp, _ := client.Financial.GLAccountsGet(params, nil)
		item := resp.Results[0]
		return item
	}()

	// params := client.General.Currencies.NewGetParams()
	// params.Top.Set(1)
	// currencies, err := client.General.Currencies.Get(params, nil)
	// log.Printf("%+v", currencies)
	// log.Println(err)
	// os.Exit(12)

	// service, resource, method/action
	service := client.GeneralJournalEntry
	body := service.NewGeneralJournalEntriesPostBody()
	body.Currency = "EUR"
	body.FinancialYear = 2017
	body.FinancialPeriod = 9
	body.JournalCode = "90"
	body.EntryNumber = 12

	line := client.GeneralJournalEntry.NewGeneralJournalEntryLine()
	line.AmountFC = 100.0
	line.AmountVATFC = 10.0
	line.VATCode = vatCode.Code
	line.VATPercentage = vatCode.Percentage
	line.Date.Time = time.Now()
	line.GLAccount = glaccount.ID
	line.Description = "TEST"
	body.GeneralJournalEntryLines = append(body.GeneralJournalEntryLines, line)

	line = client.GeneralJournalEntry.NewGeneralJournalEntryLine()
	line.AmountFC = -100.0
	line.AmountVATFC = -10.0
	line.VATCode = vatCode.Code
	line.VATPercentage = vatCode.Percentage
	line.Date = edm.DateTime{Time: time.Now()}
	line.GLAccount = glaccount.ID
	line.Description = "TEST"
	body.GeneralJournalEntryLines = append(body.GeneralJournalEntryLines, line)

	resp, err := client.GeneralJournalEntry.GeneralJournalEntriesPost(body, nil)
	if err != nil {
		panic(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println("----------------------------")
	log.Println(string(b))
	log.Println("----------------------------")
	// for _, d := range resp.Results {
	// 	log.Printf("%+v\n", d)
	// }
}
