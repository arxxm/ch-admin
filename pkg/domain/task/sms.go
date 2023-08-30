package task

import (
	"time"

	"github.com/doug-martin/goqu/v9"
)

type Sms struct {
	ID                    int64     `json:"id"`
	Status                string    `json:"status"`
	Scheme                string    `json:"scheme"`
	SenderNumber          string    `json:"sender_number"`
	SenderUserId          int64     `json:"sender_user_id"`
	SenderPrice           float64   `json:"sender_price"`
	SenderOperatorId      int64     `json:"sender_operator_id"`
	ReceiverNumber        string    `json:"receiver_number"`
	ReceiverUserId        int64     `json:"receiver_user_id"`
	ReceiverPrice         float64   `json:"receiver_price"`
	ReceiverOperatorId    int64     `json:"receiver_operator_id"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	ClientCompanyId       string    `json:"client_company_id"`
	ClientTaskId          string    `json:"client_task_id"`
	ClientTaskName        string    `json:"client_task_name"`
	ParentId              int64     `json:"parent_id"`
	SmppServerId          string    `json:"smpp_server_id"`
	Text                  string    `json:"text"`
	Country               string    `json:"country"`
	Operator              string    `json:"operator"`
	SecretCode            int64     `json:"secret_code"`
	SenderNumberId        int64     `json:"sender_number_id"`
	ReceiverNumberId      int64     `json:"receiver_number_id"`
	Timeout               int64     `json:"timeout"`
	Analyzing             bool      `json:"analyzing"`
	SmppAvatar            string    `json:"smpp_avatar"`
	IsDeleted             bool      `json:"is_deleted"`
	IsArchived            bool      `json:"is_archived"`
	TemplateId            string    `json:"template_id"`
	PrefixId              int64     `json:"prefix_id"`
	CallTypeForReceiver   string    `json:"call_type_for_receiver"`
	SenderAvatar          string    `json:"sender_avatar"`
	ReceiverAvatar        string    `json:"receiver_avatar"`
	OperatorB             string    `json:"operator_b"`
	CountryB              string    `json:"country_b"`
	SenderCountryCode     string    `json:"sender_country_code"`
	ReceiverCountryCode   string    `json:"receiver_country_code"`
	MessageId             string    `json:"message_id"`
	SenderCurrency        string    `json:"sender_currency"`
	SenderCurrencyPrice   float64   `json:"sender_currency_price"`
	ReceiverCurrency      string    `json:"receiver_currency"`
	ReceiverCurrencyPrice float64   `json:"receiver_currency_price"`
	SendersCount          int64     `json:"senders_count"`
	CallTypeForCaller     string    `json:"call_type_for_caller"`
	SenderIccId           string    `json:"sender_icc_id"`
	ReceiverIccId         string    `json:"receiver_icc_id"`
	IsFraud               bool      `json:"is_fraud"`
	OttCodeLength         int64     `json:"ott_code_length"`
	OttPackageName        string    `json:"ott_package_name"`
	OttMask               string    `json:"ott_mask"`
	OttName               string    `json:"ott_name"`
	IsOtt                 bool      `json:"is_ott"`
	SenderName            string    `json:"sender_name"`
	ReceiverName          string    `json:"receiver_name"`
	OttRegExSmsCode       string    `json:"ott_reg_ex_sms_code"`
	MinBalance            float64   `json:"min_balance"`
	MinBalanceCurrency    string    `json:"min_balance_currency"`
	IsPaired              bool      `json:"is_paired"`
	AuthType              string    `json:"auth_type"`
	Instruction           string    `json:"instruction"`
	SecretCodeString      string    `json:"secret_code_string"`
	ReceiverImsi          string    `json:"receiver_imsi"`
	ParentStatus          string    `json:"parent_status"`
	ParentCdrSenderId     string    `json:"parent_cdr_sender_id"`
	NetA                  string    `json:"netA"`
	NetB                  string    `json:"netB"`
}

func (t *Sms) TableName() string {
	return "pingocean.SmsTasks"
}

func (t *Sms) DateTimeFormat() string {
	return "2006-01-02 15:04:05.000"
}

func (t *Sms) GetInsertSqlString() string {
	ds := goqu.Insert(t.TableName()).
		Cols(
			"id",
			"status",
			"scheme",
			"senderNumber",
			"senderUserId",
			"senderPrice",
			"senderOperatorId",
			"receiverNumber",
			"receiverUserId",
			"receiverPrice",
			"receiverOperatorId",
			"createdAt",
			"updatedAt",
			"clientCompanyId",
			"clientTaskId",
			"clientTaskName",
			"parentId",
			"smppServerId",
			"text",
			"country",
			"operator",
			"secretCode",
			"senderNumberId",
			"receiverNumberId",
			"timeout",
			"analyzing",
			"smppAvatar",
			"isDeleted",
			"isArchived",
			"templateId",
			"prefixId",
			"callTypeForReceiver",
			"senderAvatar",
			"receiverAvatar",
			"operatorB",
			"countryB",
			"senderCountryCode",
			"receiverCountryCode",
			"messageId",
			"senderCurrency",
			"senderCurrencyPrice",
			"receiverCurrency",
			"receiverCurrencyPrice",
			"sendersCount",
			"callTypeForCaller",
			"senderIccId",
			"receiverIccId",
			"isFraud",
			"ottCodeLength",
			"ottPackageName",
			"ottMask",
			"ottName",
			"isOtt",
			"senderName",
			"receiverName",
			"ottRegExSmsCode",
			"minBalance",
			"minBalanceCurrency",
			"isPaired",
			"authType",
			"instruction",
			"secretCodeString",
			"receiverImsi",
			"parentStatus",
			"parentCdrSenderId",
			"netA",
			"netB",
		).
		Vals(goqu.Vals{
			t.ID,
			t.Status,
			t.Scheme,
			t.SenderNumber,
			t.SenderUserId,
			t.SenderPrice,
			t.SenderOperatorId,
			t.ReceiverNumber,
			t.ReceiverUserId,
			t.ReceiverPrice,
			t.ReceiverOperatorId,
			t.CreatedAt.Format(t.DateTimeFormat()),
			t.UpdatedAt.Format(t.DateTimeFormat()),
			t.ClientCompanyId,
			t.ClientTaskId,
			t.ClientTaskName,
			t.ParentId,
			t.SmppServerId,
			t.Text,
			t.Country,
			t.Operator,
			t.SecretCode,
			t.SenderNumberId,
			t.ReceiverNumberId,
			t.Timeout,
			t.Analyzing,
			t.SmppAvatar,
			t.IsDeleted,
			t.IsArchived,
			t.TemplateId,
			t.PrefixId,
			t.CallTypeForReceiver,
			t.SenderAvatar,
			t.ReceiverAvatar,
			t.OperatorB,
			t.CountryB,
			t.SenderCountryCode,
			t.ReceiverCountryCode,
			t.MessageId,
			t.SenderCurrency,
			t.SenderCurrencyPrice,
			t.ReceiverCurrency,
			t.ReceiverCurrencyPrice,
			t.SendersCount,
			t.CallTypeForCaller,
			t.SenderIccId,
			t.ReceiverIccId,
			t.IsFraud,
			t.OttCodeLength,
			t.OttPackageName,
			t.OttMask,
			t.OttName,
			t.IsOtt,
			t.SenderName,
			t.ReceiverName,
			t.OttRegExSmsCode,
			t.MinBalance,
			t.MinBalanceCurrency,
			t.IsPaired,
			t.AuthType,
			t.Instruction,
			t.SecretCodeString,
			t.ReceiverImsi,
			t.ParentStatus,
			t.ParentCdrSenderId,
			t.NetA,
			t.NetB,
		})

	sql, _, _ := ds.ToSQL()
	return sql
}
