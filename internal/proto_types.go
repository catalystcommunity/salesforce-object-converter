package internal

const ProtoAddress = "SalesforceAddress"
const ProtoAny = "google.protobuf.Any"
const ProtoBase64 = "google.protobuf.Any"
const ProtoBoolean = "bool"
const ProtoCombobox = "string"
const ProtoCurrency = "double"
const ProtoDataCategoryGroupReference = "string"
const ProtoDate = "string"
const ProtoDatetime = "string"
const ProtoDouble = "double"
const ProtoEmail = "string"
const ProtoEncryptedString = "string"
const ProtoId = "string"
const ProtoInt = "int64"
const ProtoLocation = "SalesforceGeolocation"
const ProtoMultiPickList = "string"
const ProtoPercent = "double"
const ProtoPhone = "string"
const ProtoPicklist = "string"
const ProtoReference = "string"
const ProtoString = "string"
const ProtoTextArea = "string"
const ProtoTime = "string"
const ProtoUrl = "string"

var SfProtoTypeMap = map[string]string{
	SfAddress:                    ProtoAddress,
	SfAny:                        ProtoAny,
	SfBase64:                     ProtoBase64,
	SfBoolean:                    ProtoBoolean,
	SfCombobox:                   ProtoCombobox,
	SfCurrency:                   ProtoCurrency,
	SfDataCategoryGroupReference: ProtoDataCategoryGroupReference,
	SfDate:                       ProtoDate,
	SfDatetime:                   ProtoDatetime,
	SfDouble:                     ProtoDouble,
	SfEmail:                      ProtoEmail,
	SfEncryptedString:            ProtoEncryptedString,
	SfId:                         ProtoId,
	SfInt:                        ProtoInt,
	SfLocation:                   ProtoLocation,
	SfMultiPickList:              ProtoMultiPickList,
	SfPercent:                    ProtoPercent,
	SfPhone:                      ProtoPhone,
	SfPicklist:                   ProtoPicklist,
	SfReference:                  ProtoReference,
	SfString:                     ProtoString,
	SfTextArea:                   ProtoTextArea,
	SfTime:                       ProtoTime,
	SfUrl:                        ProtoUrl,
}
