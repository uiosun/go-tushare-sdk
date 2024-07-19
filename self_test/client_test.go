package self_test

import (
	"errors"
	"fmt"
	"github.com/yushikuann/go-tushare-sdk/client"
	"github.com/yushikuann/go-tushare-sdk/config"
	"strings"
	"testing"
	"time"
)

var token = ""

func getToken() string {
	if token == "" {
		token = config.GetConfig("/..")["TS_TOKEN"]
	}

	return token
}

func TestClient(t *testing.T) {
	share := client.New(getToken())
	params := make(map[string]string)
	params["list_status"] = "L"
	var fields []string
	fieldStr := "ts_code,symbol,name,area,industry,market,list_date"
	fields = strings.Split(fieldStr, ",")
	data, _ := share.StockBasic(params, fields)

	fmt.Println(data.Data.Items)

	for _, item := range data.Data.Items[1:20] {
		fmt.Println(time.Parse("20060102", item[6].(string)))
	}
}

type FiledMapping struct {
	theType string
	trans   string
}

// 类型检查
func TypeCheck(fm []FiledMapping, items [][]interface{}) error {
	for _, item := range items {
		for i2, slot := range fm {
			var ok bool

			if len(fm) > len(item) {
				return errors.New(fmt.Sprintf("Data item length(%d) less than filed mapping(%d), and the value is %s", len(fm), len(item), item))
			}

			switch slot.theType {
			case "float64_and_nil":
				if item[i2] == nil {
					ok = true
				} else {
					_, ok = item[i2].(float64)
				}
			case "float64":
				_, ok = item[i2].(float64)
			case "string":
				_, ok = item[i2].(string)
			default:
				return errors.New(fmt.Sprintf("Noun type: %s", slot.theType))
			}
			if !ok {
				fmt.Println(item)
				return errors.New(fmt.Sprintf("No.%d: %s format is failed, origin value is %s, the symbol %s", i2, slot.trans, item[i2], item[0]))
			}
		}
	}

	return nil
}

func TestStkMins(t *testing.T) {
	share := client.New(getToken())
	params := make(map[string]string)
	params["ts_code"] = "000005.SZ"
	params["start_date"] = "2015-01-06 15:00:00"
	params["end_date"] = "2015-06-01 15:00:00"
	params["freq"] = "1min"
	params["limit"] = "5"
	data, err := share.StkMins(params, []string{})
	if err != nil {
		t.Fatal(err)
	}
	if len(data.Data.Items) != 5 {
		t.Fatalf("data count not has %d pieces", len(data.Data.Items))
	}

	if len(data.Data.Items) > 0 {
		mapping := []FiledMapping{
			{"string", "symbol"},
			{"string", "tradeDate"},
			{"float64", "open"},
			{"float64", "sClose"},
			{"float64", "high"},
			{"float64", "low"},
			{"float64", "vol"},
			{"float64", "amount"},
		}
		err := TypeCheck(mapping, data.Data.Items)
		if err != nil {
			t.Fatal(err)
		}
	}
}

// todo 字段不够，需要 Py 测试不存在的字段
func TestIncomeVip(t *testing.T) {
	share := client.New(getToken())
	params := make(map[string]string)
	params["start_date"] = "20240101"
	params["end_date"] = "20240401"
	data, err := share.IncomeVip(params, []string{})
	if err != nil {
		t.Fatal(err)
	}
	if len(data.Data.Items) == 0 {
		t.Fatalf("data count not has %d pieces", len(data.Data.Items))
	}

	mapping := []FiledMapping{
		{"string", "ts_code"},
		{"string", "ann_date"},
		{"string", "f_ann_date"},
		{"string", "end_date"},
		{"string", "report_type"},
		{"string", "comp_type"},
		{"string", "end_type"},
		{"float64", "basic_eps"},
		{"float64", "diluted_eps"},
		{"float64", "total_revenue"},
		{"float64", "revenue"},
		{"float64", "int_income"},
		{"float64_and_nil", "prem_earned"},
		{"float64", "comm_income"},
		{"float64", "n_commis_income"},
		{"float64", "n_oth_income"},
		{"float64", "n_oth_b_income"},
		{"float64_and_nil", "prem_income"},
		{"float64_and_nil", "out_prem"},
		{"float64_and_nil", "une_prem_reser"},
		{"float64_and_nil", "reins_income"},
		{"float64_and_nil", "n_sec_tb_income"},
		{"float64_and_nil", "n_sec_uw_income"},
		{"float64_and_nil", "n_asset_mg_income"},
		{"float64", "oth_b_income"},
		{"float64", "fv_value_chg_gain"},
		{"float64", "invest_income"},
		{"float64_and_nil", "ass_invest_income"},
		{"float64", "forex_gain"},
		{"float64", "total_cogs"},
		{"float64_and_nil", "oper_cost"},
		{"float64", "int_exp"},
		{"float64", "comm_exp"},
		{"float64", "biz_tax_surchg"},
		{"float64_and_nil", "sell_exp"},
		{"float64", "admin_exp"},
		{"float64_and_nil", "fin_exp"},
		{"float64_and_nil", "assets_impair_loss"},
		{"float64_and_nil", "prem_refund"},
		{"float64_and_nil", "compens_payout"},
		{"float64_and_nil", "reser_insur_liab"},
		{"float64_and_nil", "div_payt"},
		{"float64_and_nil", "reins_exp"},
		{"float64", "oper_exp"},
		{"float64_and_nil", "compens_payout_refu"},
		{"float64_and_nil", "insur_reser_refu"},
		{"float64_and_nil", "reins_cost_refund"},
		{"float64_and_nil", "other_bus_cost"},
		{"float64", "operate_profit"},
		{"float64", "non_oper_income"},
		{"float64", "non_oper_exp"},
		{"float64_and_nil", "nca_disploss"},
		{"float64", "total_profit"},
		{"float64", "income_tax"},
		{"float64", "n_income"},
		{"float64", "n_income_attr_p"},
		{"float64_and_nil", "minority_gain"},
		{"float64", "oth_compr_income"},
		{"float64", "t_compr_income"},
		{"float64", "compr_inc_attr_p"},
		{"float64_and_nil", "compr_inc_attr_m_s"},
		{"float64_and_nil", "ebit"},
		{"float64_and_nil", "ebitda"},
		{"float64_and_nil", "insurance_exp"},
		{"float64_and_nil", "undist_profit"},
		{"float64_and_nil", "distable_profit"},
		{"float64_and_nil", "rd_exp"},
		{"float64_and_nil", "fin_exp_int_exp"},
		{"float64_and_nil", "fin_exp_int_inc"},
		{"float64_and_nil", "transfer_surplus_rese"},
		{"float64_and_nil", "transfer_housing_imprest"},
		{"float64_and_nil", "transfer_oth"},
		{"float64_and_nil", "adj_lossgain"},
		{"float64_and_nil", "withdra_legal_surplus"},
		{"float64_and_nil", "withdra_legal_pubfund"},
		{"float64_and_nil", "withdra_biz_devfund"},
		{"float64_and_nil", "withdra_rese_fund"},
		{"float64_and_nil", "withdra_oth_ersu"},
		{"float64_and_nil", "workers_welfare"},
		{"float64_and_nil", "distr_profit_shrhder"},
		{"float64_and_nil", "prfshare_payable_dvd"},
		{"float64_and_nil", "comshare_payable_dvd"},
		{"float64_and_nil", "capit_comstock_div"},
		{"float64", "net_after_nr_lp_correct"},
		{"string", "credit_impa_loss"},
		{"float64", "net_expo_hedging_benefits"},
		{"float64", "oth_impair_loss_assets"},
		{"float64", "total_opcost"},
		{"float64", "amodcost_fin_assets"},
		{"float64", "oth_income"},
		{"float64", "asset_disp_income"},
		{"float64", "continued_net_profit"},
		{"float64", "end_net_profit"},
		{"str", "update_flag"},
	}
	err = TypeCheck(mapping, data.Data.Items)
	if err != nil {
		t.Fatal(err)
	}
}

// todo 字段不够，需要 Py 测试不存在的字段
func TestIncome(t *testing.T) {
	share := client.New(getToken())
	params := make(map[string]string)
	params["ts_code"] = "002936.SZ"
	data, err := share.Income(params, []string{})
	if err != nil {
		t.Fatal(err)
	}
	if len(data.Data.Items) == 0 {
		t.Fatalf("data count not has %d pieces", len(data.Data.Items))
	}

	mapping := []FiledMapping{
		{"string", "ts_code"},
		{"string", "ann_date"},
		{"string", "f_ann_date"},
		{"string", "end_date"},
		{"string", "report_type"},
		{"string", "comp_type"},
		{"string", "end_type"},
		{"float64", "basic_eps"},
		{"float64", "diluted_eps"},
		{"float64", "total_revenue"},
		{"float64", "revenue"},
		{"float64", "int_income"},
		{"float64_and_nil", "prem_earned"},
		{"float64", "comm_income"},
		{"float64", "n_commis_income"},
		{"float64", "n_oth_income"},
		{"float64", "n_oth_b_income"},
		{"float64_and_nil", "prem_income"},
		{"float64_and_nil", "out_prem"},
		{"float64_and_nil", "une_prem_reser"},
		{"float64_and_nil", "reins_income"},
		{"float64_and_nil", "n_sec_tb_income"},
		{"float64_and_nil", "n_sec_uw_income"},
		{"float64_and_nil", "n_asset_mg_income"},
		{"float64", "oth_b_income"},
		{"float64", "fv_value_chg_gain"},
		{"float64", "invest_income"},
		{"float64_and_nil", "ass_invest_income"},
		{"float64", "forex_gain"},
		{"float64", "total_cogs"},
		{"float64_and_nil", "oper_cost"},
		{"float64", "int_exp"},
		{"float64", "comm_exp"},
		{"float64", "biz_tax_surchg"},
		{"float64_and_nil", "sell_exp"},
		{"float64", "admin_exp"},
		{"float64_and_nil", "fin_exp"},
		{"float64_and_nil", "assets_impair_loss"},
		{"float64_and_nil", "prem_refund"},
		{"float64_and_nil", "compens_payout"},
		{"float64_and_nil", "reser_insur_liab"},
		{"float64_and_nil", "div_payt"},
		{"float64_and_nil", "reins_exp"},
		{"float64", "oper_exp"},
		{"float64_and_nil", "compens_payout_refu"},
		{"float64_and_nil", "insur_reser_refu"},
		{"float64_and_nil", "reins_cost_refund"},
		{"float64_and_nil", "other_bus_cost"},
		{"float64", "operate_profit"},
		{"float64", "non_oper_income"},
		{"float64", "non_oper_exp"},
		{"float64_and_nil", "nca_disploss"},
		{"float64", "total_profit"},
		{"float64", "income_tax"},
		{"float64", "n_income"},
		{"float64", "n_income_attr_p"},
		{"float64_and_nil", "minority_gain"},
		{"float64", "oth_compr_income"},
		{"float64", "t_compr_income"},
		{"float64", "compr_inc_attr_p"},
		{"float64_and_nil", "compr_inc_attr_m_s"},
		{"float64_and_nil", "ebit"},
		{"float64_and_nil", "ebitda"},
		{"float64_and_nil", "insurance_exp"},
		{"float64_and_nil", "undist_profit"},
		{"float64_and_nil", "distable_profit"},
		{"float64_and_nil", "rd_exp"},
		{"float64_and_nil", "fin_exp_int_exp"},
		{"float64_and_nil", "fin_exp_int_inc"},
		{"float64_and_nil", "transfer_surplus_rese"},
		{"float64_and_nil", "transfer_housing_imprest"},
		{"float64_and_nil", "transfer_oth"},
		{"float64_and_nil", "adj_lossgain"},
		{"float64_and_nil", "withdra_legal_surplus"},
		{"float64_and_nil", "withdra_legal_pubfund"},
		{"float64_and_nil", "withdra_biz_devfund"},
		{"float64_and_nil", "withdra_rese_fund"},
		{"float64_and_nil", "withdra_oth_ersu"},
		{"float64_and_nil", "workers_welfare"},
		{"float64_and_nil", "distr_profit_shrhder"},
		{"float64_and_nil", "prfshare_payable_dvd"},
		{"float64_and_nil", "comshare_payable_dvd"},
		{"float64_and_nil", "capit_comstock_div"},
		{"float64", "net_after_nr_lp_correct"},
		{"string", "credit_impa_loss"},
		{"float64", "net_expo_hedging_benefits"},
		{"float64", "oth_impair_loss_assets"},
		{"float64", "total_opcost"},
		{"float64", "amodcost_fin_assets"},
		{"float64", "oth_income"},
		{"float64", "asset_disp_income"},
		{"float64", "continued_net_profit"},
		{"float64", "end_net_profit"},
		{"str", "update_flag"},
	}
	err = TypeCheck(mapping, data.Data.Items)
	if err != nil {
		t.Fatal(err)
	}
}
